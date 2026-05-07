package com.game_engine.payment.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Payment Service - handles deposit, withdrawal, and refund operations
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class PaymentServiceGrpc {

  private PaymentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.payment.v1.PaymentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateDepositRequest,
      com.game_engine.payment.v1.CreateDepositResponse> getCreateDepositMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateDeposit",
      requestType = com.game_engine.payment.v1.CreateDepositRequest.class,
      responseType = com.game_engine.payment.v1.CreateDepositResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateDepositRequest,
      com.game_engine.payment.v1.CreateDepositResponse> getCreateDepositMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateDepositRequest, com.game_engine.payment.v1.CreateDepositResponse> getCreateDepositMethod;
    if ((getCreateDepositMethod = PaymentServiceGrpc.getCreateDepositMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getCreateDepositMethod = PaymentServiceGrpc.getCreateDepositMethod) == null) {
          PaymentServiceGrpc.getCreateDepositMethod = getCreateDepositMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.CreateDepositRequest, com.game_engine.payment.v1.CreateDepositResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateDeposit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CreateDepositRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CreateDepositResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("CreateDeposit"))
              .build();
        }
      }
    }
    return getCreateDepositMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessDepositRequest,
      com.game_engine.payment.v1.ProcessDepositResponse> getProcessDepositMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ProcessDeposit",
      requestType = com.game_engine.payment.v1.ProcessDepositRequest.class,
      responseType = com.game_engine.payment.v1.ProcessDepositResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessDepositRequest,
      com.game_engine.payment.v1.ProcessDepositResponse> getProcessDepositMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessDepositRequest, com.game_engine.payment.v1.ProcessDepositResponse> getProcessDepositMethod;
    if ((getProcessDepositMethod = PaymentServiceGrpc.getProcessDepositMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getProcessDepositMethod = PaymentServiceGrpc.getProcessDepositMethod) == null) {
          PaymentServiceGrpc.getProcessDepositMethod = getProcessDepositMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.ProcessDepositRequest, com.game_engine.payment.v1.ProcessDepositResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ProcessDeposit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.ProcessDepositRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.ProcessDepositResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("ProcessDeposit"))
              .build();
        }
      }
    }
    return getProcessDepositMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateWithdrawalRequest,
      com.game_engine.payment.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateWithdrawal",
      requestType = com.game_engine.payment.v1.CreateWithdrawalRequest.class,
      responseType = com.game_engine.payment.v1.CreateWithdrawalResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateWithdrawalRequest,
      com.game_engine.payment.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.CreateWithdrawalRequest, com.game_engine.payment.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod;
    if ((getCreateWithdrawalMethod = PaymentServiceGrpc.getCreateWithdrawalMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getCreateWithdrawalMethod = PaymentServiceGrpc.getCreateWithdrawalMethod) == null) {
          PaymentServiceGrpc.getCreateWithdrawalMethod = getCreateWithdrawalMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.CreateWithdrawalRequest, com.game_engine.payment.v1.CreateWithdrawalResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateWithdrawal"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CreateWithdrawalRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CreateWithdrawalResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("CreateWithdrawal"))
              .build();
        }
      }
    }
    return getCreateWithdrawalMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessWithdrawalRequest,
      com.game_engine.payment.v1.ProcessWithdrawalResponse> getProcessWithdrawalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ProcessWithdrawal",
      requestType = com.game_engine.payment.v1.ProcessWithdrawalRequest.class,
      responseType = com.game_engine.payment.v1.ProcessWithdrawalResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessWithdrawalRequest,
      com.game_engine.payment.v1.ProcessWithdrawalResponse> getProcessWithdrawalMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.ProcessWithdrawalRequest, com.game_engine.payment.v1.ProcessWithdrawalResponse> getProcessWithdrawalMethod;
    if ((getProcessWithdrawalMethod = PaymentServiceGrpc.getProcessWithdrawalMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getProcessWithdrawalMethod = PaymentServiceGrpc.getProcessWithdrawalMethod) == null) {
          PaymentServiceGrpc.getProcessWithdrawalMethod = getProcessWithdrawalMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.ProcessWithdrawalRequest, com.game_engine.payment.v1.ProcessWithdrawalResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ProcessWithdrawal"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.ProcessWithdrawalRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.ProcessWithdrawalResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("ProcessWithdrawal"))
              .build();
        }
      }
    }
    return getProcessWithdrawalMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.RefundPaymentRequest,
      com.game_engine.payment.v1.RefundPaymentResponse> getRefundPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RefundPayment",
      requestType = com.game_engine.payment.v1.RefundPaymentRequest.class,
      responseType = com.game_engine.payment.v1.RefundPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.RefundPaymentRequest,
      com.game_engine.payment.v1.RefundPaymentResponse> getRefundPaymentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.RefundPaymentRequest, com.game_engine.payment.v1.RefundPaymentResponse> getRefundPaymentMethod;
    if ((getRefundPaymentMethod = PaymentServiceGrpc.getRefundPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getRefundPaymentMethod = PaymentServiceGrpc.getRefundPaymentMethod) == null) {
          PaymentServiceGrpc.getRefundPaymentMethod = getRefundPaymentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.RefundPaymentRequest, com.game_engine.payment.v1.RefundPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RefundPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.RefundPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.RefundPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("RefundPayment"))
              .build();
        }
      }
    }
    return getRefundPaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentRequest,
      com.game_engine.payment.v1.GetPaymentResponse> getGetPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPayment",
      requestType = com.game_engine.payment.v1.GetPaymentRequest.class,
      responseType = com.game_engine.payment.v1.GetPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentRequest,
      com.game_engine.payment.v1.GetPaymentResponse> getGetPaymentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentRequest, com.game_engine.payment.v1.GetPaymentResponse> getGetPaymentMethod;
    if ((getGetPaymentMethod = PaymentServiceGrpc.getGetPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetPaymentMethod = PaymentServiceGrpc.getGetPaymentMethod) == null) {
          PaymentServiceGrpc.getGetPaymentMethod = getGetPaymentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetPaymentRequest, com.game_engine.payment.v1.GetPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetPayment"))
              .build();
        }
      }
    }
    return getGetPaymentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentByExternalIdRequest,
      com.game_engine.payment.v1.GetPaymentByExternalIdResponse> getGetPaymentByExternalIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPaymentByExternalId",
      requestType = com.game_engine.payment.v1.GetPaymentByExternalIdRequest.class,
      responseType = com.game_engine.payment.v1.GetPaymentByExternalIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentByExternalIdRequest,
      com.game_engine.payment.v1.GetPaymentByExternalIdResponse> getGetPaymentByExternalIdMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetPaymentByExternalIdRequest, com.game_engine.payment.v1.GetPaymentByExternalIdResponse> getGetPaymentByExternalIdMethod;
    if ((getGetPaymentByExternalIdMethod = PaymentServiceGrpc.getGetPaymentByExternalIdMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetPaymentByExternalIdMethod = PaymentServiceGrpc.getGetPaymentByExternalIdMethod) == null) {
          PaymentServiceGrpc.getGetPaymentByExternalIdMethod = getGetPaymentByExternalIdMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetPaymentByExternalIdRequest, com.game_engine.payment.v1.GetPaymentByExternalIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPaymentByExternalId"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetPaymentByExternalIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetPaymentByExternalIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetPaymentByExternalId"))
              .build();
        }
      }
    }
    return getGetPaymentByExternalIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentsRequest,
      com.game_engine.payment.v1.GetUserPaymentsResponse> getGetUserPaymentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserPayments",
      requestType = com.game_engine.payment.v1.GetUserPaymentsRequest.class,
      responseType = com.game_engine.payment.v1.GetUserPaymentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentsRequest,
      com.game_engine.payment.v1.GetUserPaymentsResponse> getGetUserPaymentsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentsRequest, com.game_engine.payment.v1.GetUserPaymentsResponse> getGetUserPaymentsMethod;
    if ((getGetUserPaymentsMethod = PaymentServiceGrpc.getGetUserPaymentsMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetUserPaymentsMethod = PaymentServiceGrpc.getGetUserPaymentsMethod) == null) {
          PaymentServiceGrpc.getGetUserPaymentsMethod = getGetUserPaymentsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetUserPaymentsRequest, com.game_engine.payment.v1.GetUserPaymentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserPayments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetUserPaymentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetUserPaymentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetUserPayments"))
              .build();
        }
      }
    }
    return getGetUserPaymentsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentSummaryRequest,
      com.game_engine.payment.v1.GetUserPaymentSummaryResponse> getGetUserPaymentSummaryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserPaymentSummary",
      requestType = com.game_engine.payment.v1.GetUserPaymentSummaryRequest.class,
      responseType = com.game_engine.payment.v1.GetUserPaymentSummaryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentSummaryRequest,
      com.game_engine.payment.v1.GetUserPaymentSummaryResponse> getGetUserPaymentSummaryMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetUserPaymentSummaryRequest, com.game_engine.payment.v1.GetUserPaymentSummaryResponse> getGetUserPaymentSummaryMethod;
    if ((getGetUserPaymentSummaryMethod = PaymentServiceGrpc.getGetUserPaymentSummaryMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetUserPaymentSummaryMethod = PaymentServiceGrpc.getGetUserPaymentSummaryMethod) == null) {
          PaymentServiceGrpc.getGetUserPaymentSummaryMethod = getGetUserPaymentSummaryMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetUserPaymentSummaryRequest, com.game_engine.payment.v1.GetUserPaymentSummaryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserPaymentSummary"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetUserPaymentSummaryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetUserPaymentSummaryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetUserPaymentSummary"))
              .build();
        }
      }
    }
    return getGetUserPaymentSummaryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedMethodsRequest,
      com.game_engine.payment.v1.GetSupportedMethodsResponse> getGetSupportedMethodsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSupportedMethods",
      requestType = com.game_engine.payment.v1.GetSupportedMethodsRequest.class,
      responseType = com.game_engine.payment.v1.GetSupportedMethodsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedMethodsRequest,
      com.game_engine.payment.v1.GetSupportedMethodsResponse> getGetSupportedMethodsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedMethodsRequest, com.game_engine.payment.v1.GetSupportedMethodsResponse> getGetSupportedMethodsMethod;
    if ((getGetSupportedMethodsMethod = PaymentServiceGrpc.getGetSupportedMethodsMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetSupportedMethodsMethod = PaymentServiceGrpc.getGetSupportedMethodsMethod) == null) {
          PaymentServiceGrpc.getGetSupportedMethodsMethod = getGetSupportedMethodsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetSupportedMethodsRequest, com.game_engine.payment.v1.GetSupportedMethodsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSupportedMethods"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetSupportedMethodsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetSupportedMethodsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetSupportedMethods"))
              .build();
        }
      }
    }
    return getGetSupportedMethodsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedCurrenciesRequest,
      com.game_engine.payment.v1.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSupportedCurrencies",
      requestType = com.game_engine.payment.v1.GetSupportedCurrenciesRequest.class,
      responseType = com.game_engine.payment.v1.GetSupportedCurrenciesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedCurrenciesRequest,
      com.game_engine.payment.v1.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.GetSupportedCurrenciesRequest, com.game_engine.payment.v1.GetSupportedCurrenciesResponse> getGetSupportedCurrenciesMethod;
    if ((getGetSupportedCurrenciesMethod = PaymentServiceGrpc.getGetSupportedCurrenciesMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getGetSupportedCurrenciesMethod = PaymentServiceGrpc.getGetSupportedCurrenciesMethod) == null) {
          PaymentServiceGrpc.getGetSupportedCurrenciesMethod = getGetSupportedCurrenciesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.GetSupportedCurrenciesRequest, com.game_engine.payment.v1.GetSupportedCurrenciesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSupportedCurrencies"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetSupportedCurrenciesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.GetSupportedCurrenciesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("GetSupportedCurrencies"))
              .build();
        }
      }
    }
    return getGetSupportedCurrenciesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.payment.v1.CancelPaymentRequest,
      com.game_engine.payment.v1.CancelPaymentResponse> getCancelPaymentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CancelPayment",
      requestType = com.game_engine.payment.v1.CancelPaymentRequest.class,
      responseType = com.game_engine.payment.v1.CancelPaymentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.payment.v1.CancelPaymentRequest,
      com.game_engine.payment.v1.CancelPaymentResponse> getCancelPaymentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.payment.v1.CancelPaymentRequest, com.game_engine.payment.v1.CancelPaymentResponse> getCancelPaymentMethod;
    if ((getCancelPaymentMethod = PaymentServiceGrpc.getCancelPaymentMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getCancelPaymentMethod = PaymentServiceGrpc.getCancelPaymentMethod) == null) {
          PaymentServiceGrpc.getCancelPaymentMethod = getCancelPaymentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.payment.v1.CancelPaymentRequest, com.game_engine.payment.v1.CancelPaymentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CancelPayment"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CancelPaymentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.payment.v1.CancelPaymentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("CancelPayment"))
              .build();
        }
      }
    }
    return getCancelPaymentMethod;
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
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Deposit operations
     * </pre>
     */
    default void createDeposit(com.game_engine.payment.v1.CreateDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateDepositResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateDepositMethod(), responseObserver);
    }

    /**
     */
    default void processDeposit(com.game_engine.payment.v1.ProcessDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessDepositResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getProcessDepositMethod(), responseObserver);
    }

    /**
     * <pre>
     * Withdrawal operations
     * </pre>
     */
    default void createWithdrawal(com.game_engine.payment.v1.CreateWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateWithdrawalResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateWithdrawalMethod(), responseObserver);
    }

    /**
     */
    default void processWithdrawal(com.game_engine.payment.v1.ProcessWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessWithdrawalResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getProcessWithdrawalMethod(), responseObserver);
    }

    /**
     * <pre>
     * Refund operation
     * </pre>
     */
    default void refundPayment(com.game_engine.payment.v1.RefundPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.RefundPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRefundPaymentMethod(), responseObserver);
    }

    /**
     * <pre>
     * Query operations
     * </pre>
     */
    default void getPayment(com.game_engine.payment.v1.GetPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPaymentMethod(), responseObserver);
    }

    /**
     */
    default void getPaymentByExternalId(com.game_engine.payment.v1.GetPaymentByExternalIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentByExternalIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPaymentByExternalIdMethod(), responseObserver);
    }

    /**
     */
    default void getUserPayments(com.game_engine.payment.v1.GetUserPaymentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserPaymentsMethod(), responseObserver);
    }

    /**
     */
    default void getUserPaymentSummary(com.game_engine.payment.v1.GetUserPaymentSummaryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentSummaryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserPaymentSummaryMethod(), responseObserver);
    }

    /**
     * <pre>
     * Supported methods and currencies
     * </pre>
     */
    default void getSupportedMethods(com.game_engine.payment.v1.GetSupportedMethodsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedMethodsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSupportedMethodsMethod(), responseObserver);
    }

    /**
     */
    default void getSupportedCurrencies(com.game_engine.payment.v1.GetSupportedCurrenciesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedCurrenciesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSupportedCurrenciesMethod(), responseObserver);
    }

    /**
     * <pre>
     * Cancel payment
     * </pre>
     */
    default void cancelPayment(com.game_engine.payment.v1.CancelPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CancelPaymentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCancelPaymentMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service PaymentService.
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
   */
  public static abstract class PaymentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return PaymentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service PaymentService.
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
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
     * Deposit operations
     * </pre>
     */
    public void createDeposit(com.game_engine.payment.v1.CreateDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateDepositResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateDepositMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void processDeposit(com.game_engine.payment.v1.ProcessDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessDepositResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getProcessDepositMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Withdrawal operations
     * </pre>
     */
    public void createWithdrawal(com.game_engine.payment.v1.CreateWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateWithdrawalResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateWithdrawalMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void processWithdrawal(com.game_engine.payment.v1.ProcessWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessWithdrawalResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getProcessWithdrawalMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Refund operation
     * </pre>
     */
    public void refundPayment(com.game_engine.payment.v1.RefundPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.RefundPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRefundPaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Query operations
     * </pre>
     */
    public void getPayment(com.game_engine.payment.v1.GetPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPaymentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPaymentByExternalId(com.game_engine.payment.v1.GetPaymentByExternalIdRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentByExternalIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPaymentByExternalIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserPayments(com.game_engine.payment.v1.GetUserPaymentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserPaymentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserPaymentSummary(com.game_engine.payment.v1.GetUserPaymentSummaryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentSummaryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserPaymentSummaryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Supported methods and currencies
     * </pre>
     */
    public void getSupportedMethods(com.game_engine.payment.v1.GetSupportedMethodsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedMethodsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSupportedMethodsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getSupportedCurrencies(com.game_engine.payment.v1.GetSupportedCurrenciesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedCurrenciesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSupportedCurrenciesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Cancel payment
     * </pre>
     */
    public void cancelPayment(com.game_engine.payment.v1.CancelPaymentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CancelPaymentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCancelPaymentMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service PaymentService.
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
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
     * Deposit operations
     * </pre>
     */
    public com.game_engine.payment.v1.CreateDepositResponse createDeposit(com.game_engine.payment.v1.CreateDepositRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateDepositMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.ProcessDepositResponse processDeposit(com.game_engine.payment.v1.ProcessDepositRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getProcessDepositMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Withdrawal operations
     * </pre>
     */
    public com.game_engine.payment.v1.CreateWithdrawalResponse createWithdrawal(com.game_engine.payment.v1.CreateWithdrawalRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.ProcessWithdrawalResponse processWithdrawal(com.game_engine.payment.v1.ProcessWithdrawalRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getProcessWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Refund operation
     * </pre>
     */
    public com.game_engine.payment.v1.RefundPaymentResponse refundPayment(com.game_engine.payment.v1.RefundPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRefundPaymentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Query operations
     * </pre>
     */
    public com.game_engine.payment.v1.GetPaymentResponse getPayment(com.game_engine.payment.v1.GetPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetPaymentByExternalIdResponse getPaymentByExternalId(com.game_engine.payment.v1.GetPaymentByExternalIdRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPaymentByExternalIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetUserPaymentsResponse getUserPayments(com.game_engine.payment.v1.GetUserPaymentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserPaymentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetUserPaymentSummaryResponse getUserPaymentSummary(com.game_engine.payment.v1.GetUserPaymentSummaryRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserPaymentSummaryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Supported methods and currencies
     * </pre>
     */
    public com.game_engine.payment.v1.GetSupportedMethodsResponse getSupportedMethods(com.game_engine.payment.v1.GetSupportedMethodsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSupportedMethodsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetSupportedCurrenciesResponse getSupportedCurrencies(com.game_engine.payment.v1.GetSupportedCurrenciesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetSupportedCurrenciesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Cancel payment
     * </pre>
     */
    public com.game_engine.payment.v1.CancelPaymentResponse cancelPayment(com.game_engine.payment.v1.CancelPaymentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCancelPaymentMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service PaymentService.
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
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
     * Deposit operations
     * </pre>
     */
    public com.game_engine.payment.v1.CreateDepositResponse createDeposit(com.game_engine.payment.v1.CreateDepositRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateDepositMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.ProcessDepositResponse processDeposit(com.game_engine.payment.v1.ProcessDepositRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getProcessDepositMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Withdrawal operations
     * </pre>
     */
    public com.game_engine.payment.v1.CreateWithdrawalResponse createWithdrawal(com.game_engine.payment.v1.CreateWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.ProcessWithdrawalResponse processWithdrawal(com.game_engine.payment.v1.ProcessWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getProcessWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Refund operation
     * </pre>
     */
    public com.game_engine.payment.v1.RefundPaymentResponse refundPayment(com.game_engine.payment.v1.RefundPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRefundPaymentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Query operations
     * </pre>
     */
    public com.game_engine.payment.v1.GetPaymentResponse getPayment(com.game_engine.payment.v1.GetPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPaymentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetPaymentByExternalIdResponse getPaymentByExternalId(com.game_engine.payment.v1.GetPaymentByExternalIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPaymentByExternalIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetUserPaymentsResponse getUserPayments(com.game_engine.payment.v1.GetUserPaymentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserPaymentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetUserPaymentSummaryResponse getUserPaymentSummary(com.game_engine.payment.v1.GetUserPaymentSummaryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserPaymentSummaryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Supported methods and currencies
     * </pre>
     */
    public com.game_engine.payment.v1.GetSupportedMethodsResponse getSupportedMethods(com.game_engine.payment.v1.GetSupportedMethodsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSupportedMethodsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.payment.v1.GetSupportedCurrenciesResponse getSupportedCurrencies(com.game_engine.payment.v1.GetSupportedCurrenciesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSupportedCurrenciesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Cancel payment
     * </pre>
     */
    public com.game_engine.payment.v1.CancelPaymentResponse cancelPayment(com.game_engine.payment.v1.CancelPaymentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCancelPaymentMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service PaymentService.
   * <pre>
   * Payment Service - handles deposit, withdrawal, and refund operations
   * </pre>
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
     * Deposit operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.CreateDepositResponse> createDeposit(
        com.game_engine.payment.v1.CreateDepositRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateDepositMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.ProcessDepositResponse> processDeposit(
        com.game_engine.payment.v1.ProcessDepositRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getProcessDepositMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Withdrawal operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.CreateWithdrawalResponse> createWithdrawal(
        com.game_engine.payment.v1.CreateWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateWithdrawalMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.ProcessWithdrawalResponse> processWithdrawal(
        com.game_engine.payment.v1.ProcessWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getProcessWithdrawalMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Refund operation
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.RefundPaymentResponse> refundPayment(
        com.game_engine.payment.v1.RefundPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRefundPaymentMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Query operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetPaymentResponse> getPayment(
        com.game_engine.payment.v1.GetPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPaymentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetPaymentByExternalIdResponse> getPaymentByExternalId(
        com.game_engine.payment.v1.GetPaymentByExternalIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPaymentByExternalIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetUserPaymentsResponse> getUserPayments(
        com.game_engine.payment.v1.GetUserPaymentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserPaymentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetUserPaymentSummaryResponse> getUserPaymentSummary(
        com.game_engine.payment.v1.GetUserPaymentSummaryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserPaymentSummaryMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Supported methods and currencies
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetSupportedMethodsResponse> getSupportedMethods(
        com.game_engine.payment.v1.GetSupportedMethodsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSupportedMethodsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.GetSupportedCurrenciesResponse> getSupportedCurrencies(
        com.game_engine.payment.v1.GetSupportedCurrenciesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSupportedCurrenciesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Cancel payment
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.payment.v1.CancelPaymentResponse> cancelPayment(
        com.game_engine.payment.v1.CancelPaymentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCancelPaymentMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_DEPOSIT = 0;
  private static final int METHODID_PROCESS_DEPOSIT = 1;
  private static final int METHODID_CREATE_WITHDRAWAL = 2;
  private static final int METHODID_PROCESS_WITHDRAWAL = 3;
  private static final int METHODID_REFUND_PAYMENT = 4;
  private static final int METHODID_GET_PAYMENT = 5;
  private static final int METHODID_GET_PAYMENT_BY_EXTERNAL_ID = 6;
  private static final int METHODID_GET_USER_PAYMENTS = 7;
  private static final int METHODID_GET_USER_PAYMENT_SUMMARY = 8;
  private static final int METHODID_GET_SUPPORTED_METHODS = 9;
  private static final int METHODID_GET_SUPPORTED_CURRENCIES = 10;
  private static final int METHODID_CANCEL_PAYMENT = 11;

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
        case METHODID_CREATE_DEPOSIT:
          serviceImpl.createDeposit((com.game_engine.payment.v1.CreateDepositRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateDepositResponse>) responseObserver);
          break;
        case METHODID_PROCESS_DEPOSIT:
          serviceImpl.processDeposit((com.game_engine.payment.v1.ProcessDepositRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessDepositResponse>) responseObserver);
          break;
        case METHODID_CREATE_WITHDRAWAL:
          serviceImpl.createWithdrawal((com.game_engine.payment.v1.CreateWithdrawalRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CreateWithdrawalResponse>) responseObserver);
          break;
        case METHODID_PROCESS_WITHDRAWAL:
          serviceImpl.processWithdrawal((com.game_engine.payment.v1.ProcessWithdrawalRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.ProcessWithdrawalResponse>) responseObserver);
          break;
        case METHODID_REFUND_PAYMENT:
          serviceImpl.refundPayment((com.game_engine.payment.v1.RefundPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.RefundPaymentResponse>) responseObserver);
          break;
        case METHODID_GET_PAYMENT:
          serviceImpl.getPayment((com.game_engine.payment.v1.GetPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentResponse>) responseObserver);
          break;
        case METHODID_GET_PAYMENT_BY_EXTERNAL_ID:
          serviceImpl.getPaymentByExternalId((com.game_engine.payment.v1.GetPaymentByExternalIdRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetPaymentByExternalIdResponse>) responseObserver);
          break;
        case METHODID_GET_USER_PAYMENTS:
          serviceImpl.getUserPayments((com.game_engine.payment.v1.GetUserPaymentsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentsResponse>) responseObserver);
          break;
        case METHODID_GET_USER_PAYMENT_SUMMARY:
          serviceImpl.getUserPaymentSummary((com.game_engine.payment.v1.GetUserPaymentSummaryRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetUserPaymentSummaryResponse>) responseObserver);
          break;
        case METHODID_GET_SUPPORTED_METHODS:
          serviceImpl.getSupportedMethods((com.game_engine.payment.v1.GetSupportedMethodsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedMethodsResponse>) responseObserver);
          break;
        case METHODID_GET_SUPPORTED_CURRENCIES:
          serviceImpl.getSupportedCurrencies((com.game_engine.payment.v1.GetSupportedCurrenciesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.GetSupportedCurrenciesResponse>) responseObserver);
          break;
        case METHODID_CANCEL_PAYMENT:
          serviceImpl.cancelPayment((com.game_engine.payment.v1.CancelPaymentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.payment.v1.CancelPaymentResponse>) responseObserver);
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
          getCreateDepositMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.CreateDepositRequest,
              com.game_engine.payment.v1.CreateDepositResponse>(
                service, METHODID_CREATE_DEPOSIT)))
        .addMethod(
          getProcessDepositMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.ProcessDepositRequest,
              com.game_engine.payment.v1.ProcessDepositResponse>(
                service, METHODID_PROCESS_DEPOSIT)))
        .addMethod(
          getCreateWithdrawalMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.CreateWithdrawalRequest,
              com.game_engine.payment.v1.CreateWithdrawalResponse>(
                service, METHODID_CREATE_WITHDRAWAL)))
        .addMethod(
          getProcessWithdrawalMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.ProcessWithdrawalRequest,
              com.game_engine.payment.v1.ProcessWithdrawalResponse>(
                service, METHODID_PROCESS_WITHDRAWAL)))
        .addMethod(
          getRefundPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.RefundPaymentRequest,
              com.game_engine.payment.v1.RefundPaymentResponse>(
                service, METHODID_REFUND_PAYMENT)))
        .addMethod(
          getGetPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetPaymentRequest,
              com.game_engine.payment.v1.GetPaymentResponse>(
                service, METHODID_GET_PAYMENT)))
        .addMethod(
          getGetPaymentByExternalIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetPaymentByExternalIdRequest,
              com.game_engine.payment.v1.GetPaymentByExternalIdResponse>(
                service, METHODID_GET_PAYMENT_BY_EXTERNAL_ID)))
        .addMethod(
          getGetUserPaymentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetUserPaymentsRequest,
              com.game_engine.payment.v1.GetUserPaymentsResponse>(
                service, METHODID_GET_USER_PAYMENTS)))
        .addMethod(
          getGetUserPaymentSummaryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetUserPaymentSummaryRequest,
              com.game_engine.payment.v1.GetUserPaymentSummaryResponse>(
                service, METHODID_GET_USER_PAYMENT_SUMMARY)))
        .addMethod(
          getGetSupportedMethodsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetSupportedMethodsRequest,
              com.game_engine.payment.v1.GetSupportedMethodsResponse>(
                service, METHODID_GET_SUPPORTED_METHODS)))
        .addMethod(
          getGetSupportedCurrenciesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.GetSupportedCurrenciesRequest,
              com.game_engine.payment.v1.GetSupportedCurrenciesResponse>(
                service, METHODID_GET_SUPPORTED_CURRENCIES)))
        .addMethod(
          getCancelPaymentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.payment.v1.CancelPaymentRequest,
              com.game_engine.payment.v1.CancelPaymentResponse>(
                service, METHODID_CANCEL_PAYMENT)))
        .build();
  }

  private static abstract class PaymentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PaymentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.payment.v1.PaymentServiceOuterClass.getDescriptor();
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
              .addMethod(getCreateDepositMethod())
              .addMethod(getProcessDepositMethod())
              .addMethod(getCreateWithdrawalMethod())
              .addMethod(getProcessWithdrawalMethod())
              .addMethod(getRefundPaymentMethod())
              .addMethod(getGetPaymentMethod())
              .addMethod(getGetPaymentByExternalIdMethod())
              .addMethod(getGetUserPaymentsMethod())
              .addMethod(getGetUserPaymentSummaryMethod())
              .addMethod(getGetSupportedMethodsMethod())
              .addMethod(getGetSupportedCurrenciesMethod())
              .addMethod(getCancelPaymentMethod())
              .build();
        }
      }
    }
    return result;
  }
}
