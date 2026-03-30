package com.game_engine.payment.service;

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

    @Value("${payment.limits.deposit.min:10}")
    private BigDecimal minDeposit;

    @Value("${payment.limits.deposit.max:10000}")
    private BigDecimal maxDeposit;

    @Value("${payment.limits.withdrawal.min:20}")
    private BigDecimal minWithdrawal;

    @Value("${payment.limits.withdrawal.max:50000}")
    private BigDecimal maxWithdrawal;

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
        // Would check with Bonus Service
        // Simplified for now - assume requirements met
        return true;
    }

    public int getRiskScore(UUID userId) {
        // Would call Risk Scoring Service via gRPC
        // Simplified for now
        return 0;
    }
}
