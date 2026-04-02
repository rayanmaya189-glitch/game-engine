package com.game_engine.payment.service;

import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Payment;
import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.repository.PaymentRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.time.Instant;
import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class PaymentServiceTest {

    @Mock
    private DepositService depositService;

    @Mock
    private WithdrawalService withdrawalService;

    @Mock
    private PaymentValidationService paymentValidationService;

    @Mock
    private GatewayResolver gatewayResolver;

    @Mock
    private PaymentRepository paymentRepository;

    @Mock
    private WalletService walletService;

    @InjectMocks
    private PaymentService paymentService;

    private UUID paymentId;
    private Payment testPayment;
    private UUID userId;

    @BeforeEach
    void setUp() {
        paymentId = UUID.randomUUID();
        userId = UUID.randomUUID();

        testPayment = Payment.builder()
                .id(paymentId)
                .userId(userId.toString())
                .externalId(UUID.randomUUID().toString())
                .type(Payment.PaymentType.DEPOSIT)
                .method(Payment.PaymentMethod.CREDIT_CARD)
                .status(Payment.PaymentStatus.PENDING)
                .amount(new BigDecimal("100.00"))
                .currency("USD")
                .build();
    }

    @Test
    void processDeposit_shouldDelegateToDepositService() {
        Deposit mockDeposit = new Deposit();
        when(depositService.processDeposit(any(), any(), any(), any(), any(), any(), any()))
                .thenReturn(mockDeposit);

        Deposit result = paymentService.processDeposit(
                userId, new BigDecimal("100.00"), "USD",
                Deposit.PaymentGateway.STRIPE, Deposit.PaymentMethod.CREDIT_CARD,
                Map.of("card", "4242"), "127.0.0.1"
        );

        assertNotNull(result);
        verify(depositService).processDeposit(eq(userId), eq(new BigDecimal("100.00")),
                eq("USD"), eq(Deposit.PaymentGateway.STRIPE), eq(Deposit.PaymentMethod.CREDIT_CARD),
                eq(Map.of("card", "4242")), eq("127.0.0.1"));
    }

    @Test
    void processWithdrawal_shouldDelegateToWithdrawalService() {
        Withdrawal mockWithdrawal = new Withdrawal();
        when(withdrawalService.processWithdrawal(any(), any(), any(), any(), any(), any()))
                .thenReturn(mockWithdrawal);

        Withdrawal result = paymentService.processWithdrawal(
                userId, new BigDecimal("50.00"), "USD",
                Deposit.PaymentGateway.STRIPE, Deposit.PaymentMethod.CREDIT_CARD,
                Map.of("account", "1234")
        );

        assertNotNull(result);
        verify(withdrawalService).processWithdrawal(any(), any(), any(), any(), any(), any());
    }

    @Test
    void approveWithdrawal_shouldDelegateToWithdrawalService() {
        UUID withdrawalId = UUID.randomUUID();
        Withdrawal mockWithdrawal = new Withdrawal();
        when(withdrawalService.approveWithdrawal(withdrawalId, userId)).thenReturn(mockWithdrawal);

        Withdrawal result = paymentService.approveWithdrawal(withdrawalId, userId);

        assertNotNull(result);
        verify(withdrawalService).approveWithdrawal(withdrawalId, userId);
    }

    @Test
    void handleDepositCallback_shouldDelegateToDepositService() {
        Map<String, String> callbackData = Map.of("transaction_id", "tx_123");
        Deposit mockDeposit = new Deposit();
        when(depositService.handleDepositCallback(Deposit.PaymentGateway.STRIPE, callbackData))
                .thenReturn(mockDeposit);

        Deposit result = paymentService.handleDepositCallback(Deposit.PaymentGateway.STRIPE, callbackData);

        assertNotNull(result);
        verify(depositService).handleDepositCallback(Deposit.PaymentGateway.STRIPE, callbackData);
    }

    @Test
    void createDeposit_shouldValidateAndSavePayment() {
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));

        Payment result = paymentService.createDeposit(
                userId.toString(), new BigDecimal("100.00"),
                Payment.PaymentMethod.CREDIT_CARD, "USD", "Test deposit"
        );

        assertNotNull(result);
        assertEquals(Payment.PaymentType.DEPOSIT, result.getType());
        assertEquals(Payment.PaymentStatus.PENDING, result.getStatus());
        assertEquals(new BigDecimal("100.00"), result.getAmount());
        verify(paymentValidationService).validateDepositAmount(new BigDecimal("100.00"));
        verify(paymentRepository).save(any(Payment.class));
    }

    @Test
    void processDepositById_shouldCompleteSuccessfully() {
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));
        when(depositService.processDeposit(any(), any(), any(), any(), any(), any(), any()))
                .thenReturn(new Deposit());

        Payment result = paymentService.processDeposit(paymentId);

        assertEquals(Payment.PaymentStatus.COMPLETED, result.getStatus());
        assertNotNull(result.getProcessedAt());
    }

    @Test
    void processDepositById_shouldFailWhenNotPending() {
        testPayment.setStatus(Payment.PaymentStatus.COMPLETED);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));

        assertThrows(IllegalStateException.class,
                () -> paymentService.processDeposit(paymentId));
    }

    @Test
    void processWithdrawalById_shouldCompleteWithSufficientBalance() {
        testPayment.setType(Payment.PaymentType.WITHDRAWAL);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));
        when(walletService.getBalance(userId, "USD")).thenReturn(new BigDecimal("500.00"));
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));

        Payment result = paymentService.processWithdrawal(paymentId);

        assertEquals(Payment.PaymentStatus.COMPLETED, result.getStatus());
        verify(walletService).debitBalance(eq(userId), eq(new BigDecimal("100.00")),
                eq("USD"), eq("WITHDRAWAL"), any());
    }

    @Test
    void processWithdrawalById_shouldFailWithInsufficientBalance() {
        testPayment.setType(Payment.PaymentType.WITHDRAWAL);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));
        when(walletService.getBalance(userId, "USD")).thenReturn(new BigDecimal("50.00"));
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));

        Payment result = paymentService.processWithdrawal(paymentId);

        assertEquals(Payment.PaymentStatus.FAILED, result.getStatus());
        assertNotNull(result.getFailureReason());
    }

    @Test
    void processWithdrawalById_shouldThrowWhenNotFound() {
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.empty());

        assertThrows(IllegalArgumentException.class,
                () -> paymentService.processWithdrawal(paymentId));
    }

    @Test
    void refundPayment_shouldRefundCompletedPayment() {
        testPayment.setStatus(Payment.PaymentStatus.COMPLETED);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));

        Payment result = paymentService.refundPayment(paymentId, new BigDecimal("100.00"));

        assertEquals(Payment.PaymentStatus.REFUNDED, result.getStatus());
        verify(walletService).creditBalance(eq(userId), eq(new BigDecimal("100.00")),
                eq("USD"), eq("REFUND"), any());
    }

    @Test
    void refundPayment_shouldPartialRefund() {
        testPayment.setStatus(Payment.PaymentStatus.COMPLETED);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));
        when(paymentRepository.save(any(Payment.class))).thenAnswer(inv -> inv.getArgument(0));

        Payment result = paymentService.refundPayment(paymentId, new BigDecimal("50.00"));

        assertEquals(Payment.PaymentStatus.PARTIALLY_REFUNDED, result.getStatus());
    }

    @Test
    void refundPayment_shouldThrowWhenNotCompleted() {
        testPayment.setStatus(Payment.PaymentStatus.PENDING);
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));

        assertThrows(IllegalStateException.class,
                () -> paymentService.refundPayment(paymentId, new BigDecimal("50.00")));
    }

    @Test
    void getPayment_shouldReturnPayment() {
        when(paymentRepository.findById(paymentId)).thenReturn(Optional.of(testPayment));

        Optional<Payment> result = paymentService.getPayment(paymentId);

        assertTrue(result.isPresent());
        assertEquals(paymentId, result.get().getId());
    }

    @Test
    void getSupportedCurrencies_shouldReturnAllCurrencies() {
        List<String> currencies = paymentService.getSupportedCurrencies();

        assertTrue(currencies.contains("USD"));
        assertTrue(currencies.contains("EUR"));
        assertTrue(currencies.contains("BTC"));
        assertTrue(currencies.size() >= 20);
    }
}
