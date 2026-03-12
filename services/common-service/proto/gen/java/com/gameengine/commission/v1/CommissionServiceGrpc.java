package com.gameengine.commission.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class CommissionServiceGrpc {

  private CommissionServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "gameengine.commission.v1.CommissionService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.SubmitClaimRequest,
      com.gameengine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitClaim",
      requestType = com.gameengine.commission.v1.SubmitClaimRequest.class,
      responseType = com.gameengine.commission.v1.SubmitClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.SubmitClaimRequest,
      com.gameengine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.SubmitClaimRequest, com.gameengine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod;
    if ((getSubmitClaimMethod = CommissionServiceGrpc.getSubmitClaimMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getSubmitClaimMethod = CommissionServiceGrpc.getSubmitClaimMethod) == null) {
          CommissionServiceGrpc.getSubmitClaimMethod = getSubmitClaimMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.SubmitClaimRequest, com.gameengine.commission.v1.SubmitClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.SubmitClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.SubmitClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("SubmitClaim"))
              .build();
        }
      }
    }
    return getSubmitClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserClaimsRequest,
      com.gameengine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserClaims",
      requestType = com.gameengine.commission.v1.GetUserClaimsRequest.class,
      responseType = com.gameengine.commission.v1.GetUserClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserClaimsRequest,
      com.gameengine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserClaimsRequest, com.gameengine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod;
    if ((getGetUserClaimsMethod = CommissionServiceGrpc.getGetUserClaimsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetUserClaimsMethod = CommissionServiceGrpc.getGetUserClaimsMethod) == null) {
          CommissionServiceGrpc.getGetUserClaimsMethod = getGetUserClaimsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetUserClaimsRequest, com.gameengine.commission.v1.GetUserClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetUserClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetUserClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetUserClaims"))
              .build();
        }
      }
    }
    return getGetUserClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetClaimsByStatusRequest,
      com.gameengine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClaimsByStatus",
      requestType = com.gameengine.commission.v1.GetClaimsByStatusRequest.class,
      responseType = com.gameengine.commission.v1.GetClaimsByStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetClaimsByStatusRequest,
      com.gameengine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetClaimsByStatusRequest, com.gameengine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod;
    if ((getGetClaimsByStatusMethod = CommissionServiceGrpc.getGetClaimsByStatusMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetClaimsByStatusMethod = CommissionServiceGrpc.getGetClaimsByStatusMethod) == null) {
          CommissionServiceGrpc.getGetClaimsByStatusMethod = getGetClaimsByStatusMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetClaimsByStatusRequest, com.gameengine.commission.v1.GetClaimsByStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClaimsByStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetClaimsByStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetClaimsByStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetClaimsByStatus"))
              .build();
        }
      }
    }
    return getGetClaimsByStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.ClaimCommissionRequest,
      com.gameengine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClaimCommission",
      requestType = com.gameengine.commission.v1.ClaimCommissionRequest.class,
      responseType = com.gameengine.commission.v1.ClaimCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.ClaimCommissionRequest,
      com.gameengine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.ClaimCommissionRequest, com.gameengine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod;
    if ((getClaimCommissionMethod = CommissionServiceGrpc.getClaimCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getClaimCommissionMethod = CommissionServiceGrpc.getClaimCommissionMethod) == null) {
          CommissionServiceGrpc.getClaimCommissionMethod = getClaimCommissionMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.ClaimCommissionRequest, com.gameengine.commission.v1.ClaimCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClaimCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.ClaimCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.ClaimCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("ClaimCommission"))
              .build();
        }
      }
    }
    return getClaimCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserSettlementsRequest,
      com.gameengine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserSettlements",
      requestType = com.gameengine.commission.v1.GetUserSettlementsRequest.class,
      responseType = com.gameengine.commission.v1.GetUserSettlementsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserSettlementsRequest,
      com.gameengine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetUserSettlementsRequest, com.gameengine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod;
    if ((getGetUserSettlementsMethod = CommissionServiceGrpc.getGetUserSettlementsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetUserSettlementsMethod = CommissionServiceGrpc.getGetUserSettlementsMethod) == null) {
          CommissionServiceGrpc.getGetUserSettlementsMethod = getGetUserSettlementsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetUserSettlementsRequest, com.gameengine.commission.v1.GetUserSettlementsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserSettlements"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetUserSettlementsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetUserSettlementsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetUserSettlements"))
              .build();
        }
      }
    }
    return getGetUserSettlementsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetSettlementByIdRequest,
      com.gameengine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSettlementById",
      requestType = com.gameengine.commission.v1.GetSettlementByIdRequest.class,
      responseType = com.gameengine.commission.v1.GetSettlementByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetSettlementByIdRequest,
      com.gameengine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetSettlementByIdRequest, com.gameengine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod;
    if ((getGetSettlementByIdMethod = CommissionServiceGrpc.getGetSettlementByIdMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetSettlementByIdMethod = CommissionServiceGrpc.getGetSettlementByIdMethod) == null) {
          CommissionServiceGrpc.getGetSettlementByIdMethod = getGetSettlementByIdMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetSettlementByIdRequest, com.gameengine.commission.v1.GetSettlementByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSettlementById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetSettlementByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetSettlementByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetSettlementById"))
              .build();
        }
      }
    }
    return getGetSettlementByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalPendingRequest,
      com.gameengine.commission.v1.GetTotalPendingResponse> getGetTotalPendingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalPending",
      requestType = com.gameengine.commission.v1.GetTotalPendingRequest.class,
      responseType = com.gameengine.commission.v1.GetTotalPendingResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalPendingRequest,
      com.gameengine.commission.v1.GetTotalPendingResponse> getGetTotalPendingMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalPendingRequest, com.gameengine.commission.v1.GetTotalPendingResponse> getGetTotalPendingMethod;
    if ((getGetTotalPendingMethod = CommissionServiceGrpc.getGetTotalPendingMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetTotalPendingMethod = CommissionServiceGrpc.getGetTotalPendingMethod) == null) {
          CommissionServiceGrpc.getGetTotalPendingMethod = getGetTotalPendingMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetTotalPendingRequest, com.gameengine.commission.v1.GetTotalPendingResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalPending"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetTotalPendingRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetTotalPendingResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetTotalPending"))
              .build();
        }
      }
    }
    return getGetTotalPendingMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalSettledRequest,
      com.gameengine.commission.v1.GetTotalSettledResponse> getGetTotalSettledMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalSettled",
      requestType = com.gameengine.commission.v1.GetTotalSettledRequest.class,
      responseType = com.gameengine.commission.v1.GetTotalSettledResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalSettledRequest,
      com.gameengine.commission.v1.GetTotalSettledResponse> getGetTotalSettledMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetTotalSettledRequest, com.gameengine.commission.v1.GetTotalSettledResponse> getGetTotalSettledMethod;
    if ((getGetTotalSettledMethod = CommissionServiceGrpc.getGetTotalSettledMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetTotalSettledMethod = CommissionServiceGrpc.getGetTotalSettledMethod) == null) {
          CommissionServiceGrpc.getGetTotalSettledMethod = getGetTotalSettledMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetTotalSettledRequest, com.gameengine.commission.v1.GetTotalSettledResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalSettled"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetTotalSettledRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetTotalSettledResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetTotalSettled"))
              .build();
        }
      }
    }
    return getGetTotalSettledMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetAgentCommissionsRequest,
      com.gameengine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAgentCommissions",
      requestType = com.gameengine.commission.v1.GetAgentCommissionsRequest.class,
      responseType = com.gameengine.commission.v1.GetAgentCommissionsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetAgentCommissionsRequest,
      com.gameengine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetAgentCommissionsRequest, com.gameengine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod;
    if ((getGetAgentCommissionsMethod = CommissionServiceGrpc.getGetAgentCommissionsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetAgentCommissionsMethod = CommissionServiceGrpc.getGetAgentCommissionsMethod) == null) {
          CommissionServiceGrpc.getGetAgentCommissionsMethod = getGetAgentCommissionsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetAgentCommissionsRequest, com.gameengine.commission.v1.GetAgentCommissionsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAgentCommissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetAgentCommissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetAgentCommissionsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetAgentCommissions"))
              .build();
        }
      }
    }
    return getGetAgentCommissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetPendingCommissionsRequest,
      com.gameengine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPendingCommissions",
      requestType = com.gameengine.commission.v1.GetPendingCommissionsRequest.class,
      responseType = com.gameengine.commission.v1.GetPendingCommissionsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetPendingCommissionsRequest,
      com.gameengine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetPendingCommissionsRequest, com.gameengine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod;
    if ((getGetPendingCommissionsMethod = CommissionServiceGrpc.getGetPendingCommissionsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetPendingCommissionsMethod = CommissionServiceGrpc.getGetPendingCommissionsMethod) == null) {
          CommissionServiceGrpc.getGetPendingCommissionsMethod = getGetPendingCommissionsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetPendingCommissionsRequest, com.gameengine.commission.v1.GetPendingCommissionsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPendingCommissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetPendingCommissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetPendingCommissionsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetPendingCommissions"))
              .build();
        }
      }
    }
    return getGetPendingCommissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetCommissionHistoryRequest,
      com.gameengine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionHistory",
      requestType = com.gameengine.commission.v1.GetCommissionHistoryRequest.class,
      responseType = com.gameengine.commission.v1.GetCommissionHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetCommissionHistoryRequest,
      com.gameengine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod() {
    io.grpc.MethodDescriptor<com.gameengine.commission.v1.GetCommissionHistoryRequest, com.gameengine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod;
    if ((getGetCommissionHistoryMethod = CommissionServiceGrpc.getGetCommissionHistoryMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionHistoryMethod = CommissionServiceGrpc.getGetCommissionHistoryMethod) == null) {
          CommissionServiceGrpc.getGetCommissionHistoryMethod = getGetCommissionHistoryMethod =
              io.grpc.MethodDescriptor.<com.gameengine.commission.v1.GetCommissionHistoryRequest, com.gameengine.commission.v1.GetCommissionHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetCommissionHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.commission.v1.GetCommissionHistoryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionHistory"))
              .build();
        }
      }
    }
    return getGetCommissionHistoryMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CommissionServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommissionServiceStub>() {
        @java.lang.Override
        public CommissionServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommissionServiceStub(channel, callOptions);
        }
      };
    return CommissionServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static CommissionServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommissionServiceBlockingV2Stub>() {
        @java.lang.Override
        public CommissionServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommissionServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return CommissionServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CommissionServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommissionServiceBlockingStub>() {
        @java.lang.Override
        public CommissionServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommissionServiceBlockingStub(channel, callOptions);
        }
      };
    return CommissionServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CommissionServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CommissionServiceFutureStub>() {
        @java.lang.Override
        public CommissionServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CommissionServiceFutureStub(channel, callOptions);
        }
      };
    return CommissionServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Commission claims
     * </pre>
     */
    default void submitClaim(com.gameengine.commission.v1.SubmitClaimRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.SubmitClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserClaims(com.gameengine.commission.v1.GetUserClaimsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getClaimsByStatus(com.gameengine.commission.v1.GetClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClaimsByStatusMethod(), responseObserver);
    }

    /**
     */
    default void claimCommission(com.gameengine.commission.v1.ClaimCommissionRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.ClaimCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimCommissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    default void getUserSettlements(com.gameengine.commission.v1.GetUserSettlementsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserSettlementsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserSettlementsMethod(), responseObserver);
    }

    /**
     */
    default void getSettlementById(com.gameengine.commission.v1.GetSettlementByIdRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetSettlementByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSettlementByIdMethod(), responseObserver);
    }

    /**
     * <pre>
     * Totals
     * </pre>
     */
    default void getTotalPending(com.gameengine.commission.v1.GetTotalPendingRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalPendingResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalPendingMethod(), responseObserver);
    }

    /**
     */
    default void getTotalSettled(com.gameengine.commission.v1.GetTotalSettledRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalSettledResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalSettledMethod(), responseObserver);
    }

    /**
     * <pre>
     * Agent commissions
     * </pre>
     */
    default void getAgentCommissions(com.gameengine.commission.v1.GetAgentCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetAgentCommissionsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAgentCommissionsMethod(), responseObserver);
    }

    /**
     */
    default void getPendingCommissions(com.gameengine.commission.v1.GetPendingCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetPendingCommissionsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPendingCommissionsMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionHistory(com.gameengine.commission.v1.GetCommissionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetCommissionHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionHistoryMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service CommissionService.
   */
  public static abstract class CommissionServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return CommissionServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service CommissionService.
   */
  public static final class CommissionServiceStub
      extends io.grpc.stub.AbstractAsyncStub<CommissionServiceStub> {
    private CommissionServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission claims
     * </pre>
     */
    public void submitClaim(com.gameengine.commission.v1.SubmitClaimRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.SubmitClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserClaims(com.gameengine.commission.v1.GetUserClaimsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClaimsByStatus(com.gameengine.commission.v1.GetClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClaimsByStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void claimCommission(com.gameengine.commission.v1.ClaimCommissionRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.ClaimCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public void getUserSettlements(com.gameengine.commission.v1.GetUserSettlementsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserSettlementsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserSettlementsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSettlementById(com.gameengine.commission.v1.GetSettlementByIdRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetSettlementByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSettlementByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Totals
     * </pre>
     */
    public void getTotalPending(com.gameengine.commission.v1.GetTotalPendingRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalPendingResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalPendingMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTotalSettled(com.gameengine.commission.v1.GetTotalSettledRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalSettledResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalSettledMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Agent commissions
     * </pre>
     */
    public void getAgentCommissions(com.gameengine.commission.v1.GetAgentCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetAgentCommissionsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAgentCommissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPendingCommissions(com.gameengine.commission.v1.GetPendingCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetPendingCommissionsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPendingCommissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionHistory(com.gameengine.commission.v1.GetCommissionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetCommissionHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionHistoryMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service CommissionService.
   */
  public static final class CommissionServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<CommissionServiceBlockingV2Stub> {
    private CommissionServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission claims
     * </pre>
     */
    public com.gameengine.commission.v1.SubmitClaimResponse submitClaim(com.gameengine.commission.v1.SubmitClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetUserClaimsResponse getUserClaims(com.gameengine.commission.v1.GetUserClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetClaimsByStatusResponse getClaimsByStatus(com.gameengine.commission.v1.GetClaimsByStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.ClaimCommissionResponse claimCommission(com.gameengine.commission.v1.ClaimCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.gameengine.commission.v1.GetUserSettlementsResponse getUserSettlements(com.gameengine.commission.v1.GetUserSettlementsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserSettlementsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetSettlementByIdResponse getSettlementById(com.gameengine.commission.v1.GetSettlementByIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSettlementByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Totals
     * </pre>
     */
    public com.gameengine.commission.v1.GetTotalPendingResponse getTotalPending(com.gameengine.commission.v1.GetTotalPendingRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTotalPendingMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetTotalSettledResponse getTotalSettled(com.gameengine.commission.v1.GetTotalSettledRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTotalSettledMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Agent commissions
     * </pre>
     */
    public com.gameengine.commission.v1.GetAgentCommissionsResponse getAgentCommissions(com.gameengine.commission.v1.GetAgentCommissionsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAgentCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetPendingCommissionsResponse getPendingCommissions(com.gameengine.commission.v1.GetPendingCommissionsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPendingCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetCommissionHistoryResponse getCommissionHistory(com.gameengine.commission.v1.GetCommissionHistoryRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service CommissionService.
   */
  public static final class CommissionServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<CommissionServiceBlockingStub> {
    private CommissionServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission claims
     * </pre>
     */
    public com.gameengine.commission.v1.SubmitClaimResponse submitClaim(com.gameengine.commission.v1.SubmitClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetUserClaimsResponse getUserClaims(com.gameengine.commission.v1.GetUserClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetClaimsByStatusResponse getClaimsByStatus(com.gameengine.commission.v1.GetClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.ClaimCommissionResponse claimCommission(com.gameengine.commission.v1.ClaimCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.gameengine.commission.v1.GetUserSettlementsResponse getUserSettlements(com.gameengine.commission.v1.GetUserSettlementsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserSettlementsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetSettlementByIdResponse getSettlementById(com.gameengine.commission.v1.GetSettlementByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSettlementByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Totals
     * </pre>
     */
    public com.gameengine.commission.v1.GetTotalPendingResponse getTotalPending(com.gameengine.commission.v1.GetTotalPendingRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalPendingMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetTotalSettledResponse getTotalSettled(com.gameengine.commission.v1.GetTotalSettledRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalSettledMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Agent commissions
     * </pre>
     */
    public com.gameengine.commission.v1.GetAgentCommissionsResponse getAgentCommissions(com.gameengine.commission.v1.GetAgentCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAgentCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetPendingCommissionsResponse getPendingCommissions(com.gameengine.commission.v1.GetPendingCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPendingCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.commission.v1.GetCommissionHistoryResponse getCommissionHistory(com.gameengine.commission.v1.GetCommissionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service CommissionService.
   */
  public static final class CommissionServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<CommissionServiceFutureStub> {
    private CommissionServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission claims
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.SubmitClaimResponse> submitClaim(
        com.gameengine.commission.v1.SubmitClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetUserClaimsResponse> getUserClaims(
        com.gameengine.commission.v1.GetUserClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetClaimsByStatusResponse> getClaimsByStatus(
        com.gameengine.commission.v1.GetClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClaimsByStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.ClaimCommissionResponse> claimCommission(
        com.gameengine.commission.v1.ClaimCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimCommissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetUserSettlementsResponse> getUserSettlements(
        com.gameengine.commission.v1.GetUserSettlementsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserSettlementsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetSettlementByIdResponse> getSettlementById(
        com.gameengine.commission.v1.GetSettlementByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSettlementByIdMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Totals
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetTotalPendingResponse> getTotalPending(
        com.gameengine.commission.v1.GetTotalPendingRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalPendingMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetTotalSettledResponse> getTotalSettled(
        com.gameengine.commission.v1.GetTotalSettledRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalSettledMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Agent commissions
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetAgentCommissionsResponse> getAgentCommissions(
        com.gameengine.commission.v1.GetAgentCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAgentCommissionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetPendingCommissionsResponse> getPendingCommissions(
        com.gameengine.commission.v1.GetPendingCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPendingCommissionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.commission.v1.GetCommissionHistoryResponse> getCommissionHistory(
        com.gameengine.commission.v1.GetCommissionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionHistoryMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SUBMIT_CLAIM = 0;
  private static final int METHODID_GET_USER_CLAIMS = 1;
  private static final int METHODID_GET_CLAIMS_BY_STATUS = 2;
  private static final int METHODID_CLAIM_COMMISSION = 3;
  private static final int METHODID_GET_USER_SETTLEMENTS = 4;
  private static final int METHODID_GET_SETTLEMENT_BY_ID = 5;
  private static final int METHODID_GET_TOTAL_PENDING = 6;
  private static final int METHODID_GET_TOTAL_SETTLED = 7;
  private static final int METHODID_GET_AGENT_COMMISSIONS = 8;
  private static final int METHODID_GET_PENDING_COMMISSIONS = 9;
  private static final int METHODID_GET_COMMISSION_HISTORY = 10;

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
        case METHODID_SUBMIT_CLAIM:
          serviceImpl.submitClaim((com.gameengine.commission.v1.SubmitClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.SubmitClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_CLAIMS:
          serviceImpl.getUserClaims((com.gameengine.commission.v1.GetUserClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_CLAIMS_BY_STATUS:
          serviceImpl.getClaimsByStatus((com.gameengine.commission.v1.GetClaimsByStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetClaimsByStatusResponse>) responseObserver);
          break;
        case METHODID_CLAIM_COMMISSION:
          serviceImpl.claimCommission((com.gameengine.commission.v1.ClaimCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.ClaimCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_USER_SETTLEMENTS:
          serviceImpl.getUserSettlements((com.gameengine.commission.v1.GetUserSettlementsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetUserSettlementsResponse>) responseObserver);
          break;
        case METHODID_GET_SETTLEMENT_BY_ID:
          serviceImpl.getSettlementById((com.gameengine.commission.v1.GetSettlementByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetSettlementByIdResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_PENDING:
          serviceImpl.getTotalPending((com.gameengine.commission.v1.GetTotalPendingRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalPendingResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_SETTLED:
          serviceImpl.getTotalSettled((com.gameengine.commission.v1.GetTotalSettledRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetTotalSettledResponse>) responseObserver);
          break;
        case METHODID_GET_AGENT_COMMISSIONS:
          serviceImpl.getAgentCommissions((com.gameengine.commission.v1.GetAgentCommissionsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetAgentCommissionsResponse>) responseObserver);
          break;
        case METHODID_GET_PENDING_COMMISSIONS:
          serviceImpl.getPendingCommissions((com.gameengine.commission.v1.GetPendingCommissionsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetPendingCommissionsResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_HISTORY:
          serviceImpl.getCommissionHistory((com.gameengine.commission.v1.GetCommissionHistoryRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.commission.v1.GetCommissionHistoryResponse>) responseObserver);
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
          getSubmitClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.SubmitClaimRequest,
              com.gameengine.commission.v1.SubmitClaimResponse>(
                service, METHODID_SUBMIT_CLAIM)))
        .addMethod(
          getGetUserClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetUserClaimsRequest,
              com.gameengine.commission.v1.GetUserClaimsResponse>(
                service, METHODID_GET_USER_CLAIMS)))
        .addMethod(
          getGetClaimsByStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetClaimsByStatusRequest,
              com.gameengine.commission.v1.GetClaimsByStatusResponse>(
                service, METHODID_GET_CLAIMS_BY_STATUS)))
        .addMethod(
          getClaimCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.ClaimCommissionRequest,
              com.gameengine.commission.v1.ClaimCommissionResponse>(
                service, METHODID_CLAIM_COMMISSION)))
        .addMethod(
          getGetUserSettlementsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetUserSettlementsRequest,
              com.gameengine.commission.v1.GetUserSettlementsResponse>(
                service, METHODID_GET_USER_SETTLEMENTS)))
        .addMethod(
          getGetSettlementByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetSettlementByIdRequest,
              com.gameengine.commission.v1.GetSettlementByIdResponse>(
                service, METHODID_GET_SETTLEMENT_BY_ID)))
        .addMethod(
          getGetTotalPendingMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetTotalPendingRequest,
              com.gameengine.commission.v1.GetTotalPendingResponse>(
                service, METHODID_GET_TOTAL_PENDING)))
        .addMethod(
          getGetTotalSettledMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetTotalSettledRequest,
              com.gameengine.commission.v1.GetTotalSettledResponse>(
                service, METHODID_GET_TOTAL_SETTLED)))
        .addMethod(
          getGetAgentCommissionsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetAgentCommissionsRequest,
              com.gameengine.commission.v1.GetAgentCommissionsResponse>(
                service, METHODID_GET_AGENT_COMMISSIONS)))
        .addMethod(
          getGetPendingCommissionsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetPendingCommissionsRequest,
              com.gameengine.commission.v1.GetPendingCommissionsResponse>(
                service, METHODID_GET_PENDING_COMMISSIONS)))
        .addMethod(
          getGetCommissionHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.commission.v1.GetCommissionHistoryRequest,
              com.gameengine.commission.v1.GetCommissionHistoryResponse>(
                service, METHODID_GET_COMMISSION_HISTORY)))
        .build();
  }

  private static abstract class CommissionServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CommissionServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.gameengine.commission.v1.CommissionServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CommissionService");
    }
  }

  private static final class CommissionServiceFileDescriptorSupplier
      extends CommissionServiceBaseDescriptorSupplier {
    CommissionServiceFileDescriptorSupplier() {}
  }

  private static final class CommissionServiceMethodDescriptorSupplier
      extends CommissionServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    CommissionServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (CommissionServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CommissionServiceFileDescriptorSupplier())
              .addMethod(getSubmitClaimMethod())
              .addMethod(getGetUserClaimsMethod())
              .addMethod(getGetClaimsByStatusMethod())
              .addMethod(getClaimCommissionMethod())
              .addMethod(getGetUserSettlementsMethod())
              .addMethod(getGetSettlementByIdMethod())
              .addMethod(getGetTotalPendingMethod())
              .addMethod(getGetTotalSettledMethod())
              .addMethod(getGetAgentCommissionsMethod())
              .addMethod(getGetPendingCommissionsMethod())
              .addMethod(getGetCommissionHistoryMethod())
              .build();
        }
      }
    }
    return result;
  }
}
