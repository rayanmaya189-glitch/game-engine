# This file is responsible for configuring your application
# and its dependencies with the aid of the Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

# Configures the endpoint
config :websocket_gateway, WebsocketGateway.Endpoint,
  url: [host: "localhost"],
  adapter: Phoenix.Endpoint.Cowboy2Adapter,
  http: [
    dispatch: [
      {:_, [
        {Phoenix.Endpoint.Cowboy2Handler, {WebsocketGateway.Endpoint, []}},
        {:_, Phoenix.Endpoint.Cowboy2WebSocketHandler, {Phoenix.Endpoint.Cowboy2WebSocket, {WebsocketGateway.Endpoint, []}}}
      ]}
    ]
  ],
  render_errors: [formats: [json: WebsocketGateway.ErrorJSON]],
  pubsub_server: WebsocketGateway.PubSub,
  live_view: [signing_salt: "websocket_gateway_secret"]

# Configures the database
config :websocket_gateway, WebsocketGateway.Repo,
  pool_size: 10,
  pool_timeout: 5000,
  timeout: 15_000

# Redis configuration
config :websocket_gateway, :redis,
  host: System.get_env("REDIS_HOST") || "localhost",
  port: String.to_integer(System.get_env("REDIS_PORT") || "6379"),
  password: System.get_env("REDIS_PASSWORD", ""),
  database: String.to_integer(System.get_env("REDIS_DB") || "0"),
  pool_size: 10

# JWT configuration
config :websocket_gateway, :jwt,
  secret_key: System.get_env("JWT_SECRET_KEY") || "your-secret-key-change-in-production",
  algorithm: "HS256",
  issuer: "game-engine",
  audience: "game-engine",
  expiration: 86_400, # 24 hours in seconds

# WebSocket configuration
config :websocket_gateway, :websocket,
  # Port configuration
  ws_port: String.to_integer(System.get_env("WS_PORT") || "8084"),
  wss_port: String.to_integer(System.get_env("WSS_PORT") || "8085"),
  
  # Connection limits
  max_connections: 1_000_000,
  max_messages_per_second: 100,
  heartbeat_interval: 30_000, # 30 seconds
  connection_timeout: 60_000, # 60 seconds
  
  # SSL configuration
  ssl_enabled: System.get_env("WSS_ENABLED", "false") == "true",
  ssl_certfile: System.get_env("SSL_CERTFILE"),
  ssl_keyfile: System.get_env("SSL_KEYFILE"),

  # Rate limiting
  rate_limit_window: 1_000, # 1 second window
  rate_limit_burst: 150

# NATS configuration
config :websocket_gateway, :nats,
  host: System.get_env("NATS_HOST") || "localhost",
  port: String.to_integer(System.get_env("NATS_PORT") || "4222"),
  username: System.get_env("NATS_USERNAME", ""),
  password: System.get_env("NATS_PASSWORD", ""),
  connect_timeout: 5_000,
  ping_interval: 30_000,
  subscriptions: [
    "game.events.*",
    "tournament.events.*",
    "jackpot.events.*"
  ]

# Phoenix channels configuration
config :websocket_gateway, :channels,
  # Game channel
  game: [
    max_players_per_room: 6,
    broadcast_timeout: 5_000,
    state_sync_interval: 100
  ],
  
  # Chat channel
  chat: [
    max_message_length: 1000,
    rate_limit: 10, # messages per minute
    mute_duration: 300, # seconds
    ban_duration: 86400 # 24 hours
  ],
  
  # Tournament channel
  tournament: [
    max_participants: 1000,
    leaderboard_update_interval: 5_000
  ],
  
  # Lobby channel
  lobby: [
    featured_games_count: 10,
    jackpot_update_interval: 5_000
  ]

# Presence tracking
config :websocket_gateway, :presence,
  pubsub_server: WebsocketGateway.PubSub,
  timeout: 25_000, # Heartbeat timeout (should be less than heartbeat_interval)

# Room management
config :websocket_gateway, :room_manager,
  max_rooms_per_user: 10,
  room_cleanup_interval: 60_000, # 1 minute
  room_idle_timeout: 300_000 # 5 minutes

# Telemetry
config :websocket_gateway, :telemetry,
  enabled: true

# Logging
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id, :user_id, :game_id, :room_id]

# Import environment specific config
import_config "#{config_env()}.exs"
