package com.game_engine.risk.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class RiskServiceGrpc {

  private RiskServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.risk.v1.RiskService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.risk.v1.GetRiskScoreRequest,
      com.game_engine.risk.v1.GetRiskScoreResponse> getGetRiskScoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetRiskScore",
      requestType = com.game_engine.risk.v1.GetRiskScoreRequest.class,
      responseType = com.game_engine.risk.v1.GetRiskScoreResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.risk.v1.GetRiskScoreRequest,
      com.game_engine.risk.v1.GetRiskScoreResponse> getGetRiskScoreMethod() {
    io.grpc.MethodDescriptor<com.game_engine.risk.v1.GetRiskScoreRequest, com.game_engine.risk.v1.GetRiskScoreResponse> getGetRiskScoreMethod;
    if ((getGetRiskScoreMethod = RiskServiceGrpc.getGetRiskScoreMethod) == null) {
      synchronized (RiskServiceGrpc.class) {
        if ((getGetRiskScoreMethod = RiskServiceGrpc.getGetRiskScoreMethod) == null) {
          RiskServiceGrpc.getGetRiskScoreMethod = getGetRiskScoreMethod =
              io.grpc.MethodDescriptor.<com.game_engine.risk.v1.GetRiskScoreRequest, com.game_engine.risk.v1.GetRiskScoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetRiskScore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.risk.v1.GetRiskScoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.risk.v1.GetRiskScoreResponse.getDefaultInstance()))
              .setSchemaDescriptor(new RiskServiceMethodDescriptorSupplier("GetRiskScore"))
              .build();
        }
      }
    }
    return getGetRiskScoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.risk.v1.UpdateRiskScoreRequest,
      com.game_engine.risk.v1.UpdateRiskScoreResponse> getUpdateRiskScoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateRiskScore",
      requestType = com.game_engine.risk.v1.UpdateRiskScoreRequest.class,
      responseType = com.game_engine.risk.v1.UpdateRiskScoreResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.risk.v1.UpdateRiskScoreRequest,
      com.game_engine.risk.v1.UpdateRiskScoreResponse> getUpdateRiskScoreMethod() {
    io.grpc.MethodDescriptor<com.game_engine.risk.v1.UpdateRiskScoreRequest, com.game_engine.risk.v1.UpdateRiskScoreResponse> getUpdateRiskScoreMethod;
    if ((getUpdateRiskScoreMethod = RiskServiceGrpc.getUpdateRiskScoreMethod) == null) {
      synchronized (RiskServiceGrpc.class) {
        if ((getUpdateRiskScoreMethod = RiskServiceGrpc.getUpdateRiskScoreMethod) == null) {
          RiskServiceGrpc.getUpdateRiskScoreMethod = getUpdateRiskScoreMethod =
              io.grpc.MethodDescriptor.<com.game_engine.risk.v1.UpdateRiskScoreRequest, com.game_engine.risk.v1.UpdateRiskScoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateRiskScore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.risk.v1.UpdateRiskScoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.risk.v1.UpdateRiskScoreResponse.getDefaultInstance()))
              .setSchemaDescriptor(new RiskServiceMethodDescriptorSupplier("UpdateRiskScore"))
              .build();
        }
      }
    }
    return getUpdateRiskScoreMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static RiskServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RiskServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RiskServiceStub>() {
        @java.lang.Override
        public RiskServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RiskServiceStub(channel, callOptions);
        }
      };
    return RiskServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static RiskServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RiskServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RiskServiceBlockingV2Stub>() {
        @java.lang.Override
        public RiskServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RiskServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return RiskServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static RiskServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RiskServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RiskServiceBlockingStub>() {
        @java.lang.Override
        public RiskServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RiskServiceBlockingStub(channel, callOptions);
        }
      };
    return RiskServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static RiskServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<RiskServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<RiskServiceFutureStub>() {
        @java.lang.Override
        public RiskServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new RiskServiceFutureStub(channel, callOptions);
        }
      };
    return RiskServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void getRiskScore(com.game_engine.risk.v1.GetRiskScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.risk.v1.GetRiskScoreResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetRiskScoreMethod(), responseObserver);
    }

    /**
     */
    default void updateRiskScore(com.game_engine.risk.v1.UpdateRiskScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.risk.v1.UpdateRiskScoreResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateRiskScoreMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service RiskService.
   */
  public static abstract class RiskServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return RiskServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service RiskService.
   */
  public static final class RiskServiceStub
      extends io.grpc.stub.AbstractAsyncStub<RiskServiceStub> {
    private RiskServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RiskServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RiskServiceStub(channel, callOptions);
    }

    /**
     */
    public void getRiskScore(com.game_engine.risk.v1.GetRiskScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.risk.v1.GetRiskScoreResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetRiskScoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateRiskScore(com.game_engine.risk.v1.UpdateRiskScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.risk.v1.UpdateRiskScoreResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateRiskScoreMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service RiskService.
   */
  public static final class RiskServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<RiskServiceBlockingV2Stub> {
    private RiskServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RiskServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RiskServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.risk.v1.GetRiskScoreResponse getRiskScore(com.game_engine.risk.v1.GetRiskScoreRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetRiskScoreMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.risk.v1.UpdateRiskScoreResponse updateRiskScore(com.game_engine.risk.v1.UpdateRiskScoreRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateRiskScoreMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service RiskService.
   */
  public static final class RiskServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<RiskServiceBlockingStub> {
    private RiskServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RiskServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RiskServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.risk.v1.GetRiskScoreResponse getRiskScore(com.game_engine.risk.v1.GetRiskScoreRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetRiskScoreMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.risk.v1.UpdateRiskScoreResponse updateRiskScore(com.game_engine.risk.v1.UpdateRiskScoreRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateRiskScoreMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service RiskService.
   */
  public static final class RiskServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<RiskServiceFutureStub> {
    private RiskServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected RiskServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new RiskServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.risk.v1.GetRiskScoreResponse> getRiskScore(
        com.game_engine.risk.v1.GetRiskScoreRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetRiskScoreMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.risk.v1.UpdateRiskScoreResponse> updateRiskScore(
        com.game_engine.risk.v1.UpdateRiskScoreRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateRiskScoreMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_RISK_SCORE = 0;
  private static final int METHODID_UPDATE_RISK_SCORE = 1;

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
        case METHODID_GET_RISK_SCORE:
          serviceImpl.getRiskScore((com.game_engine.risk.v1.GetRiskScoreRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.risk.v1.GetRiskScoreResponse>) responseObserver);
          break;
        case METHODID_UPDATE_RISK_SCORE:
          serviceImpl.updateRiskScore((com.game_engine.risk.v1.UpdateRiskScoreRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.risk.v1.UpdateRiskScoreResponse>) responseObserver);
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
          getGetRiskScoreMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.risk.v1.GetRiskScoreRequest,
              com.game_engine.risk.v1.GetRiskScoreResponse>(
                service, METHODID_GET_RISK_SCORE)))
        .addMethod(
          getUpdateRiskScoreMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.risk.v1.UpdateRiskScoreRequest,
              com.game_engine.risk.v1.UpdateRiskScoreResponse>(
                service, METHODID_UPDATE_RISK_SCORE)))
        .build();
  }

  private static abstract class RiskServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    RiskServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.risk.v1.RiskServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("RiskService");
    }
  }

  private static final class RiskServiceFileDescriptorSupplier
      extends RiskServiceBaseDescriptorSupplier {
    RiskServiceFileDescriptorSupplier() {}
  }

  private static final class RiskServiceMethodDescriptorSupplier
      extends RiskServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    RiskServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (RiskServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new RiskServiceFileDescriptorSupplier())
              .addMethod(getGetRiskScoreMethod())
              .addMethod(getUpdateRiskScoreMethod())
              .build();
        }
      }
    }
    return result;
  }
}
