package com.game-engine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "settlement")
public class Settlement {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(name = "user_id")
    private Long userId;
    
    @Column(name = "settlement_type")
    private String settlementType; // COMMISSION, REBET, INSURANCE, WINNINGS, REFUND, JACKPOT, TOURNAMENT
    
    @Column(name = "claim_id")
    private Long claimId;
    
    @Column(name = "reference_id")
    private String referenceId;
    
    @Column(name = "amount", precision = 18, scale = 2)
    private BigDecimal amount;
    
    @Column(name = "bonus_amount", precision = 18, scale = 2)
    private BigDecimal bonusAmount;
    
    @Column(name = "rake_amount", precision = 18, scale = 2)
    private BigDecimal rakeAmount;
    
    @Column(name = "net_amount", precision = 18, scale = 2)
    private BigDecimal netAmount;
    
    @Column(name = "currency")
    private String currency;
    
    @Column(name = "status")
    private String status; // PENDING, PROCESSING, COMPLETED, FAILED, CANCELLED
    
    @Column(name = "payment_method")
    private String paymentMethod; // WALLET, BANK, CRYPTO, CARD
    
    @Column(name = "transaction_id")
    private String transactionId;
    
    @Column(name = "external_transaction_id")
    private String externalTransactionId;
    
    @Column(name = "processed_by")
    private Long processedBy;
    
    @Column(name = "processed_at")
    private LocalDateTime processedAt;
    
    @Column(name = "completed_at")
    private LocalDateTime completedAt;
    
    @Column(name = "failure_reason")
    private String failureReason;
    
    @Column(name = "metadata")
    @Column(columnDefinition = "TEXT")
    private String metadata;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    public Settlement() {
        this.status = "PENDING";
        this.currency = "USD";
        this.createdAt = LocalDateTime.now();
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }
    
    public String getSettlementType() { return settlementType; }
    public void setSettlementType(String settlementType) { this.settlementType = settlementType; }
    
    public Long getClaimId() { return claimId; }
    public void setClaimId(Long claimId) { this.claimId = claimId; }
    
    public String getReferenceId() { return referenceId; }
    public void setReferenceId(String referenceId) { this.referenceId = referenceId; }
    
    public BigDecimal getAmount() { return amount; }
    public void setAmount(BigDecimal amount) { this.amount = amount; }
    
    public BigDecimal getBonusAmount() { return bonusAmount; }
    public void setBonusAmount(BigDecimal bonusAmount) { this.bonusAmount = bonusAmount; }
    
    public BigDecimal getRakeAmount() { return rakeAmount; }
    public void setRakeAmount(BigDecimal rakeAmount) { this.rakeAmount = rakeAmount; }
    
    public BigDecimal getNetAmount() { return netAmount; }
    public void setNetAmount(BigDecimal netAmount) { this.netAmount = netAmount; }
    
    public String getCurrency() { return currency; }
    public void setCurrency(String currency) { this.currency = currency; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public String getPaymentMethod() { return paymentMethod; }
    public void setPaymentMethod(String paymentMethod) { this.paymentMethod = paymentMethod; }
    
    public String getTransactionId() { return transactionId; }
    public void setTransactionId(String transactionId) { this.transactionId = transactionId; }
    
    public String getExternalTransactionId() { return externalTransactionId; }
    public void setExternalTransactionId(String externalTransactionId) { this.externalTransactionId = externalTransactionId; }
    
    public Long getProcessedBy() { return processedBy; }
    public void setProcessedBy(Long processedBy) { this.processedBy = processedBy; }
    
    public LocalDateTime getProcessedAt() { return processedAt; }
    public void setProcessedAt(LocalDateTime processedAt) { this.processedAt = processedAt; }
    
    public LocalDateTime getCompletedAt() { return completedAt; }
    public void setCompletedAt(LocalDateTime completedAt) { this.completedAt = completedAt; }
    
    public String getFailureReason() { return failureReason; }
    public void setFailureReason(String failureReason) { this.failureReason = failureReason; }
    
    public String getMetadata() { return metadata; }
    public void setMetadata(String metadata) { this.metadata = metadata; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
    
    public void calculateNetAmount() {
        BigDecimal total = amount.add(bonusAmount);
        this.netAmount = total.subtract(rakeAmount);
    }
}
