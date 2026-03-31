package com.game_engine.commission.grpc;

import com.game_engine.commission.model.*;
import com.game_engine.commission.service.CommissionService;
import com.game_engine.commission.service.CommissionConfigService;
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

@GrpcService
@Slf4j
public class GrpcCommissionService extends CommissionServiceGrpc.CommissionServiceImplBase {

    @Autowired
    private CommissionService commissionService;

    @Override
    public void createCommission(CreateCommissionRequest request, StreamObserver<CreateCommissionResponse> responseObserver) {
        try {
            Commission proto = request.getCommission();
            Commission commission = new Commission();
            commission.setAffiliateId(proto.getAffiliateId());
            commission.setMerchantId(proto.getMerchantId());
            commission.setCommissionType(proto.getCommissionType());
            commission.setNetRevenue(BigDecimal.valueOf(proto.getNetRevenue()));
            commission.setCommissionAmount(BigDecimal.valueOf(proto.getCommissionAmount()));
            commission.setPeriod(proto.getPeriod());

            Commission created = commissionService.createCommission(commission);
            responseObserver.onNext(CreateCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(created)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error creating commission", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionById(GetCommissionByIdRequest request, StreamObserver<GetCommissionByIdResponse> responseObserver) {
        try {
            Optional<Commission> commission = commissionService.getCommissionById(request.getId());
            GetCommissionByIdResponse.Builder response = GetCommissionByIdResponse.newBuilder().setFound(commission.isPresent());
            commission.ifPresent(c -> response.setCommission(toProtoCommission(c)));
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
            List<Commission> commissions = commissionService.getCommissionsByAffiliate(request.getAffiliateId());
            GetCommissionsByAffiliateResponse.Builder response = GetCommissionsByAffiliateResponse.newBuilder();
            commissions.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionsByMerchant(GetCommissionsByMerchantRequest request, StreamObserver<GetCommissionsByMerchantResponse> responseObserver) {
        try {
            List<Commission> commissions = commissionService.getCommissionsByMerchant(request.getMerchantId());
            GetCommissionsByMerchantResponse.Builder response = GetCommissionsByMerchantResponse.newBuilder();
            commissions.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionsByPeriod(GetCommissionsByPeriodRequest request, StreamObserver<GetCommissionsByPeriodResponse> responseObserver) {
        try {
            List<Commission> commissions = commissionService.getCommissionsByPeriod(request.getPeriod());
            GetCommissionsByPeriodResponse.Builder response = GetCommissionsByPeriodResponse.newBuilder();
            commissions.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCommissionByAffiliateAndPeriod(GetCommissionByAffiliateAndPeriodRequest request, StreamObserver<GetCommissionByAffiliateAndPeriodResponse> responseObserver) {
        try {
            Optional<Commission> commission = commissionService.getCommissionByAffiliateAndPeriod(request.getAffiliateId(), request.getPeriod());
            GetCommissionByAffiliateAndPeriodResponse.Builder response = GetCommissionByAffiliateAndPeriodResponse.newBuilder().setFound(commission.isPresent());
            commission.ifPresent(c -> response.setCommission(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalPaidCommission(GetTotalPaidCommissionRequest request, StreamObserver<GetTotalPaidCommissionResponse> responseObserver) {
        try {
            BigDecimal total = commissionService.getTotalPaidCommission(request.getAffiliateId());
            responseObserver.onNext(GetTotalPaidCommissionResponse.newBuilder().setTotalPaid(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalPendingCommission(GetTotalPendingCommissionRequest request, StreamObserver<GetTotalPendingCommissionResponse> responseObserver) {
        try {
            BigDecimal total = commissionService.getTotalPendingCommission(request.getAffiliateId());
            responseObserver.onNext(GetTotalPendingCommissionResponse.newBuilder().setTotalPending(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getTotalRevenueByMerchant(GetTotalRevenueByMerchantRequest request, StreamObserver<GetTotalRevenueByMerchantResponse> responseObserver) {
        try {
            BigDecimal total = commissionService.getTotalRevenueByMerchant(request.getMerchantId());
            responseObserver.onNext(GetTotalRevenueByMerchantResponse.newBuilder().setTotalRevenue(total.doubleValue()).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateRevenueShare(CalculateRevenueShareRequest request, StreamObserver<CalculateRevenueShareResponse> responseObserver) {
        try {
            String period = request.getPeriod().isEmpty() ? commissionService.generateCurrentPeriod() : request.getPeriod();
            Commission commission = commissionService.calculateRevenueShare(
                    request.getAffiliateId(), request.getMerchantId(),
                    BigDecimal.valueOf(request.getNetRevenue()),
                    BigDecimal.valueOf(request.getCommissionRate()), period);
            responseObserver.onNext(CalculateRevenueShareResponse.newBuilder()
                    .setCommission(toProtoCommission(commission)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateCPA(CalculateCPARequest request, StreamObserver<CalculateCPAResponse> responseObserver) {
        try {
            String period = request.getPeriod().isEmpty() ? commissionService.generateCurrentPeriod() : request.getPeriod();
            Commission commission = commissionService.calculateCPA(
                    request.getAffiliateId(), request.getMerchantId(),
                    request.getNewPlayers(), BigDecimal.valueOf(request.getCpaRate()), period);
            responseObserver.onNext(CalculateCPAResponse.newBuilder()
                    .setCommission(toProtoCommission(commission)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void approveCommission(ApproveCommissionRequest request, StreamObserver<ApproveCommissionResponse> responseObserver) {
        try {
            Commission commission = commissionService.approveCommission(request.getId());
            responseObserver.onNext(ApproveCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(commission)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void rejectCommission(RejectCommissionRequest request, StreamObserver<RejectCommissionResponse> responseObserver) {
        try {
            Commission commission = commissionService.rejectCommission(request.getId(), request.getReason());
            responseObserver.onNext(RejectCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(commission)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void payCommission(PayCommissionRequest request, StreamObserver<PayCommissionResponse> responseObserver) {
        try {
            Commission commission = commissionService.payCommission(request.getId());
            responseObserver.onNext(PayCommissionResponse.newBuilder()
                    .setCommission(toProtoCommission(commission)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getPendingCommissions(GetPendingCommissionsRequest request, StreamObserver<GetPendingCommissionsResponse> responseObserver) {
        try {
            List<Commission> commissions = commissionService.getPendingCommissions();
            GetPendingCommissionsResponse.Builder response = GetPendingCommissionsResponse.newBuilder();
            commissions.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAllCommissions(GetAllCommissionsRequest request, StreamObserver<GetAllCommissionsResponse> responseObserver) {
        try {
            List<Commission> commissions = commissionService.getAllCommissions();
            GetAllCommissionsResponse.Builder response = GetAllCommissionsResponse.newBuilder();
            commissions.forEach(c -> response.addCommissions(toProtoCommission(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void deleteCommission(DeleteCommissionRequest request, StreamObserver<DeleteCommissionResponse> responseObserver) {
        try {
            commissionService.deleteCommission(request.getId());
            responseObserver.onNext(DeleteCommissionResponse.newBuilder().setSuccess(true).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private CommissionServiceProto.Commission toProtoCommission(Commission c) {
        CommissionServiceProto.Commission.Builder builder = CommissionServiceProto.Commission.newBuilder()
                .setId(c.getId())
                .setAffiliateId(c.getAffiliateId())
                .setMerchantId(c.getMerchantId())
                .setStatus(c.getStatus() != null ? c.getStatus() : "")
                .setPeriod(c.getPeriod() != null ? c.getPeriod() : "");

        if (c.getCommissionType() != null) builder.setCommissionType(c.getCommissionType());
        if (c.getNetRevenue() != null) builder.setNetRevenue(c.getNetRevenue().doubleValue());
        if (c.getCommissionPercentage() != null) builder.setCommissionPercentage(c.getCommissionPercentage().doubleValue());
        if (c.getCommissionAmount() != null) builder.setCommissionAmount(c.getCommissionAmount().doubleValue());
        if (c.getCpaAmount() != null) builder.setCpaAmount(c.getCpaAmount().doubleValue());
        if (c.getCpaCount() != null) builder.setCpaCount(c.getCpaCount());
        if (c.getTotalDeposits() != null) builder.setTotalDeposits(c.getTotalDeposits().doubleValue());
        if (c.getTotalPlayers() != null) builder.setTotalPlayers(c.getTotalPlayers());
        if (c.getCalculatedAt() != null) builder.setCalculatedAt(c.getCalculatedAt().toEpochSecond(ZoneOffset.UTC));
        if (c.getApprovedAt() != null) builder.setApprovedAt(c.getApprovedAt().toEpochSecond(ZoneOffset.UTC));
        if (c.getPaidAt() != null) builder.setPaidAt(c.getPaidAt().toEpochSecond(ZoneOffset.UTC));
        if (c.getRejectionReason() != null) builder.setRejectionReason(c.getRejectionReason());
        if (c.getCreatedAt() != null) builder.setCreatedAt(c.getCreatedAt().toEpochSecond(ZoneOffset.UTC));

        return builder.build();
    }
}
