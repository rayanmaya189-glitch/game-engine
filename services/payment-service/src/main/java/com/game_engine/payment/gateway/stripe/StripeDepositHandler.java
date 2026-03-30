package com.game_engine.payment.gateway.stripe;

import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.stripe.Stripe;
import com.stripe.exception.*;
import com.stripe.model.PaymentIntent;
import com.stripe.param.PaymentIntentCreateParams;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.Map;
import java.util.UUID;

@Component
@Slf4j
public class StripeDepositHandler {

    @Value("${payment.gateway.stripe.api-key:}")
    private String apiKey;

    @Value("${payment.gateway.stripe.webhook-secret:}")
    private String webhookSecret;

    private final StripeUtils stripeUtils = new StripeUtils();

    public GatewayResponse initiateDeposit(Deposit deposit, Map<String, String> paymentDetails) {
        try {
            stripeUtils.initializeStripe(apiKey);

            PaymentIntentCreateParams.Builder paramsBuilder = PaymentIntentCreateParams.builder()
                    .setAmount(stripeUtils.currencyToMinorUnits(deposit.getAmount(), deposit.getCurrency()))
                    .setCurrency(deposit.getCurrency().toLowerCase())
                    .setMetadata(Map.of(
                        "deposit_id", deposit.getId().toString(),
                        "user_id", deposit.getUserId().toString()
                    ))
                    .setCaptureMethod(PaymentIntentCreateParams.CaptureMethod.AUTOMATIC);

            String paymentMethodId = paymentDetails.get("payment_method_id");
            if (paymentMethodId != null) {
                paramsBuilder.setPaymentMethod(paymentMethodId)
                        .setConfirm(true)
                        .setReturnUrl(paymentDetails.get("return_url"));
            } else {
                paramsBuilder.setPaymentMethodTypes(java.util.List.of("card"));
            }

            PaymentIntent intent = PaymentIntent.create(paramsBuilder.build());

            if ("requires_action".equals(intent.getStatus()) ||
                "requires_source_action".equals(intent.getStatus())) {
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

    public GatewayResponse processDepositCallback(Map<String, String> callbackData) {
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

    public GatewayResponse checkDepositStatus(UUID depositId, String gatewayTransactionId) {
        try {
            stripeUtils.initializeStripe(apiKey);
            PaymentIntent intent = PaymentIntent.retrieve(gatewayTransactionId);

            String status = stripeUtils.mapStripeStatus(intent.getStatus());

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

    public boolean healthCheck() {
        try {
            stripeUtils.initializeStripe(apiKey);
            return true;
        } catch (Exception e) {
            log.error("Stripe health check failed: {}", e.getMessage());
            return false;
        }
    }
}
