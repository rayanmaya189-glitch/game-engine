package com.game_engine.affiliate.entity;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "referral_v2")
public class Referral {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "affiliate_code", nullable = false)
    private String affiliateCode;

    @Column(name = "player_id")
    private Long playerId;

    @Column(name = "source")
    private String source; // LINK, BANNER, EMAIL, SOCIAL

    @Column(name = "status")
    private String status; // PENDING, ACTIVE, INACTIVE

    @Column(name = "first_deposit_at")
    private LocalDateTime firstDepositAt;

    @Column(name = "total_deposits", precision = 18, scale = 2)
    private BigDecimal totalDeposits;

    public Referral() {
        this.status = "PENDING";
        this.totalDeposits = BigDecimal.ZERO;
    }

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public String getAffiliateCode() { return affiliateCode; }
    public void setAffiliateCode(String affiliateCode) { this.affiliateCode = affiliateCode; }

    public Long getPlayerId() { return playerId; }
    public void setPlayerId(Long playerId) { this.playerId = playerId; }

    public String getSource() { return source; }
    public void setSource(String source) { this.source = source; }

    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }

    public LocalDateTime getFirstDepositAt() { return firstDepositAt; }
    public void setFirstDepositAt(LocalDateTime firstDepositAt) { this.firstDepositAt = firstDepositAt; }

    public BigDecimal getTotalDeposits() { return totalDeposits; }
    public void setTotalDeposits(BigDecimal totalDeposits) { this.totalDeposits = totalDeposits; }
}
