package com.game_engine.payment.service;

import com.game_engine.payment.grpc.WalletGrpcClient;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.UUID;

@Service
@RequiredArgsConstructor
@Slf4j
public class WalletService {

    private final WalletGrpcClient walletGrpcClient;

    public void creditBalance(UUID userId, BigDecimal amount, String currency,
                              String transactionType, String referenceId) {
        walletGrpcClient.creditBalance(userId, amount, currency, transactionType, referenceId);
    }

    public void debitBalance(UUID userId, BigDecimal amount, String currency,
                             String transactionType, String referenceId) {
        walletGrpcClient.debitBalance(userId, amount, currency, transactionType, referenceId);
    }

    public BigDecimal getBalance(UUID userId, String currency) {
        return walletGrpcClient.getBalance(userId, currency);
    }
}
