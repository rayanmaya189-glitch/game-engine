package com.gameengine.payment.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class PaymentServiceGrpc {

  private PaymentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "gameengine.payment.v1.PaymentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.CreatePaymentRequest,
      com.gameengine.payment.v1.CreatePaymentResponse> getCreatePaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreatePayment",
      requestType = com.gameengine.payment.v1.CreatePaymentRequest.class,
      responseType = com.gameengine.payment.v1.CreatePaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.CreatePaymentRequest,
      com.gameengine.payment.v1.CreatePaymentResponse> getCreatePaymentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.CreatePaymentRequest, com.gameengine.payment.v1.CreatePaymentResponse> getCreatePaymentMethod;
    if ((getCreatePaymentMethod = PaymentServiceGrpc.getCreatePaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getCreatePaymentMethod = PaymentServiceGrpc.getCreatePaymentMethod) == null) {
          PaymentServiceGrpc.getCreatePaymentMethod = getCreatePaymentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.CreatePaymentRequest, com.gameengine.payment.v1.CreatePaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreatePayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.CreatePaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.CreatePaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("CreatePayment"))
              .build();
        }
      }
    }
    return getCreatePaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentRequest,
      com.gameengine.payment.v1.GetPaymentResponse> getGetPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPayment",
      requestType = com.gameengine.payment.v1.GetPaymentRequest.class,
      responseType = com.gameengine.payment.v1.GetPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentRequest,
      com.gameengine.payment.v1.GetPaymentResponse> getGetPaymentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentRequest, com.gameengine.payment.v1.GetPaymentResponse> getGetPaymentMethod;
    if ((getGetPaymentMethod = PaymentServiceGrpc.getGetPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetPaymentMethod = PaymentServiceGrpc.getGetPaymentMethod) == null) {
          PaymentServiceGrpc.getGetPaymentMethod = getGetPaymentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.GetPaymentRequest, com.gameengine.payment.v1.GetPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.GetPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.GetPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetPayment"))
              .build();
        }
      }
    }
    return getGetPaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.ApprovePaymentRequest,
      com.gameengine.payment.v1.ApprovePaymentResponse> getApprovePaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ApprovePayment",
      requestType = com.gameengine.payment.v1.ApprovePaymentRequest.class,
      responseType = com.gameengine.payment.v1.ApprovePaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.ApprovePaymentRequest,
      com.gameengine.payment.v1.ApprovePaymentResponse> getApprovePaymentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.ApprovePaymentRequest, com.gameengine.payment.v1.ApprovePaymentResponse> getApprovePaymentMethod;
    if ((getApprovePaymentMethod = PaymentServiceGrpc.getApprovePaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getApprovePaymentMethod = PaymentServiceGrpc.getApprovePaymentMethod) == null) {
          PaymentServiceGrpc.getApprovePaymentMethod = getApprovePaymentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.ApprovePaymentRequest, com.gameengine.payment.v1.ApprovePaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ApprovePayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ApprovePaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ApprovePaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("ApprovePayment"))
              .build();
        }
      }
    }
    return getApprovePaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.RejectPaymentRequest,
      com.gameengine.payment.v1.RejectPaymentResponse> getRejectPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RejectPayment",
      requestType = com.gameengine.payment.v1.RejectPaymentRequest.class,
      responseType = com.gameengine.payment.v1.RejectPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.RejectPaymentRequest,
      com.gameengine.payment.v1.RejectPaymentResponse> getRejectPaymentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.RejectPaymentRequest, com.gameengine.payment.v1.RejectPaymentResponse> getRejectPaymentMethod;
    if ((getRejectPaymentMethod = PaymentServiceGrpc.getRejectPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getRejectPaymentMethod = PaymentServiceGrpc.getRejectPaymentMethod) == null) {
          PaymentServiceGrpc.getRejectPaymentMethod = getRejectPaymentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.RejectPaymentRequest, com.gameengine.payment.v1.RejectPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RejectPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.RejectPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.RejectPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("RejectPayment"))
              .build();
        }
      }
    }
    return getRejectPaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.ProcessPaymentRequest,
      com.gameengine.payment.v1.ProcessPaymentResponse> getProcessPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ProcessPayment",
      requestType = com.gameengine.payment.v1.ProcessPaymentRequest.class,
      responseType = com.gameengine.payment.v1.ProcessPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.ProcessPaymentRequest,
      com.gameengine.payment.v1.ProcessPaymentResponse> getProcessPaymentMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.ProcessPaymentRequest, com.gameengine.payment.v1.ProcessPaymentResponse> getProcessPaymentMethod;
    if ((getProcessPaymentMethod = PaymentServiceGrpc.getProcessPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getProcessPaymentMethod = PaymentServiceGrpc.getProcessPaymentMethod) == null) {
          PaymentServiceGrpc.getProcessPaymentMethod = getProcessPaymentMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.ProcessPaymentRequest, com.gameengine.payment.v1.ProcessPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ProcessPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ProcessPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ProcessPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("ProcessPayment"))
              .build();
        }
      }
    }
    return getProcessPaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.ListPaymentsRequest,
      com.gameengine.payment.v1.ListPaymentsResponse> getListPaymentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListPayments",
      requestType = com.gameengine.payment.v1.ListPaymentsRequest.class,
      responseType = com.gameengine.payment.v1.ListPaymentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.ListPaymentsRequest,
      com.gameengine.payment.v1.ListPaymentsResponse> getListPaymentsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.ListPaymentsRequest, com.gameengine.payment.v1.ListPaymentsResponse> getListPaymentsMethod;
    if ((getListPaymentsMethod = PaymentServiceGrpc.getListPaymentsMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getListPaymentsMethod = PaymentServiceGrpc.getListPaymentsMethod) == null) {
          PaymentServiceGrpc.getListPaymentsMethod = getListPaymentsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.ListPaymentsRequest, com.gameengine.payment.v1.ListPaymentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListPayments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ListPaymentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.ListPaymentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("ListPayments"))
              .build();
        }
      }
    }
    return getListPaymentsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentMethodsRequest,
      com.gameengine.payment.v1.GetPaymentMethodsResponse> getGetPaymentMethodsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPaymentMethods",
      requestType = com.gameengine.payment.v1.GetPaymentMethodsRequest.class,
      responseType = com.gameengine.payment.v1.GetPaymentMethodsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentMethodsRequest,
      com.gameengine.payment.v1.GetPaymentMethodsResponse> getGetPaymentMethodsMethod() {
    io.grpc.MethodDescriptor<com.gameengine.payment.v1.GetPaymentMethodsRequest, com.gameengine.payment.v1.GetPaymentMethodsResponse> getGetPaymentMethodsMethod;
    if ((getGetPaymentMethodsMethod = PaymentServiceGrpc.getGetPaymentMethodsMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetPaymentMethodsMethod = PaymentServiceGrpc.getGetPaymentMethodsMethod) == null) {
          PaymentServiceGrpc.getGetPaymentMethodsMethod = getGetPaymentMethodsMethod =
              io.grpc.MethodDescriptor.<com.gameengine.payment.v1.GetPaymentMethodsRequest, com.gameengine.payment.v1.GetPaymentMethodsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPaymentMethods"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.GetPaymentMethodsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.gameengine.payment.v1.GetPaymentMethodsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetPaymentMethods"))
              .build();
        }
      }
    }
    return getGetPaymentMethodsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PaymentServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceStub>() {
        @java.lang.Override
        public PaymentServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceStub(channel, callOptions);
        }
      };
    return PaymentServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static PaymentServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingV2Stub>() {
        @java.lang.Override
        public PaymentServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return PaymentServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PaymentServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingStub>() {
        @java.lang.Override
        public PaymentServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceBlockingStub(channel, callOptions);
        }
      };
    return PaymentServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PaymentServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceFutureStub>() {
        @java.lang.Override
        public PaymentServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceFutureStub(channel, callOptions);
        }
      };
    return PaymentServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Payment operations
     * </pre>
     */
    default void createPayment(com.gameengine.payment.v1.CreatePaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.CreatePaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreatePaymentMethod(), responseObserver);
    }

    /**
     */
    default void getPayment(com.gameengine.payment.v1.GetPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPaymentMethod(), responseObserver);
    }

    /**
     */
    default void approvePayment(com.gameengine.payment.v1.ApprovePaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ApprovePaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getApprovePaymentMethod(), responseObserver);
    }

    /**
     */
    default void rejectPayment(com.gameengine.payment.v1.RejectPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.RejectPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRejectPaymentMethod(), responseObserver);
    }

    /**
     */
    default void processPayment(com.gameengine.payment.v1.ProcessPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ProcessPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getProcessPaymentMethod(), responseObserver);
    }

    /**
     * <pre>
     * List and methods
     * </pre>
     */
    default void listPayments(com.gameengine.payment.v1.ListPaymentsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ListPaymentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListPaymentsMethod(), responseObserver);
    }

    /**
     */
    default void getPaymentMethods(com.gameengine.payment.v1.GetPaymentMethodsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentMethodsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPaymentMethodsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service PaymentService.
   */
  public static abstract class PaymentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return PaymentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceStub
      extends io.grpc.stub.AbstractAsyncStub<PaymentServiceStub> {
    private PaymentServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Payment operations
     * </pre>
     */
    public void createPayment(com.gameengine.payment.v1.CreatePaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.CreatePaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreatePaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPayment(com.gameengine.payment.v1.GetPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void approvePayment(com.gameengine.payment.v1.ApprovePaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ApprovePaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getApprovePaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void rejectPayment(com.gameengine.payment.v1.RejectPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.RejectPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRejectPaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void processPayment(com.gameengine.payment.v1.ProcessPaymentRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ProcessPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getProcessPaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * List and methods
     * </pre>
     */
    public void listPayments(com.gameengine.payment.v1.ListPaymentsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ListPaymentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListPaymentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPaymentMethods(com.gameengine.payment.v1.GetPaymentMethodsRequest request,
        io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentMethodsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPaymentMethodsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<PaymentServiceBlockingV2Stub> {
    private PaymentServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Payment operations
     * </pre>
     */
    public com.gameengine.payment.v1.CreatePaymentResponse createPayment(com.gameengine.payment.v1.CreatePaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreatePaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.GetPaymentResponse getPayment(com.gameengine.payment.v1.GetPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.ApprovePaymentResponse approvePayment(com.gameengine.payment.v1.ApprovePaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getApprovePaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.RejectPaymentResponse rejectPayment(com.gameengine.payment.v1.RejectPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRejectPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.ProcessPaymentResponse processPayment(com.gameengine.payment.v1.ProcessPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getProcessPaymentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * List and methods
     * </pre>
     */
    public com.gameengine.payment.v1.ListPaymentsResponse listPayments(com.gameengine.payment.v1.ListPaymentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListPaymentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.GetPaymentMethodsResponse getPaymentMethods(com.gameengine.payment.v1.GetPaymentMethodsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPaymentMethodsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<PaymentServiceBlockingStub> {
    private PaymentServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Payment operations
     * </pre>
     */
    public com.gameengine.payment.v1.CreatePaymentResponse createPayment(com.gameengine.payment.v1.CreatePaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreatePaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.GetPaymentResponse getPayment(com.gameengine.payment.v1.GetPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.ApprovePaymentResponse approvePayment(com.gameengine.payment.v1.ApprovePaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getApprovePaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.RejectPaymentResponse rejectPayment(com.gameengine.payment.v1.RejectPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRejectPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.ProcessPaymentResponse processPayment(com.gameengine.payment.v1.ProcessPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getProcessPaymentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * List and methods
     * </pre>
     */
    public com.gameengine.payment.v1.ListPaymentsResponse listPayments(com.gameengine.payment.v1.ListPaymentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListPaymentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.gameengine.payment.v1.GetPaymentMethodsResponse getPaymentMethods(com.gameengine.payment.v1.GetPaymentMethodsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPaymentMethodsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service PaymentService.
   */
  public static final class PaymentServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<PaymentServiceFutureStub> {
    private PaymentServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Payment operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.CreatePaymentResponse> createPayment(
        com.gameengine.payment.v1.CreatePaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreatePaymentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.GetPaymentResponse> getPayment(
        com.gameengine.payment.v1.GetPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPaymentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.ApprovePaymentResponse> approvePayment(
        com.gameengine.payment.v1.ApprovePaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getApprovePaymentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.RejectPaymentResponse> rejectPayment(
        com.gameengine.payment.v1.RejectPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRejectPaymentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.ProcessPaymentResponse> processPayment(
        com.gameengine.payment.v1.ProcessPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getProcessPaymentMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * List and methods
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.ListPaymentsResponse> listPayments(
        com.gameengine.payment.v1.ListPaymentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListPaymentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.gameengine.payment.v1.GetPaymentMethodsResponse> getPaymentMethods(
        com.gameengine.payment.v1.GetPaymentMethodsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPaymentMethodsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_PAYMENT = 0;
  private static final int METHODID_GET_PAYMENT = 1;
  private static final int METHODID_APPROVE_PAYMENT = 2;
  private static final int METHODID_REJECT_PAYMENT = 3;
  private static final int METHODID_PROCESS_PAYMENT = 4;
  private static final int METHODID_LIST_PAYMENTS = 5;
  private static final int METHODID_GET_PAYMENT_METHODS = 6;

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
        case METHODID_CREATE_PAYMENT:
          serviceImpl.createPayment((com.gameengine.payment.v1.CreatePaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.CreatePaymentResponse>) responseObserver);
          break;
        case METHODID_GET_PAYMENT:
          serviceImpl.getPayment((com.gameengine.payment.v1.GetPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentResponse>) responseObserver);
          break;
        case METHODID_APPROVE_PAYMENT:
          serviceImpl.approvePayment((com.gameengine.payment.v1.ApprovePaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ApprovePaymentResponse>) responseObserver);
          break;
        case METHODID_REJECT_PAYMENT:
          serviceImpl.rejectPayment((com.gameengine.payment.v1.RejectPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.RejectPaymentResponse>) responseObserver);
          break;
        case METHODID_PROCESS_PAYMENT:
          serviceImpl.processPayment((com.gameengine.payment.v1.ProcessPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ProcessPaymentResponse>) responseObserver);
          break;
        case METHODID_LIST_PAYMENTS:
          serviceImpl.listPayments((com.gameengine.payment.v1.ListPaymentsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.ListPaymentsResponse>) responseObserver);
          break;
        case METHODID_GET_PAYMENT_METHODS:
          serviceImpl.getPaymentMethods((com.gameengine.payment.v1.GetPaymentMethodsRequest) request,
              (io.grpc.stub.StreamObserver<com.gameengine.payment.v1.GetPaymentMethodsResponse>) responseObserver);
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
          getCreatePaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.CreatePaymentRequest,
              com.gameengine.payment.v1.CreatePaymentResponse>(
                service, METHODID_CREATE_PAYMENT)))
        .addMethod(
          getGetPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.GetPaymentRequest,
              com.gameengine.payment.v1.GetPaymentResponse>(
                service, METHODID_GET_PAYMENT)))
        .addMethod(
          getApprovePaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.ApprovePaymentRequest,
              com.gameengine.payment.v1.ApprovePaymentResponse>(
                service, METHODID_APPROVE_PAYMENT)))
        .addMethod(
          getRejectPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.RejectPaymentRequest,
              com.gameengine.payment.v1.RejectPaymentResponse>(
                service, METHODID_REJECT_PAYMENT)))
        .addMethod(
          getProcessPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.ProcessPaymentRequest,
              com.gameengine.payment.v1.ProcessPaymentResponse>(
                service, METHODID_PROCESS_PAYMENT)))
        .addMethod(
          getListPaymentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.ListPaymentsRequest,
              com.gameengine.payment.v1.ListPaymentsResponse>(
                service, METHODID_LIST_PAYMENTS)))
        .addMethod(
          getGetPaymentMethodsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.gameengine.payment.v1.GetPaymentMethodsRequest,
              com.gameengine.payment.v1.GetPaymentMethodsResponse>(
                service, METHODID_GET_PAYMENT_METHODS)))
        .build();
  }

  private static abstract class PaymentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PaymentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.gameengine.payment.v1.PaymentServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PaymentService");
    }
  }

  private static final class PaymentServiceFileDescriptorSupplier
      extends PaymentServiceBaseDescriptorSupplier {
    PaymentServiceFileDescriptorSupplier() {}
  }

  private static final class PaymentServiceMethodDescriptorSupplier
      extends PaymentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    PaymentServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (PaymentServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PaymentServiceFileDescriptorSupplier())
              .addMethod(getCreatePaymentMethod())
              .addMethod(getGetPaymentMethod())
              .addMethod(getApprovePaymentMethod())
              .addMethod(getRejectPaymentMethod())
              .addMethod(getProcessPaymentMethod())
              .addMethod(getListPaymentsMethod())
              .addMethod(getGetPaymentMethodsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
