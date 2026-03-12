package com.game_engine.affiliate.service;

import com.game_engine.affiliate.model.Affiliate;
import com.game_engine.affiliate.model.Referral;
import com.game_engine.affiliate.repository.AffiliateRepository;
import com.game_engine.affiliate.repository.ReferralRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class AffiliateService {
    
    @Autowired
    private AffiliateRepository affiliateRepository;
    
    @Autowired
    private ReferralRepository referralRepository;
    
    @Transactional
    public Affiliate registerAffiliate(String name, String email, String phone, Long merchantId) {
        // Check if email already exists
        if (affiliateRepository.findByEmail(email).isPresent()) {
            throw new RuntimeException("Email already registered");
        }
        
        Affiliate affiliate = new Affiliate();
        affiliate.setName(name);
        affiliate.setEmail(email);
        affiliate.setPhone(phone);
        affiliate.setMerchantId(merchantId);
        affiliate.setAffiliateCode(generateAffiliateCode());
        
        return affiliateRepository.save(affiliate);
    }
    
    public Optional<Affiliate> getAffiliateByCode(String affiliateCode) {
        return affiliateRepository.findByAffiliateCode(affiliateCode);
    }
    
    public Optional<Affiliate> getAffiliateById(Long id) {
        return affiliateRepository.findById(id);
    }
    
    public List<Affiliate> getAffiliatesByMerchant(Long merchantId) {
        return affiliateRepository.findByMerchantId(merchantId);
    }
    
    public List<Affiliate> getActiveAffiliates(Long merchantId) {
        return affiliateRepository.findActiveByMerchantId(merchantId);
    }
    
    @Transactional
    public Affiliate updateAffiliateTier(Long affiliateId, String tier) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));
        
        affiliate.setAffiliateTier(tier);
        affiliate.setUpdatedAt(LocalDateTime.now());
        
        // Update commission rates based on tier
        BigDecimal revenueShare = getRevenueShareForTier(tier);
        affiliate.setRevenueSharePercentage(revenueShare);
        
        return affiliateRepository.save(affiliate);
    }
    
    @Transactional
    public Affiliate updateAffiliateStatus(Long affiliateId, String status) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));
        
        affiliate.setStatus(status);
        affiliate.setUpdatedAt(LocalDateTime.now());
        
        return affiliateRepository.save(affiliate);
    }
    
    // Referral tracking
    @Transactional
    public Referral trackClick(String affiliateCode, String ipAddress, String userAgent, String campaignId) {
        Affiliate affiliate = affiliateRepository.findByAffiliateCode(affiliateCode)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));
        
        Referral referral = new Referral();
        referral.setAffiliate(affiliate);
        referral.setReferralCode(generateReferralCode());
        referral.setSource("CLICK");
        referral.setIpAddress(ipAddress);
        referral.setUserAgent(userAgent);
        referral.setCampaignId(campaignId);
        referral.setClickedAt(LocalDateTime.now());
        referral.setStatus("CLICKED");
        
        // Update click count
        affiliate.setTotalClicks(affiliate.getTotalClicks() + 1);
        affiliateRepository.save(affiliate);
        
        return referralRepository.save(referral);
    }
    
    @Transactional
    public Referral trackRegistration(String referralCode, String ipAddress) {
        Referral referral = referralRepository.findByReferralCode(referralCode)
                .orElseThrow(() -> new RuntimeException("Referral not found"));
        
        referral.setStatus("REGISTERED");
        referral.setRegisteredAt(LocalDateTime.now());
        referral.setIpAddress(ipAddress);
        
        // Update registration count
        Affiliate affiliate = referral.getAffiliate();
        affiliate.setTotalRegistrations(affiliate.getTotalRegistrations() + 1);
        affiliateRepository.save(affiliate);
        
        return referralRepository.save(referral);
    }
    
    @Transactional
    public Referral trackFirstDeposit(String referralCode, BigDecimal depositAmount) {
        Referral referral = referralRepository.findByReferralCode(referralCode)
                .orElseThrow(() -> new RuntimeException("Referral not found"));
        
        referral.setStatus("DEPOSITED");
        referral.setFirstDepositAt(LocalDateTime.now());
        referral.setFirstDepositAmount(depositAmount);
        
        // Update depositor count
        Affiliate affiliate = referral.getAffiliate();
        affiliate.setTotalDepositors(affiliate.getTotalDepositors() + 1);
        affiliateRepository.save(affiliate);
        
        return referralRepository.save(referral);
    }
    
    public List<Referral> getReferralsByAffiliate(Long affiliateId) {
        return referralRepository.findByAffiliateId(affiliateId);
    }
    
    public List<Referral> getReferralsByCampaign(String campaignId) {
        return referralRepository.findByCampaignId(campaignId);
    }
    
    // Commission calculation
    @Transactional
    public BigDecimal calculateCommission(Long affiliateId, BigDecimal revenue) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));
        
        // Revenue share commission
        BigDecimal revenueShare = revenue.multiply(affiliate.getRevenueSharePercentage())
                .divide(new BigDecimal("100"));
        
        // Update total revenue and commission
        affiliate.setTotalRevenue(affiliate.getTotalRevenue().add(revenue));
        affiliate.setTotalCommission(affiliate.getTotalCommission().add(revenueShare));
        affiliateRepository.save(affiliate);
        
        return revenueShare;
    }
    
    // Sub-affiliate management
    @Transactional
    public Affiliate addSubAffiliate(Long parentAffiliateId, String name, String email, String phone) {
        Affiliate parent = affiliateRepository.findById(parentAffiliateId)
                .orElseThrow(() -> new RuntimeException("Parent affiliate not found"));
        
        Affiliate subAffiliate = new Affiliate();
        subAffiliate.setName(name);
        subAffiliate.setEmail(email);
        subAffiliate.setPhone(phone);
        subAffiliate.setParentAffiliate(parent);
        subAffiliate.setMerchantId(parent.getMerchantId());
        subAffiliate.setAffiliateCode(generateAffiliateCode());
        
        // Sub-affiliates start at lower tier
        subAffiliate.setAffiliateTier("BRONZE");
        subAffiliate.setRevenueSharePercentage(new BigDecimal("10.00"));
        
        return affiliateRepository.save(subAffiliate);
    }
    
    public List<Affiliate> getSubAffiliates(Long parentAffiliateId) {
        return affiliateRepository.findByParentAffiliateId(parentAffiliateId);
    }
    
    // Reporting
    public Long getTotalClicks(Long affiliateId) {
        return referralRepository.countRegistrations(affiliateId);
    }
    
    public Long getTotalRegistrations(Long affiliateId) {
        return referralRepository.countRegistrations(affiliateId);
    }
    
    public Long getTotalDepositors(Long affiliateId) {
        return referralRepository.countDepositors(affiliateId);
    }
    
    public BigDecimal getTotalRevenue(Long affiliateId) {
        return affiliateRepository.findById(affiliateId)
                .map(Affiliate::getTotalRevenue)
                .orElse(BigDecimal.ZERO);
    }
    
    public BigDecimal getTotalCommission(Long affiliateId) {
        return affiliateRepository.findById(affiliateId)
                .map(Affiliate::getTotalCommission)
                .orElse(BigDecimal.ZERO);
    }
    
    // Helper methods
    private String generateAffiliateCode() {
        return "AFF-" + UUID.randomUUID().toString().substring(0, 8).toUpperCase();
    }
    
    private String generateReferralCode() {
        return "REF-" + UUID.randomUUID().toString().substring(0, 12).toUpperCase();
    }
    
    private BigDecimal getRevenueShareForTier(String tier) {
        return switch (tier) {
            case "DIAMOND" -> new BigDecimal("40.00");
            case "PLATINUM" -> new BigDecimal("35.00");
            case "GOLD" -> new BigDecimal("30.00");
            case "SILVER" -> new BigDecimal("25.00");
            default -> new BigDecimal("20.00");
        };
    }
}
