package com.game_engine.payment.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

/**
 * Withdrawal Entity
 * 
 * Represents a withdrawal request in the system.
 * Includes approval workflow and risk assessment integration.
 * 
 * Status Flow:
 * PENDING_REVIEW → APPROVED/PENDING_APPROVAL
 *                 → PROCESSING → COMPLETED/FAILED
 */
@Entity
@Table(name = "withdrawals", indexes = {
    @Index(name = "idx_withdrawal_user_id", columnList = "user_id"),
    @Index(name = "idx_withdrawal_status", columnList = "status"),
    @Index(name = "idx_withdrawal_gateway", columnList = "gateway"),
    @Index(name = "idx_withdrawal_created_at", columnList = "created_at")
})
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode(of = "id")
public class Withdrawal {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(name = "user_id", nullable = false)
    private UUID userId;

    @Column(name = "amount", nullable = false, precision = 18, scale = 2)
    private BigDecimal amount;

    @Column(name = "currency", nullable = false, length = 3)
    private String currency;

    @Enumerated(EnumType.STRING)
    @Column(name = "gateway", nullable = false)
    private Deposit.PaymentGateway gateway;

    @Enumerated(EnumType.STRING)
    @Column(name = "payment_method", nullable = false)
    private Deposit.PaymentMethod paymentMethod;

    @Enumerated(EnumType.STRING)
    @Column(name = "status", nullable = false)
    private WithdrawalStatus status;

    @Column(name = "payment_destination", nullable = false, columnDefinition = "TEXT")
    private String paymentDestination; // Tokenized or encrypted payment details

    @Column(name = "beneficiary_name")
    private String beneficiaryName;

    @Column(name = "gateway_transaction_id")
    private String gatewayTransactionId;

    @Column(name = "gateway_response_code")
    private String gatewayResponseCode;

    @Column(name = "gateway_response_message")
    private String gatewayResponseMessage;

    @Column(name = "risk_score")
    private Integer riskScore;

    @Enumerated(EnumType.STRING)
    @Column(name = "approval_type")
    private ApprovalType approvalType;

    @Column(name = "approved_by")
    private UUID approvedBy;

    @Column(name = "approved_at")
    private LocalDateTime approvedAt;

    @Column(name = "rejection_reason")
    private String rejectionReason;

    @Column(name = "is_first_withdrawal")
    private Boolean isFirstWithdrawal;

    @Column(name = "wagering_requirement_met")
    private Boolean wageringRequirementMet;

    @Column(name = "metadata", columnDefinition = "TEXT")
    private String metadata;

    @Column(name = "created_at", nullable = false)
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    @Column(name = "processed_at")
    private LocalDateTime processedAt;

    @Column(name = "completed_at")
    private LocalDateTime completedAt;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        if (status == null) {
            status = WithdrawalStatus.PENDING_REVIEW;
        }
    }

    @PreUpdate
    protected void onUpdate() {
        updatedAt = LocalDateTime.now();
    }

    public enum WithdrawalStatus {
        PENDING_REVIEW,      // Awaiting initial review
        PENDING_APPROVAL,   // Requires manual approval
        APPROVED,           // Approved, ready for processing
        PROCESSING,         // Payout in progress
        COMPLETED,          // Successfully paid out
        FAILED,             // Payout failed
        REJECTED,           // Rejected
        CANCELLED           // User cancelled
    }

    public enum ApprovalType {
        AUTO_APPROVED,
        MANUAL_APPROVED,
        SYSTEM_REJECTED,
        MANUAL_REJECTED
    }
}
