package com.gameengine.payment.repository;

import com.gameengine.payment.model.Payment;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.Instant;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface PaymentRepository extends JpaRepository<Payment, UUID> {

    Page<Payment> findByUserId(String userId, Pageable pageable);

    List<Payment> findByUserIdAndStatus(String userId, Payment.PaymentStatus status);

    Optional<Payment> findByExternalId(String externalId);

    @Query("SELECT p FROM Payment p WHERE p.userId = :userId AND p.type = :type AND p.status = :status " +
           "AND p.createdAt >= :startDate AND p.createdAt <= :endDate")
    List<Payment> findByUserIdAndTypeAndStatusAndDateRange(
            @Param("userId") String userId,
            @Param("type") Payment.PaymentType type,
            @Param("status") Payment.PaymentStatus status,
            @Param("startDate") Instant startDate,
            @Param("endDate") Instant endDate
    );

    @Query("SELECT SUM(p.amount) FROM Payment p WHERE p.userId = :userId AND p.type = :type " +
           "AND p.status = 'COMPLETED' AND p.createdAt >= :startDate AND p.createdAt <= :endDate")
    java.math.BigDecimal sumByUserIdAndTypeAndStatusAndDateRange(
            @Param("userId") String userId,
            @Param("type") Payment.PaymentType type,
            @Param("startDate") Instant startDate,
            @Param("endDate") Instant endDate
    );

    @Query("SELECT p FROM Payment p WHERE p.status = 'PENDING' AND p.expiresAt < :now")
    List<Payment> findExpiredPendingPayments(@Param("now") Instant now);

    @Query("SELECT p FROM Payment p WHERE p.type = :type AND p.createdAt >= :startDate")
    List<Payment> findByTypeAndDateRange(
            @Param("type") Payment.PaymentType type,
            @Param("startDate") Instant startDate
    );

    Page<Payment> findByStatus(Payment.PaymentStatus status, Pageable pageable);
}
