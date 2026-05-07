package com.game_engine.bonus.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Bonus Service - manages bonuses, claims, wagering requirements
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class BonusServiceGrpc {

  private BonusServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.bonus.v1.BonusService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusesRequest,
      com.game_engine.bonus.v1.GetActiveBonusesResponse> getGetActiveBonusesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActiveBonuses",
      requestType = com.game_engine.bonus.v1.GetActiveBonusesRequest.class,
      responseType = com.game_engine.bonus.v1.GetActiveBonusesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusesRequest,
      com.game_engine.bonus.v1.GetActiveBonusesResponse> getGetActiveBonusesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusesRequest, com.game_engine.bonus.v1.GetActiveBonusesResponse> getGetActiveBonusesMethod;
    if ((getGetActiveBonusesMethod = BonusServiceGrpc.getGetActiveBonusesMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetActiveBonusesMethod = BonusServiceGrpc.getGetActiveBonusesMethod) == null) {
          BonusServiceGrpc.getGetActiveBonusesMethod = getGetActiveBonusesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetActiveBonusesRequest, com.game_engine.bonus.v1.GetActiveBonusesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveBonuses"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetActiveBonusesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetActiveBonusesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetActiveBonuses"))
              .build();
        }
      }
    }
    return getGetActiveBonusesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusRequest,
      com.game_engine.bonus.v1.GetBonusResponse> getGetBonusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBonus",
      requestType = com.game_engine.bonus.v1.GetBonusRequest.class,
      responseType = com.game_engine.bonus.v1.GetBonusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusRequest,
      com.game_engine.bonus.v1.GetBonusResponse> getGetBonusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusRequest, com.game_engine.bonus.v1.GetBonusResponse> getGetBonusMethod;
    if ((getGetBonusMethod = BonusServiceGrpc.getGetBonusMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetBonusMethod = BonusServiceGrpc.getGetBonusMethod) == null) {
          BonusServiceGrpc.getGetBonusMethod = getGetBonusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetBonusRequest, com.game_engine.bonus.v1.GetBonusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBonus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetBonus"))
              .build();
        }
      }
    }
    return getGetBonusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateBonusRequest,
      com.game_engine.bonus.v1.CreateBonusResponse> getCreateBonusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateBonus",
      requestType = com.game_engine.bonus.v1.CreateBonusRequest.class,
      responseType = com.game_engine.bonus.v1.CreateBonusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateBonusRequest,
      com.game_engine.bonus.v1.CreateBonusResponse> getCreateBonusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateBonusRequest, com.game_engine.bonus.v1.CreateBonusResponse> getCreateBonusMethod;
    if ((getCreateBonusMethod = BonusServiceGrpc.getCreateBonusMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getCreateBonusMethod = BonusServiceGrpc.getCreateBonusMethod) == null) {
          BonusServiceGrpc.getCreateBonusMethod = getCreateBonusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.CreateBonusRequest, com.game_engine.bonus.v1.CreateBonusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateBonus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CreateBonusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CreateBonusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("CreateBonus"))
              .build();
        }
      }
    }
    return getCreateBonusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimBonusRequest,
      com.game_engine.bonus.v1.ClaimBonusResponse> getClaimBonusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClaimBonus",
      requestType = com.game_engine.bonus.v1.ClaimBonusRequest.class,
      responseType = com.game_engine.bonus.v1.ClaimBonusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimBonusRequest,
      com.game_engine.bonus.v1.ClaimBonusResponse> getClaimBonusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimBonusRequest, com.game_engine.bonus.v1.ClaimBonusResponse> getClaimBonusMethod;
    if ((getClaimBonusMethod = BonusServiceGrpc.getClaimBonusMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getClaimBonusMethod = BonusServiceGrpc.getClaimBonusMethod) == null) {
          BonusServiceGrpc.getClaimBonusMethod = getClaimBonusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.ClaimBonusRequest, com.game_engine.bonus.v1.ClaimBonusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClaimBonus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ClaimBonusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ClaimBonusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("ClaimBonus"))
              .build();
        }
      }
    }
    return getClaimBonusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CheckEligibilityRequest,
      com.game_engine.bonus.v1.CheckEligibilityResponse> getCheckEligibilityMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CheckEligibility",
      requestType = com.game_engine.bonus.v1.CheckEligibilityRequest.class,
      responseType = com.game_engine.bonus.v1.CheckEligibilityResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CheckEligibilityRequest,
      com.game_engine.bonus.v1.CheckEligibilityResponse> getCheckEligibilityMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CheckEligibilityRequest, com.game_engine.bonus.v1.CheckEligibilityResponse> getCheckEligibilityMethod;
    if ((getCheckEligibilityMethod = BonusServiceGrpc.getCheckEligibilityMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getCheckEligibilityMethod = BonusServiceGrpc.getCheckEligibilityMethod) == null) {
          BonusServiceGrpc.getCheckEligibilityMethod = getCheckEligibilityMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.CheckEligibilityRequest, com.game_engine.bonus.v1.CheckEligibilityResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CheckEligibility"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CheckEligibilityRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CheckEligibilityResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("CheckEligibility"))
              .build();
        }
      }
    }
    return getCheckEligibilityMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusHistoryRequest,
      com.game_engine.bonus.v1.GetBonusHistoryResponse> getGetBonusHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBonusHistory",
      requestType = com.game_engine.bonus.v1.GetBonusHistoryRequest.class,
      responseType = com.game_engine.bonus.v1.GetBonusHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusHistoryRequest,
      com.game_engine.bonus.v1.GetBonusHistoryResponse> getGetBonusHistoryMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusHistoryRequest, com.game_engine.bonus.v1.GetBonusHistoryResponse> getGetBonusHistoryMethod;
    if ((getGetBonusHistoryMethod = BonusServiceGrpc.getGetBonusHistoryMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetBonusHistoryMethod = BonusServiceGrpc.getGetBonusHistoryMethod) == null) {
          BonusServiceGrpc.getGetBonusHistoryMethod = getGetBonusHistoryMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetBonusHistoryRequest, com.game_engine.bonus.v1.GetBonusHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBonusHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusHistoryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetBonusHistory"))
              .build();
        }
      }
    }
    return getGetBonusHistoryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusClaimsRequest,
      com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> getGetActiveBonusClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetActiveBonusClaims",
      requestType = com.game_engine.bonus.v1.GetActiveBonusClaimsRequest.class,
      responseType = com.game_engine.bonus.v1.GetActiveBonusClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusClaimsRequest,
      com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> getGetActiveBonusClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetActiveBonusClaimsRequest, com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> getGetActiveBonusClaimsMethod;
    if ((getGetActiveBonusClaimsMethod = BonusServiceGrpc.getGetActiveBonusClaimsMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetActiveBonusClaimsMethod = BonusServiceGrpc.getGetActiveBonusClaimsMethod) == null) {
          BonusServiceGrpc.getGetActiveBonusClaimsMethod = getGetActiveBonusClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetActiveBonusClaimsRequest, com.game_engine.bonus.v1.GetActiveBonusClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveBonusClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetActiveBonusClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetActiveBonusClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetActiveBonusClaims"))
              .build();
        }
      }
    }
    return getGetActiveBonusClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ProcessWageringContributionRequest,
      com.game_engine.bonus.v1.ProcessWageringContributionResponse> getProcessWageringContributionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ProcessWageringContribution",
      requestType = com.game_engine.bonus.v1.ProcessWageringContributionRequest.class,
      responseType = com.game_engine.bonus.v1.ProcessWageringContributionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ProcessWageringContributionRequest,
      com.game_engine.bonus.v1.ProcessWageringContributionResponse> getProcessWageringContributionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ProcessWageringContributionRequest, com.game_engine.bonus.v1.ProcessWageringContributionResponse> getProcessWageringContributionMethod;
    if ((getProcessWageringContributionMethod = BonusServiceGrpc.getProcessWageringContributionMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getProcessWageringContributionMethod = BonusServiceGrpc.getProcessWageringContributionMethod) == null) {
          BonusServiceGrpc.getProcessWageringContributionMethod = getProcessWageringContributionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.ProcessWageringContributionRequest, com.game_engine.bonus.v1.ProcessWageringContributionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ProcessWageringContribution"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ProcessWageringContributionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ProcessWageringContributionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("ProcessWageringContribution"))
              .build();
        }
      }
    }
    return getProcessWageringContributionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CompleteBonusRequest,
      com.game_engine.bonus.v1.CompleteBonusResponse> getCompleteBonusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CompleteBonus",
      requestType = com.game_engine.bonus.v1.CompleteBonusRequest.class,
      responseType = com.game_engine.bonus.v1.CompleteBonusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CompleteBonusRequest,
      com.game_engine.bonus.v1.CompleteBonusResponse> getCompleteBonusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CompleteBonusRequest, com.game_engine.bonus.v1.CompleteBonusResponse> getCompleteBonusMethod;
    if ((getCompleteBonusMethod = BonusServiceGrpc.getCompleteBonusMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getCompleteBonusMethod = BonusServiceGrpc.getCompleteBonusMethod) == null) {
          BonusServiceGrpc.getCompleteBonusMethod = getCompleteBonusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.CompleteBonusRequest, com.game_engine.bonus.v1.CompleteBonusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CompleteBonus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CompleteBonusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CompleteBonusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("CompleteBonus"))
              .build();
        }
      }
    }
    return getCompleteBonusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CancelBonusRequest,
      com.game_engine.bonus.v1.CancelBonusResponse> getCancelBonusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CancelBonus",
      requestType = com.game_engine.bonus.v1.CancelBonusRequest.class,
      responseType = com.game_engine.bonus.v1.CancelBonusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CancelBonusRequest,
      com.game_engine.bonus.v1.CancelBonusResponse> getCancelBonusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CancelBonusRequest, com.game_engine.bonus.v1.CancelBonusResponse> getCancelBonusMethod;
    if ((getCancelBonusMethod = BonusServiceGrpc.getCancelBonusMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getCancelBonusMethod = BonusServiceGrpc.getCancelBonusMethod) == null) {
          BonusServiceGrpc.getCancelBonusMethod = getCancelBonusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.CancelBonusRequest, com.game_engine.bonus.v1.CancelBonusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CancelBonus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CancelBonusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CancelBonusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("CancelBonus"))
              .build();
        }
      }
    }
    return getCancelBonusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusStatsRequest,
      com.game_engine.bonus.v1.GetBonusStatsResponse> getGetBonusStatsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBonusStats",
      requestType = com.game_engine.bonus.v1.GetBonusStatsRequest.class,
      responseType = com.game_engine.bonus.v1.GetBonusStatsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusStatsRequest,
      com.game_engine.bonus.v1.GetBonusStatsResponse> getGetBonusStatsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetBonusStatsRequest, com.game_engine.bonus.v1.GetBonusStatsResponse> getGetBonusStatsMethod;
    if ((getGetBonusStatsMethod = BonusServiceGrpc.getGetBonusStatsMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetBonusStatsMethod = BonusServiceGrpc.getGetBonusStatsMethod) == null) {
          BonusServiceGrpc.getGetBonusStatsMethod = getGetBonusStatsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetBonusStatsRequest, com.game_engine.bonus.v1.GetBonusStatsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBonusStats"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusStatsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetBonusStatsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetBonusStats"))
              .build();
        }
      }
    }
    return getGetBonusStatsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BonusServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BonusServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BonusServiceStub>() {
        @java.lang.Override
        public BonusServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BonusServiceStub(channel, callOptions);
        }
      };
    return BonusServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static BonusServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BonusServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BonusServiceBlockingV2Stub>() {
        @java.lang.Override
        public BonusServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BonusServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return BonusServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BonusServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BonusServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BonusServiceBlockingStub>() {
        @java.lang.Override
        public BonusServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BonusServiceBlockingStub(channel, callOptions);
        }
      };
    return BonusServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BonusServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BonusServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BonusServiceFutureStub>() {
        @java.lang.Override
        public BonusServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BonusServiceFutureStub(channel, callOptions);
        }
      };
    return BonusServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Bonus management
     * </pre>
     */
    default void getActiveBonuses(com.game_engine.bonus.v1.GetActiveBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveBonusesMethod(), responseObserver);
    }

    /**
     */
    default void getBonus(com.game_engine.bonus.v1.GetBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBonusMethod(), responseObserver);
    }

    /**
     */
    default void createBonus(com.game_engine.bonus.v1.CreateBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateBonusMethod(), responseObserver);
    }

    /**
     */
    default void claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimBonusMethod(), responseObserver);
    }

    /**
     */
    default void checkEligibility(com.game_engine.bonus.v1.CheckEligibilityRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CheckEligibilityResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCheckEligibilityMethod(), responseObserver);
    }

    /**
     */
    default void getBonusHistory(com.game_engine.bonus.v1.GetBonusHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBonusHistoryMethod(), responseObserver);
    }

    /**
     */
    default void getActiveBonusClaims(com.game_engine.bonus.v1.GetActiveBonusClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveBonusClaimsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Wagering
     * </pre>
     */
    default void processWageringContribution(com.game_engine.bonus.v1.ProcessWageringContributionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ProcessWageringContributionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getProcessWageringContributionMethod(), responseObserver);
    }

    /**
     */
    default void completeBonus(com.game_engine.bonus.v1.CompleteBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CompleteBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCompleteBonusMethod(), responseObserver);
    }

    /**
     */
    default void cancelBonus(com.game_engine.bonus.v1.CancelBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CancelBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCancelBonusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    default void getBonusStats(com.game_engine.bonus.v1.GetBonusStatsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusStatsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBonusStatsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service BonusService.
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public static abstract class BonusServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return BonusServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service BonusService.
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public static final class BonusServiceStub
      extends io.grpc.stub.AbstractAsyncStub<BonusServiceStub> {
    private BonusServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BonusServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BonusServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Bonus management
     * </pre>
     */
    public void getActiveBonuses(com.game_engine.bonus.v1.GetActiveBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveBonusesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getBonus(com.game_engine.bonus.v1.GetBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createBonus(com.game_engine.bonus.v1.CreateBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void checkEligibility(com.game_engine.bonus.v1.CheckEligibilityRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CheckEligibilityResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCheckEligibilityMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getBonusHistory(com.game_engine.bonus.v1.GetBonusHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBonusHistoryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getActiveBonusClaims(com.game_engine.bonus.v1.GetActiveBonusClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveBonusClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Wagering
     * </pre>
     */
    public void processWageringContribution(com.game_engine.bonus.v1.ProcessWageringContributionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ProcessWageringContributionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getProcessWageringContributionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void completeBonus(com.game_engine.bonus.v1.CompleteBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CompleteBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCompleteBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void cancelBonus(com.game_engine.bonus.v1.CancelBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CancelBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCancelBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public void getBonusStats(com.game_engine.bonus.v1.GetBonusStatsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusStatsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBonusStatsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service BonusService.
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public static final class BonusServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<BonusServiceBlockingV2Stub> {
    private BonusServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BonusServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BonusServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Bonus management
     * </pre>
     */
    public com.game_engine.bonus.v1.GetActiveBonusesResponse getActiveBonuses(com.game_engine.bonus.v1.GetActiveBonusesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetActiveBonusesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetBonusResponse getBonus(com.game_engine.bonus.v1.GetBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CreateBonusResponse createBonus(com.game_engine.bonus.v1.CreateBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.ClaimBonusResponse claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CheckEligibilityResponse checkEligibility(com.game_engine.bonus.v1.CheckEligibilityRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCheckEligibilityMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetBonusHistoryResponse getBonusHistory(com.game_engine.bonus.v1.GetBonusHistoryRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBonusHistoryMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetActiveBonusClaimsResponse getActiveBonusClaims(com.game_engine.bonus.v1.GetActiveBonusClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetActiveBonusClaimsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Wagering
     * </pre>
     */
    public com.game_engine.bonus.v1.ProcessWageringContributionResponse processWageringContribution(com.game_engine.bonus.v1.ProcessWageringContributionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getProcessWageringContributionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CompleteBonusResponse completeBonus(com.game_engine.bonus.v1.CompleteBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCompleteBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CancelBonusResponse cancelBonus(com.game_engine.bonus.v1.CancelBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCancelBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.game_engine.bonus.v1.GetBonusStatsResponse getBonusStats(com.game_engine.bonus.v1.GetBonusStatsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBonusStatsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service BonusService.
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public static final class BonusServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<BonusServiceBlockingStub> {
    private BonusServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BonusServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BonusServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Bonus management
     * </pre>
     */
    public com.game_engine.bonus.v1.GetActiveBonusesResponse getActiveBonuses(com.game_engine.bonus.v1.GetActiveBonusesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveBonusesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetBonusResponse getBonus(com.game_engine.bonus.v1.GetBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CreateBonusResponse createBonus(com.game_engine.bonus.v1.CreateBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.ClaimBonusResponse claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CheckEligibilityResponse checkEligibility(com.game_engine.bonus.v1.CheckEligibilityRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCheckEligibilityMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetBonusHistoryResponse getBonusHistory(com.game_engine.bonus.v1.GetBonusHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBonusHistoryMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetActiveBonusClaimsResponse getActiveBonusClaims(com.game_engine.bonus.v1.GetActiveBonusClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveBonusClaimsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Wagering
     * </pre>
     */
    public com.game_engine.bonus.v1.ProcessWageringContributionResponse processWageringContribution(com.game_engine.bonus.v1.ProcessWageringContributionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getProcessWageringContributionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CompleteBonusResponse completeBonus(com.game_engine.bonus.v1.CompleteBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCompleteBonusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.CancelBonusResponse cancelBonus(com.game_engine.bonus.v1.CancelBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCancelBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.game_engine.bonus.v1.GetBonusStatsResponse getBonusStats(com.game_engine.bonus.v1.GetBonusStatsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBonusStatsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service BonusService.
   * <pre>
   * Bonus Service - manages bonuses, claims, wagering requirements
   * </pre>
   */
  public static final class BonusServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<BonusServiceFutureStub> {
    private BonusServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BonusServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BonusServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Bonus management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetActiveBonusesResponse> getActiveBonuses(
        com.game_engine.bonus.v1.GetActiveBonusesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveBonusesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetBonusResponse> getBonus(
        com.game_engine.bonus.v1.GetBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBonusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.CreateBonusResponse> createBonus(
        com.game_engine.bonus.v1.CreateBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateBonusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.ClaimBonusResponse> claimBonus(
        com.game_engine.bonus.v1.ClaimBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimBonusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.CheckEligibilityResponse> checkEligibility(
        com.game_engine.bonus.v1.CheckEligibilityRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCheckEligibilityMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetBonusHistoryResponse> getBonusHistory(
        com.game_engine.bonus.v1.GetBonusHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBonusHistoryMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetActiveBonusClaimsResponse> getActiveBonusClaims(
        com.game_engine.bonus.v1.GetActiveBonusClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveBonusClaimsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Wagering
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.ProcessWageringContributionResponse> processWageringContribution(
        com.game_engine.bonus.v1.ProcessWageringContributionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getProcessWageringContributionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.CompleteBonusResponse> completeBonus(
        com.game_engine.bonus.v1.CompleteBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCompleteBonusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.CancelBonusResponse> cancelBonus(
        com.game_engine.bonus.v1.CancelBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCancelBonusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Stats
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetBonusStatsResponse> getBonusStats(
        com.game_engine.bonus.v1.GetBonusStatsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBonusStatsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_ACTIVE_BONUSES = 0;
  private static final int METHODID_GET_BONUS = 1;
  private static final int METHODID_CREATE_BONUS = 2;
  private static final int METHODID_CLAIM_BONUS = 3;
  private static final int METHODID_CHECK_ELIGIBILITY = 4;
  private static final int METHODID_GET_BONUS_HISTORY = 5;
  private static final int METHODID_GET_ACTIVE_BONUS_CLAIMS = 6;
  private static final int METHODID_PROCESS_WAGERING_CONTRIBUTION = 7;
  private static final int METHODID_COMPLETE_BONUS = 8;
  private static final int METHODID_CANCEL_BONUS = 9;
  private static final int METHODID_GET_BONUS_STATS = 10;

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
        case METHODID_GET_ACTIVE_BONUSES:
          serviceImpl.getActiveBonuses((com.game_engine.bonus.v1.GetActiveBonusesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusesResponse>) responseObserver);
          break;
        case METHODID_GET_BONUS:
          serviceImpl.getBonus((com.game_engine.bonus.v1.GetBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse>) responseObserver);
          break;
        case METHODID_CREATE_BONUS:
          serviceImpl.createBonus((com.game_engine.bonus.v1.CreateBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateBonusResponse>) responseObserver);
          break;
        case METHODID_CLAIM_BONUS:
          serviceImpl.claimBonus((com.game_engine.bonus.v1.ClaimBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse>) responseObserver);
          break;
        case METHODID_CHECK_ELIGIBILITY:
          serviceImpl.checkEligibility((com.game_engine.bonus.v1.CheckEligibilityRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CheckEligibilityResponse>) responseObserver);
          break;
        case METHODID_GET_BONUS_HISTORY:
          serviceImpl.getBonusHistory((com.game_engine.bonus.v1.GetBonusHistoryRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusHistoryResponse>) responseObserver);
          break;
        case METHODID_GET_ACTIVE_BONUS_CLAIMS:
          serviceImpl.getActiveBonusClaims((com.game_engine.bonus.v1.GetActiveBonusClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetActiveBonusClaimsResponse>) responseObserver);
          break;
        case METHODID_PROCESS_WAGERING_CONTRIBUTION:
          serviceImpl.processWageringContribution((com.game_engine.bonus.v1.ProcessWageringContributionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ProcessWageringContributionResponse>) responseObserver);
          break;
        case METHODID_COMPLETE_BONUS:
          serviceImpl.completeBonus((com.game_engine.bonus.v1.CompleteBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CompleteBonusResponse>) responseObserver);
          break;
        case METHODID_CANCEL_BONUS:
          serviceImpl.cancelBonus((com.game_engine.bonus.v1.CancelBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CancelBonusResponse>) responseObserver);
          break;
        case METHODID_GET_BONUS_STATS:
          serviceImpl.getBonusStats((com.game_engine.bonus.v1.GetBonusStatsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusStatsResponse>) responseObserver);
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
          getGetActiveBonusesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetActiveBonusesRequest,
              com.game_engine.bonus.v1.GetActiveBonusesResponse>(
                service, METHODID_GET_ACTIVE_BONUSES)))
        .addMethod(
          getGetBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetBonusRequest,
              com.game_engine.bonus.v1.GetBonusResponse>(
                service, METHODID_GET_BONUS)))
        .addMethod(
          getCreateBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.CreateBonusRequest,
              com.game_engine.bonus.v1.CreateBonusResponse>(
                service, METHODID_CREATE_BONUS)))
        .addMethod(
          getClaimBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.ClaimBonusRequest,
              com.game_engine.bonus.v1.ClaimBonusResponse>(
                service, METHODID_CLAIM_BONUS)))
        .addMethod(
          getCheckEligibilityMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.CheckEligibilityRequest,
              com.game_engine.bonus.v1.CheckEligibilityResponse>(
                service, METHODID_CHECK_ELIGIBILITY)))
        .addMethod(
          getGetBonusHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetBonusHistoryRequest,
              com.game_engine.bonus.v1.GetBonusHistoryResponse>(
                service, METHODID_GET_BONUS_HISTORY)))
        .addMethod(
          getGetActiveBonusClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetActiveBonusClaimsRequest,
              com.game_engine.bonus.v1.GetActiveBonusClaimsResponse>(
                service, METHODID_GET_ACTIVE_BONUS_CLAIMS)))
        .addMethod(
          getProcessWageringContributionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.ProcessWageringContributionRequest,
              com.game_engine.bonus.v1.ProcessWageringContributionResponse>(
                service, METHODID_PROCESS_WAGERING_CONTRIBUTION)))
        .addMethod(
          getCompleteBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.CompleteBonusRequest,
              com.game_engine.bonus.v1.CompleteBonusResponse>(
                service, METHODID_COMPLETE_BONUS)))
        .addMethod(
          getCancelBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.CancelBonusRequest,
              com.game_engine.bonus.v1.CancelBonusResponse>(
                service, METHODID_CANCEL_BONUS)))
        .addMethod(
          getGetBonusStatsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetBonusStatsRequest,
              com.game_engine.bonus.v1.GetBonusStatsResponse>(
                service, METHODID_GET_BONUS_STATS)))
        .build();
  }

  private static abstract class BonusServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BonusServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.bonus.v1.BonusServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("BonusService");
    }
  }

  private static final class BonusServiceFileDescriptorSupplier
      extends BonusServiceBaseDescriptorSupplier {
    BonusServiceFileDescriptorSupplier() {}
  }

  private static final class BonusServiceMethodDescriptorSupplier
      extends BonusServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    BonusServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (BonusServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BonusServiceFileDescriptorSupplier())
              .addMethod(getGetActiveBonusesMethod())
              .addMethod(getGetBonusMethod())
              .addMethod(getCreateBonusMethod())
              .addMethod(getClaimBonusMethod())
              .addMethod(getCheckEligibilityMethod())
              .addMethod(getGetBonusHistoryMethod())
              .addMethod(getGetActiveBonusClaimsMethod())
              .addMethod(getProcessWageringContributionMethod())
              .addMethod(getCompleteBonusMethod())
              .addMethod(getCancelBonusMethod())
              .addMethod(getGetBonusStatsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
