package com.game_engine.jackpot.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class JackpotServiceGrpc {

  private JackpotServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.jackpot.v1.JackpotService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.ListJackpotsRequest,
      com.game_engine.jackpot.v1.ListJackpotsResponse> getListJackpotsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListJackpots",
      requestType = com.game_engine.jackpot.v1.ListJackpotsRequest.class,
      responseType = com.game_engine.jackpot.v1.ListJackpotsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.ListJackpotsRequest,
      com.game_engine.jackpot.v1.ListJackpotsResponse> getListJackpotsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.ListJackpotsRequest, com.game_engine.jackpot.v1.ListJackpotsResponse> getListJackpotsMethod;
    if ((getListJackpotsMethod = JackpotServiceGrpc.getListJackpotsMethod) == null) {
      synchronized (JackpotServiceGrpc.class) {
        if ((getListJackpotsMethod = JackpotServiceGrpc.getListJackpotsMethod) == null) {
          JackpotServiceGrpc.getListJackpotsMethod = getListJackpotsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.jackpot.v1.ListJackpotsRequest, com.game_engine.jackpot.v1.ListJackpotsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListJackpots"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.ListJackpotsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.ListJackpotsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JackpotServiceMethodDescriptorSupplier("ListJackpots"))
              .build();
        }
      }
    }
    return getListJackpotsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotRequest,
      com.game_engine.jackpot.v1.GetJackpotResponse> getGetJackpotMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetJackpot",
      requestType = com.game_engine.jackpot.v1.GetJackpotRequest.class,
      responseType = com.game_engine.jackpot.v1.GetJackpotResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotRequest,
      com.game_engine.jackpot.v1.GetJackpotResponse> getGetJackpotMethod() {
    io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotRequest, com.game_engine.jackpot.v1.GetJackpotResponse> getGetJackpotMethod;
    if ((getGetJackpotMethod = JackpotServiceGrpc.getGetJackpotMethod) == null) {
      synchronized (JackpotServiceGrpc.class) {
        if ((getGetJackpotMethod = JackpotServiceGrpc.getGetJackpotMethod) == null) {
          JackpotServiceGrpc.getGetJackpotMethod = getGetJackpotMethod =
              io.grpc.MethodDescriptor.<com.game_engine.jackpot.v1.GetJackpotRequest, com.game_engine.jackpot.v1.GetJackpotResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetJackpot"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetJackpotRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetJackpotResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JackpotServiceMethodDescriptorSupplier("GetJackpot"))
              .build();
        }
      }
    }
    return getGetJackpotMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetWinnersRequest,
      com.game_engine.jackpot.v1.GetWinnersResponse> getGetWinnersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetWinners",
      requestType = com.game_engine.jackpot.v1.GetWinnersRequest.class,
      responseType = com.game_engine.jackpot.v1.GetWinnersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetWinnersRequest,
      com.game_engine.jackpot.v1.GetWinnersResponse> getGetWinnersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetWinnersRequest, com.game_engine.jackpot.v1.GetWinnersResponse> getGetWinnersMethod;
    if ((getGetWinnersMethod = JackpotServiceGrpc.getGetWinnersMethod) == null) {
      synchronized (JackpotServiceGrpc.class) {
        if ((getGetWinnersMethod = JackpotServiceGrpc.getGetWinnersMethod) == null) {
          JackpotServiceGrpc.getGetWinnersMethod = getGetWinnersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.jackpot.v1.GetWinnersRequest, com.game_engine.jackpot.v1.GetWinnersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetWinners"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetWinnersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetWinnersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JackpotServiceMethodDescriptorSupplier("GetWinners"))
              .build();
        }
      }
    }
    return getGetWinnersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.JoinJackpotRequest,
      com.game_engine.jackpot.v1.JoinJackpotResponse> getJoinJackpotMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinJackpot",
      requestType = com.game_engine.jackpot.v1.JoinJackpotRequest.class,
      responseType = com.game_engine.jackpot.v1.JoinJackpotResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.JoinJackpotRequest,
      com.game_engine.jackpot.v1.JoinJackpotResponse> getJoinJackpotMethod() {
    io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.JoinJackpotRequest, com.game_engine.jackpot.v1.JoinJackpotResponse> getJoinJackpotMethod;
    if ((getJoinJackpotMethod = JackpotServiceGrpc.getJoinJackpotMethod) == null) {
      synchronized (JackpotServiceGrpc.class) {
        if ((getJoinJackpotMethod = JackpotServiceGrpc.getJoinJackpotMethod) == null) {
          JackpotServiceGrpc.getJoinJackpotMethod = getJoinJackpotMethod =
              io.grpc.MethodDescriptor.<com.game_engine.jackpot.v1.JoinJackpotRequest, com.game_engine.jackpot.v1.JoinJackpotResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinJackpot"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.JoinJackpotRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.JoinJackpotResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JackpotServiceMethodDescriptorSupplier("JoinJackpot"))
              .build();
        }
      }
    }
    return getJoinJackpotMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotHistoryRequest,
      com.game_engine.jackpot.v1.GetJackpotHistoryResponse> getGetJackpotHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetJackpotHistory",
      requestType = com.game_engine.jackpot.v1.GetJackpotHistoryRequest.class,
      responseType = com.game_engine.jackpot.v1.GetJackpotHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotHistoryRequest,
      com.game_engine.jackpot.v1.GetJackpotHistoryResponse> getGetJackpotHistoryMethod() {
    io.grpc.MethodDescriptor<com.game_engine.jackpot.v1.GetJackpotHistoryRequest, com.game_engine.jackpot.v1.GetJackpotHistoryResponse> getGetJackpotHistoryMethod;
    if ((getGetJackpotHistoryMethod = JackpotServiceGrpc.getGetJackpotHistoryMethod) == null) {
      synchronized (JackpotServiceGrpc.class) {
        if ((getGetJackpotHistoryMethod = JackpotServiceGrpc.getGetJackpotHistoryMethod) == null) {
          JackpotServiceGrpc.getGetJackpotHistoryMethod = getGetJackpotHistoryMethod =
              io.grpc.MethodDescriptor.<com.game_engine.jackpot.v1.GetJackpotHistoryRequest, com.game_engine.jackpot.v1.GetJackpotHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetJackpotHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetJackpotHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.jackpot.v1.GetJackpotHistoryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new JackpotServiceMethodDescriptorSupplier("GetJackpotHistory"))
              .build();
        }
      }
    }
    return getGetJackpotHistoryMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static JackpotServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JackpotServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JackpotServiceStub>() {
        @java.lang.Override
        public JackpotServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JackpotServiceStub(channel, callOptions);
        }
      };
    return JackpotServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static JackpotServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JackpotServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JackpotServiceBlockingV2Stub>() {
        @java.lang.Override
        public JackpotServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JackpotServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return JackpotServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static JackpotServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JackpotServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JackpotServiceBlockingStub>() {
        @java.lang.Override
        public JackpotServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JackpotServiceBlockingStub(channel, callOptions);
        }
      };
    return JackpotServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static JackpotServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<JackpotServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<JackpotServiceFutureStub>() {
        @java.lang.Override
        public JackpotServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new JackpotServiceFutureStub(channel, callOptions);
        }
      };
    return JackpotServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * List jackpots
     * </pre>
     */
    default void listJackpots(com.game_engine.jackpot.v1.ListJackpotsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.ListJackpotsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListJackpotsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get jackpot details
     * </pre>
     */
    default void getJackpot(com.game_engine.jackpot.v1.GetJackpotRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetJackpotMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get jackpot winners
     * </pre>
     */
    default void getWinners(com.game_engine.jackpot.v1.GetWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetWinnersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetWinnersMethod(), responseObserver);
    }

    /**
     * <pre>
     * Join jackpot
     * </pre>
     */
    default void joinJackpot(com.game_engine.jackpot.v1.JoinJackpotRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.JoinJackpotResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinJackpotMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get jackpot history
     * </pre>
     */
    default void getJackpotHistory(com.game_engine.jackpot.v1.GetJackpotHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetJackpotHistoryMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service JackpotService.
   */
  public static abstract class JackpotServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return JackpotServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service JackpotService.
   */
  public static final class JackpotServiceStub
      extends io.grpc.stub.AbstractAsyncStub<JackpotServiceStub> {
    private JackpotServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JackpotServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JackpotServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * List jackpots
     * </pre>
     */
    public void listJackpots(com.game_engine.jackpot.v1.ListJackpotsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.ListJackpotsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListJackpotsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get jackpot details
     * </pre>
     */
    public void getJackpot(com.game_engine.jackpot.v1.GetJackpotRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetJackpotMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get jackpot winners
     * </pre>
     */
    public void getWinners(com.game_engine.jackpot.v1.GetWinnersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetWinnersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetWinnersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Join jackpot
     * </pre>
     */
    public void joinJackpot(com.game_engine.jackpot.v1.JoinJackpotRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.JoinJackpotResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinJackpotMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get jackpot history
     * </pre>
     */
    public void getJackpotHistory(com.game_engine.jackpot.v1.GetJackpotHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetJackpotHistoryMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service JackpotService.
   */
  public static final class JackpotServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<JackpotServiceBlockingV2Stub> {
    private JackpotServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JackpotServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JackpotServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * List jackpots
     * </pre>
     */
    public com.game_engine.jackpot.v1.ListJackpotsResponse listJackpots(com.game_engine.jackpot.v1.ListJackpotsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListJackpotsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot details
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetJackpotResponse getJackpot(com.game_engine.jackpot.v1.GetJackpotRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetJackpotMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot winners
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetWinnersResponse getWinners(com.game_engine.jackpot.v1.GetWinnersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetWinnersMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Join jackpot
     * </pre>
     */
    public com.game_engine.jackpot.v1.JoinJackpotResponse joinJackpot(com.game_engine.jackpot.v1.JoinJackpotRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getJoinJackpotMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot history
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetJackpotHistoryResponse getJackpotHistory(com.game_engine.jackpot.v1.GetJackpotHistoryRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetJackpotHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service JackpotService.
   */
  public static final class JackpotServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<JackpotServiceBlockingStub> {
    private JackpotServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JackpotServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JackpotServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * List jackpots
     * </pre>
     */
    public com.game_engine.jackpot.v1.ListJackpotsResponse listJackpots(com.game_engine.jackpot.v1.ListJackpotsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListJackpotsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot details
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetJackpotResponse getJackpot(com.game_engine.jackpot.v1.GetJackpotRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetJackpotMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot winners
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetWinnersResponse getWinners(com.game_engine.jackpot.v1.GetWinnersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetWinnersMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Join jackpot
     * </pre>
     */
    public com.game_engine.jackpot.v1.JoinJackpotResponse joinJackpot(com.game_engine.jackpot.v1.JoinJackpotRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinJackpotMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get jackpot history
     * </pre>
     */
    public com.game_engine.jackpot.v1.GetJackpotHistoryResponse getJackpotHistory(com.game_engine.jackpot.v1.GetJackpotHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetJackpotHistoryMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service JackpotService.
   */
  public static final class JackpotServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<JackpotServiceFutureStub> {
    private JackpotServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected JackpotServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new JackpotServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * List jackpots
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.jackpot.v1.ListJackpotsResponse> listJackpots(
        com.game_engine.jackpot.v1.ListJackpotsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListJackpotsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get jackpot details
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.jackpot.v1.GetJackpotResponse> getJackpot(
        com.game_engine.jackpot.v1.GetJackpotRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetJackpotMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get jackpot winners
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.jackpot.v1.GetWinnersResponse> getWinners(
        com.game_engine.jackpot.v1.GetWinnersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetWinnersMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Join jackpot
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.jackpot.v1.JoinJackpotResponse> joinJackpot(
        com.game_engine.jackpot.v1.JoinJackpotRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinJackpotMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get jackpot history
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.jackpot.v1.GetJackpotHistoryResponse> getJackpotHistory(
        com.game_engine.jackpot.v1.GetJackpotHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetJackpotHistoryMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_JACKPOTS = 0;
  private static final int METHODID_GET_JACKPOT = 1;
  private static final int METHODID_GET_WINNERS = 2;
  private static final int METHODID_JOIN_JACKPOT = 3;
  private static final int METHODID_GET_JACKPOT_HISTORY = 4;

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
        case METHODID_LIST_JACKPOTS:
          serviceImpl.listJackpots((com.game_engine.jackpot.v1.ListJackpotsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.ListJackpotsResponse>) responseObserver);
          break;
        case METHODID_GET_JACKPOT:
          serviceImpl.getJackpot((com.game_engine.jackpot.v1.GetJackpotRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotResponse>) responseObserver);
          break;
        case METHODID_GET_WINNERS:
          serviceImpl.getWinners((com.game_engine.jackpot.v1.GetWinnersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetWinnersResponse>) responseObserver);
          break;
        case METHODID_JOIN_JACKPOT:
          serviceImpl.joinJackpot((com.game_engine.jackpot.v1.JoinJackpotRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.JoinJackpotResponse>) responseObserver);
          break;
        case METHODID_GET_JACKPOT_HISTORY:
          serviceImpl.getJackpotHistory((com.game_engine.jackpot.v1.GetJackpotHistoryRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.jackpot.v1.GetJackpotHistoryResponse>) responseObserver);
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
          getListJackpotsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.jackpot.v1.ListJackpotsRequest,
              com.game_engine.jackpot.v1.ListJackpotsResponse>(
                service, METHODID_LIST_JACKPOTS)))
        .addMethod(
          getGetJackpotMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.jackpot.v1.GetJackpotRequest,
              com.game_engine.jackpot.v1.GetJackpotResponse>(
                service, METHODID_GET_JACKPOT)))
        .addMethod(
          getGetWinnersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.jackpot.v1.GetWinnersRequest,
              com.game_engine.jackpot.v1.GetWinnersResponse>(
                service, METHODID_GET_WINNERS)))
        .addMethod(
          getJoinJackpotMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.jackpot.v1.JoinJackpotRequest,
              com.game_engine.jackpot.v1.JoinJackpotResponse>(
                service, METHODID_JOIN_JACKPOT)))
        .addMethod(
          getGetJackpotHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.jackpot.v1.GetJackpotHistoryRequest,
              com.game_engine.jackpot.v1.GetJackpotHistoryResponse>(
                service, METHODID_GET_JACKPOT_HISTORY)))
        .build();
  }

  private static abstract class JackpotServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    JackpotServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.jackpot.v1.JackpotServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("JackpotService");
    }
  }

  private static final class JackpotServiceFileDescriptorSupplier
      extends JackpotServiceBaseDescriptorSupplier {
    JackpotServiceFileDescriptorSupplier() {}
  }

  private static final class JackpotServiceMethodDescriptorSupplier
      extends JackpotServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    JackpotServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (JackpotServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new JackpotServiceFileDescriptorSupplier())
              .addMethod(getListJackpotsMethod())
              .addMethod(getGetJackpotMethod())
              .addMethod(getGetWinnersMethod())
              .addMethod(getJoinJackpotMethod())
              .addMethod(getGetJackpotHistoryMethod())
              .build();
        }
      }
    }
    return result;
  }
}
