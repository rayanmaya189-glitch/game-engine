package com.game-engine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "commissions")
public class Commission {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false)
    private Long affiliateId;
    
    @Column(nullable = false)
    private Long merchantId;
    
    @Column(nullable = false)
    private String commissionType; // REVENUE_SHARE, CPA, HYBRID
    
    @Column(nullable = false)
    private String status; // PENDING, APPROVED, PAID, CANCELLED
    
    @Column(nullable = false)
    private BigDecimal netRevenue;
    
    @Column(nullable = false)
    private BigDecimal commissionPercentage;
    
    @Column(nullable = false)
    private BigDecimal commissionAmount;
    
    private BigDecimal cpaAmount;
    
    private Integer cpaCount;
    
    @Column(nullable = false)
    private String period; // e.g., "2024-01" for January 2024
    
    @Column(nullable = false)
    private LocalDateTime calculatedAt;
    
    private LocalDateTime approvedAt;
    
    private LocalDateTime paidAt;
    
    private String notes;
    
    // Constructors
    public Commission() {
        this.calculatedAt = LocalDateTime.now();
        this.status = "PENDING";
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getAffiliateId() { return affiliateId; }
    public void setAffiliateId(Long affiliateId) { this.affiliateId = affiliateId; }
    
    public Long getMerchantId() { return merchantId; }
    public void setMerchantId(Long merchantId) { this.merchantId = merchantId; }
    
    public String getCommissionType() { return commissionType; }
    public void setCommissionType(String commissionType) { this.commissionType = commissionType; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public BigDecimal getNetRevenue() { return netRevenue; }
    public void setNetRevenue(BigDecimal netRevenue) { this.netRevenue = netRevenue; }
    
    public BigDecimal getCommissionPercentage() { return commissionPercentage; }
    public void setCommissionPercentage(BigDecimal commissionPercentage) { this.commissionPercentage = commissionPercentage; }
    
    public BigDecimal getCommissionAmount() { return commissionAmount; }
    public void setCommissionAmount(BigDecimal commissionAmount) { this.commissionAmount = commissionAmount; }
    
    public BigDecimal getCpaAmount() { return cpaAmount; }
    public void setCpaAmount(BigDecimal cpaAmount) { this.cpaAmount = cpaAmount; }
    
    public Integer getCpaCount() { return cpaCount; }
    public void setCpaCount(Integer cpaCount) { this.cpaCount = cpaCount; }
    
    public String getPeriod() { return period; }
    public void setPeriod(String period) { this.period = period; }
    
    public LocalDateTime getCalculatedAt() { return calculatedAt; }
    public void setCalculatedAt(LocalDateTime calculatedAt) { this.calculatedAt = calculatedAt; }
    
    public LocalDateTime getApprovedAt() { return approvedAt; }
    public void setApprovedAt(LocalDateTime approvedAt) { this.approvedAt = approvedAt; }
    
    public LocalDateTime getPaidAt() { return paidAt; }
    public void setPaidAt(LocalDateTime paidAt) { this.paidAt = paidAt; }
    
    public String getNotes() { return notes; }
    public void setNotes(String notes) { this.notes = notes; }
}
