package com.game_engine.kyc.repository;

import com.game_engine.kyc.model.KycVerification;
import com.game_engine.kyc.model.KycVerification.KycStatus;
import com.game_engine.kyc.model.KycVerification.VerificationLevel;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
public interface KycRepository extends JpaRepository<KycVerification, UUID> {

    Optional<KycVerification> findByUserId(UUID userId);

    Page<KycVerification> findByStatus(KycStatus status, Pageable pageable);

    Page<KycVerification> findByLevel(VerificationLevel level, Pageable pageable);

    Page<KycVerification> findByLevelAndStatus(VerificationLevel level, KycStatus status, Pageable pageable);

    boolean existsByUserIdAndLevelGreaterThanEqual(UUID userId, VerificationLevel level);
}
