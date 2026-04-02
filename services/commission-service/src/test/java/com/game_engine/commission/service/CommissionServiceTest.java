package com.game_engine.commission.service;

import com.game_engine.commission.entity.CommissionClaim;
import com.game_engine.commission.entity.CommissionConfig;
import com.game_engine.commission.entity.CommissionSettlement;
import com.game_engine.commission.repository.CommissionClaimRepository;
import com.game_engine.commission.repository.CommissionConfigRepository;
import com.game_engine.commission.repository.CommissionSettlementRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class CommissionServiceTest {

    @Mock
    private CommissionConfigRepository configRepository;

    @Mock
    private CommissionClaimRepository claimRepository;

    @Mock
    private CommissionSettlementRepository settlementRepository;

    @InjectMocks
    private CommissionService commissionService;

    private CommissionConfig revenueShareConfig;
    private CommissionConfig cpaConfig;

    @BeforeEach
    void setUp() {
        revenueShareConfig = new CommissionConfig();
        revenueShareConfig.setId(1L);
        revenueShareConfig.setAgentId(10L);
        revenueShareConfig.setAffiliateId(20L);
        revenueShareConfig.setCommissionType("REVENUE_SHARE");
        revenueShareConfig.setRate(new BigDecimal("0.30"));
        revenueShareConfig.setMaxCommission(new BigDecimal("10000"));
        revenueShareConfig.setIsActive(true);

        cpaConfig = new CommissionConfig();
        cpaConfig.setId(2L);
        cpaConfig.setAgentId(10L);
        cpaConfig.setAffiliateId(20L);
        cpaConfig.setCommissionType("CPA");
        cpaConfig.setRate(new BigDecimal("50.00"));
        cpaConfig.setIsActive(true);
    }

    @Nested
    @DisplayName("calculateCommission")
    class CalculateCommission {

        @Test
        @DisplayName("should calculate revenue share commission")
        void shouldCalculateRevenueShare() {
            when(configRepository.findByAgentIdAndAffiliateIdAndIsActive(10L, 20L, true))
                    .thenReturn(List.of(revenueShareConfig));

            BigDecimal result = commissionService.calculateCommission(10L, 20L, new BigDecimal("1000"), 0);

            assertEquals(0, new BigDecimal("300.00").compareTo(result));
        }

        @Test
        @DisplayName("should cap commission at max amount")
        void shouldCapCommission() {
            when(configRepository.findByAgentIdAndAffiliateIdAndIsActive(10L, 20L, true))
                    .thenReturn(List.of(revenueShareConfig));

            BigDecimal result = commissionService.calculateCommission(10L, 20L, new BigDecimal("100000"), 0);

            assertEquals(0, new BigDecimal("10000").compareTo(result));
        }

        @Test
        @DisplayName("should calculate CPA commission")
        void shouldCalculateCPA() {
            when(configRepository.findByAgentIdAndAffiliateIdAndIsActive(10L, 20L, true))
                    .thenReturn(List.of(cpaConfig));

            BigDecimal result = commissionService.calculateCommission(10L, 20L, BigDecimal.ZERO, 10);

            assertEquals(0, new BigDecimal("500.00").compareTo(result));
        }

        @Test
        @DisplayName("should throw when no config found")
        void shouldThrowWhenNoConfig() {
            when(configRepository.findByAgentIdAndAffiliateIdAndIsActive(anyLong(), anyLong(), anyBoolean()))
                    .thenReturn(List.of());

            assertThrows(RuntimeException.class, () ->
                    commissionService.calculateCommission(99L, 99L, BigDecimal.ONE, 0));
        }
    }

    @Nested
    @DisplayName("createClaim")
    class CreateClaim {

        @Test
        @DisplayName("should create a pending claim")
        void shouldCreatePendingClaim() {
            when(configRepository.findByAgentIdAndAffiliateIdAndIsActive(10L, 20L, true))
                    .thenReturn(List.of(revenueShareConfig));
            when(claimRepository.save(any(CommissionClaim.class))).thenAnswer(i -> i.getArgument(0));

            CommissionClaim claim = commissionService.createClaim(10L, 20L, 100L, "2024-01", new BigDecimal("5000"));

            assertNotNull(claim);
            assertEquals("PENDING", claim.getStatus());
            assertEquals(10L, claim.getAgentId());
            assertEquals(20L, claim.getAffiliateId());
            verify(claimRepository).save(any(CommissionClaim.class));
        }
    }

    @Nested
    @DisplayName("approveClaim")
    class ApproveClaim {

        @Test
        @DisplayName("should approve a pending claim")
        void shouldApprovePendingClaim() {
            CommissionClaim claim = new CommissionClaim();
            claim.setId(1L);
            claim.setStatus("PENDING");

            when(claimRepository.findById(1L)).thenReturn(Optional.of(claim));
            when(claimRepository.save(any())).thenReturn(claim);

            CommissionClaim result = commissionService.approveClaim(1L);

            assertEquals("APPROVED", result.getStatus());
            assertNotNull(result.getProcessedAt());
        }

        @Test
        @DisplayName("should throw when claim not found")
        void shouldThrowWhenNotFound() {
            when(claimRepository.findById(999L)).thenReturn(Optional.empty());

            assertThrows(RuntimeException.class, () -> commissionService.approveClaim(999L));
        }

        @Test
        @DisplayName("should throw when claim not pending")
        void shouldThrowWhenNotPending() {
            CommissionClaim claim = new CommissionClaim();
            claim.setStatus("APPROVED");
            when(claimRepository.findById(1L)).thenReturn(Optional.of(claim));

            assertThrows(RuntimeException.class, () -> commissionService.approveClaim(1L));
        }
    }

    @Nested
    @DisplayName("processSettlement")
    class ProcessSettlement {

        @Test
        @DisplayName("should process settlement for approved claims")
        void shouldProcessSettlement() {
            CommissionClaim claim1 = new CommissionClaim();
            claim1.setGrossRevenue(new BigDecimal("1000"));
            claim1.setCommissionAmount(new BigDecimal("300"));
            claim1.setStatus("APPROVED");

            CommissionClaim claim2 = new CommissionClaim();
            claim2.setGrossRevenue(new BigDecimal("2000"));
            claim2.setCommissionAmount(new BigDecimal("600"));
            claim2.setStatus("APPROVED");

            when(claimRepository.findByAffiliateIdAndStatus(20L, "APPROVED"))
                    .thenReturn(List.of(claim1, claim2));
            when(claimRepository.save(any())).thenReturn(claim1);
            when(settlementRepository.save(any(CommissionSettlement.class)))
                    .thenAnswer(i -> i.getArgument(0));

            LocalDateTime start = LocalDateTime.of(2024, 1, 1, 0, 0);
            LocalDateTime end = LocalDateTime.of(2024, 1, 31, 23, 59);

            CommissionSettlement settlement = commissionService.processSettlement(10L, 20L, start, end);

            assertNotNull(settlement);
            assertEquals("PAID", settlement.getStatus());
            assertEquals(0, new BigDecimal("3000").compareTo(settlement.getTotalRevenue()));
            assertEquals(0, new BigDecimal("900").compareTo(settlement.getTotalCommission()));
        }
    }
}
