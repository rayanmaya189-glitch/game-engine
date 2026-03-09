defmodule WebsocketGateway.UserSocket do
  use Phoenix.Socket

  # Transport configuration
  transport :websocket, Phoenix.Transports.WebSocket,
    timeout: 60_000,
    lifespan: :infinite

  # Longpoll transport (optional, for compatibility)
  transport :longpoll, Phoenix.Transports.LongPoll,
    window_ms: 10_000,
    lifespan: :infinite

  # JWT token parameter
  channel "game:*", WebsocketGateway.Channels.GameChannel
  channel "chat:*", WebsocketGateway.Channels.ChatChannel
  channel "tournament:*", WebsocketGateway.Channels.TournamentChannel
  channel "lobby", WebsocketGateway.Channels.LobbyChannel

  @heartbeat_interval 30_000
  @max_message_rate 100

  @impl true
  def connect(%{"token" => token}, socket, connect_info) do
    # Validate JWT token
    case WebsocketGateway.Services.Auth.validate_token(token) do
      {:ok, claims} ->
        user_id = Map.get(claims, "sub") || Map.get(claims, "user_id")
        username = Map.get(claims, "username")
        email = Map.get(claims, "email")
        role = Map.get(claims, "role", "user")

        socket = socket
          |> assign(:user_id, user_id)
          |> assign(:username, username)
          |> assign(:email, email)
          |> assign(:role, role)
          |> assign(:connected_at, DateTime.utc_now())
          |> assign(:message_count, 0)
          |> assign(:last_message_time, DateTime.utc_now())

        # Track presence
        WebsocketGateway.Services.Presence.track_user(
          self(),
          user_id,
          %{
            username: username,
            email: email,
            role: role,
            connected_at: DateTime.to_unix(DateTime.utc_now()),
            ip: get_ip(connect_info),
            user_agent: get_user_agent(connect_info)
          }
        )

        {:ok, socket}

      {:error, :token_expired} ->
        :error

      {:error, :invalid_token} ->
        :error

      {:error, reason} ->
        Logger.warning("Socket connection failed: #{inspect(reason)}")
        :error
    end
  end

  def connect(_params, _socket, _connect_info) do
    :error
  end

  @impl true
  def id(socket), do: "user:#{socket.assigns.user_id}"

  # Handle incoming messages with rate limiting
  @impl true
  def handle_in(event, params, socket) do
    # Check rate limit
    case check_rate_limit(socket) do
      :ok ->
        increment_message_count(socket)
        apply(__MODULE__, :process_message, [event, params, socket])

      {:error, :rate_limit_exceeded} ->
        {:reply, {:error, %{error: "Rate limit exceeded"}}, socket}
    end
  end

  def process_message("ping", _params, socket) do
    {:reply, {:ok, %{pong: DateTime.utc_now()}}, socket}
  end

  def process_message("heartbeat", _params, socket) do
    {:reply, {:ok, %{heartbeat: DateTime.utc_now()}}, socket}
  end

  def process_message(event, params, socket) do
    Logger.debug("Unhandled message event: #{event}, params: #{inspect(params)}")
    {:reply, {:error, %{error: "Unknown event: #{event}"}}, socket}
  end

  # Handle socket termination
  @impl true
  def terminate(reason, socket) do
    user_id = socket.assigns.user_id

    # Remove presence
    WebsocketGateway.Services.Presence.untrack_user(self(), user_id)

    # Clean up room associations
    WebsocketGateway.Services.RoomManager.handle_disconnect(user_id)

    Logger.info("Socket disconnected: user_id=#{user_id}, reason=#{inspect(reason)}")
    :ok
  end

  # Rate limiting
  defp check_rate_limit(socket) do
    now = DateTime.utc_now()
    last_time = socket.assigns.last_message_time

    # Get rate limit config
    rate_limit_window = Application.get_env(:websocket_gateway, :websocket, [])
      |> Keyword.get(:rate_limit_window, 1_000)

    elapsed = DateTime.diff(now, last_time, :millisecond)

    if elapsed > rate_limit_window do
      :ok
    else
      message_count = socket.assigns.message_count
      max_messages = Application.get_env(:websocket_gateway, :websocket, [])
        |> Keyword.get(:rate_limit_burst, 150)

      if message_count < max_messages do
        :ok
      else
        {:error, :rate_limit_exceeded}
      end
    end
  end

  defp increment_message_count(socket) do
    new_count = socket.assigns.message_count + 1
    {:noreply, assign(socket, :message_count, new_count)}
  end

  defp get_ip(connect_info) do
    case connect_info[:peer_data] do
      %{address: {a, b, c, d}} -> "#{a}.#{b}.#{c}.#{d}"
      %{address: {a, b, c, d, e, f, g, h}} -> "#{a}:#{b}:#{c}:#{d}:#{e}:#{f}:#{g}:#{h}"
      _ -> "unknown"
    end
  end

  defp get_user_agent(connect_info) do
    connect_info[:x_headers]
    |> Enum.find(fn {key, _} -> key == "user-agent" end)
    |> case do
      {_, ua} -> ua
      nil -> "unknown"
    end
  end
end
