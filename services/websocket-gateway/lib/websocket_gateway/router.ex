defmodule WebsocketGateway.Router do
  use WebsocketGateway, :router

  # Health check pipeline (no auth required)
  pipeline :health do
    plug :accepts, ["json"]
    plug WebsocketGateway.Plugs.HealthCheck
  end

  # API pipeline
  pipeline :api do
    plug :accepts, ["json"]
    plug WebsocketGateway.Plugs.RateLimiter
    plug WebsocketGateway.Plugs.RequestLogger
  end

  # WebSocket pipeline (handled by UserSocket)
  # No traditional HTTP routing for WebSocket connections

  # Health check endpoints
  scope "/health", WebsocketGateway do
    pipe_through :health

    get "/", HealthController, :index
    get "/live", HealthController, :live
    get "/ready", HealthController, :ready
  end

  # API routes for metrics and admin
  scope "/api", WebsocketGateway do
    pipe_through :api

    # Metrics
    get "/metrics", MetricsController, :index

    # Room management
    get "/rooms", RoomController, :list
    get "/rooms/:room_id", RoomController, :show
    delete "/rooms/:room_id", RoomController, :delete

    # Presence
    get "/presence", PresenceController, :index

    # Admin
    get "/stats", AdminController, :stats
  end
end
