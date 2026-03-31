package com.game_engine.affiliate.repository;

import com.game_engine.affiliate.entity.Referral;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository("referralRepositoryV2")
public interface ReferralRepository extends JpaRepository<Referral, Long> {

    List<Referral> findByAffiliateCode(String affiliateCode);

    List<Referral> findByPlayerId(Long playerId);

    List<Referral> findByStatus(String status);

    List<Referral> findByAffiliateCodeAndStatus(String affiliateCode, String status);

    long countByAffiliateCode(String affiliateCode);
}
