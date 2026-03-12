package com.game-engine.commission.service;

import com.game-engine.commission.model.CommissionConfig;
import com.game-engine.commission.repository.CommissionConfigRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
@Transactional
public class CommissionConfigService {
    
    @Autowired
    private CommissionConfigRepository commissionConfigRepository;
    
    public CommissionConfig createCommissionConfig(CommissionConfig config) {
        config.setCreatedAt(LocalDateTime.now());
        return commissionConfigRepository.save(config);
    }
    
    public Optional<CommissionConfig> getConfigById(Long id) {
        return commissionConfigRepository.findById(id);
    }
    
    public List<CommissionConfig> getConfigsByAffiliate(Long affiliateId) {
        return commissionConfigRepository.findByAffiliateId(affiliateId);
    }
    
    public List<CommissionConfig> getConfigsByMerchant(Long merchantId) {
        return commissionConfigRepository.findByMerchantId(merchantId);
    }
    
    public List<CommissionConfig> getActiveConfigsByAffiliate(Long affiliateId) {
        return commissionConfigRepository.findByAffiliateIdAndIsActive(affiliateId, true);
    }
    
    public List<CommissionConfig> getActiveConfigsByAffiliateAndMerchant(Long affiliateId, Long merchantId) {
        return commissionConfigRepository.findActiveByAffiliateAndMerchant(affiliateId, merchantId);
    }
    
    public Optional<CommissionConfig> getConfigByAffiliateAndMerchantAndType(Long affiliateId, Long merchantId, String type) {
        return commissionConfigRepository.findByAffiliateAndMerchantAndType(affiliateId, merchantId, type);
    }
    
    public CommissionConfig updateCommissionConfig(Long id, CommissionConfig updatedConfig) {
        CommissionConfig existing = commissionConfigRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission config not found"));
        
        existing.setCommissionType(updatedConfig.getCommissionType());
        existing.setRevenueShareRate(updatedConfig.getRevenueShareRate());
        existing.setCpaRate(updatedConfig.getCpaRate());
        existing.setMinPlayers(updatedConfig.getMinPlayers());
        existing.setTierRate(updatedConfig.getTierRate());
        existing.setTierThreshold(updatedConfig.getTierThreshold());
        existing.setIsActive(updatedConfig.getIsActive());
        existing.setEffectiveFrom(updatedConfig.getEffectiveFrom());
        existing.setEffectiveTo(updatedConfig.getEffectiveTo());
        existing.setUpdatedAt(LocalDateTime.now());
        
        return commissionConfigRepository.save(existing);
    }
    
    public CommissionConfig activateConfig(Long id) {
        CommissionConfig config = commissionConfigRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission config not found"));
        
        config.setIsActive(true);
        config.setUpdatedAt(LocalDateTime.now());
        return commissionConfigRepository.save(config);
    }
    
    public CommissionConfig deactivateConfig(Long id) {
        CommissionConfig config = commissionConfigRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission config not found"));
        
        config.setIsActive(false);
        config.setUpdatedAt(LocalDateTime.now());
        return commissionConfigRepository.save(config);
    }
    
    public void deleteConfig(Long id) {
        commissionConfigRepository.deleteById(id);
    }
    
    public List<CommissionConfig> getAllConfigs() {
        return commissionConfigRepository.findAll();
    }
}
