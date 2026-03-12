package com.game_engine.affiliate.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class AffiliateServiceGrpc {

  private AffiliateServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.affiliate.v1.AffiliateService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest,
      com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "TrackClick",
      requestType = com.game_engine.affiliate.v1.TrackClickRequest.class,
      responseType = com.game_engine.affiliate.v1.TrackClickResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest,
      com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.TrackClickRequest, com.game_engine.affiliate.v1.TrackClickResponse> getTrackClickMethod;
    if ((getTrackClickMethod = AffiliateServiceGrpc.getTrackClickMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getTrackClickMethod = AffiliateServiceGrpc.getTrackClickMethod) == null) {
          AffiliateServiceGrpc.getTrackClickMethod = getTrackClickMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.TrackClickRequest, com.game_engine.affiliate.v1.TrackClickResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "TrackClick"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackClickRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.TrackClickResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("TrackClick"))
              .build();
        }
      }
    }
    return getTrackClickMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest,
      com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPerformanceReport",
      requestType = com.game_engine.affiliate.v1.GetPerformanceReportRequest.class,
      responseType = com.game_engine.affiliate.v1.GetPerformanceReportResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest,
      com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetPerformanceReportRequest, com.game_engine.affiliate.v1.GetPerformanceReportResponse> getGetPerformanceReportMethod;
    if ((getGetPerformanceReportMethod = AffiliateServiceGrpc.getGetPerformanceReportMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetPerformanceReportMethod = AffiliateServiceGrpc.getGetPerformanceReportMethod) == null) {
          AffiliateServiceGrpc.getGetPerformanceReportMethod = getGetPerformanceReportMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetPerformanceReportRequest, com.game_engine.affiliate.v1.GetPerformanceReportResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPerformanceReport"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetPerformanceReportRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetPerformanceReportResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetPerformanceReport"))
              .build();
        }
      }
    }
    return getGetPerformanceReportMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest,
      com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetClickReports",
      requestType = com.game_engine.affiliate.v1.GetClickReportsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetClickReportsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest,
      com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetClickReportsRequest, com.game_engine.affiliate.v1.GetClickReportsResponse> getGetClickReportsMethod;
    if ((getGetClickReportsMethod = AffiliateServiceGrpc.getGetClickReportsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetClickReportsMethod = AffiliateServiceGrpc.getGetClickReportsMethod) == null) {
          AffiliateServiceGrpc.getGetClickReportsMethod = getGetClickReportsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetClickReportsRequest, com.game_engine.affiliate.v1.GetClickReportsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetClickReports"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetClickReportsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetClickReportsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetClickReports"))
              .build();
        }
      }
    }
    return getGetClickReportsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest,
      com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetConversionReports",
      requestType = com.game_engine.affiliate.v1.GetConversionReportsRequest.class,
      responseType = com.game_engine.affiliate.v1.GetConversionReportsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest,
      com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetConversionReportsRequest, com.game_engine.affiliate.v1.GetConversionReportsResponse> getGetConversionReportsMethod;
    if ((getGetConversionReportsMethod = AffiliateServiceGrpc.getGetConversionReportsMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetConversionReportsMethod = AffiliateServiceGrpc.getGetConversionReportsMethod) == null) {
          AffiliateServiceGrpc.getGetConversionReportsMethod = getGetConversionReportsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetConversionReportsRequest, com.game_engine.affiliate.v1.GetConversionReportsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConversionReports"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetConversionReportsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetConversionReportsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetConversionReports"))
              .build();
        }
      }
    }
    return getGetConversionReportsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
      com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAffiliateLinks",
      requestType = com.game_engine.affiliate.v1.GetAffiliateLinksRequest.class,
      responseType = com.game_engine.affiliate.v1.GetAffiliateLinksResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
      com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.GetAffiliateLinksRequest, com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getGetAffiliateLinksMethod;
    if ((getGetAffiliateLinksMethod = AffiliateServiceGrpc.getGetAffiliateLinksMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getGetAffiliateLinksMethod = AffiliateServiceGrpc.getGetAffiliateLinksMethod) == null) {
          AffiliateServiceGrpc.getGetAffiliateLinksMethod = getGetAffiliateLinksMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.GetAffiliateLinksRequest, com.game_engine.affiliate.v1.GetAffiliateLinksResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAffiliateLinks"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateLinksRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.GetAffiliateLinksResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("GetAffiliateLinks"))
              .build();
        }
      }
    }
    return getGetAffiliateLinksMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
      com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateAffiliateLink",
      requestType = com.game_engine.affiliate.v1.CreateAffiliateLinkRequest.class,
      responseType = com.game_engine.affiliate.v1.CreateAffiliateLinkResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
      com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod() {
    io.grpc.MethodDescriptor<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest, com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> getCreateAffiliateLinkMethod;
    if ((getCreateAffiliateLinkMethod = AffiliateServiceGrpc.getCreateAffiliateLinkMethod) == null) {
      synchronized (AffiliateServiceGrpc.class) {
        if ((getCreateAffiliateLinkMethod = AffiliateServiceGrpc.getCreateAffiliateLinkMethod) == null) {
          AffiliateServiceGrpc.getCreateAffiliateLinkMethod = getCreateAffiliateLinkMethod =
              io.grpc.MethodDescriptor.<com.game_engine.affiliate.v1.CreateAffiliateLinkRequest, com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateAffiliateLink"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CreateAffiliateLinkRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.affiliate.v1.CreateAffiliateLinkResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AffiliateServiceMethodDescriptorSupplier("CreateAffiliateLink"))
              .build();
        }
      }
    }
    return getCreateAffiliateLinkMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AffiliateServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceStub>() {
        @java.lang.Override
        public AffiliateServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceStub(channel, callOptions);
        }
      };
    return AffiliateServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static AffiliateServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingV2Stub>() {
        @java.lang.Override
        public AffiliateServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return AffiliateServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AffiliateServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceBlockingStub>() {
        @java.lang.Override
        public AffiliateServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceBlockingStub(channel, callOptions);
        }
      };
    return AffiliateServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AffiliateServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AffiliateServiceFutureStub>() {
        @java.lang.Override
        public AffiliateServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AffiliateServiceFutureStub(channel, callOptions);
        }
      };
    return AffiliateServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Track affiliate click
     * </pre>
     */
    default void trackClick(com.game_engine.affiliate.v1.TrackClickRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getTrackClickMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get performance report
     * </pre>
     */
    default void getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPerformanceReportMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get click reports
     * </pre>
     */
    default void getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetClickReportsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get conversion reports
     * </pre>
     */
    default void getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConversionReportsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get affiliate links
     * </pre>
     */
    default void getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAffiliateLinksMethod(), responseObserver);
    }

    /**
     * <pre>
     * Create affiliate link
     * </pre>
     */
    default void createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateAffiliateLinkMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AffiliateService.
   */
  public static abstract class AffiliateServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AffiliateServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AffiliateService.
   */
  public static final class AffiliateServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AffiliateServiceStub> {
    private AffiliateServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Track affiliate click
     * </pre>
     */
    public void trackClick(com.game_engine.affiliate.v1.TrackClickRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getTrackClickMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get performance report
     * </pre>
     */
    public void getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPerformanceReportMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get click reports
     * </pre>
     */
    public void getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetClickReportsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get conversion reports
     * </pre>
     */
    public void getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConversionReportsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get affiliate links
     * </pre>
     */
    public void getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAffiliateLinksMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Create affiliate link
     * </pre>
     */
    public void createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateAffiliateLinkMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AffiliateService.
   */
  public static final class AffiliateServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<AffiliateServiceBlockingV2Stub> {
    private AffiliateServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Track affiliate click
     * </pre>
     */
    public com.game_engine.affiliate.v1.TrackClickResponse trackClick(com.game_engine.affiliate.v1.TrackClickRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getTrackClickMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get performance report
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetPerformanceReportResponse getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPerformanceReportMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get click reports
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetClickReportsResponse getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetClickReportsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get conversion reports
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetConversionReportsResponse getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetConversionReportsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get affiliate links
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetAffiliateLinksResponse getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAffiliateLinksMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Create affiliate link
     * </pre>
     */
    public com.game_engine.affiliate.v1.CreateAffiliateLinkResponse createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateAffiliateLinkMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service AffiliateService.
   */
  public static final class AffiliateServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AffiliateServiceBlockingStub> {
    private AffiliateServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Track affiliate click
     * </pre>
     */
    public com.game_engine.affiliate.v1.TrackClickResponse trackClick(com.game_engine.affiliate.v1.TrackClickRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getTrackClickMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get performance report
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetPerformanceReportResponse getPerformanceReport(com.game_engine.affiliate.v1.GetPerformanceReportRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPerformanceReportMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get click reports
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetClickReportsResponse getClickReports(com.game_engine.affiliate.v1.GetClickReportsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetClickReportsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get conversion reports
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetConversionReportsResponse getConversionReports(com.game_engine.affiliate.v1.GetConversionReportsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConversionReportsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get affiliate links
     * </pre>
     */
    public com.game_engine.affiliate.v1.GetAffiliateLinksResponse getAffiliateLinks(com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAffiliateLinksMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Create affiliate link
     * </pre>
     */
    public com.game_engine.affiliate.v1.CreateAffiliateLinkResponse createAffiliateLink(com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateAffiliateLinkMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AffiliateService.
   */
  public static final class AffiliateServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AffiliateServiceFutureStub> {
    private AffiliateServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AffiliateServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AffiliateServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Track affiliate click
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.TrackClickResponse> trackClick(
        com.game_engine.affiliate.v1.TrackClickRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getTrackClickMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get performance report
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetPerformanceReportResponse> getPerformanceReport(
        com.game_engine.affiliate.v1.GetPerformanceReportRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPerformanceReportMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get click reports
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetClickReportsResponse> getClickReports(
        com.game_engine.affiliate.v1.GetClickReportsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetClickReportsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get conversion reports
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetConversionReportsResponse> getConversionReports(
        com.game_engine.affiliate.v1.GetConversionReportsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConversionReportsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get affiliate links
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.GetAffiliateLinksResponse> getAffiliateLinks(
        com.game_engine.affiliate.v1.GetAffiliateLinksRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAffiliateLinksMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Create affiliate link
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse> createAffiliateLink(
        com.game_engine.affiliate.v1.CreateAffiliateLinkRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateAffiliateLinkMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_TRACK_CLICK = 0;
  private static final int METHODID_GET_PERFORMANCE_REPORT = 1;
  private static final int METHODID_GET_CLICK_REPORTS = 2;
  private static final int METHODID_GET_CONVERSION_REPORTS = 3;
  private static final int METHODID_GET_AFFILIATE_LINKS = 4;
  private static final int METHODID_CREATE_AFFILIATE_LINK = 5;

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
        case METHODID_TRACK_CLICK:
          serviceImpl.trackClick((com.game_engine.affiliate.v1.TrackClickRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.TrackClickResponse>) responseObserver);
          break;
        case METHODID_GET_PERFORMANCE_REPORT:
          serviceImpl.getPerformanceReport((com.game_engine.affiliate.v1.GetPerformanceReportRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetPerformanceReportResponse>) responseObserver);
          break;
        case METHODID_GET_CLICK_REPORTS:
          serviceImpl.getClickReports((com.game_engine.affiliate.v1.GetClickReportsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetClickReportsResponse>) responseObserver);
          break;
        case METHODID_GET_CONVERSION_REPORTS:
          serviceImpl.getConversionReports((com.game_engine.affiliate.v1.GetConversionReportsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetConversionReportsResponse>) responseObserver);
          break;
        case METHODID_GET_AFFILIATE_LINKS:
          serviceImpl.getAffiliateLinks((com.game_engine.affiliate.v1.GetAffiliateLinksRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.GetAffiliateLinksResponse>) responseObserver);
          break;
        case METHODID_CREATE_AFFILIATE_LINK:
          serviceImpl.createAffiliateLink((com.game_engine.affiliate.v1.CreateAffiliateLinkRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>) responseObserver);
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
          getTrackClickMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.TrackClickRequest,
              com.game_engine.affiliate.v1.TrackClickResponse>(
                service, METHODID_TRACK_CLICK)))
        .addMethod(
          getGetPerformanceReportMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetPerformanceReportRequest,
              com.game_engine.affiliate.v1.GetPerformanceReportResponse>(
                service, METHODID_GET_PERFORMANCE_REPORT)))
        .addMethod(
          getGetClickReportsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetClickReportsRequest,
              com.game_engine.affiliate.v1.GetClickReportsResponse>(
                service, METHODID_GET_CLICK_REPORTS)))
        .addMethod(
          getGetConversionReportsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetConversionReportsRequest,
              com.game_engine.affiliate.v1.GetConversionReportsResponse>(
                service, METHODID_GET_CONVERSION_REPORTS)))
        .addMethod(
          getGetAffiliateLinksMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.GetAffiliateLinksRequest,
              com.game_engine.affiliate.v1.GetAffiliateLinksResponse>(
                service, METHODID_GET_AFFILIATE_LINKS)))
        .addMethod(
          getCreateAffiliateLinkMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.affiliate.v1.CreateAffiliateLinkRequest,
              com.game_engine.affiliate.v1.CreateAffiliateLinkResponse>(
                service, METHODID_CREATE_AFFILIATE_LINK)))
        .build();
  }

  private static abstract class AffiliateServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AffiliateServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.affiliate.v1.AffiliateServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AffiliateService");
    }
  }

  private static final class AffiliateServiceFileDescriptorSupplier
      extends AffiliateServiceBaseDescriptorSupplier {
    AffiliateServiceFileDescriptorSupplier() {}
  }

  private static final class AffiliateServiceMethodDescriptorSupplier
      extends AffiliateServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AffiliateServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (AffiliateServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AffiliateServiceFileDescriptorSupplier())
              .addMethod(getTrackClickMethod())
              .addMethod(getGetPerformanceReportMethod())
              .addMethod(getGetClickReportsMethod())
              .addMethod(getGetConversionReportsMethod())
              .addMethod(getGetAffiliateLinksMethod())
              .addMethod(getCreateAffiliateLinkMethod())
              .build();
        }
      }
    }
    return result;
  }
}
