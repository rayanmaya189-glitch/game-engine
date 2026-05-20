package com.game_engine.commission.repository;

import com.game_engine.commission.entity.CommissionClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository("commissionClaimRepositoryV2")
public interface CommissionClaimV2Repository extends JpaRepository<CommissionClaim, Long> {

    List<CommissionClaim> findByAffiliateIdAndStatus(Long affiliateId, String status);

    List<CommissionClaim> findByAgentId(Long agentId);

    List<CommissionClaim> findByAffiliateId(Long affiliateId);

    List<CommissionClaim> findByStatus(String status);
}
