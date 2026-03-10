package com.gameengine.commission.repository;

import com.gameengine.commission.model.InsuranceClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Repository
public interface InsuranceClaimRepository extends JpaRepository<InsuranceClaim, Long> {
    
    List<InsuranceClaim> findByUserId(Long userId);
    
    List<InsuranceClaim> findByUserIdAndStatus(Long userId, String status);
    
    List<InsuranceClaim> findByGameId(Long gameId);
    
    List<InsuranceClaim> findByBetId(Long betId);
    
    Optional<InsuranceClaim> findByInsurancePolicyId(String insurancePolicyId);
    
    List<InsuranceClaim> findByClaimType(String claimType);
    
    @Query("SELECT SUM(ic.claimAmount) FROM InsuranceClaim ic WHERE ic.status = 'PAID'")
    BigDecimal getTotalPaidClaims();
    
    @Query("SELECT SUM(ic.claimAmount) FROM InsuranceClaim ic WHERE ic.userId = :userId AND ic.status = 'PAID'")
    BigDecimal getTotalPaidByUser(@Param("userId") Long userId);
    
    @Query("SELECT ic FROM InsuranceClaim ic WHERE ic.status = 'PENDING' AND ic.requestedAt <= :cutoffDate")
    List<InsuranceClaim> findPendingOlderThan(@Param("cutoffDate") LocalDateTime cutoffDate);
    
    List<InsuranceClaim> findByStatus(String status);
}
