defmodule WebsocketGateway.Channels.GameChannel do
  use WebsocketGateway, :channel
  alias WebsocketGateway.Services.{Presence, RoomManager}
  alias Phoenix.Socket.Broadcast

  # Game channel topic pattern: "game:{game_id}"
  # Or with sub-topic: "game:{game_id}:table:{table_id}"

  @max_players Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:game, [])
    |> Keyword.get(:max_players_per_room, 6)

  @broadcast_timeout Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:game, [])
    |> Keyword.get(:broadcast_timeout, 5_000)

  @impl true
  def join("game:" <> game_id, params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Check if game exists and is joinable
    case RoomManager.join_room(game_id, user_id, username, @max_players) do
      {:ok, room_state} ->
        # Track presence in the game room
        Presence.track(self(), "game:#{game_id}", user_id, %{
          username: username,
          user_id: user_id,
          seat: Map.get(room_state, :seats, %{}) |> find_seat(user_id),
          joined_at: DateTime.to_unix(DateTime.utc_now())
        })

        # Subscribe to game events
        subscribe_to_game_events(game_id)

        socket = socket
          |> assign(:game_id, game_id)
          |> assign(:room_state, room_state)

        {:ok, %{room: room_state, players: get_room_players(game_id)}, socket}

      {:error, :room_full} ->
        {:error, %{reason: "Game room is full"}}

      {:error, :game_not_found} ->
        {:error, %{reason: "Game not found"}}

      {:error, reason} ->
        {:error, %{reason: inspect(reason)}}
    end
  end

  # Handle incoming events from client
  @impl true
  def handle_in("place_bet", %{"amount" => amount, "bet_type" => bet_type} = params, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id

    # Validate bet
    case validate_bet(amount, bet_type, socket) do
      :ok ->
        # Broadcast bet to game service
        broadcast_bet_event(game_id, user_id, params)
        {:reply, {:ok, %{status: "bet_placed", amount: amount, bet_type: bet_type}}, socket}

      {:error, reason} ->
        {:reply, {:error, %{reason: reason}}, socket}
    end
  end

  def handle_in("player_action", %{"action" => action} = params, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id

    # Validate action
    case validate_action(action, socket) do
      :ok ->
        broadcast_action_event(game_id, user_id, params)
        {:reply, {:ok, %{status: "action_processed", action: action}}, socket}

      {:error, reason} ->
        {:reply, {:error, %{reason: reason}}, socket}
    end
  end

  def handle_in("ready", _params, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id

    # Mark player as ready
    RoomManager.set_player_ready(game_id, user_id)

    broadcast! :player_ready, %{
      user_id: user_id,
      username: socket.assigns.username,
      ready_at: DateTime.to_unix(DateTime.utc_now())
    }

    {:noreply, socket}
  end

  def handle_in("leave_game", _params, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id

    RoomManager.leave_room(game_id, user_id)

    # Untrack presence
    Presence.untrack(self(), "game:#{game_id}", user_id)

    {:stop, :normal, %{status: "left_game"}, socket}
  end

  def handle_in("get_state", _params, socket) do
    game_id = socket.assigns.game_id
    room_state = RoomManager.get_room_state(game_id)

    {:reply, {:ok, %{room: room_state}}, socket}
  end

  def handle_in("chat_message", %{"message" => message} = params, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Forward to chat channel
    WebsocketGateway.Endpoint.broadcast("chat:game:#{game_id}", "new_message", %{
      user_id: user_id,
      username: username,
      message: message,
      timestamp: DateTime.to_unix(DateTime.utc_now())
    })

    {:noreply, socket}
  end

  # Handle game state updates from NATS
  @impl true
  def handle_info(%Broadcast{payload: payload, topic: "game.events." <> _}, socket) do
    # Broadcast game events to clients
    push(socket, "game_event", payload)
    {:noreply, socket}
  end

  def handle_info(%Broadcast{payload: payload, event: event}, socket) do
    push(socket, event, payload)
    {:noreply, socket}
  end

  # Handle member join/leave in presence
  def handle_info({:member_join, user_id, %{username: username}}, socket) do
    push(socket, "player_joined", %{
      user_id: user_id,
      username: username,
      timestamp: DateTime.to_unix(DateTime.utc_now())
    })
    {:noreply, socket}
  end

  def handle_info({:member_leave, user_id, %{username: username}}, socket) do
    push(socket, "player_left", %{
      user_id: user_id,
      username: username,
      timestamp: DateTime.to_unix(DateTime.utc_now())
    })
    {:noreply, socket}
  end

  @impl true
  def terminate(_reason, socket) do
    game_id = socket.assigns.game_id
    user_id = socket.assigns.user_id

    RoomManager.leave_room(game_id, user_id)
    Presence.untrack(self(), "game:#{game_id}", user_id)

    :ok
  end

  # Private functions
  defp validate_bet(amount, bet_type, socket) do
    with {:ok, amount} <- validate_amount(amount),
         :ok <- validate_bet_type(bet_type),
         :ok <- check_balance(amount, socket) do
      :ok
    else
      error -> error
    end
  end

  defp validate_amount(amount) when is_number(amount) and amount > 0, do: :ok
  defp validate_amount(_), do: {:error, "Invalid bet amount"}

  defp validate_bet_type(bet_type) when is_binary(bet_type), do: :ok
  defp validate_bet_type(_), do: {:error, "Invalid bet type"}

  # Check balance via NATS request to wallet service
  defp check_balance(amount, socket) do
    user_id = socket.assigns.user_id

    # Send request to wallet service via NATS
    case WebsocketGateway.NATS.Client.request("wallet.balance.check", %{user_id: user_id, amount: amount}, 5_000) do
      {:ok, response} ->
        if response["success"] == true and response["has_sufficient_balance"] == true do
          :ok
        else
          {:error, response["error"] || "Insufficient balance"}
        end
      {:error, reason} ->
        # Fail closed - deny the bet if we can't verify balance
        Logger.error("Failed to check balance: #{inspect(reason)}")
        {:error, "Unable to verify balance. Please try again."}
    end
  end

  defp validate_action(action, socket) do
    allowed_actions = ~w(hit stand double surrender split fold bet raise call check)
    if action in allowed_actions, do: :ok, else: {:error, "Invalid action"}
  end

  defp find_seat(seats, user_id) do
    seats
    |> Enum.find(fn {_seat, info} -> info.user_id == user_id end)
    |> case do
      {seat, _} -> seat
      nil -> nil
    end
  end

  defp get_room_players(game_id) do
    case Presence.list("game:#{game_id}") do
      %{metas: metas} -> Enum.map(metas, &Map.take(&1, [:user_id, :username, :seat]))
      _ -> []
    end
  end

  defp subscribe_to_game_events(game_id) do
    Phoenix.PubSub.subscribe(WebsocketGateway.PubSub, "game:events:#{game_id}")
  end

  defp broadcast_bet_event(game_id, user_id, params) do
    Phoenix.PubSub.broadcast(WebsocketGateway.PubSub, "game:events:#{game_id}", %{
      event: "bet_placed",
      user_id: user_id,
      amount: params["amount"],
      bet_type: params["bet_type"],
      timestamp: DateTime.to_unix(DateTime.utc_now())
    })
  end

  defp broadcast_action_event(game_id, user_id, params) do
    Phoenix.PubSub.broadcast(WebsocketGateway.PubSub, "game:events:#{game_id}", %{
      event: "player_action",
      user_id: user_id,
      action: params["action"],
      timestamp: DateTime.to_unix(DateTime.utc_now())
    })
  end
end
