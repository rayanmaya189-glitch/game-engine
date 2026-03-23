package com.game_engine.bonus.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

/**
 * Player Bonus
 * 
 * Tracks a bonus awarded to a specific player.
 */
@Entity
@Table(name = "player_bonuses")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class PlayerBonus {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(name = "user_id", nullable = false)
    private UUID userId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "campaign_id", nullable = false)
    private BonusCampaign campaign;

    @Column(name = "bonus_amount", nullable = false, precision = 18, scale = 2)
    private BigDecimal bonusAmount;

    @Column(name = "bonus_amount_credited", precision = 18, scale = 2)
    private BigDecimal bonusAmountCredited;

    @Column(name = "wagering_required", precision = 18, scale = 2)
    private BigDecimal wageringRequired;

    @Column(name = "wagering_progress", precision = 18, scale = 2)
    private BigDecimal wageringProgress;

    @Enumerated(EnumType.STRING)
    @Column(name = "status", nullable = false)
    private BonusStatus status;

    @Column(name = "awarded_at")
    private LocalDateTime awardedAt;

    @Column(name = "activated_at")
    private LocalDateTime activatedAt;

    @Column(name = "expires_at")
    private LocalDateTime expiresAt;

    @Column(name = "completed_at")
    private LocalDateTime completedAt;

    @Column(name = "claimed_at")
    private LocalDateTime claimedAt;

    @Column(name = "source_type")
    private String sourceType; // deposit, registration, promotion, referral, loyalty

    @Column(name = "source_id")
    private UUID sourceId;

    @Column(name = "promo_code_used")
    private String promoCodeUsed;

    // For Free Spins
    @Column(name = "free_spins_remaining")
    private Integer freeSpinsRemaining;

    @Column(name = "free_spins_game_id")
    private String freeSpinsGameId;

    @PrePersist
    protected void onCreate() {
        awardedAt = LocalDateTime.now();
        if (status == null) {
            status = BonusStatus.PENDING;
        }
        if (wageringProgress == null) {
            wageringProgress = BigDecimal.ZERO;
        }
    }

    public enum BonusStatus {
        PENDING,      // Awaiting activation (e.g., first deposit)
        ACTIVE,       // Active and can be used
        EXPIRED,      // Expired without completion
        COMPLETED,    // Wagering requirements met
        CANCELLED,    // Cancelled (bonus abuse, etc.)
        CLAIMED       // Converted to real money
    }
}
