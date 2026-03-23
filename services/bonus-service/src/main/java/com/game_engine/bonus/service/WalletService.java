package com.game_engine.bonus.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.UUID;

/**
 * Wallet Service Interface
 * 
 * Placeholder for Wallet Service communication.
 * In production, this would use gRPC to communicate with the Wallet microservice.
 */
@Service
@Slf4j
public class WalletService {

    public void creditBonusBalance(UUID userId, BigDecimal amount, String currency, String type, String reference) {
        log.info("Crediting {} {} bonus to user {} for {} - {}", amount, currency, userId, type, reference);
    }

    public void creditBalance(UUID userId, BigDecimal amount, String currency, String type, String reference) {
        log.info("Crediting {} {} to user {} for {} - {}", amount, currency, userId, type, reference);
    }

    public void convertBonusToReal(UUID userId, BigDecimal amount) {
        log.info("Converting {} bonus to real balance for user {}", amount, userId);
    }

    public BigDecimal getBonusBalance(UUID userId) {
        return BigDecimal.ZERO;
    }

    public BigDecimal getTotalBets(UUID userId, int days) {
        return BigDecimal.valueOf(1000); // Simplified
    }

    public BigDecimal getTotalWins(UUID userId, int days) {
        return BigDecimal.valueOf(800); // Simplified
    }
}
