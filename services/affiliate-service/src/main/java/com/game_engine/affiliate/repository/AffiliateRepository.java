package com.game_engine.affiliate.repository;

import com.game_engine.affiliate.entity.Affiliate;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository("affiliateRepositoryV2")
public interface AffiliateRepository extends JpaRepository<Affiliate, Long> {

    Optional<Affiliate> findByAffiliateCode(String affiliateCode);

    Optional<Affiliate> findByEmail(String email);

    List<Affiliate> findByStatus(String status);

    List<Affiliate> findByTier(String tier);

    boolean existsByAffiliateCode(String affiliateCode);

    boolean existsByEmail(String email);
}
