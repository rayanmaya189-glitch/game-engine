package com.game_engine.bonus.service;

import com.game_engine.bonus.model.Bonus;
import com.game_engine.bonus.repository.BonusRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import java.math.BigDecimal;
import java.time.Instant;
import java.util.*;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class BonusService {
    private final BonusRepository bonusRepository;

    public Bonus createBonus(Bonus bonus) {
        return bonusRepository.save(bonus);
    }

    public List<Bonus> getActiveBonuses() {
        return bonusRepository.findByStatus(Bonus.BonusStatus.ACTIVE);
    }

    public Bonus getBonusById(UUID id) {
        return bonusRepository.findById(id).orElse(null);
    }

    public BigDecimal calculateWageringContribution(BigDecimal betAmount, String gameType) {
        // Slot games contribute 100%, table games 10%
        return "SLOT".equalsIgnoreCase(gameType) ? betAmount : betAmount.multiply(BigDecimal.valueOf(0.1));
    }

    public boolean isWageringComplete(UUID bonusId, BigDecimal totalWagered, BigDecimal wageringRequirement) {
        return totalWagered.compareTo(wageringRequirement) >= 0;
    }

    @Transactional
    public Map<String, Object> claimBonus(UUID bonusId, UUID userId) {
        Map<String, Object> result = new HashMap<>();
        Bonus bonus = getBonusById(bonusId);
        
        if (bonus == null) {
            result.put("success", false);
            result.put("message", "Bonus not found");
            return result;
        }
        
        if (bonus.getStatus() != Bonus.BonusStatus.ACTIVE) {
            result.put("success", false);
            result.put("message", "Bonus is not active");
            return result;
        }
        
        // Check if bonus has reached max uses
        if (bonus.getMaxUses() != null && bonus.getCurrentUses() != null 
            && bonus.getCurrentUses() >= bonus.getMaxUses()) {
            result.put("success", false);
            result.put("message", "Bonus has reached maximum uses");
            return result;
        }
        
        // Check if bonus is still valid
        Instant now = Instant.now();
        if (bonus.getStartDate() != null && now.isBefore(bonus.getStartDate())) {
            result.put("success", false);
            result.put("message", "Bonus has not started yet");
            return result;
        }
        
        if (bonus.getEndDate() != null && now.isAfter(bonus.getEndDate())) {
            result.put("success", false);
            result.put("message", "Bonus has expired");
            return result;
        }
        
        // Increment usage count
        bonus.setCurrentUses(bonus.getCurrentUses() != null ? bonus.getCurrentUses() + 1 : 1);
        bonusRepository.save(bonus);
        
        result.put("success", true);
        result.put("message", "Bonus claimed successfully");
        result.put("bonusId", bonusId);
        result.put("userId", userId);
        result.put("bonusAmount", bonus.getAmount());
        result.put("wageringRequirement", bonus.getWageringRequirement());
        
        return result;
    }

    public Map<String, Object> checkEligibility(UUID userId) {
        Map<String, Object> result = new HashMap<>();
        List<Bonus> activeBonuses = getActiveBonuses();
        List<Bonus> eligibleBonuses = new ArrayList<>();
        
        Instant now = Instant.now();
        for (Bonus bonus : activeBonuses) {
            boolean eligible = true;
            StringBuilder reason = new StringBuilder();
            
            // Check if bonus has started
            if (bonus.getStartDate() != null && now.isBefore(bonus.getStartDate())) {
                eligible = false;
                reason.append("Bonus has not started yet; ");
            }
            
            // Check if bonus has expired
            if (bonus.getEndDate() != null && now.isAfter(bonus.getEndDate())) {
                eligible = false;
                reason.append("Bonus has expired; ");
            }
            
            // Check if bonus has reached max uses
            if (bonus.getMaxUses() != null && bonus.getCurrentUses() != null 
                && bonus.getCurrentUses() >= bonus.getMaxUses()) {
                eligible = false;
                reason.append("Bonus has reached maximum uses; ");
            }
            
            if (eligible) {
                eligibleBonuses.add(bonus);
            }
        }
        
        result.put("userId", userId);
        result.put("eligibleBonuses", eligibleBonuses);
        result.put("totalEligible", eligibleBonuses.size());
        
        return result;
    }

    public List<Map<String, Object>> getBonusHistory(UUID userId) {
        List<Map<String, Object>> history = new ArrayList<>();
        
        // TODO: Implement bonus history retrieval
        // This requires a new repository method to query the bonus_claims or user_bonuses table
        // Example implementation:
        // List<BonusClaim> claims = bonusClaimRepository.findByUserId(userId);
        // for (BonusClaim claim : claims) {
        //     Map<String, Object> entry = new HashMap<>();
        //     entry.put("bonusId", claim.getBonusId());
        //     entry.put("claimedAt", claim.getClaimedAt());
        //     entry.put("status", claim.getStatus());
        //     history.add(entry);
        // }
        
        return history;
    }
}
