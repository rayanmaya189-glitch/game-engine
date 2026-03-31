package com.game_engine.affiliate.grpc;

import com.game_engine.affiliate.model.Affiliate;
import com.game_engine.affiliate.model.Referral;
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
            Affiliate affiliate = affiliateService.registerAffiliate(
                    request.getName(), request.getEmail(), request.getPhone(), request.getMerchantId());

            responseObserver.onNext(RegisterAffiliateResponse.newBuilder()
                    .setAffiliate(toProtoAffiliate(affiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error registering affiliate", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
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
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAffiliatesByMerchant(GetAffiliatesByMerchantRequest request, StreamObserver<GetAffiliatesByMerchantResponse> responseObserver) {
        try {
            List<Affiliate> affiliates = affiliateService.getAffiliatesByMerchant(request.getMerchantId());
            GetAffiliatesByMerchantResponse.Builder response = GetAffiliatesByMerchantResponse.newBuilder();
            affiliates.forEach(a -> response.addAffiliates(toProtoAffiliate(a)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting affiliates by merchant", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getActiveAffiliates(GetActiveAffiliatesRequest request, StreamObserver<GetActiveAffiliatesResponse> responseObserver) {
        try {
            List<Affiliate> affiliates = affiliateService.getActiveAffiliates(request.getMerchantId());
            GetActiveAffiliatesResponse.Builder response = GetActiveAffiliatesResponse.newBuilder();
            affiliates.forEach(a -> response.addAffiliates(toProtoAffiliate(a)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting active affiliates", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
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
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
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
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackClick(TrackClickRequest request, StreamObserver<TrackClickResponse> responseObserver) {
        try {
            Referral referral = affiliateService.trackClick(
                    request.getAffiliateCode(), request.getIpAddress(),
                    request.getUserAgent(), request.getCampaignId());

            responseObserver.onNext(TrackClickResponse.newBuilder()
                    .setReferral(toProtoReferral(referral))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking click", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackRegistration(TrackRegistrationRequest request, StreamObserver<TrackRegistrationResponse> responseObserver) {
        try {
            Referral referral = affiliateService.trackRegistration(
                    request.getReferralCode(), request.getIpAddress());

            responseObserver.onNext(TrackRegistrationResponse.newBuilder()
                    .setReferral(toProtoReferral(referral))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking registration", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void trackFirstDeposit(TrackFirstDepositRequest request, StreamObserver<TrackFirstDepositResponse> responseObserver) {
        try {
            Referral referral = affiliateService.trackFirstDeposit(
                    request.getReferralCode(), BigDecimal.valueOf(request.getDepositAmount()));

            responseObserver.onNext(TrackFirstDepositResponse.newBuilder()
                    .setReferral(toProtoReferral(referral))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error tracking first deposit", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getReferrals(GetReferralsRequest request, StreamObserver<GetReferralsResponse> responseObserver) {
        try {
            List<Referral> referrals = affiliateService.getReferralsByAffiliate(request.getAffiliateId());
            GetReferralsResponse.Builder response = GetReferralsResponse.newBuilder();
            referrals.forEach(r -> response.addReferrals(toProtoReferral(r)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting referrals", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getCampaignReferrals(GetCampaignReferralsRequest request, StreamObserver<GetCampaignReferralsResponse> responseObserver) {
        try {
            List<Referral> referrals = affiliateService.getReferralsByCampaign(request.getCampaignId());
            GetCampaignReferralsResponse.Builder response = GetCampaignReferralsResponse.newBuilder();
            referrals.forEach(r -> response.addReferrals(toProtoReferral(r)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting campaign referrals", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void calculateCommission(CalculateCommissionRequest request, StreamObserver<CalculateCommissionResponse> responseObserver) {
        try {
            BigDecimal revenue = BigDecimal.valueOf(request.getRevenue());
            BigDecimal commission = affiliateService.calculateCommission(request.getAffiliateId(), revenue);

            responseObserver.onNext(CalculateCommissionResponse.newBuilder()
                    .setRevenue(revenue.doubleValue())
                    .setCommission(commission.doubleValue())
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error calculating commission", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void addSubAffiliate(AddSubAffiliateRequest request, StreamObserver<AddSubAffiliateResponse> responseObserver) {
        try {
            Affiliate subAffiliate = affiliateService.addSubAffiliate(
                    request.getParentAffiliateId(), request.getName(),
                    request.getEmail(), request.getPhone());

            responseObserver.onNext(AddSubAffiliateResponse.newBuilder()
                    .setSubAffiliate(toProtoAffiliate(subAffiliate))
                    .build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error adding sub-affiliate", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getSubAffiliates(GetSubAffiliatesRequest request, StreamObserver<GetSubAffiliatesResponse> responseObserver) {
        try {
            List<Affiliate> subAffiliates = affiliateService.getSubAffiliates(request.getParentAffiliateId());
            GetSubAffiliatesResponse.Builder response = GetSubAffiliatesResponse.newBuilder();
            subAffiliates.forEach(a -> response.addSubAffiliates(toProtoAffiliate(a)));

            responseObserver.onNext(response.build());
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting sub-affiliates", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    @Override
    public void getAffiliateStats(GetAffiliateStatsRequest request, StreamObserver<GetAffiliateStatsResponse> responseObserver) {
        try {
            Long affiliateId = request.getAffiliateId();
            GetAffiliateStatsResponse response = GetAffiliateStatsResponse.newBuilder()
                    .setTotalClicks(affiliateService.getTotalClicks(affiliateId))
                    .setTotalRegistrations(affiliateService.getTotalRegistrations(affiliateId))
                    .setTotalDepositors(affiliateService.getTotalDepositors(affiliateId))
                    .setTotalRevenue(affiliateService.getTotalRevenue(affiliateId).doubleValue())
                    .setTotalCommission(affiliateService.getTotalCommission(affiliateId).doubleValue())
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        } catch (Exception e) {
            log.error("Error getting affiliate stats", e);
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
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
            responseObserver.onError(io.grpc.Status.INTERNAL
                    .withDescription(e.getMessage()).asRuntimeException());
        }
    }

    private AffiliateServiceProto.Affiliate toProtoAffiliate(Affiliate affiliate) {
        AffiliateServiceProto.Affiliate.Builder builder = AffiliateServiceProto.Affiliate.newBuilder()
                .setId(affiliate.getId())
                .setAffiliateCode(affiliate.getAffiliateCode() != null ? affiliate.getAffiliateCode() : "")
                .setName(affiliate.getName() != null ? affiliate.getName() : "")
                .setEmail(affiliate.getEmail() != null ? affiliate.getEmail() : "")
                .setPhone(affiliate.getPhone() != null ? affiliate.getPhone() : "")
                .setStatus(affiliate.getStatus() != null ? affiliate.getStatus() : "")
                .setAffiliateTier(affiliate.getAffiliateTier() != null ? affiliate.getAffiliateTier() : "")
                .setMerchantId(affiliate.getMerchantId() != null ? affiliate.getMerchantId() : 0L);

        if (affiliate.getRevenueSharePercentage() != null)
            builder.setRevenueSharePercentage(affiliate.getRevenueSharePercentage().doubleValue());
        if (affiliate.getTotalClicks() != null) builder.setTotalClicks(affiliate.getTotalClicks());
        if (affiliate.getTotalRegistrations() != null) builder.setTotalRegistrations(affiliate.getTotalRegistrations());
        if (affiliate.getTotalDepositors() != null) builder.setTotalDepositors(affiliate.getTotalDepositors());
        if (affiliate.getTotalRevenue() != null) builder.setTotalRevenue(affiliate.getTotalRevenue().doubleValue());
        if (affiliate.getTotalCommission() != null) builder.setTotalCommission(affiliate.getTotalCommission().doubleValue());
        if (affiliate.getCreatedAt() != null) builder.setCreatedAt(affiliate.getCreatedAt().toEpochSecond(ZoneOffset.UTC));
        if (affiliate.getUpdatedAt() != null) builder.setUpdatedAt(affiliate.getUpdatedAt().toEpochSecond(ZoneOffset.UTC));

        return builder.build();
    }

    private AffiliateServiceProto.Referral toProtoReferral(Referral referral) {
        AffiliateServiceProto.Referral.Builder builder = AffiliateServiceProto.Referral.newBuilder()
                .setId(referral.getId())
                .setAffiliateId(referral.getAffiliate() != null ? referral.getAffiliate().getId() : 0L)
                .setReferralCode(referral.getReferralCode() != null ? referral.getReferralCode() : "")
                .setSource(referral.getSource() != null ? referral.getSource() : "")
                .setIpAddress(referral.getIpAddress() != null ? referral.getIpAddress() : "")
                .setUserAgent(referral.getUserAgent() != null ? referral.getUserAgent() : "")
                .setCampaignId(referral.getCampaignId() != null ? referral.getCampaignId() : "")
                .setStatus(referral.getStatus() != null ? referral.getStatus() : "");

        if (referral.getClickedAt() != null) builder.setClickedAt(referral.getClickedAt().toEpochSecond(ZoneOffset.UTC));
        if (referral.getRegisteredAt() != null) builder.setRegisteredAt(referral.getRegisteredAt().toEpochSecond(ZoneOffset.UTC));
        if (referral.getFirstDepositAt() != null) builder.setFirstDepositAt(referral.getFirstDepositAt().toEpochSecond(ZoneOffset.UTC));
        if (referral.getFirstDepositAmount() != null) builder.setFirstDepositAmount(referral.getFirstDepositAmount().doubleValue());

        return builder.build();
    }
}
