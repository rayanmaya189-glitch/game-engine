package com.game_engine.bonus.service;

import com.game_engine.bonus.model.BonusCampaign;
import com.game_engine.bonus.model.BonusCampaign.BonusType;
import com.game_engine.bonus.model.PlayerBonus;
import com.game_engine.bonus.model.PlayerBonus.BonusStatus;
import com.game_engine.bonus.repository.BonusCampaignRepository;
import com.game_engine.bonus.repository.PlayerBonusRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class BonusEligibilityService {

    private final BonusCampaignRepository campaignRepository;
    private final PlayerBonusRepository playerBonusRepository;
    private final WalletService walletService;

    public List<BonusCampaign> getAvailableBonuses(UUID userId, BigDecimal depositAmount) {
        LocalDateTime now = LocalDateTime.now();

        return campaignRepository.findActiveCampaigns(now).stream()
            .filter(c -> c.getMinDeposit() == null || depositAmount.compareTo(c.getMinDeposit()) >= 0)
            .filter(c -> canUserClaim(userId, c))
            .toList();
    }

    @Transactional
    public PlayerBonus awardBonus(UUID userId, UUID campaignId, String sourceType, UUID sourceId) {
        BonusCampaign campaign = campaignRepository.findById(campaignId)
                .orElseThrow(() -> new IllegalArgumentException("Campaign not found"));

        if (!canUserClaim(userId, campaign)) {
            throw new IllegalStateException("User not eligible for this bonus");
        }

        BigDecimal bonusAmount = calculateBonusAmount(campaign, sourceType, sourceId);

        PlayerBonus playerBonus = PlayerBonus.builder()
                .userId(userId)
                .campaign(campaign)
                .bonusAmount(bonusAmount)
                .bonusAmountCredited(BigDecimal.ZERO)
                .status(BonusStatus.PENDING)
                .sourceType(sourceType)
                .sourceId(sourceId)
                .build();

        if (campaign.getWageringMultiplier() != null && campaign.getWageringMultiplier() > 0) {
            playerBonus.setWageringRequired(
                bonusAmount.multiply(BigDecimal.valueOf(campaign.getWageringMultiplier()))
            );
        }

        if (campaign.getExpiryDays() != null) {
            playerBonus.setExpiresAt(LocalDateTime.now().plusDays(campaign.getExpiryDays()));
        }

        switch (campaign.getBonusType()) {
            case FREE_SPINS:
                playerBonus.setFreeSpinsRemaining(campaign.getFreeSpinCount());
                playerBonus.setFreeSpinsGameId(campaign.getFreeSpinGameId());
                playerBonus.setStatus(BonusStatus.ACTIVE);
                playerBonus.setActivatedAt(LocalDateTime.now());
                break;

            case NO_DEPOSIT:
                playerBonus.setBonusAmountCredited(bonusAmount);
                walletService.creditBonusBalance(userId, bonusAmount, "USD", "NO_DEPOSIT_BONUS", playerBonus.getId().toString());
                playerBonus.setStatus(BonusStatus.ACTIVE);
                playerBonus.setActivatedAt(LocalDateTime.now());
                break;

            case WELCOME:
            case RELOAD:
                break;

            case CASHBACK:
                break;
        }

        return playerBonusRepository.save(playerBonus);
    }

    @Transactional
    public PlayerBonus activateBonus(UUID playerBonusId) {
        PlayerBonus bonus = playerBonusRepository.findById(playerBonusId)
                .orElseThrow(() -> new IllegalArgumentException("Bonus not found"));

        if (bonus.getStatus() != BonusStatus.PENDING) {
            throw new IllegalStateException("Bonus is not pending");
        }

        bonus.setStatus(BonusStatus.ACTIVE);
        bonus.setActivatedAt(LocalDateTime.now());

        if (bonus.getCampaign().getBonusType() == BonusType.WELCOME ||
            bonus.getCampaign().getBonusType() == BonusType.RELOAD) {
            bonus.setBonusAmountCredited(bonus.getBonusAmount());
            walletService.creditBonusBalance(
                bonus.getUserId(),
                bonus.getBonusAmount(),
                "USD",
                "DEPOSIT_BONUS",
                bonus.getId().toString()
            );
        }

        return playerBonusRepository.save(bonus);
    }

    boolean canUserClaim(UUID userId, BonusCampaign campaign) {
        if (campaign.getMaxUsageCount() != null && campaign.getMaxUsageCount() == 1) {
            boolean alreadyClaimed = playerBonusRepository
                    .existsByUserIdAndCampaignId(userId, campaign.getId());
            if (alreadyClaimed) {
                return false;
            }
        }

        return true;
    }

    BigDecimal calculateBonusAmount(BonusCampaign campaign, String sourceType, UUID sourceId) {
        switch (campaign.getBonusType()) {
            case FREE_SPINS:
                return BigDecimal.ZERO;

            case WELCOME:
            case RELOAD:
                BigDecimal depositAmount = BigDecimal.valueOf(100);
                if (campaign.getMatchPercentage() != null) {
                    BigDecimal bonus = depositAmount.multiply(
                        campaign.getMatchPercentage().divide(BigDecimal.valueOf(100))
                    );
                    if (campaign.getMaxBonusAmount() != null) {
                        bonus = bonus.min(campaign.getMaxBonusAmount());
                    }
                    return bonus;
                }
                return campaign.getFixedAmount() != null ? campaign.getFixedAmount() : BigDecimal.ZERO;

            case NO_DEPOSIT:
            case REFERRAL:
                return campaign.getFixedAmount() != null ? campaign.getFixedAmount() : BigDecimal.ZERO;

            default:
                return campaign.getFixedAmount() != null ? campaign.getFixedAmount() : BigDecimal.ZERO;
        }
    }
}
