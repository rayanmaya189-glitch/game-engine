package com.game_engine.kyc.model;

import jakarta.persistence.*;
import lombok.*;
import java.time.LocalDateTime;
import java.util.UUID;

/**
 * KYC Verification Entity
 * 
 * Tracks the verification status of a user across different KYC levels.
 * 
 * KYC Levels:
 * - Level 0: Unverified (no verification attempted)
 * - Level 1: Basic (email + phone verified)
 * - Level 2: Identity (document verification passed)
 * - Level 3: Enhanced (address + source of funds verified)
 */
@Entity
@Table(name = "kyc_verifications", indexes = {
    @Index(name = "idx_kyc_user_id", columnList = "user_id"),
    @Index(name = "idx_kyc_status", columnList = "status"),
    @Index(name = "idx_kyc_level", columnList = "verification_level")
})
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode(of = "id")
public class KycVerification {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(name = "user_id", nullable = false, unique = true)
    private UUID userId;

    @Enumerated(EnumType.STRING)
    @Column(name = "verification_level", nullable = false)
    private VerificationLevel level;

    @Enumerated(EnumType.STRING)
    @Column(name = "status", nullable = false)
    private KycStatus status;

    // Level 1 - Basic Info
    @Column(name = "email_verified")
    private Boolean emailVerified;

    @Column(name = "email_verified_at")
    private LocalDateTime emailVerifiedAt;

    @Column(name = "phone_verified")
    private Boolean phoneVerified;

    @Column(name = "phone_verified_at")
    private LocalDateTime phoneVerifiedAt;

    @Column(name = "phone_number")
    private String phoneNumber;

    // Level 2 - Identity
    @Column(name = "first_name")
    private String firstName;

    @Column(name = "last_name")
    private String lastName;

    @Column(name = "date_of_birth")
    private LocalDateTime dateOfBirth;

    @Column(name = "country")
    private String country;

    @Column(name = "document_type")
    private DocumentType documentType;

    @Column(name = "document_id")
    private String documentId;

    @Column(name = "document_verified")
    private Boolean documentVerified;

    @Column(name = "document_verified_at")
    private LocalDateTime documentVerifiedAt;

    @Column(name = "face_verified")
    private Boolean faceVerified;

    @Column(name = "face_match_score")
    private Double faceMatchScore;

    // Level 3 - Enhanced Due Diligence
    @Column(name = "address_verified")
    private Boolean addressVerified;

    @Column(name = "address_document_url")
    private String addressDocumentUrl;

    @Column(name = "source_of_funds_verified")
    private Boolean sourceOfFundsVerified;

    @Column(name = "source_of_wealth_verified")
    private Boolean sourceOfWealthVerified;

    // Provider Reference
    @Column(name = "provider")
    private String provider; // sumsub, onfido

    @Column(name = "provider_verification_id")
    private String providerVerificationId;

    // Review
    @Column(name = "reviewed_by")
    private UUID reviewedBy;

    @Column(name = "reviewed_at")
    private LocalDateTime reviewedAt;

    @Column(name = "review_notes")
    private String reviewNotes;

    @Column(name = "rejection_reason")
    private String rejectionReason;

    // Timestamps
    @Column(name = "created_at", nullable = false)
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    @Column(name = "expires_at")
    private LocalDateTime expiresAt;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        if (status == null) {
            status = KycStatus.PENDING;
        }
        if (level == null) {
            level = VerificationLevel.LEVEL_0;
        }
    }

    @PreUpdate
    protected void onUpdate() {
        updatedAt = LocalDateTime.now();
    }

    public enum VerificationLevel {
        LEVEL_0,  // Unverified
        LEVEL_1,  // Basic (email + phone)
        LEVEL_2,  // Identity verified
        LEVEL_3   // Enhanced due diligence
    }

    public enum KycStatus {
        PENDING,           // Awaiting verification
        IN_PROGRESS,       // Verification in progress
        VERIFIED,          // Successfully verified
        REJECTED,          // Verification rejected
        EXPIRED,           // Verification expired
        REQUIRES_REVIEW    // Needs manual review
    }

    public enum DocumentType {
        PASSPORT,
        NATIONAL_ID,
        DRIVERS_LICENSE,
        RESIDENCE_PERMIT
    }
}
