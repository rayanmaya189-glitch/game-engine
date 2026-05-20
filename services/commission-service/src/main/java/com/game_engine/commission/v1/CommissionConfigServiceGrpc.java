package com.game_engine.commission.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * =============================================================================
 * Commission Config Service - manages commission configurations
 * =============================================================================
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class CommissionConfigServiceGrpc {

  private CommissionConfigServiceGrpc() {
  }

  public static final java.lang.String SERVICE_NAME = "game_engine.commission.v1.CommissionConfigService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionConfigRequest, com.game_engine.commission.v1.CreateCommissionConfigResponse> getCreateCommissionConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "CreateCommissionConfig", requestType = com.game_engine.commission.v1.CreateCommissionConfigRequest.class, responseType = com.game_engine.commission.v1.CreateCommissionConfigResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionConfigRequest, com.game_engine.commission.v1.CreateCommissionConfigResponse> getCreateCommissionConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.CreateCommissionConfigRequest, com.game_engine.commission.v1.CreateCommissionConfigResponse> getCreateCommissionConfigMethod;
    if ((getCreateCommissionConfigMethod = CommissionConfigServiceGrpc.getCreateCommissionConfigMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getCreateCommissionConfigMethod = CommissionConfigServiceGrpc.getCreateCommissionConfigMethod) == null) {
          CommissionConfigServiceGrpc.getCreateCommissionConfigMethod = getCreateCommissionConfigMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.CreateCommissionConfigRequest, com.game_engine.commission.v1.CreateCommissionConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateCommissionConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateCommissionConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.CreateCommissionConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("CreateCommissionConfig"))
              .build();
        }
      }
    }
    return getCreateCommissionConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByIdRequest, com.game_engine.commission.v1.GetConfigByIdResponse> getGetConfigByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetConfigById", requestType = com.game_engine.commission.v1.GetConfigByIdRequest.class, responseType = com.game_engine.commission.v1.GetConfigByIdResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByIdRequest, com.game_engine.commission.v1.GetConfigByIdResponse> getGetConfigByIdMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByIdRequest, com.game_engine.commission.v1.GetConfigByIdResponse> getGetConfigByIdMethod;
    if ((getGetConfigByIdMethod = CommissionConfigServiceGrpc.getGetConfigByIdMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetConfigByIdMethod = CommissionConfigServiceGrpc.getGetConfigByIdMethod) == null) {
          CommissionConfigServiceGrpc.getGetConfigByIdMethod = getGetConfigByIdMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetConfigByIdRequest, com.game_engine.commission.v1.GetConfigByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConfigById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("GetConfigById"))
              .build();
        }
      }
    }
    return getGetConfigByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByAffiliateRequest, com.game_engine.commission.v1.GetConfigsByAffiliateResponse> getGetConfigsByAffiliateMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetConfigsByAffiliate", requestType = com.game_engine.commission.v1.GetConfigsByAffiliateRequest.class, responseType = com.game_engine.commission.v1.GetConfigsByAffiliateResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByAffiliateRequest, com.game_engine.commission.v1.GetConfigsByAffiliateResponse> getGetConfigsByAffiliateMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByAffiliateRequest, com.game_engine.commission.v1.GetConfigsByAffiliateResponse> getGetConfigsByAffiliateMethod;
    if ((getGetConfigsByAffiliateMethod = CommissionConfigServiceGrpc.getGetConfigsByAffiliateMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetConfigsByAffiliateMethod = CommissionConfigServiceGrpc.getGetConfigsByAffiliateMethod) == null) {
          CommissionConfigServiceGrpc.getGetConfigsByAffiliateMethod = getGetConfigsByAffiliateMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetConfigsByAffiliateRequest, com.game_engine.commission.v1.GetConfigsByAffiliateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConfigsByAffiliate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigsByAffiliateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigsByAffiliateResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("GetConfigsByAffiliate"))
              .build();
        }
      }
    }
    return getGetConfigsByAffiliateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByMerchantRequest, com.game_engine.commission.v1.GetConfigsByMerchantResponse> getGetConfigsByMerchantMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetConfigsByMerchant", requestType = com.game_engine.commission.v1.GetConfigsByMerchantRequest.class, responseType = com.game_engine.commission.v1.GetConfigsByMerchantResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByMerchantRequest, com.game_engine.commission.v1.GetConfigsByMerchantResponse> getGetConfigsByMerchantMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigsByMerchantRequest, com.game_engine.commission.v1.GetConfigsByMerchantResponse> getGetConfigsByMerchantMethod;
    if ((getGetConfigsByMerchantMethod = CommissionConfigServiceGrpc.getGetConfigsByMerchantMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetConfigsByMerchantMethod = CommissionConfigServiceGrpc.getGetConfigsByMerchantMethod) == null) {
          CommissionConfigServiceGrpc.getGetConfigsByMerchantMethod = getGetConfigsByMerchantMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetConfigsByMerchantRequest, com.game_engine.commission.v1.GetConfigsByMerchantResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConfigsByMerchant"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigsByMerchantRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigsByMerchantResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("GetConfigsByMerchant"))
              .build();
        }
      }
    }
    return getGetConfigsByMerchantMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> getGetActiveConfigsByAffiliateMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetActiveConfigsByAffiliate", requestType = com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest.class, responseType = com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> getGetActiveConfigsByAffiliateMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> getGetActiveConfigsByAffiliateMethod;
    if ((getGetActiveConfigsByAffiliateMethod = CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetActiveConfigsByAffiliateMethod = CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateMethod) == null) {
          CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateMethod = getGetActiveConfigsByAffiliateMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveConfigsByAffiliate"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("GetActiveConfigsByAffiliate"))
              .build();
        }
      }
    }
    return getGetActiveConfigsByAffiliateMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> getGetActiveConfigsByAffiliateAndMerchantMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetActiveConfigsByAffiliateAndMerchant", requestType = com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest.class, responseType = com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> getGetActiveConfigsByAffiliateAndMerchantMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> getGetActiveConfigsByAffiliateAndMerchantMethod;
    if ((getGetActiveConfigsByAffiliateAndMerchantMethod = CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateAndMerchantMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetActiveConfigsByAffiliateAndMerchantMethod = CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateAndMerchantMethod) == null) {
          CommissionConfigServiceGrpc.getGetActiveConfigsByAffiliateAndMerchantMethod = getGetActiveConfigsByAffiliateAndMerchantMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetActiveConfigsByAffiliateAndMerchant"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse.getDefaultInstance()))
              .setSchemaDescriptor(
                  new CommissionConfigServiceMethodDescriptorSupplier("GetActiveConfigsByAffiliateAndMerchant"))
              .build();
        }
      }
    }
    return getGetActiveConfigsByAffiliateAndMerchantMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest, com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> getGetConfigByAffiliateAndMerchantAndTypeMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetConfigByAffiliateAndMerchantAndType", requestType = com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest.class, responseType = com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest, com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> getGetConfigByAffiliateAndMerchantAndTypeMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest, com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> getGetConfigByAffiliateAndMerchantAndTypeMethod;
    if ((getGetConfigByAffiliateAndMerchantAndTypeMethod = CommissionConfigServiceGrpc.getGetConfigByAffiliateAndMerchantAndTypeMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetConfigByAffiliateAndMerchantAndTypeMethod = CommissionConfigServiceGrpc.getGetConfigByAffiliateAndMerchantAndTypeMethod) == null) {
          CommissionConfigServiceGrpc.getGetConfigByAffiliateAndMerchantAndTypeMethod = getGetConfigByAffiliateAndMerchantAndTypeMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest, com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConfigByAffiliateAndMerchantAndType"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse.getDefaultInstance()))
              .setSchemaDescriptor(
                  new CommissionConfigServiceMethodDescriptorSupplier("GetConfigByAffiliateAndMerchantAndType"))
              .build();
        }
      }
    }
    return getGetConfigByAffiliateAndMerchantAndTypeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateCommissionConfigRequest, com.game_engine.commission.v1.UpdateCommissionConfigResponse> getUpdateCommissionConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "UpdateCommissionConfig", requestType = com.game_engine.commission.v1.UpdateCommissionConfigRequest.class, responseType = com.game_engine.commission.v1.UpdateCommissionConfigResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateCommissionConfigRequest, com.game_engine.commission.v1.UpdateCommissionConfigResponse> getUpdateCommissionConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.UpdateCommissionConfigRequest, com.game_engine.commission.v1.UpdateCommissionConfigResponse> getUpdateCommissionConfigMethod;
    if ((getUpdateCommissionConfigMethod = CommissionConfigServiceGrpc.getUpdateCommissionConfigMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getUpdateCommissionConfigMethod = CommissionConfigServiceGrpc.getUpdateCommissionConfigMethod) == null) {
          CommissionConfigServiceGrpc.getUpdateCommissionConfigMethod = getUpdateCommissionConfigMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.UpdateCommissionConfigRequest, com.game_engine.commission.v1.UpdateCommissionConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateCommissionConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.UpdateCommissionConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.UpdateCommissionConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("UpdateCommissionConfig"))
              .build();
        }
      }
    }
    return getUpdateCommissionConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.ActivateConfigRequest, com.game_engine.commission.v1.ActivateConfigResponse> getActivateConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "ActivateConfig", requestType = com.game_engine.commission.v1.ActivateConfigRequest.class, responseType = com.game_engine.commission.v1.ActivateConfigResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.ActivateConfigRequest, com.game_engine.commission.v1.ActivateConfigResponse> getActivateConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.ActivateConfigRequest, com.game_engine.commission.v1.ActivateConfigResponse> getActivateConfigMethod;
    if ((getActivateConfigMethod = CommissionConfigServiceGrpc.getActivateConfigMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getActivateConfigMethod = CommissionConfigServiceGrpc.getActivateConfigMethod) == null) {
          CommissionConfigServiceGrpc.getActivateConfigMethod = getActivateConfigMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.ActivateConfigRequest, com.game_engine.commission.v1.ActivateConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ActivateConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ActivateConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.ActivateConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("ActivateConfig"))
              .build();
        }
      }
    }
    return getActivateConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeactivateConfigRequest, com.game_engine.commission.v1.DeactivateConfigResponse> getDeactivateConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "DeactivateConfig", requestType = com.game_engine.commission.v1.DeactivateConfigRequest.class, responseType = com.game_engine.commission.v1.DeactivateConfigResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeactivateConfigRequest, com.game_engine.commission.v1.DeactivateConfigResponse> getDeactivateConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeactivateConfigRequest, com.game_engine.commission.v1.DeactivateConfigResponse> getDeactivateConfigMethod;
    if ((getDeactivateConfigMethod = CommissionConfigServiceGrpc.getDeactivateConfigMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getDeactivateConfigMethod = CommissionConfigServiceGrpc.getDeactivateConfigMethod) == null) {
          CommissionConfigServiceGrpc.getDeactivateConfigMethod = getDeactivateConfigMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.DeactivateConfigRequest, com.game_engine.commission.v1.DeactivateConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeactivateConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeactivateConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeactivateConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("DeactivateConfig"))
              .build();
        }
      }
    }
    return getDeactivateConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteConfigRequest, com.game_engine.commission.v1.DeleteConfigResponse> getDeleteConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "DeleteConfig", requestType = com.game_engine.commission.v1.DeleteConfigRequest.class, responseType = com.game_engine.commission.v1.DeleteConfigResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteConfigRequest, com.game_engine.commission.v1.DeleteConfigResponse> getDeleteConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.DeleteConfigRequest, com.game_engine.commission.v1.DeleteConfigResponse> getDeleteConfigMethod;
    if ((getDeleteConfigMethod = CommissionConfigServiceGrpc.getDeleteConfigMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getDeleteConfigMethod = CommissionConfigServiceGrpc.getDeleteConfigMethod) == null) {
          CommissionConfigServiceGrpc.getDeleteConfigMethod = getDeleteConfigMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.DeleteConfigRequest, com.game_engine.commission.v1.DeleteConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeleteConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.DeleteConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("DeleteConfig"))
              .build();
        }
      }
    }
    return getDeleteConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllConfigsRequest, com.game_engine.commission.v1.GetAllConfigsResponse> getGetAllConfigsMethod;

  @io.grpc.stub.annotations.RpcMethod(fullMethodName = SERVICE_NAME + '/'
      + "GetAllConfigs", requestType = com.game_engine.commission.v1.GetAllConfigsRequest.class, responseType = com.game_engine.commission.v1.GetAllConfigsResponse.class, methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllConfigsRequest, com.game_engine.commission.v1.GetAllConfigsResponse> getGetAllConfigsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.commission.v1.GetAllConfigsRequest, com.game_engine.commission.v1.GetAllConfigsResponse> getGetAllConfigsMethod;
    if ((getGetAllConfigsMethod = CommissionConfigServiceGrpc.getGetAllConfigsMethod) == null) {
      synchronized (CommissionConfigServiceGrpc.class) {
        if ((getGetAllConfigsMethod = CommissionConfigServiceGrpc.getGetAllConfigsMethod) == null) {
          CommissionConfigServiceGrpc.getGetAllConfigsMethod = getGetAllConfigsMethod = io.grpc.MethodDescriptor.<com.game_engine.commission.v1.GetAllConfigsRequest, com.game_engine.commission.v1.GetAllConfigsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllConfigs"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAllConfigsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.commission.v1.GetAllConfigsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CommissionConfigServiceMethodDescriptorSupplier("GetAllConfigs"))
              .build();
        }
      }
    }
    return getGetAllConfigsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CommissionConfigServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceStub> factory = new io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceStub>() {
      @java.lang.Override
      public CommissionConfigServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
        return new CommissionConfigServiceStub(channel, callOptions);
      }
    };
    return CommissionConfigServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the
   * service
   */
  public static CommissionConfigServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceBlockingV2Stub> factory = new io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceBlockingV2Stub>() {
      @java.lang.Override
      public CommissionConfigServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
        return new CommissionConfigServiceBlockingV2Stub(channel, callOptions);
      }
    };
    return CommissionConfigServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output
   * calls on the service
   */
  public static CommissionConfigServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceBlockingStub> factory = new io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceBlockingStub>() {
      @java.lang.Override
      public CommissionConfigServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
        return new CommissionConfigServiceBlockingStub(channel, callOptions);
      }
    };
    return CommissionConfigServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the
   * service
   */
  public static CommissionConfigServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceFutureStub> factory = new io.grpc.stub.AbstractStub.StubFactory<CommissionConfigServiceFutureStub>() {
      @java.lang.Override
      public CommissionConfigServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
        return new CommissionConfigServiceFutureStub(channel, callOptions);
      }
    };
    return CommissionConfigServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void createCommissionConfig(com.game_engine.commission.v1.CreateCommissionConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateCommissionConfigMethod(), responseObserver);
    }

    /**
     */
    default void getConfigById(com.game_engine.commission.v1.GetConfigByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConfigByIdMethod(), responseObserver);
    }

    /**
     */
    default void getConfigsByAffiliate(com.game_engine.commission.v1.GetConfigsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConfigsByAffiliateMethod(), responseObserver);
    }

    /**
     */
    default void getConfigsByMerchant(com.game_engine.commission.v1.GetConfigsByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByMerchantResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConfigsByMerchantMethod(), responseObserver);
    }

    /**
     */
    default void getActiveConfigsByAffiliate(com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveConfigsByAffiliateMethod(), responseObserver);
    }

    /**
     */
    default void getActiveConfigsByAffiliateAndMerchant(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetActiveConfigsByAffiliateAndMerchantMethod(),
          responseObserver);
    }

    /**
     */
    default void getConfigByAffiliateAndMerchantAndType(
        com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConfigByAffiliateAndMerchantAndTypeMethod(),
          responseObserver);
    }

    /**
     */
    default void updateCommissionConfig(com.game_engine.commission.v1.UpdateCommissionConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateCommissionConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateCommissionConfigMethod(), responseObserver);
    }

    /**
     */
    default void activateConfig(com.game_engine.commission.v1.ActivateConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ActivateConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getActivateConfigMethod(), responseObserver);
    }

    /**
     */
    default void deactivateConfig(com.game_engine.commission.v1.DeactivateConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeactivateConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeactivateConfigMethod(), responseObserver);
    }

    /**
     */
    default void deleteConfig(com.game_engine.commission.v1.DeleteConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteConfigMethod(), responseObserver);
    }

    /**
     */
    default void getAllConfigs(com.game_engine.commission.v1.GetAllConfigsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllConfigsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllConfigsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service
   * CommissionConfigService.
   * 
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public static abstract class CommissionConfigServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override
    public final io.grpc.ServerServiceDefinition bindService() {
      return CommissionConfigServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service
   * CommissionConfigService.
   * 
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public static final class CommissionConfigServiceStub
      extends io.grpc.stub.AbstractAsyncStub<CommissionConfigServiceStub> {
    private CommissionConfigServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionConfigServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionConfigServiceStub(channel, callOptions);
    }

    /**
     */
    public void createCommissionConfig(com.game_engine.commission.v1.CreateCommissionConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateCommissionConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getConfigById(com.game_engine.commission.v1.GetConfigByIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConfigByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getConfigsByAffiliate(com.game_engine.commission.v1.GetConfigsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConfigsByAffiliateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getConfigsByMerchant(com.game_engine.commission.v1.GetConfigsByMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByMerchantResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConfigsByMerchantMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getActiveConfigsByAffiliate(com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveConfigsByAffiliateMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getActiveConfigsByAffiliateAndMerchant(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetActiveConfigsByAffiliateAndMerchantMethod(), getCallOptions()), request,
          responseObserver);
    }

    /**
     */
    public void getConfigByAffiliateAndMerchantAndType(
        com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConfigByAffiliateAndMerchantAndTypeMethod(), getCallOptions()), request,
          responseObserver);
    }

    /**
     */
    public void updateCommissionConfig(com.game_engine.commission.v1.UpdateCommissionConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateCommissionConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateCommissionConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void activateConfig(com.game_engine.commission.v1.ActivateConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ActivateConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getActivateConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deactivateConfig(com.game_engine.commission.v1.DeactivateConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeactivateConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeactivateConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteConfig(com.game_engine.commission.v1.DeleteConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllConfigs(com.game_engine.commission.v1.GetAllConfigsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllConfigsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllConfigsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service
   * CommissionConfigService.
   * 
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public static final class CommissionConfigServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<CommissionConfigServiceBlockingV2Stub> {
    private CommissionConfigServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionConfigServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionConfigServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.commission.v1.CreateCommissionConfigResponse createCommissionConfig(
        com.game_engine.commission.v1.CreateCommissionConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCommissionConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigByIdResponse getConfigById(
        com.game_engine.commission.v1.GetConfigByIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigsByAffiliateResponse getConfigsByAffiliate(
        com.game_engine.commission.v1.GetConfigsByAffiliateRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigsByMerchantResponse getConfigsByMerchant(
        com.game_engine.commission.v1.GetConfigsByMerchantRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigsByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse getActiveConfigsByAffiliate(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveConfigsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse getActiveConfigsByAffiliateAndMerchant(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest request)
        throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveConfigsByAffiliateAndMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse getConfigByAffiliateAndMerchantAndType(
        com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest request)
        throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigByAffiliateAndMerchantAndTypeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.UpdateCommissionConfigResponse updateCommissionConfig(
        com.game_engine.commission.v1.UpdateCommissionConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateCommissionConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ActivateConfigResponse activateConfig(
        com.game_engine.commission.v1.ActivateConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getActivateConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeactivateConfigResponse deactivateConfig(
        com.game_engine.commission.v1.DeactivateConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeactivateConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeleteConfigResponse deleteConfig(
        com.game_engine.commission.v1.DeleteConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAllConfigsResponse getAllConfigs(
        com.game_engine.commission.v1.GetAllConfigsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllConfigsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service
   * CommissionConfigService.
   * 
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public static final class CommissionConfigServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<CommissionConfigServiceBlockingStub> {
    private CommissionConfigServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionConfigServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionConfigServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.commission.v1.CreateCommissionConfigResponse createCommissionConfig(
        com.game_engine.commission.v1.CreateCommissionConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCommissionConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigByIdResponse getConfigById(
        com.game_engine.commission.v1.GetConfigByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigsByAffiliateResponse getConfigsByAffiliate(
        com.game_engine.commission.v1.GetConfigsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigsByMerchantResponse getConfigsByMerchant(
        com.game_engine.commission.v1.GetConfigsByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigsByMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse getActiveConfigsByAffiliate(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveConfigsByAffiliateMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse getActiveConfigsByAffiliateAndMerchant(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetActiveConfigsByAffiliateAndMerchantMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse getConfigByAffiliateAndMerchantAndType(
        com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigByAffiliateAndMerchantAndTypeMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.UpdateCommissionConfigResponse updateCommissionConfig(
        com.game_engine.commission.v1.UpdateCommissionConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateCommissionConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.ActivateConfigResponse activateConfig(
        com.game_engine.commission.v1.ActivateConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getActivateConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeactivateConfigResponse deactivateConfig(
        com.game_engine.commission.v1.DeactivateConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeactivateConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.DeleteConfigResponse deleteConfig(
        com.game_engine.commission.v1.DeleteConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.commission.v1.GetAllConfigsResponse getAllConfigs(
        com.game_engine.commission.v1.GetAllConfigsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllConfigsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service
   * CommissionConfigService.
   * 
   * <pre>
   * =============================================================================
   * Commission Config Service - manages commission configurations
   * =============================================================================
   * </pre>
   */
  public static final class CommissionConfigServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<CommissionConfigServiceFutureStub> {
    private CommissionConfigServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CommissionConfigServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CommissionConfigServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.CreateCommissionConfigResponse> createCommissionConfig(
        com.game_engine.commission.v1.CreateCommissionConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateCommissionConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetConfigByIdResponse> getConfigById(
        com.game_engine.commission.v1.GetConfigByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConfigByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetConfigsByAffiliateResponse> getConfigsByAffiliate(
        com.game_engine.commission.v1.GetConfigsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConfigsByAffiliateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetConfigsByMerchantResponse> getConfigsByMerchant(
        com.game_engine.commission.v1.GetConfigsByMerchantRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConfigsByMerchantMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse> getActiveConfigsByAffiliate(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveConfigsByAffiliateMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse> getActiveConfigsByAffiliateAndMerchant(
        com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetActiveConfigsByAffiliateAndMerchantMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse> getConfigByAffiliateAndMerchantAndType(
        com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConfigByAffiliateAndMerchantAndTypeMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.UpdateCommissionConfigResponse> updateCommissionConfig(
        com.game_engine.commission.v1.UpdateCommissionConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateCommissionConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.ActivateConfigResponse> activateConfig(
        com.game_engine.commission.v1.ActivateConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getActivateConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.DeactivateConfigResponse> deactivateConfig(
        com.game_engine.commission.v1.DeactivateConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeactivateConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.DeleteConfigResponse> deleteConfig(
        com.game_engine.commission.v1.DeleteConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.commission.v1.GetAllConfigsResponse> getAllConfigs(
        com.game_engine.commission.v1.GetAllConfigsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllConfigsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_COMMISSION_CONFIG = 0;
  private static final int METHODID_GET_CONFIG_BY_ID = 1;
  private static final int METHODID_GET_CONFIGS_BY_AFFILIATE = 2;
  private static final int METHODID_GET_CONFIGS_BY_MERCHANT = 3;
  private static final int METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE = 4;
  private static final int METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE_AND_MERCHANT = 5;
  private static final int METHODID_GET_CONFIG_BY_AFFILIATE_AND_MERCHANT_AND_TYPE = 6;
  private static final int METHODID_UPDATE_COMMISSION_CONFIG = 7;
  private static final int METHODID_ACTIVATE_CONFIG = 8;
  private static final int METHODID_DEACTIVATE_CONFIG = 9;
  private static final int METHODID_DELETE_CONFIG = 10;
  private static final int METHODID_GET_ALL_CONFIGS = 11;

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
        case METHODID_CREATE_COMMISSION_CONFIG:
          serviceImpl.createCommissionConfig((com.game_engine.commission.v1.CreateCommissionConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.CreateCommissionConfigResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIG_BY_ID:
          serviceImpl.getConfigById((com.game_engine.commission.v1.GetConfigByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByIdResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIGS_BY_AFFILIATE:
          serviceImpl.getConfigsByAffiliate((com.game_engine.commission.v1.GetConfigsByAffiliateRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByAffiliateResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIGS_BY_MERCHANT:
          serviceImpl.getConfigsByMerchant((com.game_engine.commission.v1.GetConfigsByMerchantRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigsByMerchantResponse>) responseObserver);
          break;
        case METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE:
          serviceImpl.getActiveConfigsByAffiliate(
              (com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse>) responseObserver);
          break;
        case METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE_AND_MERCHANT:
          serviceImpl.getActiveConfigsByAffiliateAndMerchant(
              (com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIG_BY_AFFILIATE_AND_MERCHANT_AND_TYPE:
          serviceImpl.getConfigByAffiliateAndMerchantAndType(
              (com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse>) responseObserver);
          break;
        case METHODID_UPDATE_COMMISSION_CONFIG:
          serviceImpl.updateCommissionConfig((com.game_engine.commission.v1.UpdateCommissionConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.UpdateCommissionConfigResponse>) responseObserver);
          break;
        case METHODID_ACTIVATE_CONFIG:
          serviceImpl.activateConfig((com.game_engine.commission.v1.ActivateConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.ActivateConfigResponse>) responseObserver);
          break;
        case METHODID_DEACTIVATE_CONFIG:
          serviceImpl.deactivateConfig((com.game_engine.commission.v1.DeactivateConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeactivateConfigResponse>) responseObserver);
          break;
        case METHODID_DELETE_CONFIG:
          serviceImpl.deleteConfig((com.game_engine.commission.v1.DeleteConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.DeleteConfigResponse>) responseObserver);
          break;
        case METHODID_GET_ALL_CONFIGS:
          serviceImpl.getAllConfigs((com.game_engine.commission.v1.GetAllConfigsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.commission.v1.GetAllConfigsResponse>) responseObserver);
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
            getCreateCommissionConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.CreateCommissionConfigRequest, com.game_engine.commission.v1.CreateCommissionConfigResponse>(
                    service, METHODID_CREATE_COMMISSION_CONFIG)))
        .addMethod(
            getGetConfigByIdMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetConfigByIdRequest, com.game_engine.commission.v1.GetConfigByIdResponse>(
                    service, METHODID_GET_CONFIG_BY_ID)))
        .addMethod(
            getGetConfigsByAffiliateMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetConfigsByAffiliateRequest, com.game_engine.commission.v1.GetConfigsByAffiliateResponse>(
                    service, METHODID_GET_CONFIGS_BY_AFFILIATE)))
        .addMethod(
            getGetConfigsByMerchantMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetConfigsByMerchantRequest, com.game_engine.commission.v1.GetConfigsByMerchantResponse>(
                    service, METHODID_GET_CONFIGS_BY_MERCHANT)))
        .addMethod(
            getGetActiveConfigsByAffiliateMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetActiveConfigsByAffiliateRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateResponse>(
                    service, METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE)))
        .addMethod(
            getGetActiveConfigsByAffiliateAndMerchantMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantRequest, com.game_engine.commission.v1.GetActiveConfigsByAffiliateAndMerchantResponse>(
                    service, METHODID_GET_ACTIVE_CONFIGS_BY_AFFILIATE_AND_MERCHANT)))
        .addMethod(
            getGetConfigByAffiliateAndMerchantAndTypeMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeRequest, com.game_engine.commission.v1.GetConfigByAffiliateAndMerchantAndTypeResponse>(
                    service, METHODID_GET_CONFIG_BY_AFFILIATE_AND_MERCHANT_AND_TYPE)))
        .addMethod(
            getUpdateCommissionConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.UpdateCommissionConfigRequest, com.game_engine.commission.v1.UpdateCommissionConfigResponse>(
                    service, METHODID_UPDATE_COMMISSION_CONFIG)))
        .addMethod(
            getActivateConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.ActivateConfigRequest, com.game_engine.commission.v1.ActivateConfigResponse>(
                    service, METHODID_ACTIVATE_CONFIG)))
        .addMethod(
            getDeactivateConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.DeactivateConfigRequest, com.game_engine.commission.v1.DeactivateConfigResponse>(
                    service, METHODID_DEACTIVATE_CONFIG)))
        .addMethod(
            getDeleteConfigMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.DeleteConfigRequest, com.game_engine.commission.v1.DeleteConfigResponse>(
                    service, METHODID_DELETE_CONFIG)))
        .addMethod(
            getGetAllConfigsMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
                new MethodHandlers<com.game_engine.commission.v1.GetAllConfigsRequest, com.game_engine.commission.v1.GetAllConfigsResponse>(
                    service, METHODID_GET_ALL_CONFIGS)))
        .build();
  }

  private static abstract class CommissionConfigServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CommissionConfigServiceBaseDescriptorSupplier() {
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.commission.v1.CommissionServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CommissionConfigService");
    }
  }

  private static final class CommissionConfigServiceFileDescriptorSupplier
      extends CommissionConfigServiceBaseDescriptorSupplier {
    CommissionConfigServiceFileDescriptorSupplier() {
    }
  }

  private static final class CommissionConfigServiceMethodDescriptorSupplier
      extends CommissionConfigServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    CommissionConfigServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (CommissionConfigServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CommissionConfigServiceFileDescriptorSupplier())
              .addMethod(getCreateCommissionConfigMethod())
              .addMethod(getGetConfigByIdMethod())
              .addMethod(getGetConfigsByAffiliateMethod())
              .addMethod(getGetConfigsByMerchantMethod())
              .addMethod(getGetActiveConfigsByAffiliateMethod())
              .addMethod(getGetActiveConfigsByAffiliateAndMerchantMethod())
              .addMethod(getGetConfigByAffiliateAndMerchantAndTypeMethod())
              .addMethod(getUpdateCommissionConfigMethod())
              .addMethod(getActivateConfigMethod())
              .addMethod(getDeactivateConfigMethod())
              .addMethod(getDeleteConfigMethod())
              .addMethod(getGetAllConfigsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
