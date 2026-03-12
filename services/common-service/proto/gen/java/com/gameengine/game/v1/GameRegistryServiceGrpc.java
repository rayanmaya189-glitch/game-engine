package com.game_engine.game.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Game Registry Service - manages game catalog, categories, and game configurations
 * </pre>
 */
@io.grpc.stub.annotations.GrpcGenerated
public final class GameRegistryServiceGrpc {

  private GameRegistryServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "game_engine.game.v1.GameRegistryService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.ListGamesRequest,
      com.game_engine.game.v1.ListGamesResponse> getListGamesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListGames",
      requestType = com.game_engine.game.v1.ListGamesRequest.class,
      responseType = com.game_engine.game.v1.ListGamesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.ListGamesRequest,
      com.game_engine.game.v1.ListGamesResponse> getListGamesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.ListGamesRequest, com.game_engine.game.v1.ListGamesResponse> getListGamesMethod;
    if ((getListGamesMethod = GameRegistryServiceGrpc.getListGamesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getListGamesMethod = GameRegistryServiceGrpc.getListGamesMethod) == null) {
          GameRegistryServiceGrpc.getListGamesMethod = getListGamesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.ListGamesRequest, com.game_engine.game.v1.ListGamesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListGames"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.ListGamesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.ListGamesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("ListGames"))
              .build();
        }
      }
    }
    return getListGamesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameRequest,
      com.game_engine.game.v1.GetGameResponse> getGetGameMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetGame",
      requestType = com.game_engine.game.v1.GetGameRequest.class,
      responseType = com.game_engine.game.v1.GetGameResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameRequest,
      com.game_engine.game.v1.GetGameResponse> getGetGameMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameRequest, com.game_engine.game.v1.GetGameResponse> getGetGameMethod;
    if ((getGetGameMethod = GameRegistryServiceGrpc.getGetGameMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetGameMethod = GameRegistryServiceGrpc.getGetGameMethod) == null) {
          GameRegistryServiceGrpc.getGetGameMethod = getGetGameMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetGameRequest, com.game_engine.game.v1.GetGameResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetGame"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetGame"))
              .build();
        }
      }
    }
    return getGetGameMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameConfigRequest,
      com.game_engine.game.v1.GetGameConfigResponse> getGetGameConfigMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetGameConfig",
      requestType = com.game_engine.game.v1.GetGameConfigRequest.class,
      responseType = com.game_engine.game.v1.GetGameConfigResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameConfigRequest,
      com.game_engine.game.v1.GetGameConfigResponse> getGetGameConfigMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameConfigRequest, com.game_engine.game.v1.GetGameConfigResponse> getGetGameConfigMethod;
    if ((getGetGameConfigMethod = GameRegistryServiceGrpc.getGetGameConfigMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetGameConfigMethod = GameRegistryServiceGrpc.getGetGameConfigMethod) == null) {
          GameRegistryServiceGrpc.getGetGameConfigMethod = getGetGameConfigMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetGameConfigRequest, com.game_engine.game.v1.GetGameConfigResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetGameConfig"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameConfigRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameConfigResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetGameConfig"))
              .build();
        }
      }
    }
    return getGetGameConfigMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameURLRequest,
      com.game_engine.game.v1.GetGameURLResponse> getGetGameURLMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetGameURL",
      requestType = com.game_engine.game.v1.GetGameURLRequest.class,
      responseType = com.game_engine.game.v1.GetGameURLResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameURLRequest,
      com.game_engine.game.v1.GetGameURLResponse> getGetGameURLMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetGameURLRequest, com.game_engine.game.v1.GetGameURLResponse> getGetGameURLMethod;
    if ((getGetGameURLMethod = GameRegistryServiceGrpc.getGetGameURLMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetGameURLMethod = GameRegistryServiceGrpc.getGetGameURLMethod) == null) {
          GameRegistryServiceGrpc.getGetGameURLMethod = getGetGameURLMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetGameURLRequest, com.game_engine.game.v1.GetGameURLResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetGameURL"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameURLRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetGameURLResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetGameURL"))
              .build();
        }
      }
    }
    return getGetGameURLMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetCategoriesRequest,
      com.game_engine.game.v1.GetCategoriesResponse> getGetCategoriesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCategories",
      requestType = com.game_engine.game.v1.GetCategoriesRequest.class,
      responseType = com.game_engine.game.v1.GetCategoriesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetCategoriesRequest,
      com.game_engine.game.v1.GetCategoriesResponse> getGetCategoriesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetCategoriesRequest, com.game_engine.game.v1.GetCategoriesResponse> getGetCategoriesMethod;
    if ((getGetCategoriesMethod = GameRegistryServiceGrpc.getGetCategoriesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetCategoriesMethod = GameRegistryServiceGrpc.getGetCategoriesMethod) == null) {
          GameRegistryServiceGrpc.getGetCategoriesMethod = getGetCategoriesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetCategoriesRequest, com.game_engine.game.v1.GetCategoriesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCategories"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetCategoriesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetCategoriesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetCategories"))
              .build();
        }
      }
    }
    return getGetCategoriesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetProvidersRequest,
      com.game_engine.game.v1.GetProvidersResponse> getGetProvidersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetProviders",
      requestType = com.game_engine.game.v1.GetProvidersRequest.class,
      responseType = com.game_engine.game.v1.GetProvidersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetProvidersRequest,
      com.game_engine.game.v1.GetProvidersResponse> getGetProvidersMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetProvidersRequest, com.game_engine.game.v1.GetProvidersResponse> getGetProvidersMethod;
    if ((getGetProvidersMethod = GameRegistryServiceGrpc.getGetProvidersMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetProvidersMethod = GameRegistryServiceGrpc.getGetProvidersMethod) == null) {
          GameRegistryServiceGrpc.getGetProvidersMethod = getGetProvidersMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetProvidersRequest, com.game_engine.game.v1.GetProvidersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetProviders"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetProvidersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetProvidersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetProviders"))
              .build();
        }
      }
    }
    return getGetProvidersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.SearchGamesRequest,
      com.game_engine.game.v1.SearchGamesResponse> getSearchGamesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SearchGames",
      requestType = com.game_engine.game.v1.SearchGamesRequest.class,
      responseType = com.game_engine.game.v1.SearchGamesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.SearchGamesRequest,
      com.game_engine.game.v1.SearchGamesResponse> getSearchGamesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.SearchGamesRequest, com.game_engine.game.v1.SearchGamesResponse> getSearchGamesMethod;
    if ((getSearchGamesMethod = GameRegistryServiceGrpc.getSearchGamesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getSearchGamesMethod = GameRegistryServiceGrpc.getSearchGamesMethod) == null) {
          GameRegistryServiceGrpc.getSearchGamesMethod = getSearchGamesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.SearchGamesRequest, com.game_engine.game.v1.SearchGamesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SearchGames"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.SearchGamesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.SearchGamesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("SearchGames"))
              .build();
        }
      }
    }
    return getSearchGamesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetFeaturedGamesRequest,
      com.game_engine.game.v1.GetFeaturedGamesResponse> getGetFeaturedGamesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetFeaturedGames",
      requestType = com.game_engine.game.v1.GetFeaturedGamesRequest.class,
      responseType = com.game_engine.game.v1.GetFeaturedGamesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetFeaturedGamesRequest,
      com.game_engine.game.v1.GetFeaturedGamesResponse> getGetFeaturedGamesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetFeaturedGamesRequest, com.game_engine.game.v1.GetFeaturedGamesResponse> getGetFeaturedGamesMethod;
    if ((getGetFeaturedGamesMethod = GameRegistryServiceGrpc.getGetFeaturedGamesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetFeaturedGamesMethod = GameRegistryServiceGrpc.getGetFeaturedGamesMethod) == null) {
          GameRegistryServiceGrpc.getGetFeaturedGamesMethod = getGetFeaturedGamesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetFeaturedGamesRequest, com.game_engine.game.v1.GetFeaturedGamesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetFeaturedGames"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetFeaturedGamesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetFeaturedGamesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetFeaturedGames"))
              .build();
        }
      }
    }
    return getGetFeaturedGamesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetPopularGamesRequest,
      com.game_engine.game.v1.GetPopularGamesResponse> getGetPopularGamesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPopularGames",
      requestType = com.game_engine.game.v1.GetPopularGamesRequest.class,
      responseType = com.game_engine.game.v1.GetPopularGamesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetPopularGamesRequest,
      com.game_engine.game.v1.GetPopularGamesResponse> getGetPopularGamesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetPopularGamesRequest, com.game_engine.game.v1.GetPopularGamesResponse> getGetPopularGamesMethod;
    if ((getGetPopularGamesMethod = GameRegistryServiceGrpc.getGetPopularGamesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetPopularGamesMethod = GameRegistryServiceGrpc.getGetPopularGamesMethod) == null) {
          GameRegistryServiceGrpc.getGetPopularGamesMethod = getGetPopularGamesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetPopularGamesRequest, com.game_engine.game.v1.GetPopularGamesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPopularGames"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetPopularGamesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetPopularGamesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetPopularGames"))
              .build();
        }
      }
    }
    return getGetPopularGamesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.game_engine.game.v1.GetNewGamesRequest,
      com.game_engine.game.v1.GetNewGamesResponse> getGetNewGamesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetNewGames",
      requestType = com.game_engine.game.v1.GetNewGamesRequest.class,
      responseType = com.game_engine.game.v1.GetNewGamesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.game_engine.game.v1.GetNewGamesRequest,
      com.game_engine.game.v1.GetNewGamesResponse> getGetNewGamesMethod() {
    io.grpc.MethodDescriptor<com.game_engine.game.v1.GetNewGamesRequest, com.game_engine.game.v1.GetNewGamesResponse> getGetNewGamesMethod;
    if ((getGetNewGamesMethod = GameRegistryServiceGrpc.getGetNewGamesMethod) == null) {
      synchronized (GameRegistryServiceGrpc.class) {
        if ((getGetNewGamesMethod = GameRegistryServiceGrpc.getGetNewGamesMethod) == null) {
          GameRegistryServiceGrpc.getGetNewGamesMethod = getGetNewGamesMethod =
              io.grpc.MethodDescriptor.<com.game_engine.game.v1.GetNewGamesRequest, com.game_engine.game.v1.GetNewGamesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetNewGames"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetNewGamesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.game_engine.game.v1.GetNewGamesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new GameRegistryServiceMethodDescriptorSupplier("GetNewGames"))
              .build();
        }
      }
    }
    return getGetNewGamesMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static GameRegistryServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceStub>() {
        @java.lang.Override
        public GameRegistryServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GameRegistryServiceStub(channel, callOptions);
        }
      };
    return GameRegistryServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static GameRegistryServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceBlockingV2Stub>() {
        @java.lang.Override
        public GameRegistryServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GameRegistryServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return GameRegistryServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static GameRegistryServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceBlockingStub>() {
        @java.lang.Override
        public GameRegistryServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GameRegistryServiceBlockingStub(channel, callOptions);
        }
      };
    return GameRegistryServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static GameRegistryServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<GameRegistryServiceFutureStub>() {
        @java.lang.Override
        public GameRegistryServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new GameRegistryServiceFutureStub(channel, callOptions);
        }
      };
    return GameRegistryServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public interface AsyncService {

    /**
     */
    default void listGames(com.game_engine.game.v1.ListGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.ListGamesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListGamesMethod(), responseObserver);
    }

    /**
     */
    default void getGame(com.game_engine.game.v1.GetGameRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetGameMethod(), responseObserver);
    }

    /**
     */
    default void getGameConfig(com.game_engine.game.v1.GetGameConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameConfigResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetGameConfigMethod(), responseObserver);
    }

    /**
     */
    default void getGameURL(com.game_engine.game.v1.GetGameURLRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameURLResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetGameURLMethod(), responseObserver);
    }

    /**
     */
    default void getCategories(com.game_engine.game.v1.GetCategoriesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetCategoriesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCategoriesMethod(), responseObserver);
    }

    /**
     */
    default void getProviders(com.game_engine.game.v1.GetProvidersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetProvidersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetProvidersMethod(), responseObserver);
    }

    /**
     */
    default void searchGames(com.game_engine.game.v1.SearchGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.SearchGamesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSearchGamesMethod(), responseObserver);
    }

    /**
     */
    default void getFeaturedGames(com.game_engine.game.v1.GetFeaturedGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetFeaturedGamesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetFeaturedGamesMethod(), responseObserver);
    }

    /**
     */
    default void getPopularGames(com.game_engine.game.v1.GetPopularGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetPopularGamesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPopularGamesMethod(), responseObserver);
    }

    /**
     */
    default void getNewGames(com.game_engine.game.v1.GetNewGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetNewGamesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetNewGamesMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service GameRegistryService.
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public static abstract class GameRegistryServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return GameRegistryServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service GameRegistryService.
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public static final class GameRegistryServiceStub
      extends io.grpc.stub.AbstractAsyncStub<GameRegistryServiceStub> {
    private GameRegistryServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GameRegistryServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GameRegistryServiceStub(channel, callOptions);
    }

    /**
     */
    public void listGames(com.game_engine.game.v1.ListGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.ListGamesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListGamesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getGame(com.game_engine.game.v1.GetGameRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetGameMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getGameConfig(com.game_engine.game.v1.GetGameConfigRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameConfigResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetGameConfigMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getGameURL(com.game_engine.game.v1.GetGameURLRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameURLResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetGameURLMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getCategories(com.game_engine.game.v1.GetCategoriesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetCategoriesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCategoriesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getProviders(com.game_engine.game.v1.GetProvidersRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetProvidersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetProvidersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void searchGames(com.game_engine.game.v1.SearchGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.SearchGamesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSearchGamesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getFeaturedGames(com.game_engine.game.v1.GetFeaturedGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetFeaturedGamesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetFeaturedGamesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getPopularGames(com.game_engine.game.v1.GetPopularGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetPopularGamesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPopularGamesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getNewGames(com.game_engine.game.v1.GetNewGamesRequest request,
        io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetNewGamesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetNewGamesMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service GameRegistryService.
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public static final class GameRegistryServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<GameRegistryServiceBlockingV2Stub> {
    private GameRegistryServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GameRegistryServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GameRegistryServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.game.v1.ListGamesResponse listGames(com.game_engine.game.v1.ListGamesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getListGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameResponse getGame(com.game_engine.game.v1.GetGameRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetGameMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameConfigResponse getGameConfig(com.game_engine.game.v1.GetGameConfigRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetGameConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameURLResponse getGameURL(com.game_engine.game.v1.GetGameURLRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetGameURLMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetCategoriesResponse getCategories(com.game_engine.game.v1.GetCategoriesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetCategoriesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetProvidersResponse getProviders(com.game_engine.game.v1.GetProvidersRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetProvidersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.SearchGamesResponse searchGames(com.game_engine.game.v1.SearchGamesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getSearchGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetFeaturedGamesResponse getFeaturedGames(com.game_engine.game.v1.GetFeaturedGamesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetFeaturedGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetPopularGamesResponse getPopularGames(com.game_engine.game.v1.GetPopularGamesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetPopularGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetNewGamesResponse getNewGames(com.game_engine.game.v1.GetNewGamesRequest request) throws io.grpc.StatusException {
      return io.grpc.stub.ClientCalls.blockingV2UnaryCall(
          getChannel(), getGetNewGamesMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service GameRegistryService.
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public static final class GameRegistryServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<GameRegistryServiceBlockingStub> {
    private GameRegistryServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GameRegistryServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GameRegistryServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.game_engine.game.v1.ListGamesResponse listGames(com.game_engine.game.v1.ListGamesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameResponse getGame(com.game_engine.game.v1.GetGameRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetGameMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameConfigResponse getGameConfig(com.game_engine.game.v1.GetGameConfigRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetGameConfigMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetGameURLResponse getGameURL(com.game_engine.game.v1.GetGameURLRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetGameURLMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetCategoriesResponse getCategories(com.game_engine.game.v1.GetCategoriesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCategoriesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetProvidersResponse getProviders(com.game_engine.game.v1.GetProvidersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetProvidersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.SearchGamesResponse searchGames(com.game_engine.game.v1.SearchGamesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSearchGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetFeaturedGamesResponse getFeaturedGames(com.game_engine.game.v1.GetFeaturedGamesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetFeaturedGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetPopularGamesResponse getPopularGames(com.game_engine.game.v1.GetPopularGamesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPopularGamesMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.game_engine.game.v1.GetNewGamesResponse getNewGames(com.game_engine.game.v1.GetNewGamesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetNewGamesMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service GameRegistryService.
   * <pre>
   * Game Registry Service - manages game catalog, categories, and game configurations
   * </pre>
   */
  public static final class GameRegistryServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<GameRegistryServiceFutureStub> {
    private GameRegistryServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected GameRegistryServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new GameRegistryServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.ListGamesResponse> listGames(
        com.game_engine.game.v1.ListGamesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListGamesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetGameResponse> getGame(
        com.game_engine.game.v1.GetGameRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetGameMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetGameConfigResponse> getGameConfig(
        com.game_engine.game.v1.GetGameConfigRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetGameConfigMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetGameURLResponse> getGameURL(
        com.game_engine.game.v1.GetGameURLRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetGameURLMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetCategoriesResponse> getCategories(
        com.game_engine.game.v1.GetCategoriesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCategoriesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetProvidersResponse> getProviders(
        com.game_engine.game.v1.GetProvidersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetProvidersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.SearchGamesResponse> searchGames(
        com.game_engine.game.v1.SearchGamesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSearchGamesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetFeaturedGamesResponse> getFeaturedGames(
        com.game_engine.game.v1.GetFeaturedGamesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetFeaturedGamesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetPopularGamesResponse> getPopularGames(
        com.game_engine.game.v1.GetPopularGamesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPopularGamesMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.game_engine.game.v1.GetNewGamesResponse> getNewGames(
        com.game_engine.game.v1.GetNewGamesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetNewGamesMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_LIST_GAMES = 0;
  private static final int METHODID_GET_GAME = 1;
  private static final int METHODID_GET_GAME_CONFIG = 2;
  private static final int METHODID_GET_GAME_URL = 3;
  private static final int METHODID_GET_CATEGORIES = 4;
  private static final int METHODID_GET_PROVIDERS = 5;
  private static final int METHODID_SEARCH_GAMES = 6;
  private static final int METHODID_GET_FEATURED_GAMES = 7;
  private static final int METHODID_GET_POPULAR_GAMES = 8;
  private static final int METHODID_GET_NEW_GAMES = 9;

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
        case METHODID_LIST_GAMES:
          serviceImpl.listGames((com.game_engine.game.v1.ListGamesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.ListGamesResponse>) responseObserver);
          break;
        case METHODID_GET_GAME:
          serviceImpl.getGame((com.game_engine.game.v1.GetGameRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameResponse>) responseObserver);
          break;
        case METHODID_GET_GAME_CONFIG:
          serviceImpl.getGameConfig((com.game_engine.game.v1.GetGameConfigRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameConfigResponse>) responseObserver);
          break;
        case METHODID_GET_GAME_URL:
          serviceImpl.getGameURL((com.game_engine.game.v1.GetGameURLRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetGameURLResponse>) responseObserver);
          break;
        case METHODID_GET_CATEGORIES:
          serviceImpl.getCategories((com.game_engine.game.v1.GetCategoriesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetCategoriesResponse>) responseObserver);
          break;
        case METHODID_GET_PROVIDERS:
          serviceImpl.getProviders((com.game_engine.game.v1.GetProvidersRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetProvidersResponse>) responseObserver);
          break;
        case METHODID_SEARCH_GAMES:
          serviceImpl.searchGames((com.game_engine.game.v1.SearchGamesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.SearchGamesResponse>) responseObserver);
          break;
        case METHODID_GET_FEATURED_GAMES:
          serviceImpl.getFeaturedGames((com.game_engine.game.v1.GetFeaturedGamesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetFeaturedGamesResponse>) responseObserver);
          break;
        case METHODID_GET_POPULAR_GAMES:
          serviceImpl.getPopularGames((com.game_engine.game.v1.GetPopularGamesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetPopularGamesResponse>) responseObserver);
          break;
        case METHODID_GET_NEW_GAMES:
          serviceImpl.getNewGames((com.game_engine.game.v1.GetNewGamesRequest) request,
              (io.grpc.stub.StreamObserver<com.game_engine.game.v1.GetNewGamesResponse>) responseObserver);
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
          getListGamesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.ListGamesRequest,
              com.game_engine.game.v1.ListGamesResponse>(
                service, METHODID_LIST_GAMES)))
        .addMethod(
          getGetGameMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetGameRequest,
              com.game_engine.game.v1.GetGameResponse>(
                service, METHODID_GET_GAME)))
        .addMethod(
          getGetGameConfigMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetGameConfigRequest,
              com.game_engine.game.v1.GetGameConfigResponse>(
                service, METHODID_GET_GAME_CONFIG)))
        .addMethod(
          getGetGameURLMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetGameURLRequest,
              com.game_engine.game.v1.GetGameURLResponse>(
                service, METHODID_GET_GAME_URL)))
        .addMethod(
          getGetCategoriesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetCategoriesRequest,
              com.game_engine.game.v1.GetCategoriesResponse>(
                service, METHODID_GET_CATEGORIES)))
        .addMethod(
          getGetProvidersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetProvidersRequest,
              com.game_engine.game.v1.GetProvidersResponse>(
                service, METHODID_GET_PROVIDERS)))
        .addMethod(
          getSearchGamesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.SearchGamesRequest,
              com.game_engine.game.v1.SearchGamesResponse>(
                service, METHODID_SEARCH_GAMES)))
        .addMethod(
          getGetFeaturedGamesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetFeaturedGamesRequest,
              com.game_engine.game.v1.GetFeaturedGamesResponse>(
                service, METHODID_GET_FEATURED_GAMES)))
        .addMethod(
          getGetPopularGamesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetPopularGamesRequest,
              com.game_engine.game.v1.GetPopularGamesResponse>(
                service, METHODID_GET_POPULAR_GAMES)))
        .addMethod(
          getGetNewGamesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.game_engine.game.v1.GetNewGamesRequest,
              com.game_engine.game.v1.GetNewGamesResponse>(
                service, METHODID_GET_NEW_GAMES)))
        .build();
  }

  private static abstract class GameRegistryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    GameRegistryServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.game_engine.game.v1.GameRegistry.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("GameRegistryService");
    }
  }

  private static final class GameRegistryServiceFileDescriptorSupplier
      extends GameRegistryServiceBaseDescriptorSupplier {
    GameRegistryServiceFileDescriptorSupplier() {}
  }

  private static final class GameRegistryServiceMethodDescriptorSupplier
      extends GameRegistryServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    GameRegistryServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (GameRegistryServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new GameRegistryServiceFileDescriptorSupplier())
              .addMethod(getListGamesMethod())
              .addMethod(getGetGameMethod())
              .addMethod(getGetGameConfigMethod())
              .addMethod(getGetGameURLMethod())
              .addMethod(getGetCategoriesMethod())
              .addMethod(getGetProvidersMethod())
              .addMethod(getSearchGamesMethod())
              .addMethod(getGetFeaturedGamesMethod())
              .addMethod(getGetPopularGamesMethod())
              .addMethod(getGetNewGamesMethod())
              .build();
        }
      }
    }
    return result;
  }
}
