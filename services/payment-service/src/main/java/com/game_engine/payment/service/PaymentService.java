package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.model.Withdrawal.*;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.util.Map;
import java.util.UUID;

/**
 * Payment Service
 *
 * Facade coordinating deposit and withdrawal processing.
 * Delegates to specialized sub-services for each concern.
 */
@Service
@RequiredArgsConstructor
@Slf4j
public class PaymentService {

    private final DepositService depositService;
    private final WithdrawalService withdrawalService;
    private final PaymentValidationService paymentValidationService;
    private final GatewayResolver gatewayResolver;

    @Transactional
    public Deposit processDeposit(UUID userId, BigDecimal amount, String currency,
                                  PaymentGateway gateway, PaymentMethod paymentMethod,
                                  Map<String, String> paymentDetails, String ipAddress) {
        return depositService.processDeposit(userId, amount, currency, gateway, paymentMethod, paymentDetails, ipAddress);
    }

    @Transactional
    public Deposit handleDepositCallback(PaymentGateway gateway, Map<String, String> callbackData) {
        return depositService.handleDepositCallback(gateway, callbackData);
    }

    @Transactional
    public Withdrawal processWithdrawal(UUID userId, BigDecimal amount, String currency,
                                        PaymentGateway gateway, PaymentMethod paymentMethod,
                                        Map<String, String> payoutDetails) {
        return withdrawalService.processWithdrawal(userId, amount, currency, gateway, paymentMethod, payoutDetails);
    }

    @Transactional
    public Withdrawal approveWithdrawal(UUID withdrawalId, UUID adminUserId) {
        return withdrawalService.approveWithdrawal(withdrawalId, adminUserId);
    }

    @Transactional
    public Withdrawal rejectWithdrawal(UUID withdrawalId, UUID adminUserId, String reason) {
        return withdrawalService.rejectWithdrawal(withdrawalId, adminUserId, reason);
    }
}
