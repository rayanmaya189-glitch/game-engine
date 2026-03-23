package com.game_engine.bonus.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

/**
 * Bonus Campaign
 * 
 * Defines a bonus promotion campaign with configurable parameters.
 */
@Entity
@Table(name = "bonus_campaigns")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class BonusCampaign {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(nullable = false)
    private String name;

    @Column(name = "bonus_type", nullable = false)
    @Enumerated(EnumType.STRING)
    private BonusType bonusType;

    @Column(nullable = false)
    private String description;

    @Column(name = "is_active")
    private Boolean isActive;

    @Column(name = "start_date")
    private LocalDateTime startDate;

    @Column(name = "end_date")
    private LocalDateTime endDate;

    // Bonus Amount Configuration
    @Column(name = "match_percentage", precision = 5, scale = 2)
    private BigDecimal matchPercentage;

    @Column(name = "max_bonus_amount", precision = 18, scale = 2)
    private BigDecimal maxBonusAmount;

    @Column(name = "min_deposit", precision = 18, scale = 2)
    private BigDecimal minDeposit;

    @Column(name = "fixed_amount", precision = 18, scale = 2)
    private BigDecimal fixedAmount;

    // Free Spins Configuration
    @Column(name = "free_spin_count")
    private Integer freeSpinCount;

    @Column(name = "free_spin_game_id")
    private String freeSpinGameId;

    @Column(name = "free_spin_bet_value", precision = 18, scale = 2)
    private BigDecimal freeSpinBetValue;

    // Wagering Requirements
    @Column(name = "wagering_requirement", precision = 18, scale = 2)
    private BigDecimal wageringRequirement;

    @Column(name = "wagering_multiplier")
    private Integer wageringMultiplier; // e.g., 30x bonus amount

    // Game Weights (JSON string)
    @Column(name = "game_weights", columnDefinition = "TEXT")
    private String gameWeights; // JSON: {"slots": 1.0, "blackjack": 0.1, "roulette": 0.2}

    // Cashback Configuration
    @Column(name = "cashback_percentage", precision = 5, scale = 2)
    private BigDecimal cashbackPercentage;

    @Column(name = "cashback_period_days")
    private Integer cashbackPeriodDays;

    // Eligibility
    @Column(name = "min_kyc_level")
    private Integer minKycLevel;

    @Column(name = "max_usage_count")
    private Integer maxUsageCount;

    @Column(name = "eligible_countries", columnDefinition = "TEXT")
    private String eligibleCountries; // JSON array

    @Column(name = "promo_code")
    private String promoCode;

    // Referral Config
    @Column(name = "referrer_bonus", precision = 18, scale = 2)
    private BigDecimal referrerBonus;

    @Column(name = "referee_bonus", precision = 18, scale = 2)
    private BigDecimal refereeBonus;

    // Limits
    @Column(name = "max_cashout", precision = 18, scale = 2)
    private BigDecimal maxCashout;

    // Expiry
    @Column(name = "expiry_days")
    private Integer expiryDays;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        if (isActive == null) {
            isActive = false;
        }
    }

    public enum BonusType {
        WELCOME,
        RELOAD,
        NO_DEPOSIT,
        FREE_SPINS,
        CASHBACK,
        REFERRAL,
        LOYALTY,
        VIP
    }
}
