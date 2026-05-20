package com.game_engine.commission.grpc;

import com.game_engine.commission.model.*;
import com.game_engine.commission.service.ClaimService;
import com.game_engine.commission.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.math.BigDecimal;
import java.time.ZoneOffset;
import java.util.List;
import java.util.Optional;
import com.google.protobuf.Timestamp;

@GrpcService
@Slf4j
public class GrpcClaimService extends ClaimServiceGrpc.ClaimServiceImplBase {

    @Autowired
    private ClaimService claimService;

    // Commission Claims
    @Override
    public void submitCommissionClaim(SubmitCommissionClaimRequest request,
            StreamObserver<SubmitCommissionClaimResponse> responseObserver) {
        try {
            CommissionClaim claim = claimService.submitCommissionClaim(
                    request.getUserId(), request.getAffiliateId(), request.getCommissionId(),
                    BigDecimal.valueOf(request.getAmount()), request.getClaimReason());
            responseObserver.onNext(SubmitCommissionClaimResponse.newBuilder()
                    .setClaim(toProtoCommissionClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserCommissionClaims(GetUserCommissionClaimsRequest request,
            StreamObserver<GetUserCommissionClaimsResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = claimService.getUserCommissionClaims(request.getUserId());
            GetUserCommissionClaimsResponse.Builder response = GetUserCommissionClaimsResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoCommissionClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionClaimsByStatus(GetCommissionClaimsByStatusRequest request,
            StreamObserver<GetCommissionClaimsByStatusResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = claimService.getCommissionClaimsByStatus(request.getStatus());
            GetCommissionClaimsByStatusResponse.Builder response = GetCommissionClaimsByStatusResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoCommissionClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionClaimById(GetCommissionClaimByIdRequest request,
            StreamObserver<GetCommissionClaimByIdResponse> responseObserver) {
        try {
            Optional<CommissionClaim> claim = claimService.getCommissionClaimById(request.getId());
            GetCommissionClaimByIdResponse.Builder response = GetCommissionClaimByIdResponse.newBuilder()
                    .setFound(claim.isPresent());
            claim.ifPresent(c -> response.setClaim(toProtoCommissionClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    // Rebet Claims
    @Override
    public void createRebetClaim(CreateRebetClaimRequest request,
            StreamObserver<CreateRebetClaimResponse> responseObserver) {
        try {
            RebetClaim claim = claimService.createRebetClaim(
                    request.getUserId(), request.getBonusId(), request.getBonusCode(),
                    BigDecimal.valueOf(request.getBonusAmount()), BigDecimal.valueOf(request.getRebetRequirement()),
                    request.getGameId(), request.getBetId());
            responseObserver.onNext(CreateRebetClaimResponse.newBuilder()
                    .setClaim(toProtoRebetClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void updateRebetProgress(UpdateRebetProgressRequest request,
            StreamObserver<UpdateRebetProgressResponse> responseObserver) {
        try {
            RebetClaim claim = claimService.updateRebetProgress(request.getId(),
                    BigDecimal.valueOf(request.getAdditionalBetAmount()));
            responseObserver.onNext(UpdateRebetProgressResponse.newBuilder()
                    .setClaim(toProtoRebetClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void claimRebet(ClaimRebetRequest request, StreamObserver<ClaimRebetResponse> responseObserver) {
        try {
            RebetClaim claim = claimService.claimRebet(request.getId());
            responseObserver.onNext(ClaimRebetResponse.newBuilder()
                    .setClaim(toProtoRebetClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserRebetClaims(GetUserRebetClaimsRequest request,
            StreamObserver<GetUserRebetClaimsResponse> responseObserver) {
        try {
            List<RebetClaim> claims = claimService.getUserRebetClaims(request.getUserId());
            GetUserRebetClaimsResponse.Builder response = GetUserRebetClaimsResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoRebetClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getClaimableRebets(GetClaimableRebetsRequest request,
            StreamObserver<GetClaimableRebetsResponse> responseObserver) {
        try {
            List<RebetClaim> claims = claimService.getClaimableRebets(request.getUserId());
            GetClaimableRebetsResponse.Builder response = GetClaimableRebetsResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoRebetClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    // Insurance Claims
    @Override
    public void submitInsuranceClaim(SubmitInsuranceClaimRequest request,
            StreamObserver<SubmitInsuranceClaimResponse> responseObserver) {
        try {
            InsuranceClaim claim = claimService.submitInsuranceClaim(
                    request.getUserId(), request.getGameId(), request.getBetId(),
                    request.getInsurancePolicyId(), request.getClaimType(),
                    BigDecimal.valueOf(request.getInsuredAmount()), BigDecimal.valueOf(request.getLossAmount()),
                    request.getClaimReason(),
                    request.getEvidenceDetails().isEmpty() ? null : request.getEvidenceDetails());
            responseObserver.onNext(SubmitInsuranceClaimResponse.newBuilder()
                    .setClaim(toProtoInsuranceClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserInsuranceClaims(GetUserInsuranceClaimsRequest request,
            StreamObserver<GetUserInsuranceClaimsResponse> responseObserver) {
        try {
            List<InsuranceClaim> claims = claimService.getUserInsuranceClaims(request.getUserId());
            GetUserInsuranceClaimsResponse.Builder response = GetUserInsuranceClaimsResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoInsuranceClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getInsuranceClaimsByStatus(GetInsuranceClaimsByStatusRequest request,
            StreamObserver<GetInsuranceClaimsByStatusResponse> responseObserver) {
        try {
            List<InsuranceClaim> claims = claimService.getInsuranceClaimsByStatus(request.getStatus());
            GetInsuranceClaimsByStatusResponse.Builder response = GetInsuranceClaimsByStatusResponse.newBuilder();
            claims.forEach(c -> response.addClaims(toProtoInsuranceClaim(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    // Conversion helpers
    private com.game_engine.commission.v1.CommissionClaimMsg toProtoCommissionClaim(CommissionClaim c) {
        com.game_engine.commission.v1.CommissionClaimMsg.Builder builder = com.game_engine.commission.v1.CommissionClaimMsg
                .newBuilder()
                .setId(c.getId())
                .setStatus(c.getStatus() != null ? c.getStatus() : "");

        if (c.getUserId() != null)
            builder.setUserId(c.getUserId());
        if (c.getAffiliateId() != null)
            builder.setAffiliateId(c.getAffiliateId());
        if (c.getCommissionId() != null)
            builder.setCommissionId(c.getCommissionId());
        if (c.getClaimType() != null)
            builder.setClaimType(c.getClaimType());
        if (c.getAmount() != null)
            builder.setAmount(c.getAmount().doubleValue());
        if (c.getClaimReason() != null)
            builder.setClaimReason(c.getClaimReason());
        if (c.getAdminNote() != null)
            builder.setAdminNote(c.getAdminNote());
        if (c.getRequestedAt() != null)
            builder.setRequestedAt(c.getRequestedAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getProcessedAt() != null)
            builder.setProcessedAt(c.getProcessedAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getPaidAt() != null)
            builder.setPaidAt(c.getPaidAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getTransactionId() != null)
            builder.setTransactionId(c.getTransactionId());

        return builder.build();
    }

    private com.game_engine.commission.v1.RebetClaimMsg toProtoRebetClaim(RebetClaim c) {
        com.game_engine.commission.v1.RebetClaimMsg.Builder builder = com.game_engine.commission.v1.RebetClaimMsg
                .newBuilder()
                .setId(c.getId())
                .setStatus(c.getStatus() != null ? c.getStatus() : "")
                .setClaimable(c.isClaimable());

        if (c.getUserId() != null)
            builder.setUserId(c.getUserId());
        if (c.getBonusId() != null)
            builder.setBonusId(c.getBonusId());
        if (c.getBonusCode() != null)
            builder.setBonusCode(c.getBonusCode());
        if (c.getGameId() != null)
            builder.setGameId(c.getGameId());
        if (c.getBetId() != null)
            builder.setBetId(c.getBetId());
        if (c.getOriginalBonusAmount() != null)
            builder.setOriginalBonusAmount(c.getOriginalBonusAmount().doubleValue());
        if (c.getRebetRequirement() != null)
            builder.setRebetRequirement(c.getRebetRequirement().doubleValue());
        if (c.getCurrentRebetAmount() != null)
            builder.setCurrentRebetAmount(c.getCurrentRebetAmount().doubleValue());
        if (c.getClaimAmount() != null)
            builder.setClaimAmount(c.getClaimAmount().doubleValue());
        if (c.getExpiresAt() != null)
            builder.setExpiresAt(c.getExpiresAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getClaimedAt() != null)
            builder.setClaimedAt(c.getClaimedAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getTransactionId() != null)
            builder.setTransactionId(c.getTransactionId());

        return builder.build();
    }

    private com.game_engine.commission.v1.InsuranceClaimMsg toProtoInsuranceClaim(InsuranceClaim c) {
        com.game_engine.commission.v1.InsuranceClaimMsg.Builder builder = com.game_engine.commission.v1.InsuranceClaimMsg
                .newBuilder()
                .setId(c.getId())
                .setStatus(c.getStatus() != null ? c.getStatus() : "");

        if (c.getUserId() != null)
            builder.setUserId(c.getUserId());
        if (c.getGameId() != null)
            builder.setGameId(c.getGameId());
        if (c.getBetId() != null)
            builder.setBetId(c.getBetId());
        if (c.getInsurancePolicyId() != null)
            builder.setInsurancePolicyId(c.getInsurancePolicyId());
        if (c.getClaimType() != null)
            builder.setClaimType(c.getClaimType());
        if (c.getInsuredAmount() != null)
            builder.setInsuredAmount(c.getInsuredAmount().doubleValue());
        if (c.getLossAmount() != null)
            builder.setLossAmount(c.getLossAmount().doubleValue());
        if (c.getClaimAmount() != null)
            builder.setClaimAmount(c.getClaimAmount().doubleValue());
        if (c.getClaimReason() != null)
            builder.setClaimReason(c.getClaimReason());
        if (c.getEvidenceDetails() != null)
            builder.setEvidenceDetails(c.getEvidenceDetails());
        if (c.getAdminNote() != null)
            builder.setAdminNote(c.getAdminNote());
        if (c.getReviewedBy() != null)
            builder.setReviewedBy(c.getReviewedBy());
        if (c.getRequestedAt() != null)
            builder.setRequestedAt(c.getRequestedAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getReviewedAt() != null)
            builder.setReviewedAt(c.getReviewedAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getPaidAt() != null)
            builder.setPaidAt(c.getPaidAt().toEpochSecond(java.time.ZoneOffset.UTC));
        if (c.getTransactionId() != null)
            builder.setTransactionId(c.getTransactionId());

        return builder.build();
    }
}
