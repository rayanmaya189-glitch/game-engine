package com.game_engine.commission.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * =============================================================================
 * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
 * =============================================================================
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class ClaimServiceGrpc {

  private ClaimServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.commission.v1.ClaimService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitCommissionClaimRequest,
      com.game_engine.commission.v1.SubmitCommissionClaimResponse> getSubmitCommissionClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitCommissionClaim",
      requestType = com.game_engine.commission.v1.SubmitCommissionClaimRequest.class,
      responseType = com.game_engine.commission.v1.SubmitCommissionClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitCommissionClaimRequest,
      com.game_engine.commission.v1.SubmitCommissionClaimResponse> getSubmitCommissionClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitCommissionClaimRequest, com.game_engine.commission.v1.SubmitCommissionClaimResponse> getSubmitCommissionClaimMethod;
    if ((getSubmitCommissionClaimMethod = ClaimServiceGrpc.getSubmitCommissionClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getSubmitCommissionClaimMethod = ClaimServiceGrpc.getSubmitCommissionClaimMethod) == null) {
          ClaimServiceGrpc.getSubmitCommissionClaimMethod = getSubmitCommissionClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.SubmitCommissionClaimRequest, com.game_engine.commission.v1.SubmitCommissionClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitCommissionClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitCommissionClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitCommissionClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("SubmitCommissionClaim"))
              .build();
        }
      }
    }
    return getSubmitCommissionClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserCommissionClaimsRequest,
      com.game_engine.commission.v1.GetUserCommissionClaimsResponse> getGetUserCommissionClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserCommissionClaims",
      requestType = com.game_engine.commission.v1.GetUserCommissionClaimsRequest.class,
      responseType = com.game_engine.commission.v1.GetUserCommissionClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserCommissionClaimsRequest,
      com.game_engine.commission.v1.GetUserCommissionClaimsResponse> getGetUserCommissionClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserCommissionClaimsRequest, com.game_engine.commission.v1.GetUserCommissionClaimsResponse> getGetUserCommissionClaimsMethod;
    if ((getGetUserCommissionClaimsMethod = ClaimServiceGrpc.getGetUserCommissionClaimsMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserCommissionClaimsMethod = ClaimServiceGrpc.getGetUserCommissionClaimsMethod) == null) {
          ClaimServiceGrpc.getGetUserCommissionClaimsMethod = getGetUserCommissionClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserCommissionClaimsRequest, com.game_engine.commission.v1.GetUserCommissionClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserCommissionClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserCommissionClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserCommissionClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserCommissionClaims"))
              .build();
        }
      }
    }
    return getGetUserCommissionClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest,
      com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> getGetCommissionClaimsByStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionClaimsByStatus",
      requestType = com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest,
      com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> getGetCommissionClaimsByStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest, com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> getGetCommissionClaimsByStatusMethod;
    if ((getGetCommissionClaimsByStatusMethod = ClaimServiceGrpc.getGetCommissionClaimsByStatusMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetCommissionClaimsByStatusMethod = ClaimServiceGrpc.getGetCommissionClaimsByStatusMethod) == null) {
          ClaimServiceGrpc.getGetCommissionClaimsByStatusMethod = getGetCommissionClaimsByStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest, com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionClaimsByStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetCommissionClaimsByStatus"))
              .build();
        }
      }
    }
    return getGetCommissionClaimsByStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimByIdRequest,
      com.game_engine.commission.v1.GetCommissionClaimByIdResponse> getGetCommissionClaimByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionClaimById",
      requestType = com.game_engine.commission.v1.GetCommissionClaimByIdRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionClaimByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimByIdRequest,
      com.game_engine.commission.v1.GetCommissionClaimByIdResponse> getGetCommissionClaimByIdMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionClaimByIdRequest, com.game_engine.commission.v1.GetCommissionClaimByIdResponse> getGetCommissionClaimByIdMethod;
    if ((getGetCommissionClaimByIdMethod = ClaimServiceGrpc.getGetCommissionClaimByIdMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetCommissionClaimByIdMethod = ClaimServiceGrpc.getGetCommissionClaimByIdMethod) == null) {
          ClaimServiceGrpc.getGetCommissionClaimByIdMethod = getGetCommissionClaimByIdMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionClaimByIdRequest, com.game_engine.commission.v1.GetCommissionClaimByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionClaimById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionClaimByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionClaimByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetCommissionClaimById"))
              .build();
        }
      }
    }
    return getGetCommissionClaimByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionClaimRequest,
      com.game_engine.commission.v1.ApproveCommissionClaimResponse> getApproveCommissionClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApproveCommissionClaim",
      requestType = com.game_engine.commission.v1.ApproveCommissionClaimRequest.class,
      responseType = com.game_engine.commission.v1.ApproveCommissionClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionClaimRequest,
      com.game_engine.commission.v1.ApproveCommissionClaimResponse> getApproveCommissionClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionClaimRequest, com.game_engine.commission.v1.ApproveCommissionClaimResponse> getApproveCommissionClaimMethod;
    if ((getApproveCommissionClaimMethod = ClaimServiceGrpc.getApproveCommissionClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getApproveCommissionClaimMethod = ClaimServiceGrpc.getApproveCommissionClaimMethod) == null) {
          ClaimServiceGrpc.getApproveCommissionClaimMethod = getApproveCommissionClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ApproveCommissionClaimRequest, com.game_engine.commission.v1.ApproveCommissionClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApproveCommissionClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveCommissionClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveCommissionClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("ApproveCommissionClaim"))
              .build();
        }
      }
    }
    return getApproveCommissionClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionClaimRequest,
      com.game_engine.commission.v1.RejectCommissionClaimResponse> getRejectCommissionClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RejectCommissionClaim",
      requestType = com.game_engine.commission.v1.RejectCommissionClaimRequest.class,
      responseType = com.game_engine.commission.v1.RejectCommissionClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionClaimRequest,
      com.game_engine.commission.v1.RejectCommissionClaimResponse> getRejectCommissionClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionClaimRequest, com.game_engine.commission.v1.RejectCommissionClaimResponse> getRejectCommissionClaimMethod;
    if ((getRejectCommissionClaimMethod = ClaimServiceGrpc.getRejectCommissionClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getRejectCommissionClaimMethod = ClaimServiceGrpc.getRejectCommissionClaimMethod) == null) {
          ClaimServiceGrpc.getRejectCommissionClaimMethod = getRejectCommissionClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.RejectCommissionClaimRequest, com.game_engine.commission.v1.RejectCommissionClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RejectCommissionClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectCommissionClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectCommissionClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("RejectCommissionClaim"))
              .build();
        }
      }
    }
    return getRejectCommissionClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionClaimRequest,
      com.game_engine.commission.v1.PayCommissionClaimResponse> getPayCommissionClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PayCommissionClaim",
      requestType = com.game_engine.commission.v1.PayCommissionClaimRequest.class,
      responseType = com.game_engine.commission.v1.PayCommissionClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionClaimRequest,
      com.game_engine.commission.v1.PayCommissionClaimResponse> getPayCommissionClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionClaimRequest, com.game_engine.commission.v1.PayCommissionClaimResponse> getPayCommissionClaimMethod;
    if ((getPayCommissionClaimMethod = ClaimServiceGrpc.getPayCommissionClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getPayCommissionClaimMethod = ClaimServiceGrpc.getPayCommissionClaimMethod) == null) {
          ClaimServiceGrpc.getPayCommissionClaimMethod = getPayCommissionClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.PayCommissionClaimRequest, com.game_engine.commission.v1.PayCommissionClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PayCommissionClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayCommissionClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayCommissionClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("PayCommissionClaim"))
              .build();
        }
      }
    }
    return getPayCommissionClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateRebetClaimRequest,
      com.game_engine.commission.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateRebetClaim",
      requestType = com.game_engine.commission.v1.CreateRebetClaimRequest.class,
      responseType = com.game_engine.commission.v1.CreateRebetClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateRebetClaimRequest,
      com.game_engine.commission.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateRebetClaimRequest, com.game_engine.commission.v1.CreateRebetClaimResponse> getCreateRebetClaimMethod;
    if ((getCreateRebetClaimMethod = ClaimServiceGrpc.getCreateRebetClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getCreateRebetClaimMethod = ClaimServiceGrpc.getCreateRebetClaimMethod) == null) {
          ClaimServiceGrpc.getCreateRebetClaimMethod = getCreateRebetClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.CreateRebetClaimRequest, com.game_engine.commission.v1.CreateRebetClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateRebetClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateRebetClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateRebetClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("CreateRebetClaim"))
              .build();
        }
      }
    }
    return getCreateRebetClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateRebetProgressRequest,
      com.game_engine.commission.v1.UpdateRebetProgressResponse> getUpdateRebetProgressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateRebetProgress",
      requestType = com.game_engine.commission.v1.UpdateRebetProgressRequest.class,
      responseType = com.game_engine.commission.v1.UpdateRebetProgressResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateRebetProgressRequest,
      com.game_engine.commission.v1.UpdateRebetProgressResponse> getUpdateRebetProgressMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateRebetProgressRequest, com.game_engine.commission.v1.UpdateRebetProgressResponse> getUpdateRebetProgressMethod;
    if ((getUpdateRebetProgressMethod = ClaimServiceGrpc.getUpdateRebetProgressMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getUpdateRebetProgressMethod = ClaimServiceGrpc.getUpdateRebetProgressMethod) == null) {
          ClaimServiceGrpc.getUpdateRebetProgressMethod = getUpdateRebetProgressMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.UpdateRebetProgressRequest, com.game_engine.commission.v1.UpdateRebetProgressResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateRebetProgress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.UpdateRebetProgressRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.UpdateRebetProgressResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("UpdateRebetProgress"))
              .build();
        }
      }
    }
    return getUpdateRebetProgressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimRebetRequest,
      com.game_engine.commission.v1.ClaimRebetResponse> getClaimRebetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClaimRebet",
      requestType = com.game_engine.commission.v1.ClaimRebetRequest.class,
      responseType = com.game_engine.commission.v1.ClaimRebetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimRebetRequest,
      com.game_engine.commission.v1.ClaimRebetResponse> getClaimRebetMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimRebetRequest, com.game_engine.commission.v1.ClaimRebetResponse> getClaimRebetMethod;
    if ((getClaimRebetMethod = ClaimServiceGrpc.getClaimRebetMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getClaimRebetMethod = ClaimServiceGrpc.getClaimRebetMethod) == null) {
          ClaimServiceGrpc.getClaimRebetMethod = getClaimRebetMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ClaimRebetRequest, com.game_engine.commission.v1.ClaimRebetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClaimRebet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ClaimRebetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ClaimRebetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("ClaimRebet"))
              .build();
        }
      }
    }
    return getClaimRebetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserRebetClaimsRequest,
      com.game_engine.commission.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserRebetClaims",
      requestType = com.game_engine.commission.v1.GetUserRebetClaimsRequest.class,
      responseType = com.game_engine.commission.v1.GetUserRebetClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserRebetClaimsRequest,
      com.game_engine.commission.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserRebetClaimsRequest, com.game_engine.commission.v1.GetUserRebetClaimsResponse> getGetUserRebetClaimsMethod;
    if ((getGetUserRebetClaimsMethod = ClaimServiceGrpc.getGetUserRebetClaimsMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserRebetClaimsMethod = ClaimServiceGrpc.getGetUserRebetClaimsMethod) == null) {
          ClaimServiceGrpc.getGetUserRebetClaimsMethod = getGetUserRebetClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserRebetClaimsRequest, com.game_engine.commission.v1.GetUserRebetClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserRebetClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserRebetClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserRebetClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserRebetClaims"))
              .build();
        }
      }
    }
    return getGetUserRebetClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimableRebetsRequest,
      com.game_engine.commission.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClaimableRebets",
      requestType = com.game_engine.commission.v1.GetClaimableRebetsRequest.class,
      responseType = com.game_engine.commission.v1.GetClaimableRebetsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimableRebetsRequest,
      com.game_engine.commission.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimableRebetsRequest, com.game_engine.commission.v1.GetClaimableRebetsResponse> getGetClaimableRebetsMethod;
    if ((getGetClaimableRebetsMethod = ClaimServiceGrpc.getGetClaimableRebetsMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetClaimableRebetsMethod = ClaimServiceGrpc.getGetClaimableRebetsMethod) == null) {
          ClaimServiceGrpc.getGetClaimableRebetsMethod = getGetClaimableRebetsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetClaimableRebetsRequest, com.game_engine.commission.v1.GetClaimableRebetsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClaimableRebets"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetClaimableRebetsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetClaimableRebetsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetClaimableRebets"))
              .build();
        }
      }
    }
    return getGetClaimableRebetsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitInsuranceClaimRequest,
      com.game_engine.commission.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitInsuranceClaim",
      requestType = com.game_engine.commission.v1.SubmitInsuranceClaimRequest.class,
      responseType = com.game_engine.commission.v1.SubmitInsuranceClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitInsuranceClaimRequest,
      com.game_engine.commission.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitInsuranceClaimRequest, com.game_engine.commission.v1.SubmitInsuranceClaimResponse> getSubmitInsuranceClaimMethod;
    if ((getSubmitInsuranceClaimMethod = ClaimServiceGrpc.getSubmitInsuranceClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getSubmitInsuranceClaimMethod = ClaimServiceGrpc.getSubmitInsuranceClaimMethod) == null) {
          ClaimServiceGrpc.getSubmitInsuranceClaimMethod = getSubmitInsuranceClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.SubmitInsuranceClaimRequest, com.game_engine.commission.v1.SubmitInsuranceClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitInsuranceClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitInsuranceClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitInsuranceClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("SubmitInsuranceClaim"))
              .build();
        }
      }
    }
    return getSubmitInsuranceClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveInsuranceClaimRequest,
      com.game_engine.commission.v1.ApproveInsuranceClaimResponse> getApproveInsuranceClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApproveInsuranceClaim",
      requestType = com.game_engine.commission.v1.ApproveInsuranceClaimRequest.class,
      responseType = com.game_engine.commission.v1.ApproveInsuranceClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveInsuranceClaimRequest,
      com.game_engine.commission.v1.ApproveInsuranceClaimResponse> getApproveInsuranceClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveInsuranceClaimRequest, com.game_engine.commission.v1.ApproveInsuranceClaimResponse> getApproveInsuranceClaimMethod;
    if ((getApproveInsuranceClaimMethod = ClaimServiceGrpc.getApproveInsuranceClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getApproveInsuranceClaimMethod = ClaimServiceGrpc.getApproveInsuranceClaimMethod) == null) {
          ClaimServiceGrpc.getApproveInsuranceClaimMethod = getApproveInsuranceClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ApproveInsuranceClaimRequest, com.game_engine.commission.v1.ApproveInsuranceClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApproveInsuranceClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveInsuranceClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveInsuranceClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("ApproveInsuranceClaim"))
              .build();
        }
      }
    }
    return getApproveInsuranceClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectInsuranceClaimRequest,
      com.game_engine.commission.v1.RejectInsuranceClaimResponse> getRejectInsuranceClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RejectInsuranceClaim",
      requestType = com.game_engine.commission.v1.RejectInsuranceClaimRequest.class,
      responseType = com.game_engine.commission.v1.RejectInsuranceClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectInsuranceClaimRequest,
      com.game_engine.commission.v1.RejectInsuranceClaimResponse> getRejectInsuranceClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectInsuranceClaimRequest, com.game_engine.commission.v1.RejectInsuranceClaimResponse> getRejectInsuranceClaimMethod;
    if ((getRejectInsuranceClaimMethod = ClaimServiceGrpc.getRejectInsuranceClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getRejectInsuranceClaimMethod = ClaimServiceGrpc.getRejectInsuranceClaimMethod) == null) {
          ClaimServiceGrpc.getRejectInsuranceClaimMethod = getRejectInsuranceClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.RejectInsuranceClaimRequest, com.game_engine.commission.v1.RejectInsuranceClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RejectInsuranceClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectInsuranceClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectInsuranceClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("RejectInsuranceClaim"))
              .build();
        }
      }
    }
    return getRejectInsuranceClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayInsuranceClaimRequest,
      com.game_engine.commission.v1.PayInsuranceClaimResponse> getPayInsuranceClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PayInsuranceClaim",
      requestType = com.game_engine.commission.v1.PayInsuranceClaimRequest.class,
      responseType = com.game_engine.commission.v1.PayInsuranceClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayInsuranceClaimRequest,
      com.game_engine.commission.v1.PayInsuranceClaimResponse> getPayInsuranceClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayInsuranceClaimRequest, com.game_engine.commission.v1.PayInsuranceClaimResponse> getPayInsuranceClaimMethod;
    if ((getPayInsuranceClaimMethod = ClaimServiceGrpc.getPayInsuranceClaimMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getPayInsuranceClaimMethod = ClaimServiceGrpc.getPayInsuranceClaimMethod) == null) {
          ClaimServiceGrpc.getPayInsuranceClaimMethod = getPayInsuranceClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.PayInsuranceClaimRequest, com.game_engine.commission.v1.PayInsuranceClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PayInsuranceClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayInsuranceClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayInsuranceClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("PayInsuranceClaim"))
              .build();
        }
      }
    }
    return getPayInsuranceClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserInsuranceClaimsRequest,
      com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserInsuranceClaims",
      requestType = com.game_engine.commission.v1.GetUserInsuranceClaimsRequest.class,
      responseType = com.game_engine.commission.v1.GetUserInsuranceClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserInsuranceClaimsRequest,
      com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserInsuranceClaimsRequest, com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> getGetUserInsuranceClaimsMethod;
    if ((getGetUserInsuranceClaimsMethod = ClaimServiceGrpc.getGetUserInsuranceClaimsMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserInsuranceClaimsMethod = ClaimServiceGrpc.getGetUserInsuranceClaimsMethod) == null) {
          ClaimServiceGrpc.getGetUserInsuranceClaimsMethod = getGetUserInsuranceClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserInsuranceClaimsRequest, com.game_engine.commission.v1.GetUserInsuranceClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserInsuranceClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserInsuranceClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserInsuranceClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserInsuranceClaims"))
              .build();
        }
      }
    }
    return getGetUserInsuranceClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest,
      com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> getGetInsuranceClaimsByStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetInsuranceClaimsByStatus",
      requestType = com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest.class,
      responseType = com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest,
      com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> getGetInsuranceClaimsByStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest, com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> getGetInsuranceClaimsByStatusMethod;
    if ((getGetInsuranceClaimsByStatusMethod = ClaimServiceGrpc.getGetInsuranceClaimsByStatusMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetInsuranceClaimsByStatusMethod = ClaimServiceGrpc.getGetInsuranceClaimsByStatusMethod) == null) {
          ClaimServiceGrpc.getGetInsuranceClaimsByStatusMethod = getGetInsuranceClaimsByStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest, com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetInsuranceClaimsByStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetInsuranceClaimsByStatus"))
              .build();
        }
      }
    }
    return getGetInsuranceClaimsByStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserSettlementsRequest,
      com.game_engine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserSettlements",
      requestType = com.game_engine.commission.v1.GetUserSettlementsRequest.class,
      responseType = com.game_engine.commission.v1.GetUserSettlementsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserSettlementsRequest,
      com.game_engine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserSettlementsRequest, com.game_engine.commission.v1.GetUserSettlementsResponse> getGetUserSettlementsMethod;
    if ((getGetUserSettlementsMethod = ClaimServiceGrpc.getGetUserSettlementsMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserSettlementsMethod = ClaimServiceGrpc.getGetUserSettlementsMethod) == null) {
          ClaimServiceGrpc.getGetUserSettlementsMethod = getGetUserSettlementsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserSettlementsRequest, com.game_engine.commission.v1.GetUserSettlementsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserSettlements"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserSettlementsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserSettlementsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserSettlements"))
              .build();
        }
      }
    }
    return getGetUserSettlementsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByStatusRequest,
      com.game_engine.commission.v1.GetSettlementsByStatusResponse> getGetSettlementsByStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSettlementsByStatus",
      requestType = com.game_engine.commission.v1.GetSettlementsByStatusRequest.class,
      responseType = com.game_engine.commission.v1.GetSettlementsByStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByStatusRequest,
      com.game_engine.commission.v1.GetSettlementsByStatusResponse> getGetSettlementsByStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByStatusRequest, com.game_engine.commission.v1.GetSettlementsByStatusResponse> getGetSettlementsByStatusMethod;
    if ((getGetSettlementsByStatusMethod = ClaimServiceGrpc.getGetSettlementsByStatusMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetSettlementsByStatusMethod = ClaimServiceGrpc.getGetSettlementsByStatusMethod) == null) {
          ClaimServiceGrpc.getGetSettlementsByStatusMethod = getGetSettlementsByStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetSettlementsByStatusRequest, com.game_engine.commission.v1.GetSettlementsByStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSettlementsByStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementsByStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementsByStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetSettlementsByStatus"))
              .build();
        }
      }
    }
    return getGetSettlementsByStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByTypeRequest,
      com.game_engine.commission.v1.GetSettlementsByTypeResponse> getGetSettlementsByTypeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSettlementsByType",
      requestType = com.game_engine.commission.v1.GetSettlementsByTypeRequest.class,
      responseType = com.game_engine.commission.v1.GetSettlementsByTypeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByTypeRequest,
      com.game_engine.commission.v1.GetSettlementsByTypeResponse> getGetSettlementsByTypeMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementsByTypeRequest, com.game_engine.commission.v1.GetSettlementsByTypeResponse> getGetSettlementsByTypeMethod;
    if ((getGetSettlementsByTypeMethod = ClaimServiceGrpc.getGetSettlementsByTypeMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetSettlementsByTypeMethod = ClaimServiceGrpc.getGetSettlementsByTypeMethod) == null) {
          ClaimServiceGrpc.getGetSettlementsByTypeMethod = getGetSettlementsByTypeMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetSettlementsByTypeRequest, com.game_engine.commission.v1.GetSettlementsByTypeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSettlementsByType"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementsByTypeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementsByTypeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetSettlementsByType"))
              .build();
        }
      }
    }
    return getGetSettlementsByTypeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementByIdRequest,
      com.game_engine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSettlementById",
      requestType = com.game_engine.commission.v1.GetSettlementByIdRequest.class,
      responseType = com.game_engine.commission.v1.GetSettlementByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementByIdRequest,
      com.game_engine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetSettlementByIdRequest, com.game_engine.commission.v1.GetSettlementByIdResponse> getGetSettlementByIdMethod;
    if ((getGetSettlementByIdMethod = ClaimServiceGrpc.getGetSettlementByIdMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetSettlementByIdMethod = ClaimServiceGrpc.getGetSettlementByIdMethod) == null) {
          ClaimServiceGrpc.getGetSettlementByIdMethod = getGetSettlementByIdMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetSettlementByIdRequest, com.game_engine.commission.v1.GetSettlementByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSettlementById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetSettlementByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetSettlementById"))
              .build();
        }
      }
    }
    return getGetSettlementByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalPendingRequest,
      com.game_engine.commission.v1.GetUserTotalPendingResponse> getGetUserTotalPendingMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserTotalPending",
      requestType = com.game_engine.commission.v1.GetUserTotalPendingRequest.class,
      responseType = com.game_engine.commission.v1.GetUserTotalPendingResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalPendingRequest,
      com.game_engine.commission.v1.GetUserTotalPendingResponse> getGetUserTotalPendingMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalPendingRequest, com.game_engine.commission.v1.GetUserTotalPendingResponse> getGetUserTotalPendingMethod;
    if ((getGetUserTotalPendingMethod = ClaimServiceGrpc.getGetUserTotalPendingMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserTotalPendingMethod = ClaimServiceGrpc.getGetUserTotalPendingMethod) == null) {
          ClaimServiceGrpc.getGetUserTotalPendingMethod = getGetUserTotalPendingMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserTotalPendingRequest, com.game_engine.commission.v1.GetUserTotalPendingResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserTotalPending"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserTotalPendingRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserTotalPendingResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserTotalPending"))
              .build();
        }
      }
    }
    return getGetUserTotalPendingMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalSettledRequest,
      com.game_engine.commission.v1.GetUserTotalSettledResponse> getGetUserTotalSettledMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserTotalSettled",
      requestType = com.game_engine.commission.v1.GetUserTotalSettledRequest.class,
      responseType = com.game_engine.commission.v1.GetUserTotalSettledResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalSettledRequest,
      com.game_engine.commission.v1.GetUserTotalSettledResponse> getGetUserTotalSettledMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserTotalSettledRequest, com.game_engine.commission.v1.GetUserTotalSettledResponse> getGetUserTotalSettledMethod;
    if ((getGetUserTotalSettledMethod = ClaimServiceGrpc.getGetUserTotalSettledMethod) == null) {
      synchronized (ClaimServiceGrpc.class) {
        if ((getGetUserTotalSettledMethod = ClaimServiceGrpc.getGetUserTotalSettledMethod) == null) {
          ClaimServiceGrpc.getGetUserTotalSettledMethod = getGetUserTotalSettledMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserTotalSettledRequest, com.game_engine.commission.v1.GetUserTotalSettledResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserTotalSettled"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserTotalSettledRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserTotalSettledResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ClaimServiceMethodDescriptorSupplier("GetUserTotalSettled"))
              .build();
        }
      }
    }
    return getGetUserTotalSettledMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ClaimServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClaimServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClaimServiceStub>() {
        @java.lang.Override
        public ClaimServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClaimServiceStub(channel, callOptions);
        }
      };
    return ClaimServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static ClaimServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClaimServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClaimServiceBlockingV2Stub>() {
        @java.lang.Override
        public ClaimServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClaimServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return ClaimServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ClaimServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClaimServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClaimServiceBlockingStub>() {
        @java.lang.Override
        public ClaimServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClaimServiceBlockingStub(channel, callOptions);
        }
      };
    return ClaimServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ClaimServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ClaimServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ClaimServiceFutureStub>() {
        @java.lang.Override
        public ClaimServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ClaimServiceFutureStub(channel, callOptions);
        }
      };
    return ClaimServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Commission Claims
     * </pre>
     */
    default void submitCommissionClaim(com.game_engine.commission.v1.SubmitCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitCommissionClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserCommissionClaims(com.game_engine.commission.v1.GetUserCommissionClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserCommissionClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserCommissionClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionClaimsByStatus(com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionClaimsByStatusMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionClaimById(com.game_engine.commission.v1.GetCommissionClaimByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionClaimByIdMethod(), responseObserver);
    }

    /**
     */
    default void approveCommissionClaim(com.game_engine.commission.v1.ApproveCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApproveCommissionClaimMethod(), responseObserver);
    }

    /**
     */
    default void rejectCommissionClaim(com.game_engine.commission.v1.RejectCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRejectCommissionClaimMethod(), responseObserver);
    }

    /**
     */
    default void payCommissionClaim(com.game_engine.commission.v1.PayCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPayCommissionClaimMethod(), responseObserver);
    }

    /**
     * <pre>
     * Rebet Claims
     * </pre>
     */
    default void createRebetClaim(com.game_engine.commission.v1.CreateRebetClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateRebetClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateRebetClaimMethod(), responseObserver);
    }

    /**
     */
    default void updateRebetProgress(com.game_engine.commission.v1.UpdateRebetProgressRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateRebetProgressResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateRebetProgressMethod(), responseObserver);
    }

    /**
     */
    default void claimRebet(com.game_engine.commission.v1.ClaimRebetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimRebetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimRebetMethod(), responseObserver);
    }

    /**
     */
    default void getUserRebetClaims(com.game_engine.commission.v1.GetUserRebetClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserRebetClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserRebetClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getClaimableRebets(com.game_engine.commission.v1.GetClaimableRebetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimableRebetsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClaimableRebetsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Insurance Claims
     * </pre>
     */
    default void submitInsuranceClaim(com.game_engine.commission.v1.SubmitInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitInsuranceClaimMethod(), responseObserver);
    }

    /**
     */
    default void approveInsuranceClaim(com.game_engine.commission.v1.ApproveInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApproveInsuranceClaimMethod(), responseObserver);
    }

    /**
     */
    default void rejectInsuranceClaim(com.game_engine.commission.v1.RejectInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRejectInsuranceClaimMethod(), responseObserver);
    }

    /**
     */
    default void payInsuranceClaim(com.game_engine.commission.v1.PayInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPayInsuranceClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserInsuranceClaims(com.game_engine.commission.v1.GetUserInsuranceClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserInsuranceClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getInsuranceClaimsByStatus(com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetInsuranceClaimsByStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    default void getUserSettlements(com.game_engine.commission.v1.GetUserSettlementsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserSettlementsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserSettlementsMethod(), responseObserver);
    }

    /**
     */
    default void getSettlementsByStatus(com.game_engine.commission.v1.GetSettlementsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSettlementsByStatusMethod(), responseObserver);
    }

    /**
     */
    default void getSettlementsByType(com.game_engine.commission.v1.GetSettlementsByTypeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByTypeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSettlementsByTypeMethod(), responseObserver);
    }

    /**
     */
    default void getSettlementById(com.game_engine.commission.v1.GetSettlementByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSettlementByIdMethod(), responseObserver);
    }

    /**
     */
    default void getUserTotalPending(com.game_engine.commission.v1.GetUserTotalPendingRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalPendingResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserTotalPendingMethod(), responseObserver);
    }

    /**
     */
    default void getUserTotalSettled(com.game_engine.commission.v1.GetUserTotalSettledRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalSettledResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserTotalSettledMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service ClaimService.
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public static abstract class ClaimServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return ClaimServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service ClaimService.
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public static final class ClaimServiceStub
      extends io.grpc.stub.AbstractAsyncStub<ClaimServiceStub> {
    private ClaimServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClaimServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClaimServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission Claims
     * </pre>
     */
    public void submitCommissionClaim(com.game_engine.commission.v1.SubmitCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitCommissionClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserCommissionClaims(com.game_engine.commission.v1.GetUserCommissionClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserCommissionClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserCommissionClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionClaimsByStatus(com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionClaimsByStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionClaimById(com.game_engine.commission.v1.GetCommissionClaimByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionClaimByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void approveCommissionClaim(com.game_engine.commission.v1.ApproveCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApproveCommissionClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void rejectCommissionClaim(com.game_engine.commission.v1.RejectCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRejectCommissionClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void payCommissionClaim(com.game_engine.commission.v1.PayCommissionClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPayCommissionClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Rebet Claims
     * </pre>
     */
    public void createRebetClaim(com.game_engine.commission.v1.CreateRebetClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateRebetClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateRebetClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateRebetProgress(com.game_engine.commission.v1.UpdateRebetProgressRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateRebetProgressResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateRebetProgressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void claimRebet(com.game_engine.commission.v1.ClaimRebetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimRebetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimRebetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserRebetClaims(com.game_engine.commission.v1.GetUserRebetClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserRebetClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserRebetClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClaimableRebets(com.game_engine.commission.v1.GetClaimableRebetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimableRebetsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClaimableRebetsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Insurance Claims
     * </pre>
     */
    public void submitInsuranceClaim(com.game_engine.commission.v1.SubmitInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitInsuranceClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void approveInsuranceClaim(com.game_engine.commission.v1.ApproveInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApproveInsuranceClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void rejectInsuranceClaim(com.game_engine.commission.v1.RejectInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRejectInsuranceClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void payInsuranceClaim(com.game_engine.commission.v1.PayInsuranceClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayInsuranceClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPayInsuranceClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserInsuranceClaims(com.game_engine.commission.v1.GetUserInsuranceClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserInsuranceClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getInsuranceClaimsByStatus(com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetInsuranceClaimsByStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public void getUserSettlements(com.game_engine.commission.v1.GetUserSettlementsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserSettlementsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserSettlementsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSettlementsByStatus(com.game_engine.commission.v1.GetSettlementsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSettlementsByStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSettlementsByType(com.game_engine.commission.v1.GetSettlementsByTypeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByTypeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSettlementsByTypeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSettlementById(com.game_engine.commission.v1.GetSettlementByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSettlementByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserTotalPending(com.game_engine.commission.v1.GetUserTotalPendingRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalPendingResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserTotalPendingMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserTotalSettled(com.game_engine.commission.v1.GetUserTotalSettledRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalSettledResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserTotalSettledMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service ClaimService.
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public static final class ClaimServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<ClaimServiceBlockingV2Stub> {
    private ClaimServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClaimServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClaimServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission Claims
     * </pre>
     */
    public com.game_engine.commission.v1.SubmitCommissionClaimResponse submitCommissionClaim(com.game_engine.commission.v1.SubmitCommissionClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserCommissionClaimsResponse getUserCommissionClaims(com.game_engine.commission.v1.GetUserCommissionClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserCommissionClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse getCommissionClaimsByStatus(com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionClaimByIdResponse getCommissionClaimById(com.game_engine.commission.v1.GetCommissionClaimByIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionClaimByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ApproveCommissionClaimResponse approveCommissionClaim(com.game_engine.commission.v1.ApproveCommissionClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApproveCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectCommissionClaimResponse rejectCommissionClaim(com.game_engine.commission.v1.RejectCommissionClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRejectCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayCommissionClaimResponse payCommissionClaim(com.game_engine.commission.v1.PayCommissionClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getPayCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Rebet Claims
     * </pre>
     */
    public com.game_engine.commission.v1.CreateRebetClaimResponse createRebetClaim(com.game_engine.commission.v1.CreateRebetClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateRebetClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.UpdateRebetProgressResponse updateRebetProgress(com.game_engine.commission.v1.UpdateRebetProgressRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateRebetProgressMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ClaimRebetResponse claimRebet(com.game_engine.commission.v1.ClaimRebetRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimRebetMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserRebetClaimsResponse getUserRebetClaims(com.game_engine.commission.v1.GetUserRebetClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserRebetClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetClaimableRebetsResponse getClaimableRebets(com.game_engine.commission.v1.GetClaimableRebetsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClaimableRebetsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Insurance Claims
     * </pre>
     */
    public com.game_engine.commission.v1.SubmitInsuranceClaimResponse submitInsuranceClaim(com.game_engine.commission.v1.SubmitInsuranceClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ApproveInsuranceClaimResponse approveInsuranceClaim(com.game_engine.commission.v1.ApproveInsuranceClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApproveInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectInsuranceClaimResponse rejectInsuranceClaim(com.game_engine.commission.v1.RejectInsuranceClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRejectInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayInsuranceClaimResponse payInsuranceClaim(com.game_engine.commission.v1.PayInsuranceClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getPayInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserInsuranceClaimsResponse getUserInsuranceClaims(com.game_engine.commission.v1.GetUserInsuranceClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserInsuranceClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse getInsuranceClaimsByStatus(com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetInsuranceClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.game_engine.commission.v1.GetUserSettlementsResponse getUserSettlements(com.game_engine.commission.v1.GetUserSettlementsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserSettlementsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementsByStatusResponse getSettlementsByStatus(com.game_engine.commission.v1.GetSettlementsByStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSettlementsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementsByTypeResponse getSettlementsByType(com.game_engine.commission.v1.GetSettlementsByTypeRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSettlementsByTypeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementByIdResponse getSettlementById(com.game_engine.commission.v1.GetSettlementByIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSettlementByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserTotalPendingResponse getUserTotalPending(com.game_engine.commission.v1.GetUserTotalPendingRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserTotalPendingMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserTotalSettledResponse getUserTotalSettled(com.game_engine.commission.v1.GetUserTotalSettledRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserTotalSettledMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service ClaimService.
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public static final class ClaimServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<ClaimServiceBlockingStub> {
    private ClaimServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClaimServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClaimServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission Claims
     * </pre>
     */
    public com.game_engine.commission.v1.SubmitCommissionClaimResponse submitCommissionClaim(com.game_engine.commission.v1.SubmitCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserCommissionClaimsResponse getUserCommissionClaims(com.game_engine.commission.v1.GetUserCommissionClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserCommissionClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse getCommissionClaimsByStatus(com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionClaimByIdResponse getCommissionClaimById(com.game_engine.commission.v1.GetCommissionClaimByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionClaimByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ApproveCommissionClaimResponse approveCommissionClaim(com.game_engine.commission.v1.ApproveCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApproveCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectCommissionClaimResponse rejectCommissionClaim(com.game_engine.commission.v1.RejectCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRejectCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayCommissionClaimResponse payCommissionClaim(com.game_engine.commission.v1.PayCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPayCommissionClaimMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Rebet Claims
     * </pre>
     */
    public com.game_engine.commission.v1.CreateRebetClaimResponse createRebetClaim(com.game_engine.commission.v1.CreateRebetClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateRebetClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.UpdateRebetProgressResponse updateRebetProgress(com.game_engine.commission.v1.UpdateRebetProgressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateRebetProgressMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ClaimRebetResponse claimRebet(com.game_engine.commission.v1.ClaimRebetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimRebetMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserRebetClaimsResponse getUserRebetClaims(com.game_engine.commission.v1.GetUserRebetClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserRebetClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetClaimableRebetsResponse getClaimableRebets(com.game_engine.commission.v1.GetClaimableRebetsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClaimableRebetsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Insurance Claims
     * </pre>
     */
    public com.game_engine.commission.v1.SubmitInsuranceClaimResponse submitInsuranceClaim(com.game_engine.commission.v1.SubmitInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ApproveInsuranceClaimResponse approveInsuranceClaim(com.game_engine.commission.v1.ApproveInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApproveInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectInsuranceClaimResponse rejectInsuranceClaim(com.game_engine.commission.v1.RejectInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRejectInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayInsuranceClaimResponse payInsuranceClaim(com.game_engine.commission.v1.PayInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPayInsuranceClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserInsuranceClaimsResponse getUserInsuranceClaims(com.game_engine.commission.v1.GetUserInsuranceClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserInsuranceClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse getInsuranceClaimsByStatus(com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetInsuranceClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.game_engine.commission.v1.GetUserSettlementsResponse getUserSettlements(com.game_engine.commission.v1.GetUserSettlementsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserSettlementsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementsByStatusResponse getSettlementsByStatus(com.game_engine.commission.v1.GetSettlementsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSettlementsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementsByTypeResponse getSettlementsByType(com.game_engine.commission.v1.GetSettlementsByTypeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSettlementsByTypeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetSettlementByIdResponse getSettlementById(com.game_engine.commission.v1.GetSettlementByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSettlementByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserTotalPendingResponse getUserTotalPending(com.game_engine.commission.v1.GetUserTotalPendingRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserTotalPendingMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserTotalSettledResponse getUserTotalSettled(com.game_engine.commission.v1.GetUserTotalSettledRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserTotalSettledMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service ClaimService.
   * <pre>
   * =============================================================================
   * Claim Service - handles commission claims, rebet claims, insurance claims, settlements
   * =============================================================================
   * </pre>
   */
  public static final class ClaimServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<ClaimServiceFutureStub> {
    private ClaimServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ClaimServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ClaimServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Commission Claims
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.SubmitCommissionClaimResponse> submitCommissionClaim(
        com.game_engine.commission.v1.SubmitCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitCommissionClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserCommissionClaimsResponse> getUserCommissionClaims(
        com.game_engine.commission.v1.GetUserCommissionClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserCommissionClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse> getCommissionClaimsByStatus(
        com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionClaimsByStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionClaimByIdResponse> getCommissionClaimById(
        com.game_engine.commission.v1.GetCommissionClaimByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionClaimByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ApproveCommissionClaimResponse> approveCommissionClaim(
        com.game_engine.commission.v1.ApproveCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApproveCommissionClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.RejectCommissionClaimResponse> rejectCommissionClaim(
        com.game_engine.commission.v1.RejectCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRejectCommissionClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.PayCommissionClaimResponse> payCommissionClaim(
        com.game_engine.commission.v1.PayCommissionClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPayCommissionClaimMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Rebet Claims
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.CreateRebetClaimResponse> createRebetClaim(
        com.game_engine.commission.v1.CreateRebetClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateRebetClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.UpdateRebetProgressResponse> updateRebetProgress(
        com.game_engine.commission.v1.UpdateRebetProgressRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateRebetProgressMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ClaimRebetResponse> claimRebet(
        com.game_engine.commission.v1.ClaimRebetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimRebetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserRebetClaimsResponse> getUserRebetClaims(
        com.game_engine.commission.v1.GetUserRebetClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserRebetClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetClaimableRebetsResponse> getClaimableRebets(
        com.game_engine.commission.v1.GetClaimableRebetsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClaimableRebetsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Insurance Claims
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.SubmitInsuranceClaimResponse> submitInsuranceClaim(
        com.game_engine.commission.v1.SubmitInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitInsuranceClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ApproveInsuranceClaimResponse> approveInsuranceClaim(
        com.game_engine.commission.v1.ApproveInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApproveInsuranceClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.RejectInsuranceClaimResponse> rejectInsuranceClaim(
        com.game_engine.commission.v1.RejectInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRejectInsuranceClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.PayInsuranceClaimResponse> payInsuranceClaim(
        com.game_engine.commission.v1.PayInsuranceClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPayInsuranceClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserInsuranceClaimsResponse> getUserInsuranceClaims(
        com.game_engine.commission.v1.GetUserInsuranceClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserInsuranceClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse> getInsuranceClaimsByStatus(
        com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetInsuranceClaimsByStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Settlements
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserSettlementsResponse> getUserSettlements(
        com.game_engine.commission.v1.GetUserSettlementsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserSettlementsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetSettlementsByStatusResponse> getSettlementsByStatus(
        com.game_engine.commission.v1.GetSettlementsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSettlementsByStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetSettlementsByTypeResponse> getSettlementsByType(
        com.game_engine.commission.v1.GetSettlementsByTypeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSettlementsByTypeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetSettlementByIdResponse> getSettlementById(
        com.game_engine.commission.v1.GetSettlementByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSettlementByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserTotalPendingResponse> getUserTotalPending(
        com.game_engine.commission.v1.GetUserTotalPendingRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserTotalPendingMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserTotalSettledResponse> getUserTotalSettled(
        com.game_engine.commission.v1.GetUserTotalSettledRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserTotalSettledMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SUBMIT_COMMISSION_CLAIM = 0;
  private static final int METHODID_GET_USER_COMMISSION_CLAIMS = 1;
  private static final int METHODID_GET_COMMISSION_CLAIMS_BY_STATUS = 2;
  private static final int METHODID_GET_COMMISSION_CLAIM_BY_ID = 3;
  private static final int METHODID_APPROVE_COMMISSION_CLAIM = 4;
  private static final int METHODID_REJECT_COMMISSION_CLAIM = 5;
  private static final int METHODID_PAY_COMMISSION_CLAIM = 6;
  private static final int METHODID_CREATE_REBET_CLAIM = 7;
  private static final int METHODID_UPDATE_REBET_PROGRESS = 8;
  private static final int METHODID_CLAIM_REBET = 9;
  private static final int METHODID_GET_USER_REBET_CLAIMS = 10;
  private static final int METHODID_GET_CLAIMABLE_REBETS = 11;
  private static final int METHODID_SUBMIT_INSURANCE_CLAIM = 12;
  private static final int METHODID_APPROVE_INSURANCE_CLAIM = 13;
  private static final int METHODID_REJECT_INSURANCE_CLAIM = 14;
  private static final int METHODID_PAY_INSURANCE_CLAIM = 15;
  private static final int METHODID_GET_USER_INSURANCE_CLAIMS = 16;
  private static final int METHODID_GET_INSURANCE_CLAIMS_BY_STATUS = 17;
  private static final int METHODID_GET_USER_SETTLEMENTS = 18;
  private static final int METHODID_GET_SETTLEMENTS_BY_STATUS = 19;
  private static final int METHODID_GET_SETTLEMENTS_BY_TYPE = 20;
  private static final int METHODID_GET_SETTLEMENT_BY_ID = 21;
  private static final int METHODID_GET_USER_TOTAL_PENDING = 22;
  private static final int METHODID_GET_USER_TOTAL_SETTLED = 23;

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
        case METHODID_SUBMIT_COMMISSION_CLAIM:
          serviceImpl.submitCommissionClaim((com.game_engine.commission.v1.SubmitCommissionClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitCommissionClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_COMMISSION_CLAIMS:
          serviceImpl.getUserCommissionClaims((com.game_engine.commission.v1.GetUserCommissionClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserCommissionClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_CLAIMS_BY_STATUS:
          serviceImpl.getCommissionClaimsByStatus((com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_CLAIM_BY_ID:
          serviceImpl.getCommissionClaimById((com.game_engine.commission.v1.GetCommissionClaimByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionClaimByIdResponse>) responseObserver);
          break;
        case METHODID_APPROVE_COMMISSION_CLAIM:
          serviceImpl.approveCommissionClaim((com.game_engine.commission.v1.ApproveCommissionClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionClaimResponse>) responseObserver);
          break;
        case METHODID_REJECT_COMMISSION_CLAIM:
          serviceImpl.rejectCommissionClaim((com.game_engine.commission.v1.RejectCommissionClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionClaimResponse>) responseObserver);
          break;
        case METHODID_PAY_COMMISSION_CLAIM:
          serviceImpl.payCommissionClaim((com.game_engine.commission.v1.PayCommissionClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionClaimResponse>) responseObserver);
          break;
        case METHODID_CREATE_REBET_CLAIM:
          serviceImpl.createRebetClaim((com.game_engine.commission.v1.CreateRebetClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateRebetClaimResponse>) responseObserver);
          break;
        case METHODID_UPDATE_REBET_PROGRESS:
          serviceImpl.updateRebetProgress((com.game_engine.commission.v1.UpdateRebetProgressRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateRebetProgressResponse>) responseObserver);
          break;
        case METHODID_CLAIM_REBET:
          serviceImpl.claimRebet((com.game_engine.commission.v1.ClaimRebetRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimRebetResponse>) responseObserver);
          break;
        case METHODID_GET_USER_REBET_CLAIMS:
          serviceImpl.getUserRebetClaims((com.game_engine.commission.v1.GetUserRebetClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserRebetClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_CLAIMABLE_REBETS:
          serviceImpl.getClaimableRebets((com.game_engine.commission.v1.GetClaimableRebetsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimableRebetsResponse>) responseObserver);
          break;
        case METHODID_SUBMIT_INSURANCE_CLAIM:
          serviceImpl.submitInsuranceClaim((com.game_engine.commission.v1.SubmitInsuranceClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitInsuranceClaimResponse>) responseObserver);
          break;
        case METHODID_APPROVE_INSURANCE_CLAIM:
          serviceImpl.approveInsuranceClaim((com.game_engine.commission.v1.ApproveInsuranceClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveInsuranceClaimResponse>) responseObserver);
          break;
        case METHODID_REJECT_INSURANCE_CLAIM:
          serviceImpl.rejectInsuranceClaim((com.game_engine.commission.v1.RejectInsuranceClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectInsuranceClaimResponse>) responseObserver);
          break;
        case METHODID_PAY_INSURANCE_CLAIM:
          serviceImpl.payInsuranceClaim((com.game_engine.commission.v1.PayInsuranceClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayInsuranceClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_INSURANCE_CLAIMS:
          serviceImpl.getUserInsuranceClaims((com.game_engine.commission.v1.GetUserInsuranceClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserInsuranceClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_INSURANCE_CLAIMS_BY_STATUS:
          serviceImpl.getInsuranceClaimsByStatus((com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse>) responseObserver);
          break;
        case METHODID_GET_USER_SETTLEMENTS:
          serviceImpl.getUserSettlements((com.game_engine.commission.v1.GetUserSettlementsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserSettlementsResponse>) responseObserver);
          break;
        case METHODID_GET_SETTLEMENTS_BY_STATUS:
          serviceImpl.getSettlementsByStatus((com.game_engine.commission.v1.GetSettlementsByStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByStatusResponse>) responseObserver);
          break;
        case METHODID_GET_SETTLEMENTS_BY_TYPE:
          serviceImpl.getSettlementsByType((com.game_engine.commission.v1.GetSettlementsByTypeRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementsByTypeResponse>) responseObserver);
          break;
        case METHODID_GET_SETTLEMENT_BY_ID:
          serviceImpl.getSettlementById((com.game_engine.commission.v1.GetSettlementByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetSettlementByIdResponse>) responseObserver);
          break;
        case METHODID_GET_USER_TOTAL_PENDING:
          serviceImpl.getUserTotalPending((com.game_engine.commission.v1.GetUserTotalPendingRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalPendingResponse>) responseObserver);
          break;
        case METHODID_GET_USER_TOTAL_SETTLED:
          serviceImpl.getUserTotalSettled((com.game_engine.commission.v1.GetUserTotalSettledRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserTotalSettledResponse>) responseObserver);
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
          getSubmitCommissionClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.SubmitCommissionClaimRequest,
              com.game_engine.commission.v1.SubmitCommissionClaimResponse>(
                service, METHODID_SUBMIT_COMMISSION_CLAIM)))
        .addMethod(
          getGetUserCommissionClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserCommissionClaimsRequest,
              com.game_engine.commission.v1.GetUserCommissionClaimsResponse>(
                service, METHODID_GET_USER_COMMISSION_CLAIMS)))
        .addMethod(
          getGetCommissionClaimsByStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionClaimsByStatusRequest,
              com.game_engine.commission.v1.GetCommissionClaimsByStatusResponse>(
                service, METHODID_GET_COMMISSION_CLAIMS_BY_STATUS)))
        .addMethod(
          getGetCommissionClaimByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionClaimByIdRequest,
              com.game_engine.commission.v1.GetCommissionClaimByIdResponse>(
                service, METHODID_GET_COMMISSION_CLAIM_BY_ID)))
        .addMethod(
          getApproveCommissionClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.ApproveCommissionClaimRequest,
              com.game_engine.commission.v1.ApproveCommissionClaimResponse>(
                service, METHODID_APPROVE_COMMISSION_CLAIM)))
        .addMethod(
          getRejectCommissionClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.RejectCommissionClaimRequest,
              com.game_engine.commission.v1.RejectCommissionClaimResponse>(
                service, METHODID_REJECT_COMMISSION_CLAIM)))
        .addMethod(
          getPayCommissionClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.PayCommissionClaimRequest,
              com.game_engine.commission.v1.PayCommissionClaimResponse>(
                service, METHODID_PAY_COMMISSION_CLAIM)))
        .addMethod(
          getCreateRebetClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.CreateRebetClaimRequest,
              com.game_engine.commission.v1.CreateRebetClaimResponse>(
                service, METHODID_CREATE_REBET_CLAIM)))
        .addMethod(
          getUpdateRebetProgressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.UpdateRebetProgressRequest,
              com.game_engine.commission.v1.UpdateRebetProgressResponse>(
                service, METHODID_UPDATE_REBET_PROGRESS)))
        .addMethod(
          getClaimRebetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.ClaimRebetRequest,
              com.game_engine.commission.v1.ClaimRebetResponse>(
                service, METHODID_CLAIM_REBET)))
        .addMethod(
          getGetUserRebetClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserRebetClaimsRequest,
              com.game_engine.commission.v1.GetUserRebetClaimsResponse>(
                service, METHODID_GET_USER_REBET_CLAIMS)))
        .addMethod(
          getGetClaimableRebetsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetClaimableRebetsRequest,
              com.game_engine.commission.v1.GetClaimableRebetsResponse>(
                service, METHODID_GET_CLAIMABLE_REBETS)))
        .addMethod(
          getSubmitInsuranceClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.SubmitInsuranceClaimRequest,
              com.game_engine.commission.v1.SubmitInsuranceClaimResponse>(
                service, METHODID_SUBMIT_INSURANCE_CLAIM)))
        .addMethod(
          getApproveInsuranceClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.ApproveInsuranceClaimRequest,
              com.game_engine.commission.v1.ApproveInsuranceClaimResponse>(
                service, METHODID_APPROVE_INSURANCE_CLAIM)))
        .addMethod(
          getRejectInsuranceClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.RejectInsuranceClaimRequest,
              com.game_engine.commission.v1.RejectInsuranceClaimResponse>(
                service, METHODID_REJECT_INSURANCE_CLAIM)))
        .addMethod(
          getPayInsuranceClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.PayInsuranceClaimRequest,
              com.game_engine.commission.v1.PayInsuranceClaimResponse>(
                service, METHODID_PAY_INSURANCE_CLAIM)))
        .addMethod(
          getGetUserInsuranceClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserInsuranceClaimsRequest,
              com.game_engine.commission.v1.GetUserInsuranceClaimsResponse>(
                service, METHODID_GET_USER_INSURANCE_CLAIMS)))
        .addMethod(
          getGetInsuranceClaimsByStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetInsuranceClaimsByStatusRequest,
              com.game_engine.commission.v1.GetInsuranceClaimsByStatusResponse>(
                service, METHODID_GET_INSURANCE_CLAIMS_BY_STATUS)))
        .addMethod(
          getGetUserSettlementsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserSettlementsRequest,
              com.game_engine.commission.v1.GetUserSettlementsResponse>(
                service, METHODID_GET_USER_SETTLEMENTS)))
        .addMethod(
          getGetSettlementsByStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetSettlementsByStatusRequest,
              com.game_engine.commission.v1.GetSettlementsByStatusResponse>(
                service, METHODID_GET_SETTLEMENTS_BY_STATUS)))
        .addMethod(
          getGetSettlementsByTypeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetSettlementsByTypeRequest,
              com.game_engine.commission.v1.GetSettlementsByTypeResponse>(
                service, METHODID_GET_SETTLEMENTS_BY_TYPE)))
        .addMethod(
          getGetSettlementByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetSettlementByIdRequest,
              com.game_engine.commission.v1.GetSettlementByIdResponse>(
                service, METHODID_GET_SETTLEMENT_BY_ID)))
        .addMethod(
          getGetUserTotalPendingMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserTotalPendingRequest,
              com.game_engine.commission.v1.GetUserTotalPendingResponse>(
                service, METHODID_GET_USER_TOTAL_PENDING)))
        .addMethod(
          getGetUserTotalSettledMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserTotalSettledRequest,
              com.game_engine.commission.v1.GetUserTotalSettledResponse>(
                service, METHODID_GET_USER_TOTAL_SETTLED)))
        .build();
  }

  private static abstract class ClaimServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ClaimServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.commission.v1.CommissionServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ClaimService");
    }
  }

  private static final class ClaimServiceFileDescriptorSupplier
      extends ClaimServiceBaseDescriptorSupplier {
    ClaimServiceFileDescriptorSupplier() {}
  }

  private static final class ClaimServiceMethodDescriptorSupplier
      extends ClaimServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    ClaimServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (ClaimServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ClaimServiceFileDescriptorSupplier())
              .addMethod(getSubmitCommissionClaimMethod())
              .addMethod(getGetUserCommissionClaimsMethod())
              .addMethod(getGetCommissionClaimsByStatusMethod())
              .addMethod(getGetCommissionClaimByIdMethod())
              .addMethod(getApproveCommissionClaimMethod())
              .addMethod(getRejectCommissionClaimMethod())
              .addMethod(getPayCommissionClaimMethod())
              .addMethod(getCreateRebetClaimMethod())
              .addMethod(getUpdateRebetProgressMethod())
              .addMethod(getClaimRebetMethod())
              .addMethod(getGetUserRebetClaimsMethod())
              .addMethod(getGetClaimableRebetsMethod())
              .addMethod(getSubmitInsuranceClaimMethod())
              .addMethod(getApproveInsuranceClaimMethod())
              .addMethod(getRejectInsuranceClaimMethod())
              .addMethod(getPayInsuranceClaimMethod())
              .addMethod(getGetUserInsuranceClaimsMethod())
              .addMethod(getGetInsuranceClaimsByStatusMethod())
              .addMethod(getGetUserSettlementsMethod())
              .addMethod(getGetSettlementsByStatusMethod())
              .addMethod(getGetSettlementsByTypeMethod())
              .addMethod(getGetSettlementByIdMethod())
              .addMethod(getGetUserTotalPendingMethod())
              .addMethod(getGetUserTotalSettledMethod())
              .build();
        }
      }
    }
    return result;
  }
}
