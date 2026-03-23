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
import java.math.RoundingMode;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.UUID;

/**
 * Bonus Service
 * 
 * Manages bonus campaigns, awarding, and wagering requirements.
 */
@Service
@RequiredArgsConstructor
@Slf4j
public class BonusService {

    private final BonusCampaignRepository campaignRepository;
    private final PlayerBonusRepository playerBonusRepository;
    private final WalletService walletService;

    /**
     * Check available bonuses for a user (e.g., after deposit)
     */
    public List<BonusCampaign> getAvailableBonuses(UUID userId, BigDecimal depositAmount) {
        LocalDateTime now = LocalDateTime.now();
        
        return campaignRepository.findActiveCampaigns(now).stream()
            .filter(c -> c.getMinDeposit() == null || depositAmount.compareTo(c.getMinDeposit()) >= 0)
            .filter(c -> canUserClaim(userId, c))
            .toList();
    }

    /**
     * Award a bonus to a user
     */
    @Transactional
    public PlayerBonus awardBonus(UUID userId, UUID campaignId, String sourceType, UUID sourceId) {
        BonusCampaign campaign = campaignRepository.findById(campaignId)
                .orElseThrow(() -> new IllegalArgumentException("Campaign not found"));

        // Check eligibility
        if (!canUserClaim(userId, campaign)) {
            throw new IllegalStateException("User not eligible for this bonus");
        }

        // Calculate bonus amount based on type
        BigDecimal bonusAmount = calculateBonusAmount(campaign, sourceType, sourceId);

        // Create player bonus
        PlayerBonus playerBonus = PlayerBonus.builder()
                .userId(userId)
                .campaign(campaign)
                .bonusAmount(bonusAmount)
                .bonusAmountCredited(BigDecimal.ZERO)
                .status(BonusStatus.PENDING)
                .sourceType(sourceType)
                .sourceId(sourceId)
                .build();

        // Set wagering requirement
        if (campaign.getWageringMultiplier() != null && campaign.getWageringMultiplier() > 0) {
            playerBonus.setWageringRequired(
                bonusAmount.multiply(BigDecimal.valueOf(campaign.getWageringMultiplier()))
            );
        }

        // Set expiry
        if (campaign.getExpiryDays() != null) {
            playerBonus.setExpiresAt(LocalDateTime.now().plusDays(campaign.getExpiryDays()));
        }

        // Handle different bonus types
        switch (campaign.getBonusType()) {
            case FREE_SPINS:
                playerBonus.setFreeSpinsRemaining(campaign.getFreeSpinCount());
                playerBonus.setFreeSpinsGameId(campaign.getFreeSpinGameId());
                // Free spins don't require activation
                playerBonus.setStatus(BonusStatus.ACTIVE);
                playerBonus.setActivatedAt(LocalDateTime.now());
                break;
                
            case NO_DEPOSIT:
                // Credit immediately
                playerBonus.setBonusAmountCredited(bonusAmount);
                walletService.creditBonusBalance(userId, bonusAmount, "USD", "NO_DEPOSIT_BONUS", playerBonus.getId().toString());
                playerBonus.setStatus(BonusStatus.ACTIVE);
                playerBonus.setActivatedAt(LocalDateTime.now());
                break;
                
            case WELCOME:
            case RELOAD:
                // Pending - will activate on deposit
                break;
                
            case CASHBACK:
                // Will be calculated and awarded periodically
                break;
        }

        return playerBonusRepository.save(playerBonus);
    }

    /**
     * Activate a pending bonus (e.g., after making qualifying deposit)
     */
    @Transactional
    public PlayerBonus activateBonus(UUID playerBonusId) {
        PlayerBonus bonus = playerBonusRepository.findById(playerBonusId)
                .orElseThrow(() -> new IllegalArgumentException("Bonus not found"));

        if (bonus.getStatus() != BonusStatus.PENDING) {
            throw new IllegalStateException("Bonus is not pending");
        }

        bonus.setStatus(BonusStatus.ACTIVE);
        bonus.setActivatedAt(LocalDateTime.now());

        // Credit bonus balance for match bonuses
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

    /**
     * Process a bet and update wagering progress
     */
    @Transactional
    public void processBet(UUID userId, BigDecimal betAmount, String gameId, String gameCategory) {
        // Get active bonuses with wagering requirements
        List<PlayerBonus> activeBonuses = playerBonusRepository
                .findActiveByUserId(userId, BonusStatus.ACTIVE);

        if (activeBonuses.isEmpty()) {
            return;
        }

        // Get game weight
        double gameWeight = getGameWeight(gameCategory);

        // Calculate wagering contribution
        BigDecimal contribution = betAmount.multiply(BigDecimal.valueOf(gameWeight));

        // Distribute across bonuses (oldest first)
        for (PlayerBonus bonus : activeBonuses) {
            if (bonus.getWageringRequired() == null || 
                bonus.getWageringRequired().compareTo(BigDecimal.ZERO) <= 0) {
                continue;
            }

            BigDecimal remaining = bonus.getWageringRequired()
                    .subtract(bonus.getWageringProgress());
            
            if (remaining.compareTo(BigDecimal.ZERO) <= 0) {
                continue;
            }

            BigDecimal toApply = contribution.compareTo(remaining) > 0 ? remaining : contribution;
            bonus.setWageringProgress(bonus.getWageringProgress().add(toApply));
            contribution = contribution.subtract(toApply);

            // Check if wagering complete
            if (bonus.getWageringProgress().compareTo(bonus.getWageringRequired()) >= 0) {
                completeBonus(bonus);
            }

            playerBonusRepository.save(bonus);

            if (contribution.compareTo(BigDecimal.ZERO) <= 0) {
                break;
            }
        }
    }

    /**
     * Complete a bonus (wagering requirements met)
     */
    @Transactional
    public void completeBonus(PlayerBonus bonus) {
        bonus.setStatus(BonusStatus.COMPLETED);
        bonus.setCompletedAt(LocalDateTime.now());

        // Convert bonus to real money
        BigDecimal bonusBalance = walletService.getBonusBalance(bonus.getUserId());
        if (bonusBalance.compareTo(BigDecimal.ZERO) > 0) {
            // Check max cashout
            BigDecimal maxCashout = bonus.getCampaign().getMaxCashout();
            if (maxCashout != null && maxCashout.compareTo(BigDecimal.ZERO) > 0) {
                bonusBalance = bonusBalance.min(maxCashout);
            }
            
            walletService.convertBonusToReal(bonus.getUserId(), bonusBalance);
        }

        playerBonusRepository.save(bonus);
        log.info("Bonus completed: {} for user {}", bonus.getId(), bonus.getUserId());
    }

    /**
     * Calculate cashback for a user
     */
    @Transactional
    public BigDecimal calculateCashback(UUID userId, int periodDays) {
        // Calculate net losses for period
        BigDecimal totalBets = walletService.getTotalBets(userId, periodDays);
        BigDecimal totalWins = walletService.getTotalWins(userId, periodDays);
        BigDecimal netLosses = totalWins.subtract(totalBets);

        if (netLosses.compareTo(BigDecimal.ZERO) >= 0) {
            return BigDecimal.ZERO; // No losses = no cashback
        }

        // Get cashback campaign
        BonusCampaign cashbackCampaign = campaignRepository
                .findActiveCashbackCampaign(LocalDateTime.now());

        if (cashbackCampaign == null) {
            return BigDecimal.ZERO;
        }

        // Calculate cashback percentage
        BigDecimal percentage = cashbackCampaign.getCashbackPercentage()
                .divide(BigDecimal.valueOf(100), 4, RoundingMode.HALF_UP);
        
        BigDecimal cashback = netLosses.abs().multiply(percentage);

        // Cap at max bonus if set
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

        // Award cashback
        PlayerBonus bonus = awardBonus(userId, campaign.getId(), "cashback", null);
        bonus.setBonusAmount(cashbackAmount);
        bonus.setBonusAmountCredited(cashbackAmount);
        bonus.setStatus(BonusStatus.CLAIMED);
        bonus.setClaimedAt(LocalDateTime.now());

        // Credit as real money (no wagering on cashback typically)
        walletService.creditBalance(userId, cashbackAmount, "USD", "CASHBACK", bonus.getId().toString());

        return playerBonusRepository.save(bonus);
    }

    private boolean canUserClaim(UUID userId, BonusCampaign campaign) {
        // Check if already claimed (for one-time bonuses)
        if (campaign.getMaxUsageCount() != null && campaign.getMaxUsageCount() == 1) {
            boolean alreadyClaimed = playerBonusRepository
                    .existsByUserIdAndCampaignId(userId, campaign.getId());
            if (alreadyClaimed) {
                return false;
            }
        }

        // Check promo code requirement
        // Check country eligibility
        // Check KYC level
        
        return true;
    }

    private BigDecimal calculateBonusAmount(BonusCampaign campaign, String sourceType, UUID sourceId) {
        switch (campaign.getBonusType()) {
            case FREE_SPINS:
                return BigDecimal.ZERO; // Free spins don't have monetary value initially
                
            case WELCOME:
            case RELOAD:
                // Would get deposit amount from source
                BigDecimal depositAmount = BigDecimal.valueOf(100); // Simplified
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

    private double getGameWeight(String gameCategory) {
        // Default weights - would load from campaign config
        return switch (gameCategory.toLowerCase()) {
            case "slots" -> 1.0;
            case "table_games", "blackjack" -> 0.1;
            case "roulette" -> 0.2;
            case "poker" -> 0.05;
            default -> 1.0;
        };
    }
}
