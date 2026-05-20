package com.game_engine.commission.repository;

import com.game_engine.commission.entity.CommissionConfig;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository("commissionConfigRepositoryV2")
public interface CommissionConfigV2Repository extends JpaRepository<CommissionConfig, Long> {

    List<CommissionConfig> findByAgentIdAndAffiliateIdAndIsActive(Long agentId, Long affiliateId, Boolean isActive);

    List<CommissionConfig> findByAgentId(Long agentId);

    List<CommissionConfig> findByAffiliateId(Long affiliateId);
}
