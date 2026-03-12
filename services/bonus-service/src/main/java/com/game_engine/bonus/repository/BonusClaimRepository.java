package com.game_engine.bonus.repository;

import com.game_engine.bonus.model.BonusClaim;
import com.game_engine.bonus.model.BonusClaim.ClaimStatus;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.Instant;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface BonusClaimRepository extends JpaRepository<BonusClaim, UUID> {

    List<BonusClaim> findByUserIdOrderByClaimedAtDesc(UUID userId);

    Optional<BonusClaim> findByUserIdAndBonusIdAndStatus(UUID userId, UUID bonusId, ClaimStatus status);

    List<BonusClaim> findByUserIdAndStatus(UUID userId, ClaimStatus status);

    List<BonusClaim> findByStatus(ClaimStatus status);

    @Query("SELECT bc FROM BonusClaim bc WHERE bc.status = 'ACTIVE' AND bc.expiresAt < :now")
    List<BonusClaim> findExpiredClaims(@Param("now") Instant now);

    @Query("SELECT COUNT(bc) FROM BonusClaim bc WHERE bc.userId = :userId AND bc.bonusId = :bonusId")
    int countByUserIdAndBonusId(@Param("userId") UUID userId, @Param("bonusId") UUID bonusId);

    @Query("SELECT bc FROM BonusClaim bc WHERE bc.userId = :userId AND bc.status = 'ACTIVE'")
    List<BonusClaim> findActiveClaimsByUserId(@Param("userId") UUID userId);

    @Query("SELECT SUM(bc.bonusAmount) FROM BonusClaim bc WHERE bc.userId = :userId AND bc.status = 'COMPLETED'")
    java.math.BigDecimal sumCompletedBonusesByUserId(@Param("userId") UUID userId);
}
