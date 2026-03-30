package com.game_engine.payment.gateway.stripe;

import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Withdrawal;
import com.stripe.Stripe;
import com.stripe.exception.StripeException;
import com.stripe.model.Payout;
import com.stripe.param.PayoutCreateParams;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.Map;
import java.util.UUID;

@Component
@Slf4j
public class StripeWithdrawalHandler {

    @Value("${payment.gateway.stripe.api-key:}")
    private String apiKey;

    @Value("${payment.gateway.stripe.auto-payout:false}")
    private boolean autoPayout;

    private final StripeUtils stripeUtils = new StripeUtils();

    public GatewayResponse initiateWithdrawal(Withdrawal withdrawal, Map<String, String> payoutDetails) {
        try {
            stripeUtils.initializeStripe(apiKey);

            if (!autoPayout) {
                return GatewayResponse.builder()
                        .success(true)
                        .status("APPROVED")
                        .message("Withdrawal approved for processing")
                        .build();
            }

            PayoutCreateParams params = PayoutCreateParams.builder()
                    .setAmount(stripeUtils.currencyToMinorUnits(withdrawal.getAmount(), withdrawal.getCurrency()))
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
                    .status(stripeUtils.mapPayoutStatus(payout.getStatus()))
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

    public GatewayResponse checkWithdrawalStatus(UUID withdrawalId, String gatewayTransactionId) {
        try {
            stripeUtils.initializeStripe(apiKey);
            Payout payout = Payout.retrieve(gatewayTransactionId);

            return GatewayResponse.builder()
                    .success("paid".equals(payout.getStatus()))
                    .gatewayTransactionId(payout.getId())
                    .status(stripeUtils.mapPayoutStatus(payout.getStatus()))
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
}
