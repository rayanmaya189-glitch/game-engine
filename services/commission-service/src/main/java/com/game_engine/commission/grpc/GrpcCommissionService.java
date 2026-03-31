package com.game_engine.commission.grpc;

import com.game_engine.commission.entity.CommissionClaim;
import com.game_engine.commission.entity.CommissionConfig;
import com.game_engine.commission.entity.CommissionSettlement;
import com.game_engine.commission.service.CommissionService;
import com.game_engine.commission.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.List;
import java.util.Optional;

@GrpcService
@Slf4j
public class GrpcCommissionService extends CommissionServiceGrpc.CommissionServiceImplBase {

    @Autowired
    private CommissionService commissionService;

    @Override
    public void createCommission(CreateCommissionRequest request, StreamObserver<CreateCommissionResponse> responseObserver) {
        try {
            CommissionClaim claim = commissionService.createClaim(
                    request.getCommission().getAffiliateId(),
                    request.getCommission().getMerchantId(),
                    request.getCommission().getAffiliateId(),
                    request.getCommission().getPeriod(),
                    BigDecimal.valueOf(request.getCommission().getNetRevenue()));

            responseObserver.onNext(CreateCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error creating commission", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionById(GetCommissionByIdRequest request, StreamObserver<GetCommissionByIdResponse> responseObserver) {
        try {
            Optional<CommissionClaim> claim = commissionService.getClaimById(request.getId());
            GetCommissionByIdResponse.Builder response = GetCommissionByIdResponse.newBuilder().setFound(claim.isPresent());
            claim.ifPresent(c -> response.setCommission(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting commission", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionsByAffiliate(GetCommissionsByAffiliateRequest request, StreamObserver<GetCommissionsByAffiliateResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAffiliateId(request.getAffiliateId());
            GetCommissionsByAffiliateResponse.Builder response = GetCommissionsByAffiliateResponse.newBuilder();
            claims.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionsByMerchant(GetCommissionsByMerchantRequest request, StreamObserver<GetCommissionsByMerchantResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAgentId(request.getMerchantId());
            GetCommissionsByMerchantResponse.Builder response = GetCommissionsByMerchantResponse.newBuilder();
            claims.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionsByPeriod(GetCommissionsByPeriodRequest request, StreamObserver<GetCommissionsByPeriodResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByStatus("PENDING");
            GetCommissionsByPeriodResponse.Builder response = GetCommissionsByPeriodResponse.newBuilder();
            claims.stream()
                    .filter(c -> request.getPeriod().equals(c.getPeriod()))
                    .forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionByAffiliateAndPeriod(GetCommissionByAffiliateAndPeriodRequest request, StreamObserver<GetCommissionByAffiliateAndPeriodResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAffiliateId(request.getAffiliateId());
            Optional<CommissionClaim> match = claims.stream()
                    .filter(c -> request.getPeriod().equals(c.getPeriod()))
                    .findFirst();
            GetCommissionByAffiliateAndPeriodResponse.Builder response = GetCommissionByAffiliateAndPeriodResponse.newBuilder()
                    .setFound(match.isPresent());
            match.ifPresent(c -> response.setCommission(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalPaidCommission(GetTotalPaidCommissionRequest request, StreamObserver<GetTotalPaidCommissionResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAffiliateId(request.getAffiliateId());
            BigDecimal total = claims.stream()
                    .filter(c -> "PAID".equals(c.getStatus()))
                    .map(CommissionClaim::getCommissionAmount)
                    .reduce(BigDecimal.ZERO, BigDecimal::add);
            responseObserver.onNext(GetTotalPaidCommissionResponse.newBuilder().setTotalPaid(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalPendingCommission(GetTotalPendingCommissionRequest request, StreamObserver<GetTotalPendingCommissionResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAffiliateId(request.getAffiliateId());
            BigDecimal total = claims.stream()
                    .filter(c -> "PENDING".equals(c.getStatus()))
                    .map(CommissionClaim::getCommissionAmount)
                    .reduce(BigDecimal.ZERO, BigDecimal::add);
            responseObserver.onNext(GetTotalPendingCommissionResponse.newBuilder().setTotalPending(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalRevenueByMerchant(GetTotalRevenueByMerchantRequest request, StreamObserver<GetTotalRevenueByMerchantResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByAgentId(request.getMerchantId());
            BigDecimal total = claims.stream()
                    .map(CommissionClaim::getGrossRevenue)
                    .reduce(BigDecimal.ZERO, BigDecimal::add);
            responseObserver.onNext(GetTotalRevenueByMerchantResponse.newBuilder().setTotalRevenue(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateRevenueShare(CalculateRevenueShareRequest request, StreamObserver<CalculateRevenueShareResponse> responseObserver) {
        try {
            BigDecimal revenue = BigDecimal.valueOf(request.getNetRevenue());
            BigDecimal rate = BigDecimal.valueOf(request.getCommissionRate());
            BigDecimal commission = revenue.multiply(rate);

            CommissionClaim claim = commissionService.createClaim(
                    request.getMerchantId(), request.getAffiliateId(),
                    request.getAffiliateId(),
                    request.getPeriod().isEmpty() ? java.time.YearMonth.now().toString() : request.getPeriod(),
                    revenue);

            responseObserver.onNext(CalculateRevenueShareResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateCPA(CalculateCPARequest request, StreamObserver<CalculateCPAResponse> responseObserver) {
        try {
            BigDecimal cpaRate = BigDecimal.valueOf(request.getCpaRate());
            BigDecimal commission = cpaRate.multiply(BigDecimal.valueOf(request.getNewPlayers()));

            CommissionClaim claim = commissionService.createClaim(
                    request.getMerchantId(), request.getAffiliateId(),
                    request.getAffiliateId(),
                    request.getPeriod().isEmpty() ? java.time.YearMonth.now().toString() : request.getPeriod(),
                    BigDecimal.ZERO);

            responseObserver.onNext(CalculateCPAResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void approveCommission(ApproveCommissionRequest request, StreamObserver<ApproveCommissionResponse> responseObserver) {
        try {
            CommissionClaim claim = commissionService.approveClaim(request.getId());
            responseObserver.onNext(ApproveCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void rejectCommission(RejectCommissionRequest request, StreamObserver<RejectCommissionResponse> responseObserver) {
        try {
            CommissionClaim claim = commissionService.rejectClaim(request.getId());
            responseObserver.onNext(RejectCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void payCommission(PayCommissionRequest request, StreamObserver<PayCommissionResponse> responseObserver) {
        try {
            CommissionClaim claim = commissionService.approveClaim(request.getId());
            claim.setStatus("PAID");
            responseObserver.onNext(PayCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(claim)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getPendingCommissions(GetPendingCommissionsRequest request, StreamObserver<GetPendingCommissionsResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByStatus("PENDING");
            GetPendingCommissionsResponse.Builder response = GetPendingCommissionsResponse.newBuilder();
            claims.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAllCommissions(GetAllCommissionsRequest request, StreamObserver<GetAllCommissionsResponse> responseObserver) {
        try {
            List<CommissionClaim> claims = commissionService.getClaimsByStatus("PAID");
            GetAllCommissionsResponse.Builder response = GetAllCommissionsResponse.newBuilder();
            claims.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void deleteCommission(DeleteCommissionRequest request, StreamObserver<DeleteCommissionResponse> responseObserver) {
        try {
            responseObserver.onNext(DeleteCommissionResponse.newBuilder().setSuccess(true).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private CommissionServiceProto.Commission toProtoCommission(CommissionClaim c) {
        CommissionServiceProto.Commission.Builder builder = CommissionServiceProto.Commission.newBuilder()
                .setId(c.getId())
                .setAffiliateId(c.getAffiliateId() != null ? c.getAffiliateId() : 0L)
                .setMerchantId(c.getAgentId() != null ? c.getAgentId() : 0L)
                .setStatus(c.getStatus() != null ? c.getStatus() : "")
                .setPeriod(c.getPeriod() != null ? c.getPeriod() : "");

        if (c.getGrossRevenue() != null) builder.setNetRevenue(c.getGrossRevenue().doubleValue());
        if (c.getCommissionAmount() != null) builder.setCommissionAmount(c.getCommissionAmount().doubleValue());
        if (c.getCreatedAt() != null) builder.setCreatedAt(c.getCreatedAt().toEpochSecond(ZoneOffset.UTC));
        if (c.getProcessedAt() != null) builder.setPaidAt(c.getProcessedAt().toEpochSecond(ZoneOffset.UTC));

        return builder.build();
    }
}
