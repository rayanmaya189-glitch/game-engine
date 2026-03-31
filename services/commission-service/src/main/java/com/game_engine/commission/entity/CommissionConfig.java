package com.game_engine.commission.entity;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "commission_config_v2")
public class CommissionConfig {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "agent_id")
    private Long agentId;

    @Column(name = "affiliate_id")
    private Long affiliateId;

    @Column(name = "commission_type")
    private String commissionType; // REVENUE_SHARE, CPA, HYBRID

    @Column(name = "rate", precision = 5, scale = 4)
    private BigDecimal rate;

    @Column(name = "min_deposit", precision = 18, scale = 2)
    private BigDecimal minDeposit;

    @Column(name = "max_commission", precision = 18, scale = 2)
    private BigDecimal maxCommission;

    @Column(name = "is_active")
    private Boolean isActive;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    public CommissionConfig() {
        this.isActive = true;
        this.createdAt = LocalDateTime.now();
    }

    @PreUpdate
    public void preUpdate() {
        this.updatedAt = LocalDateTime.now();
    }

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public Long getAgentId() { return agentId; }
    public void setAgentId(Long agentId) { this.agentId = agentId; }

    public Long getAffiliateId() { return affiliateId; }
    public void setAffiliateId(Long affiliateId) { this.affiliateId = affiliateId; }

    public String getCommissionType() { return commissionType; }
    public void setCommissionType(String commissionType) { this.commissionType = commissionType; }

    public BigDecimal getRate() { return rate; }
    public void setRate(BigDecimal rate) { this.rate = rate; }

    public BigDecimal getMinDeposit() { return minDeposit; }
    public void setMinDeposit(BigDecimal minDeposit) { this.minDeposit = minDeposit; }

    public BigDecimal getMaxCommission() { return maxCommission; }
    public void setMaxCommission(BigDecimal maxCommission) { this.maxCommission = maxCommission; }

    public Boolean getIsActive() { return isActive; }
    public void setIsActive(Boolean isActive) { this.isActive = isActive; }

    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }

    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
}
