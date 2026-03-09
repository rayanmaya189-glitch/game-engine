defmodule WebsocketGateway.Channels.LobbyChannel do
  use WebsocketGateway, :channel
  alias WebsocketGateway.Services.{Presence, RoomManager}

  # Lobby channel: "lobby" - Main lobby for featured games, jackpots, etc.

  @featured_games_count Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:lobby, [])
    |> Keyword.get(:featured_games_count, 10)

  @jackpot_update_interval Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:lobby, [])
    |> Keyword.get(:jackpot_update_interval, 5_000)

  @impl true
  def join("lobby", _params, socket) do
    user_id = socket.assigns.user_id
    username = socket.assigns.username

    # Track presence in lobby
    Presence.track(self(), "lobby", user_id, %{
      username: username,
      user_id: user_id,
      joined_at: DateTime.to_unix(DateTime.utc_now())
    })

    socket = socket
      |> assign(:is_lobby, true)

    # Get lobby data
    lobby_data = get_lobby_data()

    {:ok, lobby_data, socket}
  end

  # Handle incoming events
  @impl true
  def handle_in("get_featured_games", _params, socket) do
    games = get_featured_games()
    {:reply, {:ok, %{games: games}}, socket}
  end

  def handle_in("get_jackpot", _params, socket) do
    jackpot = get_current_jackpot()
    {:reply, {:ok, %{jackpot: jackpot}}, socket}
  end

  def handle_in("get_player_counts", _params, socket) do
    counts = get_player_counts()
    {:reply, {:ok, %{counts: counts}}, socket}
  end

  def handle_in("get_online_players", _params, socket) do
    online = get_online_players()
    {:reply, {:ok, %{players: online}}, socket}
  end

  def handle_in("search_games", %{"query" => query} = params, socket) do
    limit = Map.get(params, "limit", 20)
    games = search_games(query, limit)
    {:reply, {:ok, %{games: games}}, socket}
  end

  def handle_in("get_game_categories", _params, socket) do
    categories = get_game_categories()
    {:reply, {:ok, %{categories: categories}}, socket}
  end

  def handle_in("get_new_games", _params, socket) do
    games = get_new_games()
    {:reply, {:ok, %{games: games}}, socket}
  end

  def handle_in("get_popular_games", _params, socket) do
    games = get_popular_games()
    {:reply, {:ok, %{games: games}}, socket}
  end

  def handle_in("subscribe_jackpot", _params, socket) do
    # Start periodic jackpot updates
    schedule_jackpot_updates(socket)
    {:reply, {:ok, %{status: "subscribed"}}, socket}
  end

  # Handle periodic updates
  def handle_info(:jackpot_update, socket) do
    jackpot = get_current_jackpot()
    push(socket, "jackpot_update", %{jackpot: jackpot})
    {:noreply, socket}
  end

  def handle_info(:featured_games_update, socket) do
    games = get_featured_games()
    push(socket, "featured_games_update", %{games: games})
    {:noreply, socket}
  end

  def handle_info(:player_counts_update, socket) do
    counts = get_player_counts()
    push(socket, "player_counts_update", %{counts: counts})
    {:noreply, socket}
  end

  @impl true
  def terminate(_reason, socket) do
    user_id = socket.assigns.user_id
    Presence.untrack(self(), "lobby", user_id)
    :ok
  end

  # Private functions
  defp get_lobby_data do
    %{
      featured_games: get_featured_games(),
      jackpot: get_current_jackpot(),
      player_counts: get_player_counts(),
      online_count: get_online_count(),
      new_games: get_new_games(),
      popular_games: get_popular_games(),
      categories: get_game_categories()
    }
  end

  defp get_featured_games do
    # TODO: Get from game registry service
    []
  end

  defp get_new_games do
    # TODO: Get from game registry service
    []
  end

  defp get_popular_games do
    # TODO: Get from game registry service
    []
  end

  defp get_current_jackpot do
    # TODO: Get from jackpot service via Redis
    %{
      amount: 0,
      currency: "USD",
      last_winner: nil,
      last_win_time: nil,
      games: []
    }
  end

  defp get_player_counts do
    # Get counts per game/category
    # TODO: Aggregate from presence tracking
    %{}
  end

  defp get_online_count do
    case Presence.list("lobby") do
      %{metas: metas} -> length(metas)
      _ -> 0
    end
  end

  defp get_online_players do
    case Presence.list("lobby") do
      %{metas: metas} ->
        metas
        |> Enum.map(&Map.take(&1, [:user_id, :username]))
      _ -> []
    end
  end

  defp get_game_categories do
    # TODO: Get from game registry service
    [
      %{id: "slots", name: "Slots", icon: "slot"},
      %{id: "table", name: "Table Games", icon: "table"},
      %{id: "live", name: "Live Casino", icon: "live"},
      %{id: "poker", name: "Poker", icon: "poker"},
      %{id: "arcade", name: "Arcade", icon: "arcade"}
    ]
  end

  defp search_games(query, limit) do
    # TODO: Implement search via game registry service
    []
  end

  defp schedule_jackpot_updates(socket) do
    # Schedule periodic updates
    Process.send_after(self(), :jackpot_update, @jackpot_update_interval)
    {:noreply, socket}
  end
end
