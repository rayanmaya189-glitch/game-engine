package com.game_engine.bonus.repository;

import com.game_engine.bonus.model.PlayerBonus;
import com.game_engine.bonus.model.PlayerBonus.BonusStatus;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface PlayerBonusRepository extends JpaRepository<PlayerBonus, UUID> {

    List<PlayerBonus> findByUserId(UUID userId);

    List<PlayerBonus> findByUserIdAndStatus(UUID userId, BonusStatus status);

    @Query("SELECT pb FROM PlayerBonus pb WHERE pb.userId = :userId AND pb.status = :status")
    List<PlayerBonus> findActiveByUserId(@Param("userId") UUID userId, @Param("status") BonusStatus status);

    boolean existsByUserIdAndCampaignId(UUID userId, UUID campaignId);

    Optional<PlayerBonus> findByUserIdAndCampaignId(UUID userId, UUID campaignId);

    @Query("SELECT COUNT(pb) FROM PlayerBonus pb WHERE pb.userId = :userId AND pb.campaign.id = :campaignId")
    int countByUserAndCampaign(@Param("userId") UUID userId, @Param("campaignId") UUID campaignId);
}
