package com.game_engine.payment.gateway;

import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Withdrawal;
import java.math.BigDecimal;
import java.util.Map;
import java.util.UUID;

/**
 * Payment Gateway Adapter Interface
 * 
 * Defines the contract for all payment gateway implementations.
 * Uses the Adapter pattern to support multiple gateways uniformly.
 * 
 * All implementations must be thread-safe as they may be called
 * from multiple threads simultaneously.
 */
public interface PaymentGatewayAdapter {

    /**
     * Get the gateway type this adapter handles
     */
    Deposit.PaymentGateway getGatewayType();

    /**
     * Initialize a deposit transaction
     * 
     * @param deposit The deposit entity to process
     * @param paymentDetails Tokenized payment details from the client
     * @return GatewayResponse containing the result and any redirect URL
     */
    GatewayResponse initiateDeposit(Deposit deposit, Map<String, String> paymentDetails);

    /**
     * Process a deposit callback/webhook from the gateway
     * 
     * @param callbackData Raw callback data from gateway
     * @return VerificationResult with the verified transaction status
     */
    GatewayResponse processDepositCallback(Map<String, String> callbackData);

    /**
     * Check deposit status with the gateway
     * 
     * @param depositId Our internal deposit ID
     * @param gatewayTransactionId The gateway's transaction ID
     * @return Current status from the gateway
     */
    GatewayResponse checkDepositStatus(UUID depositId, String gatewayTransactionId);

    /**
     * Initiate a withdrawal/payout
     * 
     * @param withdrawal The withdrawal entity to process
     * @param payoutDetails Encoded payout details
     * @return GatewayResponse containing the result
     */
    GatewayResponse initiateWithdrawal(Withdrawal withdrawal, Map<String, String> payoutDetails);

    /**
     * Check withdrawal status with the gateway
     * 
     * @param withdrawalId Our internal withdrawal ID
     * @param gatewayTransactionId The gateway's transaction ID
     * @return Current status from the gateway
     */
    GatewayResponse checkWithdrawalStatus(UUID withdrawalId, String gatewayTransactionId);

    /**
     * Verify webhook signature from the gateway
     * 
     * @param payload Raw request body
     * @param signature Signature from headers
     * @return true if signature is valid
     */
    boolean verifyWebhookSignature(String payload, String signature);

    /**
     * Check if gateway is available and healthy
     * 
     * @return true if gateway is operational
     */
    boolean healthCheck();

    /**
     * Get supported payment methods for this gateway
     */
    default java.util.List<Deposit.PaymentMethod> getSupportedMethods() {
        return java.util.List.of(
            Deposit.PaymentMethod.CREDIT_CARD,
            Deposit.PaymentMethod.DEBIT_CARD
        );
    }

    /**
     * Response from gateway operations
     */
    class GatewayResponse {
        private final boolean success;
        private final String gatewayTransactionId;
        private final String status;
        private final String message;
        private final String redirectUrl;
        private final Map<String, String> metadata;
        private final BigDecimal processedAmount;
        private final String errorCode;

        private GatewayResponse(Builder builder) {
            this.success = builder.success;
            this.gatewayTransactionId = builder.gatewayTransactionId;
            this.status = builder.status;
            this.message = builder.message;
            this.redirectUrl = builder.redirectUrl;
            this.metadata = builder.metadata;
            this.processedAmount = builder.processedAmount;
            this.errorCode = builder.errorCode;
        }

        public boolean isSuccess() { return success; }
        public String getGatewayTransactionId() { return gatewayTransactionId; }
        public String getStatus() { return status; }
        public String getMessage() { return message; }
        public String getRedirectUrl() { return redirectUrl; }
        public Map<String, String> getMetadata() { return metadata; }
        public BigDecimal getProcessedAmount() { return processedAmount; }
        public String getErrorCode() { return errorCode; }

        public static Builder builder() {
            return new Builder();
        }

        public static class Builder {
            private boolean success;
            private String gatewayTransactionId;
            private String status;
            private String message;
            private String redirectUrl;
            private Map<String, String> metadata;
            private BigDecimal processedAmount;
            private String errorCode;

            public Builder success(boolean success) { this.success = success; return this; }
            public Builder gatewayTransactionId(String id) { this.gatewayTransactionId = id; return this; }
            public Builder status(String status) { this.status = status; return this; }
            public Builder message(String message) { this.message = message; return this; }
            public Builder redirectUrl(String url) { this.redirectUrl = url; return this; }
            public Builder metadata(Map<String, String> m) { this.metadata = m; return this; }
            public Builder processedAmount(BigDecimal amount) { this.processedAmount = amount; return this; }
            public Builder errorCode(String code) { this.errorCode = code; return this; }

            public GatewayResponse build() {
                return new GatewayResponse(this);
            }
        }
    }
}
