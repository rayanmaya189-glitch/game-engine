package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import com.game_engine.payment.model.Payment;
import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.model.Withdrawal.*;
import com.game_engine.payment.repository.PaymentRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class PaymentService {

    private final DepositService depositService;
    private final WithdrawalService withdrawalService;
    private final PaymentValidationService paymentValidationService;
    private final GatewayResolver gatewayResolver;
    private final PaymentRepository paymentRepository;
    private final WalletService walletService;

    @Transactional
    public Deposit processDeposit(UUID userId, BigDecimal amount, String currency,
                                  PaymentGateway gateway, Deposit.PaymentMethod paymentMethod,
                                  Map<String, String> paymentDetails, String ipAddress) {
        return depositService.processDeposit(userId, amount, currency, gateway, paymentMethod, paymentDetails, ipAddress);
    }

    @Transactional
    public Payment createDeposit(String userId, BigDecimal amount, Payment.PaymentMethod method,
                                 String currency, String description) {
        paymentValidationService.validateDepositAmount(amount);

        Payment payment = Payment.builder()
                .userId(userId)
                .externalId(UUID.randomUUID().toString())
                .type(Payment.PaymentType.DEPOSIT)
                .method(method)
                .status(Payment.PaymentStatus.PENDING)
                .amount(amount)
                .currency(currency != null ? currency.toUpperCase() : "USD")
                .description(description)
                .build();

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment processDeposit(UUID paymentId) {
        Payment payment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (payment.getStatus() != Payment.PaymentStatus.PENDING) {
            throw new IllegalStateException("Payment is not in PENDING status");
        }

        payment.setStatus(Payment.PaymentStatus.PROCESSING);
        payment.setProcessedAt(Instant.now());

        try {
            Deposit.PaymentGateway gateway = mapPaymentMethodToGateway(payment.getMethod());
            Deposit.PaymentMethod depositMethod = mapToDepositPaymentMethod(payment.getMethod());

            depositService.processDeposit(
                    UUID.fromString(payment.getUserId()),
                    payment.getAmount(),
                    payment.getCurrency(),
                    gateway,
                    depositMethod,
                    Map.of("payment_token", payment.getExternalId()),
                    null
            );

            payment.setStatus(Payment.PaymentStatus.COMPLETED);
            payment.setCompletedAt(Instant.now());
        } catch (Exception e) {
            log.error("Failed to process deposit {}: {}", paymentId, e.getMessage());
            payment.setStatus(Payment.PaymentStatus.FAILED);
            payment.setFailureReason(e.getMessage());
        }

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment createWithdrawal(String userId, BigDecimal amount, Payment.PaymentMethod method,
                                    String currency, String accountDetails, String description) {
        paymentValidationService.validateWithdrawalAmount(amount);

        Payment payment = Payment.builder()
                .userId(userId)
                .externalId(UUID.randomUUID().toString())
                .type(Payment.PaymentType.WITHDRAWAL)
                .method(method)
                .status(Payment.PaymentStatus.PENDING)
                .amount(amount)
                .currency(currency != null ? currency.toUpperCase() : "USD")
                .description(description)
                .build();

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment processWithdrawal(UUID paymentId) {
        Payment payment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (payment.getStatus() != Payment.PaymentStatus.PENDING) {
            throw new IllegalStateException("Payment is not in PENDING status");
        }

        payment.setStatus(Payment.PaymentStatus.PROCESSING);
        payment.setProcessedAt(Instant.now());

        try {
            UUID userId = UUID.fromString(payment.getUserId());
            BigDecimal balance = walletService.getBalance(userId, payment.getCurrency());
            if (balance.compareTo(payment.getAmount()) < 0) {
                throw new IllegalArgumentException("Insufficient balance");
            }

            walletService.debitBalance(userId, payment.getAmount(), payment.getCurrency(),
                    "WITHDRAWAL", payment.getExternalId());

            payment.setStatus(Payment.PaymentStatus.COMPLETED);
            payment.setCompletedAt(Instant.now());
        } catch (Exception e) {
            log.error("Failed to process withdrawal {}: {}", paymentId, e.getMessage());
            payment.setStatus(Payment.PaymentStatus.FAILED);
            payment.setFailureReason(e.getMessage());
        }

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment refundPayment(UUID paymentId, BigDecimal amount) {
        Payment payment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (payment.getStatus() != Payment.PaymentStatus.COMPLETED) {
            throw new IllegalStateException("Only completed payments can be refunded");
        }

        if (amount.compareTo(payment.getAmount()) > 0) {
            throw new IllegalArgumentException("Refund amount exceeds original payment amount");
        }

        walletService.creditBalance(
                UUID.fromString(payment.getUserId()),
                amount,
                payment.getCurrency(),
                "REFUND",
                payment.getExternalId()
        );

        payment.setStatus(amount.compareTo(payment.getAmount()) == 0
                ? Payment.PaymentStatus.REFUNDED
                : Payment.PaymentStatus.PARTIALLY_REFUNDED);

        return paymentRepository.save(payment);
    }

    @Transactional
    public Deposit handleDepositCallback(PaymentGateway gateway, Map<String, String> callbackData) {
        return depositService.handleDepositCallback(gateway, callbackData);
    }

    @Transactional
    public Withdrawal processWithdrawal(UUID userId, BigDecimal amount, String currency,
                                        PaymentGateway gateway, Deposit.PaymentMethod paymentMethod,
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

    public Optional<Payment> getPayment(UUID paymentId) {
        return paymentRepository.findById(paymentId);
    }

    public Optional<Payment> getPaymentByExternalId(String externalId) {
        return paymentRepository.findByExternalId(externalId);
    }

    public Page<Payment> getUserPayments(String userId, Pageable pageable) {
        return paymentRepository.findByUserId(userId, pageable);
    }

    public BigDecimal getUserTotalDeposits(String userId, Instant startDate, Instant endDate) {
        List<Payment> payments = paymentRepository.findByUserIdAndTypeAndStatusAndDateRange(
                userId, Payment.PaymentType.DEPOSIT, Payment.PaymentStatus.COMPLETED, startDate, endDate);
        return payments.stream()
                .map(Payment::getAmount)
                .reduce(BigDecimal.ZERO, BigDecimal::add);
    }

    public BigDecimal getUserTotalWithdrawals(String userId, Instant startDate, Instant endDate) {
        List<Payment> payments = paymentRepository.findByUserIdAndTypeAndStatusAndDateRange(
                userId, Payment.PaymentType.WITHDRAWAL, Payment.PaymentStatus.COMPLETED, startDate, endDate);
        return payments.stream()
                .map(Payment::getAmount)
                .reduce(BigDecimal.ZERO, BigDecimal::add);
    }

    public List<Payment.PaymentMethod> getSupportedMethods(Payment.PaymentType type) {
        return gatewayResolver.getAvailableAdapters().stream()
                .flatMap(adapter -> mapGatewayMethodsToPaymentMethods(adapter.getSupportedMethods()).stream())
                .distinct()
                .toList();
    }

    public List<String> getSupportedCurrencies() {
        return List.of("USD", "EUR", "GBP", "JPY", "CNY", "KRW", "THB", "VND",
                "IDR", "MYR", "SGD", "AUD", "CAD", "CHF", "INR", "BRL",
                "MXN", "ZAR", "RUB", "AED", "BTC", "ETH", "USDT");
    }

    private Deposit.PaymentGateway mapPaymentMethodToGateway(Payment.PaymentMethod method) {
        return switch (method) {
            case CREDIT_CARD, DEBIT_CARD, VIRTUAL_CARD -> Deposit.PaymentGateway.STRIPE;
            case PAYPAL -> Deposit.PaymentGateway.STRIPE;
            case SKRILL -> Deposit.PaymentGateway.SKRILL;
            case NETELLER -> Deposit.PaymentGateway.NETELLER;
            case BITCOIN, ETHEREUM -> Deposit.PaymentGateway.COINBASE_COMMERCE;
            case BANK_TRANSFER, INSTANT_BANK_TRANSFER -> Deposit.PaymentGateway.LOCAL_BANK_TRANSFER;
            default -> Deposit.PaymentGateway.STRIPE;
        };
    }

    private Deposit.PaymentMethod mapToDepositPaymentMethod(Payment.PaymentMethod method) {
        return switch (method) {
            case CREDIT_CARD, DEBIT_CARD, VIRTUAL_CARD -> Deposit.PaymentMethod.CREDIT_CARD;
            case BANK_TRANSFER, INSTANT_BANK_TRANSFER -> Deposit.PaymentMethod.BANK_TRANSFER;
            case PAYPAL, SKRILL, NETELLER, MUCHBETTER, ECOPAYZ -> Deposit.PaymentMethod.E_WALLET;
            case BITCOIN, ETHEREUM -> Deposit.PaymentMethod.CRYPTOCURRENCY;
            case ASTROPAY, JETON -> Deposit.PaymentMethod.PREPAID_CARD;
            default -> Deposit.PaymentMethod.CREDIT_CARD;
        };
    }

    private List<Payment.PaymentMethod> mapGatewayMethodsToPaymentMethods(List<Deposit.PaymentMethod> gatewayMethods) {
        return gatewayMethods.stream().map(gm -> switch (gm) {
            case CREDIT_CARD -> Payment.PaymentMethod.CREDIT_CARD;
            case DEBIT_CARD -> Payment.PaymentMethod.DEBIT_CARD;
            case BANK_TRANSFER -> Payment.PaymentMethod.BANK_TRANSFER;
            case E_WALLET -> Payment.PaymentMethod.PAYPAL;
            case CRYPTOCURRENCY -> Payment.PaymentMethod.BITCOIN;
            case PREPAID_CARD -> Payment.PaymentMethod.ASTROPAY;
        }).toList();
    }
}
