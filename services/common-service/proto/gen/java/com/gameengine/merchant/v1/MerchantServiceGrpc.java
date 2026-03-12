package com.gameengine.merchant.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class MerchantServiceGrpc {

  private MerchantServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "gameengine.merchant.v1.MerchantService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListPlayersRequest,
      com.gameengine.merchant.v1.ListPlayersResponse> getListPlayersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPlayers",
      requestType = com.gameengine.merchant.v1.ListPlayersRequest.class,
      responseType = com.gameengine.merchant.v1.ListPlayersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListPlayersRequest,
      com.gameengine.merchant.v1.ListPlayersResponse> getListPlayersMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListPlayersRequest, com.gameengine.merchant.v1.ListPlayersResponse> getListPlayersMethod;
    if ((getListPlayersMethod = MerchantServiceGrpc.getListPlayersMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getListPlayersMethod = MerchantServiceGrpc.getListPlayersMethod) == null) {
          MerchantServiceGrpc.getListPlayersMethod = getListPlayersMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.ListPlayersRequest, com.gameengine.merchant.v1.ListPlayersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPlayers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListPlayersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListPlayersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("ListPlayers"))
              .build();
        }
      }
    }
    return getListPlayersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerRequest,
      com.gameengine.merchant.v1.GetPlayerResponse> getGetPlayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayer",
      requestType = com.gameengine.merchant.v1.GetPlayerRequest.class,
      responseType = com.gameengine.merchant.v1.GetPlayerResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerRequest,
      com.gameengine.merchant.v1.GetPlayerResponse> getGetPlayerMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerRequest, com.gameengine.merchant.v1.GetPlayerResponse> getGetPlayerMethod;
    if ((getGetPlayerMethod = MerchantServiceGrpc.getGetPlayerMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetPlayerMethod = MerchantServiceGrpc.getGetPlayerMethod) == null) {
          MerchantServiceGrpc.getGetPlayerMethod = getGetPlayerMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetPlayerRequest, com.gameengine.merchant.v1.GetPlayerResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetPlayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetPlayerResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetPlayer"))
              .build();
        }
      }
    }
    return getGetPlayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetRevenueReportRequest,
      com.gameengine.merchant.v1.GetRevenueReportResponse> getGetRevenueReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetRevenueReport",
      requestType = com.gameengine.merchant.v1.GetRevenueReportRequest.class,
      responseType = com.gameengine.merchant.v1.GetRevenueReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetRevenueReportRequest,
      com.gameengine.merchant.v1.GetRevenueReportResponse> getGetRevenueReportMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetRevenueReportRequest, com.gameengine.merchant.v1.GetRevenueReportResponse> getGetRevenueReportMethod;
    if ((getGetRevenueReportMethod = MerchantServiceGrpc.getGetRevenueReportMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetRevenueReportMethod = MerchantServiceGrpc.getGetRevenueReportMethod) == null) {
          MerchantServiceGrpc.getGetRevenueReportMethod = getGetRevenueReportMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetRevenueReportRequest, com.gameengine.merchant.v1.GetRevenueReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetRevenueReport"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetRevenueReportRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetRevenueReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetRevenueReport"))
              .build();
        }
      }
    }
    return getGetRevenueReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerReportRequest,
      com.gameengine.merchant.v1.GetPlayerReportResponse> getGetPlayerReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayerReport",
      requestType = com.gameengine.merchant.v1.GetPlayerReportRequest.class,
      responseType = com.gameengine.merchant.v1.GetPlayerReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerReportRequest,
      com.gameengine.merchant.v1.GetPlayerReportResponse> getGetPlayerReportMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetPlayerReportRequest, com.gameengine.merchant.v1.GetPlayerReportResponse> getGetPlayerReportMethod;
    if ((getGetPlayerReportMethod = MerchantServiceGrpc.getGetPlayerReportMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetPlayerReportMethod = MerchantServiceGrpc.getGetPlayerReportMethod) == null) {
          MerchantServiceGrpc.getGetPlayerReportMethod = getGetPlayerReportMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetPlayerReportRequest, com.gameengine.merchant.v1.GetPlayerReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayerReport"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetPlayerReportRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetPlayerReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetPlayerReport"))
              .build();
        }
      }
    }
    return getGetPlayerReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetGameReportRequest,
      com.gameengine.merchant.v1.GetGameReportResponse> getGetGameReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetGameReport",
      requestType = com.gameengine.merchant.v1.GetGameReportRequest.class,
      responseType = com.gameengine.merchant.v1.GetGameReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetGameReportRequest,
      com.gameengine.merchant.v1.GetGameReportResponse> getGetGameReportMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetGameReportRequest, com.gameengine.merchant.v1.GetGameReportResponse> getGetGameReportMethod;
    if ((getGetGameReportMethod = MerchantServiceGrpc.getGetGameReportMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetGameReportMethod = MerchantServiceGrpc.getGetGameReportMethod) == null) {
          MerchantServiceGrpc.getGetGameReportMethod = getGetGameReportMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetGameReportRequest, com.gameengine.merchant.v1.GetGameReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetGameReport"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetGameReportRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetGameReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetGameReport"))
              .build();
        }
      }
    }
    return getGetGameReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetConfigRequest,
      com.gameengine.merchant.v1.GetConfigResponse> getGetConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetConfig",
      requestType = com.gameengine.merchant.v1.GetConfigRequest.class,
      responseType = com.gameengine.merchant.v1.GetConfigResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetConfigRequest,
      com.gameengine.merchant.v1.GetConfigResponse> getGetConfigMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetConfigRequest, com.gameengine.merchant.v1.GetConfigResponse> getGetConfigMethod;
    if ((getGetConfigMethod = MerchantServiceGrpc.getGetConfigMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetConfigMethod = MerchantServiceGrpc.getGetConfigMethod) == null) {
          MerchantServiceGrpc.getGetConfigMethod = getGetConfigMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetConfigRequest, com.gameengine.merchant.v1.GetConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetConfig"))
              .build();
        }
      }
    }
    return getGetConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateConfigRequest,
      com.gameengine.merchant.v1.UpdateConfigResponse> getUpdateConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateConfig",
      requestType = com.gameengine.merchant.v1.UpdateConfigRequest.class,
      responseType = com.gameengine.merchant.v1.UpdateConfigResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateConfigRequest,
      com.gameengine.merchant.v1.UpdateConfigResponse> getUpdateConfigMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateConfigRequest, com.gameengine.merchant.v1.UpdateConfigResponse> getUpdateConfigMethod;
    if ((getUpdateConfigMethod = MerchantServiceGrpc.getUpdateConfigMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getUpdateConfigMethod = MerchantServiceGrpc.getUpdateConfigMethod) == null) {
          MerchantServiceGrpc.getUpdateConfigMethod = getUpdateConfigMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.UpdateConfigRequest, com.gameengine.merchant.v1.UpdateConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("UpdateConfig"))
              .build();
        }
      }
    }
    return getUpdateConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.RegisterWebhookRequest,
      com.gameengine.merchant.v1.RegisterWebhookResponse> getRegisterWebhookMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegisterWebhook",
      requestType = com.gameengine.merchant.v1.RegisterWebhookRequest.class,
      responseType = com.gameengine.merchant.v1.RegisterWebhookResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.RegisterWebhookRequest,
      com.gameengine.merchant.v1.RegisterWebhookResponse> getRegisterWebhookMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.RegisterWebhookRequest, com.gameengine.merchant.v1.RegisterWebhookResponse> getRegisterWebhookMethod;
    if ((getRegisterWebhookMethod = MerchantServiceGrpc.getRegisterWebhookMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getRegisterWebhookMethod = MerchantServiceGrpc.getRegisterWebhookMethod) == null) {
          MerchantServiceGrpc.getRegisterWebhookMethod = getRegisterWebhookMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.RegisterWebhookRequest, com.gameengine.merchant.v1.RegisterWebhookResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegisterWebhook"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.RegisterWebhookRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.RegisterWebhookResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("RegisterWebhook"))
              .build();
        }
      }
    }
    return getRegisterWebhookMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListWebhooksRequest,
      com.gameengine.merchant.v1.ListWebhooksResponse> getListWebhooksMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListWebhooks",
      requestType = com.gameengine.merchant.v1.ListWebhooksRequest.class,
      responseType = com.gameengine.merchant.v1.ListWebhooksResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListWebhooksRequest,
      com.gameengine.merchant.v1.ListWebhooksResponse> getListWebhooksMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListWebhooksRequest, com.gameengine.merchant.v1.ListWebhooksResponse> getListWebhooksMethod;
    if ((getListWebhooksMethod = MerchantServiceGrpc.getListWebhooksMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getListWebhooksMethod = MerchantServiceGrpc.getListWebhooksMethod) == null) {
          MerchantServiceGrpc.getListWebhooksMethod = getListWebhooksMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.ListWebhooksRequest, com.gameengine.merchant.v1.ListWebhooksResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListWebhooks"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListWebhooksRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListWebhooksResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("ListWebhooks"))
              .build();
        }
      }
    }
    return getListWebhooksMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.DeleteWebhookRequest,
      com.gameengine.merchant.v1.DeleteWebhookResponse> getDeleteWebhookMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteWebhook",
      requestType = com.gameengine.merchant.v1.DeleteWebhookRequest.class,
      responseType = com.gameengine.merchant.v1.DeleteWebhookResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.DeleteWebhookRequest,
      com.gameengine.merchant.v1.DeleteWebhookResponse> getDeleteWebhookMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.DeleteWebhookRequest, com.gameengine.merchant.v1.DeleteWebhookResponse> getDeleteWebhookMethod;
    if ((getDeleteWebhookMethod = MerchantServiceGrpc.getDeleteWebhookMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getDeleteWebhookMethod = MerchantServiceGrpc.getDeleteWebhookMethod) == null) {
          MerchantServiceGrpc.getDeleteWebhookMethod = getDeleteWebhookMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.DeleteWebhookRequest, com.gameengine.merchant.v1.DeleteWebhookResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteWebhook"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.DeleteWebhookRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.DeleteWebhookResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("DeleteWebhook"))
              .build();
        }
      }
    }
    return getDeleteWebhookMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListAgentsRequest,
      com.gameengine.merchant.v1.ListAgentsResponse> getListAgentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListAgents",
      requestType = com.gameengine.merchant.v1.ListAgentsRequest.class,
      responseType = com.gameengine.merchant.v1.ListAgentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListAgentsRequest,
      com.gameengine.merchant.v1.ListAgentsResponse> getListAgentsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.ListAgentsRequest, com.gameengine.merchant.v1.ListAgentsResponse> getListAgentsMethod;
    if ((getListAgentsMethod = MerchantServiceGrpc.getListAgentsMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getListAgentsMethod = MerchantServiceGrpc.getListAgentsMethod) == null) {
          MerchantServiceGrpc.getListAgentsMethod = getListAgentsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.ListAgentsRequest, com.gameengine.merchant.v1.ListAgentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListAgents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListAgentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.ListAgentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("ListAgents"))
              .build();
        }
      }
    }
    return getListAgentsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetAgentRequest,
      com.gameengine.merchant.v1.GetAgentResponse> getGetAgentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAgent",
      requestType = com.gameengine.merchant.v1.GetAgentRequest.class,
      responseType = com.gameengine.merchant.v1.GetAgentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetAgentRequest,
      com.gameengine.merchant.v1.GetAgentResponse> getGetAgentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.GetAgentRequest, com.gameengine.merchant.v1.GetAgentResponse> getGetAgentMethod;
    if ((getGetAgentMethod = MerchantServiceGrpc.getGetAgentMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getGetAgentMethod = MerchantServiceGrpc.getGetAgentMethod) == null) {
          MerchantServiceGrpc.getGetAgentMethod = getGetAgentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.GetAgentRequest, com.gameengine.merchant.v1.GetAgentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAgent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetAgentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.GetAgentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("GetAgent"))
              .build();
        }
      }
    }
    return getGetAgentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.CreateAgentRequest,
      com.gameengine.merchant.v1.CreateAgentResponse> getCreateAgentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAgent",
      requestType = com.gameengine.merchant.v1.CreateAgentRequest.class,
      responseType = com.gameengine.merchant.v1.CreateAgentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.CreateAgentRequest,
      com.gameengine.merchant.v1.CreateAgentResponse> getCreateAgentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.CreateAgentRequest, com.gameengine.merchant.v1.CreateAgentResponse> getCreateAgentMethod;
    if ((getCreateAgentMethod = MerchantServiceGrpc.getCreateAgentMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getCreateAgentMethod = MerchantServiceGrpc.getCreateAgentMethod) == null) {
          MerchantServiceGrpc.getCreateAgentMethod = getCreateAgentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.CreateAgentRequest, com.gameengine.merchant.v1.CreateAgentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAgent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.CreateAgentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.CreateAgentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("CreateAgent"))
              .build();
        }
      }
    }
    return getCreateAgentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentRequest,
      com.gameengine.merchant.v1.UpdateAgentResponse> getUpdateAgentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateAgent",
      requestType = com.gameengine.merchant.v1.UpdateAgentRequest.class,
      responseType = com.gameengine.merchant.v1.UpdateAgentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentRequest,
      com.gameengine.merchant.v1.UpdateAgentResponse> getUpdateAgentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentRequest, com.gameengine.merchant.v1.UpdateAgentResponse> getUpdateAgentMethod;
    if ((getUpdateAgentMethod = MerchantServiceGrpc.getUpdateAgentMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getUpdateAgentMethod = MerchantServiceGrpc.getUpdateAgentMethod) == null) {
          MerchantServiceGrpc.getUpdateAgentMethod = getUpdateAgentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.UpdateAgentRequest, com.gameengine.merchant.v1.UpdateAgentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateAgent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateAgentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateAgentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("UpdateAgent"))
              .build();
        }
      }
    }
    return getUpdateAgentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentStatusRequest,
      com.gameengine.merchant.v1.UpdateAgentStatusResponse> getUpdateAgentStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateAgentStatus",
      requestType = com.gameengine.merchant.v1.UpdateAgentStatusRequest.class,
      responseType = com.gameengine.merchant.v1.UpdateAgentStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentStatusRequest,
      com.gameengine.merchant.v1.UpdateAgentStatusResponse> getUpdateAgentStatusMethod() {
    io.grpc.MethodDescriptor<com.gameengine.merchant.v1.UpdateAgentStatusRequest, com.gameengine.merchant.v1.UpdateAgentStatusResponse> getUpdateAgentStatusMethod;
    if ((getUpdateAgentStatusMethod = MerchantServiceGrpc.getUpdateAgentStatusMethod) == null) {
      synchronized (MerchantServiceGrpc.class) {
        if ((getUpdateAgentStatusMethod = MerchantServiceGrpc.getUpdateAgentStatusMethod) == null) {
          MerchantServiceGrpc.getUpdateAgentStatusMethod = getUpdateAgentStatusMethod =
              io.grpc.MethodDescriptor.<com.gameengine.merchant.v1.UpdateAgentStatusRequest, com.gameengine.merchant.v1.UpdateAgentStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateAgentStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateAgentStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.merchant.v1.UpdateAgentStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new MerchantServiceMethodDescriptorSupplier("UpdateAgentStatus"))
              .build();
        }
      }
    }
    return getUpdateAgentStatusMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MerchantServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MerchantServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MerchantServiceStub>() {
        @java.lang.Override
        public MerchantServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MerchantServiceStub(channel, callOptions);
        }
      };
    return MerchantServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static MerchantServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MerchantServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MerchantServiceBlockingV2Stub>() {
        @java.lang.Override
        public MerchantServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MerchantServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return MerchantServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MerchantServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MerchantServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MerchantServiceBlockingStub>() {
        @java.lang.Override
        public MerchantServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MerchantServiceBlockingStub(channel, callOptions);
        }
      };
    return MerchantServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MerchantServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MerchantServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MerchantServiceFutureStub>() {
        @java.lang.Override
        public MerchantServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MerchantServiceFutureStub(channel, callOptions);
        }
      };
    return MerchantServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Player management
     * </pre>
     */
    default void listPlayers(com.gameengine.merchant.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPlayersMethod(), responseObserver);
    }

    /**
     */
    default void getPlayer(com.gameengine.merchant.v1.GetPlayerRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerMethod(), responseObserver);
    }

    /**
     * <pre>
     * Reports
     * </pre>
     */
    default void getRevenueReport(com.gameengine.merchant.v1.GetRevenueReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetRevenueReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetRevenueReportMethod(), responseObserver);
    }

    /**
     */
    default void getPlayerReport(com.gameengine.merchant.v1.GetPlayerReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerReportMethod(), responseObserver);
    }

    /**
     */
    default void getGameReport(com.gameengine.merchant.v1.GetGameReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetGameReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetGameReportMethod(), responseObserver);
    }

    /**
     * <pre>
     * Configuration
     * </pre>
     */
    default void getConfig(com.gameengine.merchant.v1.GetConfigRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConfigMethod(), responseObserver);
    }

    /**
     */
    default void updateConfig(com.gameengine.merchant.v1.UpdateConfigRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateConfigMethod(), responseObserver);
    }

    /**
     * <pre>
     * Webhooks
     * </pre>
     */
    default void registerWebhook(com.gameengine.merchant.v1.RegisterWebhookRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.RegisterWebhookResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterWebhookMethod(), responseObserver);
    }

    /**
     */
    default void listWebhooks(com.gameengine.merchant.v1.ListWebhooksRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListWebhooksResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListWebhooksMethod(), responseObserver);
    }

    /**
     */
    default void deleteWebhook(com.gameengine.merchant.v1.DeleteWebhookRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.DeleteWebhookResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteWebhookMethod(), responseObserver);
    }

    /**
     * <pre>
     * Agent management
     * </pre>
     */
    default void listAgents(com.gameengine.merchant.v1.ListAgentsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListAgentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListAgentsMethod(), responseObserver);
    }

    /**
     */
    default void getAgent(com.gameengine.merchant.v1.GetAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetAgentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAgentMethod(), responseObserver);
    }

    /**
     */
    default void createAgent(com.gameengine.merchant.v1.CreateAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.CreateAgentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAgentMethod(), responseObserver);
    }

    /**
     */
    default void updateAgent(com.gameengine.merchant.v1.UpdateAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateAgentMethod(), responseObserver);
    }

    /**
     */
    default void updateAgentStatus(com.gameengine.merchant.v1.UpdateAgentStatusRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateAgentStatusMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service MerchantService.
   */
  public static abstract class MerchantServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return MerchantServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service MerchantService.
   */
  public static final class MerchantServiceStub
      extends io.grpc.stub.AbstractAsyncStub<MerchantServiceStub> {
    private MerchantServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MerchantServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MerchantServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Player management
     * </pre>
     */
    public void listPlayers(com.gameengine.merchant.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayer(com.gameengine.merchant.v1.GetPlayerRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Reports
     * </pre>
     */
    public void getRevenueReport(com.gameengine.merchant.v1.GetRevenueReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetRevenueReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetRevenueReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayerReport(com.gameengine.merchant.v1.GetPlayerReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getGameReport(com.gameengine.merchant.v1.GetGameReportRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetGameReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetGameReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Configuration
     * </pre>
     */
    public void getConfig(com.gameengine.merchant.v1.GetConfigRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateConfig(com.gameengine.merchant.v1.UpdateConfigRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Webhooks
     * </pre>
     */
    public void registerWebhook(com.gameengine.merchant.v1.RegisterWebhookRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.RegisterWebhookResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterWebhookMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listWebhooks(com.gameengine.merchant.v1.ListWebhooksRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListWebhooksResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListWebhooksMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteWebhook(com.gameengine.merchant.v1.DeleteWebhookRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.DeleteWebhookResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteWebhookMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Agent management
     * </pre>
     */
    public void listAgents(com.gameengine.merchant.v1.ListAgentsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListAgentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListAgentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAgent(com.gameengine.merchant.v1.GetAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetAgentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAgentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createAgent(com.gameengine.merchant.v1.CreateAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.CreateAgentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAgentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateAgent(com.gameengine.merchant.v1.UpdateAgentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateAgentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateAgentStatus(com.gameengine.merchant.v1.UpdateAgentStatusRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateAgentStatusMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service MerchantService.
   */
  public static final class MerchantServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<MerchantServiceBlockingV2Stub> {
    private MerchantServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MerchantServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MerchantServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Player management
     * </pre>
     */
    public com.gameengine.merchant.v1.ListPlayersResponse listPlayers(com.gameengine.merchant.v1.ListPlayersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetPlayerResponse getPlayer(com.gameengine.merchant.v1.GetPlayerRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reports
     * </pre>
     */
    public com.gameengine.merchant.v1.GetRevenueReportResponse getRevenueReport(com.gameengine.merchant.v1.GetRevenueReportRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetRevenueReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetPlayerReportResponse getPlayerReport(com.gameengine.merchant.v1.GetPlayerReportRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetGameReportResponse getGameReport(com.gameengine.merchant.v1.GetGameReportRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetGameReportMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Configuration
     * </pre>
     */
    public com.gameengine.merchant.v1.GetConfigResponse getConfig(com.gameengine.merchant.v1.GetConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateConfigResponse updateConfig(com.gameengine.merchant.v1.UpdateConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateConfigMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Webhooks
     * </pre>
     */
    public com.gameengine.merchant.v1.RegisterWebhookResponse registerWebhook(com.gameengine.merchant.v1.RegisterWebhookRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRegisterWebhookMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.ListWebhooksResponse listWebhooks(com.gameengine.merchant.v1.ListWebhooksRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListWebhooksMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.DeleteWebhookResponse deleteWebhook(com.gameengine.merchant.v1.DeleteWebhookRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getDeleteWebhookMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Agent management
     * </pre>
     */
    public com.gameengine.merchant.v1.ListAgentsResponse listAgents(com.gameengine.merchant.v1.ListAgentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListAgentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetAgentResponse getAgent(com.gameengine.merchant.v1.GetAgentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.CreateAgentResponse createAgent(com.gameengine.merchant.v1.CreateAgentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateAgentResponse updateAgent(com.gameengine.merchant.v1.UpdateAgentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateAgentStatusResponse updateAgentStatus(com.gameengine.merchant.v1.UpdateAgentStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateAgentStatusMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service MerchantService.
   */
  public static final class MerchantServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<MerchantServiceBlockingStub> {
    private MerchantServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MerchantServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MerchantServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Player management
     * </pre>
     */
    public com.gameengine.merchant.v1.ListPlayersResponse listPlayers(com.gameengine.merchant.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetPlayerResponse getPlayer(com.gameengine.merchant.v1.GetPlayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reports
     * </pre>
     */
    public com.gameengine.merchant.v1.GetRevenueReportResponse getRevenueReport(com.gameengine.merchant.v1.GetRevenueReportRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetRevenueReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetPlayerReportResponse getPlayerReport(com.gameengine.merchant.v1.GetPlayerReportRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerReportMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetGameReportResponse getGameReport(com.gameengine.merchant.v1.GetGameReportRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetGameReportMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Configuration
     * </pre>
     */
    public com.gameengine.merchant.v1.GetConfigResponse getConfig(com.gameengine.merchant.v1.GetConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateConfigResponse updateConfig(com.gameengine.merchant.v1.UpdateConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateConfigMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Webhooks
     * </pre>
     */
    public com.gameengine.merchant.v1.RegisterWebhookResponse registerWebhook(com.gameengine.merchant.v1.RegisterWebhookRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterWebhookMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.ListWebhooksResponse listWebhooks(com.gameengine.merchant.v1.ListWebhooksRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListWebhooksMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.DeleteWebhookResponse deleteWebhook(com.gameengine.merchant.v1.DeleteWebhookRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteWebhookMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Agent management
     * </pre>
     */
    public com.gameengine.merchant.v1.ListAgentsResponse listAgents(com.gameengine.merchant.v1.ListAgentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAgentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.GetAgentResponse getAgent(com.gameengine.merchant.v1.GetAgentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.CreateAgentResponse createAgent(com.gameengine.merchant.v1.CreateAgentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateAgentResponse updateAgent(com.gameengine.merchant.v1.UpdateAgentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateAgentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.merchant.v1.UpdateAgentStatusResponse updateAgentStatus(com.gameengine.merchant.v1.UpdateAgentStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateAgentStatusMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service MerchantService.
   */
  public static final class MerchantServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<MerchantServiceFutureStub> {
    private MerchantServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MerchantServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MerchantServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Player management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.ListPlayersResponse> listPlayers(
        com.gameengine.merchant.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetPlayerResponse> getPlayer(
        com.gameengine.merchant.v1.GetPlayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Reports
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetRevenueReportResponse> getRevenueReport(
        com.gameengine.merchant.v1.GetRevenueReportRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetRevenueReportMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetPlayerReportResponse> getPlayerReport(
        com.gameengine.merchant.v1.GetPlayerReportRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerReportMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetGameReportResponse> getGameReport(
        com.gameengine.merchant.v1.GetGameReportRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetGameReportMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Configuration
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetConfigResponse> getConfig(
        com.gameengine.merchant.v1.GetConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.UpdateConfigResponse> updateConfig(
        com.gameengine.merchant.v1.UpdateConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateConfigMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Webhooks
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.RegisterWebhookResponse> registerWebhook(
        com.gameengine.merchant.v1.RegisterWebhookRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterWebhookMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.ListWebhooksResponse> listWebhooks(
        com.gameengine.merchant.v1.ListWebhooksRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListWebhooksMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.DeleteWebhookResponse> deleteWebhook(
        com.gameengine.merchant.v1.DeleteWebhookRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteWebhookMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Agent management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.ListAgentsResponse> listAgents(
        com.gameengine.merchant.v1.ListAgentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListAgentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.GetAgentResponse> getAgent(
        com.gameengine.merchant.v1.GetAgentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAgentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.CreateAgentResponse> createAgent(
        com.gameengine.merchant.v1.CreateAgentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAgentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.UpdateAgentResponse> updateAgent(
        com.gameengine.merchant.v1.UpdateAgentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateAgentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.merchant.v1.UpdateAgentStatusResponse> updateAgentStatus(
        com.gameengine.merchant.v1.UpdateAgentStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateAgentStatusMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_PLAYERS = 0;
  private static final int METHODID_GET_PLAYER = 1;
  private static final int METHODID_GET_REVENUE_REPORT = 2;
  private static final int METHODID_GET_PLAYER_REPORT = 3;
  private static final int METHODID_GET_GAME_REPORT = 4;
  private static final int METHODID_GET_CONFIG = 5;
  private static final int METHODID_UPDATE_CONFIG = 6;
  private static final int METHODID_REGISTER_WEBHOOK = 7;
  private static final int METHODID_LIST_WEBHOOKS = 8;
  private static final int METHODID_DELETE_WEBHOOK = 9;
  private static final int METHODID_LIST_AGENTS = 10;
  private static final int METHODID_GET_AGENT = 11;
  private static final int METHODID_CREATE_AGENT = 12;
  private static final int METHODID_UPDATE_AGENT = 13;
  private static final int METHODID_UPDATE_AGENT_STATUS = 14;

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
        case METHODID_LIST_PLAYERS:
          serviceImpl.listPlayers((com.gameengine.merchant.v1.ListPlayersRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListPlayersResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER:
          serviceImpl.getPlayer((com.gameengine.merchant.v1.GetPlayerRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerResponse>) responseObserver);
          break;
        case METHODID_GET_REVENUE_REPORT:
          serviceImpl.getRevenueReport((com.gameengine.merchant.v1.GetRevenueReportRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetRevenueReportResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER_REPORT:
          serviceImpl.getPlayerReport((com.gameengine.merchant.v1.GetPlayerReportRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetPlayerReportResponse>) responseObserver);
          break;
        case METHODID_GET_GAME_REPORT:
          serviceImpl.getGameReport((com.gameengine.merchant.v1.GetGameReportRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetGameReportResponse>) responseObserver);
          break;
        case METHODID_GET_CONFIG:
          serviceImpl.getConfig((com.gameengine.merchant.v1.GetConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetConfigResponse>) responseObserver);
          break;
        case METHODID_UPDATE_CONFIG:
          serviceImpl.updateConfig((com.gameengine.merchant.v1.UpdateConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateConfigResponse>) responseObserver);
          break;
        case METHODID_REGISTER_WEBHOOK:
          serviceImpl.registerWebhook((com.gameengine.merchant.v1.RegisterWebhookRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.RegisterWebhookResponse>) responseObserver);
          break;
        case METHODID_LIST_WEBHOOKS:
          serviceImpl.listWebhooks((com.gameengine.merchant.v1.ListWebhooksRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListWebhooksResponse>) responseObserver);
          break;
        case METHODID_DELETE_WEBHOOK:
          serviceImpl.deleteWebhook((com.gameengine.merchant.v1.DeleteWebhookRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.DeleteWebhookResponse>) responseObserver);
          break;
        case METHODID_LIST_AGENTS:
          serviceImpl.listAgents((com.gameengine.merchant.v1.ListAgentsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.ListAgentsResponse>) responseObserver);
          break;
        case METHODID_GET_AGENT:
          serviceImpl.getAgent((com.gameengine.merchant.v1.GetAgentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.GetAgentResponse>) responseObserver);
          break;
        case METHODID_CREATE_AGENT:
          serviceImpl.createAgent((com.gameengine.merchant.v1.CreateAgentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.CreateAgentResponse>) responseObserver);
          break;
        case METHODID_UPDATE_AGENT:
          serviceImpl.updateAgent((com.gameengine.merchant.v1.UpdateAgentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentResponse>) responseObserver);
          break;
        case METHODID_UPDATE_AGENT_STATUS:
          serviceImpl.updateAgentStatus((com.gameengine.merchant.v1.UpdateAgentStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.merchant.v1.UpdateAgentStatusResponse>) responseObserver);
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
          getListPlayersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.ListPlayersRequest,
              com.gameengine.merchant.v1.ListPlayersResponse>(
                service, METHODID_LIST_PLAYERS)))
        .addMethod(
          getGetPlayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetPlayerRequest,
              com.gameengine.merchant.v1.GetPlayerResponse>(
                service, METHODID_GET_PLAYER)))
        .addMethod(
          getGetRevenueReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetRevenueReportRequest,
              com.gameengine.merchant.v1.GetRevenueReportResponse>(
                service, METHODID_GET_REVENUE_REPORT)))
        .addMethod(
          getGetPlayerReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetPlayerReportRequest,
              com.gameengine.merchant.v1.GetPlayerReportResponse>(
                service, METHODID_GET_PLAYER_REPORT)))
        .addMethod(
          getGetGameReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetGameReportRequest,
              com.gameengine.merchant.v1.GetGameReportResponse>(
                service, METHODID_GET_GAME_REPORT)))
        .addMethod(
          getGetConfigMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetConfigRequest,
              com.gameengine.merchant.v1.GetConfigResponse>(
                service, METHODID_GET_CONFIG)))
        .addMethod(
          getUpdateConfigMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.UpdateConfigRequest,
              com.gameengine.merchant.v1.UpdateConfigResponse>(
                service, METHODID_UPDATE_CONFIG)))
        .addMethod(
          getRegisterWebhookMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.RegisterWebhookRequest,
              com.gameengine.merchant.v1.RegisterWebhookResponse>(
                service, METHODID_REGISTER_WEBHOOK)))
        .addMethod(
          getListWebhooksMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.ListWebhooksRequest,
              com.gameengine.merchant.v1.ListWebhooksResponse>(
                service, METHODID_LIST_WEBHOOKS)))
        .addMethod(
          getDeleteWebhookMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.DeleteWebhookRequest,
              com.gameengine.merchant.v1.DeleteWebhookResponse>(
                service, METHODID_DELETE_WEBHOOK)))
        .addMethod(
          getListAgentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.ListAgentsRequest,
              com.gameengine.merchant.v1.ListAgentsResponse>(
                service, METHODID_LIST_AGENTS)))
        .addMethod(
          getGetAgentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.GetAgentRequest,
              com.gameengine.merchant.v1.GetAgentResponse>(
                service, METHODID_GET_AGENT)))
        .addMethod(
          getCreateAgentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.CreateAgentRequest,
              com.gameengine.merchant.v1.CreateAgentResponse>(
                service, METHODID_CREATE_AGENT)))
        .addMethod(
          getUpdateAgentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.UpdateAgentRequest,
              com.gameengine.merchant.v1.UpdateAgentResponse>(
                service, METHODID_UPDATE_AGENT)))
        .addMethod(
          getUpdateAgentStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.merchant.v1.UpdateAgentStatusRequest,
              com.gameengine.merchant.v1.UpdateAgentStatusResponse>(
                service, METHODID_UPDATE_AGENT_STATUS)))
        .build();
  }

  private static abstract class MerchantServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MerchantServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.gameengine.merchant.v1.MerchantServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MerchantService");
    }
  }

  private static final class MerchantServiceFileDescriptorSupplier
      extends MerchantServiceBaseDescriptorSupplier {
    MerchantServiceFileDescriptorSupplier() {}
  }

  private static final class MerchantServiceMethodDescriptorSupplier
      extends MerchantServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    MerchantServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (MerchantServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MerchantServiceFileDescriptorSupplier())
              .addMethod(getListPlayersMethod())
              .addMethod(getGetPlayerMethod())
              .addMethod(getGetRevenueReportMethod())
              .addMethod(getGetPlayerReportMethod())
              .addMethod(getGetGameReportMethod())
              .addMethod(getGetConfigMethod())
              .addMethod(getUpdateConfigMethod())
              .addMethod(getRegisterWebhookMethod())
              .addMethod(getListWebhooksMethod())
              .addMethod(getDeleteWebhookMethod())
              .addMethod(getListAgentsMethod())
              .addMethod(getGetAgentMethod())
              .addMethod(getCreateAgentMethod())
              .addMethod(getUpdateAgentMethod())
              .addMethod(getUpdateAgentStatusMethod())
              .build();
        }
      }
    }
    return result;
  }
}
