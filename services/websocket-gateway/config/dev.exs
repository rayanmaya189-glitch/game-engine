import Config

# For development, we don't run any checks and only start
# relevant repositories for easier debugging

# Configure your database
config :websocket_gateway, WebsocketGateway.Repo,
  username: "postgres",
  password: "postgres",
  hostname: "localhost",
  database: "websocket_gateway_dev",
  pool_size: 5,
  stacktrace: true,
  show_sensitive_data_on_connection_error: true

# Redis for development
config :websocket_gateway, :redis,
  host: "localhost",
  port: 6379,
  password: "",
  database: 0,
  pool_size: 5

# JWT for development
config :websocket_gateway, :jwt,
  secret_key: "dev-secret-key-change-in-development-only",
  algorithm: "HS256",
  issuer: "game_engine-dev",
  audience: "game_engine-dev",
  expiration: 86400

# WebSocket for development
config :websocket_gateway, :websocket,
  ws_port: 8084,
  wss_port: 8085,
  max_connections: 10000,
  max_messages_per_second: 100,
  heartbeat_interval: 30000,
  connection_timeout: 60000,
  ssl_enabled: false

# NATS for development
config :websocket_gateway, :nats,
  host: "localhost",
  port: 4222,
  username: "",
  password: "",
  connect_timeout: 5000,
  ping_interval: 30000,
  subscriptions: [
    "game.events.*",
    "tournament.events.*",
    "jackpot.events.*"
  ]

# Watch and recompile in dev
config :websocket_gateway, WebsocketGateway.Endpoint,
  http: [ip: {0, 0, 0, 0}, port: 4000],
  debug_errors: true,
  code_reloader: true,
  check_origin: false,
  watchers: []

# The watchers key for running code reloaders
config :websocket_gateway, WebsocketGateway.Endpoint,
  watchers: [
    node: [
      "--no-warnings",
      "-e",
      "IO.puts('Phoenix is running. Connect to ws://localhost:8084 for WebSocket')"
    ]
  ]

# Enable dev routes for dashboard
config :websocket_gateway, dev_routes: true

# Do not include metadata nor timestamps in development logs
config :logger, :console, format: "[$level] $message\n"

# Set a higher stacktrace during development
config :phoenix, :stacktrace_depth, 20

# Initialize plugs at runtime for faster dev compilation
config :phoenix, :plug_init_mode, :runtime

# Enable helpful but potentially expensive debug logs
config :phoenix, :debug_application_key_checks, false
