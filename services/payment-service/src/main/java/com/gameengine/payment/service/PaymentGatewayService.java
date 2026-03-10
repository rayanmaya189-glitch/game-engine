package com.gameengine.payment.service;

import com.gameengine.payment.model.Payment;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;

import java.math.BigDecimal;
import java.util.Map;
import java.util.UUID;

@Service
@Slf4j
public class PaymentGatewayService {

    private final WebClient webClient;
    
    @Value("${payment.stripe.enabled:false}")
    private boolean stripeEnabled;
    
    @Value("${payment.paypal.enabled:false}")
    private boolean paypalEnabled;
    
    @Value("${payment.crypto.enabled:false}")
    private boolean cryptoEnabled;

    public PaymentGatewayService(WebClient.Builder webClientBuilder) {
        this.webClient = webClientBuilder.build();
    }

    public PaymentGatewayResponse processDeposit(Payment payment) {
        return switch (payment.getMethod()) {
            case CREDIT_CARD, DEBIT_CARD, VIRTUAL_CARD -> processCardDeposit(payment);
            case PAYPAL -> processPayPalDeposit(payment);
            case SKRILL -> processSkrillDeposit(payment);
            case NETELLER -> processNetellerDeposit(payment);
            case BITCOIN, ETHEREUM -> processCryptoDeposit(payment);
            case BANK_TRANSFER, INSTANT_BANK_TRANSFER -> processBankTransfer(payment);
            default -> processGenericDeposit(payment);
        };
    }

    public PaymentGatewayResponse processWithdrawal(Payment payment) {
        return switch (payment.getMethod()) {
            case CREDIT_CARD, DEBIT_CARD -> processCardWithdrawal(payment);
            case PAYPAL -> processPayPalWithdrawal(payment);
            case SKRILL -> processSkrillWithdrawal(payment);
            case NETELLER -> processNetellerWithdrawal(payment);
            case BITCOIN, ETHEREUM -> processCryptoWithdrawal(payment);
            case BANK_TRANSFER, INSTANT_BANK_TRANSFER -> processBankWithdrawal(payment);
            default -> processGenericWithdrawal(payment);
        };
    }

    public PaymentGatewayResponse processRefund(Payment originalPayment, BigDecimal amount) {
        log.info("Processing refund for payment: {} amount: {}", originalPayment.getExternalId(), amount);
        
        return PaymentGatewayResponse.builder()
                .success(true)
                .externalId(UUID.randomUUID().toString())
                .status(Payment.PaymentStatus.REFUNDED)
                .message("Refund processed successfully")
                .build();
    }

    public PaymentGatewayResponse checkPaymentStatus(String externalId, Payment.PaymentMethod method) {
        log.info("Checking payment status for: {} method: {}", externalId, method);
        
        return PaymentGatewayResponse.builder()
                .success(true)
                .externalId(externalId)
                .status(Payment.PaymentStatus.COMPLETED)
                .message("Payment confirmed")
                .build();
    }

    private PaymentGatewayResponse processCardDeposit(Payment payment) {
        log.info("Processing card deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        
        if (!stripeEnabled) {
            return createMockResponse(payment);
        }
        
        // Stripe integration would go here
        // This is a placeholder for actual Stripe API integration
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processPayPalDeposit(Payment payment) {
        log.info("Processing PayPal deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        
        if (!paypalEnabled) {
            return createMockResponse(payment);
        }
        
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processSkrillDeposit(Payment payment) {
        log.info("Processing Skrill deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processNetellerDeposit(Payment payment) {
        log.info("Processing Neteller deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processCryptoDeposit(Payment payment) {
        log.info("Processing crypto deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        
        if (!cryptoEnabled) {
            return createMockResponse(payment);
        }
        
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processBankTransfer(Payment payment) {
        log.info("Processing bank transfer: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processGenericDeposit(Payment payment) {
        log.info("Processing generic deposit: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processCardWithdrawal(Payment payment) {
        log.info("Processing card withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processPayPalWithdrawal(Payment payment) {
        log.info("Processing PayPal withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processSkrillWithdrawal(Payment payment) {
        log.info("Processing Skrill withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processNetellerWithdrawal(Payment payment) {
        log.info("Processing Neteller withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processCryptoWithdrawal(Payment payment) {
        log.info("Processing crypto withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processBankWithdrawal(Payment payment) {
        log.info("Processing bank withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse processGenericWithdrawal(Payment payment) {
        log.info("Processing generic withdrawal: {} amount: {}", payment.getExternalId(), payment.getAmount());
        return createMockResponse(payment);
    }

    private PaymentGatewayResponse createMockResponse(Payment payment) {
        return PaymentGatewayResponse.builder()
                .success(true)
                .externalId(payment.getExternalId())
                .status(Payment.PaymentStatus.COMPLETED)
                .message("Payment processed successfully")
                .gatewayTransactionId("GW-" + UUID.randomUUID().toString().substring(0, 8))
                .build();
    }

    public BigDecimal convertCurrency(BigDecimal amount, String fromCurrency, String toCurrency) {
        log.info("Converting {} {} to {}", amount, fromCurrency, toCurrency);
        // Mock exchange rates - in production, integrate with a currency exchange API
        Map<String, BigDecimal> rates = Map.of(
                "USD", BigDecimal.ONE,
                "EUR", BigDecimal.valueOf(0.92),
                "GBP", BigDecimal.valueOf(0.79),
                "BTC", BigDecimal.valueOf(0.000015),
                "ETH", BigDecimal.valueOf(0.00035)
        );
        
        BigDecimal fromRate = rates.getOrDefault(fromCurrency, BigDecimal.ONE);
        BigDecimal toRate = rates.getOrDefault(toCurrency, BigDecimal.ONE);
        
        return amount.multiply(toRate).divide(fromRate, 4, java.math.RoundingMode.HALF_UP);
    }

    @lombok.Data
    @lombok.Builder
    public static class PaymentGatewayResponse {
        private boolean success;
        private String externalId;
        private Payment.PaymentStatus status;
        private String message;
        private String gatewayTransactionId;
        private Map<String, Object> metadata;
    }
}
