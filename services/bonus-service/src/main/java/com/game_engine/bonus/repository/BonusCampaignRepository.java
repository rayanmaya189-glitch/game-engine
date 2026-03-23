package com.game_engine.bonus.repository;

import com.game_engine.bonus.model.BonusCampaign;
import com.game_engine.bonus.model.BonusCampaign.BonusType;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface BonusCampaignRepository extends JpaRepository<BonusCampaign, UUID> {

    @Query("SELECT c FROM BonusCampaign c WHERE c.isActive = true AND c.startDate <= :now AND (c.endDate IS NULL OR c.endDate >= :now)")
    List<BonusCampaign> findActiveCampaigns(LocalDateTime now);

    @Query("SELECT c FROM BonusCampaign c WHERE c.isActive = true AND c.bonusType = :type AND c.startDate <= :now AND (c.endDate IS NULL OR c.endDate >= :now)")
    List<BonusCampaign> findActiveByType(BonusType type, LocalDateTime now);

    @Query("SELECT c FROM BonusCampaign c WHERE c.bonusType = 'CASHBACK' AND c.isActive = true AND c.startDate <= :now AND (c.endDate IS NULL OR c.endDate >= :now)")
    Optional<BonusCampaign> findActiveCashbackCampaign(LocalDateTime now);

    Optional<BonusCampaign> findByPromoCode(String promoCode);
}
