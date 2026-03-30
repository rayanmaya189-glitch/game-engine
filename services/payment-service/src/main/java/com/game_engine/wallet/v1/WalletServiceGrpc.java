package com.game_engine.wallet.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Wallet Service - handles player balances, deposits, withdrawals, and betting
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.60.0)",
    comments = "Source: game_engine/wallet/v1/wallet_service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class WalletServiceGrpc {

  private WalletServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.wallet.v1.WalletService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetBalanceRequest,
      com.game_engine.wallet.v1.GetBalanceResponse> getGetBalanceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBalance",
      requestType = com.game_engine.wallet.v1.GetBalanceRequest.class,
      responseType = com.game_engine.wallet.v1.GetBalanceResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetBalanceRequest,
      com.game_engine.wallet.v1.GetBalanceResponse> getGetBalanceMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetBalanceRequest, com.game_engine.wallet.v1.GetBalanceResponse> getGetBalanceMethod;
    if ((getGetBalanceMethod = WalletServiceGrpc.getGetBalanceMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getGetBalanceMethod = WalletServiceGrpc.getGetBalanceMethod) == null) {
          WalletServiceGrpc.getGetBalanceMethod = getGetBalanceMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.GetBalanceRequest, com.game_engine.wallet.v1.GetBalanceResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBalance"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetBalanceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetBalanceResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("GetBalance"))
              .build();
        }
      }
    }
    return getGetBalanceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetAllBalancesRequest,
      com.game_engine.wallet.v1.GetAllBalancesResponse> getGetAllBalancesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAllBalances",
      requestType = com.game_engine.wallet.v1.GetAllBalancesRequest.class,
      responseType = com.game_engine.wallet.v1.GetAllBalancesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetAllBalancesRequest,
      com.game_engine.wallet.v1.GetAllBalancesResponse> getGetAllBalancesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetAllBalancesRequest, com.game_engine.wallet.v1.GetAllBalancesResponse> getGetAllBalancesMethod;
    if ((getGetAllBalancesMethod = WalletServiceGrpc.getGetAllBalancesMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getGetAllBalancesMethod = WalletServiceGrpc.getGetAllBalancesMethod) == null) {
          WalletServiceGrpc.getGetAllBalancesMethod = getGetAllBalancesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.GetAllBalancesRequest, com.game_engine.wallet.v1.GetAllBalancesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllBalances"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetAllBalancesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetAllBalancesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("GetAllBalances"))
              .build();
        }
      }
    }
    return getGetAllBalancesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetTransactionHistoryRequest,
      com.game_engine.wallet.v1.GetTransactionHistoryResponse> getGetTransactionHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTransactionHistory",
      requestType = com.game_engine.wallet.v1.GetTransactionHistoryRequest.class,
      responseType = com.game_engine.wallet.v1.GetTransactionHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetTransactionHistoryRequest,
      com.game_engine.wallet.v1.GetTransactionHistoryResponse> getGetTransactionHistoryMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetTransactionHistoryRequest, com.game_engine.wallet.v1.GetTransactionHistoryResponse> getGetTransactionHistoryMethod;
    if ((getGetTransactionHistoryMethod = WalletServiceGrpc.getGetTransactionHistoryMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getGetTransactionHistoryMethod = WalletServiceGrpc.getGetTransactionHistoryMethod) == null) {
          WalletServiceGrpc.getGetTransactionHistoryMethod = getGetTransactionHistoryMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.GetTransactionHistoryRequest, com.game_engine.wallet.v1.GetTransactionHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTransactionHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetTransactionHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetTransactionHistoryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("GetTransactionHistory"))
              .build();
        }
      }
    }
    return getGetTransactionHistoryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateDepositRequest,
      com.game_engine.wallet.v1.CreateDepositResponse> getCreateDepositMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateDeposit",
      requestType = com.game_engine.wallet.v1.CreateDepositRequest.class,
      responseType = com.game_engine.wallet.v1.CreateDepositResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateDepositRequest,
      com.game_engine.wallet.v1.CreateDepositResponse> getCreateDepositMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateDepositRequest, com.game_engine.wallet.v1.CreateDepositResponse> getCreateDepositMethod;
    if ((getCreateDepositMethod = WalletServiceGrpc.getCreateDepositMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getCreateDepositMethod = WalletServiceGrpc.getCreateDepositMethod) == null) {
          WalletServiceGrpc.getCreateDepositMethod = getCreateDepositMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.CreateDepositRequest, com.game_engine.wallet.v1.CreateDepositResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateDeposit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateDepositRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateDepositResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("CreateDeposit"))
              .build();
        }
      }
    }
    return getCreateDepositMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmDepositRequest,
      com.game_engine.wallet.v1.ConfirmDepositResponse> getConfirmDepositMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ConfirmDeposit",
      requestType = com.game_engine.wallet.v1.ConfirmDepositRequest.class,
      responseType = com.game_engine.wallet.v1.ConfirmDepositResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmDepositRequest,
      com.game_engine.wallet.v1.ConfirmDepositResponse> getConfirmDepositMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmDepositRequest, com.game_engine.wallet.v1.ConfirmDepositResponse> getConfirmDepositMethod;
    if ((getConfirmDepositMethod = WalletServiceGrpc.getConfirmDepositMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getConfirmDepositMethod = WalletServiceGrpc.getConfirmDepositMethod) == null) {
          WalletServiceGrpc.getConfirmDepositMethod = getConfirmDepositMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.ConfirmDepositRequest, com.game_engine.wallet.v1.ConfirmDepositResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ConfirmDeposit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ConfirmDepositRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ConfirmDepositResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("ConfirmDeposit"))
              .build();
        }
      }
    }
    return getConfirmDepositMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateWithdrawalRequest,
      com.game_engine.wallet.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateWithdrawal",
      requestType = com.game_engine.wallet.v1.CreateWithdrawalRequest.class,
      responseType = com.game_engine.wallet.v1.CreateWithdrawalResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateWithdrawalRequest,
      com.game_engine.wallet.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateWithdrawalRequest, com.game_engine.wallet.v1.CreateWithdrawalResponse> getCreateWithdrawalMethod;
    if ((getCreateWithdrawalMethod = WalletServiceGrpc.getCreateWithdrawalMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getCreateWithdrawalMethod = WalletServiceGrpc.getCreateWithdrawalMethod) == null) {
          WalletServiceGrpc.getCreateWithdrawalMethod = getCreateWithdrawalMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.CreateWithdrawalRequest, com.game_engine.wallet.v1.CreateWithdrawalResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateWithdrawal"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateWithdrawalRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateWithdrawalResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("CreateWithdrawal"))
              .build();
        }
      }
    }
    return getCreateWithdrawalMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmWithdrawalRequest,
      com.game_engine.wallet.v1.ConfirmWithdrawalResponse> getConfirmWithdrawalMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ConfirmWithdrawal",
      requestType = com.game_engine.wallet.v1.ConfirmWithdrawalRequest.class,
      responseType = com.game_engine.wallet.v1.ConfirmWithdrawalResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmWithdrawalRequest,
      com.game_engine.wallet.v1.ConfirmWithdrawalResponse> getConfirmWithdrawalMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ConfirmWithdrawalRequest, com.game_engine.wallet.v1.ConfirmWithdrawalResponse> getConfirmWithdrawalMethod;
    if ((getConfirmWithdrawalMethod = WalletServiceGrpc.getConfirmWithdrawalMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getConfirmWithdrawalMethod = WalletServiceGrpc.getConfirmWithdrawalMethod) == null) {
          WalletServiceGrpc.getConfirmWithdrawalMethod = getConfirmWithdrawalMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.ConfirmWithdrawalRequest, com.game_engine.wallet.v1.ConfirmWithdrawalResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ConfirmWithdrawal"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ConfirmWithdrawalRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ConfirmWithdrawalResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("ConfirmWithdrawal"))
              .build();
        }
      }
    }
    return getConfirmWithdrawalMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.PlaceBetRequest,
      com.game_engine.wallet.v1.PlaceBetResponse> getPlaceBetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PlaceBet",
      requestType = com.game_engine.wallet.v1.PlaceBetRequest.class,
      responseType = com.game_engine.wallet.v1.PlaceBetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.PlaceBetRequest,
      com.game_engine.wallet.v1.PlaceBetResponse> getPlaceBetMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.PlaceBetRequest, com.game_engine.wallet.v1.PlaceBetResponse> getPlaceBetMethod;
    if ((getPlaceBetMethod = WalletServiceGrpc.getPlaceBetMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getPlaceBetMethod = WalletServiceGrpc.getPlaceBetMethod) == null) {
          WalletServiceGrpc.getPlaceBetMethod = getPlaceBetMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.PlaceBetRequest, com.game_engine.wallet.v1.PlaceBetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PlaceBet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.PlaceBetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.PlaceBetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("PlaceBet"))
              .build();
        }
      }
    }
    return getPlaceBetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.SettleBetRequest,
      com.game_engine.wallet.v1.SettleBetResponse> getSettleBetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SettleBet",
      requestType = com.game_engine.wallet.v1.SettleBetRequest.class,
      responseType = com.game_engine.wallet.v1.SettleBetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.SettleBetRequest,
      com.game_engine.wallet.v1.SettleBetResponse> getSettleBetMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.SettleBetRequest, com.game_engine.wallet.v1.SettleBetResponse> getSettleBetMethod;
    if ((getSettleBetMethod = WalletServiceGrpc.getSettleBetMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getSettleBetMethod = WalletServiceGrpc.getSettleBetMethod) == null) {
          WalletServiceGrpc.getSettleBetMethod = getSettleBetMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.SettleBetRequest, com.game_engine.wallet.v1.SettleBetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SettleBet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.SettleBetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.SettleBetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("SettleBet"))
              .build();
        }
      }
    }
    return getSettleBetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CancelBetRequest,
      com.game_engine.wallet.v1.CancelBetResponse> getCancelBetMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CancelBet",
      requestType = com.game_engine.wallet.v1.CancelBetRequest.class,
      responseType = com.game_engine.wallet.v1.CancelBetResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CancelBetRequest,
      com.game_engine.wallet.v1.CancelBetResponse> getCancelBetMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CancelBetRequest, com.game_engine.wallet.v1.CancelBetResponse> getCancelBetMethod;
    if ((getCancelBetMethod = WalletServiceGrpc.getCancelBetMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getCancelBetMethod = WalletServiceGrpc.getCancelBetMethod) == null) {
          WalletServiceGrpc.getCancelBetMethod = getCancelBetMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.CancelBetRequest, com.game_engine.wallet.v1.CancelBetResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CancelBet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CancelBetRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CancelBetResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("CancelBet"))
              .build();
        }
      }
    }
    return getCancelBetMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateBonusCreditRequest,
      com.game_engine.wallet.v1.CreateBonusCreditResponse> getCreateBonusCreditMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateBonusCredit",
      requestType = com.game_engine.wallet.v1.CreateBonusCreditRequest.class,
      responseType = com.game_engine.wallet.v1.CreateBonusCreditResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateBonusCreditRequest,
      com.game_engine.wallet.v1.CreateBonusCreditResponse> getCreateBonusCreditMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.CreateBonusCreditRequest, com.game_engine.wallet.v1.CreateBonusCreditResponse> getCreateBonusCreditMethod;
    if ((getCreateBonusCreditMethod = WalletServiceGrpc.getCreateBonusCreditMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getCreateBonusCreditMethod = WalletServiceGrpc.getCreateBonusCreditMethod) == null) {
          WalletServiceGrpc.getCreateBonusCreditMethod = getCreateBonusCreditMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.CreateBonusCreditRequest, com.game_engine.wallet.v1.CreateBonusCreditResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateBonusCredit"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateBonusCreditRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.CreateBonusCreditResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("CreateBonusCredit"))
              .build();
        }
      }
    }
    return getCreateBonusCreditMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ReverseTransactionRequest,
      com.game_engine.wallet.v1.ReverseTransactionResponse> getReverseTransactionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ReverseTransaction",
      requestType = com.game_engine.wallet.v1.ReverseTransactionRequest.class,
      responseType = com.game_engine.wallet.v1.ReverseTransactionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ReverseTransactionRequest,
      com.game_engine.wallet.v1.ReverseTransactionResponse> getReverseTransactionMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.ReverseTransactionRequest, com.game_engine.wallet.v1.ReverseTransactionResponse> getReverseTransactionMethod;
    if ((getReverseTransactionMethod = WalletServiceGrpc.getReverseTransactionMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getReverseTransactionMethod = WalletServiceGrpc.getReverseTransactionMethod) == null) {
          WalletServiceGrpc.getReverseTransactionMethod = getReverseTransactionMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.ReverseTransactionRequest, com.game_engine.wallet.v1.ReverseTransactionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ReverseTransaction"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ReverseTransactionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.ReverseTransactionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("ReverseTransaction"))
              .build();
        }
      }
    }
    return getReverseTransactionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetPendingBetsRequest,
      com.game_engine.wallet.v1.GetPendingBetsResponse> getGetPendingBetsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPendingBets",
      requestType = com.game_engine.wallet.v1.GetPendingBetsRequest.class,
      responseType = com.game_engine.wallet.v1.GetPendingBetsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetPendingBetsRequest,
      com.game_engine.wallet.v1.GetPendingBetsResponse> getGetPendingBetsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.wallet.v1.GetPendingBetsRequest, com.game_engine.wallet.v1.GetPendingBetsResponse> getGetPendingBetsMethod;
    if ((getGetPendingBetsMethod = WalletServiceGrpc.getGetPendingBetsMethod) == null) {
      synchronized (WalletServiceGrpc.class) {
        if ((getGetPendingBetsMethod = WalletServiceGrpc.getGetPendingBetsMethod) == null) {
          WalletServiceGrpc.getGetPendingBetsMethod = getGetPendingBetsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.wallet.v1.GetPendingBetsRequest, com.game_engine.wallet.v1.GetPendingBetsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPendingBets"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetPendingBetsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.wallet.v1.GetPendingBetsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletServiceMethodDescriptorSupplier("GetPendingBets"))
              .build();
        }
      }
    }
    return getGetPendingBetsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static WalletServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletServiceStub>() {
        @java.lang.Override
        public WalletServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletServiceStub(channel, callOptions);
        }
      };
    return WalletServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static WalletServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletServiceBlockingStub>() {
        @java.lang.Override
        public WalletServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletServiceBlockingStub(channel, callOptions);
        }
      };
    return WalletServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static WalletServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletServiceFutureStub>() {
        @java.lang.Override
        public WalletServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletServiceFutureStub(channel, callOptions);
        }
      };
    return WalletServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Wallet Service - handles player balances, deposits, withdrawals, and betting
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void getBalance(com.game_engine.wallet.v1.GetBalanceRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetBalanceResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBalanceMethod(), responseObserver);
    }

    /**
     */
    default void getAllBalances(com.game_engine.wallet.v1.GetAllBalancesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetAllBalancesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllBalancesMethod(), responseObserver);
    }

    /**
     */
    default void getTransactionHistory(com.game_engine.wallet.v1.GetTransactionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetTransactionHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTransactionHistoryMethod(), responseObserver);
    }

    /**
     */
    default void createDeposit(com.game_engine.wallet.v1.CreateDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateDepositResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateDepositMethod(), responseObserver);
    }

    /**
     */
    default void confirmDeposit(com.game_engine.wallet.v1.ConfirmDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmDepositResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getConfirmDepositMethod(), responseObserver);
    }

    /**
     */
    default void createWithdrawal(com.game_engine.wallet.v1.CreateWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateWithdrawalResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateWithdrawalMethod(), responseObserver);
    }

    /**
     */
    default void confirmWithdrawal(com.game_engine.wallet.v1.ConfirmWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmWithdrawalResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getConfirmWithdrawalMethod(), responseObserver);
    }

    /**
     */
    default void placeBet(com.game_engine.wallet.v1.PlaceBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.PlaceBetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPlaceBetMethod(), responseObserver);
    }

    /**
     */
    default void settleBet(com.game_engine.wallet.v1.SettleBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.SettleBetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSettleBetMethod(), responseObserver);
    }

    /**
     */
    default void cancelBet(com.game_engine.wallet.v1.CancelBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CancelBetResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCancelBetMethod(), responseObserver);
    }

    /**
     */
    default void createBonusCredit(com.game_engine.wallet.v1.CreateBonusCreditRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateBonusCreditResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateBonusCreditMethod(), responseObserver);
    }

    /**
     */
    default void reverseTransaction(com.game_engine.wallet.v1.ReverseTransactionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ReverseTransactionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getReverseTransactionMethod(), responseObserver);
    }

    /**
     */
    default void getPendingBets(com.game_engine.wallet.v1.GetPendingBetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetPendingBetsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPendingBetsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service WalletService.
   * <pre>
   * Wallet Service - handles player balances, deposits, withdrawals, and betting
   * </pre>
   */
  public static abstract class WalletServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return WalletServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service WalletService.
   * <pre>
   * Wallet Service - handles player balances, deposits, withdrawals, and betting
   * </pre>
   */
  public static final class WalletServiceStub
      extends io.grpc.stub.AbstractAsyncStub<WalletServiceStub> {
    private WalletServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletServiceStub(channel, callOptions);
    }

    /**
     */
    public void getBalance(com.game_engine.wallet.v1.GetBalanceRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetBalanceResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBalanceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllBalances(com.game_engine.wallet.v1.GetAllBalancesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetAllBalancesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllBalancesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTransactionHistory(com.game_engine.wallet.v1.GetTransactionHistoryRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetTransactionHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTransactionHistoryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createDeposit(com.game_engine.wallet.v1.CreateDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateDepositResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateDepositMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void confirmDeposit(com.game_engine.wallet.v1.ConfirmDepositRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmDepositResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getConfirmDepositMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createWithdrawal(com.game_engine.wallet.v1.CreateWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateWithdrawalResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateWithdrawalMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void confirmWithdrawal(com.game_engine.wallet.v1.ConfirmWithdrawalRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmWithdrawalResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getConfirmWithdrawalMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void placeBet(com.game_engine.wallet.v1.PlaceBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.PlaceBetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPlaceBetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void settleBet(com.game_engine.wallet.v1.SettleBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.SettleBetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSettleBetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void cancelBet(com.game_engine.wallet.v1.CancelBetRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CancelBetResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCancelBetMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createBonusCredit(com.game_engine.wallet.v1.CreateBonusCreditRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateBonusCreditResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateBonusCreditMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void reverseTransaction(com.game_engine.wallet.v1.ReverseTransactionRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ReverseTransactionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getReverseTransactionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPendingBets(com.game_engine.wallet.v1.GetPendingBetsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetPendingBetsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPendingBetsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service WalletService.
   * <pre>
   * Wallet Service - handles player balances, deposits, withdrawals, and betting
   * </pre>
   */
  public static final class WalletServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<WalletServiceBlockingStub> {
    private WalletServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.wallet.v1.GetBalanceResponse getBalance(com.game_engine.wallet.v1.GetBalanceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBalanceMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.GetAllBalancesResponse getAllBalances(com.game_engine.wallet.v1.GetAllBalancesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllBalancesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.GetTransactionHistoryResponse getTransactionHistory(com.game_engine.wallet.v1.GetTransactionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTransactionHistoryMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.CreateDepositResponse createDeposit(com.game_engine.wallet.v1.CreateDepositRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateDepositMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.ConfirmDepositResponse confirmDeposit(com.game_engine.wallet.v1.ConfirmDepositRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getConfirmDepositMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.CreateWithdrawalResponse createWithdrawal(com.game_engine.wallet.v1.CreateWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.ConfirmWithdrawalResponse confirmWithdrawal(com.game_engine.wallet.v1.ConfirmWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getConfirmWithdrawalMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.PlaceBetResponse placeBet(com.game_engine.wallet.v1.PlaceBetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPlaceBetMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.SettleBetResponse settleBet(com.game_engine.wallet.v1.SettleBetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSettleBetMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.CancelBetResponse cancelBet(com.game_engine.wallet.v1.CancelBetRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCancelBetMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.CreateBonusCreditResponse createBonusCredit(com.game_engine.wallet.v1.CreateBonusCreditRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateBonusCreditMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.ReverseTransactionResponse reverseTransaction(com.game_engine.wallet.v1.ReverseTransactionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getReverseTransactionMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.wallet.v1.GetPendingBetsResponse getPendingBets(com.game_engine.wallet.v1.GetPendingBetsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPendingBetsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service WalletService.
   * <pre>
   * Wallet Service - handles player balances, deposits, withdrawals, and betting
   * </pre>
   */
  public static final class WalletServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<WalletServiceFutureStub> {
    private WalletServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.GetBalanceResponse> getBalance(
        com.game_engine.wallet.v1.GetBalanceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBalanceMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.GetAllBalancesResponse> getAllBalances(
        com.game_engine.wallet.v1.GetAllBalancesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllBalancesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.GetTransactionHistoryResponse> getTransactionHistory(
        com.game_engine.wallet.v1.GetTransactionHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTransactionHistoryMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.CreateDepositResponse> createDeposit(
        com.game_engine.wallet.v1.CreateDepositRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateDepositMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.ConfirmDepositResponse> confirmDeposit(
        com.game_engine.wallet.v1.ConfirmDepositRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getConfirmDepositMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.CreateWithdrawalResponse> createWithdrawal(
        com.game_engine.wallet.v1.CreateWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateWithdrawalMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.ConfirmWithdrawalResponse> confirmWithdrawal(
        com.game_engine.wallet.v1.ConfirmWithdrawalRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getConfirmWithdrawalMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.PlaceBetResponse> placeBet(
        com.game_engine.wallet.v1.PlaceBetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPlaceBetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.SettleBetResponse> settleBet(
        com.game_engine.wallet.v1.SettleBetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSettleBetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.CancelBetResponse> cancelBet(
        com.game_engine.wallet.v1.CancelBetRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCancelBetMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.CreateBonusCreditResponse> createBonusCredit(
        com.game_engine.wallet.v1.CreateBonusCreditRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateBonusCreditMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.ReverseTransactionResponse> reverseTransaction(
        com.game_engine.wallet.v1.ReverseTransactionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getReverseTransactionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.wallet.v1.GetPendingBetsResponse> getPendingBets(
        com.game_engine.wallet.v1.GetPendingBetsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPendingBetsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_BALANCE = 0;
  private static final int METHODID_GET_ALL_BALANCES = 1;
  private static final int METHODID_GET_TRANSACTION_HISTORY = 2;
  private static final int METHODID_CREATE_DEPOSIT = 3;
  private static final int METHODID_CONFIRM_DEPOSIT = 4;
  private static final int METHODID_CREATE_WITHDRAWAL = 5;
  private static final int METHODID_CONFIRM_WITHDRAWAL = 6;
  private static final int METHODID_PLACE_BET = 7;
  private static final int METHODID_SETTLE_BET = 8;
  private static final int METHODID_CANCEL_BET = 9;
  private static final int METHODID_CREATE_BONUS_CREDIT = 10;
  private static final int METHODID_REVERSE_TRANSACTION = 11;
  private static final int METHODID_GET_PENDING_BETS = 12;

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
        case METHODID_GET_BALANCE:
          serviceImpl.getBalance((com.game_engine.wallet.v1.GetBalanceRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetBalanceResponse>) responseObserver);
          break;
        case METHODID_GET_ALL_BALANCES:
          serviceImpl.getAllBalances((com.game_engine.wallet.v1.GetAllBalancesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetAllBalancesResponse>) responseObserver);
          break;
        case METHODID_GET_TRANSACTION_HISTORY:
          serviceImpl.getTransactionHistory((com.game_engine.wallet.v1.GetTransactionHistoryRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetTransactionHistoryResponse>) responseObserver);
          break;
        case METHODID_CREATE_DEPOSIT:
          serviceImpl.createDeposit((com.game_engine.wallet.v1.CreateDepositRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateDepositResponse>) responseObserver);
          break;
        case METHODID_CONFIRM_DEPOSIT:
          serviceImpl.confirmDeposit((com.game_engine.wallet.v1.ConfirmDepositRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmDepositResponse>) responseObserver);
          break;
        case METHODID_CREATE_WITHDRAWAL:
          serviceImpl.createWithdrawal((com.game_engine.wallet.v1.CreateWithdrawalRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateWithdrawalResponse>) responseObserver);
          break;
        case METHODID_CONFIRM_WITHDRAWAL:
          serviceImpl.confirmWithdrawal((com.game_engine.wallet.v1.ConfirmWithdrawalRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ConfirmWithdrawalResponse>) responseObserver);
          break;
        case METHODID_PLACE_BET:
          serviceImpl.placeBet((com.game_engine.wallet.v1.PlaceBetRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.PlaceBetResponse>) responseObserver);
          break;
        case METHODID_SETTLE_BET:
          serviceImpl.settleBet((com.game_engine.wallet.v1.SettleBetRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.SettleBetResponse>) responseObserver);
          break;
        case METHODID_CANCEL_BET:
          serviceImpl.cancelBet((com.game_engine.wallet.v1.CancelBetRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CancelBetResponse>) responseObserver);
          break;
        case METHODID_CREATE_BONUS_CREDIT:
          serviceImpl.createBonusCredit((com.game_engine.wallet.v1.CreateBonusCreditRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.CreateBonusCreditResponse>) responseObserver);
          break;
        case METHODID_REVERSE_TRANSACTION:
          serviceImpl.reverseTransaction((com.game_engine.wallet.v1.ReverseTransactionRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.ReverseTransactionResponse>) responseObserver);
          break;
        case METHODID_GET_PENDING_BETS:
          serviceImpl.getPendingBets((com.game_engine.wallet.v1.GetPendingBetsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.wallet.v1.GetPendingBetsResponse>) responseObserver);
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
          getGetBalanceMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.GetBalanceRequest,
              com.game_engine.wallet.v1.GetBalanceResponse>(
                service, METHODID_GET_BALANCE)))
        .addMethod(
          getGetAllBalancesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.GetAllBalancesRequest,
              com.game_engine.wallet.v1.GetAllBalancesResponse>(
                service, METHODID_GET_ALL_BALANCES)))
        .addMethod(
          getGetTransactionHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.GetTransactionHistoryRequest,
              com.game_engine.wallet.v1.GetTransactionHistoryResponse>(
                service, METHODID_GET_TRANSACTION_HISTORY)))
        .addMethod(
          getCreateDepositMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.CreateDepositRequest,
              com.game_engine.wallet.v1.CreateDepositResponse>(
                service, METHODID_CREATE_DEPOSIT)))
        .addMethod(
          getConfirmDepositMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.ConfirmDepositRequest,
              com.game_engine.wallet.v1.ConfirmDepositResponse>(
                service, METHODID_CONFIRM_DEPOSIT)))
        .addMethod(
          getCreateWithdrawalMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.CreateWithdrawalRequest,
              com.game_engine.wallet.v1.CreateWithdrawalResponse>(
                service, METHODID_CREATE_WITHDRAWAL)))
        .addMethod(
          getConfirmWithdrawalMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.ConfirmWithdrawalRequest,
              com.game_engine.wallet.v1.ConfirmWithdrawalResponse>(
                service, METHODID_CONFIRM_WITHDRAWAL)))
        .addMethod(
          getPlaceBetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.PlaceBetRequest,
              com.game_engine.wallet.v1.PlaceBetResponse>(
                service, METHODID_PLACE_BET)))
        .addMethod(
          getSettleBetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.SettleBetRequest,
              com.game_engine.wallet.v1.SettleBetResponse>(
                service, METHODID_SETTLE_BET)))
        .addMethod(
          getCancelBetMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.CancelBetRequest,
              com.game_engine.wallet.v1.CancelBetResponse>(
                service, METHODID_CANCEL_BET)))
        .addMethod(
          getCreateBonusCreditMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.CreateBonusCreditRequest,
              com.game_engine.wallet.v1.CreateBonusCreditResponse>(
                service, METHODID_CREATE_BONUS_CREDIT)))
        .addMethod(
          getReverseTransactionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.ReverseTransactionRequest,
              com.game_engine.wallet.v1.ReverseTransactionResponse>(
                service, METHODID_REVERSE_TRANSACTION)))
        .addMethod(
          getGetPendingBetsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.wallet.v1.GetPendingBetsRequest,
              com.game_engine.wallet.v1.GetPendingBetsResponse>(
                service, METHODID_GET_PENDING_BETS)))
        .build();
  }

  private static abstract class WalletServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    WalletServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.wallet.v1.WalletServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("WalletService");
    }
  }

  private static final class WalletServiceFileDescriptorSupplier
      extends WalletServiceBaseDescriptorSupplier {
    WalletServiceFileDescriptorSupplier() {}
  }

  private static final class WalletServiceMethodDescriptorSupplier
      extends WalletServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    WalletServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (WalletServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new WalletServiceFileDescriptorSupplier())
              .addMethod(getGetBalanceMethod())
              .addMethod(getGetAllBalancesMethod())
              .addMethod(getGetTransactionHistoryMethod())
              .addMethod(getCreateDepositMethod())
              .addMethod(getConfirmDepositMethod())
              .addMethod(getCreateWithdrawalMethod())
              .addMethod(getConfirmWithdrawalMethod())
              .addMethod(getPlaceBetMethod())
              .addMethod(getSettleBetMethod())
              .addMethod(getCancelBetMethod())
              .addMethod(getCreateBonusCreditMethod())
              .addMethod(getReverseTransactionMethod())
              .addMethod(getGetPendingBetsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
