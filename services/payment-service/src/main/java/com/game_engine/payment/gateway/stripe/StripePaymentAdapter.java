package com.game_engine.payment.gateway.stripe;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Withdrawal;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Map;
import java.util.UUID;

/**
 * Stripe Payment Gateway Adapter
 *
 * Delegates to specialized handlers for deposits and withdrawals.
 * PCI DSS: Uses Stripe's hosted fields - no card data touches our servers.
 */
@Component
@RequiredArgsConstructor
@Slf4j
public class StripePaymentAdapter implements PaymentGatewayAdapter {

    private final StripeDepositHandler depositHandler;
    private final StripeWithdrawalHandler withdrawalHandler;

    @Override
    public Deposit.PaymentGateway getGatewayType() {
        return Deposit.PaymentGateway.STRIPE;
    }

    @Override
    public GatewayResponse initiateDeposit(Deposit deposit, Map<String, String> paymentDetails) {
        return depositHandler.initiateDeposit(deposit, paymentDetails);
    }

    @Override
    public GatewayResponse processDepositCallback(Map<String, String> callbackData) {
        return depositHandler.processDepositCallback(callbackData);
    }

    @Override
    public GatewayResponse checkDepositStatus(UUID depositId, String gatewayTransactionId) {
        return depositHandler.checkDepositStatus(depositId, gatewayTransactionId);
    }

    @Override
    public GatewayResponse initiateWithdrawal(Withdrawal withdrawal, Map<String, String> payoutDetails) {
        return withdrawalHandler.initiateWithdrawal(withdrawal, payoutDetails);
    }

    @Override
    public GatewayResponse checkWithdrawalStatus(UUID withdrawalId, String gatewayTransactionId) {
        return withdrawalHandler.checkWithdrawalStatus(withdrawalId, gatewayTransactionId);
    }

    @Override
    public boolean verifyWebhookSignature(String payload, String signature) {
        return depositHandler.verifyWebhookSignature(payload, signature);
    }

    @Override
    public boolean healthCheck() {
        return depositHandler.healthCheck();
    }

    @Override
    public List<Deposit.PaymentMethod> getSupportedMethods() {
        return List.of(
            Deposit.PaymentMethod.CREDIT_CARD,
            Deposit.PaymentMethod.DEBIT_CARD
        );
    }
}
