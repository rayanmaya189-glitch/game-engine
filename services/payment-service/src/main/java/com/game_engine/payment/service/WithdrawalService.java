package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.model.Withdrawal.*;
import com.game_engine.payment.repository.WithdrawalRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class WithdrawalService {

    private final WithdrawalRepository withdrawalRepository;
    private final WalletService walletService;
    private final PaymentValidationService paymentValidationService;
    private final GatewayResolver gatewayResolver;

    @Value("${payment.limits.withdrawal.auto-approve-threshold:1000}")
    private BigDecimal autoApproveThreshold;

    @Value("${payment.risk.threshold.medium:50}")
    private int mediumRiskThreshold;

    /**
     * Process a new withdrawal request
     *
     * Flow:
     * 1. Validate amount and limits
     * 2. Check KYC verification
     * 3. Verify wagering requirements
     * 4. Check balance availability
     * 5. Get risk score
     * 6. Auto-approve or flag for manual review
     * 7. Process if approved
     */
    @Transactional
    public Withdrawal processWithdrawal(UUID userId, BigDecimal amount, String currency,
                                        Deposit.PaymentGateway gateway, Deposit.PaymentMethod paymentMethod,
                                        Map<String, String> payoutDetails) {

        // 1. Validate amount
        paymentValidationService.validateWithdrawalAmount(amount);

        // 2. Check balance
        BigDecimal balance = walletService.getBalance(userId, currency);
        if (balance.compareTo(amount) < 0) {
            throw new IllegalArgumentException("Insufficient balance");
        }

        // 3. Check if first withdrawal (requires extra verification)
        boolean isFirstWithdrawal = !withdrawalRepository.existsByUserIdAndStatusIn(
                userId, List.of(WithdrawalStatus.COMPLETED, WithdrawalStatus.APPROVED));

        // 4. Check wagering requirements
        boolean wageringMet = paymentValidationService.checkWageringRequirements(userId, amount);

        // 5. Get risk score (would call Risk Service)
        int riskScore = paymentValidationService.getRiskScore(userId);

        // 6. Determine approval type
        ApprovalType approvalType;
        WithdrawalStatus status;

        if (amount.compareTo(autoApproveThreshold) < 0 &&
            riskScore < mediumRiskThreshold &&
            !isFirstWithdrawal &&
            wageringMet) {
            approvalType = ApprovalType.AUTO_APPROVED;
            status = WithdrawalStatus.APPROVED;
        } else {
            approvalType = ApprovalType.MANUAL_APPROVED;
            status = isFirstWithdrawal || riskScore >= mediumRiskThreshold
                ? WithdrawalStatus.PENDING_APPROVAL
                : WithdrawalStatus.APPROVED;
        }

        // 7. Create withdrawal record
        Withdrawal withdrawal = Withdrawal.builder()
                .userId(userId)
                .amount(amount)
                .currency(currency.toUpperCase())
                .gateway(gateway)
                .paymentMethod(paymentMethod)
                .paymentDestination(payoutDetails.get("destination_token"))
                .beneficiaryName(payoutDetails.get("beneficiary_name"))
                .status(status)
                .approvalType(approvalType)
                .isFirstWithdrawal(isFirstWithdrawal)
                .wageringRequirementMet(wageringMet)
                .riskScore(riskScore)
                .build();

        withdrawal = withdrawalRepository.save(withdrawal);

        // 8. If auto-approved, process immediately
        if (status == WithdrawalStatus.APPROVED) {
            processApprovedWithdrawal(withdrawal);
        }

        return withdrawal;
    }

    /**
     * Process an approved withdrawal
     */
    @Transactional
    public void processApprovedWithdrawal(Withdrawal withdrawal) {
        // Debit wallet
        walletService.debitBalance(withdrawal.getUserId(), withdrawal.getAmount(), withdrawal.getCurrency(),
                "WITHDRAWAL", withdrawal.getId().toString());

        // Initiate payout with gateway
        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(withdrawal.getGateway());
        Map<String, String> payoutDetails = Map.of(
            "destination_token", withdrawal.getPaymentDestination(),
            "beneficiary_name", withdrawal.getBeneficiaryName() != null && !withdrawal.getBeneficiaryName().isEmpty() ? withdrawal.getBeneficiaryName() : ""
        );

        GatewayResponse response = adapter.initiateWithdrawal(withdrawal, payoutDetails);

        withdrawal.setGatewayTransactionId(response.getGatewayTransactionId());

        if (response.isSuccess()) {
            withdrawal.setStatus(WithdrawalStatus.PROCESSING);
            withdrawal.setProcessedAt(LocalDateTime.now());
        } else {
            withdrawal.setStatus(WithdrawalStatus.FAILED);
            withdrawal.setGatewayResponseMessage(response.getMessage());
        }

        withdrawalRepository.save(withdrawal);
    }

    /**
     * Approve a withdrawal (admin action)
     */
    @Transactional
    public Withdrawal approveWithdrawal(UUID withdrawalId, UUID adminUserId) {
        Withdrawal withdrawal = withdrawalRepository.findById(withdrawalId)
                .orElseThrow(() -> new IllegalArgumentException("Withdrawal not found"));

        if (withdrawal.getStatus() != WithdrawalStatus.PENDING_APPROVAL) {
            throw new IllegalStateException("Withdrawal is not pending approval");
        }

        withdrawal.setStatus(WithdrawalStatus.APPROVED);
        withdrawal.setApprovedBy(adminUserId);
        withdrawal.setApprovedAt(LocalDateTime.now());
        withdrawal.setApprovalType(ApprovalType.MANUAL_APPROVED);

        withdrawal = withdrawalRepository.save(withdrawal);

        // Process the withdrawal
        processApprovedWithdrawal(withdrawal);

        return withdrawal;
    }

    /**
     * Reject a withdrawal (admin action)
     */
    @Transactional
    public Withdrawal rejectWithdrawal(UUID withdrawalId, UUID adminUserId, String reason) {
        Withdrawal withdrawal = withdrawalRepository.findById(withdrawalId)
                .orElseThrow(() -> new IllegalArgumentException("Withdrawal not found"));

        if (withdrawal.getStatus() != WithdrawalStatus.PENDING_APPROVAL &&
            withdrawal.getStatus() != WithdrawalStatus.PENDING_REVIEW) {
            throw new IllegalStateException("Withdrawal cannot be rejected");
        }

        // Refund balance if already debited
        if (withdrawal.getStatus() == WithdrawalStatus.APPROVED ||
            withdrawal.getStatus() == WithdrawalStatus.PROCESSING) {
            walletService.creditBalance(withdrawal.getUserId(), withdrawal.getAmount(), withdrawal.getCurrency(),
                    "WITHDRAWAL_REVERSAL", withdrawal.getId().toString());
        }

        withdrawal.setStatus(WithdrawalStatus.REJECTED);
        withdrawal.setApprovedBy(adminUserId);
        withdrawal.setApprovedAt(LocalDateTime.now());
        withdrawal.setApprovalType(ApprovalType.MANUAL_REJECTED);
        withdrawal.setRejectionReason(reason);

        return withdrawalRepository.save(withdrawal);
    }
}
