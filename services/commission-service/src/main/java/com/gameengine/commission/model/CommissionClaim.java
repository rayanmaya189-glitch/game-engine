package com.game-engine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "commission_claim")
public class CommissionClaim {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(name = "user_id")
    private Long userId;
    
    @Column(name = "affiliate_id")
    private Long affiliateId;
    
    @Column(name = "commission_id")
    private Long commissionId;
    
    @Column(name = "claim_type")
    private String claimType; // COMMISSION, REBET, INSURANCE
    
    @Column(name = "amount", precision = 18, scale = 2)
    private BigDecimal amount;
    
    @Column(name = "status")
    private String status; // PENDING, APPROVED, REJECTED, PAID, CANCELLED
    
    @Column(name = "claim_reason")
    private String claimReason;
    
    @Column(name = "admin_note")
    private String adminNote;
    
    @Column(name = "requested_at")
    private LocalDateTime requestedAt;
    
    @Column(name = "processed_at")
    private LocalDateTime processedAt;
    
    @Column(name = "paid_at")
    private LocalDateTime paidAt;
    
    @Column(name = "transaction_id")
    private String transactionId;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    public CommissionClaim() {
        this.status = "PENDING";
        this.createdAt = LocalDateTime.now();
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }
    
    public Long getAffiliateId() { return affiliateId; }
    public void setAffiliateId(Long affiliateId) { this.affiliateId = affiliateId; }
    
    public Long getCommissionId() { return commissionId; }
    public void setCommissionId(Long commissionId) { this.commissionId = commissionId; }
    
    public String getClaimType() { return claimType; }
    public void setClaimType(String claimType) { this.claimType = claimType; }
    
    public BigDecimal getAmount() { return amount; }
    public void setAmount(BigDecimal amount) { this.amount = amount; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public String getClaimReason() { return claimReason; }
    public void setClaimReason(String claimReason) { this.claimReason = claimReason; }
    
    public String getAdminNote() { return adminNote; }
    public void setAdminNote(String adminNote) { this.adminNote = adminNote; }
    
    public LocalDateTime getRequestedAt() { return requestedAt; }
    public void setRequestedAt(LocalDateTime requestedAt) { this.requestedAt = requestedAt; }
    
    public LocalDateTime getProcessedAt() { return processedAt; }
    public void setProcessedAt(LocalDateTime processedAt) { this.processedAt = processedAt; }
    
    public LocalDateTime getPaidAt() { return paidAt; }
    public void setPaidAt(LocalDateTime paidAt) { this.paidAt = paidAt; }
    
    public String getTransactionId() { return transactionId; }
    public void setTransactionId(String transactionId) { this.transactionId = transactionId; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
}
