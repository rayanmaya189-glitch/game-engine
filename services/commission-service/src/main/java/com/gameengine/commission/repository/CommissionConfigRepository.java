package com.gameengine.commission.repository;

import com.gameengine.commission.model.CommissionConfig;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CommissionConfigRepository extends JpaRepository<CommissionConfig, Long> {
    
    List<CommissionConfig> findByAffiliateId(Long affiliateId);
    
    List<CommissionConfig> findByMerchantId(Long merchantId);
    
    List<CommissionConfig> findByAffiliateIdAndIsActive(Long affiliateId, Boolean isActive);
    
    @Query("SELECT cc FROM CommissionConfig cc WHERE cc.affiliateId = :affiliateId AND cc.merchantId = :merchantId AND cc.isActive = true")
    List<CommissionConfig> findActiveByAffiliateAndMerchant(@Param("affiliateId") Long affiliateId, @Param("merchantId") Long merchantId);
    
    @Query("SELECT cc FROM CommissionConfig cc WHERE cc.affiliateId = :affiliateId AND cc.merchantId = :merchantId AND cc.commissionType = :type AND cc.isActive = true")
    Optional<CommissionConfig> findByAffiliateAndMerchantAndType(
            @Param("affiliateId") Long affiliateId, 
            @Param("merchantId") Long merchantId, 
            @Param("type") String type);
}
