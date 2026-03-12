package com.game-engine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "insurance_claim")
public class InsuranceClaim {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(name = "user_id")
    private Long userId;
    
    @Column(name = "game_id")
    private Long gameId;
    
    @Column(name = "bet_id")
    private Long betId;
    
    @Column(name = "insurance_policy_id")
    private String insurancePolicyId;
    
    @Column(name = "claim_type")
    private String claimType; // GAME_LOSS, SYSTEM_FAILURE, TECHNNICAL_ERROR, FRAUD_PROTECTION
    
    @Column(name = "insured_amount", precision = 18, scale = 2)
    private BigDecimal insuredAmount;
    
    @Column(name = "loss_amount", precision = 18, scale = 2)
    private BigDecimal lossAmount;
    
    @Column(name = "claim_amount", precision = 18, scale = 2)
    private BigDecimal claimAmount;
    
    @Column(name = "status")
    private String status; // PENDING, REVIEWING, APPROVED, REJECTED, PAID, CANCELLED
    
    @Column(name = "claim_reason")
    private String claimReason;
    
    @Column(name = "evidence_details")
    private String evidenceDetails;
    
    @Column(name = "admin_note")
    private String adminNote;
    
    @Column(name = "reviewed_by")
    private Long reviewedBy;
    
    @Column(name = "requested_at")
    private LocalDateTime requestedAt;
    
    @Column(name = "reviewed_at")
    private LocalDateTime reviewedAt;
    
    @Column(name = "paid_at")
    private LocalDateTime paidAt;
    
    @Column(name = "transaction_id")
    private String transactionId;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    public InsuranceClaim() {
        this.status = "PENDING";
        this.createdAt = LocalDateTime.now();
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }
    
    public Long getGameId() { return gameId; }
    public void setGameId(Long gameId) { this.gameId = gameId; }
    
    public Long getBetId() { return betId; }
    public void setBetId(Long betId) { this.betId = betId; }
    
    public String getInsurancePolicyId() { return insurancePolicyId; }
    public void setInsurancePolicyId(String insurancePolicyId) { this.insurancePolicyId = insurancePolicyId; }
    
    public String getClaimType() { return claimType; }
    public void setClaimType(String claimType) { this.claimType = claimType; }
    
    public BigDecimal getInsuredAmount() { return insuredAmount; }
    public void setInsuredAmount(BigDecimal insuredAmount) { this.insuredAmount = insuredAmount; }
    
    public BigDecimal getLossAmount() { return lossAmount; }
    public void setLossAmount(BigDecimal lossAmount) { this.lossAmount = lossAmount; }
    
    public BigDecimal getClaimAmount() { return claimAmount; }
    public void setClaimAmount(BigDecimal claimAmount) { this.claimAmount = claimAmount; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public String getClaimReason() { return claimReason; }
    public void setClaimReason(String claimReason) { this.claimReason = claimReason; }
    
    public String getEvidenceDetails() { return evidenceDetails; }
    public void setEvidenceDetails(String evidenceDetails) { this.evidenceDetails = evidenceDetails; }
    
    public String getAdminNote() { return adminNote; }
    public void setAdminNote(String adminNote) { this.adminNote = adminNote; }
    
    public Long getReviewedBy() { return reviewedBy; }
    public void setReviewedBy(Long reviewedBy) { this.reviewedBy = reviewedBy; }
    
    public LocalDateTime getRequestedAt() { return requestedAt; }
    public void setRequestedAt(LocalDateTime requestedAt) { this.requestedAt = requestedAt; }
    
    public LocalDateTime getReviewedAt() { return reviewedAt; }
    public void setReviewedAt(LocalDateTime reviewedAt) { this.reviewedAt = reviewedAt; }
    
    public LocalDateTime getPaidAt() { return paidAt; }
    public void setPaidAt(LocalDateTime paidAt) { this.paidAt = paidAt; }
    
    public String getTransactionId() { return transactionId; }
    public void setTransactionId(String transactionId) { this.transactionId = transactionId; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
}
