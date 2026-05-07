package com.game_engine.tournament.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Tournament Service - manages tournaments, registrations, and leaderboards
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class TournamentServiceGrpc {

  private TournamentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.tournament.v1.TournamentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CreateTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getCreateTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateTournament",
      requestType = com.game_engine.tournament.v1.CreateTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.TournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CreateTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getCreateTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CreateTournamentRequest, com.game_engine.tournament.v1.TournamentResponse> getCreateTournamentMethod;
    if ((getCreateTournamentMethod = TournamentServiceGrpc.getCreateTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getCreateTournamentMethod = TournamentServiceGrpc.getCreateTournamentMethod) == null) {
          TournamentServiceGrpc.getCreateTournamentMethod = getCreateTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.CreateTournamentRequest, com.game_engine.tournament.v1.TournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.CreateTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.TournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("CreateTournament"))
              .build();
        }
      }
    }
    return getCreateTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getGetTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTournament",
      requestType = com.game_engine.tournament.v1.GetTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.TournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getGetTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentRequest, com.game_engine.tournament.v1.TournamentResponse> getGetTournamentMethod;
    if ((getGetTournamentMethod = TournamentServiceGrpc.getGetTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetTournamentMethod = TournamentServiceGrpc.getGetTournamentMethod) == null) {
          TournamentServiceGrpc.getGetTournamentMethod = getGetTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetTournamentRequest, com.game_engine.tournament.v1.TournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.TournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetTournament"))
              .build();
        }
      }
    }
    return getGetTournamentMethod;
  }

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

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getUpdateTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateTournament",
      requestType = com.game_engine.tournament.v1.UpdateTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.TournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateTournamentRequest,
      com.game_engine.tournament.v1.TournamentResponse> getUpdateTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdateTournamentRequest, com.game_engine.tournament.v1.TournamentResponse> getUpdateTournamentMethod;
    if ((getUpdateTournamentMethod = TournamentServiceGrpc.getUpdateTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getUpdateTournamentMethod = TournamentServiceGrpc.getUpdateTournamentMethod) == null) {
          TournamentServiceGrpc.getUpdateTournamentMethod = getUpdateTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.UpdateTournamentRequest, com.game_engine.tournament.v1.TournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UpdateTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.TournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("UpdateTournament"))
              .build();
        }
      }
    }
    return getUpdateTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CancelTournamentRequest,
      com.game_engine.tournament.v1.CancelTournamentResponse> getCancelTournamentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CancelTournament",
      requestType = com.game_engine.tournament.v1.CancelTournamentRequest.class,
      responseType = com.game_engine.tournament.v1.CancelTournamentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CancelTournamentRequest,
      com.game_engine.tournament.v1.CancelTournamentResponse> getCancelTournamentMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.CancelTournamentRequest, com.game_engine.tournament.v1.CancelTournamentResponse> getCancelTournamentMethod;
    if ((getCancelTournamentMethod = TournamentServiceGrpc.getCancelTournamentMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getCancelTournamentMethod = TournamentServiceGrpc.getCancelTournamentMethod) == null) {
          TournamentServiceGrpc.getCancelTournamentMethod = getCancelTournamentMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.CancelTournamentRequest, com.game_engine.tournament.v1.CancelTournamentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CancelTournament"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.CancelTournamentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.CancelTournamentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("CancelTournament"))
              .build();
        }
      }
    }
    return getCancelTournamentMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.RegisterUserRequest,
      com.game_engine.tournament.v1.RegisterUserResponse> getRegisterUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RegisterUser",
      requestType = com.game_engine.tournament.v1.RegisterUserRequest.class,
      responseType = com.game_engine.tournament.v1.RegisterUserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.RegisterUserRequest,
      com.game_engine.tournament.v1.RegisterUserResponse> getRegisterUserMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.RegisterUserRequest, com.game_engine.tournament.v1.RegisterUserResponse> getRegisterUserMethod;
    if ((getRegisterUserMethod = TournamentServiceGrpc.getRegisterUserMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getRegisterUserMethod = TournamentServiceGrpc.getRegisterUserMethod) == null) {
          TournamentServiceGrpc.getRegisterUserMethod = getRegisterUserMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.RegisterUserRequest, com.game_engine.tournament.v1.RegisterUserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RegisterUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.RegisterUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.RegisterUserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("RegisterUser"))
              .build();
        }
      }
    }
    return getRegisterUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UnregisterUserRequest,
      com.game_engine.tournament.v1.UnregisterUserResponse> getUnregisterUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UnregisterUser",
      requestType = com.game_engine.tournament.v1.UnregisterUserRequest.class,
      responseType = com.game_engine.tournament.v1.UnregisterUserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UnregisterUserRequest,
      com.game_engine.tournament.v1.UnregisterUserResponse> getUnregisterUserMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UnregisterUserRequest, com.game_engine.tournament.v1.UnregisterUserResponse> getUnregisterUserMethod;
    if ((getUnregisterUserMethod = TournamentServiceGrpc.getUnregisterUserMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getUnregisterUserMethod = TournamentServiceGrpc.getUnregisterUserMethod) == null) {
          TournamentServiceGrpc.getUnregisterUserMethod = getUnregisterUserMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.UnregisterUserRequest, com.game_engine.tournament.v1.UnregisterUserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UnregisterUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UnregisterUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UnregisterUserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("UnregisterUser"))
              .build();
        }
      }
    }
    return getUnregisterUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentUsersRequest,
      com.game_engine.tournament.v1.GetTournamentUsersResponse> getGetTournamentUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTournamentUsers",
      requestType = com.game_engine.tournament.v1.GetTournamentUsersRequest.class,
      responseType = com.game_engine.tournament.v1.GetTournamentUsersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentUsersRequest,
      com.game_engine.tournament.v1.GetTournamentUsersResponse> getGetTournamentUsersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetTournamentUsersRequest, com.game_engine.tournament.v1.GetTournamentUsersResponse> getGetTournamentUsersMethod;
    if ((getGetTournamentUsersMethod = TournamentServiceGrpc.getGetTournamentUsersMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetTournamentUsersMethod = TournamentServiceGrpc.getGetTournamentUsersMethod) == null) {
          TournamentServiceGrpc.getGetTournamentUsersMethod = getGetTournamentUsersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetTournamentUsersRequest, com.game_engine.tournament.v1.GetTournamentUsersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTournamentUsers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetTournamentUsersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetTournamentUsersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetTournamentUsers"))
              .build();
        }
      }
    }
    return getGetTournamentUsersMethod;
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

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdatePlayerScoreRequest,
      com.game_engine.tournament.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdatePlayerScore",
      requestType = com.game_engine.tournament.v1.UpdatePlayerScoreRequest.class,
      responseType = com.game_engine.tournament.v1.UpdatePlayerScoreResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdatePlayerScoreRequest,
      com.game_engine.tournament.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.UpdatePlayerScoreRequest, com.game_engine.tournament.v1.UpdatePlayerScoreResponse> getUpdatePlayerScoreMethod;
    if ((getUpdatePlayerScoreMethod = TournamentServiceGrpc.getUpdatePlayerScoreMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getUpdatePlayerScoreMethod = TournamentServiceGrpc.getUpdatePlayerScoreMethod) == null) {
          TournamentServiceGrpc.getUpdatePlayerScoreMethod = getUpdatePlayerScoreMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.UpdatePlayerScoreRequest, com.game_engine.tournament.v1.UpdatePlayerScoreResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdatePlayerScore"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UpdatePlayerScoreRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.UpdatePlayerScoreResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("UpdatePlayerScore"))
              .build();
        }
      }
    }
    return getUpdatePlayerScoreMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetUserTournamentsRequest,
      com.game_engine.tournament.v1.GetUserTournamentsResponse> getGetUserTournamentsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserTournaments",
      requestType = com.game_engine.tournament.v1.GetUserTournamentsRequest.class,
      responseType = com.game_engine.tournament.v1.GetUserTournamentsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetUserTournamentsRequest,
      com.game_engine.tournament.v1.GetUserTournamentsResponse> getGetUserTournamentsMethod() {
    io.grpc.MethodDescriptor<com.game_engine.tournament.v1.GetUserTournamentsRequest, com.game_engine.tournament.v1.GetUserTournamentsResponse> getGetUserTournamentsMethod;
    if ((getGetUserTournamentsMethod = TournamentServiceGrpc.getGetUserTournamentsMethod) == null) {
      synchronized (TournamentServiceGrpc.class) {
        if ((getGetUserTournamentsMethod = TournamentServiceGrpc.getGetUserTournamentsMethod) == null) {
          TournamentServiceGrpc.getGetUserTournamentsMethod = getGetUserTournamentsMethod =
              io.grpc.MethodDescriptor.<com.game_engine.tournament.v1.GetUserTournamentsRequest, com.game_engine.tournament.v1.GetUserTournamentsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserTournaments"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetUserTournamentsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.tournament.v1.GetUserTournamentsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new TournamentServiceMethodDescriptorSupplier("GetUserTournaments"))
              .build();
        }
      }
    }
    return getGetUserTournamentsMethod;
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
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Tournament management
     * </pre>
     */
    default void createTournament(com.game_engine.tournament.v1.CreateTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateTournamentMethod(), responseObserver);
    }

    /**
     */
    default void getTournament(com.game_engine.tournament.v1.GetTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTournamentMethod(), responseObserver);
    }

    /**
     */
    default void listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListTournamentsMethod(), responseObserver);
    }

    /**
     */
    default void updateTournament(com.game_engine.tournament.v1.UpdateTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateTournamentMethod(), responseObserver);
    }

    /**
     */
    default void cancelTournament(com.game_engine.tournament.v1.CancelTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.CancelTournamentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCancelTournamentMethod(), responseObserver);
    }

    /**
     * <pre>
     * Registration
     * </pre>
     */
    default void registerUser(com.game_engine.tournament.v1.RegisterUserRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.RegisterUserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRegisterUserMethod(), responseObserver);
    }

    /**
     */
    default void unregisterUser(com.game_engine.tournament.v1.UnregisterUserRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UnregisterUserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUnregisterUserMethod(), responseObserver);
    }

    /**
     */
    default void getTournamentUsers(com.game_engine.tournament.v1.GetTournamentUsersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentUsersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTournamentUsersMethod(), responseObserver);
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
    default void updatePlayerScore(com.game_engine.tournament.v1.UpdatePlayerScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdatePlayerScoreResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdatePlayerScoreMethod(), responseObserver);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    default void getUserTournaments(com.game_engine.tournament.v1.GetUserTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetUserTournamentsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserTournamentsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service TournamentService.
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
   */
  public static abstract class TournamentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return TournamentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service TournamentService.
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
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
     * Tournament management
     * </pre>
     */
    public void createTournament(com.game_engine.tournament.v1.CreateTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTournament(com.game_engine.tournament.v1.GetTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListTournamentsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateTournament(com.game_engine.tournament.v1.UpdateTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void cancelTournament(com.game_engine.tournament.v1.CancelTournamentRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.CancelTournamentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCancelTournamentMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Registration
     * </pre>
     */
    public void registerUser(com.game_engine.tournament.v1.RegisterUserRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.RegisterUserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRegisterUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void unregisterUser(com.game_engine.tournament.v1.UnregisterUserRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UnregisterUserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUnregisterUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getTournamentUsers(com.game_engine.tournament.v1.GetTournamentUsersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentUsersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTournamentUsersMethod(), getCallOptions()), request, responseObserver);
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
    public void updatePlayerScore(com.game_engine.tournament.v1.UpdatePlayerScoreRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdatePlayerScoreResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdatePlayerScoreMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public void getUserTournaments(com.game_engine.tournament.v1.GetUserTournamentsRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetUserTournamentsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserTournamentsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service TournamentService.
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
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
     * Tournament management
     * </pre>
     */
    public com.game_engine.tournament.v1.TournamentResponse createTournament(com.game_engine.tournament.v1.CreateTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCreateTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.TournamentResponse getTournament(com.game_engine.tournament.v1.GetTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.ListTournamentsResponse listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListTournamentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.TournamentResponse updateTournament(com.game_engine.tournament.v1.UpdateTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdateTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.CancelTournamentResponse cancelTournament(com.game_engine.tournament.v1.CancelTournamentRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getCancelTournamentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Registration
     * </pre>
     */
    public com.game_engine.tournament.v1.RegisterUserResponse registerUser(com.game_engine.tournament.v1.RegisterUserRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getRegisterUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.UnregisterUserResponse unregisterUser(com.game_engine.tournament.v1.UnregisterUserRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUnregisterUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.GetTournamentUsersResponse getTournamentUsers(com.game_engine.tournament.v1.GetTournamentUsersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetTournamentUsersMethod(), getCallOptions(), request);
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
    public com.game_engine.tournament.v1.UpdatePlayerScoreResponse updatePlayerScore(com.game_engine.tournament.v1.UpdatePlayerScoreRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getUpdatePlayerScoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.game_engine.tournament.v1.GetUserTournamentsResponse getUserTournaments(com.game_engine.tournament.v1.GetUserTournamentsRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetUserTournamentsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service TournamentService.
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
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
     * Tournament management
     * </pre>
     */
    public com.game_engine.tournament.v1.TournamentResponse createTournament(com.game_engine.tournament.v1.CreateTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.TournamentResponse getTournament(com.game_engine.tournament.v1.GetTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.ListTournamentsResponse listTournaments(com.game_engine.tournament.v1.ListTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListTournamentsMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.TournamentResponse updateTournament(com.game_engine.tournament.v1.UpdateTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateTournamentMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.CancelTournamentResponse cancelTournament(com.game_engine.tournament.v1.CancelTournamentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCancelTournamentMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Registration
     * </pre>
     */
    public com.game_engine.tournament.v1.RegisterUserResponse registerUser(com.game_engine.tournament.v1.RegisterUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRegisterUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.UnregisterUserResponse unregisterUser(com.game_engine.tournament.v1.UnregisterUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUnregisterUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.tournament.v1.GetTournamentUsersResponse getTournamentUsers(com.game_engine.tournament.v1.GetTournamentUsersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTournamentUsersMethod(), getCallOptions(), request);
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
    public com.game_engine.tournament.v1.UpdatePlayerScoreResponse updatePlayerScore(com.game_engine.tournament.v1.UpdatePlayerScoreRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdatePlayerScoreMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.game_engine.tournament.v1.GetUserTournamentsResponse getUserTournaments(com.game_engine.tournament.v1.GetUserTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserTournamentsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service TournamentService.
   * <pre>
   * Tournament Service - manages tournaments, registrations, and leaderboards
   * </pre>
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
     * Tournament management
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.TournamentResponse> createTournament(
        com.game_engine.tournament.v1.CreateTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateTournamentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.TournamentResponse> getTournament(
        com.game_engine.tournament.v1.GetTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTournamentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.ListTournamentsResponse> listTournaments(
        com.game_engine.tournament.v1.ListTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListTournamentsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.TournamentResponse> updateTournament(
        com.game_engine.tournament.v1.UpdateTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateTournamentMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.CancelTournamentResponse> cancelTournament(
        com.game_engine.tournament.v1.CancelTournamentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCancelTournamentMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Registration
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.RegisterUserResponse> registerUser(
        com.game_engine.tournament.v1.RegisterUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRegisterUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.UnregisterUserResponse> unregisterUser(
        com.game_engine.tournament.v1.UnregisterUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUnregisterUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.GetTournamentUsersResponse> getTournamentUsers(
        com.game_engine.tournament.v1.GetTournamentUsersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTournamentUsersMethod(), getCallOptions()), request);
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
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.UpdatePlayerScoreResponse> updatePlayerScore(
        com.game_engine.tournament.v1.UpdatePlayerScoreRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdatePlayerScoreMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * User tournaments
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.tournament.v1.GetUserTournamentsResponse> getUserTournaments(
        com.game_engine.tournament.v1.GetUserTournamentsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserTournamentsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_TOURNAMENT = 0;
  private static final int METHODID_GET_TOURNAMENT = 1;
  private static final int METHODID_LIST_TOURNAMENTS = 2;
  private static final int METHODID_UPDATE_TOURNAMENT = 3;
  private static final int METHODID_CANCEL_TOURNAMENT = 4;
  private static final int METHODID_REGISTER_USER = 5;
  private static final int METHODID_UNREGISTER_USER = 6;
  private static final int METHODID_GET_TOURNAMENT_USERS = 7;
  private static final int METHODID_GET_LEADERBOARD = 8;
  private static final int METHODID_UPDATE_PLAYER_SCORE = 9;
  private static final int METHODID_GET_USER_TOURNAMENTS = 10;

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
        case METHODID_CREATE_TOURNAMENT:
          serviceImpl.createTournament((com.game_engine.tournament.v1.CreateTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse>) responseObserver);
          break;
        case METHODID_GET_TOURNAMENT:
          serviceImpl.getTournament((com.game_engine.tournament.v1.GetTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse>) responseObserver);
          break;
        case METHODID_LIST_TOURNAMENTS:
          serviceImpl.listTournaments((com.game_engine.tournament.v1.ListTournamentsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.ListTournamentsResponse>) responseObserver);
          break;
        case METHODID_UPDATE_TOURNAMENT:
          serviceImpl.updateTournament((com.game_engine.tournament.v1.UpdateTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.TournamentResponse>) responseObserver);
          break;
        case METHODID_CANCEL_TOURNAMENT:
          serviceImpl.cancelTournament((com.game_engine.tournament.v1.CancelTournamentRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.CancelTournamentResponse>) responseObserver);
          break;
        case METHODID_REGISTER_USER:
          serviceImpl.registerUser((com.game_engine.tournament.v1.RegisterUserRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.RegisterUserResponse>) responseObserver);
          break;
        case METHODID_UNREGISTER_USER:
          serviceImpl.unregisterUser((com.game_engine.tournament.v1.UnregisterUserRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UnregisterUserResponse>) responseObserver);
          break;
        case METHODID_GET_TOURNAMENT_USERS:
          serviceImpl.getTournamentUsers((com.game_engine.tournament.v1.GetTournamentUsersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetTournamentUsersResponse>) responseObserver);
          break;
        case METHODID_GET_LEADERBOARD:
          serviceImpl.getLeaderboard((com.game_engine.tournament.v1.GetLeaderboardRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetLeaderboardResponse>) responseObserver);
          break;
        case METHODID_UPDATE_PLAYER_SCORE:
          serviceImpl.updatePlayerScore((com.game_engine.tournament.v1.UpdatePlayerScoreRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.UpdatePlayerScoreResponse>) responseObserver);
          break;
        case METHODID_GET_USER_TOURNAMENTS:
          serviceImpl.getUserTournaments((com.game_engine.tournament.v1.GetUserTournamentsRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.tournament.v1.GetUserTournamentsResponse>) responseObserver);
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
          getCreateTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.CreateTournamentRequest,
              com.game_engine.tournament.v1.TournamentResponse>(
                service, METHODID_CREATE_TOURNAMENT)))
        .addMethod(
          getGetTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetTournamentRequest,
              com.game_engine.tournament.v1.TournamentResponse>(
                service, METHODID_GET_TOURNAMENT)))
        .addMethod(
          getListTournamentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.ListTournamentsRequest,
              com.game_engine.tournament.v1.ListTournamentsResponse>(
                service, METHODID_LIST_TOURNAMENTS)))
        .addMethod(
          getUpdateTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.UpdateTournamentRequest,
              com.game_engine.tournament.v1.TournamentResponse>(
                service, METHODID_UPDATE_TOURNAMENT)))
        .addMethod(
          getCancelTournamentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.CancelTournamentRequest,
              com.game_engine.tournament.v1.CancelTournamentResponse>(
                service, METHODID_CANCEL_TOURNAMENT)))
        .addMethod(
          getRegisterUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.RegisterUserRequest,
              com.game_engine.tournament.v1.RegisterUserResponse>(
                service, METHODID_REGISTER_USER)))
        .addMethod(
          getUnregisterUserMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.UnregisterUserRequest,
              com.game_engine.tournament.v1.UnregisterUserResponse>(
                service, METHODID_UNREGISTER_USER)))
        .addMethod(
          getGetTournamentUsersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetTournamentUsersRequest,
              com.game_engine.tournament.v1.GetTournamentUsersResponse>(
                service, METHODID_GET_TOURNAMENT_USERS)))
        .addMethod(
          getGetLeaderboardMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetLeaderboardRequest,
              com.game_engine.tournament.v1.GetLeaderboardResponse>(
                service, METHODID_GET_LEADERBOARD)))
        .addMethod(
          getUpdatePlayerScoreMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.UpdatePlayerScoreRequest,
              com.game_engine.tournament.v1.UpdatePlayerScoreResponse>(
                service, METHODID_UPDATE_PLAYER_SCORE)))
        .addMethod(
          getGetUserTournamentsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.tournament.v1.GetUserTournamentsRequest,
              com.game_engine.tournament.v1.GetUserTournamentsResponse>(
                service, METHODID_GET_USER_TOURNAMENTS)))
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
              .addMethod(getCreateTournamentMethod())
              .addMethod(getGetTournamentMethod())
              .addMethod(getListTournamentsMethod())
              .addMethod(getUpdateTournamentMethod())
              .addMethod(getCancelTournamentMethod())
              .addMethod(getRegisterUserMethod())
              .addMethod(getUnregisterUserMethod())
              .addMethod(getGetTournamentUsersMethod())
              .addMethod(getGetLeaderboardMethod())
              .addMethod(getUpdatePlayerScoreMethod())
              .addMethod(getGetUserTournamentsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
