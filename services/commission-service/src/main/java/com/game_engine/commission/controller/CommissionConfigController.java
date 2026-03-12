package com.game_engine.commission.controller;

import com.game_engine.commission.model.CommissionConfig;
import com.game_engine.commission.service.CommissionConfigService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/commission-configs")
public class CommissionConfigController {
    
    @Autowired
    private CommissionConfigService commissionConfigService;
    
    @PostMapping
    public ResponseEntity<CommissionConfig> createCommissionConfig(@RequestBody CommissionConfig config) {
        CommissionConfig created = commissionConfigService.createCommissionConfig(config);
        return ResponseEntity.status(HttpStatus.CREATED).body(created);
    }
    
    @GetMapping("/{id}")
    public ResponseEntity<CommissionConfig> getConfigById(@PathVariable Long id) {
        return commissionConfigService.getConfigById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @GetMapping("/affiliate/{affiliateId}")
    public ResponseEntity<List<CommissionConfig>> getConfigsByAffiliate(@PathVariable Long affiliateId) {
        List<CommissionConfig> configs = commissionConfigService.getConfigsByAffiliate(affiliateId);
        return ResponseEntity.ok(configs);
    }
    
    @GetMapping("/merchant/{merchantId}")
    public ResponseEntity<List<CommissionConfig>> getConfigsByMerchant(@PathVariable Long merchantId) {
        List<CommissionConfig> configs = commissionConfigService.getConfigsByMerchant(merchantId);
        return ResponseEntity.ok(configs);
    }
    
    @GetMapping("/affiliate/{affiliateId}/active")
    public ResponseEntity<List<CommissionConfig>> getActiveConfigsByAffiliate(@PathVariable Long affiliateId) {
        List<CommissionConfig> configs = commissionConfigService.getActiveConfigsByAffiliate(affiliateId);
        return ResponseEntity.ok(configs);
    }
    
    @GetMapping("/affiliate/{affiliateId}/merchant/{merchantId}/active")
    public ResponseEntity<List<CommissionConfig>> getActiveConfigsByAffiliateAndMerchant(
            @PathVariable Long affiliateId, @PathVariable Long merchantId) {
        List<CommissionConfig> configs = commissionConfigService.getActiveConfigsByAffiliateAndMerchant(affiliateId, merchantId);
        return ResponseEntity.ok(configs);
    }
    
    @GetMapping("/affiliate/{affiliateId}/merchant/{merchantId}/type/{type}")
    public ResponseEntity<CommissionConfig> getConfigByAffiliateAndMerchantAndType(
            @PathVariable Long affiliateId, @PathVariable Long merchantId, @PathVariable String type) {
        return commissionConfigService.getConfigByAffiliateAndMerchantAndType(affiliateId, merchantId, type)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    @PutMapping("/{id}")
    public ResponseEntity<CommissionConfig> updateCommissionConfig(
            @PathVariable Long id, @RequestBody CommissionConfig config) {
        CommissionConfig updated = commissionConfigService.updateCommissionConfig(id, config);
        return ResponseEntity.ok(updated);
    }
    
    @PostMapping("/{id}/activate")
    public ResponseEntity<CommissionConfig> activateConfig(@PathVariable Long id) {
        CommissionConfig config = commissionConfigService.activateConfig(id);
        return ResponseEntity.ok(config);
    }
    
    @PostMapping("/{id}/deactivate")
    public ResponseEntity<CommissionConfig> deactivateConfig(@PathVariable Long id) {
        CommissionConfig config = commissionConfigService.deactivateConfig(id);
        return ResponseEntity.ok(config);
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteConfig(@PathVariable Long id) {
        commissionConfigService.deleteConfig(id);
        return ResponseEntity.noContent().build();
    }
    
    @GetMapping
    public ResponseEntity<List<CommissionConfig>> getAllConfigs() {
        List<CommissionConfig> configs = commissionConfigService.getAllConfigs();
        return ResponseEntity.ok(configs);
    }
}
