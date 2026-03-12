package com.game_engine.commission.model;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity
@Table(name = "commissions")
@Setter
@Getter
public class Commission {

    // Alias methods for backwards compatibility with CommissionService
    public void setCommissionRate(BigDecimal rate) { setCommissionPercentage(rate); }
    public BigDecimal getCommissionRate() { return getCommissionPercentage(); }
    public void setNewPlayers(int players) { setCpaCount(players); }
    public int getNewPlayers() { return getCpaCount() != null ? getCpaCount() : 0; }

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false)
    private Long affiliateId;
    
    @Column(nullable = false)
    private Long merchantId;
    
    @Column(nullable = false)
    private String commissionType; // REVENUE_SHARE, CPA, HYBRID
    
    @Column(nullable = false)
    private String status; // PENDING, APPROVED, PAID, CANCELLED
    
    @Column(nullable = false)
    private BigDecimal netRevenue;
    
    @Column(nullable = false)
    private BigDecimal commissionPercentage;
    
    @Column(nullable = false)
    private BigDecimal commissionAmount;
    
    private BigDecimal cpaAmount;
    
    private Integer cpaCount;
    
    private BigDecimal totalDeposits;
    
    private Integer totalPlayers;
    
    private LocalDateTime rejectedAt;
    
    private String rejectionReason;
    
    @Column(nullable = false)
    private String period; // e.g., "2024-01" for January 2024
    
    @Column(nullable = false)
    private LocalDateTime calculatedAt;
    
    private LocalDateTime approvedAt;
    
    private LocalDateTime paidAt;
    
    private String notes;

    @Column(name = "created_at")
    private LocalDateTime createdAt;
    
    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
    
    // Constructors
    public Commission() {
        this.calculatedAt = LocalDateTime.now();
        this.status = "PENDING";
        this.createdAt = LocalDateTime.now();
    }
    
}
