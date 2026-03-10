package com.gameengine.commission.controller;

import com.gameengine.commission.model.Commission;
import com.gameengine.commission.service.CommissionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/v1/commissions")
public class CommissionController {
    
    @Autowired
    private CommissionService commissionService;
    
    @PostMapping
    public ResponseEntity<Commission> createCommission(@RequestBody Commission commission) {
        Commission created = commissionService.createCommission(commission);
        return ResponseEntity.status(HttpStatus.CREATED).body(created);
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<Commission> getCommissionById(@PathVariable Long id) {
        return commissionService.getCommissionById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @GetMapping("/affiliate/{affiliateId}")
    public ResponseEntity<List<Commission>> getCommissionsByAffiliate(@PathVariable Long affiliateId) {
        List<Commission> commissions = commissionService.getCommissionsByAffiliate(affiliateId);
        return ResponseEntity.ok(commissions);
    }
    
    @GetMapping("/merchant/{merchantId}")
    public ResponseEntity<List<Commission>> getCommissionsByMerchant(@PathVariable Long merchantId) {
        List<Commission> commissions = commissionService.getCommissionsByMerchant(merchantId);
        return ResponseEntity.ok(commissions);
    }
    
    @GetMapping("/period/{period}")
    public ResponseEntity<List<Commission>> getCommissionsByPeriod(@PathVariable String period) {
        List<Commission> commissions = commissionService.getCommissionsByPeriod(period);
        return ResponseEntity.ok(commissions);
    }
    
    @GetMapping("/affiliate/{affiliateId}/period/{period}")
    public ResponseEntity<Commission> getCommissionByAffiliateAndPeriod(
            @PathVariable Long affiliateId, @PathVariable String period) {
        return commissionService.getCommissionByAffiliateAndPeriod(affiliateId, period)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @GetMapping("/affiliate/{affiliateId}/total-paid")
    public ResponseEntity<Map<String, BigDecimal>> getTotalPaidCommission(@PathVariable Long affiliateId) {
        BigDecimal total = commissionService.getTotalPaidCommission(affiliateId);
        return ResponseEntity.ok(Map.of("totalPaid", total));
    }
    
    @GetMapping("/affiliate/{affiliateId}/total-pending")
    public ResponseEntity<Map<String, BigDecimal>> getTotalPendingCommission(@PathVariable Long affiliateId) {
        BigDecimal total = commissionService.getTotalPendingCommission(affiliateId);
        return ResponseEntity.ok(Map.of("totalPending", total));
    }
    
    @GetMapping("/merchant/{merchantId}/total-revenue")
    public ResponseEntity<Map<String, BigDecimal>> getTotalRevenueByMerchant(@PathVariable Long merchantId) {
        BigDecimal total = commissionService.getTotalRevenueByMerchant(merchantId);
        return ResponseEntity.ok(Map.of("totalRevenue", total));
    }
    
    @PostMapping("/calculate-revenue-share")
    public ResponseEntity<Commission> calculateRevenueShare(@RequestBody Map<String, Object> request) {
        Long affiliateId = Long.valueOf(request.get("affiliateId").toString());
        Long merchantId = Long.valueOf(request.get("merchantId").toString());
        BigDecimal netRevenue = new BigDecimal(request.get("netRevenue").toString());
        BigDecimal commissionRate = new BigDecimal(request.get("commissionRate").toString());
        String period = request.get("period") != null ? request.get("period").toString() : commissionService.generateCurrentPeriod();
        
        Commission commission = commissionService.calculateRevenueShare(affiliateId, merchantId, netRevenue, commissionRate, period);
        return ResponseEntity.ok(commission);
    }
    
    @PostMapping("/calculate-cpa")
    public ResponseEntity<Commission> calculateCPA(@RequestBody Map<String, Object> request) {
        Long affiliateId = Long.valueOf(request.get("affiliateId").toString());
        Long merchantId = Long.valueOf(request.get("merchantId").toString());
        int newPlayers = Integer.parseInt(request.get("newPlayers").toString());
        BigDecimal cpaRate = new BigDecimal(request.get("cpaRate").toString());
        String period = request.get("period") != null ? request.get("period").toString() : commissionService.generateCurrentPeriod();
        
        Commission commission = commissionService.calculateCPA(affiliateId, merchantId, newPlayers, cpaRate, period);
        return ResponseEntity.ok(commission);
    }
    
    @PostMapping("/{id}/approve")
    public ResponseEntity<Commission> approveCommission(@PathVariable Long id) {
        Commission commission = commissionService.approveCommission(id);
        return ResponseEntity.ok(commission);
    }
    
    @PostMapping("/{id}/reject")
    public ResponseEntity<Commission> rejectCommission(@PathVariable Long id, @RequestBody Map<String, String> request) {
        String reason = request.get("reason");
        Commission commission = commissionService.rejectCommission(id, reason);
        return ResponseEntity.ok(commission);
    }
    
    @PostMapping("/{id}/pay")
    public ResponseEntity<Commission> payCommission(@PathVariable Long id) {
        Commission commission = commissionService.payCommission(id);
        return ResponseEntity.ok(commission);
    }
    
    @GetMapping("/pending")
    public ResponseEntity<List<Commission>> getPendingCommissions() {
        List<Commission> commissions = commissionService.getPendingCommissions();
        return ResponseEntity.ok(commissions);
    }
    
    @GetMapping
    public ResponseEntity<List<Commission>> getAllCommissions() {
        List<Commission> commissions = commissionService.getAllCommissions();
        return ResponseEntity.ok(commissions);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteCommission(@PathVariable Long id) {
        commissionService.deleteCommission(id);
        return ResponseEntity.noContent().build();
    }
}
