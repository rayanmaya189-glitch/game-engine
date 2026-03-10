package com.gameengine.commission.repository;

import com.gameengine.commission.model.CommissionClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Repository
public interface CommissionClaimRepository extends JpaRepository<CommissionClaim, Long> {
    
    List<CommissionClaim> findByUserId(Long userId);
    
    List<CommissionClaim> findByAffiliateId(Long affiliateId);
    
    List<CommissionClaim> findByUserIdAndStatus(Long userId, String status);
    
    List<CommissionClaim> findByClaimType(String claimType);
    
    Optional<CommissionClaim> findByTransactionId(String transactionId);
    
    @Query("SELECT SUM(cc.amount) FROM CommissionClaim cc WHERE cc.userId = :userId AND cc.status = 'PAID'")
    BigDecimal getTotalPaidAmount(@Param("userId") Long userId);
    
    @Query("SELECT SUM(cc.amount) FROM CommissionClaim cc WHERE cc.userId = :userId AND cc.status = 'PENDING'")
    BigDecimal getTotalPendingAmount(@Param("userId") Long userId);
    
    @Query("SELECT cc FROM CommissionClaim cc WHERE cc.status = 'PENDING' AND cc.requestedAt <= :cutoffDate")
    List<CommissionClaim> findPendingOlderThan(@Param("cutoffDate") LocalDateTime cutoffDate);
    
    List<CommissionClaim> findByStatus(String status);
}
