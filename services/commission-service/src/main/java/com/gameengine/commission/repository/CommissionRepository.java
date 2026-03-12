package com.game-engine.commission.repository;

import com.game-engine.commission.model.Commission;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Repository
public interface CommissionRepository extends JpaRepository<Commission, Long> {
    
    List<Commission> findByAffiliateId(Long affiliateId);
    
    List<Commission> findByMerchantId(Long merchantId);
    
    List<Commission> findByAffiliateIdAndStatus(Long affiliateId, String status);
    
    List<Commission> findByPeriod(String period);
    
    @Query("SELECT c FROM Commission c WHERE c.affiliateId = :affiliateId AND c.period = :period")
    Optional<Commission> findByAffiliateAndPeriod(@Param("affiliateId") Long affiliateId, @Param("period") String period);
    
    @Query("SELECT SUM(c.commissionAmount) FROM Commission c WHERE c.affiliateId = :affiliateId AND c.status = 'PAID'")
    BigDecimal getTotalPaidCommission(@Param("affiliateId") Long affiliateId);
    
    @Query("SELECT SUM(c.commissionAmount) FROM Commission c WHERE c.affiliateId = :affiliateId AND c.status = 'PENDING'")
    BigDecimal getTotalPendingCommission(@Param("affiliateId") Long affiliateId);
    
    @Query("SELECT SUM(c.netRevenue) FROM Commission c WHERE c.merchantId = :merchantId")
    BigDecimal getTotalRevenueByMerchant(@Param("merchantId") Long merchantId);
    
    @Query("SELECT c FROM Commission c WHERE c.status = 'PENDING' AND c.calculatedAt <= :cutoffDate")
    List<Commission> findPendingOlderThan(@Param("cutoffDate") LocalDateTime cutoffDate);
}
