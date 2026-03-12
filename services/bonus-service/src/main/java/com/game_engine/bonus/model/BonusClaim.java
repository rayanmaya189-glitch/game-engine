package com.game_engine.bonus.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "bonus_claims")
@Getter @Setter
@NoArgsConstructor @AllArgsConstructor @Builder
public class BonusClaim {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(name = "user_id", nullable = false)
    private UUID userId;

    @Column(name = "bonus_id", nullable = false)
    private UUID bonusId;

    @Column(name = "bonus_amount", nullable = false)
    private BigDecimal bonusAmount;

    @Column(name = "wagering_requirement", nullable = false)
    private BigDecimal wageringRequirement;

    @Column(name = "wagering_contributed")
    private BigDecimal wageringContributed;

    @Column(name = "winnings_amount")
    private BigDecimal winningsAmount;

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    private ClaimStatus status;

    @Column(name = "claimed_at", nullable = false)
    private Instant claimedAt;

    @Column(name = "completed_at")
    private Instant completedAt;

    @Column(name = "expires_at")
    private Instant expiresAt;

    @Column(name = "cancelled_at")
    private Instant cancelledAt;

    @Column(name = "cancellation_reason")
    private String cancellationReason;

    public enum ClaimStatus {
        ACTIVE, COMPLETED, CANCELLED, EXPIRED
    }

    public boolean isWageringComplete() {
        if (wageringContributed == null || wageringRequirement == null) {
            return false;
        }
        return wageringContributed.compareTo(wageringRequirement) >= 0;
    }

    public BigDecimal getRemainingWagering() {
        if (wageringContributed == null || wageringRequirement == null) {
            return wageringRequirement;
        }
        return wageringRequirement.subtract(wageringContributed).max(BigDecimal.ZERO);
    }
}
