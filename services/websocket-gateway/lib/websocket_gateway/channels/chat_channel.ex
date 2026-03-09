defmodule WebsocketGateway.Channels.ChatChannel do
  use WebsocketGateway, :channel
  alias WebsocketGateway.Services.{Presence, RoomManager}

  # Chat channel topics:
  # "chat:room" - Public chat
  # "chat:game:{game_id}" - Game table chat
  # "chat:tournament:{tournament_id}" - Tournament chat
  # "chat:lobby" - Lobby chat

  @max_message_length Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:chat, [])
    |> Keyword.get(:max_message_length, 1000)

  @rate_limit Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:chat, [])
    |> Keyword.get(:rate_limit, 10)

  @impl true
  def join("chat:" <> room_type, params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username
    role = socket.assigns.role

    # Determine the full room ID
    room_id = case room_type do
      "room" -> "lobby"
      "game:" <> game_id -> "game:#{game_id}"
      "tournament:" <> tournament_id -> "tournament:#{tournament_id}"
      "lobby" -> "lobby"
      other -> other
    end

    # Track presence in chat room
    Presence.track(self(), "chat:#{room_id}", user_id, %{
      username: username,
      user_id: user_id,
      role: role,
      joined_at: DateTime.to_unix(DateTime.utc_now())
    })

    socket = socket
      |> assign(:chat_room_id, room_id)
      |> assign(:chat_role, role)

    # Get chat history
    history = get_chat_history(room_id)

    {:ok, %{room: room_id, history: history, users: get_chat_users(room_id)}, socket}
  end

  # Handle incoming messages
  @impl true
  def handle_in("send_message", %{"message" => message} = params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username
    chat_room_id = socket.assigns.chat_room_id
    chat_role = socket.assigns.chat_role

    # Check rate limit
    case check_rate_limit(socket) do
      :ok ->
        # Validate message
        case validate_message(message) do
          :ok ->
            message_data = %{
              id: generate_message_id(),
              user_id: user_id,
              username: username,
              message: message,
              room_id: chat_room_id,
              timestamp: DateTime.to_unix(DateTime.utc_now()),
              role: chat_role
            }

            # Save to history
            save_message(chat_room_id, message_data)

            # Broadcast to room
            broadcast! :new_message, message_data

            {:reply, {:ok, %{status: "sent", message_id: message_data.id}}, socket}

          {:error, reason} ->
            {:reply, {:error, %{reason: reason}}, socket}
        end

      {:error, :rate_limit_exceeded} ->
        {:reply, {:error, %{reason: "Rate limit exceeded. Please wait."}}, socket}
    end
  end

  def handle_in("typing", _params, socket) do
    username = socket.assigns.username
    chat_room_id = socket.assigns.chat_room_id

    broadcast! :user_typing, %{
      user_id: socket.assigns.user_id,
      username: username,
      room_id: chat_room_id
    }

    {:noreply, socket}
  end

  def handle_in("get_history", %{"limit" => limit} = params, socket) do
    chat_room_id = socket.assigns.chat_room_id
    history = get_chat_history(chat_room_id, limit)

    {:reply, {:ok, %{history: history}}, socket}
  end

  def handle_in("get_history", _params, socket) do
    handle_in("get_history", %{"limit" => 50}, socket)
  end

  # Moderator actions (admin/moderator only)
  def handle_in("mute_user", %{"user_id" => target_user_id, "duration" => duration}, socket) do
    if is_moderator(socket) do
      mute_duration = duration || Application.get_env(:channels, [])
        |> Keyword.get(:chat, [])
        |> Keyword.get(:mute_duration, websocket_gateway, :300)

      RoomManager.mute_user(socket.assigns.chat_room_id, target_user_id, mute_duration)

      broadcast! :user_muted, %{
        user_id: target_user_id,
        duration: mute_duration,
        moderator_id: socket.assigns.user_id
      }

      {:reply, {:ok, %{status: "muted"}}, socket}
    else
      {:reply, {:error, %{reason: "Moderator access required"}}, socket}
    end
  end

  def handle_in("ban_user", %{"user_id" => target_user_id, "reason" => reason}, socket) do
    if is_moderator(socket) do
      ban_duration = Application.get_env(:websocket_gateway, :channels, [])
        |> Keyword.get(:chat, [])
        |> Keyword.get(:ban_duration, 86400)

      RoomManager.ban_user(socket.assigns.chat_room_id, target_user_id, ban_duration, reason)

      broadcast! :user_banned, %{
        user_id: target_user_id,
        reason: reason,
        moderator_id: socket.assigns.user_id
      }

      {:reply, {:ok, %{status: "banned"}}, socket}
    else
      {:reply, {:error, %{reason: "Moderator access required"}}, socket}
    end
  end

  def handle_in("unmute_user", %{"user_id" => target_user_id}, socket) do
    if is_moderator(socket) do
      RoomManager.unmute_user(socket.assigns.chat_room_id, target_user_id)

      broadcast! :user_unmuted, %{
        user_id: target_user_id,
        moderator_id: socket.assigns.user_id
      }

      {:reply, {:ok, %{status: "unmuted"}}, socket}
    else
      {:reply, {:error, %{reason: "Moderator access required"}}, socket}
    end
  end

  def handle_in("get_users", _params, socket) do
    users = get_chat_users(socket.assigns.chat_room_id)
    {:reply, {:ok, %{users: users}}, socket}
  end

  @impl true
  def terminate(_reason, socket) do
    chat_room_id = socket.assigns.chat_room_id
    user_id = socket.assigns.user_id

    Presence.untrack(self(), "chat:#{chat_room_id}", user_id)
    :ok
  end

  # Private functions
  defp validate_message(message) when is_binary(message) do
    len = String.length(message)
    if len > 0 and len <= @max_message_length do
      :ok
    else
      {:error, "Message must be between 1 and #{@max_message_length} characters"}
    end
  end
  defp validate_message(_), do: {:error, "Invalid message format"}

  defp check_rate_limit(socket) do
    # Simple in-memory rate limiting
    now = DateTime.utc_now()
    last_message = socket.assigns[:last_message_time] || DateTime.add(now, -100)

    case DateTime.diff(now, last_message, :second) do
      seconds when seconds >= 60 / @rate_limit ->
        {:ok, socket}
      _ ->
        {:error, :rate_limit_exceeded}
    end
  end

  defp is_moderator(socket) do
    socket.assigns.chat_role in ["admin", "moderator"]
  end

  defp generate_message_id do
    :crypto.strong_rand_bytes(8) |> Base.encode16()
  end

  defp get_chat_history(room_id, limit \\ 50) do
    # TODO: Get from Redis
    []
  end

  defp save_message(room_id, message_data) do
    # TODO: Save to Redis with TTL
    :ok
  end

  defp get_chat_users(room_id) do
    case Presence.list("chat:#{room_id}") do
      %{metas: metas} -> Enum.map(metas, &Map.take(&1, [:user_id, :username, :role]))
      _ -> []
    end
  end
end
