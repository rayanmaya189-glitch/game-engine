package com.game_engine.commission.repository;

import com.game_engine.commission.entity.CommissionClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository("commissionClaimRepositoryV2")
public interface CommissionClaimRepository extends JpaRepository<CommissionClaim, Long> {

    List<CommissionClaim> findByAgentId(Long agentId);

    List<CommissionClaim> findByAffiliateId(Long affiliateId);

    List<CommissionClaim> findByStatus(String status);

    Optional<CommissionClaim> findByClaimId(String claimId);

    List<CommissionClaim> findByAgentIdAndStatus(Long agentId, String status);

    List<CommissionClaim> findByAffiliateIdAndStatus(Long affiliateId, String status);

    List<CommissionClaim> findByPeriod(String period);
}
