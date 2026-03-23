package com.game_engine.kyc;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * KYC Service Application
 * 
 * Know Your Customer verification service with document processing.
 * Supports multiple verification levels:
 * - Level 1: Email + Phone verification
 * - Level 2: Identity document verification
 * - Level 3: Enhanced due diligence
 * 
 * Ports:
 * - HTTP: 8082
 * - gRPC: 9013
 */
@SpringBootApplication
public class KycServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(KycServiceApplication.class, args);
    }
}
