package com.game_engine.auth.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Auth Service - handles user authentication, registration, and session management
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class AuthServiceGrpc {

  private AuthServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.auth.v1.AuthService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.RegisterRequest,
      com.game_engine.auth.v1.RegisterResponse> getRegisterMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Register",
      requestType = com.game_engine.auth.v1.RegisterRequest.class,
      responseType = com.game_engine.auth.v1.RegisterResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.RegisterRequest,
      com.game_engine.auth.v1.RegisterResponse> getRegisterMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.RegisterRequest, com.game_engine.auth.v1.RegisterResponse> getRegisterMethod;
    if ((getRegisterMethod = AuthServiceGrpc.getRegisterMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getRegisterMethod = AuthServiceGrpc.getRegisterMethod) == null) {
          AuthServiceGrpc.getRegisterMethod = getRegisterMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.RegisterRequest, com.game_engine.auth.v1.RegisterResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Register"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.RegisterRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.RegisterResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("Register"))
              .build();
        }
      }
    }
    return getRegisterMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.LoginRequest,
      com.game_engine.auth.v1.LoginResponse> getLoginMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Login",
      requestType = com.game_engine.auth.v1.LoginRequest.class,
      responseType = com.game_engine.auth.v1.LoginResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.LoginRequest,
      com.game_engine.auth.v1.LoginResponse> getLoginMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.LoginRequest, com.game_engine.auth.v1.LoginResponse> getLoginMethod;
    if ((getLoginMethod = AuthServiceGrpc.getLoginMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getLoginMethod = AuthServiceGrpc.getLoginMethod) == null) {
          AuthServiceGrpc.getLoginMethod = getLoginMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.LoginRequest, com.game_engine.auth.v1.LoginResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Login"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.LoginRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.LoginResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("Login"))
              .build();
        }
      }
    }
    return getLoginMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.RefreshTokenRequest,
      com.game_engine.auth.v1.RefreshTokenResponse> getRefreshTokenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RefreshToken",
      requestType = com.game_engine.auth.v1.RefreshTokenRequest.class,
      responseType = com.game_engine.auth.v1.RefreshTokenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.RefreshTokenRequest,
      com.game_engine.auth.v1.RefreshTokenResponse> getRefreshTokenMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.RefreshTokenRequest, com.game_engine.auth.v1.RefreshTokenResponse> getRefreshTokenMethod;
    if ((getRefreshTokenMethod = AuthServiceGrpc.getRefreshTokenMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getRefreshTokenMethod = AuthServiceGrpc.getRefreshTokenMethod) == null) {
          AuthServiceGrpc.getRefreshTokenMethod = getRefreshTokenMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.RefreshTokenRequest, com.game_engine.auth.v1.RefreshTokenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RefreshToken"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.RefreshTokenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.RefreshTokenResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("RefreshToken"))
              .build();
        }
      }
    }
    return getRefreshTokenMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.LogoutRequest,
      com.google.protobuf.Empty> getLogoutMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Logout",
      requestType = com.game_engine.auth.v1.LogoutRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.LogoutRequest,
      com.google.protobuf.Empty> getLogoutMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.LogoutRequest, com.google.protobuf.Empty> getLogoutMethod;
    if ((getLogoutMethod = AuthServiceGrpc.getLogoutMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getLogoutMethod = AuthServiceGrpc.getLogoutMethod) == null) {
          AuthServiceGrpc.getLogoutMethod = getLogoutMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.LogoutRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Logout"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.LogoutRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("Logout"))
              .build();
        }
      }
    }
    return getLogoutMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyEmailRequest,
      com.game_engine.auth.v1.VerifyEmailResponse> getVerifyEmailMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "VerifyEmail",
      requestType = com.game_engine.auth.v1.VerifyEmailRequest.class,
      responseType = com.game_engine.auth.v1.VerifyEmailResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyEmailRequest,
      com.game_engine.auth.v1.VerifyEmailResponse> getVerifyEmailMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyEmailRequest, com.game_engine.auth.v1.VerifyEmailResponse> getVerifyEmailMethod;
    if ((getVerifyEmailMethod = AuthServiceGrpc.getVerifyEmailMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getVerifyEmailMethod = AuthServiceGrpc.getVerifyEmailMethod) == null) {
          AuthServiceGrpc.getVerifyEmailMethod = getVerifyEmailMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.VerifyEmailRequest, com.game_engine.auth.v1.VerifyEmailResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "VerifyEmail"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.VerifyEmailRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.VerifyEmailResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("VerifyEmail"))
              .build();
        }
      }
    }
    return getVerifyEmailMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyPhoneRequest,
      com.game_engine.auth.v1.VerifyPhoneResponse> getVerifyPhoneMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "VerifyPhone",
      requestType = com.game_engine.auth.v1.VerifyPhoneRequest.class,
      responseType = com.game_engine.auth.v1.VerifyPhoneResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyPhoneRequest,
      com.game_engine.auth.v1.VerifyPhoneResponse> getVerifyPhoneMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.VerifyPhoneRequest, com.game_engine.auth.v1.VerifyPhoneResponse> getVerifyPhoneMethod;
    if ((getVerifyPhoneMethod = AuthServiceGrpc.getVerifyPhoneMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getVerifyPhoneMethod = AuthServiceGrpc.getVerifyPhoneMethod) == null) {
          AuthServiceGrpc.getVerifyPhoneMethod = getVerifyPhoneMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.VerifyPhoneRequest, com.game_engine.auth.v1.VerifyPhoneResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "VerifyPhone"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.VerifyPhoneRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.VerifyPhoneResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("VerifyPhone"))
              .build();
        }
      }
    }
    return getVerifyPhoneMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.Enable2FARequest,
      com.game_engine.auth.v1.Enable2FAResponse> getEnable2FAMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Enable2FA",
      requestType = com.game_engine.auth.v1.Enable2FARequest.class,
      responseType = com.game_engine.auth.v1.Enable2FAResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.Enable2FARequest,
      com.game_engine.auth.v1.Enable2FAResponse> getEnable2FAMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.Enable2FARequest, com.game_engine.auth.v1.Enable2FAResponse> getEnable2FAMethod;
    if ((getEnable2FAMethod = AuthServiceGrpc.getEnable2FAMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getEnable2FAMethod = AuthServiceGrpc.getEnable2FAMethod) == null) {
          AuthServiceGrpc.getEnable2FAMethod = getEnable2FAMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.Enable2FARequest, com.game_engine.auth.v1.Enable2FAResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Enable2FA"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.Enable2FARequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.Enable2FAResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("Enable2FA"))
              .build();
        }
      }
    }
    return getEnable2FAMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.Verify2FARequest,
      com.game_engine.auth.v1.Verify2FAResponse> getVerify2FAMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Verify2FA",
      requestType = com.game_engine.auth.v1.Verify2FARequest.class,
      responseType = com.game_engine.auth.v1.Verify2FAResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.Verify2FARequest,
      com.game_engine.auth.v1.Verify2FAResponse> getVerify2FAMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.Verify2FARequest, com.game_engine.auth.v1.Verify2FAResponse> getVerify2FAMethod;
    if ((getVerify2FAMethod = AuthServiceGrpc.getVerify2FAMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getVerify2FAMethod = AuthServiceGrpc.getVerify2FAMethod) == null) {
          AuthServiceGrpc.getVerify2FAMethod = getVerify2FAMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.Verify2FARequest, com.game_engine.auth.v1.Verify2FAResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Verify2FA"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.Verify2FARequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.Verify2FAResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("Verify2FA"))
              .build();
        }
      }
    }
    return getVerify2FAMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResetPassword",
      requestType = com.game_engine.auth.v1.ResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.ResetPasswordRequest,
      com.google.protobuf.Empty> getResetPasswordMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.ResetPasswordRequest, com.google.protobuf.Empty> getResetPasswordMethod;
    if ((getResetPasswordMethod = AuthServiceGrpc.getResetPasswordMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getResetPasswordMethod = AuthServiceGrpc.getResetPasswordMethod) == null) {
          AuthServiceGrpc.getResetPasswordMethod = getResetPasswordMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.ResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.ResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("ResetPassword"))
              .build();
        }
      }
    }
    return getResetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.ConfirmResetPasswordRequest,
      com.google.protobuf.Empty> getConfirmResetPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ConfirmResetPassword",
      requestType = com.game_engine.auth.v1.ConfirmResetPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.ConfirmResetPasswordRequest,
      com.google.protobuf.Empty> getConfirmResetPasswordMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.ConfirmResetPasswordRequest, com.google.protobuf.Empty> getConfirmResetPasswordMethod;
    if ((getConfirmResetPasswordMethod = AuthServiceGrpc.getConfirmResetPasswordMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getConfirmResetPasswordMethod = AuthServiceGrpc.getConfirmResetPasswordMethod) == null) {
          AuthServiceGrpc.getConfirmResetPasswordMethod = getConfirmResetPasswordMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.ConfirmResetPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ConfirmResetPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.ConfirmResetPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("ConfirmResetPassword"))
              .build();
        }
      }
    }
    return getConfirmResetPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.ValidateTokenRequest,
      com.game_engine.auth.v1.ValidateTokenResponse> getValidateTokenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ValidateToken",
      requestType = com.game_engine.auth.v1.ValidateTokenRequest.class,
      responseType = com.game_engine.auth.v1.ValidateTokenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.ValidateTokenRequest,
      com.game_engine.auth.v1.ValidateTokenResponse> getValidateTokenMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.ValidateTokenRequest, com.game_engine.auth.v1.ValidateTokenResponse> getValidateTokenMethod;
    if ((getValidateTokenMethod = AuthServiceGrpc.getValidateTokenMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getValidateTokenMethod = AuthServiceGrpc.getValidateTokenMethod) == null) {
          AuthServiceGrpc.getValidateTokenMethod = getValidateTokenMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.ValidateTokenRequest, com.game_engine.auth.v1.ValidateTokenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ValidateToken"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.ValidateTokenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.ValidateTokenResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("ValidateToken"))
              .build();
        }
      }
    }
    return getValidateTokenMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.auth.v1.ChangePasswordRequest,
      com.google.protobuf.Empty> getChangePasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ChangePassword",
      requestType = com.game_engine.auth.v1.ChangePasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.auth.v1.ChangePasswordRequest,
      com.google.protobuf.Empty> getChangePasswordMethod() {
    io.grpc.MethodDescriptor<com.game_engine.auth.v1.ChangePasswordRequest, com.google.protobuf.Empty> getChangePasswordMethod;
    if ((getChangePasswordMethod = AuthServiceGrpc.getChangePasswordMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getChangePasswordMethod = AuthServiceGrpc.getChangePasswordMethod) == null) {
          AuthServiceGrpc.getChangePasswordMethod = getChangePasswordMethod =
              io.grpc.MethodDescriptor.<com.game_engine.auth.v1.ChangePasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ChangePassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.auth.v1.ChangePasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("ChangePassword"))
              .build();
        }
      }
    }
    return getChangePasswordMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AuthServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceStub>() {
        @java.lang.Override
        public AuthServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceStub(channel, callOptions);
        }
      };
    return AuthServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static AuthServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingV2Stub>() {
        @java.lang.Override
        public AuthServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return AuthServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AuthServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingStub>() {
        @java.lang.Override
        public AuthServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceBlockingStub(channel, callOptions);
        }
      };
    return AuthServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AuthServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceFutureStub>() {
        @java.lang.Override
        public AuthServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceFutureStub(channel, callOptions);
        }
      };
    return AuthServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Register new user
     * </pre>
     */
    default void register(com.game_engine.auth.v1.RegisterRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RegisterResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterMethod(), responseObserver);
    }

    /**
     * <pre>
     * Login with credentials
     * </pre>
     */
    default void login(com.game_engine.auth.v1.LoginRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.LoginResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLoginMethod(), responseObserver);
    }

    /**
     * <pre>
     * Refresh token
     * </pre>
     */
    default void refreshToken(com.game_engine.auth.v1.RefreshTokenRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RefreshTokenResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRefreshTokenMethod(), responseObserver);
    }

    /**
     * <pre>
     * Logout
     * </pre>
     */
    default void logout(com.game_engine.auth.v1.LogoutRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLogoutMethod(), responseObserver);
    }

    /**
     * <pre>
     * Verify email
     * </pre>
     */
    default void verifyEmail(com.game_engine.auth.v1.VerifyEmailRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyEmailResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getVerifyEmailMethod(), responseObserver);
    }

    /**
     * <pre>
     * Verify phone
     * </pre>
     */
    default void verifyPhone(com.game_engine.auth.v1.VerifyPhoneRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyPhoneResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getVerifyPhoneMethod(), responseObserver);
    }

    /**
     * <pre>
     * Enable 2FA
     * </pre>
     */
    default void enable2FA(com.game_engine.auth.v1.Enable2FARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Enable2FAResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getEnable2FAMethod(), responseObserver);
    }

    /**
     * <pre>
     * Verify 2FA
     * </pre>
     */
    default void verify2FA(com.game_engine.auth.v1.Verify2FARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Verify2FAResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getVerify2FAMethod(), responseObserver);
    }

    /**
     * <pre>
     * Reset password
     * </pre>
     */
    default void resetPassword(com.game_engine.auth.v1.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     * Confirm reset password
     * </pre>
     */
    default void confirmResetPassword(com.game_engine.auth.v1.ConfirmResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getConfirmResetPasswordMethod(), responseObserver);
    }

    /**
     * <pre>
     * Validate token (internal)
     * </pre>
     */
    default void validateToken(com.game_engine.auth.v1.ValidateTokenRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.ValidateTokenResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getValidateTokenMethod(), responseObserver);
    }

    /**
     * <pre>
     * Change password
     * </pre>
     */
    default void changePassword(com.game_engine.auth.v1.ChangePasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getChangePasswordMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AuthService.
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public static abstract class AuthServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AuthServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AuthService.
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public static final class AuthServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AuthServiceStub> {
    private AuthServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Register new user
     * </pre>
     */
    public void register(com.game_engine.auth.v1.RegisterRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RegisterResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Login with credentials
     * </pre>
     */
    public void login(com.game_engine.auth.v1.LoginRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.LoginResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Refresh token
     * </pre>
     */
    public void refreshToken(com.game_engine.auth.v1.RefreshTokenRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RefreshTokenResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRefreshTokenMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Logout
     * </pre>
     */
    public void logout(com.game_engine.auth.v1.LogoutRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Verify email
     * </pre>
     */
    public void verifyEmail(com.game_engine.auth.v1.VerifyEmailRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyEmailResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getVerifyEmailMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Verify phone
     * </pre>
     */
    public void verifyPhone(com.game_engine.auth.v1.VerifyPhoneRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyPhoneResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getVerifyPhoneMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Enable 2FA
     * </pre>
     */
    public void enable2FA(com.game_engine.auth.v1.Enable2FARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Enable2FAResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getEnable2FAMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Verify 2FA
     * </pre>
     */
    public void verify2FA(com.game_engine.auth.v1.Verify2FARequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Verify2FAResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getVerify2FAMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Reset password
     * </pre>
     */
    public void resetPassword(com.game_engine.auth.v1.ResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Confirm reset password
     * </pre>
     */
    public void confirmResetPassword(com.game_engine.auth.v1.ConfirmResetPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getConfirmResetPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Validate token (internal)
     * </pre>
     */
    public void validateToken(com.game_engine.auth.v1.ValidateTokenRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.auth.v1.ValidateTokenResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getValidateTokenMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Change password
     * </pre>
     */
    public void changePassword(com.game_engine.auth.v1.ChangePasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getChangePasswordMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AuthService.
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public static final class AuthServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<AuthServiceBlockingV2Stub> {
    private AuthServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Register new user
     * </pre>
     */
    public com.game_engine.auth.v1.RegisterResponse register(com.game_engine.auth.v1.RegisterRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRegisterMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Login with credentials
     * </pre>
     */
    public com.game_engine.auth.v1.LoginResponse login(com.game_engine.auth.v1.LoginRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getLoginMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Refresh token
     * </pre>
     */
    public com.game_engine.auth.v1.RefreshTokenResponse refreshToken(com.game_engine.auth.v1.RefreshTokenRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRefreshTokenMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Logout
     * </pre>
     */
    public com.google.protobuf.Empty logout(com.game_engine.auth.v1.LogoutRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getLogoutMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify email
     * </pre>
     */
    public com.game_engine.auth.v1.VerifyEmailResponse verifyEmail(com.game_engine.auth.v1.VerifyEmailRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getVerifyEmailMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify phone
     * </pre>
     */
    public com.game_engine.auth.v1.VerifyPhoneResponse verifyPhone(com.game_engine.auth.v1.VerifyPhoneRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getVerifyPhoneMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Enable 2FA
     * </pre>
     */
    public com.game_engine.auth.v1.Enable2FAResponse enable2FA(com.game_engine.auth.v1.Enable2FARequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getEnable2FAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify 2FA
     * </pre>
     */
    public com.game_engine.auth.v1.Verify2FAResponse verify2FA(com.game_engine.auth.v1.Verify2FARequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getVerify2FAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reset password
     * </pre>
     */
    public com.google.protobuf.Empty resetPassword(com.game_engine.auth.v1.ResetPasswordRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Confirm reset password
     * </pre>
     */
    public com.google.protobuf.Empty confirmResetPassword(com.game_engine.auth.v1.ConfirmResetPasswordRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getConfirmResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Validate token (internal)
     * </pre>
     */
    public com.game_engine.auth.v1.ValidateTokenResponse validateToken(com.game_engine.auth.v1.ValidateTokenRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getValidateTokenMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Change password
     * </pre>
     */
    public com.google.protobuf.Empty changePassword(com.game_engine.auth.v1.ChangePasswordRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getChangePasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service AuthService.
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public static final class AuthServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AuthServiceBlockingStub> {
    private AuthServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Register new user
     * </pre>
     */
    public com.game_engine.auth.v1.RegisterResponse register(com.game_engine.auth.v1.RegisterRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Login with credentials
     * </pre>
     */
    public com.game_engine.auth.v1.LoginResponse login(com.game_engine.auth.v1.LoginRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLoginMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Refresh token
     * </pre>
     */
    public com.game_engine.auth.v1.RefreshTokenResponse refreshToken(com.game_engine.auth.v1.RefreshTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRefreshTokenMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Logout
     * </pre>
     */
    public com.google.protobuf.Empty logout(com.game_engine.auth.v1.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLogoutMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify email
     * </pre>
     */
    public com.game_engine.auth.v1.VerifyEmailResponse verifyEmail(com.game_engine.auth.v1.VerifyEmailRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getVerifyEmailMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify phone
     * </pre>
     */
    public com.game_engine.auth.v1.VerifyPhoneResponse verifyPhone(com.game_engine.auth.v1.VerifyPhoneRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getVerifyPhoneMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Enable 2FA
     * </pre>
     */
    public com.game_engine.auth.v1.Enable2FAResponse enable2FA(com.game_engine.auth.v1.Enable2FARequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getEnable2FAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Verify 2FA
     * </pre>
     */
    public com.game_engine.auth.v1.Verify2FAResponse verify2FA(com.game_engine.auth.v1.Verify2FARequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getVerify2FAMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Reset password
     * </pre>
     */
    public com.google.protobuf.Empty resetPassword(com.game_engine.auth.v1.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Confirm reset password
     * </pre>
     */
    public com.google.protobuf.Empty confirmResetPassword(com.game_engine.auth.v1.ConfirmResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getConfirmResetPasswordMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Validate token (internal)
     * </pre>
     */
    public com.game_engine.auth.v1.ValidateTokenResponse validateToken(com.game_engine.auth.v1.ValidateTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getValidateTokenMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Change password
     * </pre>
     */
    public com.google.protobuf.Empty changePassword(com.game_engine.auth.v1.ChangePasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getChangePasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AuthService.
   * <pre>
   * Auth Service - handles user authentication, registration, and session management
   * </pre>
   */
  public static final class AuthServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AuthServiceFutureStub> {
    private AuthServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Register new user
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.RegisterResponse> register(
        com.game_engine.auth.v1.RegisterRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Login with credentials
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.LoginResponse> login(
        com.game_engine.auth.v1.LoginRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLoginMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Refresh token
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.RefreshTokenResponse> refreshToken(
        com.game_engine.auth.v1.RefreshTokenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRefreshTokenMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Logout
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> logout(
        com.game_engine.auth.v1.LogoutRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLogoutMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Verify email
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.VerifyEmailResponse> verifyEmail(
        com.game_engine.auth.v1.VerifyEmailRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getVerifyEmailMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Verify phone
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.VerifyPhoneResponse> verifyPhone(
        com.game_engine.auth.v1.VerifyPhoneRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getVerifyPhoneMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Enable 2FA
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.Enable2FAResponse> enable2FA(
        com.game_engine.auth.v1.Enable2FARequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getEnable2FAMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Verify 2FA
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.Verify2FAResponse> verify2FA(
        com.game_engine.auth.v1.Verify2FARequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getVerify2FAMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Reset password
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> resetPassword(
        com.game_engine.auth.v1.ResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Confirm reset password
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> confirmResetPassword(
        com.game_engine.auth.v1.ConfirmResetPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getConfirmResetPasswordMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Validate token (internal)
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.auth.v1.ValidateTokenResponse> validateToken(
        com.game_engine.auth.v1.ValidateTokenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getValidateTokenMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Change password
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> changePassword(
        com.game_engine.auth.v1.ChangePasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getChangePasswordMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_REGISTER = 0;
  private static final int METHODID_LOGIN = 1;
  private static final int METHODID_REFRESH_TOKEN = 2;
  private static final int METHODID_LOGOUT = 3;
  private static final int METHODID_VERIFY_EMAIL = 4;
  private static final int METHODID_VERIFY_PHONE = 5;
  private static final int METHODID_ENABLE2FA = 6;
  private static final int METHODID_VERIFY2FA = 7;
  private static final int METHODID_RESET_PASSWORD = 8;
  private static final int METHODID_CONFIRM_RESET_PASSWORD = 9;
  private static final int METHODID_VALIDATE_TOKEN = 10;
  private static final int METHODID_CHANGE_PASSWORD = 11;

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
        case METHODID_REGISTER:
          serviceImpl.register((com.game_engine.auth.v1.RegisterRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RegisterResponse>) responseObserver);
          break;
        case METHODID_LOGIN:
          serviceImpl.login((com.game_engine.auth.v1.LoginRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.LoginResponse>) responseObserver);
          break;
        case METHODID_REFRESH_TOKEN:
          serviceImpl.refreshToken((com.game_engine.auth.v1.RefreshTokenRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.RefreshTokenResponse>) responseObserver);
          break;
        case METHODID_LOGOUT:
          serviceImpl.logout((com.game_engine.auth.v1.LogoutRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_VERIFY_EMAIL:
          serviceImpl.verifyEmail((com.game_engine.auth.v1.VerifyEmailRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyEmailResponse>) responseObserver);
          break;
        case METHODID_VERIFY_PHONE:
          serviceImpl.verifyPhone((com.game_engine.auth.v1.VerifyPhoneRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.VerifyPhoneResponse>) responseObserver);
          break;
        case METHODID_ENABLE2FA:
          serviceImpl.enable2FA((com.game_engine.auth.v1.Enable2FARequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Enable2FAResponse>) responseObserver);
          break;
        case METHODID_VERIFY2FA:
          serviceImpl.verify2FA((com.game_engine.auth.v1.Verify2FARequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.Verify2FAResponse>) responseObserver);
          break;
        case METHODID_RESET_PASSWORD:
          serviceImpl.resetPassword((com.game_engine.auth.v1.ResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_CONFIRM_RESET_PASSWORD:
          serviceImpl.confirmResetPassword((com.game_engine.auth.v1.ConfirmResetPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_VALIDATE_TOKEN:
          serviceImpl.validateToken((com.game_engine.auth.v1.ValidateTokenRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.auth.v1.ValidateTokenResponse>) responseObserver);
          break;
        case METHODID_CHANGE_PASSWORD:
          serviceImpl.changePassword((com.game_engine.auth.v1.ChangePasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
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
          getRegisterMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.RegisterRequest,
              com.game_engine.auth.v1.RegisterResponse>(
                service, METHODID_REGISTER)))
        .addMethod(
          getLoginMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.LoginRequest,
              com.game_engine.auth.v1.LoginResponse>(
                service, METHODID_LOGIN)))
        .addMethod(
          getRefreshTokenMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.RefreshTokenRequest,
              com.game_engine.auth.v1.RefreshTokenResponse>(
                service, METHODID_REFRESH_TOKEN)))
        .addMethod(
          getLogoutMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.LogoutRequest,
              com.google.protobuf.Empty>(
                service, METHODID_LOGOUT)))
        .addMethod(
          getVerifyEmailMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.VerifyEmailRequest,
              com.game_engine.auth.v1.VerifyEmailResponse>(
                service, METHODID_VERIFY_EMAIL)))
        .addMethod(
          getVerifyPhoneMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.VerifyPhoneRequest,
              com.game_engine.auth.v1.VerifyPhoneResponse>(
                service, METHODID_VERIFY_PHONE)))
        .addMethod(
          getEnable2FAMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.Enable2FARequest,
              com.game_engine.auth.v1.Enable2FAResponse>(
                service, METHODID_ENABLE2FA)))
        .addMethod(
          getVerify2FAMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.Verify2FARequest,
              com.game_engine.auth.v1.Verify2FAResponse>(
                service, METHODID_VERIFY2FA)))
        .addMethod(
          getResetPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.ResetPasswordRequest,
              com.google.protobuf.Empty>(
                service, METHODID_RESET_PASSWORD)))
        .addMethod(
          getConfirmResetPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.ConfirmResetPasswordRequest,
              com.google.protobuf.Empty>(
                service, METHODID_CONFIRM_RESET_PASSWORD)))
        .addMethod(
          getValidateTokenMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.ValidateTokenRequest,
              com.game_engine.auth.v1.ValidateTokenResponse>(
                service, METHODID_VALIDATE_TOKEN)))
        .addMethod(
          getChangePasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.auth.v1.ChangePasswordRequest,
              com.google.protobuf.Empty>(
                service, METHODID_CHANGE_PASSWORD)))
        .build();
  }

  private static abstract class AuthServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AuthServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.auth.v1.AuthServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AuthService");
    }
  }

  private static final class AuthServiceFileDescriptorSupplier
      extends AuthServiceBaseDescriptorSupplier {
    AuthServiceFileDescriptorSupplier() {}
  }

  private static final class AuthServiceMethodDescriptorSupplier
      extends AuthServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AuthServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (AuthServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AuthServiceFileDescriptorSupplier())
              .addMethod(getRegisterMethod())
              .addMethod(getLoginMethod())
              .addMethod(getRefreshTokenMethod())
              .addMethod(getLogoutMethod())
              .addMethod(getVerifyEmailMethod())
              .addMethod(getVerifyPhoneMethod())
              .addMethod(getEnable2FAMethod())
              .addMethod(getVerify2FAMethod())
              .addMethod(getResetPasswordMethod())
              .addMethod(getConfirmResetPasswordMethod())
              .addMethod(getValidateTokenMethod())
              .addMethod(getChangePasswordMethod())
              .build();
        }
      }
    }
    return result;
  }
}
