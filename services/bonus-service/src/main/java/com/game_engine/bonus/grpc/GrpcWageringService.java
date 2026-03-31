package com.game_engine.bonus.grpc;

import com.game_engine.bonus.service.BonusService;
import com.game_engine.bonus.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;

import java.util.List;
import java.util.Map;
import java.util.UUID;

@GrpcService
@RequiredArgsConstructor
@Slf4j
public class GrpcWageringService extends BonusServiceGrpc.BonusServiceImplBase {

    private final BonusService bonusService;

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
}
