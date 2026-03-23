package com.game_engine.bonus;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Bonus Service Application
 * 
 * Manages all bonus types and wagering requirements.
 * 
 * Supported Bonus Types:
 * - Welcome Bonus (first deposit match)
 * - Reload Bonus (subsequent deposits)
 * - No-Deposit Bonus (registration bonus)
 * - Free Spins
 * - Cashback
 * - Referral Bonus
 * 
 * Port: 8084
 */
@SpringBootApplication
public class BonusServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(BonusServiceApplication.class, args);
    }
}
