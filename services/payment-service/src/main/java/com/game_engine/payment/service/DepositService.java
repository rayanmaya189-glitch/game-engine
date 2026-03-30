package com.game_engine.payment.service;

import com.game_engine.payment.gateway.PaymentGatewayAdapter;
import com.game_engine.payment.gateway.PaymentGatewayAdapter.GatewayResponse;
import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.*;
import com.game_engine.payment.repository.DepositRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.Map;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class DepositService {

    private final DepositRepository depositRepository;
    private final WalletService walletService;
    private final PaymentValidationService paymentValidationService;
    private final GatewayResolver gatewayResolver;

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
        paymentValidationService.validateDepositAmount(amount);

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
        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(gateway);

        // 4. Initiate deposit with gateway
        GatewayResponse response = adapter.initiateDeposit(deposit, paymentDetails);

        // 5. Update deposit based on response
        deposit.setGatewayTransactionId(response.getGatewayTransactionId());
        deposit.setStatus(gatewayResolver.mapResponseStatusToDepositStatus(response.getStatus()));

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
        PaymentGatewayAdapter adapter = gatewayResolver.getGatewayAdapter(gateway);
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
}
