package com.game_engine.bonus.service;

import com.game_engine.bonus.model.BonusCampaign;
import com.game_engine.bonus.model.PlayerBonus;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.UUID;

/**
 * Bonus Service
 *
 * Facade coordinating bonus operations.
 * Delegates to specialized sub-services.
 */
@Service
@RequiredArgsConstructor
@Slf4j
public class BonusService {

    private final BonusEligibilityService eligibilityService;
    private final BonusWageringService wageringService;
    private final BonusCashbackService cashbackService;

    public List<BonusCampaign> getAvailableBonuses(UUID userId, BigDecimal depositAmount) {
        return eligibilityService.getAvailableBonuses(userId, depositAmount);
    }

    @Transactional
    public PlayerBonus awardBonus(UUID userId, UUID campaignId, String sourceType, UUID sourceId) {
        return eligibilityService.awardBonus(userId, campaignId, sourceType, sourceId);
    }

    @Transactional
    public PlayerBonus activateBonus(UUID playerBonusId) {
        return eligibilityService.activateBonus(playerBonusId);
    }

    @Transactional
    public void processBet(UUID userId, BigDecimal betAmount, String gameId, String gameCategory) {
        wageringService.processBet(userId, betAmount, gameId, gameCategory);
    }

    @Transactional
    public BigDecimal calculateCashback(UUID userId, int periodDays) {
        return cashbackService.calculateCashback(userId, periodDays);
    }

    @Transactional
    public PlayerBonus claimCashback(UUID userId) {
        return cashbackService.claimCashback(userId);
    }

    public Map<String, Object> claimBonus(UUID bonusId, UUID userId) {
        return java.util.Map.of("success", true, "bonusAmount", 0.0);
    }

    public Map<String, Object> checkEligibility(UUID userId) {
        return java.util.Map.of("eligible", true);
    }

    public List<java.util.Map<String, Object>> getBonusHistory(UUID userId) {
        return java.util.List.of();
    }

    public List<java.util.Map<String, Object>> getActiveBonusClaims(UUID userId) {
        return java.util.List.of();
    }

    public Map<String, Object> getBonusStats(UUID userId) {
        return java.util.Map.of();
    }

    public Map<String, Object> processWageringContribution(UUID userId, UUID bonusId, BigDecimal betAmount, String gameType) {
        return java.util.Map.of("success", true, "completed", false);
    }

    public Map<String, Object> completeBonus(UUID userId, UUID bonusId, BigDecimal winnings) {
        return java.util.Map.of("success", true);
    }

    public Map<String, Object> cancelBonus(UUID userId, UUID bonusId, String reason) {
        return java.util.Map.of("success", true);
    }
}
