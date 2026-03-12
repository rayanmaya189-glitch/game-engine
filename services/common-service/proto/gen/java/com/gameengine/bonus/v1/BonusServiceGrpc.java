package com.game_engine.bonus.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class BonusServiceGrpc {

  private BonusServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.bonus.v1.BonusService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ListBonusesRequest,
      com.game_engine.bonus.v1.ListBonusesResponse> getListBonusesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListBonuses",
      requestType = com.game_engine.bonus.v1.ListBonusesRequest.class,
      responseType = com.game_engine.bonus.v1.ListBonusesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ListBonusesRequest,
      com.game_engine.bonus.v1.ListBonusesResponse> getListBonusesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ListBonusesRequest, com.game_engine.bonus.v1.ListBonusesResponse> getListBonusesMethod;
    if ((getListBonusesMethod = BonusServiceGrpc.getListBonusesMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getListBonusesMethod = BonusServiceGrpc.getListBonusesMethod) == null) {
          BonusServiceGrpc.getListBonusesMethod = getListBonusesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.ListBonusesRequest, com.game_engine.bonus.v1.ListBonusesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListBonuses"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ListBonusesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ListBonusesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("ListBonuses"))
              .build();
        }
      }
    }
    return getListBonusesMethod;
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

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserBonusesRequest,
      com.game_engine.bonus.v1.GetUserBonusesResponse> getGetUserBonusesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserBonuses",
      requestType = com.game_engine.bonus.v1.GetUserBonusesRequest.class,
      responseType = com.game_engine.bonus.v1.GetUserBonusesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserBonusesRequest,
      com.game_engine.bonus.v1.GetUserBonusesResponse> getGetUserBonusesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserBonusesRequest, com.game_engine.bonus.v1.GetUserBonusesResponse> getGetUserBonusesMethod;
    if ((getGetUserBonusesMethod = BonusServiceGrpc.getGetUserBonusesMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetUserBonusesMethod = BonusServiceGrpc.getGetUserBonusesMethod) == null) {
          BonusServiceGrpc.getGetUserBonusesMethod = getGetUserBonusesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetUserBonusesRequest, com.game_engine.bonus.v1.GetUserBonusesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserBonuses"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserBonusesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserBonusesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetUserBonuses"))
              .build();
        }
      }
    }
    return getGetUserBonusesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateRebetClaimRequest,
      com.game_engine.bonus.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateRebetClaim",
      requestType = com.game_engine.bonus.v1.CreateRebetClaimRequest.class,
      responseType = com.game_engine.bonus.v1.CreateRebetClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateRebetClaimRequest,
      com.game_engine.bonus.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.CreateRebetClaimRequest, com.game_engine.bonus.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod;
    if ((getCreateRebetClaimMethod = BonusServiceGrpc.getCreateRebetClaimMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getCreateRebetClaimMethod = BonusServiceGrpc.getCreateRebetClaimMethod) == null) {
          BonusServiceGrpc.getCreateRebetClaimMethod = getCreateRebetClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.CreateRebetClaimRequest, com.game_engine.bonus.v1.CreateRebetClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateRebetClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CreateRebetClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.CreateRebetClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("CreateRebetClaim"))
              .build();
        }
      }
    }
    return getCreateRebetClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserRebetClaimsRequest,
      com.game_engine.bonus.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserRebetClaims",
      requestType = com.game_engine.bonus.v1.GetUserRebetClaimsRequest.class,
      responseType = com.game_engine.bonus.v1.GetUserRebetClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserRebetClaimsRequest,
      com.game_engine.bonus.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserRebetClaimsRequest, com.game_engine.bonus.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod;
    if ((getGetUserRebetClaimsMethod = BonusServiceGrpc.getGetUserRebetClaimsMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetUserRebetClaimsMethod = BonusServiceGrpc.getGetUserRebetClaimsMethod) == null) {
          BonusServiceGrpc.getGetUserRebetClaimsMethod = getGetUserRebetClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetUserRebetClaimsRequest, com.game_engine.bonus.v1.GetUserRebetClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserRebetClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserRebetClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserRebetClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetUserRebetClaims"))
              .build();
        }
      }
    }
    return getGetUserRebetClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetClaimableRebetsRequest,
      com.game_engine.bonus.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClaimableRebets",
      requestType = com.game_engine.bonus.v1.GetClaimableRebetsRequest.class,
      responseType = com.game_engine.bonus.v1.GetClaimableRebetsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetClaimableRebetsRequest,
      com.game_engine.bonus.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetClaimableRebetsRequest, com.game_engine.bonus.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod;
    if ((getGetClaimableRebetsMethod = BonusServiceGrpc.getGetClaimableRebetsMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetClaimableRebetsMethod = BonusServiceGrpc.getGetClaimableRebetsMethod) == null) {
          BonusServiceGrpc.getGetClaimableRebetsMethod = getGetClaimableRebetsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetClaimableRebetsRequest, com.game_engine.bonus.v1.GetClaimableRebetsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClaimableRebets"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetClaimableRebetsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetClaimableRebetsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetClaimableRebets"))
              .build();
        }
      }
    }
    return getGetClaimableRebetsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimRebetRequest,
      com.game_engine.bonus.v1.ClaimRebetResponse> getClaimRebetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClaimRebet",
      requestType = com.game_engine.bonus.v1.ClaimRebetRequest.class,
      responseType = com.game_engine.bonus.v1.ClaimRebetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimRebetRequest,
      com.game_engine.bonus.v1.ClaimRebetResponse> getClaimRebetMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.ClaimRebetRequest, com.game_engine.bonus.v1.ClaimRebetResponse> getClaimRebetMethod;
    if ((getClaimRebetMethod = BonusServiceGrpc.getClaimRebetMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getClaimRebetMethod = BonusServiceGrpc.getClaimRebetMethod) == null) {
          BonusServiceGrpc.getClaimRebetMethod = getClaimRebetMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.ClaimRebetRequest, com.game_engine.bonus.v1.ClaimRebetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClaimRebet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ClaimRebetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.ClaimRebetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("ClaimRebet"))
              .build();
        }
      }
    }
    return getClaimRebetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.SubmitInsuranceClaimRequest,
      com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitInsuranceClaim",
      requestType = com.game_engine.bonus.v1.SubmitInsuranceClaimRequest.class,
      responseType = com.game_engine.bonus.v1.SubmitInsuranceClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.SubmitInsuranceClaimRequest,
      com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.SubmitInsuranceClaimRequest, com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod;
    if ((getSubmitInsuranceClaimMethod = BonusServiceGrpc.getSubmitInsuranceClaimMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getSubmitInsuranceClaimMethod = BonusServiceGrpc.getSubmitInsuranceClaimMethod) == null) {
          BonusServiceGrpc.getSubmitInsuranceClaimMethod = getSubmitInsuranceClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.SubmitInsuranceClaimRequest, com.game_engine.bonus.v1.SubmitInsuranceClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitInsuranceClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.SubmitInsuranceClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.SubmitInsuranceClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("SubmitInsuranceClaim"))
              .build();
        }
      }
    }
    return getSubmitInsuranceClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest,
      com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserInsuranceClaims",
      requestType = com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest.class,
      responseType = com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest,
      com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest, com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod;
    if ((getGetUserInsuranceClaimsMethod = BonusServiceGrpc.getGetUserInsuranceClaimsMethod) == null) {
      synchronized (BonusServiceGrpc.class) {
        if ((getGetUserInsuranceClaimsMethod = BonusServiceGrpc.getGetUserInsuranceClaimsMethod) == null) {
          BonusServiceGrpc.getGetUserInsuranceClaimsMethod = getGetUserInsuranceClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest, com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserInsuranceClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BonusServiceMethodDescriptorSupplier("GetUserInsuranceClaims"))
              .build();
        }
      }
    }
    return getGetUserInsuranceClaimsMethod;
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
   */
  public interface AsyncService {

    /**
     * <pre>
     * List available bonuses
     * </pre>
     */
    default void listBonuses(com.game_engine.bonus.v1.ListBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ListBonusesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListBonusesMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get bonus details
     * </pre>
     */
    default void getBonus(com.game_engine.bonus.v1.GetBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBonusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Claim a bonus
     * </pre>
     */
    default void claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimBonusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get user's claimed bonuses
     * </pre>
     */
    default void getUserBonuses(com.game_engine.bonus.v1.GetUserBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserBonusesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserBonusesMethod(), responseObserver);
    }

    /**
     * <pre>
     * Rebet claim operations
     * </pre>
     */
    default void createRebetClaim(com.game_engine.bonus.v1.CreateRebetClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateRebetClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateRebetClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserRebetClaims(com.game_engine.bonus.v1.GetUserRebetClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserRebetClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserRebetClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getClaimableRebets(com.game_engine.bonus.v1.GetClaimableRebetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetClaimableRebetsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClaimableRebetsMethod(), responseObserver);
    }

    /**
     */
    default void claimRebet(com.game_engine.bonus.v1.ClaimRebetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimRebetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimRebetMethod(), responseObserver);
    }

    /**
     * <pre>
     * Insurance claim operations
     * </pre>
     */
    default void submitInsuranceClaim(com.game_engine.bonus.v1.SubmitInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitInsuranceClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserInsuranceClaims(com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserInsuranceClaimsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service BonusService.
   */
  public static abstract class BonusServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return BonusServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service BonusService.
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
     * List available bonuses
     * </pre>
     */
    public void listBonuses(com.game_engine.bonus.v1.ListBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ListBonusesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListBonusesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get bonus details
     * </pre>
     */
    public void getBonus(com.game_engine.bonus.v1.GetBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Claim a bonus
     * </pre>
     */
    public void claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimBonusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get user's claimed bonuses
     * </pre>
     */
    public void getUserBonuses(com.game_engine.bonus.v1.GetUserBonusesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserBonusesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserBonusesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Rebet claim operations
     * </pre>
     */
    public void createRebetClaim(com.game_engine.bonus.v1.CreateRebetClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateRebetClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateRebetClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserRebetClaims(com.game_engine.bonus.v1.GetUserRebetClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserRebetClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserRebetClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClaimableRebets(com.game_engine.bonus.v1.GetClaimableRebetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetClaimableRebetsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClaimableRebetsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void claimRebet(com.game_engine.bonus.v1.ClaimRebetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimRebetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimRebetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Insurance claim operations
     * </pre>
     */
    public void submitInsuranceClaim(com.game_engine.bonus.v1.SubmitInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitInsuranceClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserInsuranceClaims(com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserInsuranceClaimsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service BonusService.
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
     * List available bonuses
     * </pre>
     */
    public com.game_engine.bonus.v1.ListBonusesResponse listBonuses(com.game_engine.bonus.v1.ListBonusesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListBonusesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get bonus details
     * </pre>
     */
    public com.game_engine.bonus.v1.GetBonusResponse getBonus(com.game_engine.bonus.v1.GetBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Claim a bonus
     * </pre>
     */
    public com.game_engine.bonus.v1.ClaimBonusResponse claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get user's claimed bonuses
     * </pre>
     */
    public com.game_engine.bonus.v1.GetUserBonusesResponse getUserBonuses(com.game_engine.bonus.v1.GetUserBonusesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserBonusesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Rebet claim operations
     * </pre>
     */
    public com.game_engine.bonus.v1.CreateRebetClaimResponse createRebetClaim(com.game_engine.bonus.v1.CreateRebetClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateRebetClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetUserRebetClaimsResponse getUserRebetClaims(com.game_engine.bonus.v1.GetUserRebetClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserRebetClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetClaimableRebetsResponse getClaimableRebets(com.game_engine.bonus.v1.GetClaimableRebetsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClaimableRebetsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.ClaimRebetResponse claimRebet(com.game_engine.bonus.v1.ClaimRebetRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimRebetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Insurance claim operations
     * </pre>
     */
    public com.game_engine.bonus.v1.SubmitInsuranceClaimResponse submitInsuranceClaim(com.game_engine.bonus.v1.SubmitInsuranceClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse getUserInsuranceClaims(com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserInsuranceClaimsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service BonusService.
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
     * List available bonuses
     * </pre>
     */
    public com.game_engine.bonus.v1.ListBonusesResponse listBonuses(com.game_engine.bonus.v1.ListBonusesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListBonusesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get bonus details
     * </pre>
     */
    public com.game_engine.bonus.v1.GetBonusResponse getBonus(com.game_engine.bonus.v1.GetBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Claim a bonus
     * </pre>
     */
    public com.game_engine.bonus.v1.ClaimBonusResponse claimBonus(com.game_engine.bonus.v1.ClaimBonusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimBonusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get user's claimed bonuses
     * </pre>
     */
    public com.game_engine.bonus.v1.GetUserBonusesResponse getUserBonuses(com.game_engine.bonus.v1.GetUserBonusesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserBonusesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Rebet claim operations
     * </pre>
     */
    public com.game_engine.bonus.v1.CreateRebetClaimResponse createRebetClaim(com.game_engine.bonus.v1.CreateRebetClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateRebetClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetUserRebetClaimsResponse getUserRebetClaims(com.game_engine.bonus.v1.GetUserRebetClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserRebetClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetClaimableRebetsResponse getClaimableRebets(com.game_engine.bonus.v1.GetClaimableRebetsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClaimableRebetsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.ClaimRebetResponse claimRebet(com.game_engine.bonus.v1.ClaimRebetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimRebetMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Insurance claim operations
     * </pre>
     */
    public com.game_engine.bonus.v1.SubmitInsuranceClaimResponse submitInsuranceClaim(com.game_engine.bonus.v1.SubmitInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse getUserInsuranceClaims(com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserInsuranceClaimsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service BonusService.
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
     * List available bonuses
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.ListBonusesResponse> listBonuses(
        com.game_engine.bonus.v1.ListBonusesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListBonusesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get bonus details
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetBonusResponse> getBonus(
        com.game_engine.bonus.v1.GetBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBonusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Claim a bonus
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.ClaimBonusResponse> claimBonus(
        com.game_engine.bonus.v1.ClaimBonusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimBonusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get user's claimed bonuses
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetUserBonusesResponse> getUserBonuses(
        com.game_engine.bonus.v1.GetUserBonusesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserBonusesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Rebet claim operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.CreateRebetClaimResponse> createRebetClaim(
        com.game_engine.bonus.v1.CreateRebetClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateRebetClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetUserRebetClaimsResponse> getUserRebetClaims(
        com.game_engine.bonus.v1.GetUserRebetClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserRebetClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetClaimableRebetsResponse> getClaimableRebets(
        com.game_engine.bonus.v1.GetClaimableRebetsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClaimableRebetsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.ClaimRebetResponse> claimRebet(
        com.game_engine.bonus.v1.ClaimRebetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimRebetMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Insurance claim operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.SubmitInsuranceClaimResponse> submitInsuranceClaim(
        com.game_engine.bonus.v1.SubmitInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitInsuranceClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse> getUserInsuranceClaims(
        com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserInsuranceClaimsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_BONUSES = 0;
  private static final int METHODID_GET_BONUS = 1;
  private static final int METHODID_CLAIM_BONUS = 2;
  private static final int METHODID_GET_USER_BONUSES = 3;
  private static final int METHODID_CREATE_REBET_CLAIM = 4;
  private static final int METHODID_GET_USER_REBET_CLAIMS = 5;
  private static final int METHODID_GET_CLAIMABLE_REBETS = 6;
  private static final int METHODID_CLAIM_REBET = 7;
  private static final int METHODID_SUBMIT_INSURANCE_CLAIM = 8;
  private static final int METHODID_GET_USER_INSURANCE_CLAIMS = 9;

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
        case METHODID_LIST_BONUSES:
          serviceImpl.listBonuses((com.game_engine.bonus.v1.ListBonusesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ListBonusesResponse>) responseObserver);
          break;
        case METHODID_GET_BONUS:
          serviceImpl.getBonus((com.game_engine.bonus.v1.GetBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetBonusResponse>) responseObserver);
          break;
        case METHODID_CLAIM_BONUS:
          serviceImpl.claimBonus((com.game_engine.bonus.v1.ClaimBonusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimBonusResponse>) responseObserver);
          break;
        case METHODID_GET_USER_BONUSES:
          serviceImpl.getUserBonuses((com.game_engine.bonus.v1.GetUserBonusesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserBonusesResponse>) responseObserver);
          break;
        case METHODID_CREATE_REBET_CLAIM:
          serviceImpl.createRebetClaim((com.game_engine.bonus.v1.CreateRebetClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.CreateRebetClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_REBET_CLAIMS:
          serviceImpl.getUserRebetClaims((com.game_engine.bonus.v1.GetUserRebetClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserRebetClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_CLAIMABLE_REBETS:
          serviceImpl.getClaimableRebets((com.game_engine.bonus.v1.GetClaimableRebetsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetClaimableRebetsResponse>) responseObserver);
          break;
        case METHODID_CLAIM_REBET:
          serviceImpl.claimRebet((com.game_engine.bonus.v1.ClaimRebetRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.ClaimRebetResponse>) responseObserver);
          break;
        case METHODID_SUBMIT_INSURANCE_CLAIM:
          serviceImpl.submitInsuranceClaim((com.game_engine.bonus.v1.SubmitInsuranceClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.SubmitInsuranceClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_INSURANCE_CLAIMS:
          serviceImpl.getUserInsuranceClaims((com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse>) responseObserver);
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
          getListBonusesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.ListBonusesRequest,
              com.game_engine.bonus.v1.ListBonusesResponse>(
                service, METHODID_LIST_BONUSES)))
        .addMethod(
          getGetBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetBonusRequest,
              com.game_engine.bonus.v1.GetBonusResponse>(
                service, METHODID_GET_BONUS)))
        .addMethod(
          getClaimBonusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.ClaimBonusRequest,
              com.game_engine.bonus.v1.ClaimBonusResponse>(
                service, METHODID_CLAIM_BONUS)))
        .addMethod(
          getGetUserBonusesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetUserBonusesRequest,
              com.game_engine.bonus.v1.GetUserBonusesResponse>(
                service, METHODID_GET_USER_BONUSES)))
        .addMethod(
          getCreateRebetClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.CreateRebetClaimRequest,
              com.game_engine.bonus.v1.CreateRebetClaimResponse>(
                service, METHODID_CREATE_REBET_CLAIM)))
        .addMethod(
          getGetUserRebetClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetUserRebetClaimsRequest,
              com.game_engine.bonus.v1.GetUserRebetClaimsResponse>(
                service, METHODID_GET_USER_REBET_CLAIMS)))
        .addMethod(
          getGetClaimableRebetsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetClaimableRebetsRequest,
              com.game_engine.bonus.v1.GetClaimableRebetsResponse>(
                service, METHODID_GET_CLAIMABLE_REBETS)))
        .addMethod(
          getClaimRebetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.ClaimRebetRequest,
              com.game_engine.bonus.v1.ClaimRebetResponse>(
                service, METHODID_CLAIM_REBET)))
        .addMethod(
          getSubmitInsuranceClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.SubmitInsuranceClaimRequest,
              com.game_engine.bonus.v1.SubmitInsuranceClaimResponse>(
                service, METHODID_SUBMIT_INSURANCE_CLAIM)))
        .addMethod(
          getGetUserInsuranceClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.bonus.v1.GetUserInsuranceClaimsRequest,
              com.game_engine.bonus.v1.GetUserInsuranceClaimsResponse>(
                service, METHODID_GET_USER_INSURANCE_CLAIMS)))
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
              .addMethod(getListBonusesMethod())
              .addMethod(getGetBonusMethod())
              .addMethod(getClaimBonusMethod())
              .addMethod(getGetUserBonusesMethod())
              .addMethod(getCreateRebetClaimMethod())
              .addMethod(getGetUserRebetClaimsMethod())
              .addMethod(getGetClaimableRebetsMethod())
              .addMethod(getClaimRebetMethod())
              .addMethod(getSubmitInsuranceClaimMethod())
              .addMethod(getGetUserInsuranceClaimsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
