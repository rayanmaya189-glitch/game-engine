package com.game_engine.commission.repository;

import com.game_engine.commission.entity.CommissionSettlement;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CommissionSettlementRepository extends JpaRepository<CommissionSettlement, Long> {

    List<CommissionSettlement> findByAgentId(Long agentId);

    List<CommissionSettlement> findByAffiliateId(Long affiliateId);

    List<CommissionSettlement> findByStatus(String status);

    Optional<CommissionSettlement> findBySettlementId(String settlementId);

    List<CommissionSettlement> findByAgentIdAndAffiliateId(Long agentId, Long affiliateId);
}
