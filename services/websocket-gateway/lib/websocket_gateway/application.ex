defmodule WebsocketGateway.Application do
  @moduledoc """
  Application module for WebsocketGateway.

  This is the main entry point for the Phoenix application.
  It starts the supervision tree which includes:
  - Phoenix Endpoint (WebSocket handler)
  - PubSub for channel communication
  - Presence tracking
  - NATS client
  - Room manager
  - Telemetry
  """

  use Application

  @impl true
  def start(_type, _args) do
    # Load configuration
    websocket_config = Application.get_env(:websocket_gateway, :websocket, [])
    ws_port = Keyword.get(websocket_config, :ws_port, 8084)
    wss_port = Keyword.get(websocket_config, :wss_port, 8085)

    children = [
      # Start the PubSub system
      {Phoenix.PubSub, name: WebsocketGateway.PubSub},

      # Start the Endpoint (http/https)
      WebsocketGateway.Endpoint,

      # Presence tracking
      {WebsocketGateway.Services.Presence, []},

      # Room manager
      {WebsocketGateway.Services.RoomManager, []},

      # NATS client for event subscription
      {WebsocketGateway.NATS.Client, []},

      # Telemetry
      WebsocketGateway.Telemetry,

      # Redis connection
      {Redix, redis_config()}
    ]
    |> maybe_start_websocket(websocket_config)

    # Start metrics endpoint in production
    children = if Application.get_env(:websocket_gateway, :telemetry, [])[:enabled] do
      children ++ [WebsocketGateway.Telemetry.Endpoint]
    else
      children
    end

    opts = [strategy: :one_for_one, name: WebsocketGateway.Supervisor]
    Supervisor.start_link(children, opts)
  end

  defp redis_config do
    config = Application.get_env(:websocket_gateway, :redis, [])
    [
      host: Keyword.get(config, :host, "localhost"),
      port: Keyword.get(config, :port, 6379),
      password: Keyword.get(config, :password, ""),
      database: Keyword.get(config, :database, 0),
      pool_size: Keyword.get(config, :pool_size, 10)
    ]
  end

  defp maybe_start_websocket(children, config) do
    ws_port = Keyword.get(config, :ws_port, 8084)

    # Configure HTTP endpoint with WebSocket support
    http_config = [
      ip: {0, 0, 0, 0},
      port: ws_port,
      compress: true,
      stream_window: 100,
      idle_timeout: 60_000,
      max_connections: Keyword.get(config, :max_connections, 1_000_000)
    ]

    # Add SSL config if WSS is enabled
    http_config = if Keyword.get(config, :ssl_enabled, false) do
      Keyword.put(http_config, :transport_options, [
        :inet6,
        certfile: Keyword.get(config, :ssl_certfile, ""),
        keyfile: Keyword.get(config, :ssl_keyfile, ""),
        cacertfile: Keyword.get(config, :ssl_cacertfile, "")
      ])
    else
      http_config
    end

    # Update endpoint config
    Application.put_env(:websocket_gateway, WebsocketGateway.Endpoint,
      http: http_config,
      debug_errors: Application.get_env(:websocket_gateway, :debug_errors, false),
      secret_key_base: Application.get_env(:websocket_gateway, :secret_key_base,
        String.duplicate("secret", 8)),
      render_errors: [formats: [json: WebsocketGateway.ErrorJSON]],
      pubsub_server: WebsocketGateway.PubSub,
      live_view: [signing_salt: "websocket_gateway"]
    )

    children
  end

  @impl true
  def config_change(changed, _new, removed) do
    WebsocketGateway.Endpoint.config_change(changed, removed)
    :ok
  end
end
