package com.gameengine.affiliate.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "affiliates")
public class Affiliate {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false, unique = true)
    private String affiliateCode;
    
    @Column(nullable = false)
    private String name;
    
    @Column(nullable = false, unique = true)
    private String email;
    
    private String phone;
    
    @Column(nullable = false)
    private String status; // ACTIVE, INACTIVE, SUSPENDED
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_affiliate_id")
    private Affiliate parentAffiliate;
    
    @Column(nullable = false)
    private String affiliateTier; // BRONZE, SILVER, GOLD, PLATINUM, DIAMOND
    
    @Column(nullable = false)
    private BigDecimal revenueSharePercentage;
    
    private BigDecimal cpaAmount; // Cost Per Acquisition
    
    private Integer cpaThreshold; // Number of players needed for CPA
    
    private String paymentMethod;
    
    private String paymentDetails;
    
    @Column(nullable = false)
    private LocalDateTime createdAt;
    
    private LocalDateTime updatedAt;
    
    @Column(nullable = false)
    private Long merchantId;
    
    // Statistics (denormalized for performance)
    private Integer totalClicks;
    private Integer totalRegistrations;
    private Integer totalDepositors;
    private BigDecimal totalRevenue;
    private BigDecimal totalCommission;
    
    // Constructors
    public Affiliate() {
        this.createdAt = LocalDateTime.now();
        this.status = "ACTIVE";
        this.affiliateTier = "BRONZE";
        this.revenueSharePercentage = new BigDecimal("20.00");
        this.totalClicks = 0;
        this.totalRegistrations = 0;
        this.totalDepositors = 0;
        this.totalRevenue = BigDecimal.ZERO;
        this.totalCommission = BigDecimal.ZERO;
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public String getAffiliateCode() { return affiliateCode; }
    public void setAffiliateCode(String affiliateCode) { this.affiliateCode = affiliateCode; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public String getPhone() { return phone; }
    public void setPhone(String phone) { this.phone = phone; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public Affiliate getParentAffiliate() { return parentAffiliate; }
    public void setParentAffiliate(Affiliate parentAffiliate) { this.parentAffiliate = parentAffiliate; }
    
    public String getAffiliateTier() { return affiliateTier; }
    public void setAffiliateTier(String affiliateTier) { this.affiliateTier = affiliateTier; }
    
    public BigDecimal getRevenueSharePercentage() { return revenueSharePercentage; }
    public void setRevenueSharePercentage(BigDecimal revenueSharePercentage) { this.revenueSharePercentage = revenueSharePercentage; }
    
    public BigDecimal getCpaAmount() { return cpaAmount; }
    public void setCpaAmount(BigDecimal cpaAmount) { this.cpaAmount = cpaAmount; }
    
    public Integer getCpaThreshold() { return cpaThreshold; }
    public void setCpaThreshold(Integer cpaThreshold) { this.cpaThreshold = cpaThreshold; }
    
    public String getPaymentMethod() { return paymentMethod; }
    public void setPaymentMethod(String paymentMethod) { this.paymentMethod = paymentMethod; }
    
    public String getPaymentDetails() { return paymentDetails; }
    public void setPaymentDetails(String paymentDetails) { this.paymentDetails = paymentDetails; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
    
    public Long getMerchantId() { return merchantId; }
    public void setMerchantId(Long merchantId) { this.merchantId = merchantId; }
    
    public Integer getTotalClicks() { return totalClicks; }
    public void setTotalClicks(Integer totalClicks) { this.totalClicks = totalClicks; }
    
    public Integer getTotalRegistrations() { return totalRegistrations; }
    public void setTotalRegistrations(Integer totalRegistrations) { this.totalRegistrations = totalRegistrations; }
    
    public Integer getTotalDepositors() { return totalDepositors; }
    public void setTotalDepositors(Integer totalDepositors) { this.totalDepositors = totalDepositors; }
    
    public BigDecimal getTotalRevenue() { return totalRevenue; }
    public void setTotalRevenue(BigDecimal totalRevenue) { this.totalRevenue = totalRevenue; }
    
    public BigDecimal getTotalCommission() { return totalCommission; }
    public void setTotalCommission(BigDecimal totalCommission) { this.totalCommission = totalCommission; }
}
