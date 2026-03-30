package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import com.game_engine.payment.model.Payment;
import com.game_engine.payment.model.Withdrawal;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@Service
@Slf4j
public class PaymentGatewayService {

    private final GatewayResolver gatewayResolver;
    private final CurrencyExchangeService currencyExchangeService;

    public PaymentGatewayService(GatewayResolver gatewayResolver,
                                  CurrencyExchangeService currencyExchangeService) {
        this.gatewayResolver = gatewayResolver;
        this.currencyExchangeService = currencyExchangeService;
    }

    public PaymentGatewayResponse processDeposit(Deposit deposit, Map<String, String> paymentDetails) {
        log.info("Processing deposit via gateway: {} amount: {} {}",
                deposit.getGateway(), deposit.getAmount(), deposit.getCurrency());

        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(deposit.getGateway());
        GatewayResponse response = adapter.initiateDeposit(deposit, paymentDetails);

        return PaymentGatewayResponse.builder()
                .success(response.isSuccess())
                .externalId(deposit.getId().toString())
                .status(mapToPaymentStatus(response.getStatus()))
                .message(response.getMessage())
                .gatewayTransactionId(response.getGatewayTransactionId())
                .redirectUrl(response.getRedirectUrl())
                .metadata(response.getMetadata() != null
                        ? Map.copyOf(response.getMetadata())
                        : Map.of())
                .build();
    }

    public PaymentGatewayResponse processWithdrawal(Withdrawal withdrawal, Map<String, String> payoutDetails) {
        log.info("Processing withdrawal via gateway: {} amount: {} {}",
                withdrawal.getGateway(), withdrawal.getAmount(), withdrawal.getCurrency());

        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(withdrawal.getGateway());
        GatewayResponse response = adapter.initiateWithdrawal(withdrawal, payoutDetails);

        return PaymentGatewayResponse.builder()
                .success(response.isSuccess())
                .externalId(withdrawal.getId().toString())
                .status(mapToPaymentStatus(response.getStatus()))
                .message(response.getMessage())
                .gatewayTransactionId(response.getGatewayTransactionId())
                .build();
    }

    public PaymentGatewayResponse checkDepositStatus(Deposit deposit) {
        log.info("Checking deposit status: {} gateway: {}", deposit.getId(), deposit.getGateway());

        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(deposit.getGateway());
        GatewayResponse response = adapter.checkDepositStatus(deposit.getId(), deposit.getGatewayTransactionId());

        return PaymentGatewayResponse.builder()
                .success(response.isSuccess())
                .externalId(deposit.getId().toString())
                .status(mapToPaymentStatus(response.getStatus()))
                .message(response.getMessage())
                .gatewayTransactionId(response.getGatewayTransactionId())
                .build();
    }

    public PaymentGatewayResponse checkWithdrawalStatus(Withdrawal withdrawal) {
        log.info("Checking withdrawal status: {} gateway: {}", withdrawal.getId(), withdrawal.getGateway());

        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(withdrawal.getGateway());
        GatewayResponse response = adapter.checkWithdrawalStatus(withdrawal.getId(), withdrawal.getGatewayTransactionId());

        return PaymentGatewayResponse.builder()
                .success(response.isSuccess())
                .externalId(withdrawal.getId().toString())
                .status(mapToPaymentStatus(response.getStatus()))
                .message(response.getMessage())
                .gatewayTransactionId(response.getGatewayTransactionId())
                .build();
    }

    public PaymentGatewayResponse processRefund(Deposit originalDeposit, BigDecimal amount) {
        log.info("Processing refund for deposit: {} amount: {}", originalDeposit.getId(), amount);

        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(originalDeposit.getGateway());
        Map<String, String> refundDetails = Map.of(
                "original_transaction_id", originalDeposit.getGatewayTransactionId(),
                "amount", amount.toPlainString()
        );

        GatewayResponse response = adapter.processDepositCallback(Map.of(
                "type", "refund",
                "data.object.id", originalDeposit.getGatewayTransactionId(),
                "data.object.amount", amount.toPlainString(),
                "data.object.status", "refunded"
        ));

        return PaymentGatewayResponse.builder()
                .success(response.isSuccess())
                .externalId(originalDeposit.getId().toString())
                .status(Payment.PaymentStatus.REFUNDED)
                .message("Refund processed successfully")
                .gatewayTransactionId(response.getGatewayTransactionId())
                .build();
    }

    public boolean verifyWebhookSignature(PaymentGateway gateway, String payload, String signature) {
        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(gateway);
        return adapter.verifyWebhookSignature(payload, signature);
    }

    public boolean healthCheck(PaymentGateway gateway) {
        try {
            PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(gateway);
            return adapter.healthCheck();
        } catch (Exception e) {
            log.error("Health check failed for gateway {}: {}", gateway, e.getMessage());
            return false;
        }
    }

    public List<PaymentGateway> getAvailableGateways() {
        return gatewayResolver.getAvailableAdapters().stream()
                .filter(adapter -> {
                    try {
                        return adapter.healthCheck();
                    } catch (Exception e) {
                        return false;
                    }
                })
                .map(PaymentGatewayAdapter::getGatewayType)
                .toList();
    }

    public BigDecimal convertCurrency(BigDecimal amount, String fromCurrency, String toCurrency) {
        return currencyExchangeService.convert(amount, fromCurrency, toCurrency);
    }

    private Payment.PaymentStatus mapToPaymentStatus(String status) {
        if (status == null) return Payment.PaymentStatus.PENDING;
        return switch (status.toUpperCase()) {
            case "COMPLETED", "SUCCEEDED", "PAID" -> Payment.PaymentStatus.COMPLETED;
            case "PROCESSING", "IN_TRANSIT" -> Payment.PaymentStatus.PROCESSING;
            case "PENDING", "PENDING_VERIFICATION", "REQUIRES_ACTION" -> Payment.PaymentStatus.PENDING;
            case "FAILED" -> Payment.PaymentStatus.FAILED;
            case "CANCELLED", "CANCELED" -> Payment.PaymentStatus.CANCELLED;
            case "REFUNDED" -> Payment.PaymentStatus.REFUNDED;
            case "EXPIRED" -> Payment.PaymentStatus.EXPIRED;
            default -> Payment.PaymentStatus.PENDING;
        };
    }

    @lombok.Data
    @lombok.Builder
    public static class PaymentGatewayResponse {
        private boolean success;
        private String externalId;
        private Payment.PaymentStatus status;
        private String message;
        private String gatewayTransactionId;
        private String redirectUrl;
        private Map<String, Object> metadata;
    }
}
