package com.game-engine.affiliate.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "affiliate_referrals")
public class Referral {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "affiliate_id", nullable = false)
    private Affiliate affiliate;
    
    @Column(nullable = false, unique = true)
    private String referralCode;
    
    @Column(nullable = false)
    private String source; // CLICK, REGISTRATION, DEPOSIT
    
    private String ipAddress;
    
    private String userAgent;
    
    private String campaignId;
    
    private String campaignName;
    
    @Column(nullable = false)
    private LocalDateTime clickedAt;
    
    private LocalDateTime registeredAt;
    
    private LocalDateTime firstDepositAt;
    
    private BigDecimal firstDepositAmount;
    
    private String status; // CLICKED, REGISTERED, DEPOSITED
    
    // Constructors
    public Referral() {}
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Affiliate getAffiliate() { return affiliate; }
    public void setAffiliate(Affiliate affiliate) { this.affiliate = affiliate; }
    
    public String getReferralCode() { return referralCode; }
    public void setReferralCode(String referralCode) { this.referralCode = referralCode; }
    
    public String getSource() { return source; }
    public void setSource(String source) { this.source = source; }
    
    public String getIpAddress() { return ipAddress; }
    public void setIpAddress(String ipAddress) { this.ipAddress = ipAddress; }
    
    public String getUserAgent() { return userAgent; }
    public void setUserAgent(String userAgent) { this.userAgent = userAgent; }
    
    public String getCampaignId() { return campaignId; }
    public void setCampaignId(String campaignId) { this.campaignId = campaignId; }
    
    public String getCampaignName() { return campaignName; }
    public void setCampaignName(String campaignName) { this.campaignName = campaignName; }
    
    public LocalDateTime getClickedAt() { return clickedAt; }
    public void setClickedAt(LocalDateTime clickedAt) { this.clickedAt = clickedAt; }
    
    public LocalDateTime getRegisteredAt() { return registeredAt; }
    public void setRegisteredAt(LocalDateTime registeredAt) { this.registeredAt = registeredAt; }
    
    public LocalDateTime getFirstDepositAt() { return firstDepositAt; }
    public void setFirstDepositAt(LocalDateTime firstDepositAt) { this.firstDepositAt = firstDepositAt; }
    
    public BigDecimal getFirstDepositAmount() { return firstDepositAmount; }
    public void setFirstDepositAmount(BigDecimal firstDepositAmount) { this.firstDepositAmount = firstDepositAmount; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
}
