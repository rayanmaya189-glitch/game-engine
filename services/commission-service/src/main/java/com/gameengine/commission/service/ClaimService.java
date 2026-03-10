package com.gameengine.commission.service;

import com.gameengine.commission.model.*;
import com.gameengine.commission.repository.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@Transactional
public class ClaimService {
    
    @Autowired
    private CommissionClaimRepository commissionClaimRepository;
    
    @Autowired
    private RebetClaimRepository rebetClaimRepository;
    
    @Autowired
    private InsuranceClaimRepository insuranceClaimRepository;
    
    @Autowired
    private SettlementRepository settlementRepository;
    
    @Autowired
    private CommissionRepository commissionRepository;
    
    // Commission Claim Methods
    public CommissionClaim submitCommissionClaim(Long userId, Long affiliateId, Long commissionId, 
                                                   BigDecimal amount, String claimReason) {
        Commission claim = commissionRepository.findById(commissionId)
                .orElseThrow(() -> new RuntimeException("Commission not found"));
        
        if (claim.getStatus().equals("PAID")) {
            throw new RuntimeException("Commission already paid");
        }
        
        CommissionClaim commissionClaim = new CommissionClaim();
        commissionClaim.setUserId(userId);
        commissionClaim.setAffiliateId(affiliateId);
        commissionClaim.setCommissionId(commissionId);
        commissionClaim.setClaimType("COMMISSION");
        commissionClaim.setAmount(amount);
        commissionClaim.setClaimReason(claimReason);
        commissionClaim.setRequestedAt(LocalDateTime.now());
        commissionClaim.setStatus("PENDING");
        
        return commissionClaimRepository.save(commissionClaim);
    }
    
    public List<CommissionClaim> getUserCommissionClaims(Long userId) {
        return commissionClaimRepository.findByUserId(userId);
    }
    
    public List<CommissionClaim> getCommissionClaimsByStatus(String status) {
        return commissionClaimRepository.findByStatus(status);
    }
    
    public Optional<CommissionClaim> getCommissionClaimById(Long id) {
        return commissionClaimRepository.findById(id);
    }
    
    public BigDecimal getUserTotalPendingClaims(Long userId) {
        BigDecimal total = commissionClaimRepository.getTotalPendingAmount(userId);
        return total != null ? total : BigDecimal.ZERO;
    }
    
    // Rebet Claim Methods
    public RebetClaim createRebetClaim(Long userId, Long bonusId, String bonusCode, 
                                        BigDecimal bonusAmount, BigDecimal rebetRequirement,
                                        Long gameId, Long betId) {
        RebetClaim rebetClaim = new RebetClaim();
        rebetClaim.setUserId(userId);
        rebetClaim.setBonusId(bonusId);
        rebetClaim.setBonusCode(bonusCode);
        rebetClaim.setOriginalBonusAmount(bonusAmount);
        rebetClaim.setRebetRequirement(rebetRequirement);
        rebetClaim.setGameId(gameId);
        rebetClaim.setBetId(betId);
        rebetClaim.setStatus("IN_PROGRESS");
        rebetClaim.setExpiresAt(LocalDateTime.now().plusDays(30));
        
        return rebetClaimRepository.save(rebetClaim);
    }
    
    public RebetClaim updateRebetProgress(Long claimId, BigDecimal additionalBetAmount) {
        RebetClaim rebetClaim = rebetClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Rebet claim not found"));
        
        rebetClaim.setCurrentRebetAmount(rebetClaim.getCurrentRebetAmount().add(additionalBetAmount));
        rebetClaim.setUpdatedAt(LocalDateTime.now());
        
        if (rebetClaim.getCurrentRebetAmount().compareTo(rebetClaim.getRebetRequirement()) >= 0) {
            rebetClaim.setStatus("CLAIMABLE");
            rebetClaim.setClaimAmount(rebetClaim.getOriginalBonusAmount());
        }
        
        return rebetClaimRepository.save(rebetClaim);
    }
    
    public RebetClaim claimRebet(Long claimId) {
        RebetClaim rebetClaim = rebetClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Rebet claim not found"));
        
        if (!rebetClaim.isClaimable()) {
            throw new RuntimeException("Rebet requirement not met");
        }
        
        rebetClaim.setStatus("CLAIMED");
        rebetClaim.setClaimedAt(LocalDateTime.now());
        rebetClaim.setTransactionId(UUID.randomUUID().toString());
        
        // Create settlement
        createSettlement(rebetClaim.getUserId(), "REBET", claimId, rebetClaim.getClaimAmount(), 
                         BigDecimal.ZERO, BigDecimal.ZERO, "WALLET", rebetClaim.getTransactionId());
        
        return rebetClaimRepository.save(rebetClaim);
    }
    
    public List<RebetClaim> getUserRebetClaims(Long userId) {
        return rebetClaimRepository.findByUserId(userId);
    }
    
    public List<RebetClaim> getClaimableRebets(Long userId) {
        return rebetClaimRepository.findByUserIdAndStatus(userId, "CLAIMABLE");
    }
    
    // Insurance Claim Methods
    public InsuranceClaim submitInsuranceClaim(Long userId, Long gameId, Long betId,
                                                 String insurancePolicyId, String claimType,
                                                 BigDecimal insuredAmount, BigDecimal lossAmount,
                                                 String claimReason, String evidenceDetails) {
        InsuranceClaim insuranceClaim = new InsuranceClaim();
        insuranceClaim.setUserId(userId);
        insuranceClaim.setGameId(gameId);
        insuranceClaim.setBetId(betId);
        insuranceClaim.setInsurancePolicyId(insurancePolicyId);
        insuranceClaim.setClaimType(claimType);
        insuranceClaim.setInsuredAmount(insuredAmount);
        insuranceClaim.setLossAmount(lossAmount);
        insuranceClaim.setClaimAmount(insuredAmount.min(lossAmount));
        insuranceClaim.setClaimReason(claimReason);
        insuranceClaim.setEvidenceDetails(evidenceDetails);
        insuranceClaim.setRequestedAt(LocalDateTime.now());
        insuranceClaim.setStatus("PENDING");
        
        return insuranceClaimRepository.save(insuranceClaim);
    }
    
    public InsuranceClaim approveInsuranceClaim(Long claimId, Long reviewedBy, String adminNote) {
        InsuranceClaim insuranceClaim = insuranceClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Insurance claim not found"));
        
        insuranceClaim.setStatus("APPROVED");
        insuranceClaim.setReviewedBy(reviewedBy);
        insuranceClaim.setAdminNote(adminNote);
        insuranceClaim.setReviewedAt(LocalDateTime.now());
        
        return insuranceClaimRepository.save(insuranceClaim);
    }
    
    public InsuranceClaim rejectInsuranceClaim(Long claimId, Long reviewedBy, String adminNote) {
        InsuranceClaim insuranceClaim = insuranceClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Insurance claim not found"));
        
        insuranceClaim.setStatus("REJECTED");
        insuranceClaim.setReviewedBy(reviewedBy);
        insuranceClaim.setAdminNote(adminNote);
        insuranceClaim.setReviewedAt(LocalDateTime.now());
        
        return insuranceClaimRepository.save(insuranceClaim);
    }
    
    public InsuranceClaim payInsuranceClaim(Long claimId) {
        InsuranceClaim insuranceClaim = insuranceClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Insurance claim not found"));
        
        if (!"APPROVED".equals(insuranceClaim.getStatus())) {
            throw new RuntimeException("Insurance claim must be approved before payment");
        }
        
        insuranceClaim.setStatus("PAID");
        insuranceClaim.setPaidAt(LocalDateTime.now());
        insuranceClaim.setTransactionId(UUID.randomUUID().toString());
        
        // Create settlement
        createSettlement(insuranceClaim.getUserId(), "INSURANCE", claimId, insuranceClaim.getClaimAmount(),
                         BigDecimal.ZERO, BigDecimal.ZERO, "WALLET", insuranceClaim.getTransactionId());
        
        return insuranceClaimRepository.save(insuranceClaim);
    }
    
    public List<InsuranceClaim> getUserInsuranceClaims(Long userId) {
        return insuranceClaimRepository.findByUserId(userId);
    }
    
    public List<InsuranceClaim> getInsuranceClaimsByStatus(String status) {
        return insuranceClaimRepository.findByStatus(status);
    }
    
    // Settlement Methods
    public Settlement createSettlement(Long userId, String settlementType, Long claimId,
                                        BigDecimal amount, BigDecimal bonusAmount, BigDecimal rakeAmount,
                                        String paymentMethod, String transactionId) {
        Settlement settlement = new Settlement();
        settlement.setUserId(userId);
        settlement.setSettlementType(settlementType);
        settlement.setClaimId(claimId);
        settlement.setAmount(amount);
        settlement.setBonusAmount(bonusAmount);
        settlement.setRakeAmount(rakeAmount);
        settlement.calculateNetAmount();
        settlement.setPaymentMethod(paymentMethod);
        settlement.setTransactionId(transactionId);
        settlement.setStatus("COMPLETED");
        settlement.setCompletedAt(LocalDateTime.now());
        
        return settlementRepository.save(settlement);
    }
    
    public List<Settlement> getUserSettlements(Long userId) {
        return settlementRepository.findByUserId(userId);
    }
    
    public List<Settlement> getSettlementsByStatus(String status) {
        return settlementRepository.findByStatus(status);
    }
    
    public List<Settlement> getSettlementsByType(String type) {
        return settlementRepository.findBySettlementType(type);
    }
    
    public Optional<Settlement> getSettlementById(Long id) {
        return settlementRepository.findById(id);
    }
    
    public BigDecimal getUserTotalSettled(Long userId) {
        BigDecimal total = settlementRepository.getTotalSettledAmount(userId);
        return total != null ? total : BigDecimal.ZERO;
    }
    
    // Claim approval workflow
    public CommissionClaim approveCommissionClaim(Long claimId, String adminNote) {
        CommissionClaim claim = commissionClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Claim not found"));
        
        claim.setStatus("APPROVED");
        claim.setAdminNote(adminNote);
        claim.setProcessedAt(LocalDateTime.now());
        
        return commissionClaimRepository.save(claim);
    }
    
    public CommissionClaim rejectCommissionClaim(Long claimId, String adminNote) {
        CommissionClaim claim = commissionClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Claim not found"));
        
        claim.setStatus("REJECTED");
        claim.setAdminNote(adminNote);
        claim.setProcessedAt(LocalDateTime.now());
        
        return commissionClaimRepository.save(claim);
    }
    
    public CommissionClaim payCommissionClaim(Long claimId) {
        CommissionClaim claim = commissionClaimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Claim not found"));
        
        if (!"APPROVED".equals(claim.getStatus())) {
            throw new RuntimeException("Claim must be approved before payment");
        }
        
        claim.setStatus("PAID");
        claim.setPaidAt(LocalDateTime.now());
        claim.setTransactionId(UUID.randomUUID().toString());
        
        // Create settlement
        createSettlement(claim.getUserId(), "COMMISSION", claimId, claim.getAmount(),
                         BigDecimal.ZERO, BigDecimal.ZERO, "WALLET", claim.getTransactionId());
        
        return commissionClaimRepository.save(claim);
    }
}
