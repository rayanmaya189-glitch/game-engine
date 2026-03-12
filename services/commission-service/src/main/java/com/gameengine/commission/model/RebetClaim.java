package com.game-engine.commission.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "rebet_claim")
public class RebetClaim {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(name = "user_id")
    private Long userId;
    
    @Column(name = "bonus_id")
    private Long bonusId;
    
    @Column(name = "bonus_code")
    private String bonusCode;
    
    @Column(name = "game_id")
    private Long gameId;
    
    @Column(name = "bet_id")
    private Long betId;
    
    @Column(name = "original_bonus_amount", precision = 18, scale = 2)
    private BigDecimal originalBonusAmount;
    
    @Column(name = "rebet_requirement", precision = 18, scale = 2)
    private BigDecimal rebetRequirement;
    
    @Column(name = "current_rebet_amount", precision = 18, scale = 2)
    private BigDecimal currentRebetAmount;
    
    @Column(name = "claim_amount", precision = 18, scale = 2)
    private BigDecimal claimAmount;
    
    @Column(name = "status")
    private String status; // PENDING, IN_PROGRESS, CLAIMABLE, CLAIMED, EXPIRED, CANCELLED
    
    @Column(name = "claim_reason")
    private String claimReason;
    
    @Column(name = "expires_at")
    private LocalDateTime expiresAt;
    
    @Column(name = "claimed_at")
    private LocalDateTime claimedAt;
    
    @Column(name = "transaction_id")
    private String transactionId;
    
    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    public RebetClaim() {
        this.status = "PENDING";
        this.createdAt = LocalDateTime.now();
        this.currentRebetAmount = BigDecimal.ZERO;
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public Long getUserId() { return userId; }
    public void setUserId(Long userId) { this.userId = userId; }
    
    public Long getBonusId() { return bonusId; }
    public void setBonusId(Long bonusId) { this.bonusId = bonusId; }
    
    public String getBonusCode() { return bonusCode; }
    public void setBonusCode(String bonusCode) { this.bonusCode = bonusCode; }
    
    public Long getGameId() { return gameId; }
    public void setGameId(Long gameId) { this.gameId = gameId; }
    
    public Long getBetId() { return betId; }
    public void setBetId(Long betId) { this.betId = betId; }
    
    public BigDecimal getOriginalBonusAmount() { return originalBonusAmount; }
    public void setOriginalBonusAmount(BigDecimal originalBonusAmount) { this.originalBonusAmount = originalBonusAmount; }
    
    public BigDecimal getRebetRequirement() { return rebetRequirement; }
    public void setRebetRequirement(BigDecimal rebetRequirement) { this.rebetRequirement = rebetRequirement; }
    
    public BigDecimal getCurrentRebetAmount() { return currentRebetAmount; }
    public void setCurrentRebetAmount(BigDecimal currentRebetAmount) { this.currentRebetAmount = currentRebetAmount; }
    
    public BigDecimal getClaimAmount() { return claimAmount; }
    public void setClaimAmount(BigDecimal claimAmount) { this.claimAmount = claimAmount; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public String getClaimReason() { return claimReason; }
    public void setClaimReason(String claimReason) { this.claimReason = claimReason; }
    
    public LocalDateTime getExpiresAt() { return expiresAt; }
    public void setExpiresAt(LocalDateTime expiresAt) { this.expiresAt = expiresAt; }
    
    public LocalDateTime getClaimedAt() { return claimedAt; }
    public void setClaimedAt(LocalDateTime claimedAt) { this.claimedAt = claimedAt; }
    
    public String getTransactionId() { return transactionId; }
    public void setTransactionId(String transactionId) { this.transactionId = transactionId; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
    
    public boolean isClaimable() {
        return "CLAIMABLE".equals(status) || 
               (currentRebetAmount.compareTo(rebetRequirement) >= 0);
    }
}
