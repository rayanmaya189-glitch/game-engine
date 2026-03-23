package com.game_engine.payment.repository;

import com.game_engine.payment.model.Deposit;
import com.game_engine.payment.model.Deposit.DepositStatus;
import com.game_engine.payment.model.Deposit.PaymentGateway;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface DepositRepository extends JpaRepository<Deposit, UUID> {

    Optional<Deposit> findByGatewayTransactionId(String gatewayTransactionId);

    Page<Deposit> findByUserId(UUID userId, Pageable pageable);

    Page<Deposit> findByStatus(DepositStatus status, Pageable pageable);

    Page<Deposit> findByUserIdAndStatus(UUID userId, DepositStatus status, Pageable pageable);

    List<Deposit> findByGatewayAndCreatedAtBetween(PaymentGateway gateway, 
                                                    LocalDateTime start, 
                                                    LocalDateTime end);

    @Query("SELECT d FROM Deposit d WHERE d.userId = :userId AND d.createdAt >= :since AND d.status = :status")
    List<Deposit> findRecentByUserAndStatus(@Param("userId") UUID userId,
                                            @Param("since") LocalDateTime since,
                                            @Param("status") DepositStatus status);

    @Query("SELECT SUM(d.amount) FROM Deposit d WHERE d.userId = :userId AND d.status = 'COMPLETED' AND d.createdAt >= :since")
    java.math.BigDecimal sumDepositsByUserSince(@Param("userId") UUID userId, @Param("since") LocalDateTime since);

    @Query("SELECT COUNT(d) FROM Deposit d WHERE d.userId = :userId AND d.status IN :statuses AND d.createdAt >= :since")
    int countDepositsByUserSince(@Param("userId") UUID userId,
                                  @Param("statuses") List<DepositStatus> statuses,
                                  @Param("since") LocalDateTime since);
}
