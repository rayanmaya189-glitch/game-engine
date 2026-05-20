package com.game_engine.commission.repository;

import com.game_engine.commission.model.CommissionClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.util.List;
import java.util.Optional;

@Repository("commissionClaimRepositoryV1")
public interface CommissionClaimRepository extends JpaRepository<CommissionClaim, Long> {

    List<CommissionClaim> findByUserId(Long userId);

    List<CommissionClaim> findByStatus(String status);

    @Query("SELECT SUM(c.amount) FROM CommissionClaim c WHERE c.userId = :userId AND c.status = 'PENDING'")
    BigDecimal getTotalPendingAmount(Long userId);
}
