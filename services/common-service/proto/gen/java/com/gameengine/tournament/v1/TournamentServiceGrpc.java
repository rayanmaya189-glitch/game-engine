package com.game_engine.tournament.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class TournamentServiceGrpc {

  private TournamentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.tournament.v1.TournamentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.ListTournamentsRequest,
      com.game_engine.tournament.v1.ListTournamentsResponse> getListTournamentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListTournaments",
      requestType = com.game_engine.tournament.v1.ListTournamentsRequest.class,
      responseType = com.game_engine.tournament.v1.ListTournamentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.ListTournamentsRequest,
      com.game_engine.tournament.v1.ListTournamentsResponse> getListTournamentsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.ListTournamentsRequest, com.game_engine.tournament.v1.ListTournamentsResponse> getListTournamentsMethod;
    if ((getListTournamentsMethod = TournamentServiceGrpc.getListTournamentsMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getListTournamentsMethod = TournamentServiceGrpc.getListTournamentsMethod) == null) {
          TournamentServiceGrpc.getListTournamentsMethod = getListTournamentsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.ListTournamentsRequest, com.game_engine.tournament.v1.ListTournamentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListTournaments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.ListTournamentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.ListTournamentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("ListTournaments"))
              .build();
        }
      }
    }
    return getListTournamentsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest,
      com.game_engine.tournament.v1.GetTournamentResponse> getGetTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTournament",
      requestType = com.game_engine.tournament.v1.GetTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.GetTournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest,
      com.game_engine.tournament.v1.GetTournamentResponse> getGetTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest, com.game_engine.tournament.v1.GetTournamentResponse> getGetTournamentMethod;
    if ((getGetTournamentMethod = TournamentServiceGrpc.getGetTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetTournamentMethod = TournamentServiceGrpc.getGetTournamentMethod) == null) {
          TournamentServiceGrpc.getGetTournamentMethod = getGetTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetTournamentRequest, com.game_engine.tournament.v1.GetTournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetTournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetTournament"))
              .build();
        }
      }
    }
    return getGetTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.JoinTournamentRequest,
      com.game_engine.tournament.v1.JoinTournamentResponse> getJoinTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "JoinTournament",
      requestType = com.game_engine.tournament.v1.JoinTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.JoinTournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.JoinTournamentRequest,
      com.game_engine.tournament.v1.JoinTournamentResponse> getJoinTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.JoinTournamentRequest, com.game_engine.tournament.v1.JoinTournamentResponse> getJoinTournamentMethod;
    if ((getJoinTournamentMethod = TournamentServiceGrpc.getJoinTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getJoinTournamentMethod = TournamentServiceGrpc.getJoinTournamentMethod) == null) {
          TournamentServiceGrpc.getJoinTournamentMethod = getJoinTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.JoinTournamentRequest, com.game_engine.tournament.v1.JoinTournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "JoinTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.JoinTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.JoinTournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("JoinTournament"))
              .build();
        }
      }
    }
    return getJoinTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.LeaveTournamentRequest,
      com.game_engine.tournament.v1.LeaveTournamentResponse> getLeaveTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LeaveTournament",
      requestType = com.game_engine.tournament.v1.LeaveTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.LeaveTournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.LeaveTournamentRequest,
      com.game_engine.tournament.v1.LeaveTournamentResponse> getLeaveTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.LeaveTournamentRequest, com.game_engine.tournament.v1.LeaveTournamentResponse> getLeaveTournamentMethod;
    if ((getLeaveTournamentMethod = TournamentServiceGrpc.getLeaveTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getLeaveTournamentMethod = TournamentServiceGrpc.getLeaveTournamentMethod) == null) {
          TournamentServiceGrpc.getLeaveTournamentMethod = getLeaveTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.LeaveTournamentRequest, com.game_engine.tournament.v1.LeaveTournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "LeaveTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.LeaveTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.LeaveTournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("LeaveTournament"))
              .build();
        }
      }
    }
    return getLeaveTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetLeaderboardRequest,
      com.game_engine.tournament.v1.GetLeaderboardResponse> getGetLeaderboardMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetLeaderboard",
      requestType = com.game_engine.tournament.v1.GetLeaderboardRequest.class,
      responseType = com.game_engine.tournament.v1.GetLeaderboardResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetLeaderboardRequest,
      com.game_engine.tournament.v1.GetLeaderboardResponse> getGetLeaderboardMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetLeaderboardRequest, com.game_engine.tournament.v1.GetLeaderboardResponse> getGetLeaderboardMethod;
    if ((getGetLeaderboardMethod = TournamentServiceGrpc.getGetLeaderboardMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetLeaderboardMethod = TournamentServiceGrpc.getGetLeaderboardMethod) == null) {
          TournamentServiceGrpc.getGetLeaderboardMethod = getGetLeaderboardMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetLeaderboardRequest, com.game_engine.tournament.v1.GetLeaderboardResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetLeaderboard"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetLeaderboardRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetLeaderboardResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetLeaderboard"))
              .build();
        }
      }
    }
    return getGetLeaderboardMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateScoreRequest,
      com.game_engine.tournament.v1.UpdateScoreResponse> getUpdateScoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateScore",
      requestType = com.game_engine.tournament.v1.UpdateScoreRequest.class,
      responseType = com.game_engine.tournament.v1.UpdateScoreResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateScoreRequest,
      com.game_engine.tournament.v1.UpdateScoreResponse> getUpdateScoreMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateScoreRequest, com.game_engine.tournament.v1.UpdateScoreResponse> getUpdateScoreMethod;
    if ((getUpdateScoreMethod = TournamentServiceGrpc.getUpdateScoreMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getUpdateScoreMethod = TournamentServiceGrpc.getUpdateScoreMethod) == null) {
          TournamentServiceGrpc.getUpdateScoreMethod = getUpdateScoreMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.UpdateScoreRequest, com.game_engine.tournament.v1.UpdateScoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateScore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UpdateScoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UpdateScoreResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("UpdateScore"))
              .build();
        }
      }
    }
    return getUpdateScoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetMyTournamentsRequest,
      com.game_engine.tournament.v1.GetMyTournamentsResponse> getGetMyTournamentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetMyTournaments",
      requestType = com.game_engine.tournament.v1.GetMyTournamentsRequest.class,
      responseType = com.game_engine.tournament.v1.GetMyTournamentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetMyTournamentsRequest,
      com.game_engine.tournament.v1.GetMyTournamentsResponse> getGetMyTournamentsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetMyTournamentsRequest, com.game_engine.tournament.v1.GetMyTournamentsResponse> getGetMyTournamentsMethod;
    if ((getGetMyTournamentsMethod = TournamentServiceGrpc.getGetMyTournamentsMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetMyTournamentsMethod = TournamentServiceGrpc.getGetMyTournamentsMethod) == null) {
          TournamentServiceGrpc.getGetMyTournamentsMethod = getGetMyTournamentsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetMyTournamentsRequest, com.game_engine.tournament.v1.GetMyTournamentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetMyTournaments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetMyTournamentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetMyTournamentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetMyTournaments"))
              .build();
        }
      }
    }
    return getGetMyTournamentsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static TournamentServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TournamentServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TournamentServiceStub>() {
        @java.lang.Override
        public TournamentServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TournamentServiceStub(channel, callOptions);
        }
      };
    return TournamentServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static TournamentServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TournamentServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TournamentServiceBlockingV2Stub>() {
        @java.lang.Override
        public TournamentServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TournamentServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return TournamentServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static TournamentServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TournamentServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TournamentServiceBlockingStub>() {
        @java.lang.Override
        public TournamentServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TournamentServiceBlockingStub(channel, callOptions);
        }
      };
    return TournamentServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static TournamentServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<TournamentServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<TournamentServiceFutureStub>() {
        @java.lang.Override
        public TournamentServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new TournamentServiceFutureStub(channel, callOptions);
        }
      };
    return TournamentServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Tournament operations
     * </pre>
     */
    default void listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListTournamentsMethod(), responseObserver);
    }

    /**
     */
    default void getTournament(com.game_engine.tournament.v1.GetTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTournamentMethod(), responseObserver);
    }

    /**
     */
    default void joinTournament(com.game_engine.tournament.v1.JoinTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.JoinTournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getJoinTournamentMethod(), responseObserver);
    }

    /**
     */
    default void leaveTournament(com.game_engine.tournament.v1.LeaveTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.LeaveTournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLeaveTournamentMethod(), responseObserver);
    }

    /**
     * <pre>
     * Leaderboard
     * </pre>
     */
    default void getLeaderboard(com.game_engine.tournament.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetLeaderboardMethod(), responseObserver);
    }

    /**
     */
    default void updateScore(com.game_engine.tournament.v1.UpdateScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdateScoreResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateScoreMethod(), responseObserver);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    default void getMyTournaments(com.game_engine.tournament.v1.GetMyTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetMyTournamentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetMyTournamentsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service TournamentService.
   */
  public static abstract class TournamentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return TournamentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service TournamentService.
   */
  public static final class TournamentServiceStub
      extends io.grpc.stub.AbstractAsyncStub<TournamentServiceStub> {
    private TournamentServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TournamentServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TournamentServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Tournament operations
     * </pre>
     */
    public void listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListTournamentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTournament(com.game_engine.tournament.v1.GetTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void joinTournament(com.game_engine.tournament.v1.JoinTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.JoinTournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getJoinTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void leaveTournament(com.game_engine.tournament.v1.LeaveTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.LeaveTournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLeaveTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Leaderboard
     * </pre>
     */
    public void getLeaderboard(com.game_engine.tournament.v1.GetLeaderboardRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetLeaderboardResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetLeaderboardMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateScore(com.game_engine.tournament.v1.UpdateScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdateScoreResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateScoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public void getMyTournaments(com.game_engine.tournament.v1.GetMyTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetMyTournamentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetMyTournamentsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service TournamentService.
   */
  public static final class TournamentServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<TournamentServiceBlockingV2Stub> {
    private TournamentServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TournamentServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TournamentServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Tournament operations
     * </pre>
     */
    public com.game_engine.tournament.v1.ListTournamentsResponse listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListTournamentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.GetTournamentResponse getTournament(com.game_engine.tournament.v1.GetTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.JoinTournamentResponse joinTournament(com.game_engine.tournament.v1.JoinTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getJoinTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.LeaveTournamentResponse leaveTournament(com.game_engine.tournament.v1.LeaveTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getLeaveTournamentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Leaderboard
     * </pre>
     */
    public com.game_engine.tournament.v1.GetLeaderboardResponse getLeaderboard(com.game_engine.tournament.v1.GetLeaderboardRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.UpdateScoreResponse updateScore(com.game_engine.tournament.v1.UpdateScoreRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateScoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.game_engine.tournament.v1.GetMyTournamentsResponse getMyTournaments(com.game_engine.tournament.v1.GetMyTournamentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetMyTournamentsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service TournamentService.
   */
  public static final class TournamentServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<TournamentServiceBlockingStub> {
    private TournamentServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TournamentServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TournamentServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Tournament operations
     * </pre>
     */
    public com.game_engine.tournament.v1.ListTournamentsResponse listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListTournamentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.GetTournamentResponse getTournament(com.game_engine.tournament.v1.GetTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.JoinTournamentResponse joinTournament(com.game_engine.tournament.v1.JoinTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getJoinTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.LeaveTournamentResponse leaveTournament(com.game_engine.tournament.v1.LeaveTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLeaveTournamentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Leaderboard
     * </pre>
     */
    public com.game_engine.tournament.v1.GetLeaderboardResponse getLeaderboard(com.game_engine.tournament.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetLeaderboardMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.UpdateScoreResponse updateScore(com.game_engine.tournament.v1.UpdateScoreRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateScoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.game_engine.tournament.v1.GetMyTournamentsResponse getMyTournaments(com.game_engine.tournament.v1.GetMyTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetMyTournamentsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service TournamentService.
   */
  public static final class TournamentServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<TournamentServiceFutureStub> {
    private TournamentServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected TournamentServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new TournamentServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Tournament operations
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.ListTournamentsResponse> listTournaments(
        com.game_engine.tournament.v1.ListTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListTournamentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.GetTournamentResponse> getTournament(
        com.game_engine.tournament.v1.GetTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTournamentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.JoinTournamentResponse> joinTournament(
        com.game_engine.tournament.v1.JoinTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getJoinTournamentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.LeaveTournamentResponse> leaveTournament(
        com.game_engine.tournament.v1.LeaveTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLeaveTournamentMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Leaderboard
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.GetLeaderboardResponse> getLeaderboard(
        com.game_engine.tournament.v1.GetLeaderboardRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetLeaderboardMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.UpdateScoreResponse> updateScore(
        com.game_engine.tournament.v1.UpdateScoreRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateScoreMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.GetMyTournamentsResponse> getMyTournaments(
        com.game_engine.tournament.v1.GetMyTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetMyTournamentsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_TOURNAMENTS = 0;
  private static final int METHODID_GET_TOURNAMENT = 1;
  private static final int METHODID_JOIN_TOURNAMENT = 2;
  private static final int METHODID_LEAVE_TOURNAMENT = 3;
  private static final int METHODID_GET_LEADERBOARD = 4;
  private static final int METHODID_UPDATE_SCORE = 5;
  private static final int METHODID_GET_MY_TOURNAMENTS = 6;

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
        case METHODID_LIST_TOURNAMENTS:
          serviceImpl.listTournaments((com.game_engine.tournament.v1.ListTournamentsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse>) responseObserver);
          break;
        case METHODID_GET_TOURNAMENT:
          serviceImpl.getTournament((com.game_engine.tournament.v1.GetTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentResponse>) responseObserver);
          break;
        case METHODID_JOIN_TOURNAMENT:
          serviceImpl.joinTournament((com.game_engine.tournament.v1.JoinTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.JoinTournamentResponse>) responseObserver);
          break;
        case METHODID_LEAVE_TOURNAMENT:
          serviceImpl.leaveTournament((com.game_engine.tournament.v1.LeaveTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.LeaveTournamentResponse>) responseObserver);
          break;
        case METHODID_GET_LEADERBOARD:
          serviceImpl.getLeaderboard((com.game_engine.tournament.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_UPDATE_SCORE:
          serviceImpl.updateScore((com.game_engine.tournament.v1.UpdateScoreRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdateScoreResponse>) responseObserver);
          break;
        case METHODID_GET_MY_TOURNAMENTS:
          serviceImpl.getMyTournaments((com.game_engine.tournament.v1.GetMyTournamentsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetMyTournamentsResponse>) responseObserver);
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
          getListTournamentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.ListTournamentsRequest,
              com.game_engine.tournament.v1.ListTournamentsResponse>(
                service, METHODID_LIST_TOURNAMENTS)))
        .addMethod(
          getGetTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetTournamentRequest,
              com.game_engine.tournament.v1.GetTournamentResponse>(
                service, METHODID_GET_TOURNAMENT)))
        .addMethod(
          getJoinTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.JoinTournamentRequest,
              com.game_engine.tournament.v1.JoinTournamentResponse>(
                service, METHODID_JOIN_TOURNAMENT)))
        .addMethod(
          getLeaveTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.LeaveTournamentRequest,
              com.game_engine.tournament.v1.LeaveTournamentResponse>(
                service, METHODID_LEAVE_TOURNAMENT)))
        .addMethod(
          getGetLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetLeaderboardRequest,
              com.game_engine.tournament.v1.GetLeaderboardResponse>(
                service, METHODID_GET_LEADERBOARD)))
        .addMethod(
          getUpdateScoreMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.UpdateScoreRequest,
              com.game_engine.tournament.v1.UpdateScoreResponse>(
                service, METHODID_UPDATE_SCORE)))
        .addMethod(
          getGetMyTournamentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetMyTournamentsRequest,
              com.game_engine.tournament.v1.GetMyTournamentsResponse>(
                service, METHODID_GET_MY_TOURNAMENTS)))
        .build();
  }

  private static abstract class TournamentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    TournamentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.tournament.v1.TournamentServiceOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("TournamentService");
    }
  }

  private static final class TournamentServiceFileDescriptorSupplier
      extends TournamentServiceBaseDescriptorSupplier {
    TournamentServiceFileDescriptorSupplier() {}
  }

  private static final class TournamentServiceMethodDescriptorSupplier
      extends TournamentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    TournamentServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (TournamentServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new TournamentServiceFileDescriptorSupplier())
              .addMethod(getListTournamentsMethod())
              .addMethod(getGetTournamentMethod())
              .addMethod(getJoinTournamentMethod())
              .addMethod(getLeaveTournamentMethod())
              .addMethod(getGetLeaderboardMethod())
              .addMethod(getUpdateScoreMethod())
              .addMethod(getGetMyTournamentsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
