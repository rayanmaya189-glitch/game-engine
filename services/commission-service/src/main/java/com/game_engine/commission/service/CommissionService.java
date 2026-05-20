package com.game_engine.commission.service;

import com.game_engine.commission.entity.CommissionClaim;
import com.game_engine.commission.entity.CommissionConfig;
import com.game_engine.commission.entity.CommissionSettlement;
import com.game_engine.commission.repository.CommissionClaimV2Repository;
import com.game_engine.commission.repository.CommissionConfigV2Repository;
import com.game_engine.commission.repository.CommissionSettlementRepository;
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
public class CommissionService {

    @Autowired
    @Qualifier("commissionConfigRepositoryV2")
    private CommissionConfigV2Repository configRepository;

    @Autowired
    @Qualifier("commissionClaimRepositoryV2")
    private CommissionClaimV2Repository claimRepository;

    @Autowired
    private CommissionSettlementRepository settlementRepository;

    public BigDecimal calculateCommission(Long agentId, Long affiliateId, BigDecimal grossRevenue, int newPlayers) {
        Optional<CommissionConfig> configOpt = configRepository
                .findByAgentIdAndAffiliateIdAndIsActive(agentId, affiliateId, true)
                .stream().findFirst();

        if (configOpt.isEmpty()) {
            throw new RuntimeException("No active commission config found for agent/affiliate");
        }

        CommissionConfig config = configOpt.get();
        BigDecimal commission;

        switch (config.getCommissionType()) {
            case "REVENUE_SHARE" -> commission = grossRevenue.multiply(config.getRate());
            case "CPA" -> commission = config.getRate().multiply(BigDecimal.valueOf(newPlayers));
            case "HYBRID" -> {
                BigDecimal revShare = grossRevenue.multiply(config.getRate());
                BigDecimal cpa = config.getRate().multiply(BigDecimal.valueOf(newPlayers));
                commission = revShare.add(cpa);
            }
            default -> throw new RuntimeException("Unknown commission type: " + config.getCommissionType());
        }

        if (config.getMaxCommission() != null && commission.compareTo(config.getMaxCommission()) > 0) {
            commission = config.getMaxCommission();
        }

        return commission;
    }

    public CommissionClaim createClaim(Long agentId, Long affiliateId, Long playerId,
                                        String period, BigDecimal grossRevenue) {
        BigDecimal commissionAmount = calculateCommission(agentId, affiliateId, grossRevenue, 0);

        CommissionClaim claim = new CommissionClaim();
        claim.setClaimId(UUID.randomUUID().toString());
        claim.setAgentId(agentId);
        claim.setAffiliateId(affiliateId);
        claim.setPlayerId(playerId);
        claim.setPeriod(period);
        claim.setGrossRevenue(grossRevenue);
        claim.setCommissionAmount(commissionAmount);
        claim.setStatus("PENDING");

        return claimRepository.save(claim);
    }

    public CommissionClaim approveClaim(Long claimId) {
        CommissionClaim claim = claimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Commission claim not found"));

        if (!"PENDING".equals(claim.getStatus())) {
            throw new RuntimeException("Claim is not in PENDING status");
        }

        claim.setStatus("APPROVED");
        claim.setProcessedAt(LocalDateTime.now());
        return claimRepository.save(claim);
    }

    public CommissionClaim rejectClaim(Long claimId) {
        CommissionClaim claim = claimRepository.findById(claimId)
                .orElseThrow(() -> new RuntimeException("Commission claim not found"));

        if (!"PENDING".equals(claim.getStatus())) {
            throw new RuntimeException("Claim is not in PENDING status");
        }

        claim.setStatus("REJECTED");
        claim.setProcessedAt(LocalDateTime.now());
        return claimRepository.save(claim);
    }

    public CommissionSettlement processSettlement(Long agentId, Long affiliateId,
                                                    LocalDateTime periodStart, LocalDateTime periodEnd) {
        List<CommissionClaim> approvedClaims = claimRepository.findByAffiliateIdAndStatus(affiliateId, "APPROVED");

        BigDecimal totalRevenue = approvedClaims.stream()
                .map(CommissionClaim::getGrossRevenue)
                .reduce(BigDecimal.ZERO, BigDecimal::add);

        BigDecimal totalCommission = approvedClaims.stream()
                .map(CommissionClaim::getCommissionAmount)
                .reduce(BigDecimal.ZERO, BigDecimal::add);

        CommissionSettlement settlement = new CommissionSettlement();
        settlement.setSettlementId(UUID.randomUUID().toString());
        settlement.setAgentId(agentId);
        settlement.setAffiliateId(affiliateId);
        settlement.setPeriodStart(periodStart);
        settlement.setPeriodEnd(periodEnd);
        settlement.setTotalRevenue(totalRevenue);
        settlement.setTotalCommission(totalCommission);
        settlement.setStatus("PAID");
        settlement.setPaidAt(LocalDateTime.now());

        for (CommissionClaim claim : approvedClaims) {
            claim.setStatus("PAID");
            claim.setProcessedAt(LocalDateTime.now());
            claimRepository.save(claim);
        }

        return settlementRepository.save(settlement);
    }

    public Optional<CommissionClaim> getClaimById(Long id) {
        return claimRepository.findById(id);
    }

    public List<CommissionClaim> getClaimsByAgentId(Long agentId) {
        return claimRepository.findByAgentId(agentId);
    }

    public List<CommissionClaim> getClaimsByAffiliateId(Long affiliateId) {
        return claimRepository.findByAffiliateId(affiliateId);
    }

    public List<CommissionClaim> getClaimsByStatus(String status) {
        return claimRepository.findByStatus(status);
    }

    public List<CommissionSettlement> getSettlementsByAffiliateId(Long affiliateId) {
        return settlementRepository.findByAffiliateId(affiliateId);
    }

    public List<CommissionSettlement> getSettlementsByAgentId(Long agentId) {
        return settlementRepository.findByAgentId(agentId);
    }

    public CommissionConfig createConfig(CommissionConfig config) {
        return configRepository.save(config);
    }

    public Optional<CommissionConfig> getConfigById(Long id) {
        return configRepository.findById(id);
    }

    public List<CommissionConfig> getConfigsByAgentId(Long agentId) {
        return configRepository.findByAgentId(agentId);
    }

    public List<CommissionConfig> getConfigsByAffiliateId(Long affiliateId) {
        return configRepository.findByAffiliateId(affiliateId);
    }

    public CommissionConfig updateConfig(Long id, CommissionConfig updatedConfig) {
        CommissionConfig existing = configRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Commission config not found"));

        existing.setCommissionType(updatedConfig.getCommissionType());
        existing.setRate(updatedConfig.getRate());
        existing.setMinDeposit(updatedConfig.getMinDeposit());
        existing.setMaxCommission(updatedConfig.getMaxCommission());
        existing.setIsActive(updatedConfig.getIsActive());
        existing.setUpdatedAt(LocalDateTime.now());

        return configRepository.save(existing);
    }

    public void deleteConfig(Long id) {
        configRepository.deleteById(id);
    }
}
