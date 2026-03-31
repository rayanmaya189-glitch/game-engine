package com.game_engine.affiliate.service;

import com.game_engine.affiliate.entity.Affiliate;
import com.game_engine.affiliate.entity.Referral;
import com.game_engine.affiliate.repository.AffiliateRepository;
import com.game_engine.affiliate.repository.ReferralRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@Transactional
public class AffiliateService {

    @Autowired
    @Qualifier("affiliateRepositoryV2")
    private AffiliateRepository affiliateRepository;

    @Autowired
    @Qualifier("referralRepositoryV2")
    private ReferralRepository referralRepository;

    public Affiliate registerAffiliate(String name, String email) {
        if (affiliateRepository.findByEmail(email).isPresent()) {
            throw new RuntimeException("Email already registered");
        }

        Affiliate affiliate = new Affiliate();
        affiliate.setName(name);
        affiliate.setEmail(email);
        affiliate.setAffiliateCode("AFF-" + UUID.randomUUID().toString().substring(0, 8).toUpperCase());

        return affiliateRepository.save(affiliate);
    }

    public Referral trackReferral(String affiliateCode, Long playerId, String source) {
        Affiliate affiliate = affiliateRepository.findByAffiliateCode(affiliateCode)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));

        Referral referral = new Referral();
        referral.setAffiliateCode(affiliateCode);
        referral.setPlayerId(playerId);
        referral.setSource(source);
        referral.setStatus("ACTIVE");

        affiliate.setTotalReferrals(affiliate.getTotalReferrals() + 1);
        affiliateRepository.save(affiliate);

        return referralRepository.save(referral);
    }

    public BigDecimal calculateCommissions(Long affiliateId) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));

        List<Referral> referrals = referralRepository.findByAffiliateCode(affiliate.getAffiliateCode());
        BigDecimal totalDeposits = referrals.stream()
                .map(Referral::getTotalDeposits)
                .reduce(BigDecimal.ZERO, BigDecimal::add);

        BigDecimal commission = totalDeposits.multiply(affiliate.getCommissionRate());

        affiliate.setTotalRevenue(totalDeposits);
        affiliateRepository.save(affiliate);

        return commission;
    }

    public AffiliateStats getAffiliateStats(Long affiliateId) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));

        List<Referral> referrals = referralRepository.findByAffiliateCode(affiliate.getAffiliateCode());
        long totalReferrals = referrals.size();
        long activeReferrals = referrals.stream().filter(r -> "ACTIVE".equals(r.getStatus())).count();
        BigDecimal totalDeposits = referrals.stream()
                .map(r -> r.getTotalDeposits() != null ? r.getTotalDeposits() : BigDecimal.ZERO)
                .reduce(BigDecimal.ZERO, BigDecimal::add);
        BigDecimal totalCommission = totalDeposits.multiply(affiliate.getCommissionRate());

        return new AffiliateStats(totalReferrals, activeReferrals, totalDeposits, totalCommission);
    }

    public Optional<Affiliate> getAffiliateByCode(String affiliateCode) {
        return affiliateRepository.findByAffiliateCode(affiliateCode);
    }

    public Optional<Affiliate> getAffiliateById(Long id) {
        return affiliateRepository.findById(id);
    }

    public List<Affiliate> getAllAffiliates() {
        return affiliateRepository.findAll();
    }

    public List<Affiliate> getAffiliatesByStatus(String status) {
        return affiliateRepository.findByStatus(status);
    }

    public List<Referral> getReferralsByAffiliateCode(String affiliateCode) {
        return referralRepository.findByAffiliateCode(affiliateCode);
    }

    public Affiliate updateAffiliateTier(Long affiliateId, String tier) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));

        affiliate.setTier(tier);
        affiliate.setCommissionRate(getRateForTier(tier));
        return affiliateRepository.save(affiliate);
    }

    public Affiliate updateAffiliateStatus(Long affiliateId, String status) {
        Affiliate affiliate = affiliateRepository.findById(affiliateId)
                .orElseThrow(() -> new RuntimeException("Affiliate not found"));

        affiliate.setStatus(status);
        return affiliateRepository.save(affiliate);
    }

    private BigDecimal getRateForTier(String tier) {
        return switch (tier) {
            case "DIAMOND" -> new BigDecimal("0.4000");
            case "PLATINUM" -> new BigDecimal("0.3500");
            case "GOLD" -> new BigDecimal("0.3000");
            case "SILVER" -> new BigDecimal("0.2500");
            default -> new BigDecimal("0.2000");
        };
    }

    public static class AffiliateStats {
        private final long totalReferrals;
        private final long activeReferrals;
        private final BigDecimal totalDeposits;
        private final BigDecimal totalCommission;

        public AffiliateStats(long totalReferrals, long activeReferrals, BigDecimal totalDeposits, BigDecimal totalCommission) {
            this.totalReferrals = totalReferrals;
            this.activeReferrals = activeReferrals;
            this.totalDeposits = totalDeposits;
            this.totalCommission = totalCommission;
        }

        public long getTotalReferrals() { return totalReferrals; }
        public long getActiveReferrals() { return activeReferrals; }
        public BigDecimal getTotalDeposits() { return totalDeposits; }
        public BigDecimal getTotalCommission() { return totalCommission; }
    }
}
