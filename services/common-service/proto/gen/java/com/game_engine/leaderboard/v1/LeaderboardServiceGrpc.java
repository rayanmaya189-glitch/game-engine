package com.game_engine.leaderboard.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * LeaderboardService manages leaderboards and player rankings
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class LeaderboardServiceGrpc {

  private LeaderboardServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.leaderboard.v1.LeaderboardService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetDailyLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetDailyLeaderboard",
      requestType = com.game_engine.leaderboard.v1.GetLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.GetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetDailyLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetDailyLeaderboardMethod;
    if ((getGetDailyLeaderboardMethod = LeaderboardServiceGrpc.getGetDailyLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getGetDailyLeaderboardMethod = LeaderboardServiceGrpc.getGetDailyLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getGetDailyLeaderboardMethod = getGetDailyLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetDailyLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("GetDailyLeaderboard"))
              .build();
        }
      }
    }
    return getGetDailyLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetWeeklyLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetWeeklyLeaderboard",
      requestType = com.game_engine.leaderboard.v1.GetLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.GetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetWeeklyLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetWeeklyLeaderboardMethod;
    if ((getGetWeeklyLeaderboardMethod = LeaderboardServiceGrpc.getGetWeeklyLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getGetWeeklyLeaderboardMethod = LeaderboardServiceGrpc.getGetWeeklyLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getGetWeeklyLeaderboardMethod = getGetWeeklyLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetWeeklyLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("GetWeeklyLeaderboard"))
              .build();
        }
      }
    }
    return getGetWeeklyLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetMonthlyLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetMonthlyLeaderboard",
      requestType = com.game_engine.leaderboard.v1.GetLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.GetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetMonthlyLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetMonthlyLeaderboardMethod;
    if ((getGetMonthlyLeaderboardMethod = LeaderboardServiceGrpc.getGetMonthlyLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getGetMonthlyLeaderboardMethod = LeaderboardServiceGrpc.getGetMonthlyLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getGetMonthlyLeaderboardMethod = getGetMonthlyLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetMonthlyLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("GetMonthlyLeaderboard"))
              .build();
        }
      }
    }
    return getGetMonthlyLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetAllTimeLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAllTimeLeaderboard",
      requestType = com.game_engine.leaderboard.v1.GetLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.GetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest,
      com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetAllTimeLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse> getGetAllTimeLeaderboardMethod;
    if ((getGetAllTimeLeaderboardMethod = LeaderboardServiceGrpc.getGetAllTimeLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getGetAllTimeLeaderboardMethod = LeaderboardServiceGrpc.getGetAllTimeLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getGetAllTimeLeaderboardMethod = getGetAllTimeLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.GetLeaderboardRequest, com.game_engine.leaderboard.v1.GetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllTimeLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("GetAllTimeLeaderboard"))
              .build();
        }
      }
    }
    return getGetAllTimeLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetPlayerRankRequest,
      com.game_engine.leaderboard.v1.GetPlayerRankResponse> getGetPlayerRankMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPlayerRank",
      requestType = com.game_engine.leaderboard.v1.GetPlayerRankRequest.class,
      responseType = com.game_engine.leaderboard.v1.GetPlayerRankResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetPlayerRankRequest,
      com.game_engine.leaderboard.v1.GetPlayerRankResponse> getGetPlayerRankMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.GetPlayerRankRequest, com.game_engine.leaderboard.v1.GetPlayerRankResponse> getGetPlayerRankMethod;
    if ((getGetPlayerRankMethod = LeaderboardServiceGrpc.getGetPlayerRankMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getGetPlayerRankMethod = LeaderboardServiceGrpc.getGetPlayerRankMethod) == null) {
          LeaderboardServiceGrpc.getGetPlayerRankMethod = getGetPlayerRankMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.GetPlayerRankRequest, com.game_engine.leaderboard.v1.GetPlayerRankResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPlayerRank"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetPlayerRankRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.GetPlayerRankResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("GetPlayerRank"))
              .build();
        }
      }
    }
    return getGetPlayerRankMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest,
      com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePlayerScore",
      requestType = com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest.class,
      responseType = com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest,
      com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest, com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod;
    if ((getUpdatePlayerScoreMethod = LeaderboardServiceGrpc.getUpdatePlayerScoreMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getUpdatePlayerScoreMethod = LeaderboardServiceGrpc.getUpdatePlayerScoreMethod) == null) {
          LeaderboardServiceGrpc.getUpdatePlayerScoreMethod = getUpdatePlayerScoreMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest, com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePlayerScore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("UpdatePlayerScore"))
              .build();
        }
      }
    }
    return getUpdatePlayerScoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.DistributePrizesRequest,
      com.game_engine.leaderboard.v1.DistributePrizesResponse> getDistributePrizesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DistributePrizes",
      requestType = com.game_engine.leaderboard.v1.DistributePrizesRequest.class,
      responseType = com.game_engine.leaderboard.v1.DistributePrizesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.DistributePrizesRequest,
      com.game_engine.leaderboard.v1.DistributePrizesResponse> getDistributePrizesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.DistributePrizesRequest, com.game_engine.leaderboard.v1.DistributePrizesResponse> getDistributePrizesMethod;
    if ((getDistributePrizesMethod = LeaderboardServiceGrpc.getDistributePrizesMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getDistributePrizesMethod = LeaderboardServiceGrpc.getDistributePrizesMethod) == null) {
          LeaderboardServiceGrpc.getDistributePrizesMethod = getDistributePrizesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.DistributePrizesRequest, com.game_engine.leaderboard.v1.DistributePrizesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DistributePrizes"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.DistributePrizesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.DistributePrizesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("DistributePrizes"))
              .build();
        }
      }
    }
    return getDistributePrizesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.SyncLeaderboardRequest,
      com.game_engine.leaderboard.v1.SyncLeaderboardResponse> getSyncLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SyncLeaderboard",
      requestType = com.game_engine.leaderboard.v1.SyncLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.SyncLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.SyncLeaderboardRequest,
      com.game_engine.leaderboard.v1.SyncLeaderboardResponse> getSyncLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.SyncLeaderboardRequest, com.game_engine.leaderboard.v1.SyncLeaderboardResponse> getSyncLeaderboardMethod;
    if ((getSyncLeaderboardMethod = LeaderboardServiceGrpc.getSyncLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getSyncLeaderboardMethod = LeaderboardServiceGrpc.getSyncLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getSyncLeaderboardMethod = getSyncLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.SyncLeaderboardRequest, com.game_engine.leaderboard.v1.SyncLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SyncLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.SyncLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.SyncLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("SyncLeaderboard"))
              .build();
        }
      }
    }
    return getSyncLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.ResetLeaderboardRequest,
      com.game_engine.leaderboard.v1.ResetLeaderboardResponse> getResetLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ResetLeaderboard",
      requestType = com.game_engine.leaderboard.v1.ResetLeaderboardRequest.class,
      responseType = com.game_engine.leaderboard.v1.ResetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.ResetLeaderboardRequest,
      com.game_engine.leaderboard.v1.ResetLeaderboardResponse> getResetLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.leaderboard.v1.ResetLeaderboardRequest, com.game_engine.leaderboard.v1.ResetLeaderboardResponse> getResetLeaderboardMethod;
    if ((getResetLeaderboardMethod = LeaderboardServiceGrpc.getResetLeaderboardMethod) == null) {
      synchronized (LeaderboardServiceGrpc.class) {
        if ((getResetLeaderboardMethod = LeaderboardServiceGrpc.getResetLeaderboardMethod) == null) {
          LeaderboardServiceGrpc.getResetLeaderboardMethod = getResetLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.leaderboard.v1.ResetLeaderboardRequest, com.game_engine.leaderboard.v1.ResetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ResetLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.ResetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.leaderboard.v1.ResetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new LeaderboardServiceMethodDescriptorSupplier("ResetLeaderboard"))
              .build();
        }
      }
    }
    return getResetLeaderboardMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static LeaderboardServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceStub>() {
        @java.lang.Override
        public LeaderboardServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LeaderboardServiceStub(channel, callOptions);
        }
      };
    return LeaderboardServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static LeaderboardServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceBlockingV2Stub>() {
        @java.lang.Override
        public LeaderboardServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LeaderboardServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return LeaderboardServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static LeaderboardServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceBlockingStub>() {
        @java.lang.Override
        public LeaderboardServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LeaderboardServiceBlockingStub(channel, callOptions);
        }
      };
    return LeaderboardServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static LeaderboardServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<LeaderboardServiceFutureStub>() {
        @java.lang.Override
        public LeaderboardServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new LeaderboardServiceFutureStub(channel, callOptions);
        }
      };
    return LeaderboardServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void getDailyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetDailyLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void getWeeklyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetWeeklyLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void getMonthlyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetMonthlyLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void getAllTimeLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllTimeLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void getPlayerRank(com.game_engine.leaderboard.v1.GetPlayerRankRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetPlayerRankResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPlayerRankMethod(), responseObserver);
    }

    /**
     */
    default void updatePlayerScore(com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePlayerScoreMethod(), responseObserver);
    }

    /**
     */
    default void distributePrizes(com.game_engine.leaderboard.v1.DistributePrizesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.DistributePrizesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDistributePrizesMethod(), responseObserver);
    }

    /**
     */
    default void syncLeaderboard(com.game_engine.leaderboard.v1.SyncLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.SyncLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSyncLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void resetLeaderboard(com.game_engine.leaderboard.v1.ResetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.ResetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getResetLeaderboardMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service LeaderboardService.
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public static abstract class LeaderboardServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return LeaderboardServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service LeaderboardService.
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public static final class LeaderboardServiceStub
      extends io.grpc.stub.AbstractAsyncStub<LeaderboardServiceStub> {
    private LeaderboardServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LeaderboardServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LeaderboardServiceStub(channel, callOptions);
    }

    /**
     */
    public void getDailyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetDailyLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getWeeklyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetWeeklyLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getMonthlyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetMonthlyLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAllTimeLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllTimeLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPlayerRank(com.game_engine.leaderboard.v1.GetPlayerRankRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetPlayerRankResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPlayerRankMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updatePlayerScore(com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePlayerScoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void distributePrizes(com.game_engine.leaderboard.v1.DistributePrizesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.DistributePrizesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDistributePrizesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void syncLeaderboard(com.game_engine.leaderboard.v1.SyncLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.SyncLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSyncLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void resetLeaderboard(com.game_engine.leaderboard.v1.ResetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.ResetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getResetLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service LeaderboardService.
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public static final class LeaderboardServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<LeaderboardServiceBlockingV2Stub> {
    private LeaderboardServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LeaderboardServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LeaderboardServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getDailyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetDailyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getWeeklyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetWeeklyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getMonthlyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetMonthlyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getAllTimeLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetAllTimeLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetPlayerRankResponse getPlayerRank(com.game_engine.leaderboard.v1.GetPlayerRankRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPlayerRankMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse updatePlayerScore(com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePlayerScoreMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.DistributePrizesResponse distributePrizes(com.game_engine.leaderboard.v1.DistributePrizesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getDistributePrizesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.SyncLeaderboardResponse syncLeaderboard(com.game_engine.leaderboard.v1.SyncLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSyncLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.ResetLeaderboardResponse resetLeaderboard(com.game_engine.leaderboard.v1.ResetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getResetLeaderboardMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service LeaderboardService.
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public static final class LeaderboardServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<LeaderboardServiceBlockingStub> {
    private LeaderboardServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LeaderboardServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LeaderboardServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getDailyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetDailyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getWeeklyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetWeeklyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getMonthlyLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetMonthlyLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetLeaderboardResponse getAllTimeLeaderboard(com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllTimeLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.GetPlayerRankResponse getPlayerRank(com.game_engine.leaderboard.v1.GetPlayerRankRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPlayerRankMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse updatePlayerScore(com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePlayerScoreMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.DistributePrizesResponse distributePrizes(com.game_engine.leaderboard.v1.DistributePrizesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDistributePrizesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.SyncLeaderboardResponse syncLeaderboard(com.game_engine.leaderboard.v1.SyncLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSyncLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.leaderboard.v1.ResetLeaderboardResponse resetLeaderboard(com.game_engine.leaderboard.v1.ResetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getResetLeaderboardMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service LeaderboardService.
   * <pre>
   * LeaderboardService manages leaderboards and player rankings
   * </pre>
   */
  public static final class LeaderboardServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<LeaderboardServiceFutureStub> {
    private LeaderboardServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected LeaderboardServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new LeaderboardServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.GetLeaderboardResponse> getDailyLeaderboard(
        com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetDailyLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.GetLeaderboardResponse> getWeeklyLeaderboard(
        com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetWeeklyLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.GetLeaderboardResponse> getMonthlyLeaderboard(
        com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetMonthlyLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.GetLeaderboardResponse> getAllTimeLeaderboard(
        com.game_engine.leaderboard.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllTimeLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.GetPlayerRankResponse> getPlayerRank(
        com.game_engine.leaderboard.v1.GetPlayerRankRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPlayerRankMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse> updatePlayerScore(
        com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePlayerScoreMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.DistributePrizesResponse> distributePrizes(
        com.game_engine.leaderboard.v1.DistributePrizesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDistributePrizesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.SyncLeaderboardResponse> syncLeaderboard(
        com.game_engine.leaderboard.v1.SyncLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSyncLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.leaderboard.v1.ResetLeaderboardResponse> resetLeaderboard(
        com.game_engine.leaderboard.v1.ResetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getResetLeaderboardMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_DAILY_LEADERBOARD = 0;
  private static final int METHODID_GET_WEEKLY_LEADERBOARD = 1;
  private static final int METHODID_GET_MONTHLY_LEADERBOARD = 2;
  private static final int METHODID_GET_ALL_TIME_LEADERBOARD = 3;
  private static final int METHODID_GET_PLAYER_RANK = 4;
  private static final int METHODID_UPDATE_PLAYER_SCORE = 5;
  private static final int METHODID_DISTRIBUTE_PRIZES = 6;
  private static final int METHODID_SYNC_LEADERBOARD = 7;
  private static final int METHODID_RESET_LEADERBOARD = 8;

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
        case METHODID_GET_DAILY_LEADERBOARD:
          serviceImpl.getDailyLeaderboard((com.game_engine.leaderboard.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_GET_WEEKLY_LEADERBOARD:
          serviceImpl.getWeeklyLeaderboard((com.game_engine.leaderboard.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_GET_MONTHLY_LEADERBOARD:
          serviceImpl.getMonthlyLeaderboard((com.game_engine.leaderboard.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_GET_ALL_TIME_LEADERBOARD:
          serviceImpl.getAllTimeLeaderboard((com.game_engine.leaderboard.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_GET_PLAYER_RANK:
          serviceImpl.getPlayerRank((com.game_engine.leaderboard.v1.GetPlayerRankRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.GetPlayerRankResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PLAYER_SCORE:
          serviceImpl.updatePlayerScore((com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse>) responseObserver);
          break;
        case METHODID_DISTRIBUTE_PRIZES:
          serviceImpl.distributePrizes((com.game_engine.leaderboard.v1.DistributePrizesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.DistributePrizesResponse>) responseObserver);
          break;
        case METHODID_SYNC_LEADERBOARD:
          serviceImpl.syncLeaderboard((com.game_engine.leaderboard.v1.SyncLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.SyncLeaderboardResponse>) responseObserver);
          break;
        case METHODID_RESET_LEADERBOARD:
          serviceImpl.resetLeaderboard((com.game_engine.leaderboard.v1.ResetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.leaderboard.v1.ResetLeaderboardResponse>) responseObserver);
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
          getGetDailyLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.GetLeaderboardRequest,
              com.game_engine.leaderboard.v1.GetLeaderboardResponse>(
                service, METHODID_GET_DAILY_LEADERBOARD)))
        .addMethod(
          getGetWeeklyLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.GetLeaderboardRequest,
              com.game_engine.leaderboard.v1.GetLeaderboardResponse>(
                service, METHODID_GET_WEEKLY_LEADERBOARD)))
        .addMethod(
          getGetMonthlyLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.GetLeaderboardRequest,
              com.game_engine.leaderboard.v1.GetLeaderboardResponse>(
                service, METHODID_GET_MONTHLY_LEADERBOARD)))
        .addMethod(
          getGetAllTimeLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.GetLeaderboardRequest,
              com.game_engine.leaderboard.v1.GetLeaderboardResponse>(
                service, METHODID_GET_ALL_TIME_LEADERBOARD)))
        .addMethod(
          getGetPlayerRankMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.GetPlayerRankRequest,
              com.game_engine.leaderboard.v1.GetPlayerRankResponse>(
                service, METHODID_GET_PLAYER_RANK)))
        .addMethod(
          getUpdatePlayerScoreMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.UpdatePlayerScoreRequest,
              com.game_engine.leaderboard.v1.UpdatePlayerScoreResponse>(
                service, METHODID_UPDATE_PLAYER_SCORE)))
        .addMethod(
          getDistributePrizesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.DistributePrizesRequest,
              com.game_engine.leaderboard.v1.DistributePrizesResponse>(
                service, METHODID_DISTRIBUTE_PRIZES)))
        .addMethod(
          getSyncLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.SyncLeaderboardRequest,
              com.game_engine.leaderboard.v1.SyncLeaderboardResponse>(
                service, METHODID_SYNC_LEADERBOARD)))
        .addMethod(
          getResetLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.leaderboard.v1.ResetLeaderboardRequest,
              com.game_engine.leaderboard.v1.ResetLeaderboardResponse>(
                service, METHODID_RESET_LEADERBOARD)))
        .build();
  }

  private static abstract class LeaderboardServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    LeaderboardServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.leaderboard.v1.LeaderboardServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("LeaderboardService");
    }
  }

  private static final class LeaderboardServiceFileDescriptorSupplier
      extends LeaderboardServiceBaseDescriptorSupplier {
    LeaderboardServiceFileDescriptorSupplier() {}
  }

  private static final class LeaderboardServiceMethodDescriptorSupplier
      extends LeaderboardServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    LeaderboardServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (LeaderboardServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new LeaderboardServiceFileDescriptorSupplier())
              .addMethod(getGetDailyLeaderboardMethod())
              .addMethod(getGetWeeklyLeaderboardMethod())
              .addMethod(getGetMonthlyLeaderboardMethod())
              .addMethod(getGetAllTimeLeaderboardMethod())
              .addMethod(getGetPlayerRankMethod())
              .addMethod(getUpdatePlayerScoreMethod())
              .addMethod(getDistributePrizesMethod())
              .addMethod(getSyncLeaderboardMethod())
              .addMethod(getResetLeaderboardMethod())
              .build();
        }
      }
    }
    return result;
  }
}
