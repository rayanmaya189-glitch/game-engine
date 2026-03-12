package com.game_engine.bonus.service;

import com.game_engine.bonus.model.Bonus;
import com.game_engine.bonus.model.BonusClaim;
import com.game_engine.bonus.model.BonusClaim.ClaimStatus;
import com.game_engine.bonus.repository.BonusRepository;
import com.game_engine.bonus.repository.BonusClaimRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import java.math.BigDecimal;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.*;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class BonusService {
    private final BonusRepository bonusRepository;
    private final BonusClaimRepository bonusClaimRepository;

    private static final int DEFAULT_BONUS_VALIDITY_DAYS = 30;

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
        
        // Check if user has already claimed this bonus
        int existingClaims = bonusClaimRepository.countByUserIdAndBonusId(userId, bonusId);
        if (existingClaims > 0) {
            result.put("success", false);
            result.put("message", "You have already claimed this bonus");
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
        
        // Calculate bonus amount
        BigDecimal bonusAmount = bonus.getAmount();
        if (bonus.getPercentage() != null && bonus.getMinDeposit() != null) {
            // For deposit bonuses, calculate percentage
            BigDecimal calculatedAmount = bonus.getMinDeposit().multiply(bonus.getPercentage()).divide(BigDecimal.valueOf(100));
            if (bonus.getMaxAmount() != null && calculatedAmount.compareTo(bonus.getMaxAmount()) > 0) {
                calculatedAmount = bonus.getMaxAmount();
            }
            bonusAmount = calculatedAmount;
        }
        
        // Calculate wagering requirement
        BigDecimal wageringRequirement = BigDecimal.ZERO;
        if (bonus.getWageringRequirement() != null && bonusAmount != null) {
            wageringRequirement = bonusAmount.multiply(BigDecimal.valueOf(bonus.getWageringRequirement()));
        }
        
        // Calculate expiry date (default 30 days)
        Instant expiresAt = now.plus(DEFAULT_BONUS_VALIDITY_DAYS, ChronoUnit.DAYS);
        
        // Create bonus claim record
        BonusClaim claim = BonusClaim.builder()
            .userId(userId)
            .bonusId(bonusId)
            .bonusAmount(bonusAmount)
            .wageringRequirement(wageringRequirement)
            .wageringContributed(BigDecimal.ZERO)
            .status(ClaimStatus.ACTIVE)
            .claimedAt(now)
            .expiresAt(expiresAt)
            .build();
        bonusClaimRepository.save(claim);
        
        // Increment usage count
        bonus.setCurrentUses(bonus.getCurrentUses() != null ? bonus.getCurrentUses() + 1 : 1);
        bonusRepository.save(bonus);
        
        result.put("success", true);
        result.put("message", "Bonus claimed successfully");
        result.put("bonusId", bonusId);
        result.put("userId", userId);
        result.put("bonusAmount", bonusAmount);
        result.put("wageringRequirement", wageringRequirement);
        result.put("expiresAt", expiresAt);
        result.put("claimId", claim.getId());
        
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
        List<BonusClaim> claims = bonusClaimRepository.findByUserIdOrderByClaimedAtDesc(userId);
        List<Map<String, Object>> history = new ArrayList<>();
        
        for (BonusClaim claim : claims) {
            Map<String, Object> entry = new HashMap<>();
            entry.put("claimId", claim.getId());
            entry.put("bonusId", claim.getBonusId());
            entry.put("bonusAmount", claim.getBonusAmount());
            entry.put("wageringRequirement", claim.getWageringRequirement());
            entry.put("wageringContributed", claim.getWageringContributed() != null ? claim.getWageringContributed() : BigDecimal.ZERO);
            entry.put("winningsAmount", claim.getWinningsAmount());
            entry.put("status", claim.getStatus());
            entry.put("claimedAt", claim.getClaimedAt());
            entry.put("completedAt", claim.getCompletedAt());
            entry.put("expiresAt", claim.getExpiresAt());
            
            // Get bonus details
            Bonus bonus = getBonusById(claim.getBonusId());
            if (bonus != null) {
                entry.put("bonusName", bonus.getName());
                entry.put("bonusType", bonus.getType());
            }
            
            history.add(entry);
        }
        
        return history;
    }

    @Transactional
    public Map<String, Object> processWageringContribution(UUID userId, UUID bonusId, BigDecimal betAmount, String gameType) {
        Map<String, Object> result = new HashMap<>();
        
        Optional<BonusClaim> claimOpt = bonusClaimRepository.findByUserIdAndBonusIdAndStatus(userId, bonusId, ClaimStatus.ACTIVE);
        if (claimOpt.isEmpty()) {
            result.put("success", false);
            result.put("message", "No active bonus claim found");
            return result;
        }
        
        BonusClaim claim = claimOpt.get();
        
        // Check if expired
        if (claim.getExpiresAt() != null && Instant.now().isAfter(claim.getExpiresAt())) {
            claim.setStatus(ClaimStatus.EXPIRED);
            bonusClaimRepository.save(claim);
            result.put("success", false);
            result.put("message", "Bonus has expired");
            return result;
        }
        
        // Calculate wagering contribution based on game type
        BigDecimal contribution = calculateWageringContribution(betAmount, gameType);
        
        // Update wagering contributed
        BigDecimal currentContributed = claim.getWageringContributed() != null ? claim.getWageringContributed() : BigDecimal.ZERO;
        claim.setWageringContributed(currentContributed.add(contribution));
        bonusClaimRepository.save(claim);
        
        result.put("success", true);
        result.put("message", "Wagering contribution added");
        result.put("contribution", contribution);
        result.put("totalContributed", claim.getWageringContributed());
        result.put("remaining", claim.getRemainingWagering());
        result.put("wageringComplete", claim.isWageringComplete());
        
        return result;
    }

    @Transactional
    public Map<String, Object> completeBonus(UUID userId, UUID bonusId, BigDecimal winnings) {
        Map<String, Object> result = new HashMap<>();
        
        Optional<BonusClaim> claimOpt = bonusClaimRepository.findByUserIdAndBonusIdAndStatus(userId, bonusId, ClaimStatus.ACTIVE);
        if (claimOpt.isEmpty()) {
            result.put("success", false);
            result.put("message", "No active bonus claim found");
            return result;
        }
        
        BonusClaim claim = claimOpt.get();
        
        if (!claim.isWageringComplete()) {
            result.put("success", false);
            result.put("message", "Wagering requirements not yet met");
            return result;
        }
        
        claim.setStatus(ClaimStatus.COMPLETED);
        claim.setCompletedAt(Instant.now());
        claim.setWinningsAmount(winnings);
        bonusClaimRepository.save(claim);
        
        result.put("success", true);
        result.put("message", "Bonus completed successfully");
        result.put("bonusAmount", claim.getBonusAmount());
        result.put("winnings", winnings);
        result.put("totalAmount", claim.getBonusAmount().add(winnings));
        
        return result;
    }

    @Transactional
    public Map<String, Object> cancelBonus(UUID userId, UUID bonusId, String reason) {
        Map<String, Object> result = new HashMap<>();
        
        Optional<BonusClaim> claimOpt = bonusClaimRepository.findByUserIdAndBonusIdAndStatus(userId, bonusId, ClaimStatus.ACTIVE);
        if (claimOpt.isEmpty()) {
            result.put("success", false);
            result.put("message", "No active bonus claim found");
            return result;
        }
        
        BonusClaim claim = claimOpt.get();
        claim.setStatus(ClaimStatus.CANCELLED);
        claim.setCancelledAt(Instant.now());
        claim.setCancellationReason(reason);
        bonusClaimRepository.save(claim);
        
        result.put("success", true);
        result.put("message", "Bonus cancelled");
        result.put("reason", reason);
        
        return result;
    }

    public List<Map<String, Object>> getActiveBonusClaims(UUID userId) {
        List<BonusClaim> claims = bonusClaimRepository.findActiveClaimsByUserId(userId);
        List<Map<String, Object>> activeClaims = new ArrayList<>();
        
        for (BonusClaim claim : claims) {
            Map<String, Object> entry = new HashMap<>();
            entry.put("claimId", claim.getId());
            entry.put("bonusId", claim.getBonusId());
            entry.put("bonusAmount", claim.getBonusAmount());
            entry.put("wageringRequirement", claim.getWageringRequirement());
            entry.put("wageringContributed", claim.getWageringContributed() != null ? claim.getWageringContributed() : BigDecimal.ZERO);
            entry.put("remaining", claim.getRemainingWagering());
            entry.put("expiresAt", claim.getExpiresAt());
            
            Bonus bonus = getBonusById(claim.getBonusId());
            if (bonus != null) {
                entry.put("bonusName", bonus.getName());
                entry.put("bonusType", bonus.getType());
                entry.put("allowedGames", bonus.getAllowedGames());
            }
            
            activeClaims.add(entry);
        }
        
        return activeClaims;
    }

    @Transactional
    public void processExpiredBonuses() {
        List<BonusClaim> expiredClaims = bonusClaimRepository.findExpiredClaims(Instant.now());
        for (BonusClaim claim : expiredClaims) {
            claim.setStatus(ClaimStatus.EXPIRED);
            bonusClaimRepository.save(claim);
        }
    }

    public Map<String, Object> getBonusStats(UUID userId) {
        Map<String, Object> stats = new HashMap<>();
        
        List<BonusClaim> allClaims = bonusClaimRepository.findByUserIdOrderByClaimedAtDesc(userId);
        List<BonusClaim> activeClaims = bonusClaimRepository.findByUserIdAndStatus(userId, ClaimStatus.ACTIVE);
        List<BonusClaim> completedClaims = bonusClaimRepository.findByUserIdAndStatus(userId, ClaimStatus.COMPLETED);
        
        stats.put("totalBonusesClaimed", allClaims.size());
        stats.put("activeBonuses", activeClaims.size());
        stats.put("completedBonuses", completedClaims.size());
        stats.put("totalBonusAmount", allClaims.stream()
            .map(BonusClaim::getBonusAmount)
            .filter(Objects::nonNull)
            .reduce(BigDecimal.ZERO, BigDecimal::add));
        stats.put("totalWinnings", completedClaims.stream()
            .map(BonusClaim::getWinningsAmount)
            .filter(Objects::nonNull)
            .reduce(BigDecimal.ZERO, BigDecimal::add));
        
        return stats;
    }
}
