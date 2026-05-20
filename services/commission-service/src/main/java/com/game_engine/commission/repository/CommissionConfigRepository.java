package com.game_engine.commission.repository;

import com.game_engine.commission.model.CommissionConfig;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository("commissionConfigRepositoryV2")
public interface CommissionConfigRepository extends JpaRepository<CommissionConfig, Long> {

    List<CommissionConfig> findByAgentId(Long agentId);

    List<CommissionConfig> findByAffiliateId(Long affiliateId);

    List<CommissionConfig> findByAgentIdAndIsActive(Long agentId, Boolean isActive);

    List<com.game_engine.commission.model.CommissionConfig> findByAffiliateIdAndIsActive(Long affiliateId,
            Boolean isActive);

    List<CommissionConfig> findByAgentIdAndAffiliateIdAndIsActive(Long agentId, Long affiliateId, Boolean isActive);

    @Query("SELECT c FROM CommissionConfig c WHERE c.merchantId = :merchantId")
    List<CommissionConfig> findByMerchantId(Long merchantId);

    @Query("SELECT c FROM CommissionConfig c WHERE c.affiliateId = :affiliateId AND c.merchantId = :merchantId AND c.isActive = true")
    List<CommissionConfig> findActiveByAffiliateAndMerchant(Long affiliateId, Long merchantId);

    @Query("SELECT c FROM CommissionConfig c WHERE c.affiliateId = :affiliateId AND c.merchantId = :merchantId AND c.type = :type")
    Optional<CommissionConfig> findByAffiliateAndMerchantAndType(Long affiliateId, Long merchantId, String type);
}
