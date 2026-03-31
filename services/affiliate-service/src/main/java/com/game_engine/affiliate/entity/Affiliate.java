package com.game_engine.affiliate.entity;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "affiliate_v2")
public class Affiliate {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "affiliate_code", nullable = false, unique = true)
    private String affiliateCode;

    @Column(name = "name", nullable = false)
    private String name;

    @Column(name = "email", nullable = false, unique = true)
    private String email;

    @Column(name = "status")
    private String status; // ACTIVE, INACTIVE, SUSPENDED

    @Column(name = "tier")
    private String tier; // BRONZE, SILVER, GOLD, PLATINUM, DIAMOND

    @Column(name = "commission_rate", precision = 5, scale = 4)
    private BigDecimal commissionRate;

    @Column(name = "total_referrals")
    private Integer totalReferrals;

    @Column(name = "total_revenue", precision = 18, scale = 2)
    private BigDecimal totalRevenue;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    public Affiliate() {
        this.status = "ACTIVE";
        this.tier = "BRONZE";
        this.commissionRate = new BigDecimal("0.2000");
        this.totalReferrals = 0;
        this.totalRevenue = BigDecimal.ZERO;
        this.createdAt = LocalDateTime.now();
    }

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public String getAffiliateCode() { return affiliateCode; }
    public void setAffiliateCode(String affiliateCode) { this.affiliateCode = affiliateCode; }

    public String getName() { return name; }
    public void setName(String name) { this.name = name; }

    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }

    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }

    public String getTier() { return tier; }
    public void setTier(String tier) { this.tier = tier; }

    public BigDecimal getCommissionRate() { return commissionRate; }
    public void setCommissionRate(BigDecimal commissionRate) { this.commissionRate = commissionRate; }

    public Integer getTotalReferrals() { return totalReferrals; }
    public void setTotalReferrals(Integer totalReferrals) { this.totalReferrals = totalReferrals; }

    public BigDecimal getTotalRevenue() { return totalRevenue; }
    public void setTotalRevenue(BigDecimal totalRevenue) { this.totalRevenue = totalRevenue; }

    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
}
