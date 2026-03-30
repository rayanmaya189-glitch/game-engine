package com.game_engine.payment.gateway.stripe;

import com.stripe.Stripe;

import java.math.BigDecimal;

public class StripeUtils {

    public void initializeStripe(String apiKey) {
        if (Stripe.apiKey == null || Stripe.apiKey.isEmpty()) {
            Stripe.apiKey = apiKey;
        }
    }

    public long currencyToMinorUnits(BigDecimal amount, String currency) {
        int[] zeroDecimalCurrencies = { "JPY", "KRW", "VND" };

        for (String curr : zeroDecimalCurrencies) {
            if (curr.equals(currency.toUpperCase())) {
                return amount.longValue();
            }
        }

        return amount.multiply(BigDecimal.valueOf(100)).longValue();
    }

    public String mapStripeStatus(String stripeStatus) {
        return switch (stripeStatus) {
            case "succeeded" -> "COMPLETED";
            case "processing" -> "PROCESSING";
            case "requires_action", "requires_source_action", "requires_payment_method" -> "PENDING_VERIFICATION";
            case "canceled" -> "CANCELLED";
            default -> "PROCESSING";
        };
    }

    public String mapPayoutStatus(String stripeStatus) {
        return switch (stripeStatus) {
            case "paid" -> "COMPLETED";
            case "pending", "in_transit" -> "PROCESSING";
            case "failed", "canceled" -> "FAILED";
            default -> "PROCESSING";
        };
    }
}
