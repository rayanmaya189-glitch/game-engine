package com.game_engine.payment.repository;

import com.game_engine.payment.model.Withdrawal;
import com.game_engine.payment.model.Withdrawal.WithdrawalStatus;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface WithdrawalRepository extends JpaRepository<Withdrawal, UUID> {

    Page<Withdrawal> findByUserId(UUID userId, Pageable pageable);

    Page<Withdrawal> findByStatus(WithdrawalStatus status, Pageable pageable);

    Page<Withdrawal> findByUserIdAndStatus(UUID userId, WithdrawalStatus status, Pageable pageable);

    boolean existsByUserIdAndStatusIn(UUID userId, List<WithdrawalStatus> statuses);

    @Query("SELECT w FROM Withdrawal w WHERE w.userId = :userId AND w.createdAt >= :since AND w.status IN :statuses")
    List<Withdrawal> findRecentByUserAndStatus(@Param("userId") UUID userId,
                                                @Param("since") LocalDateTime since,
                                                @Param("statuses") List<WithdrawalStatus> statuses);

    @Query("SELECT SUM(w.amount) FROM Withdrawal w WHERE w.userId = :userId AND w.status IN ('COMPLETED', 'APPROVED', 'PROCESSING') AND w.createdAt >= :since")
    BigDecimal sumPendingWithdrawalsByUserSince(@Param("userId") UUID userId, @Param("since") LocalDateTime since);

    @Query("SELECT COUNT(w) FROM Withdrawal w WHERE w.userId = :userId AND w.status = 'COMPLETED' AND w.createdAt >= :since")
    int countCompletedWithdrawalsByUserSince(@Param("userId") UUID userId, @Param("since") LocalDateTime since);

    List<Withdrawal> findByStatusInAndCreatedAtBefore(List<WithdrawalStatus> statuses, LocalDateTime before);
}
