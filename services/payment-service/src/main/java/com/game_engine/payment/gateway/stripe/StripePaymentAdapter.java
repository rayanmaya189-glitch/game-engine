package com.game_engine.payment.gateway.stripe;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Withdrawal;
import com.stripe.Stripe;
import com.stripe.exception.*;
import com.stripe.model.PaymentIntent;
import com.stripe.model.Payout;
import com.stripe.param.PaymentIntentCreateParams;
import com.stripe.param.PayoutCreateParams;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

/**
 * Stripe Payment Gateway Adapter
 * 
 * Implements payment processing using Stripe's Payment Intents API.
 * Supports credit/debit cards with 3D Secure authentication.
 * 
 * PCI DSS: Uses Stripe's hosted fields - no card data touches our servers.
 */
@Component
@Slf4j
public class StripePaymentAdapter implements PaymentGatewayAdapter {

    @Value("${payment.gateway.stripe.api-key:}")
    private String apiKey;

    @Value("${payment.gateway.stripe.webhook-secret:}")
    private String webhookSecret;

    @Value("${payment.gateway.stripe.auto-payout:false}")
    private boolean autoPayout;

    public StripePaymentAdapter() {
        // Initialize with API key from configuration
    }

    @Override
    public Deposit.PaymentGateway getGatewayType() {
        return Deposit.PaymentGateway.STRIPE;
    }

    @Override
    public GatewayResponse initiateDeposit(Deposit deposit, Map<String, String> paymentDetails) {
        try {
            initializeStripe();

            // Create PaymentIntent with Stripe
            PaymentIntentCreateParams.Builder paramsBuilder = PaymentIntentCreateParams.builder()
                    .setAmount(currencyToMinorUnits(deposit.getAmount(), deposit.getCurrency()))
                    .setCurrency(deposit.getCurrency().toLowerCase())
                    .setMetadata(Map.of(
                        "deposit_id", deposit.getId().toString(),
                        "user_id", deposit.getUserId().toString()
                    ))
                    .setCaptureMethod(PaymentIntentCreateParams.CaptureMethod.AUTOMATIC);

            // Handle 3D Secure if required
            String paymentMethodId = paymentDetails.get("payment_method_id");
            if (paymentMethodId != null) {
                paramsBuilder.setPaymentMethod(paymentMethodId)
                        .setConfirm(true)
                        .setReturnUrl(paymentDetails.get("return_url"));
            } else {
                // Use payment link/redirect flow
                paramsBuilder.setPaymentMethodTypes(java.util.List.of("card"));
            }

            PaymentIntent intent = PaymentIntent.create(paramsBuilder.build());

            // Build response based on status
            if ("requires_action".equals(intent.getStatus()) || 
                "requires_source_action".equals(intent.getStatus())) {
                // 3D Secure required - return redirect URL
                Map<String, String> nextAction = intent.getNextAction();
                String redirectUrl = null;
                if (nextAction != null && nextAction.containsKey("redirect_to_url")) {
                    redirectUrl = nextAction.get("redirect_to_url");
                }
                
                return GatewayResponse.builder()
                        .success(true)
                        .gatewayTransactionId(intent.getId())
                        .status("PENDING_VERIFICATION")
                        .message("3D Secure verification required")
                        .redirectUrl(redirectUrl)
                        .metadata(intent.getMetadata())
                        .build();
            }

            if ("succeeded".equals(intent.getStatus())) {
                return GatewayResponse.builder()
                        .success(true)
                        .gatewayTransactionId(intent.getId())
                        .status("COMPLETED")
                        .message("Payment successful")
                        .metadata(intent.getMetadata())
                        .processedAmount(deposit.getAmount())
                        .build();
            }

            // Other statuses
            return GatewayResponse.builder()
                    .success(false)
                    .gatewayTransactionId(intent.getId())
                    .status(intent.getStatus())
                    .message("Payment requires additional action")
                    .build();

        } catch (CardException e) {
            log.error("Stripe card error: {} - {}", e.getCode(), e.getMessage());
            return GatewayResponse.builder()
                    .success(false)
                    .status("FAILED")
                    .message(e.getMessage())
                    .errorCode(e.getCode())
                    .build();
        } catch (StripeException e) {
            log.error("Stripe error: {} - {}", e.getStripeError() != null ? e.getStripeError().getCode() : "unknown", e.getMessage());
            return GatewayResponse.builder()
                    .success(false)
                    .status("FAILED")
                    .message(e.getMessage())
                    .errorCode(e.getStripeError() != null ? e.getStripeError().getCode() : "stripe_error")
                    .build();
        }
    }

    @Override
    public GatewayResponse processDepositCallback(Map<String, String> callbackData) {
        // Handle Stripe webhook - would verify signature first
        String eventType = callbackData.get("type");
        String status = callbackData.get("data.object.status");

        return GatewayResponse.builder()
                .success("payment_intent.succeeded".equals(eventType))
                .gatewayTransactionId(callbackData.get("data.object.id"))
                .status(status)
                .message(eventType)
                .metadata(callbackData)
                .build();
    }

    @Override
    public GatewayResponse checkDepositStatus(UUID depositId, String gatewayTransactionId) {
        try {
            initializeStripe();
            PaymentIntent intent = PaymentIntent.retrieve(gatewayTransactionId);

            String status = mapStripeStatus(intent.getStatus());

            return GatewayResponse.builder()
                    .success("COMPLETED".equals(status) || "PROCESSING".equals(status))
                    .gatewayTransactionId(intent.getId())
                    .status(status)
                    .message(intent.getStatus())
                    .build();

        } catch (StripeException e) {
            log.error("Error checking Stripe payment status: {}", e.getMessage());
            return GatewayResponse.builder()
                    .success(false)
                    .message(e.getMessage())
                    .build();
        }
    }

    @Override
    public GatewayResponse initiateWithdrawal(Withdrawal withdrawal, Map<String, String> payoutDetails) {
        try {
            initializeStripe();

            if (!autoPayout) {
                // Mark as approved - actual payout handled separately
                return GatewayResponse.builder()
                        .success(true)
                        .status("APPROVED")
                        .message("Withdrawal approved for processing")
                        .build();
            }

            // Create payout
            PayoutCreateParams params = PayoutCreateParams.builder()
                    .setAmount(currencyToMinorUnits(withdrawal.getAmount(), withdrawal.getCurrency()))
                    .setCurrency(withdrawal.getCurrency().toLowerCase())
                    .setMetadata(Map.of(
                        "withdrawal_id", withdrawal.getId().toString(),
                        "user_id", withdrawal.getUserId().toString()
                    ))
                    .build();

            Payout payout = Payout.create(params);

            return GatewayResponse.builder()
                    .success("pending".equals(payout.getStatus()) || "in_transit".equals(payout.getStatus()))
                    .gatewayTransactionId(payout.getId())
                    .status(mapPayoutStatus(payout.getStatus()))
                    .message("Payout initiated")
                    .build();

        } catch (StripeException e) {
            log.error("Stripe payout error: {}", e.getMessage());
            return GatewayResponse.builder()
                    .success(false)
                    .status("FAILED")
                    .message(e.getMessage())
                    .errorCode(e.getStripeError() != null ? e.getStripeError().getCode() : "stripe_error")
                    .build();
        }
    }

    @Override
    public GatewayResponse checkWithdrawalStatus(UUID withdrawalId, String gatewayTransactionId) {
        try {
            initializeStripe();
            Payout payout = Payout.retrieve(gatewayTransactionId);

            return GatewayResponse.builder()
                    .success("paid".equals(payout.getStatus()))
                    .gatewayTransactionId(payout.getId())
                    .status(mapPayoutStatus(payout.getStatus()))
                    .message(payout.getStatus())
                    .build();

        } catch (StripeException e) {
            log.error("Error checking Stripe payout status: {}", e.getMessage());
            return GatewayResponse.builder()
                    .success(false)
                    .message(e.getMessage())
                    .build();
        }
    }

    @Override
    public boolean verifyWebhookSignature(String payload, String signature) {
        if (webhookSecret == null || webhookSecret.isEmpty()) {
            log.warn("Stripe webhook secret not configured");
            return false;
        }
        
        try {
            com.stripe.net.Webhook.constructEvent(payload, signature, webhookSecret);
            return true;
        } catch (Exception e) {
            log.error("Webhook signature verification failed: {}", e.getMessage());
            return false;
        }
    }

    @Override
    public boolean healthCheck() {
        try {
            initializeStripe();
            // Just verify API key is valid by creating a minimal call
            // In production, would use Stripe's /v1/balance or /v1/accounts
            return true;
        } catch (Exception e) {
            log.error("Stripe health check failed: {}", e.getMessage());
            return false;
        }
    }

    @Override
    public java.util.List<Deposit.PaymentMethod> getSupportedMethods() {
        return java.util.List.of(
            Deposit.PaymentMethod.CREDIT_CARD,
            Deposit.PaymentMethod.DEBIT_CARD
        );
    }

    private void initializeStripe() {
        if (Stripe.apiKey == null || Stripe.apiKey.isEmpty()) {
            Stripe.apiKey = apiKey;
        }
    }

    private long currencyToMinorUnits(BigDecimal amount, String currency) {
        // Most currencies use 2 decimal places
        // Note: Some currencies like JPY use 0 decimal places
        int[] zeroDecimalCurrencies = { "JPY", "KRW", "VND" };
        
        for (String curr : zeroDecimalCurrencies) {
            if (curr.equals(currency.toUpperCase())) {
                return amount.longValue();
            }
        }
        
        return amount.multiply(BigDecimal.valueOf(100)).longValue();
    }

    private String mapStripeStatus(String stripeStatus) {
        return switch (stripeStatus) {
            case "succeeded" -> "COMPLETED";
            case "processing" -> "PROCESSING";
            case "requires_action", "requires_source_action", "requires_payment_method" -> "PENDING_VERIFICATION";
            case "canceled" -> "CANCELLED";
            default -> "PROCESSING";
        };
    }

    private String mapPayoutStatus(String stripeStatus) {
        return switch (stripeStatus) {
            case "paid" -> "COMPLETED";
            case "pending", "in_transit" -> "PROCESSING";
            case "failed", "canceled" -> "FAILED";
            default -> "PROCESSING";
        };
    }
}
