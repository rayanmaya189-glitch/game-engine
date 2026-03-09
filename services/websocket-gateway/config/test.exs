import Config

# Test configuration

# Configure your database
config :websocket_gateway, WebsocketGateway.Repo,
  username: "postgres",
  password: "postgres",
  hostname: "localhost",
  database: "websocket_gateway_test#{System.get_env("MIX_TEST_PARTITION")}",
  pool: Ecto.Adapters.SQL.Sandbox,
  pool_size: 5

# Redis for testing
config :websocket_gateway, :redis,
  host: "localhost",
  port: 6379,
  password: "",
  database: 1,
  pool_size: 2

# JWT for testing
config :websocket_gateway, :jwt,
  secret_key: "test-secret-key",
  algorithm: "HS256",
  issuer: "game-engine-test",
  audience: "game-engine-test",
  expiration: 3600

# WebSocket for testing
config :websocket_gateway, :websocket,
  ws_port: 8084,
  wss_port: 8085,
  max_connections: 1000,
  max_messages_per_second: 100,
  heartbeat_interval: 30000,
  connection_timeout: 60000,
  ssl_enabled: false

# NATS for testing (disabled by default)
config :websocket_gateway, :nats,
  host: "localhost",
  port: 4222,
  username: "",
  password: "",
  connect_timeout: 5000,
  ping_interval: 30000,
  subscriptions: []

# Endpoint configuration
config :websocket_gateway, WebsocketGateway.Endpoint,
  http: [ip: {127, 0, 0, 1}, port: 4002],
  server: false

# Print only warnings and errors during tests
config :logger, level: :warning

# Initialize plugs at runtime for faster test compilation
config :phoenix, :plug_init_mode, :runtime

# Disable coverage and other dev-only features
config :phoenix, :debug_application_key_checks, false
