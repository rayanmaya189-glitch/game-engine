package com.game_engine.payment.grpc;

import com.game_engine.common.v1.Money;
import com.game_engine.wallet.v1.*;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

@Component
@Slf4j
public class WalletGrpcClient {

    @Value("${wallet.service.host:wallet-service}")
    private String walletServiceHost;

    @Value("${wallet.service.grpc.port:9009}")
    private int walletServicePort;

    @Value("${wallet.service.timeout-seconds:10}")
    private long timeoutSeconds;

    private ManagedChannel channel;
    private WalletServiceGrpc.WalletServiceBlockingStub blockingStub;

    @PostConstruct
    public void init() {
        log.info("Connecting to Wallet Service at {}:{}", walletServiceHost, walletServicePort);
        channel = ManagedChannelBuilder
                .forAddress(walletServiceHost, walletServicePort)
                .usePlaintext()
                .build();
        blockingStub = WalletServiceGrpc.newBlockingStub(channel);
    }

    @PreDestroy
    public void shutdown() {
        if (channel != null && !channel.isShutdown()) {
            try {
                channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
                channel.shutdownNow();
                Thread.currentThread().interrupt();
            }
        }
    }

    public void creditBalance(UUID userId, BigDecimal amount, String currency,
                              String transactionType, String referenceId) {
        log.info("Crediting {} {} to user {} via gRPC - type: {} ref: {}",
                amount, currency, userId, transactionType, referenceId);

        try {
            Money money = buildMoney(amount, currency);

            if ("DEPOSIT".equals(transactionType)) {
                ConfirmDepositRequest request = ConfirmDepositRequest.newBuilder()
                        .setTransactionId(referenceId)
                        .setPaymentReference(referenceId)
                        .setProviderStatus("COMPLETED")
                        .build();

                ConfirmDepositResponse response = blockingStub
                        .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                        .confirmDeposit(request);

                if (!response.getSuccess()) {
                    throw new WalletOperationException("Failed to credit balance: " + response.getMessage());
                }
                log.info("Balance credited successfully for user {}: {}", userId, response.getMessage());
            } else {
                CreateDepositRequest request = CreateDepositRequest.newBuilder()
                        .setUserId(userId.toString())
                        .setAmount(money)
                        .build();

                CreateDepositResponse response = blockingStub
                        .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                        .createDeposit(request);

                log.info("Deposit created for user {}: {}", userId, response.getPaymentReference());
            }
        } catch (StatusRuntimeException e) {
            log.error("gRPC error crediting balance for user {}: {}", userId, e.getStatus());
            throw new WalletOperationException("Wallet service unavailable: " + e.getStatus().getDescription(), e);
        }
    }

    public void debitBalance(UUID userId, BigDecimal amount, String currency,
                             String transactionType, String referenceId) {
        log.info("Debiting {} {} from user {} via gRPC - type: {} ref: {}",
                amount, currency, userId, transactionType, referenceId);

        try {
            Money money = buildMoney(amount, currency);

            if ("WITHDRAWAL".equals(transactionType) || "WITHDRAWAL_REVERSAL".equals(transactionType)) {
                ConfirmWithdrawalRequest request = ConfirmWithdrawalRequest.newBuilder()
                        .setTransactionId(referenceId)
                        .setStatus(com.game_engine.common.v1.TransactionStatus.TRANSACTION_STATUS_COMPLETED)
                        .build();

                ConfirmWithdrawalResponse response = blockingStub
                        .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                        .confirmWithdrawal(request);

                if (!response.getSuccess()) {
                    throw new WalletOperationException("Failed to debit balance: " + response.getMessage());
                }
                log.info("Balance debited successfully for user {}: {}", userId, response.getMessage());
            } else {
                CreateWithdrawalRequest request = CreateWithdrawalRequest.newBuilder()
                        .setUserId(userId.toString())
                        .setAmount(money)
                        .setWithdrawalMethodId(referenceId)
                        .build();

                CreateWithdrawalResponse response = blockingStub
                        .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                        .createWithdrawal(request);

                log.info("Withdrawal created for user {}: {}", userId, response.getMessage());
            }
        } catch (StatusRuntimeException e) {
            log.error("gRPC error debiting balance for user {}: {}", userId, e.getStatus());
            throw new WalletOperationException("Wallet service unavailable: " + e.getStatus().getDescription(), e);
        }
    }

    public BigDecimal getBalance(UUID userId, String currency) {
        log.info("Getting balance for user {} currency {} via gRPC", userId, currency);

        try {
            GetBalanceRequest request = GetBalanceRequest.newBuilder()
                    .setUserId(userId.toString())
                    .setBalanceType(com.game_engine.common.v1.BalanceType.BALANCE_TYPE_REAL)
                    .build();

            GetBalanceResponse response = blockingStub
                    .withDeadlineAfter(timeoutSeconds, TimeUnit.SECONDS)
                    .getBalance(request);

            Money balance = response.getBalance();
            return fromMinorUnits(balance.getAmount(), currency);
        } catch (StatusRuntimeException e) {
            log.error("gRPC error getting balance for user {}: {}", userId, e.getStatus());
            throw new WalletOperationException("Wallet service unavailable: " + e.getStatus().getDescription(), e);
        }
    }

    private Money buildMoney(BigDecimal amount, String currency) {
        long minorUnits = toMinorUnits(amount, currency);
        com.game_engine.common.v1.Currency currencyEnum = mapCurrency(currency);

        return Money.newBuilder()
                .setAmount(minorUnits)
                .setCurrency(currencyEnum)
                .setDisplayAmount(amount.toPlainString())
                .build();
    }

    private long toMinorUnits(BigDecimal amount, String currency) {
        int exponent = getCurrencyExponent(currency);
        return amount.multiply(BigDecimal.TEN.pow(exponent)).setScale(0, RoundingMode.HALF_UP).longValue();
    }

    private BigDecimal fromMinorUnits(long minorUnits, String currency) {
        int exponent = getCurrencyExponent(currency);
        return BigDecimal.valueOf(minorUnits).divide(BigDecimal.TEN.pow(exponent), exponent, RoundingMode.HALF_UP);
    }

    private int getCurrencyExponent(String currency) {
        if (currency == null) return 2;
        return switch (currency.toUpperCase()) {
            case "JPY", "KRW", "VND" -> 0;
            case "BTC", "ETH", "USDT" -> 8;
            default -> 2;
        };
    }

    private com.game_engine.common.v1.Currency mapCurrency(String currency) {
        if (currency == null) return com.game_engine.common.v1.Currency.CURRENCY_UNSPECIFIED;
        return switch (currency.toUpperCase()) {
            case "USD" -> com.game_engine.common.v1.Currency.CURRENCY_USD;
            case "EUR" -> com.game_engine.common.v1.Currency.CURRENCY_EUR;
            case "GBP" -> com.game_engine.common.v1.Currency.CURRENCY_GBP;
            case "JPY" -> com.game_engine.common.v1.Currency.CURRENCY_JPY;
            case "CNY" -> com.game_engine.common.v1.Currency.CURRENCY_CNY;
            case "KRW" -> com.game_engine.common.v1.Currency.CURRENCY_KRW;
            case "THB" -> com.game_engine.common.v1.Currency.CURRENCY_THB;
            case "VND" -> com.game_engine.common.v1.Currency.CURRENCY_VND;
            case "IDR" -> com.game_engine.common.v1.Currency.CURRENCY_IDR;
            case "MYR" -> com.game_engine.common.v1.Currency.CURRENCY_MYR;
            case "SGD" -> com.game_engine.common.v1.Currency.CURRENCY_SGD;
            case "AUD" -> com.game_engine.common.v1.Currency.CURRENCY_AUD;
            case "CAD" -> com.game_engine.common.v1.Currency.CURRENCY_CAD;
            case "CHF" -> com.game_engine.common.v1.Currency.CURRENCY_CHF;
            case "INR" -> com.game_engine.common.v1.Currency.CURRENCY_INR;
            case "BRL" -> com.game_engine.common.v1.Currency.CURRENCY_BRL;
            case "MXN" -> com.game_engine.common.v1.Currency.CURRENCY_MXN;
            case "ZAR" -> com.game_engine.common.v1.Currency.CURRENCY_ZAR;
            case "RUB" -> com.game_engine.common.v1.Currency.CURRENCY_RUB;
            case "AED" -> com.game_engine.common.v1.Currency.CURRENCY_AED;
            case "BTC" -> com.game_engine.common.v1.Currency.CURRENCY_BTC;
            case "ETH" -> com.game_engine.common.v1.Currency.CURRENCY_ETH;
            case "USDT" -> com.game_engine.common.v1.Currency.CURRENCY_USDT;
            default -> com.game_engine.common.v1.Currency.CURRENCY_UNSPECIFIED;
        };
    }

    public static class WalletOperationException extends RuntimeException {
        public WalletOperationException(String message) {
            super(message);
        }

        public WalletOperationException(String message, Throwable cause) {
            super(message, cause);
        }
    }
}
