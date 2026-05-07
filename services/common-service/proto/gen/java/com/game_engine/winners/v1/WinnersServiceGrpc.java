package com.game_engine.winners.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * WinnersService manages winner showcases and privacy settings
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class WinnersServiceGrpc {

  private WinnersServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.winners.v1.WinnersService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetRecentWinnersRequest,
      com.game_engine.winners.v1.GetRecentWinnersResponse> getGetRecentWinnersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetRecentWinners",
      requestType = com.game_engine.winners.v1.GetRecentWinnersRequest.class,
      responseType = com.game_engine.winners.v1.GetRecentWinnersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetRecentWinnersRequest,
      com.game_engine.winners.v1.GetRecentWinnersResponse> getGetRecentWinnersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetRecentWinnersRequest, com.game_engine.winners.v1.GetRecentWinnersResponse> getGetRecentWinnersMethod;
    if ((getGetRecentWinnersMethod = WinnersServiceGrpc.getGetRecentWinnersMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getGetRecentWinnersMethod = WinnersServiceGrpc.getGetRecentWinnersMethod) == null) {
          WinnersServiceGrpc.getGetRecentWinnersMethod = getGetRecentWinnersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.GetRecentWinnersRequest, com.game_engine.winners.v1.GetRecentWinnersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetRecentWinners"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetRecentWinnersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetRecentWinnersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("GetRecentWinners"))
              .build();
        }
      }
    }
    return getGetRecentWinnersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetBigWinsRequest,
      com.game_engine.winners.v1.GetBigWinsResponse> getGetBigWinsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBigWins",
      requestType = com.game_engine.winners.v1.GetBigWinsRequest.class,
      responseType = com.game_engine.winners.v1.GetBigWinsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetBigWinsRequest,
      com.game_engine.winners.v1.GetBigWinsResponse> getGetBigWinsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetBigWinsRequest, com.game_engine.winners.v1.GetBigWinsResponse> getGetBigWinsMethod;
    if ((getGetBigWinsMethod = WinnersServiceGrpc.getGetBigWinsMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getGetBigWinsMethod = WinnersServiceGrpc.getGetBigWinsMethod) == null) {
          WinnersServiceGrpc.getGetBigWinsMethod = getGetBigWinsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.GetBigWinsRequest, com.game_engine.winners.v1.GetBigWinsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBigWins"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetBigWinsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetBigWinsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("GetBigWins"))
              .build();
        }
      }
    }
    return getGetBigWinsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetJackpotWinnersRequest,
      com.game_engine.winners.v1.GetJackpotWinnersResponse> getGetJackpotWinnersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetJackpotWinners",
      requestType = com.game_engine.winners.v1.GetJackpotWinnersRequest.class,
      responseType = com.game_engine.winners.v1.GetJackpotWinnersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetJackpotWinnersRequest,
      com.game_engine.winners.v1.GetJackpotWinnersResponse> getGetJackpotWinnersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetJackpotWinnersRequest, com.game_engine.winners.v1.GetJackpotWinnersResponse> getGetJackpotWinnersMethod;
    if ((getGetJackpotWinnersMethod = WinnersServiceGrpc.getGetJackpotWinnersMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getGetJackpotWinnersMethod = WinnersServiceGrpc.getGetJackpotWinnersMethod) == null) {
          WinnersServiceGrpc.getGetJackpotWinnersMethod = getGetJackpotWinnersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.GetJackpotWinnersRequest, com.game_engine.winners.v1.GetJackpotWinnersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetJackpotWinners"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetJackpotWinnersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetJackpotWinnersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("GetJackpotWinners"))
              .build();
        }
      }
    }
    return getGetJackpotWinnersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.RecordWinRequest,
      com.game_engine.winners.v1.RecordWinResponse> getRecordWinMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RecordWin",
      requestType = com.game_engine.winners.v1.RecordWinRequest.class,
      responseType = com.game_engine.winners.v1.RecordWinResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.RecordWinRequest,
      com.game_engine.winners.v1.RecordWinResponse> getRecordWinMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.RecordWinRequest, com.game_engine.winners.v1.RecordWinResponse> getRecordWinMethod;
    if ((getRecordWinMethod = WinnersServiceGrpc.getRecordWinMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getRecordWinMethod = WinnersServiceGrpc.getRecordWinMethod) == null) {
          WinnersServiceGrpc.getRecordWinMethod = getRecordWinMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.RecordWinRequest, com.game_engine.winners.v1.RecordWinResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RecordWin"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.RecordWinRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.RecordWinResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("RecordWin"))
              .build();
        }
      }
    }
    return getRecordWinMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetPrivacySettingsRequest,
      com.game_engine.winners.v1.GetPrivacySettingsResponse> getGetPrivacySettingsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPrivacySettings",
      requestType = com.game_engine.winners.v1.GetPrivacySettingsRequest.class,
      responseType = com.game_engine.winners.v1.GetPrivacySettingsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetPrivacySettingsRequest,
      com.game_engine.winners.v1.GetPrivacySettingsResponse> getGetPrivacySettingsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.GetPrivacySettingsRequest, com.game_engine.winners.v1.GetPrivacySettingsResponse> getGetPrivacySettingsMethod;
    if ((getGetPrivacySettingsMethod = WinnersServiceGrpc.getGetPrivacySettingsMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getGetPrivacySettingsMethod = WinnersServiceGrpc.getGetPrivacySettingsMethod) == null) {
          WinnersServiceGrpc.getGetPrivacySettingsMethod = getGetPrivacySettingsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.GetPrivacySettingsRequest, com.game_engine.winners.v1.GetPrivacySettingsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPrivacySettings"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetPrivacySettingsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.GetPrivacySettingsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("GetPrivacySettings"))
              .build();
        }
      }
    }
    return getGetPrivacySettingsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.winners.v1.UpdatePrivacySettingsRequest,
      com.game_engine.winners.v1.UpdatePrivacySettingsResponse> getUpdatePrivacySettingsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePrivacySettings",
      requestType = com.game_engine.winners.v1.UpdatePrivacySettingsRequest.class,
      responseType = com.game_engine.winners.v1.UpdatePrivacySettingsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.winners.v1.UpdatePrivacySettingsRequest,
      com.game_engine.winners.v1.UpdatePrivacySettingsResponse> getUpdatePrivacySettingsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.winners.v1.UpdatePrivacySettingsRequest, com.game_engine.winners.v1.UpdatePrivacySettingsResponse> getUpdatePrivacySettingsMethod;
    if ((getUpdatePrivacySettingsMethod = WinnersServiceGrpc.getUpdatePrivacySettingsMethod) == null) {
      synchronized (WinnersServiceGrpc.class) {
        if ((getUpdatePrivacySettingsMethod = WinnersServiceGrpc.getUpdatePrivacySettingsMethod) == null) {
          WinnersServiceGrpc.getUpdatePrivacySettingsMethod = getUpdatePrivacySettingsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.winners.v1.UpdatePrivacySettingsRequest, com.game_engine.winners.v1.UpdatePrivacySettingsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePrivacySettings"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.UpdatePrivacySettingsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.winners.v1.UpdatePrivacySettingsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WinnersServiceMethodDescriptorSupplier("UpdatePrivacySettings"))
              .build();
        }
      }
    }
    return getUpdatePrivacySettingsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static WinnersServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WinnersServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WinnersServiceStub>() {
        @java.lang.Override
        public WinnersServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WinnersServiceStub(channel, callOptions);
        }
      };
    return WinnersServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static WinnersServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WinnersServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WinnersServiceBlockingV2Stub>() {
        @java.lang.Override
        public WinnersServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WinnersServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return WinnersServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static WinnersServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WinnersServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WinnersServiceBlockingStub>() {
        @java.lang.Override
        public WinnersServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WinnersServiceBlockingStub(channel, callOptions);
        }
      };
    return WinnersServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static WinnersServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WinnersServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WinnersServiceFutureStub>() {
        @java.lang.Override
        public WinnersServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WinnersServiceFutureStub(channel, callOptions);
        }
      };
    return WinnersServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void getRecentWinners(com.game_engine.winners.v1.GetRecentWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetRecentWinnersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetRecentWinnersMethod(), responseObserver);
    }

    /**
     */
    default void getBigWins(com.game_engine.winners.v1.GetBigWinsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetBigWinsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBigWinsMethod(), responseObserver);
    }

    /**
     */
    default void getJackpotWinners(com.game_engine.winners.v1.GetJackpotWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetJackpotWinnersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetJackpotWinnersMethod(), responseObserver);
    }

    /**
     */
    default void recordWin(com.game_engine.winners.v1.RecordWinRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.RecordWinResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRecordWinMethod(), responseObserver);
    }

    /**
     */
    default void getPrivacySettings(com.game_engine.winners.v1.GetPrivacySettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetPrivacySettingsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPrivacySettingsMethod(), responseObserver);
    }

    /**
     */
    default void updatePrivacySettings(com.game_engine.winners.v1.UpdatePrivacySettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.UpdatePrivacySettingsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePrivacySettingsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service WinnersService.
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public static abstract class WinnersServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return WinnersServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service WinnersService.
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public static final class WinnersServiceStub
      extends io.grpc.stub.AbstractAsyncStub<WinnersServiceStub> {
    private WinnersServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WinnersServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WinnersServiceStub(channel, callOptions);
    }

    /**
     */
    public void getRecentWinners(com.game_engine.winners.v1.GetRecentWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetRecentWinnersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetRecentWinnersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getBigWins(com.game_engine.winners.v1.GetBigWinsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetBigWinsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBigWinsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getJackpotWinners(com.game_engine.winners.v1.GetJackpotWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetJackpotWinnersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetJackpotWinnersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void recordWin(com.game_engine.winners.v1.RecordWinRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.RecordWinResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRecordWinMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPrivacySettings(com.game_engine.winners.v1.GetPrivacySettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetPrivacySettingsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPrivacySettingsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePrivacySettings(com.game_engine.winners.v1.UpdatePrivacySettingsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.winners.v1.UpdatePrivacySettingsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePrivacySettingsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service WinnersService.
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public static final class WinnersServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<WinnersServiceBlockingV2Stub> {
    private WinnersServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WinnersServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WinnersServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.winners.v1.GetRecentWinnersResponse getRecentWinners(com.game_engine.winners.v1.GetRecentWinnersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetRecentWinnersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetBigWinsResponse getBigWins(com.game_engine.winners.v1.GetBigWinsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetBigWinsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetJackpotWinnersResponse getJackpotWinners(com.game_engine.winners.v1.GetJackpotWinnersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetJackpotWinnersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.RecordWinResponse recordWin(com.game_engine.winners.v1.RecordWinRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRecordWinMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetPrivacySettingsResponse getPrivacySettings(com.game_engine.winners.v1.GetPrivacySettingsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPrivacySettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.UpdatePrivacySettingsResponse updatePrivacySettings(com.game_engine.winners.v1.UpdatePrivacySettingsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePrivacySettingsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service WinnersService.
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public static final class WinnersServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<WinnersServiceBlockingStub> {
    private WinnersServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WinnersServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WinnersServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.winners.v1.GetRecentWinnersResponse getRecentWinners(com.game_engine.winners.v1.GetRecentWinnersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetRecentWinnersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetBigWinsResponse getBigWins(com.game_engine.winners.v1.GetBigWinsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBigWinsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetJackpotWinnersResponse getJackpotWinners(com.game_engine.winners.v1.GetJackpotWinnersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetJackpotWinnersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.RecordWinResponse recordWin(com.game_engine.winners.v1.RecordWinRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRecordWinMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.GetPrivacySettingsResponse getPrivacySettings(com.game_engine.winners.v1.GetPrivacySettingsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPrivacySettingsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.winners.v1.UpdatePrivacySettingsResponse updatePrivacySettings(com.game_engine.winners.v1.UpdatePrivacySettingsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePrivacySettingsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service WinnersService.
   * <pre>
   * WinnersService manages winner showcases and privacy settings
   * </pre>
   */
  public static final class WinnersServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<WinnersServiceFutureStub> {
    private WinnersServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WinnersServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WinnersServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.GetRecentWinnersResponse> getRecentWinners(
        com.game_engine.winners.v1.GetRecentWinnersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetRecentWinnersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.GetBigWinsResponse> getBigWins(
        com.game_engine.winners.v1.GetBigWinsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBigWinsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.GetJackpotWinnersResponse> getJackpotWinners(
        com.game_engine.winners.v1.GetJackpotWinnersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetJackpotWinnersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.RecordWinResponse> recordWin(
        com.game_engine.winners.v1.RecordWinRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRecordWinMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.GetPrivacySettingsResponse> getPrivacySettings(
        com.game_engine.winners.v1.GetPrivacySettingsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPrivacySettingsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.winners.v1.UpdatePrivacySettingsResponse> updatePrivacySettings(
        com.game_engine.winners.v1.UpdatePrivacySettingsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePrivacySettingsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_RECENT_WINNERS = 0;
  private static final int METHODID_GET_BIG_WINS = 1;
  private static final int METHODID_GET_JACKPOT_WINNERS = 2;
  private static final int METHODID_RECORD_WIN = 3;
  private static final int METHODID_GET_PRIVACY_SETTINGS = 4;
  private static final int METHODID_UPDATE_PRIVACY_SETTINGS = 5;

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
        case METHODID_GET_RECENT_WINNERS:
          serviceImpl.getRecentWinners((com.game_engine.winners.v1.GetRecentWinnersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetRecentWinnersResponse>) responseObserver);
          break;
        case METHODID_GET_BIG_WINS:
          serviceImpl.getBigWins((com.game_engine.winners.v1.GetBigWinsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetBigWinsResponse>) responseObserver);
          break;
        case METHODID_GET_JACKPOT_WINNERS:
          serviceImpl.getJackpotWinners((com.game_engine.winners.v1.GetJackpotWinnersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetJackpotWinnersResponse>) responseObserver);
          break;
        case METHODID_RECORD_WIN:
          serviceImpl.recordWin((com.game_engine.winners.v1.RecordWinRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.RecordWinResponse>) responseObserver);
          break;
        case METHODID_GET_PRIVACY_SETTINGS:
          serviceImpl.getPrivacySettings((com.game_engine.winners.v1.GetPrivacySettingsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.GetPrivacySettingsResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PRIVACY_SETTINGS:
          serviceImpl.updatePrivacySettings((com.game_engine.winners.v1.UpdatePrivacySettingsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.winners.v1.UpdatePrivacySettingsResponse>) responseObserver);
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
          getGetRecentWinnersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.GetRecentWinnersRequest,
              com.game_engine.winners.v1.GetRecentWinnersResponse>(
                service, METHODID_GET_RECENT_WINNERS)))
        .addMethod(
          getGetBigWinsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.GetBigWinsRequest,
              com.game_engine.winners.v1.GetBigWinsResponse>(
                service, METHODID_GET_BIG_WINS)))
        .addMethod(
          getGetJackpotWinnersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.GetJackpotWinnersRequest,
              com.game_engine.winners.v1.GetJackpotWinnersResponse>(
                service, METHODID_GET_JACKPOT_WINNERS)))
        .addMethod(
          getRecordWinMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.RecordWinRequest,
              com.game_engine.winners.v1.RecordWinResponse>(
                service, METHODID_RECORD_WIN)))
        .addMethod(
          getGetPrivacySettingsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.GetPrivacySettingsRequest,
              com.game_engine.winners.v1.GetPrivacySettingsResponse>(
                service, METHODID_GET_PRIVACY_SETTINGS)))
        .addMethod(
          getUpdatePrivacySettingsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.winners.v1.UpdatePrivacySettingsRequest,
              com.game_engine.winners.v1.UpdatePrivacySettingsResponse>(
                service, METHODID_UPDATE_PRIVACY_SETTINGS)))
        .build();
  }

  private static abstract class WinnersServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    WinnersServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.winners.v1.WinnersServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("WinnersService");
    }
  }

  private static final class WinnersServiceFileDescriptorSupplier
      extends WinnersServiceBaseDescriptorSupplier {
    WinnersServiceFileDescriptorSupplier() {}
  }

  private static final class WinnersServiceMethodDescriptorSupplier
      extends WinnersServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    WinnersServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (WinnersServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new WinnersServiceFileDescriptorSupplier())
              .addMethod(getGetRecentWinnersMethod())
              .addMethod(getGetBigWinsMethod())
              .addMethod(getGetJackpotWinnersMethod())
              .addMethod(getRecordWinMethod())
              .addMethod(getGetPrivacySettingsMethod())
              .addMethod(getUpdatePrivacySettingsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
