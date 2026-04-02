package com.game_engine.bonus.service;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class BonusServiceTest {

    @Mock
    private BonusEligibilityService eligibilityService;

    @Mock
    private BonusWageringService wageringService;

    @Mock
    private BonusCashbackService cashbackService;

    private BonusService bonusService;

    @BeforeEach
    void setUp() {
        bonusService = new BonusService(eligibilityService, wageringService, cashbackService);
    }

    @Nested
    @DisplayName("getAvailableBonuses")
    class GetAvailableBonuses {

        @Test
        @DisplayName("should return available bonuses for user")
        void shouldReturnAvailableBonuses() {
            UUID userId = UUID.randomUUID();
            BigDecimal depositAmount = new BigDecimal("100.00");

            BonusCampaign campaign = BonusCampaign.builder()
                    .id(UUID.randomUUID())
                    .name("Welcome Bonus")
                    .isActive(true)
                    .matchPercentage(new BigDecimal("100"))
                    .maxBonusAmount(new BigDecimal("500"))
                    .build();

            when(eligibilityService.getAvailableBonuses(userId, depositAmount))
                    .thenReturn(List.of(campaign));

            List<BonusCampaign> result = bonusService.getAvailableBonuses(userId, depositAmount);

            assertEquals(1, result.size());
            assertEquals("Welcome Bonus", result.get(0).getName());
            verify(eligibilityService).getAvailableBonuses(userId, depositAmount);
        }

        @Test
        @DisplayName("should return empty list when no bonuses available")
        void shouldReturnEmptyWhenNoBonuses() {
            UUID userId = UUID.randomUUID();
            when(eligibilityService.getAvailableBonuses(any(), any())).thenReturn(List.of());

            List<BonusCampaign> result = bonusService.getAvailableBonuses(userId, BigDecimal.ZERO);

            assertTrue(result.isEmpty());
        }
    }

    @Nested
    @DisplayName("activateBonus")
    class ActivateBonus {

        @Test
        @DisplayName("should activate a pending bonus")
        void shouldActivatePendingBonus() {
            UUID bonusId = UUID.randomUUID();
            PlayerBonus bonus = PlayerBonus.builder()
                    .id(bonusId)
                    .status(PlayerBonus.BonusStatus.ACTIVE)
                    .build();

            when(eligibilityService.activateBonus(bonusId)).thenReturn(bonus);

            PlayerBonus result = bonusService.activateBonus(bonusId);

            assertNotNull(result);
            assertEquals(PlayerBonus.BonusStatus.ACTIVE, result.getStatus());
            verify(eligibilityService).activateBonus(bonusId);
        }
    }

    @Nested
    @DisplayName("processWagering")
    class ProcessWagering {

        @Test
        @DisplayName("should process bet for wagering")
        void shouldProcessBet() {
            UUID userId = UUID.randomUUID();
            BigDecimal betAmount = new BigDecimal("50.00");

            bonusService.processBet(userId, betAmount, "game-1", "slots");

            verify(wageringService).processBet(userId, betAmount, "game-1", "slots");
        }

        @Test
        @DisplayName("should handle zero bet amount")
        void shouldHandleZeroBet() {
            UUID userId = UUID.randomUUID();

            bonusService.processBet(userId, BigDecimal.ZERO, "game-1", "slots");

            verify(wageringService).processBet(userId, BigDecimal.ZERO, "game-1", "slots");
        }
    }

    @Nested
    @DisplayName("calculateCashback")
    class CalculateCashback {

        @Test
        @DisplayName("should calculate cashback for user")
        void shouldCalculateCashback() {
            UUID userId = UUID.randomUUID();
            BigDecimal expectedCashback = new BigDecimal("20.00");

            when(cashbackService.calculateCashback(userId, 7)).thenReturn(expectedCashback);

            BigDecimal result = bonusService.calculateCashback(userId, 7);

            assertEquals(0, expectedCashback.compareTo(result));
        }

        @Test
        @DisplayName("should return zero when no losses")
        void shouldReturnZeroWhenNoLosses() {
            UUID userId = UUID.randomUUID();
            when(cashbackService.calculateCashback(userId, 30)).thenReturn(BigDecimal.ZERO);

            BigDecimal result = bonusService.calculateCashback(userId, 30);

            assertEquals(0, BigDecimal.ZERO.compareTo(result));
        }
    }
}
