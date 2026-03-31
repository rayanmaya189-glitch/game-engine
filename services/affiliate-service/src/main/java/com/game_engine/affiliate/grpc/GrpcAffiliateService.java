package com.game_engine.affiliate.grpc;

import com.game_engine.affiliate.entity.Affiliate;
import com.game_engine.affiliate.entity.Referral;
import com.game_engine.affiliate.service.AffiliateService;
import com.game_engine.affiliate.v1.*;
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
public class GrpcAffiliateService extends AffiliateServiceGrpc.AffiliateServiceImplBase {

    @Autowired
    private AffiliateService affiliateService;

    @Override
    public void registerAffiliate(RegisterAffiliateRequest request, StreamObserver<RegisterAffiliateResponse> responseObserver) {
        try {
            Affiliate affiliate = affiliateService.registerAffiliate(request.getName(), request.getEmail());
            responseObserver.onNext(RegisterAffiliateResponse.newBuilder()
                    .setAffiliate(toProtoAffiliate(affiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error registering affiliate", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAffiliateByCode(GetAffiliateByCodeRequest request, StreamObserver<GetAffiliateByCodeResponse> responseObserver) {
        try {
            Optional<Affiliate> affiliate = affiliateService.getAffiliateByCode(request.getAffiliateCode());
            GetAffiliateByCodeResponse.Builder response = GetAffiliateByCodeResponse.newBuilder()
                    .setFound(affiliate.isPresent());
            affiliate.ifPresent(a -> response.setAffiliate(toProtoAffiliate(a)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting affiliate by code", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAffiliatesByMerchant(GetAffiliatesByMerchantRequest request, StreamObserver<GetAffiliatesByMerchantResponse> responseObserver) {
        try {
            List<Affiliate> affiliates = affiliateService.getAllAffiliates();
            GetAffiliatesByMerchantResponse.Builder response = GetAffiliatesByMerchantResponse.newBuilder();
            affiliates.forEach(a -> response.addAffiliates(toProtoAffiliate(a)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting affiliates by merchant", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getActiveAffiliates(GetActiveAffiliatesRequest request, StreamObserver<GetActiveAffiliatesResponse> responseObserver) {
        try {
            List<Affiliate> affiliates = affiliateService.getAffiliatesByStatus("ACTIVE");
            GetActiveAffiliatesResponse.Builder response = GetActiveAffiliatesResponse.newBuilder();
            affiliates.forEach(a -> response.addAffiliates(toProtoAffiliate(a)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting active affiliates", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void updateAffiliateTier(UpdateAffiliateTierRequest request, StreamObserver<UpdateAffiliateTierResponse> responseObserver) {
        try {
            Affiliate affiliate = affiliateService.updateAffiliateTier(request.getAffiliateId(), request.getTier());
            responseObserver.onNext(UpdateAffiliateTierResponse.newBuilder()
                    .setAffiliate(toProtoAffiliate(affiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error updating affiliate tier", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void updateAffiliateStatus(UpdateAffiliateStatusRequest request, StreamObserver<UpdateAffiliateStatusResponse> responseObserver) {
        try {
            Affiliate affiliate = affiliateService.updateAffiliateStatus(request.getAffiliateId(), request.getStatus());
            responseObserver.onNext(UpdateAffiliateStatusResponse.newBuilder()
                    .setAffiliate(toProtoAffiliate(affiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error updating affiliate status", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackClick(TrackClickRequest request, StreamObserver<TrackClickResponse> responseObserver) {
        try {
            Referral referral = affiliateService.trackReferral(
                    request.getAffiliateCode(), 0L, "LINK");
            responseObserver.onNext(TrackClickResponse.newBuilder()
                    .setReferral(toProtoReferral(referral))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking click", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackRegistration(TrackRegistrationRequest request, StreamObserver<TrackRegistrationResponse> responseObserver) {
        try {
            List<Referral> referrals = affiliateService.getReferralsByAffiliateCode("");
            Optional<Referral> match = referrals.stream().findFirst();
            TrackRegistrationResponse.Builder response = TrackRegistrationResponse.newBuilder();
            match.ifPresent(r -> response.setReferral(toProtoReferral(r)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking registration", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackFirstDeposit(TrackFirstDepositRequest request, StreamObserver<TrackFirstDepositResponse> responseObserver) {
        try {
            List<Referral> referrals = affiliateService.getReferralsByAffiliateCode(request.getReferralCode());
            Optional<Referral> match = referrals.stream().findFirst();
            TrackFirstDepositResponse.Builder response = TrackFirstDepositResponse.newBuilder();
            match.ifPresent(r -> {
                r.setFirstDepositAt(java.time.LocalDateTime.now());
                r.setTotalDeposits(BigDecimal.valueOf(request.getDepositAmount()));
                r.setStatus("ACTIVE");
                response.setReferral(toProtoReferral(r));
            });
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking first deposit", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getReferrals(GetReferralsRequest request, StreamObserver<GetReferralsResponse> responseObserver) {
        try {
            Optional<Affiliate> affiliate = affiliateService.getAffiliateById(request.getAffiliateId());
            List<Referral> referrals = affiliate
                    .map(a -> affiliateService.getReferralsByAffiliateCode(a.getAffiliateCode()))
                    .orElse(List.of());
            GetReferralsResponse.Builder response = GetReferralsResponse.newBuilder();
            referrals.forEach(r -> response.addReferrals(toProtoReferral(r)));
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting referrals", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCampaignReferrals(GetCampaignReferralsRequest request, StreamObserver<GetCampaignReferralsResponse> responseObserver) {
        try {
            GetCampaignReferralsResponse.Builder response = GetCampaignReferralsResponse.newBuilder();
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting campaign referrals", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateCommission(CalculateCommissionRequest request, StreamObserver<CalculateCommissionResponse> responseObserver) {
        try {
            BigDecimal revenue = BigDecimal.valueOf(request.getRevenue());
            BigDecimal commission = affiliateService.calculateCommissions(request.getAffiliateId());
            responseObserver.onNext(CalculateCommissionResponse.newBuilder()
                    .setRevenue(revenue.doubleValue())
                    .setCommission(commission.doubleValue())
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error calculating commission", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void addSubAffiliate(AddSubAffiliateRequest request, StreamObserver<AddSubAffiliateResponse> responseObserver) {
        try {
            Affiliate subAffiliate = affiliateService.registerAffiliate(request.getName(), request.getEmail());
            subAffiliate.setCommissionRate(new BigDecimal("0.1000"));
            subAffiliate.setTier("BRONZE");
            responseObserver.onNext(AddSubAffiliateResponse.newBuilder()
                    .setSubAffiliate(toProtoAffiliate(subAffiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error adding sub-affiliate", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSubAffiliates(GetSubAffiliatesRequest request, StreamObserver<GetSubAffiliatesResponse> responseObserver) {
        try {
            GetSubAffiliatesResponse.Builder response = GetSubAffiliatesResponse.newBuilder();
            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting sub-affiliates", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAffiliateStats(GetAffiliateStatsRequest request, StreamObserver<GetAffiliateStatsResponse> responseObserver) {
        try {
            AffiliateService.AffiliateStats stats = affiliateService.getAffiliateStats(request.getAffiliateId());
            GetAffiliateStatsResponse response = GetAffiliateStatsResponse.newBuilder()
                    .setTotalClicks(stats.getTotalReferrals())
                    .setTotalRegistrations(stats.getTotalReferrals())
                    .setTotalDepositors(stats.getActiveReferrals())
                    .setTotalRevenue(stats.getTotalDeposits().doubleValue())
                    .setTotalCommission(stats.getTotalCommission().doubleValue())
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting affiliate stats", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void redirectToRegistration(RedirectToRegistrationRequest request, StreamObserver<RedirectToRegistrationResponse> responseObserver) {
        try {
            RedirectToRegistrationResponse response = RedirectToRegistrationResponse.newBuilder()
                    .setRedirectUrl("/register?ref=" + request.getReferralCode())
                    .build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error redirecting to registration", e);
            responseObserver.onError(io.grpc.Status.INTERNAL.withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private AffiliateServiceProto.Affiliate toProtoAffiliate(Affiliate affiliate) {
        AffiliateServiceProto.Affiliate.Builder builder = AffiliateServiceProto.Affiliate.newBuilder()
                .setId(affiliate.getId())
                .setAffiliateCode(affiliate.getAffiliateCode() != null ? affiliate.getAffiliateCode() : "")
                .setName(affiliate.getName() != null ? affiliate.getName() : "")
                .setEmail(affiliate.getEmail() != null ? affiliate.getEmail() : "")
                .setStatus(affiliate.getStatus() != null ? affiliate.getStatus() : "")
                .setAffiliateTier(affiliate.getTier() != null ? affiliate.getTier() : "");

        if (affiliate.getCommissionRate() != null)
            builder.setRevenueSharePercentage(affiliate.getCommissionRate().multiply(new BigDecimal("100")).doubleValue());
        if (affiliate.getTotalReferrals() != null) builder.setTotalRegistrations(affiliate.getTotalReferrals());
        if (affiliate.getTotalRevenue() != null) builder.setTotalRevenue(affiliate.getTotalRevenue().doubleValue());
        if (affiliate.getCreatedAt() != null) builder.setCreatedAt(affiliate.getCreatedAt().toEpochSecond(ZoneOffset.UTC));

        return builder.build();
    }

    private AffiliateServiceProto.Referral toProtoReferral(Referral referral) {
        AffiliateServiceProto.Referral.Builder builder = AffiliateServiceProto.Referral.newBuilder()
                .setId(referral.getId())
                .setReferralCode(referral.getAffiliateCode() != null ? referral.getAffiliateCode() : "")
                .setSource(referral.getSource() != null ? referral.getSource() : "")
                .setStatus(referral.getStatus() != null ? referral.getStatus() : "");

        if (referral.getPlayerId() != null) builder.setAffiliateId(referral.getPlayerId());
        if (referral.getFirstDepositAt() != null) builder.setFirstDepositAt(referral.getFirstDepositAt().toEpochSecond(ZoneOffset.UTC));
        if (referral.getTotalDeposits() != null) builder.setFirstDepositAmount(referral.getTotalDeposits().doubleValue());

        return builder.build();
    }
}
