package com.game_engine.payment.controller;

import com.game_engine.payment.model.Payment;
import com.game_engine.payment.service.PaymentService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Sort;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.math.BigDecimal;
import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/payments")
@RequiredArgsConstructor
@Slf4j
public class PaymentController {

    private final PaymentService paymentService;

    @PostMapping("/deposits")
    public ResponseEntity<Payment> createDeposit(@RequestBody DepositRequest request) {
        log.info("Creating deposit for user: {} amount: {}", request.userId(), request.amount());
        
        Payment payment = paymentService.createDeposit(
                request.userId(),
                request.amount(),
                request.method(),
                request.currency(),
                request.description()
        );
        
        return ResponseEntity.ok(payment);
    }

    @PostMapping("/deposits/{paymentId}/process")
    public ResponseEntity<Payment> processDeposit(@PathVariable UUID paymentId) {
        log.info("Processing deposit: {}", paymentId);
        
        Payment payment = paymentService.processDeposit(paymentId);
        return ResponseEntity.ok(payment);
    }

    @PostMapping("/withdrawals")
    public ResponseEntity<Payment> createWithdrawal(@RequestBody WithdrawalRequest request) {
        log.info("Creating withdrawal for user: {} amount: {}", request.userId(), request.amount());
        
        Payment payment = paymentService.createWithdrawal(
                request.userId(),
                request.amount(),
                request.method(),
                request.currency(),
                request.accountDetails(),
                request.description()
        );
        
        return ResponseEntity.ok(payment);
    }

    @PostMapping("/withdrawals/{paymentId}/process")
    public ResponseEntity<Payment> processWithdrawal(@PathVariable UUID paymentId) {
        log.info("Processing withdrawal: {}", paymentId);
        
        Payment payment = paymentService.processWithdrawal(paymentId);
        return ResponseEntity.ok(payment);
    }

    @PostMapping("/{paymentId}/refund")
    public ResponseEntity<Payment> refundPayment(
            @PathVariable UUID paymentId,
            @RequestBody RefundRequest request) {
        log.info("Refunding payment: {} amount: {}", paymentId, request.amount());
        
        Payment payment = paymentService.refundPayment(paymentId, request.amount());
        return ResponseEntity.ok(payment);
    }

    @GetMapping("/{paymentId}")
    public ResponseEntity<Payment> getPayment(@PathVariable UUID paymentId) {
        return paymentService.getPayment(paymentId)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    @GetMapping("/external/{externalId}")
    public ResponseEntity<Payment> getPaymentByExternalId(@PathVariable String externalId) {
        return paymentService.getPaymentByExternalId(externalId)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    @GetMapping("/users/{userId}")
    public ResponseEntity<Page<Payment>> getUserPayments(
            @PathVariable String userId,
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "20") int size,
            @RequestParam(defaultValue = "createdAt") String sortBy,
            @RequestParam(defaultValue = "desc") String sortDir) {
        
        Sort sort = sortDir.equalsIgnoreCase("asc") 
                ? Sort.by(sortBy).ascending() 
                : Sort.by(sortBy).descending();
        
        PageRequest pageRequest = PageRequest.of(page, size, sort);
        Page<Payment> payments = paymentService.getUserPayments(userId, pageRequest);
        
        return ResponseEntity.ok(payments);
    }

    @GetMapping("/users/{userId}/summary")
    public ResponseEntity<Map<String, BigDecimal>> getUserPaymentSummary(
            @PathVariable String userId,
            @RequestParam(required = false, defaultValue = "30") int days) {
        
        Instant startDate = Instant.now().minus(days, ChronoUnit.DAYS);
        Instant endDate = Instant.now();
        
        BigDecimal totalDeposits = paymentService.getUserTotalDeposits(userId, startDate, endDate);
        BigDecimal totalWithdrawals = paymentService.getUserTotalWithdrawals(userId, startDate, endDate);
        
        return ResponseEntity.ok(Map.of(
                "totalDeposits", totalDeposits,
                "totalWithdrawals", totalWithdrawals,
                "netAmount", totalDeposits.subtract(totalWithdrawals)
        ));
    }

    @GetMapping("/methods")
    public ResponseEntity<Map<String, List<String>>> getSupportedMethods() {
        return ResponseEntity.ok(Map.of(
                "deposit", paymentService.getSupportedMethods(Payment.PaymentType.DEPOSIT)
                        .stream().map(Enum::name).toList(),
                "withdrawal", paymentService.getSupportedMethods(Payment.PaymentType.WITHDRAWAL)
                        .stream().map(Enum::name).toList()
        ));
    }

    @GetMapping("/currencies")
    public ResponseEntity<List<String>> getSupportedCurrencies() {
        return ResponseEntity.ok(paymentService.getSupportedCurrencies());
    }

    @PostMapping("/{paymentId}/cancel")
    public ResponseEntity<Payment> cancelPayment(@PathVariable UUID paymentId) {
        log.info("Cancelling payment: {}", paymentId);
        
        return paymentService.getPayment(paymentId)
                .map(payment -> {
                    if (payment.getStatus() == Payment.PaymentStatus.PENDING) {
                        payment.setStatus(Payment.PaymentStatus.CANCELLED);
                        return ResponseEntity.ok(payment);
                    }
                    return ResponseEntity.badRequest().<Payment>build();
                })
                .orElse(ResponseEntity.notFound().build());
    }

    // Request DTOs
    public record DepositRequest(
            String userId,
            BigDecimal amount,
            Payment.PaymentMethod method,
            String currency,
            String description
    ) {}

    public record WithdrawalRequest(
            String userId,
            BigDecimal amount,
            Payment.PaymentMethod method,
            String currency,
            String accountDetails,
            String description
    ) {}

    public record RefundRequest(
            BigDecimal amount
    ) {}
}
