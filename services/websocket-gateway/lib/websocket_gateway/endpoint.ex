defmodule WebsocketGateway.Endpoint do
  use Phoenix.Endpoint, otp_app: :websocket_gateway

  # Socket handler
  socket "/socket", WebsocketGateway.UserSocket,
    websocket: true,
    longpoll: false

  # Health check endpoint (no auth required)
  plug Plug.HealthCheck

  # Health check for liveness
  plug WebsocketGateway.Plugs.HealthCheck

  # Rate limiting
  plug WebsocketGateway.Plugs.RateLimiter

  # Request logging
  plug WebsocketGateway.Plugs.RequestLogger

  # Metrics endpoint
  if code_reloading? do
    plug Phoenix.CodeReloader
    plug Phoenix.Ecto.CheckRepoStatus, otp_app: :websocket_gateway
  end

  plug Phoenix.LiveDashboard.RequestLogger,
    param_key: "request_logger",
    cookie_key: "request_logger"

  plug Plug.Static,
    at: "/",
    from: :websocket_gateway,
    gzip: false,
    only: ~w(css fonts images js favicon.ico robots.txt)

  # Code reloading
  plug Phoenix.CodeReloader

  plug Phoenix.TokenVerifier,
    endpoint: __MODULE__,
    token_salt: "websocket_gateway",
    max_age: 60 * 60 * 24 * 7 # 7 days

  plug Plug.Parsers,
    parsers: [:urlencoded, :multipart, :json],
    pass: ["*/*"],
    json_decoder: Phoenix.json_library()

  plug Plug.MethodOverride
  plug Plug.Head
  plug Plug.Session,
    store: :cookie,
    key: "_websocket_gateway_key",
    signing_salt: "websocket_gateway",
    same_site: "Lax"

  plug WebsocketGateway.Router
end
