package com.game_engine.agent.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class AgentServiceGrpc {

  private AgentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.agent.v1.AgentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.agent.v1.ListPlayersRequest,
      com.game_engine.agent.v1.ListPlayersResponse> getListPlayersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPlayers",
      requestType = com.game_engine.agent.v1.ListPlayersRequest.class,
      responseType = com.game_engine.agent.v1.ListPlayersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.agent.v1.ListPlayersRequest,
      com.game_engine.agent.v1.ListPlayersResponse> getListPlayersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.agent.v1.ListPlayersRequest, com.game_engine.agent.v1.ListPlayersResponse> getListPlayersMethod;
    if ((getListPlayersMethod = AgentServiceGrpc.getListPlayersMethod) == null) {
      synchronized (AgentServiceGrpc.class) {
        if ((getListPlayersMethod = AgentServiceGrpc.getListPlayersMethod) == null) {
          AgentServiceGrpc.getListPlayersMethod = getListPlayersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.agent.v1.ListPlayersRequest, com.game_engine.agent.v1.ListPlayersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPlayers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.ListPlayersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.ListPlayersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AgentServiceMethodDescriptorSupplier("ListPlayers"))
              .build();
        }
      }
    }
    return getListPlayersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetPlayerRequest,
      com.game_engine.agent.v1.GetPlayerResponse> getGetPlayerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayer",
      requestType = com.game_engine.agent.v1.GetPlayerRequest.class,
      responseType = com.game_engine.agent.v1.GetPlayerResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetPlayerRequest,
      com.game_engine.agent.v1.GetPlayerResponse> getGetPlayerMethod() {
    io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetPlayerRequest, com.game_engine.agent.v1.GetPlayerResponse> getGetPlayerMethod;
    if ((getGetPlayerMethod = AgentServiceGrpc.getGetPlayerMethod) == null) {
      synchronized (AgentServiceGrpc.class) {
        if ((getGetPlayerMethod = AgentServiceGrpc.getGetPlayerMethod) == null) {
          AgentServiceGrpc.getGetPlayerMethod = getGetPlayerMethod =
              io.grpc.MethodDescriptor.<com.game_engine.agent.v1.GetPlayerRequest, com.game_engine.agent.v1.GetPlayerResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.GetPlayerRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.GetPlayerResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AgentServiceMethodDescriptorSupplier("GetPlayer"))
              .build();
        }
      }
    }
    return getGetPlayerMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.agent.v1.UpdatePlayerLimitRequest,
      com.game_engine.agent.v1.UpdatePlayerLimitResponse> getUpdatePlayerLimitMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePlayerLimit",
      requestType = com.game_engine.agent.v1.UpdatePlayerLimitRequest.class,
      responseType = com.game_engine.agent.v1.UpdatePlayerLimitResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.agent.v1.UpdatePlayerLimitRequest,
      com.game_engine.agent.v1.UpdatePlayerLimitResponse> getUpdatePlayerLimitMethod() {
    io.grpc.MethodDescriptor<com.game_engine.agent.v1.UpdatePlayerLimitRequest, com.game_engine.agent.v1.UpdatePlayerLimitResponse> getUpdatePlayerLimitMethod;
    if ((getUpdatePlayerLimitMethod = AgentServiceGrpc.getUpdatePlayerLimitMethod) == null) {
      synchronized (AgentServiceGrpc.class) {
        if ((getUpdatePlayerLimitMethod = AgentServiceGrpc.getUpdatePlayerLimitMethod) == null) {
          AgentServiceGrpc.getUpdatePlayerLimitMethod = getUpdatePlayerLimitMethod =
              io.grpc.MethodDescriptor.<com.game_engine.agent.v1.UpdatePlayerLimitRequest, com.game_engine.agent.v1.UpdatePlayerLimitResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePlayerLimit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.UpdatePlayerLimitRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.UpdatePlayerLimitResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AgentServiceMethodDescriptorSupplier("UpdatePlayerLimit"))
              .build();
        }
      }
    }
    return getUpdatePlayerLimitMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetDashboardRequest,
      com.game_engine.agent.v1.GetDashboardResponse> getGetDashboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDashboard",
      requestType = com.game_engine.agent.v1.GetDashboardRequest.class,
      responseType = com.game_engine.agent.v1.GetDashboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetDashboardRequest,
      com.game_engine.agent.v1.GetDashboardResponse> getGetDashboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.agent.v1.GetDashboardRequest, com.game_engine.agent.v1.GetDashboardResponse> getGetDashboardMethod;
    if ((getGetDashboardMethod = AgentServiceGrpc.getGetDashboardMethod) == null) {
      synchronized (AgentServiceGrpc.class) {
        if ((getGetDashboardMethod = AgentServiceGrpc.getGetDashboardMethod) == null) {
          AgentServiceGrpc.getGetDashboardMethod = getGetDashboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.agent.v1.GetDashboardRequest, com.game_engine.agent.v1.GetDashboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDashboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.GetDashboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.agent.v1.GetDashboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AgentServiceMethodDescriptorSupplier("GetDashboard"))
              .build();
        }
      }
    }
    return getGetDashboardMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AgentServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AgentServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AgentServiceStub>() {
        @java.lang.Override
        public AgentServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AgentServiceStub(channel, callOptions);
        }
      };
    return AgentServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static AgentServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AgentServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AgentServiceBlockingV2Stub>() {
        @java.lang.Override
        public AgentServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AgentServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return AgentServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AgentServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AgentServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AgentServiceBlockingStub>() {
        @java.lang.Override
        public AgentServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AgentServiceBlockingStub(channel, callOptions);
        }
      };
    return AgentServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AgentServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AgentServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AgentServiceFutureStub>() {
        @java.lang.Override
        public AgentServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AgentServiceFutureStub(channel, callOptions);
        }
      };
    return AgentServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * List downline players
     * </pre>
     */
    default void listPlayers(com.game_engine.agent.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPlayersMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get player details
     * </pre>
     */
    default void getPlayer(com.game_engine.agent.v1.GetPlayerRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetPlayerResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerMethod(), responseObserver);
    }

    /**
     * <pre>
     * Update player limits
     * </pre>
     */
    default void updatePlayerLimit(com.game_engine.agent.v1.UpdatePlayerLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.UpdatePlayerLimitResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePlayerLimitMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get agent dashboard
     * </pre>
     */
    default void getDashboard(com.game_engine.agent.v1.GetDashboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetDashboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDashboardMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AgentService.
   */
  public static abstract class AgentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AgentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AgentService.
   */
  public static final class AgentServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AgentServiceStub> {
    private AgentServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AgentServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AgentServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * List downline players
     * </pre>
     */
    public void listPlayers(com.game_engine.agent.v1.ListPlayersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.ListPlayersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get player details
     * </pre>
     */
    public void getPlayer(com.game_engine.agent.v1.GetPlayerRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetPlayerResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Update player limits
     * </pre>
     */
    public void updatePlayerLimit(com.game_engine.agent.v1.UpdatePlayerLimitRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.UpdatePlayerLimitResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePlayerLimitMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get agent dashboard
     * </pre>
     */
    public void getDashboard(com.game_engine.agent.v1.GetDashboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetDashboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDashboardMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AgentService.
   */
  public static final class AgentServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<AgentServiceBlockingV2Stub> {
    private AgentServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AgentServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AgentServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * List downline players
     * </pre>
     */
    public com.game_engine.agent.v1.ListPlayersResponse listPlayers(com.game_engine.agent.v1.ListPlayersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get player details
     * </pre>
     */
    public com.game_engine.agent.v1.GetPlayerResponse getPlayer(com.game_engine.agent.v1.GetPlayerRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Update player limits
     * </pre>
     */
    public com.game_engine.agent.v1.UpdatePlayerLimitResponse updatePlayerLimit(com.game_engine.agent.v1.UpdatePlayerLimitRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePlayerLimitMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get agent dashboard
     * </pre>
     */
    public com.game_engine.agent.v1.GetDashboardResponse getDashboard(com.game_engine.agent.v1.GetDashboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetDashboardMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service AgentService.
   */
  public static final class AgentServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AgentServiceBlockingStub> {
    private AgentServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AgentServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AgentServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * List downline players
     * </pre>
     */
    public com.game_engine.agent.v1.ListPlayersResponse listPlayers(com.game_engine.agent.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListPlayersMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get player details
     * </pre>
     */
    public com.game_engine.agent.v1.GetPlayerResponse getPlayer(com.game_engine.agent.v1.GetPlayerRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Update player limits
     * </pre>
     */
    public com.game_engine.agent.v1.UpdatePlayerLimitResponse updatePlayerLimit(com.game_engine.agent.v1.UpdatePlayerLimitRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePlayerLimitMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get agent dashboard
     * </pre>
     */
    public com.game_engine.agent.v1.GetDashboardResponse getDashboard(com.game_engine.agent.v1.GetDashboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDashboardMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AgentService.
   */
  public static final class AgentServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AgentServiceFutureStub> {
    private AgentServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AgentServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AgentServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * List downline players
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.agent.v1.ListPlayersResponse> listPlayers(
        com.game_engine.agent.v1.ListPlayersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListPlayersMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get player details
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.agent.v1.GetPlayerResponse> getPlayer(
        com.game_engine.agent.v1.GetPlayerRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Update player limits
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.agent.v1.UpdatePlayerLimitResponse> updatePlayerLimit(
        com.game_engine.agent.v1.UpdatePlayerLimitRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePlayerLimitMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get agent dashboard
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.agent.v1.GetDashboardResponse> getDashboard(
        com.game_engine.agent.v1.GetDashboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDashboardMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_PLAYERS = 0;
  private static final int METHODID_GET_PLAYER = 1;
  private static final int METHODID_UPDATE_PLAYER_LIMIT = 2;
  private static final int METHODID_GET_DASHBOARD = 3;

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
          serviceImpl.listPlayers((com.game_engine.agent.v1.ListPlayersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.agent.v1.ListPlayersResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER:
          serviceImpl.getPlayer((com.game_engine.agent.v1.GetPlayerRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetPlayerResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PLAYER_LIMIT:
          serviceImpl.updatePlayerLimit((com.game_engine.agent.v1.UpdatePlayerLimitRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.agent.v1.UpdatePlayerLimitResponse>) responseObserver);
          break;
        case METHODID_GET_DASHBOARD:
          serviceImpl.getDashboard((com.game_engine.agent.v1.GetDashboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.agent.v1.GetDashboardResponse>) responseObserver);
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
              com.game_engine.agent.v1.ListPlayersRequest,
              com.game_engine.agent.v1.ListPlayersResponse>(
                service, METHODID_LIST_PLAYERS)))
        .addMethod(
          getGetPlayerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.agent.v1.GetPlayerRequest,
              com.game_engine.agent.v1.GetPlayerResponse>(
                service, METHODID_GET_PLAYER)))
        .addMethod(
          getUpdatePlayerLimitMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.agent.v1.UpdatePlayerLimitRequest,
              com.game_engine.agent.v1.UpdatePlayerLimitResponse>(
                service, METHODID_UPDATE_PLAYER_LIMIT)))
        .addMethod(
          getGetDashboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.agent.v1.GetDashboardRequest,
              com.game_engine.agent.v1.GetDashboardResponse>(
                service, METHODID_GET_DASHBOARD)))
        .build();
  }

  private static abstract class AgentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AgentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.agent.v1.AgentServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AgentService");
    }
  }

  private static final class AgentServiceFileDescriptorSupplier
      extends AgentServiceBaseDescriptorSupplier {
    AgentServiceFileDescriptorSupplier() {}
  }

  private static final class AgentServiceMethodDescriptorSupplier
      extends AgentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AgentServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (AgentServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AgentServiceFileDescriptorSupplier())
              .addMethod(getListPlayersMethod())
              .addMethod(getGetPlayerMethod())
              .addMethod(getUpdatePlayerLimitMethod())
              .addMethod(getGetDashboardMethod())
              .build();
        }
      }
    }
    return result;
  }
}
