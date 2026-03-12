package com.game_engine.affiliate.controller;

import com.game_engine.affiliate.model.Affiliate;
import com.game_engine.affiliate.model.Referral;
import com.game_engine.affiliate.service.AffiliateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/v1/affiliate")
public class AffiliateController {
    
    @Autowired
    private AffiliateService affiliateService;
    
    // Affiliate registration
    @PostMapping("/register")
    public ResponseEntity<Affiliate> registerAffiliate(@RequestBody Map<String, String> request) {
        String name = request.get("name");
        String email = request.get("email");
        String phone = request.get("phone");
        Long merchantId = Long.parseLong(request.get("merchantId"));
        
        Affiliate affiliate = affiliateService.registerAffiliate(name, email, phone, merchantId);
        return ResponseEntity.ok(affiliate);
    }
    
    // Get affiliate by code
    @GetMapping("/{affiliateCode}")
    public ResponseEntity<Affiliate> getAffiliateByCode(@PathVariable String affiliateCode) {
        return affiliateService.getAffiliateByCode(affiliateCode)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    // Get all affiliates for a merchant
    @GetMapping("/merchant/{merchantId}")
    public ResponseEntity<List<Affiliate>> getAffiliatesByMerchant(@PathVariable Long merchantId) {
        List<Affiliate> affiliates = affiliateService.getAffiliatesByMerchant(merchantId);
        return ResponseEntity.ok(affiliates);
    }
    
    // Get active affiliates
    @GetMapping("/merchant/{merchantId}/active")
    public ResponseEntity<List<Affiliate>> getActiveAffiliates(@PathVariable Long merchantId) {
        List<Affiliate> affiliates = affiliateService.getActiveAffiliates(merchantId);
        return ResponseEntity.ok(affiliates);
    }
    
    // Update affiliate tier
    @PutMapping("/{affiliateId}/tier")
    public ResponseEntity<Affiliate> updateAffiliateTier(
            @PathVariable Long affiliateId,
            @RequestBody Map<String, String> request) {
        String tier = request.get("tier");
        Affiliate affiliate = affiliateService.updateAffiliateTier(affiliateId, tier);
        return ResponseEntity.ok(affiliate);
    }
    
    // Update affiliate status
    @PutMapping("/{affiliateId}/status")
    public ResponseEntity<Affiliate> updateAffiliateStatus(
            @PathVariable Long affiliateId,
            @RequestBody Map<String, String> request) {
        String status = request.get("status");
        Affiliate affiliate = affiliateService.updateAffiliateStatus(affiliateId, status);
        return ResponseEntity.ok(affiliate);
    }
    
    // Track click
    @PostMapping("/track/click")
    public ResponseEntity<Referral> trackClick(@RequestBody Map<String, String> request) {
        String affiliateCode = request.get("affiliateCode");
        String ipAddress = request.get("ipAddress");
        String userAgent = request.get("userAgent");
        String campaignId = request.get("campaignId");
        
        Referral referral = affiliateService.trackClick(affiliateCode, ipAddress, userAgent, campaignId);
        return ResponseEntity.ok(referral);
    }
    
    // Track registration
    @PostMapping("/track/registration")
    public ResponseEntity<Referral> trackRegistration(@RequestBody Map<String, String> request) {
        String referralCode = request.get("referralCode");
        String ipAddress = request.get("ipAddress");
        
        Referral referral = affiliateService.trackRegistration(referralCode, ipAddress);
        return ResponseEntity.ok(referral);
    }
    
    // Track first deposit
    @PostMapping("/track/deposit")
    public ResponseEntity<Referral> trackFirstDeposit(
            @RequestBody Map<String, Object> request) {
        String referralCode = (String) request.get("referralCode");
        BigDecimal depositAmount = new BigDecimal(request.get("depositAmount").toString());
        
        Referral referral = affiliateService.trackFirstDeposit(referralCode, depositAmount);
        return ResponseEntity.ok(referral);
    }
    
    // Get referrals by affiliate
    @GetMapping("/{affiliateId}/referrals")
    public ResponseEntity<List<Referral>> getReferrals(@PathVariable Long affiliateId) {
        List<Referral> referrals = affiliateService.getReferralsByAffiliate(affiliateId);
        return ResponseEntity.ok(referrals);
    }
    
    // Get referrals by campaign
    @GetMapping("/campaign/{campaignId}/referrals")
    public ResponseEntity<List<Referral>> getCampaignReferrals(@PathVariable String campaignId) {
        List<Referral> referrals = affiliateService.getReferralsByCampaign(campaignId);
        return ResponseEntity.ok(referrals);
    }
    
    // Calculate commission
    @PostMapping("/{affiliateId}/commission")
    public ResponseEntity<Map<String, BigDecimal>> calculateCommission(
            @PathVariable Long affiliateId,
            @RequestBody Map<String, Object> request) {
        BigDecimal revenue = new BigDecimal(request.get("revenue").toString());
        BigDecimal commission = affiliateService.calculateCommission(affiliateId, revenue);
        
        Map<String, BigDecimal> response = new HashMap<>();
        response.put("revenue", revenue);
        response.put("commission", commission);
        
        return ResponseEntity.ok(response);
    }
    
    // Add sub-affiliate
    @PostMapping("/{parentAffiliateId}/sub-affiliate")
    public ResponseEntity<Affiliate> addSubAffiliate(
            @PathVariable Long parentAffiliateId,
            @RequestBody Map<String, String> request) {
        String name = request.get("name");
        String email = request.get("email");
        String phone = request.get("phone");
        
        Affiliate subAffiliate = affiliateService.addSubAffiliate(parentAffiliateId, name, email, phone);
        return ResponseEntity.ok(subAffiliate);
    }
    
    // Get sub-affiliates
    @GetMapping("/{parentAffiliateId}/sub-affiliates")
    public ResponseEntity<List<Affiliate>> getSubAffiliates(@PathVariable Long parentAffiliateId) {
        List<Affiliate> subAffiliates = affiliateService.getSubAffiliates(parentAffiliateId);
        return ResponseEntity.ok(subAffiliates);
    }
    
    // Get affiliate statistics
    @GetMapping("/{affiliateId}/stats")
    public ResponseEntity<Map<String, Object>> getAffiliateStats(@PathVariable Long affiliateId) {
        Map<String, Object> stats = new HashMap<>();
        stats.put("totalClicks", affiliateService.getTotalClicks(affiliateId));
        stats.put("totalRegistrations", affiliateService.getTotalRegistrations(affiliateId));
        stats.put("totalDepositors", affiliateService.getTotalDepositors(affiliateId));
        stats.put("totalRevenue", affiliateService.getTotalRevenue(affiliateId));
        stats.put("totalCommission", affiliateService.getTotalCommission(affiliateId));
        
        return ResponseEntity.ok(stats);
    }
    
    // Redirect to registration with affiliate code
    @GetMapping("/ref/{referralCode}")
    public ResponseEntity<Map<String, String>> redirectToRegistration(@PathVariable String referralCode) {
        Map<String, String> response = new HashMap<>();
        response.put("redirectUrl", "/register?ref=" + referralCode);
        return ResponseEntity.ok(response);
    }
}
