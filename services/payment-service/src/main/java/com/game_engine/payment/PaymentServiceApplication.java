package com.game_engine.payment;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableScheduling;

/**
 * Payment Service Application
 * 
 * Handles all payment processing including deposits, withdrawals,
 * and reconciliation across multiple payment gateways.
 * 
 * Supported Gateways:
 * - Stripe (Cards)
 * - Adyen (Cards + Local Methods)
 * - Skrill (E-wallet)
 * - Neteller (E-wallet)
 * - Paysafecard (Prepaid)
 * - Crypto (Coinbase, BitPay)
 * 
 * Ports:
 * - HTTP: 8081
 * - gRPC: 9012
 */
@SpringBootApplication
@EnableScheduling
public class PaymentServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(PaymentServiceApplication.class, args);
    }
}
