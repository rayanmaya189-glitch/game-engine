package com.game_engine.commission.repository;

import com.game_engine.commission.model.RebetClaim;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface RebetClaimRepository extends JpaRepository<RebetClaim, Long> {
    
    List<RebetClaim> findByUserId(Long userId);
    
    List<RebetClaim> findByUserIdAndStatus(Long userId, String status);
    
    List<RebetClaim> findByBonusId(Long bonusId);
    
    Optional<RebetClaim> findByBonusCode(String bonusCode);
    
    Optional<RebetClaim> findByBetId(Long betId);
    
    @Query("SELECT rc FROM RebetClaim rc WHERE rc.status = 'CLAIMABLE' AND rc.expiresAt > CURRENT_TIMESTAMP")
    List<RebetClaim> findExpiredClaimable();
    
    @Query("SELECT rc FROM RebetClaim rc WHERE rc.status = 'IN_PROGRESS' AND rc.expiresAt <= CURRENT_TIMESTAMP")
    List<RebetClaim> findExpired();
    
    List<RebetClaim> findByStatus(String status);
}
