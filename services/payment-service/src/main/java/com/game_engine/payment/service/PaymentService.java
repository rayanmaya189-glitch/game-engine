package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.model.Withdrawal.*;
import com.game_engine.payment.repository.DepositRepository;
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

/**
 * Payment Service
 * 
 * Core business logic for deposit and withdrawal processing.
 * Coordinates between payment gateways, wallet service, and compliance systems.
 */
@Service
@RequiredArgsConstructor
@Slf4j
public class PaymentService {

    private final DepositRepository depositRepository;
    private final WithdrawalRepository withdrawalRepository;
    private final List<PaymentGatewayAdapter> gatewayAdapters;
    private final WalletService walletService;

    // Configuration
    @Value("${payment.limits.deposit.min:10}")
    private BigDecimal minDeposit;

    @Value("${payment.limits.deposit.max:10000}")
    private BigDecimal maxDeposit;

    @Value("${payment.limits.withdrawal.min:20}")
    private BigDecimal minWithdrawal;

    @Value("${payment.limits.withdrawal.max:50000}")
    private BigDecimal maxWithdrawal;

    @Value("${payment.limits.withdrawal.auto-approve-threshold:1000}")
    private BigDecimal autoApproveThreshold;

    @Value("${payment.risk.threshold.medium:50}")
    private int mediumRiskThreshold;

    /**
     * Process a new deposit request
     * 
     * Flow:
     * 1. Validate amount and limits
     * 2. Check KYC requirements
     * 3. Get risk score
     * 4. Select appropriate gateway
     * 5. Initiate deposit with gateway
     * 6. Return redirect URL or status
     */
    @Transactional
    public Deposit processDeposit(UUID userId, BigDecimal amount, String currency,
                                   PaymentGateway gateway, PaymentMethod paymentMethod,
                                   Map<String, String> paymentDetails, String ipAddress) {
        
        // 1. Validate amount
        validateDepositAmount(amount);

        // 2. Create deposit record
        Deposit deposit = Deposit.builder()
                .userId(userId)
                .amount(amount)
                .currency(currency.toUpperCase())
                .gateway(gateway)
                .paymentMethod(paymentMethod)
                .status(DepositStatus.INITIATED)
                .ipAddress(ipAddress)
                .paymentToken(paymentDetails.get("payment_token"))
                .build();

        deposit = depositRepository.save(deposit);

        // 3. Get gateway adapter
        PaymentGatewayAdapter adapter = getGatewayAdapter(gateway);

        // 4. Initiate deposit with gateway
        GatewayResponse response = adapter.initiateDeposit(deposit, paymentDetails);

        // 5. Update deposit based on response
        deposit.setGatewayTransactionId(response.getGatewayTransactionId());
        deposit.setStatus(mapResponseStatusToDepositStatus(response.getStatus()));

        if (response.getMessage() != null) {
            deposit.setGatewayResponseMessage(response.getMessage());
        }

        // Handle redirect for 3DS
        if (response.getRedirectUrl() != null) {
            deposit.setRedirectUrl(response.getRedirectUrl());
            deposit.setStatus(DepositStatus.PENDING_VERIFICATION);
        }

        deposit = depositRepository.save(deposit);

        // 6. If immediate success, credit wallet
        if (response.isSuccess() && "COMPLETED".equals(response.getStatus())) {
            completeDeposit(deposit);
        }

        return deposit;
    }

    /**
     * Handle deposit webhook callback from payment gateway
     */
    @Transactional
    public Deposit handleDepositCallback(PaymentGateway gateway, Map<String, String> callbackData) {
        PaymentGatewayAdapter adapter = getGatewayAdapter(gateway);
        GatewayResponse response = adapter.processDepositCallback(callbackData);

        String gatewayTransactionId = response.getGatewayTransactionId();
        Deposit deposit = depositRepository.findByGatewayTransactionId(gatewayTransactionId)
                .orElseThrow(() -> new IllegalArgumentException("Deposit not found: " + gatewayTransactionId));

        if (response.isSuccess() && "COMPLETED".equals(response.getStatus())) {
            completeDeposit(deposit);
        } else {
            deposit.setStatus(DepositStatus.FAILED);
            deposit.setGatewayResponseMessage(response.getMessage());
        }

        return depositRepository.save(deposit);
    }

    /**
     * Complete a successful deposit - credit the wallet
     */
    @Transactional
    public void completeDeposit(Deposit deposit) {
        deposit.setStatus(DepositStatus.COMPLETED);
        deposit.setCompletedAt(LocalDateTime.now());

        // Credit wallet
        walletService.creditBalance(deposit.getUserId(), deposit.getAmount(), deposit.getCurrency(),
                "DEPOSIT", deposit.getId().toString());

        depositRepository.save(deposit);
        log.info("Deposit completed: {} - {} {}", deposit.getId(), deposit.getAmount(), deposit.getCurrency());
    }

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
                                        PaymentGateway gateway, PaymentMethod paymentMethod,
                                        Map<String, String> payoutDetails) {
        
        // 1. Validate amount
        validateWithdrawalAmount(amount);

        // 2. Check balance
        BigDecimal balance = walletService.getBalance(userId, currency);
        if (balance.compareTo(amount) < 0) {
            throw new IllegalArgumentException("Insufficient balance");
        }

        // 3. Check if first withdrawal (requires extra verification)
        boolean isFirstWithdrawal = !withdrawalRepository.existsByUserIdAndStatusIn(
                userId, List.of(WithdrawalStatus.COMPLETED, WithdrawalStatus.APPROVED));

        // 4. Check wagering requirements
        boolean wageringMet = checkWageringRequirements(userId, amount);

        // 5. Get risk score (would call Risk Service)
        int riskScore = getRiskScore(userId);

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
        PaymentGatewayAdapter adapter = getGatewayAdapter(withdrawal.getGateway());
        Map<String, String> payoutDetails = Map.of(
            "destination_token", withdrawal.getPaymentDestination(),
            "beneficiary_name", withdrawal.getBeneficiaryName() != "" ? withdrawal.getBeneficiaryName() : ""
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

    private void validateDepositAmount(BigDecimal amount) {
        if (amount.compareTo(minDeposit) < 0) {
            throw new IllegalArgumentException("Minimum deposit is " + minDeposit);
        }
        if (amount.compareTo(maxDeposit) > 0) {
            throw new IllegalArgumentException("Maximum deposit is " + maxDeposit);
        }
    }

    private void validateWithdrawalAmount(BigDecimal amount) {
        if (amount.compareTo(minWithdrawal) < 0) {
            throw new IllegalArgumentException("Minimum withdrawal is " + minWithdrawal);
        }
        if (amount.compareTo(maxWithdrawal) > 0) {
            throw new IllegalArgumentException("Maximum withdrawal is " + maxWithdrawal);
        }
    }

    private boolean checkWageringRequirements(UUID userId, BigDecimal amount) {
        // Would check with Bonus Service
        // Simplified for now - assume requirements met
        return true;
    }

    private int getRiskScore(UUID userId) {
        // Would call Risk Scoring Service via gRPC
        // Simplified for now
        return 0;
    }

    private PaymentGatewayAdapter getGatewayAdapter(PaymentGateway gateway) {
        return gatewayAdapters.stream()
                .filter(a -> a.getGatewayType() == gateway)
                .findFirst()
                .orElseThrow(() -> new IllegalArgumentException("Gateway not supported: " + gateway));
    }

    private DepositStatus mapResponseStatusToDepositStatus(String responseStatus) {
        if (responseStatus == null) return DepositStatus.FAILED;
        
        return switch (responseStatus) {
            case "COMPLETED" -> DepositStatus.COMPLETED;
            case "PROCESSING" -> DepositStatus.PROCESSING;
            case "PENDING_VERIFICATION" -> DepositStatus.PENDING_VERIFICATION;
            case "FAILED" -> DepositStatus.FAILED;
            case "CANCELLED" -> DepositStatus.CANCELLED;
            default -> DepositStatus.PROCESSING;
        };
    }
}
