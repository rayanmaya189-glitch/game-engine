package com.game_engine.payment.service;

import com.game_engine.payment.model.Payment;
import com.game_engine.payment.repository.PaymentRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class PaymentService {

    private final PaymentRepository paymentRepository;
    private final PaymentGatewayService paymentGatewayService;

    @Value("${payment.limits.min-deposit:10.00}")
    private BigDecimal minDeposit;

    @Value("${payment.limits.max-deposit:50000.00}")
    private BigDecimal maxDeposit;

    @Value("${payment.limits.min-withdrawal:20.00}")
    private BigDecimal minWithdrawal;

    @Value("${payment.limits.max-withdrawal:50000.00}")
    private BigDecimal maxWithdrawal;

    @Value("${payment.currency.default:USD}")
    private String defaultCurrency;

    @Transactional
    public Payment createDeposit(String userId, BigDecimal amount, Payment.PaymentMethod method, 
                                   String currency, String description) {
        validateDepositAmount(amount);
        
        String externalId = "DEP-" + UUID.randomUUID().toString();
        
        Payment payment = Payment.builder()
                .userId(userId)
                .externalId(externalId)
                .type(Payment.PaymentType.DEPOSIT)
                .method(method)
                .status(Payment.PaymentStatus.PENDING)
                .amount(amount)
                .currency(currency != null ? currency : defaultCurrency)
                .description(description)
                .expiresAt(Instant.now().plus(30, ChronoUnit.MINUTES))
                .retries(0)
                .build();

        Payment savedPayment = paymentRepository.save(payment);
        log.info("Created deposit payment: {} for user: {} amount: {}", externalId, userId, amount);

        return savedPayment;
    }

    @Transactional
    public Payment processDeposit(UUID paymentId) {
        Payment payment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (payment.getStatus() != Payment.PaymentStatus.PENDING) {
            throw new IllegalStateException("Payment is not in pending status");
        }

        payment.setStatus(Payment.PaymentStatus.PROCESSING);
        payment = paymentRepository.save(payment);

        try {
            PaymentGatewayService.PaymentGatewayResponse response = paymentGatewayService.processDeposit(payment);

            if (response.isSuccess()) {
                payment.setStatus(Payment.PaymentStatus.COMPLETED);
                payment.setCompletedAt(Instant.now());
                payment.setGatewayResponse(response.getMessage());
            } else {
                payment.setStatus(Payment.PaymentStatus.FAILED);
                payment.setFailureReason(response.getMessage());
            }
        } catch (Exception e) {
            log.error("Error processing deposit: {}", e.getMessage());
            payment.setStatus(Payment.PaymentStatus.FAILED);
            payment.setFailureReason(e.getMessage());
        }

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment createWithdrawal(String userId, BigDecimal amount, Payment.PaymentMethod method,
                                     String currency, String accountDetails, String description) {
        validateWithdrawalAmount(amount);
        
        String externalId = "WDR-" + UUID.randomUUID().toString();
        
        // Calculate fee (example: 2% for most methods, 1% for crypto)
        BigDecimal fee = calculateFee(amount, method);
        BigDecimal netAmount = amount.subtract(fee);

        Payment payment = Payment.builder()
                .userId(userId)
                .externalId(externalId)
                .type(Payment.PaymentType.WITHDRAWAL)
                .method(method)
                .status(Payment.PaymentStatus.PENDING)
                .amount(amount)
                .currency(currency != null ? currency : defaultCurrency)
                .fee(fee)
                .netAmount(netAmount)
                .description(description)
                .metadata(accountDetails)
                .expiresAt(Instant.now().plus(24, ChronoUnit.HOURS))
                .retries(0)
                .build();

        Payment savedPayment = paymentRepository.save(payment);
        log.info("Created withdrawal payment: {} for user: {} amount: {}", externalId, userId, amount);

        return savedPayment;
    }

    @Transactional
    public Payment processWithdrawal(UUID paymentId) {
        Payment payment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (payment.getStatus() != Payment.PaymentStatus.PENDING) {
            throw new IllegalStateException("Payment is not in pending status");
        }

        payment.setStatus(Payment.PaymentStatus.PROCESSING);
        payment = paymentRepository.save(payment);

        try {
            PaymentGatewayService.PaymentGatewayResponse response = paymentGatewayService.processWithdrawal(payment);

            if (response.isSuccess()) {
                payment.setStatus(Payment.PaymentStatus.COMPLETED);
                payment.setCompletedAt(Instant.now());
                payment.setGatewayResponse(response.getMessage());
            } else {
                payment.setStatus(Payment.PaymentStatus.FAILED);
                payment.setFailureReason(response.getMessage());
            }
        } catch (Exception e) {
            log.error("Error processing withdrawal: {}", e.getMessage());
            payment.setStatus(Payment.PaymentStatus.FAILED);
            payment.setFailureReason(e.getMessage());
        }

        return paymentRepository.save(payment);
    }

    @Transactional
    public Payment refundPayment(UUID paymentId, BigDecimal amount) {
        Payment originalPayment = paymentRepository.findById(paymentId)
                .orElseThrow(() -> new IllegalArgumentException("Payment not found: " + paymentId));

        if (originalPayment.getStatus() != Payment.PaymentStatus.COMPLETED) {
            throw new IllegalStateException("Can only refund completed payments");
        }

        if (amount.compareTo(originalPayment.getAmount()) > 0) {
            throw new IllegalArgumentException("Refund amount cannot exceed original payment amount");
        }

        String externalId = "REF-" + UUID.randomUUID().toString();
        
        Payment refundPayment = Payment.builder()
                .userId(originalPayment.getUserId())
                .externalId(externalId)
                .type(Payment.PaymentType.REFUND)
                .method(originalPayment.getMethod())
                .status(Payment.PaymentStatus.PENDING)
                .amount(amount)
                .currency(originalPayment.getCurrency())
                .description("Refund for payment: " + originalPayment.getExternalId())
                .metadata(paymentId.toString())
                .build();

        try {
            PaymentGatewayService.PaymentGatewayResponse response = 
                    paymentGatewayService.processRefund(originalPayment, amount);

            if (response.isSuccess()) {
                refundPayment.setStatus(Payment.PaymentStatus.COMPLETED);
                refundPayment.setCompletedAt(Instant.now());
                
                originalPayment.setStatus(Payment.PaymentStatus.REFUNDED);
                paymentRepository.save(originalPayment);
            } else {
                refundPayment.setStatus(Payment.PaymentStatus.FAILED);
                refundPayment.setFailureReason(response.getMessage());
            }
        } catch (Exception e) {
            log.error("Error processing refund: {}", e.getMessage());
            refundPayment.setStatus(Payment.PaymentStatus.FAILED);
            refundPayment.setFailureReason(e.getMessage());
        }

        return paymentRepository.save(refundPayment);
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

    public List<Payment> getUserPaymentsByStatus(String userId, Payment.PaymentStatus status) {
        return paymentRepository.findByUserIdAndStatus(userId, status);
    }

    public BigDecimal getUserTotalDeposits(String userId, Instant startDate, Instant endDate) {
        BigDecimal total = paymentRepository.sumByUserIdAndTypeAndStatusAndDateRange(
                userId, Payment.PaymentType.DEPOSIT, startDate, endDate);
        return total != null ? total : BigDecimal.ZERO;
    }

    public BigDecimal getUserTotalWithdrawals(String userId, Instant startDate, Instant endDate) {
        BigDecimal total = paymentRepository.sumByUserIdAndTypeAndStatusAndDateRange(
                userId, Payment.PaymentType.WITHDRAWAL, startDate, endDate);
        return total != null ? total : BigDecimal.ZERO;
    }

    @Scheduled(fixedRate = 60000) // Run every minute
    @Transactional
    public void processExpiredPayments() {
        List<Payment> expiredPayments = paymentRepository.findExpiredPendingPayments(Instant.now());
        
        for (Payment payment : expiredPayments) {
            log.info("Expiring payment: {}", payment.getExternalId());
            payment.setStatus(Payment.PaymentStatus.EXPIRED);
            paymentRepository.save(payment);
        }
    }

    @Scheduled(fixedRate = 300000) // Run every 5 minutes
    @Transactional
    public void retryFailedPayments() {
        List<Payment> pendingPayments = paymentRepository.findByTypeAndDateRange(
                Payment.PaymentType.DEPOSIT, Instant.now().minus(1, ChronoUnit.HOURS));
        
        for (Payment payment : pendingPayments) {
            if (payment.getStatus() == Payment.PaymentStatus.FAILED && payment.getRetries() < 3) {
                log.info("Retrying payment: {}", payment.getExternalId());
                payment.setRetries(payment.getRetries() + 1);
                
                try {
                    PaymentGatewayService.PaymentGatewayResponse response = 
                            paymentGatewayService.processDeposit(payment);
                    
                    if (response.isSuccess()) {
                        payment.setStatus(Payment.PaymentStatus.COMPLETED);
                        payment.setCompletedAt(Instant.now());
                    } else {
                        payment.setFailureReason(response.getMessage());
                    }
                } catch (Exception e) {
                    log.error("Error retrying payment: {}", e.getMessage());
                    payment.setFailureReason(e.getMessage());
                }
                
                paymentRepository.save(payment);
            }
        }
    }

    private void validateDepositAmount(BigDecimal amount) {
        if (amount.compareTo(minDeposit) < 0) {
            throw new IllegalArgumentException("Minimum deposit amount is " + minDeposit);
        }
        if (amount.compareTo(maxDeposit) > 0) {
            throw new IllegalArgumentException("Maximum deposit amount is " + maxDeposit);
        }
    }

    private void validateWithdrawalAmount(BigDecimal amount) {
        if (amount.compareTo(minWithdrawal) < 0) {
            throw new IllegalArgumentException("Minimum withdrawal amount is " + minWithdrawal);
        }
        if (amount.compareTo(maxWithdrawal) > 0) {
            throw new IllegalArgumentException("Maximum withdrawal amount is " + maxWithdrawal);
        }
    }

    private BigDecimal calculateFee(BigDecimal amount, Payment.PaymentMethod method) {
        // Example fee structure - should be configurable
        return switch (method) {
            case BITCOIN, ETHEREUM -> amount.multiply(BigDecimal.valueOf(0.01)); // 1% for crypto
            case BANK_TRANSFER -> amount.multiply(BigDecimal.valueOf(0.005)).max(BigDecimal.valueOf(5)); // 0.5% min $5
            default -> amount.multiply(BigDecimal.valueOf(0.025)); // 2.5% for e-wallets/cards
        };
    }

    public List<String> getSupportedCurrencies() {
        return List.of("USD", "EUR", "GBP", "BTC", "ETH");
    }

    public List<Payment.PaymentMethod> getSupportedMethods(Payment.PaymentType type) {
        return switch (type) {
            case DEPOSIT -> List.of(
                    Payment.PaymentMethod.CREDIT_CARD,
                    Payment.PaymentMethod.DEBIT_CARD,
                    Payment.PaymentMethod.PAYPAL,
                    Payment.PaymentMethod.SKRILL,
                    Payment.PaymentMethod.NETELLER,
                    Payment.PaymentMethod.BITCOIN,
                    Payment.PaymentMethod.ETHEREUM,
                    Payment.PaymentMethod.BANK_TRANSFER,
                    Payment.PaymentMethod.INSTANT_BANK_TRANSFER
            );
            case WITHDRAWAL -> List.of(
                    Payment.PaymentMethod.CREDIT_CARD,
                    Payment.PaymentMethod.DEBIT_CARD,
                    Payment.PaymentMethod.PAYPAL,
                    Payment.PaymentMethod.SKRILL,
                    Payment.PaymentMethod.NETELLER,
                    Payment.PaymentMethod.BITCOIN,
                    Payment.PaymentMethod.ETHEREUM,
                    Payment.PaymentMethod.BANK_TRANSFER
            );
            default -> List.of();
        };
    }
}
