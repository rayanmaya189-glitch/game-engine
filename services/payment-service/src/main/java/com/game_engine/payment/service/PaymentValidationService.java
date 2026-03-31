package com.game_engine.payment.service;

import com.game_engine.payment.grpc.BonusGrpcClient;
import com.game_engine.payment.grpc.RiskGrpcClient;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class PaymentValidationService {

    private final BonusGrpcClient bonusGrpcClient;
    private final RiskGrpcClient riskGrpcClient;

    @Value("${payment.limits.deposit.min:10}")
    private BigDecimal minDeposit;

    @Value("${payment.limits.deposit.max:10000}")
    private BigDecimal maxDeposit;

    @Value("${payment.limits.withdrawal.min:20}")
    private BigDecimal minWithdrawal;

    @Value("${payment.limits.withdrawal.max:50000}")
    private BigDecimal maxWithdrawal;

    @Value("${payment.risk.default-score:0}")
    private int defaultRiskScore;

    @Value("${payment.wagering.default-met:true}")
    private boolean defaultWageringMet;

    public void validateDepositAmount(BigDecimal amount) {
        if (amount.compareTo(minDeposit) < 0) {
            throw new IllegalArgumentException("Minimum deposit is " + minDeposit);
        }
        if (amount.compareTo(maxDeposit) > 0) {
            throw new IllegalArgumentException("Maximum deposit is " + maxDeposit);
        }
    }

    public void validateWithdrawalAmount(BigDecimal amount) {
        if (amount.compareTo(minWithdrawal) < 0) {
            throw new IllegalArgumentException("Minimum withdrawal is " + minWithdrawal);
        }
        if (amount.compareTo(maxWithdrawal) > 0) {
            throw new IllegalArgumentException("Maximum withdrawal is " + maxWithdrawal);
        }
    }

    public boolean checkWageringRequirements(UUID userId, BigDecimal amount) {
        try {
            return bonusGrpcClient.checkWageringRequirements(userId, amount.doubleValue());
        } catch (Exception e) {
            log.error("Failed to check wagering requirements for user {}: {}", userId, e.getMessage());
            return defaultWageringMet;
        }
    }

    public int getRiskScore(UUID userId) {
        try {
            return riskGrpcClient.getRiskScore(userId);
        } catch (Exception e) {
            log.error("Failed to get risk score for user {}: {}", userId, e.getMessage());
            return defaultRiskScore;
        }
    }
}
