import Config

# Production configuration
# Use environment variables for sensitive values

# Database configuration
config :websocket_gateway, WebsocketGateway.Repo,
  pool_size: String.to_integer(System.get_env("DB_POOL_SIZE") || "10"),
  pool_timeout: 5000,
  timeout: 15_000,
  ssl: System.get_env("DB_SSL_ENABLED") == "true",
  ssl_opts: [
    verify: :verify_peer,
    cacertfile: System.get_env("DB_CA_CERTFILE"),
    server_name_indication: System.get_env("DB_SERVER_NAME"),
    customize_hostname_check: [
      match_fun: :public_key.pkix_verify_hostname_match_fun(:https)
    ]
  ]

# Redis configuration
config :websocket_gateway, :redis,
  host: System.get_env("REDIS_HOST"),
  port: String.to_integer(System.get_env("REDIS_PORT") || "6379"),
  password: System.get_env("REDIS_PASSWORD", ""),
  database: String.to_integer(System.get_env("REDIS_DB") || "0"),
  pool_size: String.to_integer(System.get_env("REDIS_POOL_SIZE") || "20"),
  socket_opts: [:binary, active: false, packet: :raw, recbuf: 16384, sndbuf: 16384]

# JWT configuration - MUST be set via environment in production
jwt_secret = System.get_env("JWT_SECRET_KEY") || raise """
JWT_SECRET_KEY environment variable must be set in production
"""

config :websocket_gateway, :jwt,
  secret_key: jwt_secret,
  algorithm: "HS256",
  issuer: "game-engine",
  audience: "game-engine",
  expiration: 86_400

# WebSocket configuration for production
ws_port = String.to_integer(System.get_env("WS_PORT") || "8084")
wss_enabled = System.get_env("WSS_ENABLED", "false") == "true"
wss_port = String.to_integer(System.get_env("WSS_PORT") || "8085")

websocket_config = [
  ws_port: ws_port,
  wss_port: wss_port,
  max_connections: 1_000_000,
  max_messages_per_second: 100,
  heartbeat_interval: 30_000,
  connection_timeout: 60_000,
  ssl_enabled: wss_enabled,
  rate_limit_window: 1_000,
  rate_limit_burst: 150
]

websocket_config =
  if wss_enabled do
    Keyword.put(websocket_config, :ssl_certfile, System.get_env("SSL_CERTFILE"))
    |> Keyword.put(:ssl_keyfile, System.get_env("SSL_KEYFILE"))
    |> Keyword.put(:ssl_cacertfile, System.get_env("SSL_CACERTFILE"))
  else
    websocket_config
  end

config :websocket_gateway, :websocket, websocket_config

# NATS configuration
config :websocket_gateway, :nats,
  host: System.get_env("NATS_HOST"),
  port: String.to_integer(System.get_env("NATS_PORT") || "4222"),
  username: System.get_env("NATS_USERNAME", ""),
  password: System.get_env("NATS_PASSWORD", ""),
  connect_timeout: 10_000,
  ping_interval: 30_000,
  subscriptions: [
    "game.events.*",
    "tournament.events.*",
    "jackpot.events.*"
  ]

# Endpoint configuration
config :websocket_gateway, WebsocketGateway.Endpoint,
  url: [host: System.get_env("APP_HOST") || "localhost", port: ws_port],
  http: [
    ip: {0, 0, 0, 0},
    port: ws_port,
    compress: true,
    stream_window: 100,
    idle_timeout: 60_000
  ],
  secret_key_base: System.get_env("SECRET_KEY_BASE") || raise("SECRET_KEY_BASE must be set"),
  render_errors: [formats: [json: WebsocketGateway.ErrorJSON]],
  pubsub_server: WebsocketGateway.PubSub,
  live_view: [signing_salt: System.get_env("LIVE_VIEW_SALT") || raise("LIVE_VIEW_SALT must be set")]

# Phoenix production configuration
config :phoenix, :serve_endpoints, true

# Logging - JSON format for production
config :logger, :console,
  format: {:logger_formatter, :format, "$time $metadata[$level] $message\n"},
  metadata: [:request_id, :user_id, :game_id, :room_id],
  level: :info

# Disable debug routes in production
config :websocket_gateway, dev_routes: false

# Do not adapter runtime code compilation in production
config :phoenix, :plug_init_mode, :compile

# Telemetry - production metrics
config :websocket_gateway, :telemetry, enabled: true

# Cache static assets
config :websocket_gateway, WebsocketGateway.Endpoint,
  cache_static_manifest: "priv/static/cache_manifest.json"
