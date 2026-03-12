package com.game_engine.affiliate.repository;

import com.game_engine.affiliate.model.Affiliate;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.util.List;
import java.util.Optional;

@Repository
public interface AffiliateRepository extends JpaRepository<Affiliate, Long> {
    
    Optional<Affiliate> findByAffiliateCode(String affiliateCode);
    
    Optional<Affiliate> findByEmail(String email);
    
    List<Affiliate> findByMerchantId(Long merchantId);
    
    List<Affiliate> findByStatus(String status);
    
    List<Affiliate> findByParentAffiliateId(Long parentAffiliateId);
    
    @Query("SELECT a FROM Affiliate a WHERE a.merchantId = :merchantId AND a.status = 'ACTIVE'")
    List<Affiliate> findActiveByMerchantId(@Param("merchantId") Long merchantId);
    
    @Query("SELECT a FROM Affiliate a WHERE a.affiliateTier = :tier")
    List<Affiliate> findByTier(@Param("tier") String tier);
    
    @Query("SELECT SUM(a.totalRevenue) FROM Affiliate a WHERE a.merchantId = :merchantId")
    BigDecimal getTotalRevenueByMerchant(@Param("merchantId") Long merchantId);
    
    @Query("SELECT COUNT(a) FROM Affiliate a WHERE a.merchantId = :merchantId AND a.status = 'ACTIVE'")
    Long countActiveAffiliates(@Param("merchantId") Long merchantId);
    
    boolean existsByAffiliateCode(String affiliateCode);
    
    boolean existsByEmail(String email);
}
