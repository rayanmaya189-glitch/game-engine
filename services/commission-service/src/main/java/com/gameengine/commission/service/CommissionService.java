package com.gameengine.commission.service;

import com.gameengine.commission.model.Commission;
import com.gameengine.commission.repository.CommissionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.YearMonth;
import java.util.List;
import java.util.Optional;

@Service
@Transactional
public class CommissionService {
    
    @Autowired
    private CommissionRepository commissionRepository;
    
    public Commission createCommission(Commission commission) {
        commission.setCreatedAt(LocalDateTime.now());
        commission.setStatus("PENDING");
        return commissionRepository.save(commission);
    }
    
    public Optional<Commission> getCommissionById(Long id) {
        return commissionRepository.findById(id);
    }
    
    public List<Commission> getCommissionsByAffiliate(Long affiliateId) {
        return commissionRepository.findByAffiliateId(affiliateId);
    }
    
    public List<Commission> getCommissionsByMerchant(Long merchantId) {
        return commissionRepository.findByMerchantId(merchantId);
    }
    
    public List<Commission> getCommissionsByPeriod(String period) {
        return commissionRepository.findByPeriod(period);
    }
    
    public Optional<Commission> getCommissionByAffiliateAndPeriod(Long affiliateId, String period) {
        return commissionRepository.findByAffiliateAndPeriod(affiliateId, period);
    }
    
    public BigDecimal getTotalPaidCommission(Long affiliateId) {
        BigDecimal total = commissionRepository.getTotalPaidCommission(affiliateId);
        return total != null ? total : BigDecimal.ZERO;
    }
    
    public BigDecimal getTotalPendingCommission(Long affiliateId) {
        BigDecimal total = commissionRepository.getTotalPendingCommission(affiliateId);
        return total != null ? total : BigDecimal.ZERO;
    }
    
    public BigDecimal getTotalRevenueByMerchant(Long merchantId) {
        BigDecimal total = commissionRepository.getTotalRevenueByMerchant(merchantId);
        return total != null ? total : BigDecimal.ZERO;
    }
    
    public Commission calculateRevenueShare(Long affiliateId, Long merchantId, BigDecimal netRevenue, 
                                              BigDecimal commissionRate, String period) {
        Optional<Commission> existing = commissionRepository.findByAffiliateAndPeriod(affiliateId, period);
        
        Commission commission;
        if (existing.isPresent()) {
            commission = existing.get();
            commission.setNetRevenue(commission.getNetRevenue().add(netRevenue));
            commission.setCommissionAmount(commission.getNetRevenue().multiply(commissionRate));
            commission.setCalculatedAt(LocalDateTime.now());
        } else {
            commission = new Commission();
            commission.setAffiliateId(affiliateId);
            commission.setMerchantId(merchantId);
            commission.setPeriod(period);
            commission.setNetRevenue(netRevenue);
            commission.setCommissionRate(commissionRate);
            commission.setCommissionAmount(netRevenue.multiply(commissionRate));
            commission.setCpaAmount(BigDecimal.ZERO);
            commission.setTotalDeposits(BigDecimal.ZERO);
            commission.setTotalPlayers(0);
            commission.setNewPlayers(0);
        }
        
        return commissionRepository.save(commission);
    }
    
    public Commission calculateCPA(Long affiliateId, Long merchantId, int newPlayers, 
                                     BigDecimal cpaRate, String period) {
        Optional<Commission> existing = commissionRepository.findByAffiliateAndPeriod(affiliateId, period);
        
        Commission commission;
        if (existing.isPresent()) {
            commission = existing.get();
            commission.setNewPlayers(commission.getNewPlayers() + newPlayers);
            commission.setCpaAmount(BigDecimal.valueOf(commission.getNewPlayers()).multiply(cpaRate));
            commission.setCalculatedAt(LocalDateTime.now());
        } else {
            commission = new Commission();
            commission.setAffiliateId(affiliateId);
            commission.setMerchantId(merchantId);
            commission.setPeriod(period);
            commission.setNetRevenue(BigDecimal.ZERO);
            commission.setCommissionRate(BigDecimal.ZERO);
            commission.setCommissionAmount(BigDecimal.ZERO);
            commission.setCpaAmount(BigDecimal.valueOf(newPlayers).multiply(cpaRate));
            commission.setTotalDeposits(BigDecimal.ZERO);
            commission.setTotalPlayers(0);
            commission.setNewPlayers(newPlayers);
        }
        
        return commissionRepository.save(commission);
    }
    
    public Commission approveCommission(Long id) {
        Commission commission = commissionRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission not found"));
        commission.setStatus("APPROVED");
        commission.setApprovedAt(LocalDateTime.now());
        return commissionRepository.save(commission);
    }
    
    public Commission rejectCommission(Long id, String reason) {
        Commission commission = commissionRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission not found"));
        commission.setStatus("REJECTED");
        commission.setRejectedAt(LocalDateTime.now());
        commission.setRejectionReason(reason);
        return commissionRepository.save(commission);
    }
    
    public Commission payCommission(Long id) {
        Commission commission = commissionRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission not found"));
        
        if (!"APPROVED".equals(commission.getStatus())) {
            throw new RuntimeException("Commission must be approved before payment");
        }
        
        commission.setStatus("PAID");
        commission.setPaidAt(LocalDateTime.now());
        return commissionRepository.save(commission);
    }
    
    public List<Commission> getPendingCommissions() {
        LocalDateTime cutoffDate = LocalDateTime.now().minusDays(30);
        return commissionRepository.findPendingOlderThan(cutoffDate);
    }
    
    public String generatePeriod(LocalDate date) {
        YearMonth ym = YearMonth.from(date);
        return ym.toString();
    }
    
    public String generateCurrentPeriod() {
        return generatePeriod(LocalDate.now());
    }
    
    public List<Commission> getAllCommissions() {
        return commissionRepository.findAll();
    }
    
    public void deleteCommission(Long id) {
        commissionRepository.deleteById(id);
    }
}
