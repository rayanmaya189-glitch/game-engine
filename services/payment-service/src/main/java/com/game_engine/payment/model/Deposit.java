package com.game_engine.payment.model;

import jakarta.persistence.*;
import lombok.*;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.UUID;

/**
 * Deposit Entity
 * 
 * Represents a deposit transaction in the system.
 * Follows PCI DSS requirements - no raw card data is stored.
 * 
 * Status Flow:
 * INITIATED → PROCESSING → COMPLETED/FAILED
 *           → PENDING_VERIFICATION (3DS/SCA)
 */
@Entity
@Table(name = "deposits", indexes = {
    @Index(name = "idx_deposit_user_id", columnList = "user_id"),
    @Index(name = "idx_deposit_status", columnList = "status"),
    @Index(name = "idx_deposit_gateway", columnList = "gateway"),
    @Index(name = "idx_deposit_created_at", columnList = "created_at")
})
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode(of = "id")
public class Deposit {

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
    private PaymentGateway gateway;

    @Enumerated(EnumType.STRING)
    @Column(name = "payment_method", nullable = false)
    private PaymentMethod paymentMethod;

    @Enumerated(EnumType.STRING)
    @Column(name = "status", nullable = false)
    private DepositStatus status;

    @Column(name = "gateway_transaction_id")
    private String gatewayTransactionId;

    @Column(name = "gateway_response_code")
    private String gatewayResponseCode;

    @Column(name = "gateway_response_message")
    private String gatewayResponseMessage;

    @Column(name = "payment_token")
    private String paymentToken; // Tokenized payment reference

    @Column(name = "redirect_url")
    private String redirectUrl; // For 3DS/SCA redirect

    @Column(name = "ip_address")
    private String ipAddress;

    @Column(name = "user_agent")
    private String userAgent;

    @Column(name = "kyc_level_required")
    private Integer kycLevelRequired;

    @Column(name = "risk_score")
    private Integer riskScore;

    @Column(name = "metadata", columnDefinition = "TEXT")
    private String metadata; // JSON string for additional data

    @Column(name = "created_at", nullable = false)
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;

    @Column(name = "completed_at")
    private LocalDateTime completedAt;

    @PrePersist
    protected void onCreate() {
        createdAt = LocalDateTime.now();
        if (status == null) {
            status = DepositStatus.INITIATED;
        }
    }

    @PreUpdate
    protected void onUpdate() {
        updatedAt = LocalDateTime.now();
    }

    public enum PaymentGateway {
        STRIPE,
        ADYEN,
        SKRILL,
        NETELLER,
        PAYSAFECARD,
        COINBASE_COMMERCE,
        BITPAY,
        LOCAL_BANK_TRANSFER
    }

    public enum PaymentMethod {
        CREDIT_CARD,
        DEBIT_CARD,
        BANK_TRANSFER,
        E_WALLET,
        CRYPTOCURRENCY,
        PREPAID_CARD
    }

    public enum DepositStatus {
        INITIATED,           // Created, awaiting payment
        PROCESSING,         // Payment in progress
        PENDING_VERIFICATION, // 3DS/SCA required
        COMPLETED,          // Successfully processed
        FAILED,             // Payment failed
        CANCELLED,          // User cancelled
        EXPIRED             // Payment timeout
    }
}
