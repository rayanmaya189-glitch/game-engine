package com.game_engine.commission.grpc;

import com.game_engine.commission.model.*;
import com.game_engine.commission.service.CommissionConfigService;
import com.game_engine.commission.v1.*;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.List;
import java.util.Optional;

@GrpcService
@Slf4j
public class GrpcCommissionConfigService extends CommissionConfigServiceGrpc.CommissionConfigServiceImplBase {

    @Autowired
    private CommissionConfigService commissionConfigService;

    @Override
    public void createCommissionConfig(CreateCommissionConfigRequest request, StreamObserver<CreateCommissionConfigResponse> responseObserver) {
        try {
            CommissionConfig config = fromProtoConfig(request.getConfig());
            CommissionConfig created = commissionConfigService.createCommissionConfig(config);
            responseObserver.onNext(CreateCommissionConfigResponse.newBuilder()
                    .setConfig(toProtoConfig(created)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getConfigById(GetConfigByIdRequest request, StreamObserver<GetConfigByIdResponse> responseObserver) {
        try {
            Optional<CommissionConfig> config = commissionConfigService.getConfigById(request.getId());
            GetConfigByIdResponse.Builder response = GetConfigByIdResponse.newBuilder().setFound(config.isPresent());
            config.ifPresent(c -> response.setConfig(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getConfigsByAffiliate(GetConfigsByAffiliateRequest request, StreamObserver<GetConfigsByAffiliateResponse> responseObserver) {
        try {
            List<CommissionConfig> configs = commissionConfigService.getConfigsByAffiliate(request.getAffiliateId());
            GetConfigsByAffiliateResponse.Builder response = GetConfigsByAffiliateResponse.newBuilder();
            configs.forEach(c -> response.addConfigs(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getConfigsByMerchant(GetConfigsByMerchantRequest request, StreamObserver<GetConfigsByMerchantResponse> responseObserver) {
        try {
            List<CommissionConfig> configs = commissionConfigService.getConfigsByMerchant(request.getMerchantId());
            GetConfigsByMerchantResponse.Builder response = GetConfigsByMerchantResponse.newBuilder();
            configs.forEach(c -> response.addConfigs(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getActiveConfigsByAffiliate(GetActiveConfigsByAffiliateRequest request, StreamObserver<GetActiveConfigsByAffiliateResponse> responseObserver) {
        try {
            List<CommissionConfig> configs = commissionConfigService.getActiveConfigsByAffiliate(request.getAffiliateId());
            GetActiveConfigsByAffiliateResponse.Builder response = GetActiveConfigsByAffiliateResponse.newBuilder();
            configs.forEach(c -> response.addConfigs(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getActiveConfigsByAffiliateAndMerchant(GetActiveConfigsByAffiliateAndMerchantRequest request, StreamObserver<GetActiveConfigsByAffiliateAndMerchantResponse> responseObserver) {
        try {
            List<CommissionConfig> configs = commissionConfigService.getActiveConfigsByAffiliateAndMerchant(request.getAffiliateId(), request.getMerchantId());
            GetActiveConfigsByAffiliateAndMerchantResponse.Builder response = GetActiveConfigsByAffiliateAndMerchantResponse.newBuilder();
            configs.forEach(c -> response.addConfigs(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getConfigByAffiliateAndMerchantAndType(GetConfigByAffiliateAndMerchantAndTypeRequest request, StreamObserver<GetConfigByAffiliateAndMerchantAndTypeResponse> responseObserver) {
        try {
            Optional<CommissionConfig> config = commissionConfigService.getConfigByAffiliateAndMerchantAndType(
                    request.getAffiliateId(), request.getMerchantId(), request.getType());
            GetConfigByAffiliateAndMerchantAndTypeResponse.Builder response = GetConfigByAffiliateAndMerchantAndTypeResponse.newBuilder()
                    .setFound(config.isPresent());
            config.ifPresent(c -> response.setConfig(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void updateCommissionConfig(UpdateCommissionConfigRequest request, StreamObserver<UpdateCommissionConfigResponse> responseObserver) {
        try {
            CommissionConfig config = fromProtoConfig(request.getConfig());
            CommissionConfig updated = commissionConfigService.updateCommissionConfig(request.getId(), config);
            responseObserver.onNext(UpdateCommissionConfigResponse.newBuilder()
                    .setConfig(toProtoConfig(updated)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void activateConfig(ActivateConfigRequest request, StreamObserver<ActivateConfigResponse> responseObserver) {
        try {
            CommissionConfig config = commissionConfigService.activateConfig(request.getId());
            responseObserver.onNext(ActivateConfigResponse.newBuilder()
                    .setConfig(toProtoConfig(config)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void deactivateConfig(DeactivateConfigRequest request, StreamObserver<DeactivateConfigResponse> responseObserver) {
        try {
            CommissionConfig config = commissionConfigService.deactivateConfig(request.getId());
            responseObserver.onNext(DeactivateConfigResponse.newBuilder()
                    .setConfig(toProtoConfig(config)).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void deleteConfig(DeleteConfigRequest request, StreamObserver<DeleteConfigResponse> responseObserver) {
        try {
            commissionConfigService.deleteConfig(request.getId());
            responseObserver.onNext(DeleteConfigResponse.newBuilder().setSuccess(true).build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAllConfigs(GetAllConfigsRequest request, StreamObserver<GetAllConfigsResponse> responseObserver) {
        try {
            List<CommissionConfig> configs = commissionConfigService.getAllConfigs();
            GetAllConfigsResponse.Builder response = GetAllConfigsResponse.newBuilder();
            configs.forEach(c -> response.addConfigs(toProtoConfig(c)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private CommissionServiceProto.CommissionConfig toProtoConfig(CommissionConfig c) {
        CommissionServiceProto.CommissionConfig.Builder builder = CommissionServiceProto.CommissionConfig.newBuilder()
                .setId(c.getId());

        if (c.getMerchantId() != null) builder.setMerchantId(c.getMerchantId());
        if (c.getAffiliateId() != null) builder.setAffiliateId(c.getAffiliateId());
        if (c.getCommissionType() != null) builder.setCommissionType(c.getCommissionType());
        if (c.getRevenueShareRate() != null) builder.setRevenueShareRate(c.getRevenueShareRate().doubleValue());
        if (c.getCpaRate() != null) builder.setCpaRate(c.getCpaRate().doubleValue());
        if (c.getMinPlayers() != null) builder.setMinPlayers(c.getMinPlayers());
        if (c.getTierRate() != null) builder.setTierRate(c.getTierRate().doubleValue());
        if (c.getTierThreshold() != null) builder.setTierThreshold(c.getTierThreshold());
        if (c.getIsActive() != null) builder.setIsActive(c.getIsActive());
        if (c.getEffectiveFrom() != null) builder.setEffectiveFrom(c.getEffectiveFrom().toEpochSecond(ZoneOffset.UTC));
        if (c.getEffectiveTo() != null) builder.setEffectiveTo(c.getEffectiveTo().toEpochSecond(ZoneOffset.UTC));
        if (c.getCreatedAt() != null) builder.setCreatedAt(c.getCreatedAt().toEpochSecond(ZoneOffset.UTC));
        if (c.getUpdatedAt() != null) builder.setUpdatedAt(c.getUpdatedAt().toEpochSecond(ZoneOffset.UTC));

        return builder.build();
    }

    private CommissionConfig fromProtoConfig(CommissionServiceProto.CommissionConfig proto) {
        CommissionConfig config = new CommissionConfig();
        config.setMerchantId(proto.getMerchantId());
        config.setAffiliateId(proto.getAffiliateId());
        config.setCommissionType(proto.getCommissionType());
        config.setRevenueShareRate(BigDecimal.valueOf(proto.getRevenueShareRate()));
        config.setCpaRate(BigDecimal.valueOf(proto.getCpaRate()));
        config.setMinPlayers(proto.getMinPlayers());
        config.setTierRate(BigDecimal.valueOf(proto.getTierRate()));
        config.setTierThreshold(proto.getTierThreshold());
        config.setIsActive(proto.getIsActive());
        if (proto.getEffectiveFrom() != 0) config.setEffectiveFrom(LocalDateTime.ofEpochSecond(proto.getEffectiveFrom(), 0, ZoneOffset.UTC));
        if (proto.getEffectiveTo() != 0) config.setEffectiveTo(LocalDateTime.ofEpochSecond(proto.getEffectiveTo(), 0, ZoneOffset.UTC));
        return config;
    }
}
