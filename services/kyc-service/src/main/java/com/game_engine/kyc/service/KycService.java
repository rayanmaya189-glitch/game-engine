package com.game_engine.kyc.service;

import com.game_engine.kyc.model.KycVerification;
import com.game_engine.kyc.model.KycVerification.*;
import com.game_engine.kyc.repository.KycRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.Optional;
import java.util.UUID;

/**
 * KYC Service
 * 
 * Manages the Know Your Customer verification process.
 * Supports verification levels from basic (email/phone) to full identity verification.
 */
@Service
@RequiredArgsConstructor
@Slf4j
public class KycService {

    private final KycRepository kycRepository;
    private final DocumentStorageService documentStorage;

    // KYC Triggers
    @Value("${kyc.trigger.withdrawal:2000}")
    private double withdrawalTriggerAmount;

    @Value("${kyc.trigger.deposit-cumulative:2000}")
    private double depositCumulativeTrigger;

    @Value("${kyc.trigger.deposit-high:10000}")
    private double highDepositTrigger;

    /**
     * Get KYC verification status for a user
     */
    public Optional<KycVerification> getVerification(UUID userId) {
        return kycRepository.findByUserId(userId);
    }

    /**
     * Get current verification level for a user
     */
    public VerificationLevel getVerificationLevel(UUID userId) {
        return kycRepository.findByUserId(userId)
                .map(KycVerification::getLevel)
                .orElse(VerificationLevel.LEVEL_0);
    }

    /**
     * Check if user meets required KYC level
     */
    public boolean hasRequiredLevel(UUID userId, VerificationLevel requiredLevel) {
        VerificationLevel currentLevel = getVerificationLevel(userId);
        return currentLevel.ordinal() >= requiredLevel.ordinal();
    }

    /**
     * Start Level 1 verification (Basic - email + phone)
     */
    @Transactional
    public KycVerification startLevel1Verification(UUID userId, String phoneNumber) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseGet(() -> createNewVerification(userId));

        verification.setPhoneNumber(phoneNumber);
        verification.setStatus(KycStatus.IN_PROGRESS);
        verification.setLevel(VerificationLevel.LEVEL_1);

        return kycRepository.save(verification);
    }

    /**
     * Complete phone verification (Level 1)
     */
    @Transactional
    public KycVerification verifyPhone(UUID userId, String verificationCode) {
        // Would verify the OTP code here
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalStateException("KYC verification not found"));

        verification.setPhoneVerified(true);
        verification.setPhoneVerifiedAt(LocalDateTime.now());
        updateLevelIfComplete(verification);

        return kycRepository.save(verification);
    }

    /**
     * Start Level 2 verification (Identity - document upload)
     */
    @Transactional
    public KycVerification startLevel2Verification(UUID userId, String firstName, String lastName,
                                                   LocalDateTime dateOfBirth, String country,
                                                   DocumentType documentType) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseGet(() -> createNewVerification(userId));

        verification.setFirstName(firstName);
        verification.setLastName(lastName);
        verification.setDateOfBirth(dateOfBirth);
        verification.setCountry(country);
        verification.setDocumentType(documentType);
        verification.setStatus(KycStatus.IN_PROGRESS);
        verification.setLevel(VerificationLevel.LEVEL_2);

        return kycRepository.save(verification);
    }

    /**
     * Submit document for verification
     */
    @Transactional
    public KycVerification submitDocument(UUID userId, byte[] documentFront, byte[] documentBack,
                                         byte[] selfie) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalStateException("KYC verification not found"));

        // Upload documents to S3
        String docFrontUrl = documentStorage.storeDocument(userId, "document_front", documentFront);
        String docBackUrl = documentStorage.storeDocument(userId, "document_back", documentBack);
        String selfieUrl = documentStorage.storeDocument(userId, "selfie", selfie);

        // Would call Sumsub/Onfido here for verification
        // For now, mark as pending review
        verification.setStatus(KycStatus.REQUIRES_REVIEW);

        return kycRepository.save(verification);
    }

    /**
     * Process webhook from KYC provider (Sumsub/Onfido)
     */
    @Transactional
    public KycVerification processProviderWebhook(UUID userId, String provider, String status,
                                                   Double faceMatchScore) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalStateException("KYC verification not found"));

        verification.setProvider(provider);
        
        switch (status.toLowerCase()) {
            case "completed", "verified", "success":
                verification.setStatus(KycStatus.VERIFIED);
                verification.setDocumentVerified(true);
                verification.setDocumentVerifiedAt(LocalDateTime.now());
                if (faceMatchScore != null) {
                    verification.setFaceVerified(true);
                    verification.setFaceMatchScore(faceMatchScore);
                }
                break;
            case "pending", "in_progress":
                verification.setStatus(KycStatus.IN_PROGRESS);
                break;
            case "failed", "rejected":
                verification.setStatus(KycStatus.REJECTED);
                break;
            default:
                verification.setStatus(KycStatus.REQUIRES_REVIEW);
        }

        updateLevelIfComplete(verification);
        return kycRepository.save(verification);
    }

    /**
     * Start Level 3 verification (Enhanced Due Diligence)
     */
    @Transactional
    public KycVerification startLevel3Verification(UUID userId) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalStateException("Level 2 verification required first"));

        if (verification.getLevel() != VerificationLevel.LEVEL_2) {
            throw new IllegalStateException("Must complete Level 2 verification first");
        }

        verification.setLevel(VerificationLevel.LEVEL_3);
        verification.setStatus(KycStatus.IN_PROGRESS);

        return kycRepository.save(verification);
    }

    /**
     * Submit proof of address document
     */
    @Transactional
    public KycVerification submitAddressProof(UUID userId, byte[] document) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalStateException("KYC verification not found"));

        String url = documentStorage.storeDocument(userId, "address_proof", document);
        verification.setAddressDocumentUrl(url);
        verification.setAddressVerified(false); // Awaiting review
        verification.setStatus(KycStatus.REQUIRES_REVIEW);

        return kycRepository.save(verification);
    }

    /**
     * Admin: Approve verification
     */
    @Transactional
    public KycVerification approve(UUID userId, UUID adminUserId, String notes) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalArgumentException("Verification not found"));

        verification.setStatus(KycStatus.VERIFIED);
        verification.setReviewedBy(adminUserId);
        verification.setReviewedAt(LocalDateTime.now());
        verification.setReviewNotes(notes);

        // Set appropriate verified flags based on level
        switch (verification.getLevel()) {
            case LEVEL_1:
                verification.setPhoneVerified(true);
                verification.setPhoneVerifiedAt(LocalDateTime.now());
                break;
            case LEVEL_2:
                verification.setDocumentVerified(true);
                verification.setDocumentVerifiedAt(LocalDateTime.now());
                verification.setFaceVerified(true);
                break;
            case LEVEL_3:
                verification.setAddressVerified(true);
                verification.setSourceOfFundsVerified(true);
                break;
        }

        return kycRepository.save(verification);
    }

    /**
     * Admin: Reject verification
     */
    @Transactional
    public KycVerification reject(UUID userId, UUID adminUserId, String reason) {
        KycVerification verification = kycRepository.findByUserId(userId)
                .orElseThrow(() -> new IllegalArgumentException("Verification not found"));

        verification.setStatus(KycStatus.REJECTED);
        verification.setReviewedBy(adminUserId);
        verification.setReviewedAt(LocalDateTime.now());
        verification.setRejectionReason(reason);

        return kycRepository.save(verification);
    }

    /**
     * Check if KYC is required based on triggers
     */
    public VerificationLevel getRequiredLevelForAction(double totalDeposits, double withdrawalAmount) {
        if (withdrawalAmount > 0) {
            return VerificationLevel.LEVEL_2;
        }
        if (totalDeposits >= highDepositTrigger) {
            return VerificationLevel.LEVEL_3;
        }
        if (totalDeposits >= depositCumulativeTrigger) {
            return VerificationLevel.LEVEL_2;
        }
        return VerificationLevel.LEVEL_0;
    }

    private KycVerification createNewVerification(UUID userId) {
        return KycVerification.builder()
                .userId(userId)
                .level(VerificationLevel.LEVEL_0)
                .status(KycStatus.PENDING)
                .build();
    }

    private void updateLevelIfComplete(KycVerification verification) {
        switch (verification.getLevel()) {
            case LEVEL_0:
                if (verification.getEmailVerified() != null && verification.getPhoneVerified() != null) {
                    if (verification.getEmailVerified() && verification.getPhoneVerified()) {
                        verification.setLevel(VerificationLevel.LEVEL_1);
                        verification.setStatus(KycStatus.VERIFIED);
                    }
                }
                break;
            case LEVEL_1:
                if (Boolean.TRUE.equals(verification.getDocumentVerified())) {
                    verification.setLevel(VerificationLevel.LEVEL_2);
                }
                break;
            case LEVEL_2:
                if (Boolean.TRUE.equals(verification.getAddressVerified()) &&
                    Boolean.TRUE.equals(verification.getSourceOfFundsVerified())) {
                    verification.setLevel(VerificationLevel.LEVEL_3);
                }
                break;
        }
    }
}
