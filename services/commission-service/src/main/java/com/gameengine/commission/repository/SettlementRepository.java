package com.gameengine.commission.repository;

import com.gameengine.commission.model.Settlement;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Repository
public interface SettlementRepository extends JpaRepository<Settlement, Long> {
    
    List<Settlement> findByUserId(Long userId);
    
    List<Settlement> findByUserIdAndStatus(Long userId, String status);
    
    List<Settlement> findBySettlementType(String settlementType);
    
    Optional<Settlement> findByTransactionId(String transactionId);
    
    Optional<Settlement> findByReferenceId(String referenceId);
    
    List<Settlement> findByClaimId(Long claimId);
    
    @Query("SELECT SUM(s.netAmount) FROM Settlement s WHERE s.userId = :userId AND s.status = 'COMPLETED'")
    BigDecimal getTotalSettledAmount(@Param("userId") Long userId);
    
    @Query("SELECT SUM(s.netAmount) FROM Settlement s WHERE s.settlementType = :type AND s.status = 'COMPLETED'")
    BigDecimal getTotalByType(@Param("type") String type);
    
    @Query("SELECT s FROM Settlement s WHERE s.status = 'PENDING' AND s.createdAt <= :cutoffDate")
    List<Settlement> findPendingOlderThan(@Param("cutoffDate") LocalDateTime cutoffDate);
    
    List<Settlement> findByStatus(String status);
    
    @Query("SELECT s FROM Settlement s WHERE s.createdAt BETWEEN :startDate AND :endDate")
    List<Settlement> findByDateRange(@Param("startDate") LocalDateTime startDate, @Param("endDate") LocalDateTime endDate);
}
