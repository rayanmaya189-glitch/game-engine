package com.game_engine.user.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * User Service - handles player profiles, KYC, limits, and settings management
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class UserServiceGrpc {

  private UserServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.user.v1.UserService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.GetProfileRequest,
      com.game_engine.user.v1.GetProfileResponse> getGetProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetProfile",
      requestType = com.game_engine.user.v1.GetProfileRequest.class,
      responseType = com.game_engine.user.v1.GetProfileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.GetProfileRequest,
      com.game_engine.user.v1.GetProfileResponse> getGetProfileMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.GetProfileRequest, com.game_engine.user.v1.GetProfileResponse> getGetProfileMethod;
    if ((getGetProfileMethod = UserServiceGrpc.getGetProfileMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getGetProfileMethod = UserServiceGrpc.getGetProfileMethod) == null) {
          UserServiceGrpc.getGetProfileMethod = getGetProfileMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.GetProfileRequest, com.game_engine.user.v1.GetProfileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetProfileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetProfileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("GetProfile"))
              .build();
        }
      }
    }
    return getGetProfileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdateProfileRequest,
      com.game_engine.user.v1.UpdateProfileResponse> getUpdateProfileMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateProfile",
      requestType = com.game_engine.user.v1.UpdateProfileRequest.class,
      responseType = com.game_engine.user.v1.UpdateProfileResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdateProfileRequest,
      com.game_engine.user.v1.UpdateProfileResponse> getUpdateProfileMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdateProfileRequest, com.game_engine.user.v1.UpdateProfileResponse> getUpdateProfileMethod;
    if ((getUpdateProfileMethod = UserServiceGrpc.getUpdateProfileMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getUpdateProfileMethod = UserServiceGrpc.getUpdateProfileMethod) == null) {
          UserServiceGrpc.getUpdateProfileMethod = getUpdateProfileMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.UpdateProfileRequest, com.game_engine.user.v1.UpdateProfileResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateProfile"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdateProfileRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdateProfileResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("UpdateProfile"))
              .build();
        }
      }
    }
    return getUpdateProfileMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.GetKYCStatusRequest,
      com.game_engine.user.v1.GetKYCStatusResponse> getGetKYCStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetKYCStatus",
      requestType = com.game_engine.user.v1.GetKYCStatusRequest.class,
      responseType = com.game_engine.user.v1.GetKYCStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.GetKYCStatusRequest,
      com.game_engine.user.v1.GetKYCStatusResponse> getGetKYCStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.GetKYCStatusRequest, com.game_engine.user.v1.GetKYCStatusResponse> getGetKYCStatusMethod;
    if ((getGetKYCStatusMethod = UserServiceGrpc.getGetKYCStatusMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getGetKYCStatusMethod = UserServiceGrpc.getGetKYCStatusMethod) == null) {
          UserServiceGrpc.getGetKYCStatusMethod = getGetKYCStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.GetKYCStatusRequest, com.game_engine.user.v1.GetKYCStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetKYCStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetKYCStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetKYCStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("GetKYCStatus"))
              .build();
        }
      }
    }
    return getGetKYCStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.SubmitKYCRequest,
      com.game_engine.user.v1.SubmitKYCResponse> getSubmitKYCMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SubmitKYC",
      requestType = com.game_engine.user.v1.SubmitKYCRequest.class,
      responseType = com.game_engine.user.v1.SubmitKYCResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.SubmitKYCRequest,
      com.game_engine.user.v1.SubmitKYCResponse> getSubmitKYCMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.SubmitKYCRequest, com.game_engine.user.v1.SubmitKYCResponse> getSubmitKYCMethod;
    if ((getSubmitKYCMethod = UserServiceGrpc.getSubmitKYCMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getSubmitKYCMethod = UserServiceGrpc.getSubmitKYCMethod) == null) {
          UserServiceGrpc.getSubmitKYCMethod = getSubmitKYCMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.SubmitKYCRequest, com.game_engine.user.v1.SubmitKYCResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SubmitKYC"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SubmitKYCRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SubmitKYCResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("SubmitKYC"))
              .build();
        }
      }
    }
    return getSubmitKYCMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerSettingsRequest,
      com.game_engine.user.v1.GetPlayerSettingsResponse> getGetPlayerSettingsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayerSettings",
      requestType = com.game_engine.user.v1.GetPlayerSettingsRequest.class,
      responseType = com.game_engine.user.v1.GetPlayerSettingsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerSettingsRequest,
      com.game_engine.user.v1.GetPlayerSettingsResponse> getGetPlayerSettingsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerSettingsRequest, com.game_engine.user.v1.GetPlayerSettingsResponse> getGetPlayerSettingsMethod;
    if ((getGetPlayerSettingsMethod = UserServiceGrpc.getGetPlayerSettingsMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getGetPlayerSettingsMethod = UserServiceGrpc.getGetPlayerSettingsMethod) == null) {
          UserServiceGrpc.getGetPlayerSettingsMethod = getGetPlayerSettingsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.GetPlayerSettingsRequest, com.game_engine.user.v1.GetPlayerSettingsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayerSettings"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerSettingsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerSettingsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("GetPlayerSettings"))
              .build();
        }
      }
    }
    return getGetPlayerSettingsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerSettingsRequest,
      com.game_engine.user.v1.UpdatePlayerSettingsResponse> getUpdatePlayerSettingsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePlayerSettings",
      requestType = com.game_engine.user.v1.UpdatePlayerSettingsRequest.class,
      responseType = com.game_engine.user.v1.UpdatePlayerSettingsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerSettingsRequest,
      com.game_engine.user.v1.UpdatePlayerSettingsResponse> getUpdatePlayerSettingsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerSettingsRequest, com.game_engine.user.v1.UpdatePlayerSettingsResponse> getUpdatePlayerSettingsMethod;
    if ((getUpdatePlayerSettingsMethod = UserServiceGrpc.getUpdatePlayerSettingsMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getUpdatePlayerSettingsMethod = UserServiceGrpc.getUpdatePlayerSettingsMethod) == null) {
          UserServiceGrpc.getUpdatePlayerSettingsMethod = getUpdatePlayerSettingsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.UpdatePlayerSettingsRequest, com.game_engine.user.v1.UpdatePlayerSettingsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePlayerSettings"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdatePlayerSettingsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdatePlayerSettingsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("UpdatePlayerSettings"))
              .build();
        }
      }
    }
    return getUpdatePlayerSettingsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerByAdminRequest,
      com.game_engine.user.v1.GetPlayerByAdminResponse> getGetPlayerByAdminMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayerByAdmin",
      requestType = com.game_engine.user.v1.GetPlayerByAdminRequest.class,
      responseType = com.game_engine.user.v1.GetPlayerByAdminResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerByAdminRequest,
      com.game_engine.user.v1.GetPlayerByAdminResponse> getGetPlayerByAdminMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerByAdminRequest, com.game_engine.user.v1.GetPlayerByAdminResponse> getGetPlayerByAdminMethod;
    if ((getGetPlayerByAdminMethod = UserServiceGrpc.getGetPlayerByAdminMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getGetPlayerByAdminMethod = UserServiceGrpc.getGetPlayerByAdminMethod) == null) {
          UserServiceGrpc.getGetPlayerByAdminMethod = getGetPlayerByAdminMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.GetPlayerByAdminRequest, com.game_engine.user.v1.GetPlayerByAdminResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayerByAdmin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerByAdminRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerByAdminResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("GetPlayerByAdmin"))
              .build();
        }
      }
    }
    return getGetPlayerByAdminMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.ListPlayersRequest,
      com.game_engine.user.v1.ListPlayersResponse> getListPlayersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPlayers",
      requestType = com.game_engine.user.v1.ListPlayersRequest.class,
      responseType = com.game_engine.user.v1.ListPlayersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.ListPlayersRequest,
      com.game_engine.user.v1.ListPlayersResponse> getListPlayersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.ListPlayersRequest, com.game_engine.user.v1.ListPlayersResponse> getListPlayersMethod;
    if ((getListPlayersMethod = UserServiceGrpc.getListPlayersMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getListPlayersMethod = UserServiceGrpc.getListPlayersMethod) == null) {
          UserServiceGrpc.getListPlayersMethod = getListPlayersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.ListPlayersRequest, com.game_engine.user.v1.ListPlayersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPlayers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.ListPlayersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.ListPlayersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("ListPlayers"))
              .build();
        }
      }
    }
    return getListPlayersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerStatusRequest,
      com.game_engine.user.v1.UpdatePlayerStatusResponse> getUpdatePlayerStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePlayerStatus",
      requestType = com.game_engine.user.v1.UpdatePlayerStatusRequest.class,
      responseType = com.game_engine.user.v1.UpdatePlayerStatusResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerStatusRequest,
      com.game_engine.user.v1.UpdatePlayerStatusResponse> getUpdatePlayerStatusMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.UpdatePlayerStatusRequest, com.game_engine.user.v1.UpdatePlayerStatusResponse> getUpdatePlayerStatusMethod;
    if ((getUpdatePlayerStatusMethod = UserServiceGrpc.getUpdatePlayerStatusMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getUpdatePlayerStatusMethod = UserServiceGrpc.getUpdatePlayerStatusMethod) == null) {
          UserServiceGrpc.getUpdatePlayerStatusMethod = getUpdatePlayerStatusMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.UpdatePlayerStatusRequest, com.game_engine.user.v1.UpdatePlayerStatusResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePlayerStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdatePlayerStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.UpdatePlayerStatusResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("UpdatePlayerStatus"))
              .build();
        }
      }
    }
    return getUpdatePlayerStatusMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.SetDepositLimitRequest,
      com.game_engine.user.v1.SetDepositLimitResponse> getSetDepositLimitMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetDepositLimit",
      requestType = com.game_engine.user.v1.SetDepositLimitRequest.class,
      responseType = com.game_engine.user.v1.SetDepositLimitResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.SetDepositLimitRequest,
      com.game_engine.user.v1.SetDepositLimitResponse> getSetDepositLimitMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.SetDepositLimitRequest, com.game_engine.user.v1.SetDepositLimitResponse> getSetDepositLimitMethod;
    if ((getSetDepositLimitMethod = UserServiceGrpc.getSetDepositLimitMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getSetDepositLimitMethod = UserServiceGrpc.getSetDepositLimitMethod) == null) {
          UserServiceGrpc.getSetDepositLimitMethod = getSetDepositLimitMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.SetDepositLimitRequest, com.game_engine.user.v1.SetDepositLimitResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetDepositLimit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetDepositLimitRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetDepositLimitResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("SetDepositLimit"))
              .build();
        }
      }
    }
    return getSetDepositLimitMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.SetBetLimitRequest,
      com.game_engine.user.v1.SetBetLimitResponse> getSetBetLimitMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetBetLimit",
      requestType = com.game_engine.user.v1.SetBetLimitRequest.class,
      responseType = com.game_engine.user.v1.SetBetLimitResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.SetBetLimitRequest,
      com.game_engine.user.v1.SetBetLimitResponse> getSetBetLimitMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.SetBetLimitRequest, com.game_engine.user.v1.SetBetLimitResponse> getSetBetLimitMethod;
    if ((getSetBetLimitMethod = UserServiceGrpc.getSetBetLimitMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getSetBetLimitMethod = UserServiceGrpc.getSetBetLimitMethod) == null) {
          UserServiceGrpc.getSetBetLimitMethod = getSetBetLimitMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.SetBetLimitRequest, com.game_engine.user.v1.SetBetLimitResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetBetLimit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetBetLimitRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetBetLimitResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("SetBetLimit"))
              .build();
        }
      }
    }
    return getSetBetLimitMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.SetLossLimitRequest,
      com.game_engine.user.v1.SetLossLimitResponse> getSetLossLimitMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetLossLimit",
      requestType = com.game_engine.user.v1.SetLossLimitRequest.class,
      responseType = com.game_engine.user.v1.SetLossLimitResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.SetLossLimitRequest,
      com.game_engine.user.v1.SetLossLimitResponse> getSetLossLimitMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.SetLossLimitRequest, com.game_engine.user.v1.SetLossLimitResponse> getSetLossLimitMethod;
    if ((getSetLossLimitMethod = UserServiceGrpc.getSetLossLimitMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getSetLossLimitMethod = UserServiceGrpc.getSetLossLimitMethod) == null) {
          UserServiceGrpc.getSetLossLimitMethod = getSetLossLimitMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.SetLossLimitRequest, com.game_engine.user.v1.SetLossLimitResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetLossLimit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetLossLimitRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SetLossLimitResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("SetLossLimit"))
              .build();
        }
      }
    }
    return getSetLossLimitMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerLimitsRequest,
      com.game_engine.user.v1.GetPlayerLimitsResponse> getGetPlayerLimitsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayerLimits",
      requestType = com.game_engine.user.v1.GetPlayerLimitsRequest.class,
      responseType = com.game_engine.user.v1.GetPlayerLimitsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerLimitsRequest,
      com.game_engine.user.v1.GetPlayerLimitsResponse> getGetPlayerLimitsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.GetPlayerLimitsRequest, com.game_engine.user.v1.GetPlayerLimitsResponse> getGetPlayerLimitsMethod;
    if ((getGetPlayerLimitsMethod = UserServiceGrpc.getGetPlayerLimitsMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getGetPlayerLimitsMethod = UserServiceGrpc.getGetPlayerLimitsMethod) == null) {
          UserServiceGrpc.getGetPlayerLimitsMethod = getGetPlayerLimitsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.GetPlayerLimitsRequest, com.game_engine.user.v1.GetPlayerLimitsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayerLimits"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerLimitsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.GetPlayerLimitsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("GetPlayerLimits"))
              .build();
        }
      }
    }
    return getGetPlayerLimitsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.user.v1.SelfExcludeRequest,
      com.game_engine.user.v1.SelfExcludeResponse> getSelfExcludeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SelfExclude",
      requestType = com.game_engine.user.v1.SelfExcludeRequest.class,
      responseType = com.game_engine.user.v1.SelfExcludeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.user.v1.SelfExcludeRequest,
      com.game_engine.user.v1.SelfExcludeResponse> getSelfExcludeMethod() {
    io.grpc.MethodDescriptor<com.game_engine.user.v1.SelfExcludeRequest, com.game_engine.user.v1.SelfExcludeResponse> getSelfExcludeMethod;
    if ((getSelfExcludeMethod = UserServiceGrpc.getSelfExcludeMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getSelfExcludeMethod = UserServiceGrpc.getSelfExcludeMethod) == null) {
          UserServiceGrpc.getSelfExcludeMethod = getSelfExcludeMethod =
              io.grpc.MethodDescriptor.<com.game_engine.user.v1.SelfExcludeRequest, com.game_engine.user.v1.SelfExcludeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SelfExclude"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SelfExcludeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.user.v1.SelfExcludeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("SelfExclude"))
              .build();
        }
      }
    }
    return getSelfExcludeMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceStub>() {
        @java.lang.Override
        public UserServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceStub(channel, callOptions);
        }
      };
    return UserServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static UserServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingV2Stub>() {
        @java.lang.Override
        public UserServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return UserServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingStub>() {
        @java.lang.Override
        public UserServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceBlockingStub(channel, callOptions);
        }
      };
    return UserServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceFutureStub>() {
        @java.lang.Override
        public UserServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceFutureStub(channel, callOptions);
        }
      };
    return UserServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void getProfile(com.game_engine.user.v1.GetProfileRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetProfileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetProfileMethod(), responseObserver);
    }

    /**
     */
    default void updateProfile(com.game_engine.user.v1.UpdateProfileRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdateProfileResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateProfileMethod(), responseObserver);
    }

    /**
     */
    default void getKYCStatus(com.game_engine.user.v1.GetKYCStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetKYCStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetKYCStatusMethod(), responseObserver);
    }

    /**
     */
    default void submitKYC(com.game_engine.user.v1.SubmitKYCRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SubmitKYCResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSubmitKYCMethod(), responseObserver);
    }

    /**
     */
    default void getPlayerSettings(com.game_engine.user.v1.GetPlayerSettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerSettingsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerSettingsMethod(), responseObserver);
    }

    /**
     */
    default void updatePlayerSettings(com.game_engine.user.v1.UpdatePlayerSettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerSettingsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePlayerSettingsMethod(), responseObserver);
    }

    /**
     */
    default void getPlayerByAdmin(com.game_engine.user.v1.GetPlayerByAdminRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerByAdminResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerByAdminMethod(), responseObserver);
    }

    /**
     */
    default void listPlayers(com.game_engine.user.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPlayersMethod(), responseObserver);
    }

    /**
     */
    default void updatePlayerStatus(com.game_engine.user.v1.UpdatePlayerStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerStatusResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePlayerStatusMethod(), responseObserver);
    }

    /**
     * <pre>
     * Responsible Gaming - Limits management
     * </pre>
     */
    default void setDepositLimit(com.game_engine.user.v1.SetDepositLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetDepositLimitResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetDepositLimitMethod(), responseObserver);
    }

    /**
     */
    default void setBetLimit(com.game_engine.user.v1.SetBetLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetBetLimitResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetBetLimitMethod(), responseObserver);
    }

    /**
     */
    default void setLossLimit(com.game_engine.user.v1.SetLossLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetLossLimitResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetLossLimitMethod(), responseObserver);
    }

    /**
     */
    default void getPlayerLimits(com.game_engine.user.v1.GetPlayerLimitsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerLimitsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerLimitsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Self-Exclusion
     * </pre>
     */
    default void selfExclude(com.game_engine.user.v1.SelfExcludeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SelfExcludeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSelfExcludeMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service UserService.
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public static abstract class UserServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return UserServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service UserService.
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public static final class UserServiceStub
      extends io.grpc.stub.AbstractAsyncStub<UserServiceStub> {
    private UserServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceStub(channel, callOptions);
    }

    /**
     */
    public void getProfile(com.game_engine.user.v1.GetProfileRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetProfileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetProfileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateProfile(com.game_engine.user.v1.UpdateProfileRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdateProfileResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getKYCStatus(com.game_engine.user.v1.GetKYCStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetKYCStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetKYCStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void submitKYC(com.game_engine.user.v1.SubmitKYCRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SubmitKYCResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSubmitKYCMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayerSettings(com.game_engine.user.v1.GetPlayerSettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerSettingsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerSettingsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePlayerSettings(com.game_engine.user.v1.UpdatePlayerSettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerSettingsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePlayerSettingsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayerByAdmin(com.game_engine.user.v1.GetPlayerByAdminRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerByAdminResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerByAdminMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listPlayers(com.game_engine.user.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePlayerStatus(com.game_engine.user.v1.UpdatePlayerStatusRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerStatusResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePlayerStatusMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Responsible Gaming - Limits management
     * </pre>
     */
    public void setDepositLimit(com.game_engine.user.v1.SetDepositLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetDepositLimitResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetDepositLimitMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setBetLimit(com.game_engine.user.v1.SetBetLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetBetLimitResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetBetLimitMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void setLossLimit(com.game_engine.user.v1.SetLossLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetLossLimitResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetLossLimitMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayerLimits(com.game_engine.user.v1.GetPlayerLimitsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerLimitsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerLimitsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Self-Exclusion
     * </pre>
     */
    public void selfExclude(com.game_engine.user.v1.SelfExcludeRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.user.v1.SelfExcludeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSelfExcludeMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service UserService.
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public static final class UserServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<UserServiceBlockingV2Stub> {
    private UserServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.user.v1.GetProfileResponse getProfile(com.game_engine.user.v1.GetProfileRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdateProfileResponse updateProfile(com.game_engine.user.v1.UpdateProfileRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetKYCStatusResponse getKYCStatus(com.game_engine.user.v1.GetKYCStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetKYCStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SubmitKYCResponse submitKYC(com.game_engine.user.v1.SubmitKYCRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSubmitKYCMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerSettingsResponse getPlayerSettings(com.game_engine.user.v1.GetPlayerSettingsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerSettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdatePlayerSettingsResponse updatePlayerSettings(com.game_engine.user.v1.UpdatePlayerSettingsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePlayerSettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerByAdminResponse getPlayerByAdmin(com.game_engine.user.v1.GetPlayerByAdminRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerByAdminMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.ListPlayersResponse listPlayers(com.game_engine.user.v1.ListPlayersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdatePlayerStatusResponse updatePlayerStatus(com.game_engine.user.v1.UpdatePlayerStatusRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePlayerStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Responsible Gaming - Limits management
     * </pre>
     */
    public com.game_engine.user.v1.SetDepositLimitResponse setDepositLimit(com.game_engine.user.v1.SetDepositLimitRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSetDepositLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SetBetLimitResponse setBetLimit(com.game_engine.user.v1.SetBetLimitRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSetBetLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SetLossLimitResponse setLossLimit(com.game_engine.user.v1.SetLossLimitRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSetLossLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerLimitsResponse getPlayerLimits(com.game_engine.user.v1.GetPlayerLimitsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerLimitsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Self-Exclusion
     * </pre>
     */
    public com.game_engine.user.v1.SelfExcludeResponse selfExclude(com.game_engine.user.v1.SelfExcludeRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSelfExcludeMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service UserService.
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public static final class UserServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<UserServiceBlockingStub> {
    private UserServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.user.v1.GetProfileResponse getProfile(com.game_engine.user.v1.GetProfileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdateProfileResponse updateProfile(com.game_engine.user.v1.UpdateProfileRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateProfileMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetKYCStatusResponse getKYCStatus(com.game_engine.user.v1.GetKYCStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetKYCStatusMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SubmitKYCResponse submitKYC(com.game_engine.user.v1.SubmitKYCRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSubmitKYCMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerSettingsResponse getPlayerSettings(com.game_engine.user.v1.GetPlayerSettingsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerSettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdatePlayerSettingsResponse updatePlayerSettings(com.game_engine.user.v1.UpdatePlayerSettingsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePlayerSettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerByAdminResponse getPlayerByAdmin(com.game_engine.user.v1.GetPlayerByAdminRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerByAdminMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.ListPlayersResponse listPlayers(com.game_engine.user.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.UpdatePlayerStatusResponse updatePlayerStatus(com.game_engine.user.v1.UpdatePlayerStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePlayerStatusMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Responsible Gaming - Limits management
     * </pre>
     */
    public com.game_engine.user.v1.SetDepositLimitResponse setDepositLimit(com.game_engine.user.v1.SetDepositLimitRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetDepositLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SetBetLimitResponse setBetLimit(com.game_engine.user.v1.SetBetLimitRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetBetLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.SetLossLimitResponse setLossLimit(com.game_engine.user.v1.SetLossLimitRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetLossLimitMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.user.v1.GetPlayerLimitsResponse getPlayerLimits(com.game_engine.user.v1.GetPlayerLimitsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerLimitsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Self-Exclusion
     * </pre>
     */
    public com.game_engine.user.v1.SelfExcludeResponse selfExclude(com.game_engine.user.v1.SelfExcludeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSelfExcludeMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service UserService.
   * <pre>
   * User Service - handles player profiles, KYC, limits, and settings management
   * </pre>
   */
  public static final class UserServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<UserServiceFutureStub> {
    private UserServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.GetProfileResponse> getProfile(
        com.game_engine.user.v1.GetProfileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetProfileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.UpdateProfileResponse> updateProfile(
        com.game_engine.user.v1.UpdateProfileRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateProfileMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.GetKYCStatusResponse> getKYCStatus(
        com.game_engine.user.v1.GetKYCStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetKYCStatusMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.SubmitKYCResponse> submitKYC(
        com.game_engine.user.v1.SubmitKYCRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSubmitKYCMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.GetPlayerSettingsResponse> getPlayerSettings(
        com.game_engine.user.v1.GetPlayerSettingsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerSettingsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.UpdatePlayerSettingsResponse> updatePlayerSettings(
        com.game_engine.user.v1.UpdatePlayerSettingsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePlayerSettingsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.GetPlayerByAdminResponse> getPlayerByAdmin(
        com.game_engine.user.v1.GetPlayerByAdminRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerByAdminMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.ListPlayersResponse> listPlayers(
        com.game_engine.user.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.UpdatePlayerStatusResponse> updatePlayerStatus(
        com.game_engine.user.v1.UpdatePlayerStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePlayerStatusMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Responsible Gaming - Limits management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.SetDepositLimitResponse> setDepositLimit(
        com.game_engine.user.v1.SetDepositLimitRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetDepositLimitMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.SetBetLimitResponse> setBetLimit(
        com.game_engine.user.v1.SetBetLimitRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetBetLimitMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.SetLossLimitResponse> setLossLimit(
        com.game_engine.user.v1.SetLossLimitRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetLossLimitMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.GetPlayerLimitsResponse> getPlayerLimits(
        com.game_engine.user.v1.GetPlayerLimitsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerLimitsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Self-Exclusion
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.user.v1.SelfExcludeResponse> selfExclude(
        com.game_engine.user.v1.SelfExcludeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSelfExcludeMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_PROFILE = 0;
  private static final int METHODID_UPDATE_PROFILE = 1;
  private static final int METHODID_GET_KYCSTATUS = 2;
  private static final int METHODID_SUBMIT_KYC = 3;
  private static final int METHODID_GET_PLAYER_SETTINGS = 4;
  private static final int METHODID_UPDATE_PLAYER_SETTINGS = 5;
  private static final int METHODID_GET_PLAYER_BY_ADMIN = 6;
  private static final int METHODID_LIST_PLAYERS = 7;
  private static final int METHODID_UPDATE_PLAYER_STATUS = 8;
  private static final int METHODID_SET_DEPOSIT_LIMIT = 9;
  private static final int METHODID_SET_BET_LIMIT = 10;
  private static final int METHODID_SET_LOSS_LIMIT = 11;
  private static final int METHODID_GET_PLAYER_LIMITS = 12;
  private static final int METHODID_SELF_EXCLUDE = 13;

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
        case METHODID_GET_PROFILE:
          serviceImpl.getProfile((com.game_engine.user.v1.GetProfileRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetProfileResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PROFILE:
          serviceImpl.updateProfile((com.game_engine.user.v1.UpdateProfileRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdateProfileResponse>) responseObserver);
          break;
        case METHODID_GET_KYCSTATUS:
          serviceImpl.getKYCStatus((com.game_engine.user.v1.GetKYCStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetKYCStatusResponse>) responseObserver);
          break;
        case METHODID_SUBMIT_KYC:
          serviceImpl.submitKYC((com.game_engine.user.v1.SubmitKYCRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.SubmitKYCResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER_SETTINGS:
          serviceImpl.getPlayerSettings((com.game_engine.user.v1.GetPlayerSettingsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerSettingsResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PLAYER_SETTINGS:
          serviceImpl.updatePlayerSettings((com.game_engine.user.v1.UpdatePlayerSettingsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerSettingsResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER_BY_ADMIN:
          serviceImpl.getPlayerByAdmin((com.game_engine.user.v1.GetPlayerByAdminRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerByAdminResponse>) responseObserver);
          break;
        case METHODID_LIST_PLAYERS:
          serviceImpl.listPlayers((com.game_engine.user.v1.ListPlayersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.ListPlayersResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PLAYER_STATUS:
          serviceImpl.updatePlayerStatus((com.game_engine.user.v1.UpdatePlayerStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.UpdatePlayerStatusResponse>) responseObserver);
          break;
        case METHODID_SET_DEPOSIT_LIMIT:
          serviceImpl.setDepositLimit((com.game_engine.user.v1.SetDepositLimitRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetDepositLimitResponse>) responseObserver);
          break;
        case METHODID_SET_BET_LIMIT:
          serviceImpl.setBetLimit((com.game_engine.user.v1.SetBetLimitRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetBetLimitResponse>) responseObserver);
          break;
        case METHODID_SET_LOSS_LIMIT:
          serviceImpl.setLossLimit((com.game_engine.user.v1.SetLossLimitRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.SetLossLimitResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER_LIMITS:
          serviceImpl.getPlayerLimits((com.game_engine.user.v1.GetPlayerLimitsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.GetPlayerLimitsResponse>) responseObserver);
          break;
        case METHODID_SELF_EXCLUDE:
          serviceImpl.selfExclude((com.game_engine.user.v1.SelfExcludeRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.user.v1.SelfExcludeResponse>) responseObserver);
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
          getGetProfileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.GetProfileRequest,
              com.game_engine.user.v1.GetProfileResponse>(
                service, METHODID_GET_PROFILE)))
        .addMethod(
          getUpdateProfileMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.UpdateProfileRequest,
              com.game_engine.user.v1.UpdateProfileResponse>(
                service, METHODID_UPDATE_PROFILE)))
        .addMethod(
          getGetKYCStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.GetKYCStatusRequest,
              com.game_engine.user.v1.GetKYCStatusResponse>(
                service, METHODID_GET_KYCSTATUS)))
        .addMethod(
          getSubmitKYCMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.SubmitKYCRequest,
              com.game_engine.user.v1.SubmitKYCResponse>(
                service, METHODID_SUBMIT_KYC)))
        .addMethod(
          getGetPlayerSettingsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.GetPlayerSettingsRequest,
              com.game_engine.user.v1.GetPlayerSettingsResponse>(
                service, METHODID_GET_PLAYER_SETTINGS)))
        .addMethod(
          getUpdatePlayerSettingsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.UpdatePlayerSettingsRequest,
              com.game_engine.user.v1.UpdatePlayerSettingsResponse>(
                service, METHODID_UPDATE_PLAYER_SETTINGS)))
        .addMethod(
          getGetPlayerByAdminMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.GetPlayerByAdminRequest,
              com.game_engine.user.v1.GetPlayerByAdminResponse>(
                service, METHODID_GET_PLAYER_BY_ADMIN)))
        .addMethod(
          getListPlayersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.ListPlayersRequest,
              com.game_engine.user.v1.ListPlayersResponse>(
                service, METHODID_LIST_PLAYERS)))
        .addMethod(
          getUpdatePlayerStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.UpdatePlayerStatusRequest,
              com.game_engine.user.v1.UpdatePlayerStatusResponse>(
                service, METHODID_UPDATE_PLAYER_STATUS)))
        .addMethod(
          getSetDepositLimitMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.SetDepositLimitRequest,
              com.game_engine.user.v1.SetDepositLimitResponse>(
                service, METHODID_SET_DEPOSIT_LIMIT)))
        .addMethod(
          getSetBetLimitMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.SetBetLimitRequest,
              com.game_engine.user.v1.SetBetLimitResponse>(
                service, METHODID_SET_BET_LIMIT)))
        .addMethod(
          getSetLossLimitMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.SetLossLimitRequest,
              com.game_engine.user.v1.SetLossLimitResponse>(
                service, METHODID_SET_LOSS_LIMIT)))
        .addMethod(
          getGetPlayerLimitsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.GetPlayerLimitsRequest,
              com.game_engine.user.v1.GetPlayerLimitsResponse>(
                service, METHODID_GET_PLAYER_LIMITS)))
        .addMethod(
          getSelfExcludeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.user.v1.SelfExcludeRequest,
              com.game_engine.user.v1.SelfExcludeResponse>(
                service, METHODID_SELF_EXCLUDE)))
        .build();
  }

  private static abstract class UserServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.user.v1.UserServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("UserService");
    }
  }

  private static final class UserServiceFileDescriptorSupplier
      extends UserServiceBaseDescriptorSupplier {
    UserServiceFileDescriptorSupplier() {}
  }

  private static final class UserServiceMethodDescriptorSupplier
      extends UserServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    UserServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (UserServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserServiceFileDescriptorSupplier())
              .addMethod(getGetProfileMethod())
              .addMethod(getUpdateProfileMethod())
              .addMethod(getGetKYCStatusMethod())
              .addMethod(getSubmitKYCMethod())
              .addMethod(getGetPlayerSettingsMethod())
              .addMethod(getUpdatePlayerSettingsMethod())
              .addMethod(getGetPlayerByAdminMethod())
              .addMethod(getListPlayersMethod())
              .addMethod(getUpdatePlayerStatusMethod())
              .addMethod(getSetDepositLimitMethod())
              .addMethod(getSetBetLimitMethod())
              .addMethod(getSetLossLimitMethod())
              .addMethod(getGetPlayerLimitsMethod())
              .addMethod(getSelfExcludeMethod())
              .build();
        }
      }
    }
    return result;
  }
}
