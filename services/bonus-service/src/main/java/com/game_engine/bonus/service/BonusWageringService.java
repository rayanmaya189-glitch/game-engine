package com.game_engine.bonus.service;

import com.game_engine.bonus.model.PlayerBonus;
import com.game_engine.bonus.model.PlayerBonus.BonusStatus;
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
public class BonusWageringService {

    private final PlayerBonusRepository playerBonusRepository;
    private final WalletService walletService;

    /**
     * Process a bet and update wagering progress
     */
    @Transactional
    public void processBet(UUID userId, BigDecimal betAmount, String gameId, String gameCategory) {
        List<PlayerBonus> activeBonuses = playerBonusRepository
                .findActiveByUserId(userId, BonusStatus.ACTIVE);

        if (activeBonuses.isEmpty()) {
            return;
        }

        double gameWeight = getGameWeight(gameCategory);
        BigDecimal contribution = betAmount.multiply(BigDecimal.valueOf(gameWeight));

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

        BigDecimal bonusBalance = walletService.getBonusBalance(bonus.getUserId());
        if (bonusBalance.compareTo(BigDecimal.ZERO) > 0) {
            BigDecimal maxCashout = bonus.getCampaign().getMaxCashout();
            if (maxCashout != null && maxCashout.compareTo(BigDecimal.ZERO) > 0) {
                bonusBalance = bonusBalance.min(maxCashout);
            }

            walletService.convertBonusToReal(bonus.getUserId(), bonusBalance);
        }

        playerBonusRepository.save(bonus);
        log.info("Bonus completed: {} for user {}", bonus.getId(), bonus.getUserId());
    }

    public double getGameWeight(String gameCategory) {
        return switch (gameCategory.toLowerCase()) {
            case "slots" -> 1.0;
            case "table_games", "blackjack" -> 0.1;
            case "roulette" -> 0.2;
            case "poker" -> 0.05;
            default -> 1.0;
        };
    }
}
