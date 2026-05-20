package com.game_engine.commission.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * =============================================================================
 * Commission Service - manages commission calculations and configs
 * =============================================================================
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class CommissionServiceGrpc {

  private CommissionServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.commission.v1.CommissionService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionRequest,
      com.game_engine.commission.v1.CreateCommissionResponse> getCreateCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateCommission",
      requestType = com.game_engine.commission.v1.CreateCommissionRequest.class,
      responseType = com.game_engine.commission.v1.CreateCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionRequest,
      com.game_engine.commission.v1.CreateCommissionResponse> getCreateCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionRequest, com.game_engine.commission.v1.CreateCommissionResponse> getCreateCommissionMethod;
    if ((getCreateCommissionMethod = CommissionServiceGrpc.getCreateCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getCreateCommissionMethod = CommissionServiceGrpc.getCreateCommissionMethod) == null) {
          CommissionServiceGrpc.getCreateCommissionMethod = getCreateCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.CreateCommissionRequest, com.game_engine.commission.v1.CreateCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("CreateCommission"))
              .build();
        }
      }
    }
    return getCreateCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByIdRequest,
      com.game_engine.commission.v1.GetCommissionByIdResponse> getGetCommissionByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionById",
      requestType = com.game_engine.commission.v1.GetCommissionByIdRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByIdRequest,
      com.game_engine.commission.v1.GetCommissionByIdResponse> getGetCommissionByIdMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByIdRequest, com.game_engine.commission.v1.GetCommissionByIdResponse> getGetCommissionByIdMethod;
    if ((getGetCommissionByIdMethod = CommissionServiceGrpc.getGetCommissionByIdMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionByIdMethod = CommissionServiceGrpc.getGetCommissionByIdMethod) == null) {
          CommissionServiceGrpc.getGetCommissionByIdMethod = getGetCommissionByIdMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionByIdRequest, com.game_engine.commission.v1.GetCommissionByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionById"))
              .build();
        }
      }
    }
    return getGetCommissionByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByAffiliateRequest,
      com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> getGetCommissionsByAffiliateMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionsByAffiliate",
      requestType = com.game_engine.commission.v1.GetCommissionsByAffiliateRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionsByAffiliateResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByAffiliateRequest,
      com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> getGetCommissionsByAffiliateMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByAffiliateRequest, com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> getGetCommissionsByAffiliateMethod;
    if ((getGetCommissionsByAffiliateMethod = CommissionServiceGrpc.getGetCommissionsByAffiliateMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionsByAffiliateMethod = CommissionServiceGrpc.getGetCommissionsByAffiliateMethod) == null) {
          CommissionServiceGrpc.getGetCommissionsByAffiliateMethod = getGetCommissionsByAffiliateMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionsByAffiliateRequest, com.game_engine.commission.v1.GetCommissionsByAffiliateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionsByAffiliate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByAffiliateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByAffiliateResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionsByAffiliate"))
              .build();
        }
      }
    }
    return getGetCommissionsByAffiliateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByMerchantRequest,
      com.game_engine.commission.v1.GetCommissionsByMerchantResponse> getGetCommissionsByMerchantMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionsByMerchant",
      requestType = com.game_engine.commission.v1.GetCommissionsByMerchantRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionsByMerchantResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByMerchantRequest,
      com.game_engine.commission.v1.GetCommissionsByMerchantResponse> getGetCommissionsByMerchantMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByMerchantRequest, com.game_engine.commission.v1.GetCommissionsByMerchantResponse> getGetCommissionsByMerchantMethod;
    if ((getGetCommissionsByMerchantMethod = CommissionServiceGrpc.getGetCommissionsByMerchantMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionsByMerchantMethod = CommissionServiceGrpc.getGetCommissionsByMerchantMethod) == null) {
          CommissionServiceGrpc.getGetCommissionsByMerchantMethod = getGetCommissionsByMerchantMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionsByMerchantRequest, com.game_engine.commission.v1.GetCommissionsByMerchantResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionsByMerchant"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByMerchantRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByMerchantResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionsByMerchant"))
              .build();
        }
      }
    }
    return getGetCommissionsByMerchantMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByPeriodRequest,
      com.game_engine.commission.v1.GetCommissionsByPeriodResponse> getGetCommissionsByPeriodMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionsByPeriod",
      requestType = com.game_engine.commission.v1.GetCommissionsByPeriodRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionsByPeriodResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByPeriodRequest,
      com.game_engine.commission.v1.GetCommissionsByPeriodResponse> getGetCommissionsByPeriodMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionsByPeriodRequest, com.game_engine.commission.v1.GetCommissionsByPeriodResponse> getGetCommissionsByPeriodMethod;
    if ((getGetCommissionsByPeriodMethod = CommissionServiceGrpc.getGetCommissionsByPeriodMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionsByPeriodMethod = CommissionServiceGrpc.getGetCommissionsByPeriodMethod) == null) {
          CommissionServiceGrpc.getGetCommissionsByPeriodMethod = getGetCommissionsByPeriodMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionsByPeriodRequest, com.game_engine.commission.v1.GetCommissionsByPeriodResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionsByPeriod"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByPeriodRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionsByPeriodResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionsByPeriod"))
              .build();
        }
      }
    }
    return getGetCommissionsByPeriodMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest,
      com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> getGetCommissionByAffiliateAndPeriodMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionByAffiliateAndPeriod",
      requestType = com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest,
      com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> getGetCommissionByAffiliateAndPeriodMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest, com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> getGetCommissionByAffiliateAndPeriodMethod;
    if ((getGetCommissionByAffiliateAndPeriodMethod = CommissionServiceGrpc.getGetCommissionByAffiliateAndPeriodMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionByAffiliateAndPeriodMethod = CommissionServiceGrpc.getGetCommissionByAffiliateAndPeriodMethod) == null) {
          CommissionServiceGrpc.getGetCommissionByAffiliateAndPeriodMethod = getGetCommissionByAffiliateAndPeriodMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest, com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionByAffiliateAndPeriod"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetCommissionByAffiliateAndPeriod"))
              .build();
        }
      }
    }
    return getGetCommissionByAffiliateAndPeriodMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPaidCommissionRequest,
      com.game_engine.commission.v1.GetTotalPaidCommissionResponse> getGetTotalPaidCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalPaidCommission",
      requestType = com.game_engine.commission.v1.GetTotalPaidCommissionRequest.class,
      responseType = com.game_engine.commission.v1.GetTotalPaidCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPaidCommissionRequest,
      com.game_engine.commission.v1.GetTotalPaidCommissionResponse> getGetTotalPaidCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPaidCommissionRequest, com.game_engine.commission.v1.GetTotalPaidCommissionResponse> getGetTotalPaidCommissionMethod;
    if ((getGetTotalPaidCommissionMethod = CommissionServiceGrpc.getGetTotalPaidCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetTotalPaidCommissionMethod = CommissionServiceGrpc.getGetTotalPaidCommissionMethod) == null) {
          CommissionServiceGrpc.getGetTotalPaidCommissionMethod = getGetTotalPaidCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetTotalPaidCommissionRequest, com.game_engine.commission.v1.GetTotalPaidCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalPaidCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalPaidCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalPaidCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetTotalPaidCommission"))
              .build();
        }
      }
    }
    return getGetTotalPaidCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPendingCommissionRequest,
      com.game_engine.commission.v1.GetTotalPendingCommissionResponse> getGetTotalPendingCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalPendingCommission",
      requestType = com.game_engine.commission.v1.GetTotalPendingCommissionRequest.class,
      responseType = com.game_engine.commission.v1.GetTotalPendingCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPendingCommissionRequest,
      com.game_engine.commission.v1.GetTotalPendingCommissionResponse> getGetTotalPendingCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalPendingCommissionRequest, com.game_engine.commission.v1.GetTotalPendingCommissionResponse> getGetTotalPendingCommissionMethod;
    if ((getGetTotalPendingCommissionMethod = CommissionServiceGrpc.getGetTotalPendingCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetTotalPendingCommissionMethod = CommissionServiceGrpc.getGetTotalPendingCommissionMethod) == null) {
          CommissionServiceGrpc.getGetTotalPendingCommissionMethod = getGetTotalPendingCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetTotalPendingCommissionRequest, com.game_engine.commission.v1.GetTotalPendingCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalPendingCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalPendingCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalPendingCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetTotalPendingCommission"))
              .build();
        }
      }
    }
    return getGetTotalPendingCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest,
      com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> getGetTotalRevenueByMerchantMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalRevenueByMerchant",
      requestType = com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest.class,
      responseType = com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest,
      com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> getGetTotalRevenueByMerchantMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest, com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> getGetTotalRevenueByMerchantMethod;
    if ((getGetTotalRevenueByMerchantMethod = CommissionServiceGrpc.getGetTotalRevenueByMerchantMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetTotalRevenueByMerchantMethod = CommissionServiceGrpc.getGetTotalRevenueByMerchantMethod) == null) {
          CommissionServiceGrpc.getGetTotalRevenueByMerchantMethod = getGetTotalRevenueByMerchantMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest, com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalRevenueByMerchant"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetTotalRevenueByMerchant"))
              .build();
        }
      }
    }
    return getGetTotalRevenueByMerchantMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateRevenueShareRequest,
      com.game_engine.commission.v1.CalculateRevenueShareResponse> getCalculateRevenueShareMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CalculateRevenueShare",
      requestType = com.game_engine.commission.v1.CalculateRevenueShareRequest.class,
      responseType = com.game_engine.commission.v1.CalculateRevenueShareResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateRevenueShareRequest,
      com.game_engine.commission.v1.CalculateRevenueShareResponse> getCalculateRevenueShareMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateRevenueShareRequest, com.game_engine.commission.v1.CalculateRevenueShareResponse> getCalculateRevenueShareMethod;
    if ((getCalculateRevenueShareMethod = CommissionServiceGrpc.getCalculateRevenueShareMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getCalculateRevenueShareMethod = CommissionServiceGrpc.getCalculateRevenueShareMethod) == null) {
          CommissionServiceGrpc.getCalculateRevenueShareMethod = getCalculateRevenueShareMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.CalculateRevenueShareRequest, com.game_engine.commission.v1.CalculateRevenueShareResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CalculateRevenueShare"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CalculateRevenueShareRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CalculateRevenueShareResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("CalculateRevenueShare"))
              .build();
        }
      }
    }
    return getCalculateRevenueShareMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateCPARequest,
      com.game_engine.commission.v1.CalculateCPAResponse> getCalculateCPAMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CalculateCPA",
      requestType = com.game_engine.commission.v1.CalculateCPARequest.class,
      responseType = com.game_engine.commission.v1.CalculateCPAResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateCPARequest,
      com.game_engine.commission.v1.CalculateCPAResponse> getCalculateCPAMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.CalculateCPARequest, com.game_engine.commission.v1.CalculateCPAResponse> getCalculateCPAMethod;
    if ((getCalculateCPAMethod = CommissionServiceGrpc.getCalculateCPAMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getCalculateCPAMethod = CommissionServiceGrpc.getCalculateCPAMethod) == null) {
          CommissionServiceGrpc.getCalculateCPAMethod = getCalculateCPAMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.CalculateCPARequest, com.game_engine.commission.v1.CalculateCPAResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CalculateCPA"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CalculateCPARequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CalculateCPAResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("CalculateCPA"))
              .build();
        }
      }
    }
    return getCalculateCPAMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionRequest,
      com.game_engine.commission.v1.ApproveCommissionResponse> getApproveCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApproveCommission",
      requestType = com.game_engine.commission.v1.ApproveCommissionRequest.class,
      responseType = com.game_engine.commission.v1.ApproveCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionRequest,
      com.game_engine.commission.v1.ApproveCommissionResponse> getApproveCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ApproveCommissionRequest, com.game_engine.commission.v1.ApproveCommissionResponse> getApproveCommissionMethod;
    if ((getApproveCommissionMethod = CommissionServiceGrpc.getApproveCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getApproveCommissionMethod = CommissionServiceGrpc.getApproveCommissionMethod) == null) {
          CommissionServiceGrpc.getApproveCommissionMethod = getApproveCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ApproveCommissionRequest, com.game_engine.commission.v1.ApproveCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApproveCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ApproveCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("ApproveCommission"))
              .build();
        }
      }
    }
    return getApproveCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionRequest,
      com.game_engine.commission.v1.RejectCommissionResponse> getRejectCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RejectCommission",
      requestType = com.game_engine.commission.v1.RejectCommissionRequest.class,
      responseType = com.game_engine.commission.v1.RejectCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionRequest,
      com.game_engine.commission.v1.RejectCommissionResponse> getRejectCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.RejectCommissionRequest, com.game_engine.commission.v1.RejectCommissionResponse> getRejectCommissionMethod;
    if ((getRejectCommissionMethod = CommissionServiceGrpc.getRejectCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getRejectCommissionMethod = CommissionServiceGrpc.getRejectCommissionMethod) == null) {
          CommissionServiceGrpc.getRejectCommissionMethod = getRejectCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.RejectCommissionRequest, com.game_engine.commission.v1.RejectCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RejectCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.RejectCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("RejectCommission"))
              .build();
        }
      }
    }
    return getRejectCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionRequest,
      com.game_engine.commission.v1.PayCommissionResponse> getPayCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PayCommission",
      requestType = com.game_engine.commission.v1.PayCommissionRequest.class,
      responseType = com.game_engine.commission.v1.PayCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionRequest,
      com.game_engine.commission.v1.PayCommissionResponse> getPayCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.PayCommissionRequest, com.game_engine.commission.v1.PayCommissionResponse> getPayCommissionMethod;
    if ((getPayCommissionMethod = CommissionServiceGrpc.getPayCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getPayCommissionMethod = CommissionServiceGrpc.getPayCommissionMethod) == null) {
          CommissionServiceGrpc.getPayCommissionMethod = getPayCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.PayCommissionRequest, com.game_engine.commission.v1.PayCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PayCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.PayCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("PayCommission"))
              .build();
        }
      }
    }
    return getPayCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetPendingCommissionsRequest,
      com.game_engine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPendingCommissions",
      requestType = com.game_engine.commission.v1.GetPendingCommissionsRequest.class,
      responseType = com.game_engine.commission.v1.GetPendingCommissionsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetPendingCommissionsRequest,
      com.game_engine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetPendingCommissionsRequest, com.game_engine.commission.v1.GetPendingCommissionsResponse> getGetPendingCommissionsMethod;
    if ((getGetPendingCommissionsMethod = CommissionServiceGrpc.getGetPendingCommissionsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetPendingCommissionsMethod = CommissionServiceGrpc.getGetPendingCommissionsMethod) == null) {
          CommissionServiceGrpc.getGetPendingCommissionsMethod = getGetPendingCommissionsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetPendingCommissionsRequest, com.game_engine.commission.v1.GetPendingCommissionsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPendingCommissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetPendingCommissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetPendingCommissionsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetPendingCommissions"))
              .build();
        }
      }
    }
    return getGetPendingCommissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllCommissionsRequest,
      com.game_engine.commission.v1.GetAllCommissionsResponse> getGetAllCommissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAllCommissions",
      requestType = com.game_engine.commission.v1.GetAllCommissionsRequest.class,
      responseType = com.game_engine.commission.v1.GetAllCommissionsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllCommissionsRequest,
      com.game_engine.commission.v1.GetAllCommissionsResponse> getGetAllCommissionsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllCommissionsRequest, com.game_engine.commission.v1.GetAllCommissionsResponse> getGetAllCommissionsMethod;
    if ((getGetAllCommissionsMethod = CommissionServiceGrpc.getGetAllCommissionsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetAllCommissionsMethod = CommissionServiceGrpc.getGetAllCommissionsMethod) == null) {
          CommissionServiceGrpc.getGetAllCommissionsMethod = getGetAllCommissionsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetAllCommissionsRequest, com.game_engine.commission.v1.GetAllCommissionsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllCommissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAllCommissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAllCommissionsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetAllCommissions"))
              .build();
        }
      }
    }
    return getGetAllCommissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteCommissionRequest,
      com.game_engine.commission.v1.DeleteCommissionResponse> getDeleteCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteCommission",
      requestType = com.game_engine.commission.v1.DeleteCommissionRequest.class,
      responseType = com.game_engine.commission.v1.DeleteCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteCommissionRequest,
      com.game_engine.commission.v1.DeleteCommissionResponse> getDeleteCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteCommissionRequest, com.game_engine.commission.v1.DeleteCommissionResponse> getDeleteCommissionMethod;
    if ((getDeleteCommissionMethod = CommissionServiceGrpc.getDeleteCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getDeleteCommissionMethod = CommissionServiceGrpc.getDeleteCommissionMethod) == null) {
          CommissionServiceGrpc.getDeleteCommissionMethod = getDeleteCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.DeleteCommissionRequest, com.game_engine.commission.v1.DeleteCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeleteCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeleteCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("DeleteCommission"))
              .build();
        }
      }
    }
    return getDeleteCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitClaimRequest,
      com.game_engine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitClaim",
      requestType = com.game_engine.commission.v1.SubmitClaimRequest.class,
      responseType = com.game_engine.commission.v1.SubmitClaimResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitClaimRequest,
      com.game_engine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.SubmitClaimRequest, com.game_engine.commission.v1.SubmitClaimResponse> getSubmitClaimMethod;
    if ((getSubmitClaimMethod = CommissionServiceGrpc.getSubmitClaimMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getSubmitClaimMethod = CommissionServiceGrpc.getSubmitClaimMethod) == null) {
          CommissionServiceGrpc.getSubmitClaimMethod = getSubmitClaimMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.SubmitClaimRequest, com.game_engine.commission.v1.SubmitClaimResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitClaim"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitClaimRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.SubmitClaimResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("SubmitClaim"))
              .build();
        }
      }
    }
    return getSubmitClaimMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserClaimsRequest,
      com.game_engine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserClaims",
      requestType = com.game_engine.commission.v1.GetUserClaimsRequest.class,
      responseType = com.game_engine.commission.v1.GetUserClaimsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserClaimsRequest,
      com.game_engine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetUserClaimsRequest, com.game_engine.commission.v1.GetUserClaimsResponse> getGetUserClaimsMethod;
    if ((getGetUserClaimsMethod = CommissionServiceGrpc.getGetUserClaimsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetUserClaimsMethod = CommissionServiceGrpc.getGetUserClaimsMethod) == null) {
          CommissionServiceGrpc.getGetUserClaimsMethod = getGetUserClaimsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetUserClaimsRequest, com.game_engine.commission.v1.GetUserClaimsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserClaims"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserClaimsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetUserClaimsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetUserClaims"))
              .build();
        }
      }
    }
    return getGetUserClaimsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimsByStatusRequest,
      com.game_engine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClaimsByStatus",
      requestType = com.game_engine.commission.v1.GetClaimsByStatusRequest.class,
      responseType = com.game_engine.commission.v1.GetClaimsByStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimsByStatusRequest,
      com.game_engine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetClaimsByStatusRequest, com.game_engine.commission.v1.GetClaimsByStatusResponse> getGetClaimsByStatusMethod;
    if ((getGetClaimsByStatusMethod = CommissionServiceGrpc.getGetClaimsByStatusMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetClaimsByStatusMethod = CommissionServiceGrpc.getGetClaimsByStatusMethod) == null) {
          CommissionServiceGrpc.getGetClaimsByStatusMethod = getGetClaimsByStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetClaimsByStatusRequest, com.game_engine.commission.v1.GetClaimsByStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClaimsByStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetClaimsByStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetClaimsByStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetClaimsByStatus"))
              .build();
        }
      }
    }
    return getGetClaimsByStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimCommissionRequest,
      com.game_engine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ClaimCommission",
      requestType = com.game_engine.commission.v1.ClaimCommissionRequest.class,
      responseType = com.game_engine.commission.v1.ClaimCommissionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimCommissionRequest,
      com.game_engine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ClaimCommissionRequest, com.game_engine.commission.v1.ClaimCommissionResponse> getClaimCommissionMethod;
    if ((getClaimCommissionMethod = CommissionServiceGrpc.getClaimCommissionMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getClaimCommissionMethod = CommissionServiceGrpc.getClaimCommissionMethod) == null) {
          CommissionServiceGrpc.getClaimCommissionMethod = getClaimCommissionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ClaimCommissionRequest, com.game_engine.commission.v1.ClaimCommissionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ClaimCommission"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ClaimCommissionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ClaimCommissionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("ClaimCommission"))
              .build();
        }
      }
    }
    return getClaimCommissionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAgentCommissionsRequest,
      com.game_engine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAgentCommissions",
      requestType = com.game_engine.commission.v1.GetAgentCommissionsRequest.class,
      responseType = com.game_engine.commission.v1.GetAgentCommissionsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAgentCommissionsRequest,
      com.game_engine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAgentCommissionsRequest, com.game_engine.commission.v1.GetAgentCommissionsResponse> getGetAgentCommissionsMethod;
    if ((getGetAgentCommissionsMethod = CommissionServiceGrpc.getGetAgentCommissionsMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetAgentCommissionsMethod = CommissionServiceGrpc.getGetAgentCommissionsMethod) == null) {
          CommissionServiceGrpc.getGetAgentCommissionsMethod = getGetAgentCommissionsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetAgentCommissionsRequest, com.game_engine.commission.v1.GetAgentCommissionsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAgentCommissions"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAgentCommissionsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAgentCommissionsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionServiceMethodDescriptorSupplier("GetAgentCommissions"))
              .build();
        }
      }
    }
    return getGetAgentCommissionsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionHistoryRequest,
      com.game_engine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCommissionHistory",
      requestType = com.game_engine.commission.v1.GetCommissionHistoryRequest.class,
      responseType = com.game_engine.commission.v1.GetCommissionHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionHistoryRequest,
      com.game_engine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetCommissionHistoryRequest, com.game_engine.commission.v1.GetCommissionHistoryResponse> getGetCommissionHistoryMethod;
    if ((getGetCommissionHistoryMethod = CommissionServiceGrpc.getGetCommissionHistoryMethod) == null) {
      synchronized (CommissionServiceGrpc.class) {
        if ((getGetCommissionHistoryMethod = CommissionServiceGrpc.getGetCommissionHistoryMethod) == null) {
          CommissionServiceGrpc.getGetCommissionHistoryMethod = getGetCommissionHistoryMethod =
              io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetCommissionHistoryRequest, com.game_engine.commission.v1.GetCommissionHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCommissionHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetCommissionHistoryResponse.getDefaultInstance()))
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
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Commission operations
     * </pre>
     */
    default void createCommission(com.game_engine.commission.v1.CreateCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateCommissionMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionById(com.game_engine.commission.v1.GetCommissionByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionByIdMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionsByAffiliate(com.game_engine.commission.v1.GetCommissionsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionsByAffiliateMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionsByMerchant(com.game_engine.commission.v1.GetCommissionsByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByMerchantResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionsByMerchantMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionsByPeriod(com.game_engine.commission.v1.GetCommissionsByPeriodRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByPeriodResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionsByPeriodMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionByAffiliateAndPeriod(com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionByAffiliateAndPeriodMethod(), responseObserver);
    }

    /**
     */
    default void getTotalPaidCommission(com.game_engine.commission.v1.GetTotalPaidCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPaidCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalPaidCommissionMethod(), responseObserver);
    }

    /**
     */
    default void getTotalPendingCommission(com.game_engine.commission.v1.GetTotalPendingCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPendingCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalPendingCommissionMethod(), responseObserver);
    }

    /**
     */
    default void getTotalRevenueByMerchant(com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalRevenueByMerchantMethod(), responseObserver);
    }

    /**
     * <pre>
     * Calculation
     * </pre>
     */
    default void calculateRevenueShare(com.game_engine.commission.v1.CalculateRevenueShareRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateRevenueShareResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCalculateRevenueShareMethod(), responseObserver);
    }

    /**
     */
    default void calculateCPA(com.game_engine.commission.v1.CalculateCPARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateCPAResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCalculateCPAMethod(), responseObserver);
    }

    /**
     * <pre>
     * Approval and payment
     * </pre>
     */
    default void approveCommission(com.game_engine.commission.v1.ApproveCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApproveCommissionMethod(), responseObserver);
    }

    /**
     */
    default void rejectCommission(com.game_engine.commission.v1.RejectCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRejectCommissionMethod(), responseObserver);
    }

    /**
     */
    default void payCommission(com.game_engine.commission.v1.PayCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPayCommissionMethod(), responseObserver);
    }

    /**
     * <pre>
     * Queries
     * </pre>
     */
    default void getPendingCommissions(com.game_engine.commission.v1.GetPendingCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetPendingCommissionsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPendingCommissionsMethod(), responseObserver);
    }

    /**
     */
    default void getAllCommissions(com.game_engine.commission.v1.GetAllCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllCommissionsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllCommissionsMethod(), responseObserver);
    }

    /**
     */
    default void deleteCommission(com.game_engine.commission.v1.DeleteCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteCommissionMethod(), responseObserver);
    }

    /**
     */
    default void submitClaim(com.game_engine.commission.v1.SubmitClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitClaimResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitClaimMethod(), responseObserver);
    }

    /**
     */
    default void getUserClaims(com.game_engine.commission.v1.GetUserClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserClaimsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserClaimsMethod(), responseObserver);
    }

    /**
     */
    default void getClaimsByStatus(com.game_engine.commission.v1.GetClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClaimsByStatusMethod(), responseObserver);
    }

    /**
     */
    default void claimCommission(com.game_engine.commission.v1.ClaimCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimCommissionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getClaimCommissionMethod(), responseObserver);
    }

    /**
     */
    default void getAgentCommissions(com.game_engine.commission.v1.GetAgentCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAgentCommissionsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAgentCommissionsMethod(), responseObserver);
    }

    /**
     */
    default void getCommissionHistory(com.game_engine.commission.v1.GetCommissionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCommissionHistoryMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service CommissionService.
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
   */
  public static abstract class CommissionServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return CommissionServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service CommissionService.
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
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
     * Commission operations
     * </pre>
     */
    public void createCommission(com.game_engine.commission.v1.CreateCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionById(com.game_engine.commission.v1.GetCommissionByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionsByAffiliate(com.game_engine.commission.v1.GetCommissionsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionsByAffiliateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionsByMerchant(com.game_engine.commission.v1.GetCommissionsByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByMerchantResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionsByMerchantMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionsByPeriod(com.game_engine.commission.v1.GetCommissionsByPeriodRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByPeriodResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionsByPeriodMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionByAffiliateAndPeriod(com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionByAffiliateAndPeriodMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTotalPaidCommission(com.game_engine.commission.v1.GetTotalPaidCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPaidCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalPaidCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTotalPendingCommission(com.game_engine.commission.v1.GetTotalPendingCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPendingCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalPendingCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTotalRevenueByMerchant(com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalRevenueByMerchantMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Calculation
     * </pre>
     */
    public void calculateRevenueShare(com.game_engine.commission.v1.CalculateRevenueShareRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateRevenueShareResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCalculateRevenueShareMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void calculateCPA(com.game_engine.commission.v1.CalculateCPARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateCPAResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCalculateCPAMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Approval and payment
     * </pre>
     */
    public void approveCommission(com.game_engine.commission.v1.ApproveCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApproveCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void rejectCommission(com.game_engine.commission.v1.RejectCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRejectCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void payCommission(com.game_engine.commission.v1.PayCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPayCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Queries
     * </pre>
     */
    public void getPendingCommissions(com.game_engine.commission.v1.GetPendingCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetPendingCommissionsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPendingCommissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllCommissions(com.game_engine.commission.v1.GetAllCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllCommissionsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllCommissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteCommission(com.game_engine.commission.v1.DeleteCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void submitClaim(com.game_engine.commission.v1.SubmitClaimRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitClaimResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitClaimMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserClaims(com.game_engine.commission.v1.GetUserClaimsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserClaimsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserClaimsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getClaimsByStatus(com.game_engine.commission.v1.GetClaimsByStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimsByStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClaimsByStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void claimCommission(com.game_engine.commission.v1.ClaimCommissionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimCommissionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getClaimCommissionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAgentCommissions(com.game_engine.commission.v1.GetAgentCommissionsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAgentCommissionsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAgentCommissionsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCommissionHistory(com.game_engine.commission.v1.GetCommissionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCommissionHistoryMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service CommissionService.
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
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
     * Commission operations
     * </pre>
     */
    public com.game_engine.commission.v1.CreateCommissionResponse createCommission(com.game_engine.commission.v1.CreateCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionByIdResponse getCommissionById(com.game_engine.commission.v1.GetCommissionByIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByAffiliateResponse getCommissionsByAffiliate(com.game_engine.commission.v1.GetCommissionsByAffiliateRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByMerchantResponse getCommissionsByMerchant(com.game_engine.commission.v1.GetCommissionsByMerchantRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionsByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByPeriodResponse getCommissionsByPeriod(com.game_engine.commission.v1.GetCommissionsByPeriodRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionsByPeriodMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse getCommissionByAffiliateAndPeriod(com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionByAffiliateAndPeriodMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalPaidCommissionResponse getTotalPaidCommission(com.game_engine.commission.v1.GetTotalPaidCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTotalPaidCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalPendingCommissionResponse getTotalPendingCommission(com.game_engine.commission.v1.GetTotalPendingCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTotalPendingCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse getTotalRevenueByMerchant(com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTotalRevenueByMerchantMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Calculation
     * </pre>
     */
    public com.game_engine.commission.v1.CalculateRevenueShareResponse calculateRevenueShare(com.game_engine.commission.v1.CalculateRevenueShareRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCalculateRevenueShareMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.CalculateCPAResponse calculateCPA(com.game_engine.commission.v1.CalculateCPARequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCalculateCPAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Approval and payment
     * </pre>
     */
    public com.game_engine.commission.v1.ApproveCommissionResponse approveCommission(com.game_engine.commission.v1.ApproveCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApproveCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectCommissionResponse rejectCommission(com.game_engine.commission.v1.RejectCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRejectCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayCommissionResponse payCommission(com.game_engine.commission.v1.PayCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getPayCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Queries
     * </pre>
     */
    public com.game_engine.commission.v1.GetPendingCommissionsResponse getPendingCommissions(com.game_engine.commission.v1.GetPendingCommissionsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPendingCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAllCommissionsResponse getAllCommissions(com.game_engine.commission.v1.GetAllCommissionsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAllCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeleteCommissionResponse deleteCommission(com.game_engine.commission.v1.DeleteCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getDeleteCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.SubmitClaimResponse submitClaim(com.game_engine.commission.v1.SubmitClaimRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserClaimsResponse getUserClaims(com.game_engine.commission.v1.GetUserClaimsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetClaimsByStatusResponse getClaimsByStatus(com.game_engine.commission.v1.GetClaimsByStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ClaimCommissionResponse claimCommission(com.game_engine.commission.v1.ClaimCommissionRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getClaimCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAgentCommissionsResponse getAgentCommissions(com.game_engine.commission.v1.GetAgentCommissionsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAgentCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionHistoryResponse getCommissionHistory(com.game_engine.commission.v1.GetCommissionHistoryRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCommissionHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service CommissionService.
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
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
     * Commission operations
     * </pre>
     */
    public com.game_engine.commission.v1.CreateCommissionResponse createCommission(com.game_engine.commission.v1.CreateCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionByIdResponse getCommissionById(com.game_engine.commission.v1.GetCommissionByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByAffiliateResponse getCommissionsByAffiliate(com.game_engine.commission.v1.GetCommissionsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByMerchantResponse getCommissionsByMerchant(com.game_engine.commission.v1.GetCommissionsByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionsByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionsByPeriodResponse getCommissionsByPeriod(com.game_engine.commission.v1.GetCommissionsByPeriodRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionsByPeriodMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse getCommissionByAffiliateAndPeriod(com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionByAffiliateAndPeriodMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalPaidCommissionResponse getTotalPaidCommission(com.game_engine.commission.v1.GetTotalPaidCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalPaidCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalPendingCommissionResponse getTotalPendingCommission(com.game_engine.commission.v1.GetTotalPendingCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalPendingCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse getTotalRevenueByMerchant(com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalRevenueByMerchantMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Calculation
     * </pre>
     */
    public com.game_engine.commission.v1.CalculateRevenueShareResponse calculateRevenueShare(com.game_engine.commission.v1.CalculateRevenueShareRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCalculateRevenueShareMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.CalculateCPAResponse calculateCPA(com.game_engine.commission.v1.CalculateCPARequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCalculateCPAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Approval and payment
     * </pre>
     */
    public com.game_engine.commission.v1.ApproveCommissionResponse approveCommission(com.game_engine.commission.v1.ApproveCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApproveCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.RejectCommissionResponse rejectCommission(com.game_engine.commission.v1.RejectCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRejectCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.PayCommissionResponse payCommission(com.game_engine.commission.v1.PayCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPayCommissionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Queries
     * </pre>
     */
    public com.game_engine.commission.v1.GetPendingCommissionsResponse getPendingCommissions(com.game_engine.commission.v1.GetPendingCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPendingCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAllCommissionsResponse getAllCommissions(com.game_engine.commission.v1.GetAllCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeleteCommissionResponse deleteCommission(com.game_engine.commission.v1.DeleteCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.SubmitClaimResponse submitClaim(com.game_engine.commission.v1.SubmitClaimRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitClaimMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetUserClaimsResponse getUserClaims(com.game_engine.commission.v1.GetUserClaimsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserClaimsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetClaimsByStatusResponse getClaimsByStatus(com.game_engine.commission.v1.GetClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClaimsByStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ClaimCommissionResponse claimCommission(com.game_engine.commission.v1.ClaimCommissionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getClaimCommissionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAgentCommissionsResponse getAgentCommissions(com.game_engine.commission.v1.GetAgentCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAgentCommissionsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetCommissionHistoryResponse getCommissionHistory(com.game_engine.commission.v1.GetCommissionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCommissionHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service CommissionService.
   * <pre>
   * =============================================================================
   * Commission Service - manages commission calculations and configs
   * =============================================================================
   * </pre>
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
     * Commission operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.CreateCommissionResponse> createCommission(
        com.game_engine.commission.v1.CreateCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionByIdResponse> getCommissionById(
        com.game_engine.commission.v1.GetCommissionByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionsByAffiliateResponse> getCommissionsByAffiliate(
        com.game_engine.commission.v1.GetCommissionsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionsByAffiliateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionsByMerchantResponse> getCommissionsByMerchant(
        com.game_engine.commission.v1.GetCommissionsByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionsByMerchantMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionsByPeriodResponse> getCommissionsByPeriod(
        com.game_engine.commission.v1.GetCommissionsByPeriodRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionsByPeriodMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse> getCommissionByAffiliateAndPeriod(
        com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionByAffiliateAndPeriodMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetTotalPaidCommissionResponse> getTotalPaidCommission(
        com.game_engine.commission.v1.GetTotalPaidCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalPaidCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetTotalPendingCommissionResponse> getTotalPendingCommission(
        com.game_engine.commission.v1.GetTotalPendingCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalPendingCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse> getTotalRevenueByMerchant(
        com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalRevenueByMerchantMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Calculation
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.CalculateRevenueShareResponse> calculateRevenueShare(
        com.game_engine.commission.v1.CalculateRevenueShareRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCalculateRevenueShareMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.CalculateCPAResponse> calculateCPA(
        com.game_engine.commission.v1.CalculateCPARequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCalculateCPAMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Approval and payment
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ApproveCommissionResponse> approveCommission(
        com.game_engine.commission.v1.ApproveCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApproveCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.RejectCommissionResponse> rejectCommission(
        com.game_engine.commission.v1.RejectCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRejectCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.PayCommissionResponse> payCommission(
        com.game_engine.commission.v1.PayCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPayCommissionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Queries
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetPendingCommissionsResponse> getPendingCommissions(
        com.game_engine.commission.v1.GetPendingCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPendingCommissionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetAllCommissionsResponse> getAllCommissions(
        com.game_engine.commission.v1.GetAllCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllCommissionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.DeleteCommissionResponse> deleteCommission(
        com.game_engine.commission.v1.DeleteCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.SubmitClaimResponse> submitClaim(
        com.game_engine.commission.v1.SubmitClaimRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitClaimMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetUserClaimsResponse> getUserClaims(
        com.game_engine.commission.v1.GetUserClaimsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserClaimsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetClaimsByStatusResponse> getClaimsByStatus(
        com.game_engine.commission.v1.GetClaimsByStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClaimsByStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ClaimCommissionResponse> claimCommission(
        com.game_engine.commission.v1.ClaimCommissionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getClaimCommissionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetAgentCommissionsResponse> getAgentCommissions(
        com.game_engine.commission.v1.GetAgentCommissionsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAgentCommissionsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetCommissionHistoryResponse> getCommissionHistory(
        com.game_engine.commission.v1.GetCommissionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCommissionHistoryMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_COMMISSION = 0;
  private static final int METHODID_GET_COMMISSION_BY_ID = 1;
  private static final int METHODID_GET_COMMISSIONS_BY_AFFILIATE = 2;
  private static final int METHODID_GET_COMMISSIONS_BY_MERCHANT = 3;
  private static final int METHODID_GET_COMMISSIONS_BY_PERIOD = 4;
  private static final int METHODID_GET_COMMISSION_BY_AFFILIATE_AND_PERIOD = 5;
  private static final int METHODID_GET_TOTAL_PAID_COMMISSION = 6;
  private static final int METHODID_GET_TOTAL_PENDING_COMMISSION = 7;
  private static final int METHODID_GET_TOTAL_REVENUE_BY_MERCHANT = 8;
  private static final int METHODID_CALCULATE_REVENUE_SHARE = 9;
  private static final int METHODID_CALCULATE_CPA = 10;
  private static final int METHODID_APPROVE_COMMISSION = 11;
  private static final int METHODID_REJECT_COMMISSION = 12;
  private static final int METHODID_PAY_COMMISSION = 13;
  private static final int METHODID_GET_PENDING_COMMISSIONS = 14;
  private static final int METHODID_GET_ALL_COMMISSIONS = 15;
  private static final int METHODID_DELETE_COMMISSION = 16;
  private static final int METHODID_SUBMIT_CLAIM = 17;
  private static final int METHODID_GET_USER_CLAIMS = 18;
  private static final int METHODID_GET_CLAIMS_BY_STATUS = 19;
  private static final int METHODID_CLAIM_COMMISSION = 20;
  private static final int METHODID_GET_AGENT_COMMISSIONS = 21;
  private static final int METHODID_GET_COMMISSION_HISTORY = 22;

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
        case METHODID_CREATE_COMMISSION:
          serviceImpl.createCommission((com.game_engine.commission.v1.CreateCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_BY_ID:
          serviceImpl.getCommissionById((com.game_engine.commission.v1.GetCommissionByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByIdResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSIONS_BY_AFFILIATE:
          serviceImpl.getCommissionsByAffiliate((com.game_engine.commission.v1.GetCommissionsByAffiliateRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByAffiliateResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSIONS_BY_MERCHANT:
          serviceImpl.getCommissionsByMerchant((com.game_engine.commission.v1.GetCommissionsByMerchantRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByMerchantResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSIONS_BY_PERIOD:
          serviceImpl.getCommissionsByPeriod((com.game_engine.commission.v1.GetCommissionsByPeriodRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionsByPeriodResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_BY_AFFILIATE_AND_PERIOD:
          serviceImpl.getCommissionByAffiliateAndPeriod((com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_PAID_COMMISSION:
          serviceImpl.getTotalPaidCommission((com.game_engine.commission.v1.GetTotalPaidCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPaidCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_PENDING_COMMISSION:
          serviceImpl.getTotalPendingCommission((com.game_engine.commission.v1.GetTotalPendingCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalPendingCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_REVENUE_BY_MERCHANT:
          serviceImpl.getTotalRevenueByMerchant((com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse>) responseObserver);
          break;
        case METHODID_CALCULATE_REVENUE_SHARE:
          serviceImpl.calculateRevenueShare((com.game_engine.commission.v1.CalculateRevenueShareRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateRevenueShareResponse>) responseObserver);
          break;
        case METHODID_CALCULATE_CPA:
          serviceImpl.calculateCPA((com.game_engine.commission.v1.CalculateCPARequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CalculateCPAResponse>) responseObserver);
          break;
        case METHODID_APPROVE_COMMISSION:
          serviceImpl.approveCommission((com.game_engine.commission.v1.ApproveCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ApproveCommissionResponse>) responseObserver);
          break;
        case METHODID_REJECT_COMMISSION:
          serviceImpl.rejectCommission((com.game_engine.commission.v1.RejectCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.RejectCommissionResponse>) responseObserver);
          break;
        case METHODID_PAY_COMMISSION:
          serviceImpl.payCommission((com.game_engine.commission.v1.PayCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.PayCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_PENDING_COMMISSIONS:
          serviceImpl.getPendingCommissions((com.game_engine.commission.v1.GetPendingCommissionsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetPendingCommissionsResponse>) responseObserver);
          break;
        case METHODID_GET_ALL_COMMISSIONS:
          serviceImpl.getAllCommissions((com.game_engine.commission.v1.GetAllCommissionsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllCommissionsResponse>) responseObserver);
          break;
        case METHODID_DELETE_COMMISSION:
          serviceImpl.deleteCommission((com.game_engine.commission.v1.DeleteCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteCommissionResponse>) responseObserver);
          break;
        case METHODID_SUBMIT_CLAIM:
          serviceImpl.submitClaim((com.game_engine.commission.v1.SubmitClaimRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.SubmitClaimResponse>) responseObserver);
          break;
        case METHODID_GET_USER_CLAIMS:
          serviceImpl.getUserClaims((com.game_engine.commission.v1.GetUserClaimsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetUserClaimsResponse>) responseObserver);
          break;
        case METHODID_GET_CLAIMS_BY_STATUS:
          serviceImpl.getClaimsByStatus((com.game_engine.commission.v1.GetClaimsByStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetClaimsByStatusResponse>) responseObserver);
          break;
        case METHODID_CLAIM_COMMISSION:
          serviceImpl.claimCommission((com.game_engine.commission.v1.ClaimCommissionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ClaimCommissionResponse>) responseObserver);
          break;
        case METHODID_GET_AGENT_COMMISSIONS:
          serviceImpl.getAgentCommissions((com.game_engine.commission.v1.GetAgentCommissionsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAgentCommissionsResponse>) responseObserver);
          break;
        case METHODID_GET_COMMISSION_HISTORY:
          serviceImpl.getCommissionHistory((com.game_engine.commission.v1.GetCommissionHistoryRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetCommissionHistoryResponse>) responseObserver);
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
          getCreateCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.CreateCommissionRequest,
              com.game_engine.commission.v1.CreateCommissionResponse>(
                service, METHODID_CREATE_COMMISSION)))
        .addMethod(
          getGetCommissionByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionByIdRequest,
              com.game_engine.commission.v1.GetCommissionByIdResponse>(
                service, METHODID_GET_COMMISSION_BY_ID)))
        .addMethod(
          getGetCommissionsByAffiliateMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionsByAffiliateRequest,
              com.game_engine.commission.v1.GetCommissionsByAffiliateResponse>(
                service, METHODID_GET_COMMISSIONS_BY_AFFILIATE)))
        .addMethod(
          getGetCommissionsByMerchantMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionsByMerchantRequest,
              com.game_engine.commission.v1.GetCommissionsByMerchantResponse>(
                service, METHODID_GET_COMMISSIONS_BY_MERCHANT)))
        .addMethod(
          getGetCommissionsByPeriodMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionsByPeriodRequest,
              com.game_engine.commission.v1.GetCommissionsByPeriodResponse>(
                service, METHODID_GET_COMMISSIONS_BY_PERIOD)))
        .addMethod(
          getGetCommissionByAffiliateAndPeriodMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodRequest,
              com.game_engine.commission.v1.GetCommissionByAffiliateAndPeriodResponse>(
                service, METHODID_GET_COMMISSION_BY_AFFILIATE_AND_PERIOD)))
        .addMethod(
          getGetTotalPaidCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetTotalPaidCommissionRequest,
              com.game_engine.commission.v1.GetTotalPaidCommissionResponse>(
                service, METHODID_GET_TOTAL_PAID_COMMISSION)))
        .addMethod(
          getGetTotalPendingCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetTotalPendingCommissionRequest,
              com.game_engine.commission.v1.GetTotalPendingCommissionResponse>(
                service, METHODID_GET_TOTAL_PENDING_COMMISSION)))
        .addMethod(
          getGetTotalRevenueByMerchantMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetTotalRevenueByMerchantRequest,
              com.game_engine.commission.v1.GetTotalRevenueByMerchantResponse>(
                service, METHODID_GET_TOTAL_REVENUE_BY_MERCHANT)))
        .addMethod(
          getCalculateRevenueShareMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.CalculateRevenueShareRequest,
              com.game_engine.commission.v1.CalculateRevenueShareResponse>(
                service, METHODID_CALCULATE_REVENUE_SHARE)))
        .addMethod(
          getCalculateCPAMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.CalculateCPARequest,
              com.game_engine.commission.v1.CalculateCPAResponse>(
                service, METHODID_CALCULATE_CPA)))
        .addMethod(
          getApproveCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.ApproveCommissionRequest,
              com.game_engine.commission.v1.ApproveCommissionResponse>(
                service, METHODID_APPROVE_COMMISSION)))
        .addMethod(
          getRejectCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.RejectCommissionRequest,
              com.game_engine.commission.v1.RejectCommissionResponse>(
                service, METHODID_REJECT_COMMISSION)))
        .addMethod(
          getPayCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.PayCommissionRequest,
              com.game_engine.commission.v1.PayCommissionResponse>(
                service, METHODID_PAY_COMMISSION)))
        .addMethod(
          getGetPendingCommissionsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetPendingCommissionsRequest,
              com.game_engine.commission.v1.GetPendingCommissionsResponse>(
                service, METHODID_GET_PENDING_COMMISSIONS)))
        .addMethod(
          getGetAllCommissionsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetAllCommissionsRequest,
              com.game_engine.commission.v1.GetAllCommissionsResponse>(
                service, METHODID_GET_ALL_COMMISSIONS)))
        .addMethod(
          getDeleteCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.DeleteCommissionRequest,
              com.game_engine.commission.v1.DeleteCommissionResponse>(
                service, METHODID_DELETE_COMMISSION)))
        .addMethod(
          getSubmitClaimMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.SubmitClaimRequest,
              com.game_engine.commission.v1.SubmitClaimResponse>(
                service, METHODID_SUBMIT_CLAIM)))
        .addMethod(
          getGetUserClaimsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetUserClaimsRequest,
              com.game_engine.commission.v1.GetUserClaimsResponse>(
                service, METHODID_GET_USER_CLAIMS)))
        .addMethod(
          getGetClaimsByStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetClaimsByStatusRequest,
              com.game_engine.commission.v1.GetClaimsByStatusResponse>(
                service, METHODID_GET_CLAIMS_BY_STATUS)))
        .addMethod(
          getClaimCommissionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.ClaimCommissionRequest,
              com.game_engine.commission.v1.ClaimCommissionResponse>(
                service, METHODID_CLAIM_COMMISSION)))
        .addMethod(
          getGetAgentCommissionsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetAgentCommissionsRequest,
              com.game_engine.commission.v1.GetAgentCommissionsResponse>(
                service, METHODID_GET_AGENT_COMMISSIONS)))
        .addMethod(
          getGetCommissionHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.commission.v1.GetCommissionHistoryRequest,
              com.game_engine.commission.v1.GetCommissionHistoryResponse>(
                service, METHODID_GET_COMMISSION_HISTORY)))
        .build();
  }

  private static abstract class CommissionServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CommissionServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.commission.v1.CommissionServiceOuterClass.getDescriptor();
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
              .addMethod(getCreateCommissionMethod())
              .addMethod(getGetCommissionByIdMethod())
              .addMethod(getGetCommissionsByAffiliateMethod())
              .addMethod(getGetCommissionsByMerchantMethod())
              .addMethod(getGetCommissionsByPeriodMethod())
              .addMethod(getGetCommissionByAffiliateAndPeriodMethod())
              .addMethod(getGetTotalPaidCommissionMethod())
              .addMethod(getGetTotalPendingCommissionMethod())
              .addMethod(getGetTotalRevenueByMerchantMethod())
              .addMethod(getCalculateRevenueShareMethod())
              .addMethod(getCalculateCPAMethod())
              .addMethod(getApproveCommissionMethod())
              .addMethod(getRejectCommissionMethod())
              .addMethod(getPayCommissionMethod())
              .addMethod(getGetPendingCommissionsMethod())
              .addMethod(getGetAllCommissionsMethod())
              .addMethod(getDeleteCommissionMethod())
              .addMethod(getSubmitClaimMethod())
              .addMethod(getGetUserClaimsMethod())
              .addMethod(getGetClaimsByStatusMethod())
              .addMethod(getClaimCommissionMethod())
              .addMethod(getGetAgentCommissionsMethod())
              .addMethod(getGetCommissionHistoryMethod())
              .build();
        }
      }
    }
    return result;
  }
}
