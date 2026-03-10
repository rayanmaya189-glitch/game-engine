package com.gameengine.payment.model;

import jakarta.persistence.*;
import lombok.*;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import java.math.BigDecimal;
import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "payments", indexes = {
    @Index(name = "idx_payment_user_id", columnList = "userId"),
    @Index(name = "idx_payment_status", columnList = "status"),
    @Index(name = "idx_payment_type", columnList = "type"),
    @Index(name = "idx_payment_external_id", columnList = "externalId")
})
@EntityListeners(AuditingEntityListener.class)
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Payment {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    @Column(nullable = false)
    private String userId;

    @Column(nullable = false)
    private String externalId;

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    private PaymentType type;

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    private PaymentMethod method;

    @Enumerated(EnumType.STRING)
    @Column(nullable = false)
    private PaymentStatus status;

    @Column(nullable = false, precision = 19, scale = 4)
    private BigDecimal amount;

    @Column(nullable = false, length = 3)
    private String currency;

    @Column(precision = 19, scale = 4)
    private BigDecimal convertedAmount;

    @Column(length = 3)
    private String convertedCurrency;

    @Column(precision = 19, scale = 4)
    private BigDecimal fee;

    @Column(precision = 19, scale = 4)
    private BigDecimal netAmount;

    @Column(length = 500)
    private String description;

    @Column(length = 500)
    private String metadata;

    @Column(length = 100)
    private String paymentGateway;

    @Column(length = 500)
    private String gatewayResponse;

    @Column(length = 500)
    private String failureReason;

    @Column
    private Integer retries;

    @Column
    private Instant processedAt;

    @Column
    private Instant completedAt;

    @Column
    private Instant expiresAt;

    @CreatedDate
    @Column(nullable = false, updatable = false)
    private Instant createdAt;

    @LastModifiedDate
    @Column(nullable = false)
    private Instant updatedAt;

    public enum PaymentType {
        DEPOSIT,
        WITHDRAWAL,
        REFUND,
        TRANSFER
    }

    public enum PaymentMethod {
        CREDIT_CARD,
        DEBIT_CARD,
        VIRTUAL_CARD,
        PAYPAL,
        SKRILL,
        NETELLER,
        BITCOIN,
        ETHEREUM,
        BANK_TRANSFER,
        INSTANT_BANK_TRANSFER,
        ECOPAYZ,
        MUCHBETTER,
        ASTROPAY,
        JETON
    }

    public enum PaymentStatus {
        PENDING,
        PROCESSING,
        COMPLETED,
        FAILED,
        CANCELLED,
        EXPIRED,
        REFUNDED,
        PARTIALLY_REFUNDED
    }
}
