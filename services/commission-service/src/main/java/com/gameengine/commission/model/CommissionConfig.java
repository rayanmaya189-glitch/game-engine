package com.gameengine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "commission_config")
public class CommissionConfig {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(name = "merchant_id")
    private Long merchantId;
    
    @Column(name = "affiliate_id")
    private Long affiliateId;
    
    @Column(name = "commission_type")
    private String commissionType; // REVENUE_SHARE, CPA, HYBRID
    
    @Column(name = "revenue_share_rate", precision = 5, scale = 4)
    private BigDecimal revenueShareRate;
    
    @Column(name = "cpa_rate", precision = 10, scale = 2)
    private BigDecimal cpaRate;
    
    @Column(name = "min_players")
    private Integer minPlayers;
    
    @Column(name = "tier_rate", precision = 5, scale = 4)
    private BigDecimal tierRate;
    
    @Column(name = "tier_threshold")
    private Integer tierThreshold;
    
    @Column(name = "is_active")
    private Boolean isActive;
    
    @Column(name = "effective_from")
    private LocalDateTime effectiveFrom;
    
    @Column(name = "effective_to")
    private LocalDateTime effectiveTo;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    public CommissionConfig() {
        this.isActive = true;
        this.createdAt = LocalDateTime.now();
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getMerchantId() { return merchantId; }
    public void setMerchantId(Long merchantId) { this.merchantId = merchantId; }
    
    public Long getAffiliateId() { return affiliateId; }
    public void setAffiliateId(Long affiliateId) { this.affiliateId = affiliateId; }
    
    public String getCommissionType() { return commissionType; }
    public void setCommissionType(String commissionType) { this.commissionType = commissionType; }
    
    public BigDecimal getRevenueShareRate() { return revenueShareRate; }
    public void setRevenueShareRate(BigDecimal revenueShareRate) { this.revenueShareRate = revenueShareRate; }
    
    public BigDecimal getCpaRate() { return cpaRate; }
    public void setCpaRate(BigDecimal cpaRate) { this.cpaRate = cpaRate; }
    
    public Integer getMinPlayers() { return minPlayers; }
    public void setMinPlayers(Integer minPlayers) { this.minPlayers = minPlayers; }
    
    public BigDecimal getTierRate() { return tierRate; }
    public void setTierRate(BigDecimal tierRate) { this.tierRate = tierRate; }
    
    public Integer getTierThreshold() { return tierThreshold; }
    public void setTierThreshold(Integer tierThreshold) { this.tierThreshold = tierThreshold; }
    
    public Boolean getIsActive() { return isActive; }
    public void setIsActive(Boolean isActive) { this.isActive = isActive; }
    
    public LocalDateTime getEffectiveFrom() { return effectiveFrom; }
    public void setEffectiveFrom(LocalDateTime effectiveFrom) { this.effectiveFrom = effectiveFrom; }
    
    public LocalDateTime getEffectiveTo() { return effectiveTo; }
    public void setEffectiveTo(LocalDateTime effectiveTo) { this.effectiveTo = effectiveTo; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
}
