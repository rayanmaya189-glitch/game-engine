package com.game_engine.bonus.service;

import com.game_engine.bonus.model.BonusCampaign;
import com.game_engine.bonus.model.PlayerBonus;
import com.game_engine.bonus.model.PlayerBonus.BonusStatus;
import com.game_engine.bonus.repository.BonusCampaignRepository;
import com.game_engine.bonus.repository.PlayerBonusRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.time.LocalDateTime;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class BonusCashbackService {

    private final BonusCampaignRepository campaignRepository;
    private final PlayerBonusRepository playerBonusRepository;
    private final WalletService walletService;
    private final BonusEligibilityService eligibilityService;

    /**
     * Calculate cashback for a user
     */
    @Transactional
    public BigDecimal calculateCashback(UUID userId, int periodDays) {
        BigDecimal totalBets = walletService.getTotalBets(userId, periodDays);
        BigDecimal totalWins = walletService.getTotalWins(userId, periodDays);
        BigDecimal netLosses = totalWins.subtract(totalBets);

        if (netLosses.compareTo(BigDecimal.ZERO) >= 0) {
            return BigDecimal.ZERO;
        }

        BonusCampaign cashbackCampaign = campaignRepository
                .findActiveCashbackCampaign(LocalDateTime.now());

        if (cashbackCampaign == null) {
            return BigDecimal.ZERO;
        }

        BigDecimal percentage = cashbackCampaign.getCashbackPercentage()
                .divide(BigDecimal.valueOf(100), 4, RoundingMode.HALF_UP);

        BigDecimal cashback = netLosses.abs().multiply(percentage);

        if (cashbackCampaign.getMaxBonusAmount() != null) {
            cashback = cashback.min(cashbackCampaign.getMaxBonusAmount());
        }

        return cashback;
    }

    /**
     * Claim cashback
     */
    @Transactional
    public PlayerBonus claimCashback(UUID userId) {
        BonusCampaign campaign = campaignRepository.findActiveCashbackCampaign(LocalDateTime.now());
        if (campaign == null) {
            throw new IllegalStateException("No active cashback campaign");
        }

        int periodDays = campaign.getCashbackPeriodDays() != null ? campaign.getCashbackPeriodDays() : 7;
        BigDecimal cashbackAmount = calculateCashback(userId, periodDays);

        if (cashbackAmount.compareTo(BigDecimal.ZERO) <= 0) {
            throw new IllegalStateException("No cashback available");
        }

        PlayerBonus bonus = eligibilityService.awardBonus(userId, campaign.getId(), "cashback", null);
        bonus.setBonusAmount(cashbackAmount);
        bonus.setBonusAmountCredited(cashbackAmount);
        bonus.setStatus(BonusStatus.CLAIMED);
        bonus.setClaimedAt(LocalDateTime.now());

        walletService.creditBalance(userId, cashbackAmount, "USD", "CASHBACK", bonus.getId().toString());

        return playerBonusRepository.save(bonus);
    }
}
