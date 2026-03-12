package com.game_engine.bonus.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "bonuses")
@Getter @Setter
@NoArgsConstructor @AllArgsConstructor @Builder
public class Bonus {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    private String name;
    private String description;

    @Enumerated(EnumType.STRING)
    private BonusType type;

    @Enumerated(EnumType.STRING)
    private BonusStatus status;

    private BigDecimal amount;
    private BigDecimal percentage;
    private BigDecimal maxAmount;
    private BigDecimal minDeposit;

    private Integer wageringRequirement;
    private BigDecimal maxBet;
    private String allowedGames;

    private Instant startDate;
    private Instant endDate;
    private Integer maxUses;
    private Integer currentUses;

    @Enumerated(EnumType.STRING)
    private VIPLevel vipLevel;

    private Instant createdAt;
    private Instant updatedAt;

    public enum BonusType {
        WELCOME, DEPOSIT, NO_DEPOSIT, FREE_SPINS, CASHBACK, RELOAD, LOYALTY
    }

    public enum BonusStatus {
        ACTIVE, INACTIVE, EXPIRED
    }

    public enum VIPLevel {
        NONE, BRONZE, SILVER, GOLD, PLATINUM, DIAMOND
    }
}
