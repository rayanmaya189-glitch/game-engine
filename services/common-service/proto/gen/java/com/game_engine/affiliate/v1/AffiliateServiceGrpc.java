package com.game_engine.affiliate.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Affiliate Service - manages affiliate registration, tracking, and commissions
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class AffiliateServiceGrpc {

  private AffiliateServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.affiliate.v1.AffiliateService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RegisterAffiliateRequest,
      com.game_engine.affiliate.v1.RegisterAffiliateResponse> getRegisterAffiliateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegisterAffiliate",
      requestType = com.game_engine.affiliate.v1.RegisterAffiliateRequest.class,
      responseType = com.game_engine.affiliate.v1.RegisterAffiliateResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RegisterAffiliateRequest,
      com.game_engine.affiliate.v1.RegisterAffiliateResponse> getRegisterAffiliateMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RegisterAffiliateRequest, com.game_engine.affiliate.v1.RegisterAffiliateResponse> getRegisterAffiliateMethod;
    if ((getRegisterAffiliateMethod = AffiliateServiceGrpc.getRegisterAffiliateMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getRegisterAffiliateMethod = AffiliateServiceGrpc.getRegisterAffiliateMethod) == null) {
          AffiliateServiceGrpc.getRegisterAffiliateMethod = getRegisterAffiliateMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.RegisterAffiliateRequest, com.game_engine.affiliate.v1.RegisterAffiliateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegisterAffiliate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.RegisterAffiliateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.RegisterAffiliateResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("RegisterAffiliate"))
              .build();
        }
      }
    }
    return getRegisterAffiliateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateByCodeRequest,
      com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> getGetAffiliateByCodeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAffiliateByCode",
      requestType = com.game_engine.affiliate.v1.GetAffiliateByCodeRequest.class,
      responseType = com.game_engine.affiliate.v1.GetAffiliateByCodeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateByCodeRequest,
      com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> getGetAffiliateByCodeMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateByCodeRequest, com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> getGetAffiliateByCodeMethod;
    if ((getGetAffiliateByCodeMethod = AffiliateServiceGrpc.getGetAffiliateByCodeMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetAffiliateByCodeMethod = AffiliateServiceGrpc.getGetAffiliateByCodeMethod) == null) {
          AffiliateServiceGrpc.getGetAffiliateByCodeMethod = getGetAffiliateByCodeMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetAffiliateByCodeRequest, com.game_engine.affiliate.v1.GetAffiliateByCodeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAffiliateByCode"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateByCodeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateByCodeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetAffiliateByCode"))
              .build();
        }
      }
    }
    return getGetAffiliateByCodeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest,
      com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> getGetAffiliatesByMerchantMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAffiliatesByMerchant",
      requestType = com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest.class,
      responseType = com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest,
      com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> getGetAffiliatesByMerchantMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest, com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> getGetAffiliatesByMerchantMethod;
    if ((getGetAffiliatesByMerchantMethod = AffiliateServiceGrpc.getGetAffiliatesByMerchantMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetAffiliatesByMerchantMethod = AffiliateServiceGrpc.getGetAffiliatesByMerchantMethod) == null) {
          AffiliateServiceGrpc.getGetAffiliatesByMerchantMethod = getGetAffiliatesByMerchantMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest, com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAffiliatesByMerchant"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetAffiliatesByMerchant"))
              .build();
        }
      }
    }
    return getGetAffiliatesByMerchantMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetActiveAffiliatesRequest,
      com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> getGetActiveAffiliatesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActiveAffiliates",
      requestType = com.game_engine.affiliate.v1.GetActiveAffiliatesRequest.class,
      responseType = com.game_engine.affiliate.v1.GetActiveAffiliatesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetActiveAffiliatesRequest,
      com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> getGetActiveAffiliatesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetActiveAffiliatesRequest, com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> getGetActiveAffiliatesMethod;
    if ((getGetActiveAffiliatesMethod = AffiliateServiceGrpc.getGetActiveAffiliatesMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetActiveAffiliatesMethod = AffiliateServiceGrpc.getGetActiveAffiliatesMethod) == null) {
          AffiliateServiceGrpc.getGetActiveAffiliatesMethod = getGetActiveAffiliatesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetActiveAffiliatesRequest, com.game_engine.affiliate.v1.GetActiveAffiliatesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveAffiliates"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetActiveAffiliatesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetActiveAffiliatesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetActiveAffiliates"))
              .build();
        }
      }
    }
    return getGetActiveAffiliatesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateTierRequest,
      com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> getUpdateAffiliateTierMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateAffiliateTier",
      requestType = com.game_engine.affiliate.v1.UpdateAffiliateTierRequest.class,
      responseType = com.game_engine.affiliate.v1.UpdateAffiliateTierResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateTierRequest,
      com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> getUpdateAffiliateTierMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateTierRequest, com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> getUpdateAffiliateTierMethod;
    if ((getUpdateAffiliateTierMethod = AffiliateServiceGrpc.getUpdateAffiliateTierMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getUpdateAffiliateTierMethod = AffiliateServiceGrpc.getUpdateAffiliateTierMethod) == null) {
          AffiliateServiceGrpc.getUpdateAffiliateTierMethod = getUpdateAffiliateTierMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.UpdateAffiliateTierRequest, com.game_engine.affiliate.v1.UpdateAffiliateTierResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateAffiliateTier"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.UpdateAffiliateTierRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.UpdateAffiliateTierResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("UpdateAffiliateTier"))
              .build();
        }
      }
    }
    return getUpdateAffiliateTierMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest,
      com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> getUpdateAffiliateStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateAffiliateStatus",
      requestType = com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest.class,
      responseType = com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest,
      com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> getUpdateAffiliateStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest, com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> getUpdateAffiliateStatusMethod;
    if ((getUpdateAffiliateStatusMethod = AffiliateServiceGrpc.getUpdateAffiliateStatusMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getUpdateAffiliateStatusMethod = AffiliateServiceGrpc.getUpdateAffiliateStatusMethod) == null) {
          AffiliateServiceGrpc.getUpdateAffiliateStatusMethod = getUpdateAffiliateStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest, com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateAffiliateStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("UpdateAffiliateStatus"))
              .build();
        }
      }
    }
    return getUpdateAffiliateStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest,
      com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "TrackClick",
      requestType = com.game_engine.affiliate.v1.TrackClickRequest.class,
      responseType = com.game_engine.affiliate.v1.TrackClickResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest,
      com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest, com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod;
    if ((getTrackClickMethod = AffiliateServiceGrpc.getTrackClickMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getTrackClickMethod = AffiliateServiceGrpc.getTrackClickMethod) == null) {
          AffiliateServiceGrpc.getTrackClickMethod = getTrackClickMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.TrackClickRequest, com.game_engine.affiliate.v1.TrackClickResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "TrackClick"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackClickRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackClickResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("TrackClick"))
              .build();
        }
      }
    }
    return getTrackClickMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackRegistrationRequest,
      com.game_engine.affiliate.v1.TrackRegistrationResponse> getTrackRegistrationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "TrackRegistration",
      requestType = com.game_engine.affiliate.v1.TrackRegistrationRequest.class,
      responseType = com.game_engine.affiliate.v1.TrackRegistrationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackRegistrationRequest,
      com.game_engine.affiliate.v1.TrackRegistrationResponse> getTrackRegistrationMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackRegistrationRequest, com.game_engine.affiliate.v1.TrackRegistrationResponse> getTrackRegistrationMethod;
    if ((getTrackRegistrationMethod = AffiliateServiceGrpc.getTrackRegistrationMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getTrackRegistrationMethod = AffiliateServiceGrpc.getTrackRegistrationMethod) == null) {
          AffiliateServiceGrpc.getTrackRegistrationMethod = getTrackRegistrationMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.TrackRegistrationRequest, com.game_engine.affiliate.v1.TrackRegistrationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "TrackRegistration"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackRegistrationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackRegistrationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("TrackRegistration"))
              .build();
        }
      }
    }
    return getTrackRegistrationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackFirstDepositRequest,
      com.game_engine.affiliate.v1.TrackFirstDepositResponse> getTrackFirstDepositMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "TrackFirstDeposit",
      requestType = com.game_engine.affiliate.v1.TrackFirstDepositRequest.class,
      responseType = com.game_engine.affiliate.v1.TrackFirstDepositResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackFirstDepositRequest,
      com.game_engine.affiliate.v1.TrackFirstDepositResponse> getTrackFirstDepositMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackFirstDepositRequest, com.game_engine.affiliate.v1.TrackFirstDepositResponse> getTrackFirstDepositMethod;
    if ((getTrackFirstDepositMethod = AffiliateServiceGrpc.getTrackFirstDepositMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getTrackFirstDepositMethod = AffiliateServiceGrpc.getTrackFirstDepositMethod) == null) {
          AffiliateServiceGrpc.getTrackFirstDepositMethod = getTrackFirstDepositMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.TrackFirstDepositRequest, com.game_engine.affiliate.v1.TrackFirstDepositResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "TrackFirstDeposit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackFirstDepositRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackFirstDepositResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("TrackFirstDeposit"))
              .build();
        }
      }
    }
    return getTrackFirstDepositMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetReferralsRequest,
      com.game_engine.affiliate.v1.GetReferralsResponse> getGetReferralsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetReferrals",
      requestType = com.game_engine.affiliate.v1.GetReferralsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetReferralsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetReferralsRequest,
      com.game_engine.affiliate.v1.GetReferralsResponse> getGetReferralsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetReferralsRequest, com.game_engine.affiliate.v1.GetReferralsResponse> getGetReferralsMethod;
    if ((getGetReferralsMethod = AffiliateServiceGrpc.getGetReferralsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetReferralsMethod = AffiliateServiceGrpc.getGetReferralsMethod) == null) {
          AffiliateServiceGrpc.getGetReferralsMethod = getGetReferralsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetReferralsRequest, com.game_engine.affiliate.v1.GetReferralsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetReferrals"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetReferralsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetReferralsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetReferrals"))
              .build();
        }
      }
    }
    return getGetReferralsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetCampaignReferralsRequest,
      com.game_engine.affiliate.v1.GetCampaignReferralsResponse> getGetCampaignReferralsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCampaignReferrals",
      requestType = com.game_engine.affiliate.v1.GetCampaignReferralsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetCampaignReferralsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetCampaignReferralsRequest,
      com.game_engine.affiliate.v1.GetCampaignReferralsResponse> getGetCampaignReferralsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetCampaignReferralsRequest, com.game_engine.affiliate.v1.GetCampaignReferralsResponse> getGetCampaignReferralsMethod;
    if ((getGetCampaignReferralsMethod = AffiliateServiceGrpc.getGetCampaignReferralsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetCampaignReferralsMethod = AffiliateServiceGrpc.getGetCampaignReferralsMethod) == null) {
          AffiliateServiceGrpc.getGetCampaignReferralsMethod = getGetCampaignReferralsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetCampaignReferralsRequest, com.game_engine.affiliate.v1.GetCampaignReferralsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCampaignReferrals"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetCampaignReferralsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetCampaignReferralsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetCampaignReferrals"))
              .build();
        }
      }
    }
    return getGetCampaignReferralsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CalculateCommissionRequest,
      com.game_engine.affiliate.v1.CalculateCommissionResponse> getCalculateCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CalculateCommission",
      requestType = com.game_engine.affiliate.v1.CalculateCommissionRequest.class,
      responseType = com.game_engine.affiliate.v1.CalculateCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CalculateCommissionRequest,
      com.game_engine.affiliate.v1.CalculateCommissionResponse> getCalculateCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CalculateCommissionRequest, com.game_engine.affiliate.v1.CalculateCommissionResponse> getCalculateCommissionMethod;
    if ((getCalculateCommissionMethod = AffiliateServiceGrpc.getCalculateCommissionMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getCalculateCommissionMethod = AffiliateServiceGrpc.getCalculateCommissionMethod) == null) {
          AffiliateServiceGrpc.getCalculateCommissionMethod = getCalculateCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.CalculateCommissionRequest, com.game_engine.affiliate.v1.CalculateCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CalculateCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CalculateCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CalculateCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("CalculateCommission"))
              .build();
        }
      }
    }
    return getCalculateCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.AddSubAffiliateRequest,
      com.game_engine.affiliate.v1.AddSubAffiliateResponse> getAddSubAffiliateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AddSubAffiliate",
      requestType = com.game_engine.affiliate.v1.AddSubAffiliateRequest.class,
      responseType = com.game_engine.affiliate.v1.AddSubAffiliateResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.AddSubAffiliateRequest,
      com.game_engine.affiliate.v1.AddSubAffiliateResponse> getAddSubAffiliateMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.AddSubAffiliateRequest, com.game_engine.affiliate.v1.AddSubAffiliateResponse> getAddSubAffiliateMethod;
    if ((getAddSubAffiliateMethod = AffiliateServiceGrpc.getAddSubAffiliateMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getAddSubAffiliateMethod = AffiliateServiceGrpc.getAddSubAffiliateMethod) == null) {
          AffiliateServiceGrpc.getAddSubAffiliateMethod = getAddSubAffiliateMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.AddSubAffiliateRequest, com.game_engine.affiliate.v1.AddSubAffiliateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AddSubAffiliate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.AddSubAffiliateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.AddSubAffiliateResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("AddSubAffiliate"))
              .build();
        }
      }
    }
    return getAddSubAffiliateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetSubAffiliatesRequest,
      com.game_engine.affiliate.v1.GetSubAffiliatesResponse> getGetSubAffiliatesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSubAffiliates",
      requestType = com.game_engine.affiliate.v1.GetSubAffiliatesRequest.class,
      responseType = com.game_engine.affiliate.v1.GetSubAffiliatesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetSubAffiliatesRequest,
      com.game_engine.affiliate.v1.GetSubAffiliatesResponse> getGetSubAffiliatesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetSubAffiliatesRequest, com.game_engine.affiliate.v1.GetSubAffiliatesResponse> getGetSubAffiliatesMethod;
    if ((getGetSubAffiliatesMethod = AffiliateServiceGrpc.getGetSubAffiliatesMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetSubAffiliatesMethod = AffiliateServiceGrpc.getGetSubAffiliatesMethod) == null) {
          AffiliateServiceGrpc.getGetSubAffiliatesMethod = getGetSubAffiliatesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetSubAffiliatesRequest, com.game_engine.affiliate.v1.GetSubAffiliatesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSubAffiliates"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetSubAffiliatesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetSubAffiliatesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetSubAffiliates"))
              .build();
        }
      }
    }
    return getGetSubAffiliatesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateStatsRequest,
      com.game_engine.affiliate.v1.GetAffiliateStatsResponse> getGetAffiliateStatsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAffiliateStats",
      requestType = com.game_engine.affiliate.v1.GetAffiliateStatsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetAffiliateStatsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateStatsRequest,
      com.game_engine.affiliate.v1.GetAffiliateStatsResponse> getGetAffiliateStatsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateStatsRequest, com.game_engine.affiliate.v1.GetAffiliateStatsResponse> getGetAffiliateStatsMethod;
    if ((getGetAffiliateStatsMethod = AffiliateServiceGrpc.getGetAffiliateStatsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetAffiliateStatsMethod = AffiliateServiceGrpc.getGetAffiliateStatsMethod) == null) {
          AffiliateServiceGrpc.getGetAffiliateStatsMethod = getGetAffiliateStatsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetAffiliateStatsRequest, com.game_engine.affiliate.v1.GetAffiliateStatsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAffiliateStats"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateStatsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateStatsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetAffiliateStats"))
              .build();
        }
      }
    }
    return getGetAffiliateStatsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RedirectToRegistrationRequest,
      com.game_engine.affiliate.v1.RedirectToRegistrationResponse> getRedirectToRegistrationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RedirectToRegistration",
      requestType = com.game_engine.affiliate.v1.RedirectToRegistrationRequest.class,
      responseType = com.game_engine.affiliate.v1.RedirectToRegistrationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RedirectToRegistrationRequest,
      com.game_engine.affiliate.v1.RedirectToRegistrationResponse> getRedirectToRegistrationMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.RedirectToRegistrationRequest, com.game_engine.affiliate.v1.RedirectToRegistrationResponse> getRedirectToRegistrationMethod;
    if ((getRedirectToRegistrationMethod = AffiliateServiceGrpc.getRedirectToRegistrationMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getRedirectToRegistrationMethod = AffiliateServiceGrpc.getRedirectToRegistrationMethod) == null) {
          AffiliateServiceGrpc.getRedirectToRegistrationMethod = getRedirectToRegistrationMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.RedirectToRegistrationRequest, com.game_engine.affiliate.v1.RedirectToRegistrationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RedirectToRegistration"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.RedirectToRegistrationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.RedirectToRegistrationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("RedirectToRegistration"))
              .build();
        }
      }
    }
    return getRedirectToRegistrationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest,
      com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPerformanceReport",
      requestType = com.game_engine.affiliate.v1.GetPerformanceReportRequest.class,
      responseType = com.game_engine.affiliate.v1.GetPerformanceReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest,
      com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest, com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod;
    if ((getGetPerformanceReportMethod = AffiliateServiceGrpc.getGetPerformanceReportMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetPerformanceReportMethod = AffiliateServiceGrpc.getGetPerformanceReportMethod) == null) {
          AffiliateServiceGrpc.getGetPerformanceReportMethod = getGetPerformanceReportMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetPerformanceReportRequest, com.game_engine.affiliate.v1.GetPerformanceReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPerformanceReport"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetPerformanceReportRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetPerformanceReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetPerformanceReport"))
              .build();
        }
      }
    }
    return getGetPerformanceReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest,
      com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClickReports",
      requestType = com.game_engine.affiliate.v1.GetClickReportsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetClickReportsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest,
      com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest, com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod;
    if ((getGetClickReportsMethod = AffiliateServiceGrpc.getGetClickReportsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetClickReportsMethod = AffiliateServiceGrpc.getGetClickReportsMethod) == null) {
          AffiliateServiceGrpc.getGetClickReportsMethod = getGetClickReportsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetClickReportsRequest, com.game_engine.affiliate.v1.GetClickReportsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClickReports"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetClickReportsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetClickReportsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetClickReports"))
              .build();
        }
      }
    }
    return getGetClickReportsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest,
      com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetConversionReports",
      requestType = com.game_engine.affiliate.v1.GetConversionReportsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetConversionReportsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest,
      com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest, com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod;
    if ((getGetConversionReportsMethod = AffiliateServiceGrpc.getGetConversionReportsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetConversionReportsMethod = AffiliateServiceGrpc.getGetConversionReportsMethod) == null) {
          AffiliateServiceGrpc.getGetConversionReportsMethod = getGetConversionReportsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetConversionReportsRequest, com.game_engine.affiliate.v1.GetConversionReportsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConversionReports"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetConversionReportsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetConversionReportsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetConversionReports"))
              .build();
        }
      }
    }
    return getGetConversionReportsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
      com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAffiliateLinks",
      requestType = com.game_engine.affiliate.v1.GetAffiliateLinksRequest.class,
      responseType = com.game_engine.affiliate.v1.GetAffiliateLinksResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
      com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest, com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod;
    if ((getGetAffiliateLinksMethod = AffiliateServiceGrpc.getGetAffiliateLinksMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetAffiliateLinksMethod = AffiliateServiceGrpc.getGetAffiliateLinksMethod) == null) {
          AffiliateServiceGrpc.getGetAffiliateLinksMethod = getGetAffiliateLinksMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetAffiliateLinksRequest, com.game_engine.affiliate.v1.GetAffiliateLinksResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAffiliateLinks"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateLinksRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateLinksResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetAffiliateLinks"))
              .build();
        }
      }
    }
    return getGetAffiliateLinksMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
      com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAffiliateLink",
      requestType = com.game_engine.affiliate.v1.CreateAffiliateLinkRequest.class,
      responseType = com.game_engine.affiliate.v1.CreateAffiliateLinkResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
      com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest, com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod;
    if ((getCreateAffiliateLinkMethod = AffiliateServiceGrpc.getCreateAffiliateLinkMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getCreateAffiliateLinkMethod = AffiliateServiceGrpc.getCreateAffiliateLinkMethod) == null) {
          AffiliateServiceGrpc.getCreateAffiliateLinkMethod = getCreateAffiliateLinkMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest, com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAffiliateLink"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CreateAffiliateLinkRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CreateAffiliateLinkResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("CreateAffiliateLink"))
              .build();
        }
      }
    }
    return getCreateAffiliateLinkMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AffiliateServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceStub>() {
        @java.lang.Override
        public AffiliateServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceStub(channel, callOptions);
        }
      };
    return AffiliateServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static AffiliateServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingV2Stub>() {
        @java.lang.Override
        public AffiliateServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return AffiliateServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AffiliateServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingStub>() {
        @java.lang.Override
        public AffiliateServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceBlockingStub(channel, callOptions);
        }
      };
    return AffiliateServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AffiliateServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceFutureStub>() {
        @java.lang.Override
        public AffiliateServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceFutureStub(channel, callOptions);
        }
      };
    return AffiliateServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Affiliate management
     * </pre>
     */
    default void registerAffiliate(com.game_engine.affiliate.v1.RegisterAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RegisterAffiliateResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterAffiliateMethod(), responseObserver);
    }

    /**
     */
    default void getAffiliateByCode(com.game_engine.affiliate.v1.GetAffiliateByCodeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAffiliateByCodeMethod(), responseObserver);
    }

    /**
     */
    default void getAffiliatesByMerchant(com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAffiliatesByMerchantMethod(), responseObserver);
    }

    /**
     */
    default void getActiveAffiliates(com.game_engine.affiliate.v1.GetActiveAffiliatesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveAffiliatesMethod(), responseObserver);
    }

    /**
     */
    default void updateAffiliateTier(com.game_engine.affiliate.v1.UpdateAffiliateTierRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateAffiliateTierMethod(), responseObserver);
    }

    /**
     */
    default void updateAffiliateStatus(com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateAffiliateStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Tracking
     * </pre>
     */
    default void trackClick(com.game_engine.affiliate.v1.TrackClickRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getTrackClickMethod(), responseObserver);
    }

    /**
     */
    default void trackRegistration(com.game_engine.affiliate.v1.TrackRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackRegistrationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getTrackRegistrationMethod(), responseObserver);
    }

    /**
     */
    default void trackFirstDeposit(com.game_engine.affiliate.v1.TrackFirstDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackFirstDepositResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getTrackFirstDepositMethod(), responseObserver);
    }

    /**
     * <pre>
     * Referrals
     * </pre>
     */
    default void getReferrals(com.game_engine.affiliate.v1.GetReferralsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetReferralsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetReferralsMethod(), responseObserver);
    }

    /**
     */
    default void getCampaignReferrals(com.game_engine.affiliate.v1.GetCampaignReferralsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetCampaignReferralsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCampaignReferralsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Commission
     * </pre>
     */
    default void calculateCommission(com.game_engine.affiliate.v1.CalculateCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CalculateCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCalculateCommissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * Sub-affiliates
     * </pre>
     */
    default void addSubAffiliate(com.game_engine.affiliate.v1.AddSubAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.AddSubAffiliateResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAddSubAffiliateMethod(), responseObserver);
    }

    /**
     */
    default void getSubAffiliates(com.game_engine.affiliate.v1.GetSubAffiliatesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetSubAffiliatesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSubAffiliatesMethod(), responseObserver);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    default void getAffiliateStats(com.game_engine.affiliate.v1.GetAffiliateStatsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateStatsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAffiliateStatsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Redirect
     * </pre>
     */
    default void redirectToRegistration(com.game_engine.affiliate.v1.RedirectToRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RedirectToRegistrationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRedirectToRegistrationMethod(), responseObserver);
    }

    /**
     * <pre>
     * Reporting and link management
     * </pre>
     */
    default void getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPerformanceReportMethod(), responseObserver);
    }

    /**
     */
    default void getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClickReportsMethod(), responseObserver);
    }

    /**
     */
    default void getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConversionReportsMethod(), responseObserver);
    }

    /**
     */
    default void getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAffiliateLinksMethod(), responseObserver);
    }

    /**
     */
    default void createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAffiliateLinkMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AffiliateService.
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public static abstract class AffiliateServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AffiliateServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AffiliateService.
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public static final class AffiliateServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AffiliateServiceStub> {
    private AffiliateServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Affiliate management
     * </pre>
     */
    public void registerAffiliate(com.game_engine.affiliate.v1.RegisterAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RegisterAffiliateResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterAffiliateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAffiliateByCode(com.game_engine.affiliate.v1.GetAffiliateByCodeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAffiliateByCodeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAffiliatesByMerchant(com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAffiliatesByMerchantMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getActiveAffiliates(com.game_engine.affiliate.v1.GetActiveAffiliatesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveAffiliatesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateAffiliateTier(com.game_engine.affiliate.v1.UpdateAffiliateTierRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateAffiliateTierMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateAffiliateStatus(com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateAffiliateStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Tracking
     * </pre>
     */
    public void trackClick(com.game_engine.affiliate.v1.TrackClickRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getTrackClickMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void trackRegistration(com.game_engine.affiliate.v1.TrackRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackRegistrationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getTrackRegistrationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void trackFirstDeposit(com.game_engine.affiliate.v1.TrackFirstDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackFirstDepositResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getTrackFirstDepositMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Referrals
     * </pre>
     */
    public void getReferrals(com.game_engine.affiliate.v1.GetReferralsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetReferralsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetReferralsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCampaignReferrals(com.game_engine.affiliate.v1.GetCampaignReferralsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetCampaignReferralsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCampaignReferralsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Commission
     * </pre>
     */
    public void calculateCommission(com.game_engine.affiliate.v1.CalculateCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CalculateCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCalculateCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Sub-affiliates
     * </pre>
     */
    public void addSubAffiliate(com.game_engine.affiliate.v1.AddSubAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.AddSubAffiliateResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAddSubAffiliateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSubAffiliates(com.game_engine.affiliate.v1.GetSubAffiliatesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetSubAffiliatesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSubAffiliatesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public void getAffiliateStats(com.game_engine.affiliate.v1.GetAffiliateStatsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateStatsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAffiliateStatsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Redirect
     * </pre>
     */
    public void redirectToRegistration(com.game_engine.affiliate.v1.RedirectToRegistrationRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RedirectToRegistrationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRedirectToRegistrationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Reporting and link management
     * </pre>
     */
    public void getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPerformanceReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClickReportsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConversionReportsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAffiliateLinksMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAffiliateLinkMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AffiliateService.
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public static final class AffiliateServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<AffiliateServiceBlockingV2Stub> {
    private AffiliateServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Affiliate management
     * </pre>
     */
    public com.game_engine.affiliate.v1.RegisterAffiliateResponse registerAffiliate(com.game_engine.affiliate.v1.RegisterAffiliateRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRegisterAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliateByCodeResponse getAffiliateByCode(com.game_engine.affiliate.v1.GetAffiliateByCodeRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAffiliateByCodeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse getAffiliatesByMerchant(com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAffiliatesByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetActiveAffiliatesResponse getActiveAffiliates(com.game_engine.affiliate.v1.GetActiveAffiliatesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetActiveAffiliatesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.UpdateAffiliateTierResponse updateAffiliateTier(com.game_engine.affiliate.v1.UpdateAffiliateTierRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateAffiliateTierMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse updateAffiliateStatus(com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateAffiliateStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Tracking
     * </pre>
     */
    public com.game_engine.affiliate.v1.TrackClickResponse trackClick(com.game_engine.affiliate.v1.TrackClickRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getTrackClickMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.TrackRegistrationResponse trackRegistration(com.game_engine.affiliate.v1.TrackRegistrationRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getTrackRegistrationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.TrackFirstDepositResponse trackFirstDeposit(com.game_engine.affiliate.v1.TrackFirstDepositRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getTrackFirstDepositMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Referrals
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetReferralsResponse getReferrals(com.game_engine.affiliate.v1.GetReferralsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetReferralsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetCampaignReferralsResponse getCampaignReferrals(com.game_engine.affiliate.v1.GetCampaignReferralsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCampaignReferralsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Commission
     * </pre>
     */
    public com.game_engine.affiliate.v1.CalculateCommissionResponse calculateCommission(com.game_engine.affiliate.v1.CalculateCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCalculateCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Sub-affiliates
     * </pre>
     */
    public com.game_engine.affiliate.v1.AddSubAffiliateResponse addSubAffiliate(com.game_engine.affiliate.v1.AddSubAffiliateRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getAddSubAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetSubAffiliatesResponse getSubAffiliates(com.game_engine.affiliate.v1.GetSubAffiliatesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSubAffiliatesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetAffiliateStatsResponse getAffiliateStats(com.game_engine.affiliate.v1.GetAffiliateStatsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAffiliateStatsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Redirect
     * </pre>
     */
    public com.game_engine.affiliate.v1.RedirectToRegistrationResponse redirectToRegistration(com.game_engine.affiliate.v1.RedirectToRegistrationRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRedirectToRegistrationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reporting and link management
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetPerformanceReportResponse getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPerformanceReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetClickReportsResponse getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClickReportsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetConversionReportsResponse getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetConversionReportsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliateLinksResponse getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAffiliateLinksMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.CreateAffiliateLinkResponse createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateAffiliateLinkMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service AffiliateService.
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public static final class AffiliateServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AffiliateServiceBlockingStub> {
    private AffiliateServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Affiliate management
     * </pre>
     */
    public com.game_engine.affiliate.v1.RegisterAffiliateResponse registerAffiliate(com.game_engine.affiliate.v1.RegisterAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliateByCodeResponse getAffiliateByCode(com.game_engine.affiliate.v1.GetAffiliateByCodeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAffiliateByCodeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse getAffiliatesByMerchant(com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAffiliatesByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetActiveAffiliatesResponse getActiveAffiliates(com.game_engine.affiliate.v1.GetActiveAffiliatesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveAffiliatesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.UpdateAffiliateTierResponse updateAffiliateTier(com.game_engine.affiliate.v1.UpdateAffiliateTierRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateAffiliateTierMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse updateAffiliateStatus(com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateAffiliateStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Tracking
     * </pre>
     */
    public com.game_engine.affiliate.v1.TrackClickResponse trackClick(com.game_engine.affiliate.v1.TrackClickRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getTrackClickMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.TrackRegistrationResponse trackRegistration(com.game_engine.affiliate.v1.TrackRegistrationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getTrackRegistrationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.TrackFirstDepositResponse trackFirstDeposit(com.game_engine.affiliate.v1.TrackFirstDepositRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getTrackFirstDepositMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Referrals
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetReferralsResponse getReferrals(com.game_engine.affiliate.v1.GetReferralsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetReferralsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetCampaignReferralsResponse getCampaignReferrals(com.game_engine.affiliate.v1.GetCampaignReferralsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCampaignReferralsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Commission
     * </pre>
     */
    public com.game_engine.affiliate.v1.CalculateCommissionResponse calculateCommission(com.game_engine.affiliate.v1.CalculateCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCalculateCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Sub-affiliates
     * </pre>
     */
    public com.game_engine.affiliate.v1.AddSubAffiliateResponse addSubAffiliate(com.game_engine.affiliate.v1.AddSubAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAddSubAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetSubAffiliatesResponse getSubAffiliates(com.game_engine.affiliate.v1.GetSubAffiliatesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSubAffiliatesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetAffiliateStatsResponse getAffiliateStats(com.game_engine.affiliate.v1.GetAffiliateStatsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAffiliateStatsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Redirect
     * </pre>
     */
    public com.game_engine.affiliate.v1.RedirectToRegistrationResponse redirectToRegistration(com.game_engine.affiliate.v1.RedirectToRegistrationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRedirectToRegistrationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reporting and link management
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetPerformanceReportResponse getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPerformanceReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetClickReportsResponse getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClickReportsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetConversionReportsResponse getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConversionReportsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.GetAffiliateLinksResponse getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAffiliateLinksMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.affiliate.v1.CreateAffiliateLinkResponse createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAffiliateLinkMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AffiliateService.
   * <pre>
   * Affiliate Service - manages affiliate registration, tracking, and commissions
   * </pre>
   */
  public static final class AffiliateServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AffiliateServiceFutureStub> {
    private AffiliateServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Affiliate management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.RegisterAffiliateResponse> registerAffiliate(
        com.game_engine.affiliate.v1.RegisterAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterAffiliateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetAffiliateByCodeResponse> getAffiliateByCode(
        com.game_engine.affiliate.v1.GetAffiliateByCodeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAffiliateByCodeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse> getAffiliatesByMerchant(
        com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAffiliatesByMerchantMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetActiveAffiliatesResponse> getActiveAffiliates(
        com.game_engine.affiliate.v1.GetActiveAffiliatesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveAffiliatesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.UpdateAffiliateTierResponse> updateAffiliateTier(
        com.game_engine.affiliate.v1.UpdateAffiliateTierRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateAffiliateTierMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse> updateAffiliateStatus(
        com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateAffiliateStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Tracking
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.TrackClickResponse> trackClick(
        com.game_engine.affiliate.v1.TrackClickRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getTrackClickMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.TrackRegistrationResponse> trackRegistration(
        com.game_engine.affiliate.v1.TrackRegistrationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getTrackRegistrationMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.TrackFirstDepositResponse> trackFirstDeposit(
        com.game_engine.affiliate.v1.TrackFirstDepositRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getTrackFirstDepositMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Referrals
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetReferralsResponse> getReferrals(
        com.game_engine.affiliate.v1.GetReferralsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetReferralsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetCampaignReferralsResponse> getCampaignReferrals(
        com.game_engine.affiliate.v1.GetCampaignReferralsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCampaignReferralsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Commission
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.CalculateCommissionResponse> calculateCommission(
        com.game_engine.affiliate.v1.CalculateCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCalculateCommissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Sub-affiliates
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.AddSubAffiliateResponse> addSubAffiliate(
        com.game_engine.affiliate.v1.AddSubAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAddSubAffiliateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetSubAffiliatesResponse> getSubAffiliates(
        com.game_engine.affiliate.v1.GetSubAffiliatesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSubAffiliatesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetAffiliateStatsResponse> getAffiliateStats(
        com.game_engine.affiliate.v1.GetAffiliateStatsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAffiliateStatsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Redirect
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.RedirectToRegistrationResponse> redirectToRegistration(
        com.game_engine.affiliate.v1.RedirectToRegistrationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRedirectToRegistrationMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Reporting and link management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetPerformanceReportResponse> getPerformanceReport(
        com.game_engine.affiliate.v1.GetPerformanceReportRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPerformanceReportMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetClickReportsResponse> getClickReports(
        com.game_engine.affiliate.v1.GetClickReportsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClickReportsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetConversionReportsResponse> getConversionReports(
        com.game_engine.affiliate.v1.GetConversionReportsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConversionReportsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getAffiliateLinks(
        com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAffiliateLinksMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> createAffiliateLink(
        com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAffiliateLinkMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_REGISTER_AFFILIATE = 0;
  private static final int METHODID_GET_AFFILIATE_BY_CODE = 1;
  private static final int METHODID_GET_AFFILIATES_BY_MERCHANT = 2;
  private static final int METHODID_GET_ACTIVE_AFFILIATES = 3;
  private static final int METHODID_UPDATE_AFFILIATE_TIER = 4;
  private static final int METHODID_UPDATE_AFFILIATE_STATUS = 5;
  private static final int METHODID_TRACK_CLICK = 6;
  private static final int METHODID_TRACK_REGISTRATION = 7;
  private static final int METHODID_TRACK_FIRST_DEPOSIT = 8;
  private static final int METHODID_GET_REFERRALS = 9;
  private static final int METHODID_GET_CAMPAIGN_REFERRALS = 10;
  private static final int METHODID_CALCULATE_COMMISSION = 11;
  private static final int METHODID_ADD_SUB_AFFILIATE = 12;
  private static final int METHODID_GET_SUB_AFFILIATES = 13;
  private static final int METHODID_GET_AFFILIATE_STATS = 14;
  private static final int METHODID_REDIRECT_TO_REGISTRATION = 15;
  private static final int METHODID_GET_PERFORMANCE_REPORT = 16;
  private static final int METHODID_GET_CLICK_REPORTS = 17;
  private static final int METHODID_GET_CONVERSION_REPORTS = 18;
  private static final int METHODID_GET_AFFILIATE_LINKS = 19;
  private static final int METHODID_CREATE_AFFILIATE_LINK = 20;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_REGISTER_AFFILIATE:
          serviceImpl.registerAffiliate((com.game_engine.affiliate.v1.RegisterAffiliateRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RegisterAffiliateResponse>) responseObserver);
          break;
        case METHODID_GET_AFFILIATE_BY_CODE:
          serviceImpl.getAffiliateByCode((com.game_engine.affiliate.v1.GetAffiliateByCodeRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateByCodeResponse>) responseObserver);
          break;
        case METHODID_GET_AFFILIATES_BY_MERCHANT:
          serviceImpl.getAffiliatesByMerchant((com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse>) responseObserver);
          break;
        case METHODID_GET_ACTIVE_AFFILIATES:
          serviceImpl.getActiveAffiliates((com.game_engine.affiliate.v1.GetActiveAffiliatesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetActiveAffiliatesResponse>) responseObserver);
          break;
        case METHODID_UPDATE_AFFILIATE_TIER:
          serviceImpl.updateAffiliateTier((com.game_engine.affiliate.v1.UpdateAffiliateTierRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateTierResponse>) responseObserver);
          break;
        case METHODID_UPDATE_AFFILIATE_STATUS:
          serviceImpl.updateAffiliateStatus((com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse>) responseObserver);
          break;
        case METHODID_TRACK_CLICK:
          serviceImpl.trackClick((com.game_engine.affiliate.v1.TrackClickRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse>) responseObserver);
          break;
        case METHODID_TRACK_REGISTRATION:
          serviceImpl.trackRegistration((com.game_engine.affiliate.v1.TrackRegistrationRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackRegistrationResponse>) responseObserver);
          break;
        case METHODID_TRACK_FIRST_DEPOSIT:
          serviceImpl.trackFirstDeposit((com.game_engine.affiliate.v1.TrackFirstDepositRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackFirstDepositResponse>) responseObserver);
          break;
        case METHODID_GET_REFERRALS:
          serviceImpl.getReferrals((com.game_engine.affiliate.v1.GetReferralsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetReferralsResponse>) responseObserver);
          break;
        case METHODID_GET_CAMPAIGN_REFERRALS:
          serviceImpl.getCampaignReferrals((com.game_engine.affiliate.v1.GetCampaignReferralsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetCampaignReferralsResponse>) responseObserver);
          break;
        case METHODID_CALCULATE_COMMISSION:
          serviceImpl.calculateCommission((com.game_engine.affiliate.v1.CalculateCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CalculateCommissionResponse>) responseObserver);
          break;
        case METHODID_ADD_SUB_AFFILIATE:
          serviceImpl.addSubAffiliate((com.game_engine.affiliate.v1.AddSubAffiliateRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.AddSubAffiliateResponse>) responseObserver);
          break;
        case METHODID_GET_SUB_AFFILIATES:
          serviceImpl.getSubAffiliates((com.game_engine.affiliate.v1.GetSubAffiliatesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetSubAffiliatesResponse>) responseObserver);
          break;
        case METHODID_GET_AFFILIATE_STATS:
          serviceImpl.getAffiliateStats((com.game_engine.affiliate.v1.GetAffiliateStatsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateStatsResponse>) responseObserver);
          break;
        case METHODID_REDIRECT_TO_REGISTRATION:
          serviceImpl.redirectToRegistration((com.game_engine.affiliate.v1.RedirectToRegistrationRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.RedirectToRegistrationResponse>) responseObserver);
          break;
        case METHODID_GET_PERFORMANCE_REPORT:
          serviceImpl.getPerformanceReport((com.game_engine.affiliate.v1.GetPerformanceReportRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse>) responseObserver);
          break;
        case METHODID_GET_CLICK_REPORTS:
          serviceImpl.getClickReports((com.game_engine.affiliate.v1.GetClickReportsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse>) responseObserver);
          break;
        case METHODID_GET_CONVERSION_REPORTS:
          serviceImpl.getConversionReports((com.game_engine.affiliate.v1.GetConversionReportsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse>) responseObserver);
          break;
        case METHODID_GET_AFFILIATE_LINKS:
          serviceImpl.getAffiliateLinks((com.game_engine.affiliate.v1.GetAffiliateLinksRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse>) responseObserver);
          break;
        case METHODID_CREATE_AFFILIATE_LINK:
          serviceImpl.createAffiliateLink((com.game_engine.affiliate.v1.CreateAffiliateLinkRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getRegisterAffiliateMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.RegisterAffiliateRequest,
              com.game_engine.affiliate.v1.RegisterAffiliateResponse>(
                service, METHODID_REGISTER_AFFILIATE)))
        .addMethod(
          getGetAffiliateByCodeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetAffiliateByCodeRequest,
              com.game_engine.affiliate.v1.GetAffiliateByCodeResponse>(
                service, METHODID_GET_AFFILIATE_BY_CODE)))
        .addMethod(
          getGetAffiliatesByMerchantMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetAffiliatesByMerchantRequest,
              com.game_engine.affiliate.v1.GetAffiliatesByMerchantResponse>(
                service, METHODID_GET_AFFILIATES_BY_MERCHANT)))
        .addMethod(
          getGetActiveAffiliatesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetActiveAffiliatesRequest,
              com.game_engine.affiliate.v1.GetActiveAffiliatesResponse>(
                service, METHODID_GET_ACTIVE_AFFILIATES)))
        .addMethod(
          getUpdateAffiliateTierMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.UpdateAffiliateTierRequest,
              com.game_engine.affiliate.v1.UpdateAffiliateTierResponse>(
                service, METHODID_UPDATE_AFFILIATE_TIER)))
        .addMethod(
          getUpdateAffiliateStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.UpdateAffiliateStatusRequest,
              com.game_engine.affiliate.v1.UpdateAffiliateStatusResponse>(
                service, METHODID_UPDATE_AFFILIATE_STATUS)))
        .addMethod(
          getTrackClickMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.TrackClickRequest,
              com.game_engine.affiliate.v1.TrackClickResponse>(
                service, METHODID_TRACK_CLICK)))
        .addMethod(
          getTrackRegistrationMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.TrackRegistrationRequest,
              com.game_engine.affiliate.v1.TrackRegistrationResponse>(
                service, METHODID_TRACK_REGISTRATION)))
        .addMethod(
          getTrackFirstDepositMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.TrackFirstDepositRequest,
              com.game_engine.affiliate.v1.TrackFirstDepositResponse>(
                service, METHODID_TRACK_FIRST_DEPOSIT)))
        .addMethod(
          getGetReferralsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetReferralsRequest,
              com.game_engine.affiliate.v1.GetReferralsResponse>(
                service, METHODID_GET_REFERRALS)))
        .addMethod(
          getGetCampaignReferralsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetCampaignReferralsRequest,
              com.game_engine.affiliate.v1.GetCampaignReferralsResponse>(
                service, METHODID_GET_CAMPAIGN_REFERRALS)))
        .addMethod(
          getCalculateCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.CalculateCommissionRequest,
              com.game_engine.affiliate.v1.CalculateCommissionResponse>(
                service, METHODID_CALCULATE_COMMISSION)))
        .addMethod(
          getAddSubAffiliateMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.AddSubAffiliateRequest,
              com.game_engine.affiliate.v1.AddSubAffiliateResponse>(
                service, METHODID_ADD_SUB_AFFILIATE)))
        .addMethod(
          getGetSubAffiliatesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetSubAffiliatesRequest,
              com.game_engine.affiliate.v1.GetSubAffiliatesResponse>(
                service, METHODID_GET_SUB_AFFILIATES)))
        .addMethod(
          getGetAffiliateStatsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetAffiliateStatsRequest,
              com.game_engine.affiliate.v1.GetAffiliateStatsResponse>(
                service, METHODID_GET_AFFILIATE_STATS)))
        .addMethod(
          getRedirectToRegistrationMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.RedirectToRegistrationRequest,
              com.game_engine.affiliate.v1.RedirectToRegistrationResponse>(
                service, METHODID_REDIRECT_TO_REGISTRATION)))
        .addMethod(
          getGetPerformanceReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetPerformanceReportRequest,
              com.game_engine.affiliate.v1.GetPerformanceReportResponse>(
                service, METHODID_GET_PERFORMANCE_REPORT)))
        .addMethod(
          getGetClickReportsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetClickReportsRequest,
              com.game_engine.affiliate.v1.GetClickReportsResponse>(
                service, METHODID_GET_CLICK_REPORTS)))
        .addMethod(
          getGetConversionReportsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetConversionReportsRequest,
              com.game_engine.affiliate.v1.GetConversionReportsResponse>(
                service, METHODID_GET_CONVERSION_REPORTS)))
        .addMethod(
          getGetAffiliateLinksMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
              com.game_engine.affiliate.v1.GetAffiliateLinksResponse>(
                service, METHODID_GET_AFFILIATE_LINKS)))
        .addMethod(
          getCreateAffiliateLinkMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
              com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>(
                service, METHODID_CREATE_AFFILIATE_LINK)))
        .build();
  }

  private static abstract class AffiliateServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AffiliateServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.affiliate.v1.AffiliateServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AffiliateService");
    }
  }

  private static final class AffiliateServiceFileDescriptorSupplier
      extends AffiliateServiceBaseDescriptorSupplier {
    AffiliateServiceFileDescriptorSupplier() {}
  }

  private static final class AffiliateServiceMethodDescriptorSupplier
      extends AffiliateServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AffiliateServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AffiliateServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AffiliateServiceFileDescriptorSupplier())
              .addMethod(getRegisterAffiliateMethod())
              .addMethod(getGetAffiliateByCodeMethod())
              .addMethod(getGetAffiliatesByMerchantMethod())
              .addMethod(getGetActiveAffiliatesMethod())
              .addMethod(getUpdateAffiliateTierMethod())
              .addMethod(getUpdateAffiliateStatusMethod())
              .addMethod(getTrackClickMethod())
              .addMethod(getTrackRegistrationMethod())
              .addMethod(getTrackFirstDepositMethod())
              .addMethod(getGetReferralsMethod())
              .addMethod(getGetCampaignReferralsMethod())
              .addMethod(getCalculateCommissionMethod())
              .addMethod(getAddSubAffiliateMethod())
              .addMethod(getGetSubAffiliatesMethod())
              .addMethod(getGetAffiliateStatsMethod())
              .addMethod(getRedirectToRegistrationMethod())
              .addMethod(getGetPerformanceReportMethod())
              .addMethod(getGetClickReportsMethod())
              .addMethod(getGetConversionReportsMethod())
              .addMethod(getGetAffiliateLinksMethod())
              .addMethod(getCreateAffiliateLinkMethod())
              .build();
        }
      }
    }
    return result;
  }
}
