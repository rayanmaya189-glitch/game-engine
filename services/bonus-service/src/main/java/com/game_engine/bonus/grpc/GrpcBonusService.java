package com.game_engine.bonus.grpc;

import com.game_engine.bonus.model.Bonus;
import com.game_engine.bonus.service.BonusService;
import com.game_engine.bonus.repository.BonusRepository;
import com.game_engine.bonus.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;

import java.time.Instant;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@GrpcService
@RequiredArgsConstructor
@Slf4j
public class GrpcBonusService extends BonusServiceGrpc.BonusServiceImplBase {

    private final BonusService bonusService;
    private final BonusRepository bonusRepository;

    @Override
    public void getActiveBonuses(GetActiveBonusesRequest request, StreamObserver<GetActiveBonusesResponse> responseObserver) {
        try {
            List<Bonus> bonuses = bonusRepository.findAll().stream()
                    .filter(b -> b.getStatus() == Bonus.BonusStatus.ACTIVE)
                    .toList();

            GetActiveBonusesResponse.Builder response = GetActiveBonusesResponse.newBuilder();
            for (Bonus bonus : bonuses) {
                response.addBonuses(toProtoBonus(bonus));
            }
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting active bonuses", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getBonus(GetBonusRequest request, StreamObserver<GetBonusResponse> responseObserver) {
        try {
            UUID id = UUID.fromString(request.getId());
            Bonus bonus = bonusRepository.findById(id)
                    .orElseThrow(() -> new RuntimeException("Bonus not found: " + id));

            responseObserver.onNext(GetBonusResponse.newBuilder()
                    .setBonus(toProtoBonus(bonus))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting bonus: {}", request.getId(), e);
            responseObserver.onError(io.grpc.Status.NOT_FOUND
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void createBonus(CreateBonusRequest request, StreamObserver<CreateBonusResponse> responseObserver) {
        try {
            Bonus bonus = new Bonus();
            BonusServiceProto.Bonus reqBonus = request.getBonus();
            bonus.setName(reqBonus.getName());
            bonus.setDescription(reqBonus.getDescription());
            bonus.setType(Bonus.BonusType.valueOf(reqBonus.getType()));
            bonus.setStatus(Bonus.BonusStatus.ACTIVE);
            bonus.setAmount(java.math.BigDecimal.valueOf(reqBonus.getAmount()));
            bonus.setPercentage(java.math.BigDecimal.valueOf(reqBonus.getPercentage()));
            bonus.setMaxAmount(java.math.BigDecimal.valueOf(reqBonus.getMaxAmount()));
            bonus.setMinDeposit(java.math.BigDecimal.valueOf(reqBonus.getMinDeposit()));
            bonus.setWageringRequirement(reqBonus.getWageringRequirement());
            bonus.setMaxBet(java.math.BigDecimal.valueOf(reqBonus.getMaxBet()));
            bonus.setAllowedGames(reqBonus.getAllowedGames());
            bonus.setVipLevel(Bonus.VIPLevel.valueOf(reqBonus.getVipLevel()));

            Bonus saved = bonusRepository.save(bonus);

            responseObserver.onNext(CreateBonusResponse.newBuilder()
                    .setBonus(toProtoBonus(saved))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error creating bonus", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void claimBonus(ClaimBonusRequest request, StreamObserver<ClaimBonusResponse> responseObserver) {
        try {
            UUID bonusId = UUID.fromString(request.getBonusId());
            UUID userId = UUID.fromString(request.getUserId());

            Map<String, Object> result = bonusService.claimBonus(bonusId, userId);

            ClaimBonusResponse.Builder response = ClaimBonusResponse.newBuilder()
                    .setSuccess(true)
                    .setMessage("Bonus claimed successfully");
            if (result.containsKey("bonusAmount")) {
                response.setBonusAmount(((Number) result.get("bonusAmount")).doubleValue());
            }
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error claiming bonus", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void checkEligibility(CheckEligibilityRequest request, StreamObserver<CheckEligibilityResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            Map<String, Object> eligibility = bonusService.checkEligibility(userId);

            CheckEligibilityResponse.Builder response = CheckEligibilityResponse.newBuilder()
                    .setEligible((Boolean) eligibility.getOrDefault("eligible", false));
            eligibility.forEach((k, v) -> response.putDetails(k, String.valueOf(v)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error checking eligibility", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getBonusHistory(GetBonusHistoryRequest request, StreamObserver<GetBonusHistoryResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            List<Map<String, Object>> history = bonusService.getBonusHistory(userId);

            GetBonusHistoryResponse.Builder response = GetBonusHistoryResponse.newBuilder();
            for (Map<String, Object> entry : history) {
                BonusHistoryEntry.Builder entryBuilder = BonusHistoryEntry.newBuilder()
                        .setBonusId(String.valueOf(entry.getOrDefault("bonusId", "")))
                        .setBonusName(String.valueOf(entry.getOrDefault("bonusName", "")))
                        .setStatus(String.valueOf(entry.getOrDefault("status", "")));
                if (entry.containsKey("amount")) {
                    entryBuilder.setAmount(((Number) entry.get("amount")).doubleValue());
                }
                response.addEntries(entryBuilder.build());
            }

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting bonus history", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getActiveBonusClaims(GetActiveBonusClaimsRequest request, StreamObserver<GetActiveBonusClaimsResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            List<Map<String, Object>> activeClaims = bonusService.getActiveBonusClaims(userId);

            GetActiveBonusClaimsResponse.Builder response = GetActiveBonusClaimsResponse.newBuilder();
            for (Map<String, Object> claim : activeClaims) {
                ActiveBonusClaim.Builder claimBuilder = ActiveBonusClaim.newBuilder()
                        .setBonusId(String.valueOf(claim.getOrDefault("bonusId", "")))
                        .setBonusName(String.valueOf(claim.getOrDefault("bonusName", "")))
                        .setStatus(String.valueOf(claim.getOrDefault("status", "")));
                if (claim.containsKey("amount")) claimBuilder.setAmount(((Number) claim.get("amount")).doubleValue());
                if (claim.containsKey("wageringProgress")) claimBuilder.setWageringProgress(((Number) claim.get("wageringProgress")).doubleValue());
                if (claim.containsKey("wageringRequired")) claimBuilder.setWageringRequired(((Number) claim.get("wageringRequired")).doubleValue());
                response.addClaims(claimBuilder.build());
            }

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting active bonus claims", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void processWageringContribution(ProcessWageringContributionRequest request, StreamObserver<ProcessWageringContributionResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            UUID bonusId = UUID.fromString(request.getBonusId());
            java.math.BigDecimal betAmount = java.math.BigDecimal.valueOf(request.getBetAmount());
            String gameType = request.getGameType();

            Map<String, Object> result = bonusService.processWageringContribution(userId, bonusId, betAmount, gameType);

            ProcessWageringContributionResponse.Builder response = ProcessWageringContributionResponse.newBuilder()
                    .setSuccess((Boolean) result.getOrDefault("success", false))
                    .setCompleted((Boolean) result.getOrDefault("completed", false));
            if (result.containsKey("wageringProgress")) response.setWageringProgress(((Number) result.get("wageringProgress")).doubleValue());
            if (result.containsKey("wageringRequired")) response.setWageringRequired(((Number) result.get("wageringRequired")).doubleValue());

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error processing wagering contribution", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void completeBonus(CompleteBonusRequest request, StreamObserver<CompleteBonusResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            UUID bonusId = UUID.fromString(request.getBonusId());
            java.math.BigDecimal winnings = java.math.BigDecimal.valueOf(request.getWinnings());

            Map<String, Object> result = bonusService.completeBonus(userId, bonusId, winnings);

            CompleteBonusResponse.Builder response = CompleteBonusResponse.newBuilder()
                    .setSuccess((Boolean) result.getOrDefault("success", false))
                    .setMessage(String.valueOf(result.getOrDefault("message", "Bonus completed")));
            if (result.containsKey("payoutAmount")) response.setPayoutAmount(((Number) result.get("payoutAmount")).doubleValue());

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error completing bonus", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void cancelBonus(CancelBonusRequest request, StreamObserver<CancelBonusResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            UUID bonusId = UUID.fromString(request.getBonusId());
            String reason = request.getReason();

            Map<String, Object> result = bonusService.cancelBonus(userId, bonusId, reason);

            CancelBonusResponse response = CancelBonusResponse.newBuilder()
                    .setSuccess((Boolean) result.getOrDefault("success", false))
                    .setMessage(String.valueOf(result.getOrDefault("message", "Bonus cancelled")))
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error cancelling bonus", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getBonusStats(GetBonusStatsRequest request, StreamObserver<GetBonusStatsResponse> responseObserver) {
        try {
            UUID userId = UUID.fromString(request.getUserId());
            Map<String, Object> stats = bonusService.getBonusStats(userId);

            GetBonusStatsResponse.Builder response = GetBonusStatsResponse.newBuilder();
            stats.forEach((k, v) -> response.putStats(k, String.valueOf(v)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting bonus stats", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private BonusServiceProto.Bonus toProtoBonus(Bonus bonus) {
        BonusServiceProto.Bonus.Builder builder = BonusServiceProto.Bonus.newBuilder()
                .setId(bonus.getId().toString())
                .setName(bonus.getName() != null ? bonus.getName() : "")
                .setDescription(bonus.getDescription() != null ? bonus.getDescription() : "")
                .setType(bonus.getType() != null ? bonus.getType().name() : "")
                .setStatus(bonus.getStatus() != null ? bonus.getStatus().name() : "");

        if (bonus.getAmount() != null) builder.setAmount(bonus.getAmount().doubleValue());
        if (bonus.getPercentage() != null) builder.setPercentage(bonus.getPercentage().doubleValue());
        if (bonus.getMaxAmount() != null) builder.setMaxAmount(bonus.getMaxAmount().doubleValue());
        if (bonus.getMinDeposit() != null) builder.setMinDeposit(bonus.getMinDeposit().doubleValue());
        if (bonus.getWageringRequirement() != null) builder.setWageringRequirement(bonus.getWageringRequirement());
        if (bonus.getMaxBet() != null) builder.setMaxBet(bonus.getMaxBet().doubleValue());
        if (bonus.getAllowedGames() != null) builder.setAllowedGames(bonus.getAllowedGames());
        if (bonus.getStartDate() != null) builder.setStartDate(bonus.getStartDate().toEpochMilli());
        if (bonus.getEndDate() != null) builder.setEndDate(bonus.getEndDate().toEpochMilli());
        if (bonus.getMaxUses() != null) builder.setMaxUses(bonus.getMaxUses());
        if (bonus.getCurrentUses() != null) builder.setCurrentUses(bonus.getCurrentUses());
        if (bonus.getVipLevel() != null) builder.setVipLevel(bonus.getVipLevel().name());
        if (bonus.getCreatedAt() != null) builder.setCreatedAt(bonus.getCreatedAt().toEpochMilli());
        if (bonus.getUpdatedAt() != null) builder.setUpdatedAt(bonus.getUpdatedAt().toEpochMilli());

        return builder.build();
    }
}
