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
public class GrpcClaimApprovalService extends ClaimServiceGrpc.ClaimServiceImplBase {

    @Autowired
    private ClaimService claimService;

    // Commission Claim Approval/Rejection/Payment
    @Override
    public void approveCommissionClaim(ApproveCommissionClaimRequest request,
            StreamObserver<ApproveCommissionClaimResponse> responseObserver) {
        try {
            CommissionClaim claim = claimService.approveCommissionClaim(request.getId(), request.getAdminNote());
            responseObserver.onNext(ApproveCommissionClaimResponse.newBuilder()
                    .setClaim(toProtoCommissionClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void rejectCommissionClaim(RejectCommissionClaimRequest request,
            StreamObserver<RejectCommissionClaimResponse> responseObserver) {
        try {
            CommissionClaim claim = claimService.rejectCommissionClaim(request.getId(), request.getAdminNote());
            responseObserver.onNext(RejectCommissionClaimResponse.newBuilder()
                    .setClaim(toProtoCommissionClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void payCommissionClaim(PayCommissionClaimRequest request,
            StreamObserver<PayCommissionClaimResponse> responseObserver) {
        try {
            CommissionClaim claim = claimService.payCommissionClaim(request.getId());
            responseObserver.onNext(PayCommissionClaimResponse.newBuilder()
                    .setClaim(toProtoCommissionClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    // Insurance Claim Approval/Rejection/Payment
    @Override
    public void approveInsuranceClaim(ApproveInsuranceClaimRequest request,
            StreamObserver<ApproveInsuranceClaimResponse> responseObserver) {
        try {
            InsuranceClaim claim = claimService.approveInsuranceClaim(request.getId(), request.getReviewedBy(),
                    request.getAdminNote());
            responseObserver.onNext(ApproveInsuranceClaimResponse.newBuilder()
                    .setClaim(toProtoInsuranceClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void rejectInsuranceClaim(RejectInsuranceClaimRequest request,
            StreamObserver<RejectInsuranceClaimResponse> responseObserver) {
        try {
            InsuranceClaim claim = claimService.rejectInsuranceClaim(request.getId(), request.getReviewedBy(),
                    request.getAdminNote());
            responseObserver.onNext(RejectInsuranceClaimResponse.newBuilder()
                    .setClaim(toProtoInsuranceClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void payInsuranceClaim(PayInsuranceClaimRequest request,
            StreamObserver<PayInsuranceClaimResponse> responseObserver) {
        try {
            InsuranceClaim claim = claimService.payInsuranceClaim(request.getId());
            responseObserver.onNext(PayInsuranceClaimResponse.newBuilder()
                    .setClaim(toProtoInsuranceClaim(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    // Settlements
    @Override
    public void getUserSettlements(GetUserSettlementsRequest request,
            StreamObserver<GetUserSettlementsResponse> responseObserver) {
        try {
            List<Settlement> settlements = claimService.getUserSettlements(request.getUserId());
            GetUserSettlementsResponse.Builder response = GetUserSettlementsResponse.newBuilder();
            settlements.forEach(s -> response.addSettlements(toProtoSettlement(s)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSettlementsByStatus(GetSettlementsByStatusRequest request,
            StreamObserver<GetSettlementsByStatusResponse> responseObserver) {
        try {
            List<Settlement> settlements = claimService.getSettlementsByStatus(request.getStatus());
            GetSettlementsByStatusResponse.Builder response = GetSettlementsByStatusResponse.newBuilder();
            settlements.forEach(s -> response.addSettlements(toProtoSettlement(s)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSettlementsByType(GetSettlementsByTypeRequest request,
            StreamObserver<GetSettlementsByTypeResponse> responseObserver) {
        try {
            List<Settlement> settlements = claimService.getSettlementsByType(request.getType());
            GetSettlementsByTypeResponse.Builder response = GetSettlementsByTypeResponse.newBuilder();
            settlements.forEach(s -> response.addSettlements(toProtoSettlement(s)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSettlementById(GetSettlementByIdRequest request,
            StreamObserver<GetSettlementByIdResponse> responseObserver) {
        try {
            Optional<Settlement> settlement = claimService.getSettlementById(request.getId());
            GetSettlementByIdResponse.Builder response = GetSettlementByIdResponse.newBuilder()
                    .setFound(settlement.isPresent());
            settlement.ifPresent(s -> response.setSettlement(toProtoSettlement(s)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserTotalPending(GetUserTotalPendingRequest request,
            StreamObserver<GetUserTotalPendingResponse> responseObserver) {
        try {
            BigDecimal total = claimService.getUserTotalPendingClaims(request.getUserId());
            responseObserver
                    .onNext(GetUserTotalPendingResponse.newBuilder().setTotalPending(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getUserTotalSettled(GetUserTotalSettledRequest request,
            StreamObserver<GetUserTotalSettledResponse> responseObserver) {
        try {
            BigDecimal total = claimService.getUserTotalSettled(request.getUserId());
            responseObserver
                    .onNext(GetUserTotalSettledResponse.newBuilder().setTotalSettled(total.doubleValue()).build());
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
            builder.setRequestedAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getRequestedAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
        if (c.getProcessedAt() != null)
            builder.setProcessedAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getProcessedAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
        if (c.getPaidAt() != null)
            builder.setPaidAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getPaidAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
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
            builder.setRequestedAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getRequestedAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
        if (c.getReviewedAt() != null)
            builder.setReviewedAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getReviewedAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
        if (c.getPaidAt() != null)
            builder.setPaidAt(
                    Timestamp.newBuilder()
                            .setSeconds(c.getPaidAt().toEpochSecond(ZoneOffset.UTC))
                            .build());
        if (c.getTransactionId() != null)
            builder.setTransactionId(c.getTransactionId());

        return builder.build();
    }

    private com.game_engine.commission.v1.SettlementMsg toProtoSettlement(Settlement s) {
        com.game_engine.commission.v1.SettlementMsg.Builder builder = com.game_engine.commission.v1.SettlementMsg
                .newBuilder()
                .setId(s.getId())
                .setStatus(s.getStatus() != null ? s.getStatus() : "");

        if (s.getUserId() != null)
            builder.setUserId(s.getUserId());
        if (s.getSettlementType() != null)
            builder.setSettlementType(s.getSettlementType());
        if (s.getClaimId() != null)
            builder.setClaimId(s.getClaimId());
        if (s.getAmount() != null)
            builder.setAmount(s.getAmount().doubleValue());
        if (s.getBonusAmount() != null)
            builder.setBonusAmount(s.getBonusAmount().doubleValue());
        if (s.getRakeAmount() != null)
            builder.setRakeAmount(s.getRakeAmount().doubleValue());
        if (s.getNetAmount() != null)
            builder.setNetAmount(s.getNetAmount().doubleValue());
        if (s.getPaymentMethod() != null)
            builder.setPaymentMethod(s.getPaymentMethod());
        if (s.getTransactionId() != null)
            builder.setTransactionId(s.getTransactionId());
        if (s.getCompletedAt() != null)
            builder.setCompletedAt(
                    Timestamp.newBuilder()
                            .setSeconds(s.getCompletedAt().toEpochSecond(ZoneOffset.UTC))
                            .build());

        return builder.build();
    }
}
