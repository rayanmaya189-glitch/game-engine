defmodule WebsocketGateway.Channels.TournamentChannel do
  use WebsocketGateway, :channel
  alias WebsocketGateway.Services.{Presence, RoomManager, ServiceClient}

  # Tournament channel topics:
  # "tournament:lobby" - Main tournament lobby
  # "tournament:{tournament_id}" - Specific tournament room

  @max_participants Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:tournament, [])
    |> Keyword.get(:max_participants, 1000)

  @leaderboard_update_interval Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:tournament, [])
    |> Keyword.get(:leaderboard_update_interval, 5_000)

  @impl true
  def join("tournament:lobby", _params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Track presence in tournament lobby
    Presence.track(self(), "tournament:lobby", user_id, %{
      username: username,
      user_id: user_id,
      joined_at: DateTime.to_unix(DateTime.utc_now())
    })

    socket = socket
      |> assign(:tournament_id, "lobby")
      |> assign(:is_lobby, true)

    # Get active tournaments
    tournaments = get_active_tournaments()

    {:ok, %{tournaments: tournaments, online_count: get_online_count()}, socket}
  end

  def join("tournament:" <> tournament_id, params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Check if tournament exists and is joinable
    case RoomManager.join_tournament(tournament_id, user_id, username, @max_participants) do
      {:ok, tournament_state} ->
        # Track presence
        Presence.track(self(), "tournament:#{tournament_id}", user_id, %{
          username: username,
          user_id: user_id,
          joined_at: DateTime.to_unix(DateTime.utc_now()),
          status: tournament_state.status
        })

        # Subscribe to tournament events
        subscribe_to_tournament_events(tournament_id)

        socket = socket
          |> assign(:tournament_id, tournament_id)
          |> assign(:is_lobby, false)
          |> assign(:tournament_state, tournament_state)

        {:ok, %{tournament: tournament_state, leaderboard: get_leaderboard(tournament_id)}, socket}

      {:error, :tournament_full} ->
        {:error, %{reason: "Tournament is full"}}

      {:error, :tournament_not_found} ->
        {:error, %{reason: "Tournament not found"}}

      {:error, :tournament_ended} ->
        {:error, %{reason: "Tournament has ended"}}

      {:error, reason} ->
        {:error, %{reason: inspect(reason)}}
    end
  end

  # Handle incoming events
  @impl true
  def handle_in("register", %{"buy_in" => buy_in} = params, socket) do
    tournament_id = socket.assigns.tournament_id
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Process registration
    case RoomManager.register_player(tournament_id, user_id, username, buy_in) do
      {:ok, registration} ->
        broadcast! :player_registered, %{
          user_id: user_id,
          username: username,
          buy_in: buy_in,
          timestamp: DateTime.to_unix(DateTime.utc_now())
        }

        {:reply, {:ok, %{status: "registered", registration: registration}}, socket}

      {:error, reason} ->
        {:reply, {:error, %{reason: reason}}, socket}
    end
  end

  def handle_in("unregister", _params, socket) do
    tournament_id = socket.assigns.tournament_id
    user_id = socket.assigns.user_id

    case RoomManager.unregister_player(tournament_id, user_id) do
      :ok ->
        broadcast! :player_unregistered, %{
          user_id: user_id,
          timestamp: DateTime.to_unix(DateTime.utc_now())
        }

        {:reply, {:ok, %{status: "unregistered"}}, socket}

      {:error, reason} ->
        {:reply, {:error, %{reason: reason}}, socket}
    end
  end

  def handle_in("get_leaderboard", _params, socket) do
    tournament_id = socket.assigns.tournament_id
    leaderboard = get_leaderboard(tournament_id)

    {:reply, {:ok, %{leaderboard: leaderboard}}, socket}
  end

  def handle_in("get_standings", _params, socket) do
    tournament_id = socket.assigns.tournament_id
    standings = get_standings(tournament_id)

    {:reply, {:ok, %{standings: standings}}, socket}
  end

  def handle_in("get_schedule", _params, socket) do
    schedule = get_tournament_schedule()

    {:reply, {:ok, %{schedule: schedule}}, socket}
  end

  def handle_in("leave_tournament", _params, socket) do
    tournament_id = socket.assigns.tournament_id
    user_id = socket.assigns.user_id

    RoomManager.leave_tournament(tournament_id, user_id)
    Presence.untrack(self(), "tournament:#{tournament_id}", user_id)

    {:stop, :normal, %{status: "left_tournament"}, socket}
  end

  # Handle tournament events from NATS
  @impl true
  def handle_info(%Phoenix.Socket.Broadcast{payload: payload, topic: "tournament.events." <> _}, socket) do
    push(socket, "tournament_event", payload)
    {:noreply, socket}
  end

  def handle_info({:leaderboard_update, leaderboard}, socket) do
    push(socket, "leaderboard_update", %{leaderboard: leaderboard})
    {:noreply, socket}
  end

  def handle_info({:tournament_state_update, state}, socket) do
    push(socket, "tournament_update", %{state: state})
    {:noreply, socket}
  end

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
    unless socket.assigns[:is_lobby] do
      tournament_id = socket.assigns.tournament_id
      user_id = socket.assigns.user_id

      RoomManager.leave_tournament(tournament_id, user_id)
      Presence.untrack(self(), "tournament:#{tournament_id}", user_id)
    end

    :ok
  end

  # Private functions
  defp get_active_tournaments do
    case ServiceClient.list_active_tournaments() do
      {:ok, %{"tournaments" => tournaments}} -> tournaments
      {:ok, response} -> Map.get(response, "tournaments", [])
      {:error, _} -> []
    end
  end

  defp get_tournament_schedule do
    case ServiceClient.get_tournament_schedule() do
      {:ok, %{"schedule" => schedule}} -> schedule
      {:ok, response} -> Map.get(response, "schedule", [])
      {:error, _} -> []
    end
  end

  defp get_online_count do
    case Presence.list("tournament:lobby") do
      %{metas: metas} -> length(metas)
      _ -> 0
    end
  end

  defp get_leaderboard(tournament_id) do
    case ServiceClient.get_leaderboard(tournament_id) do
      {:ok, %{"leaderboard" => leaderboard}} -> leaderboard
      {:ok, response} -> Map.get(response, "leaderboard", [])
      {:error, _} -> []
    end
  end

  defp get_standings(tournament_id) do
    case ServiceClient.get_standings(tournament_id) do
      {:ok, %{"standings" => standings}} -> standings
      {:ok, response} -> Map.get(response, "standings", [])
      {:error, _} -> []
    end
  end

  defp subscribe_to_tournament_events(tournament_id) do
    Phoenix.PubSub.subscribe(WebsocketGateway.PubSub, "tournament:events:#{tournament_id}")
  end
end
