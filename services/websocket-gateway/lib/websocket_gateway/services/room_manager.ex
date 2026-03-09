defmodule WebsocketGateway.Services.RoomManager do
  @moduledoc """
  Room Manager service for managing game rooms and tournament registrations.

  Provides functionality for:
  - Creating/joining/leaving game rooms
  - Managing room state
  - Tournament registration
  - Player management
  """

  use GenServer
  require Logger

  @max_rooms_per_user Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:max_rooms_per_user, 10)

  @room_cleanup_interval Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:room_cleanup_interval, 60_000)

  @room_idle_timeout Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:room_idle_timeout, 300_000)

  @game_config Application.get_env(:websocket_gateway, :channels, [])
    |> Keyword.get(:game, [])

  @max_players Keyword.get(@game_config, :max_players_per_room, 6)

  # Client API

  @doc """
  Start the RoomManager service.
  """
  def start_link(opts \\ []) do
    GenServer.start_link(__MODULE__, opts, name: __MODULE__)
  end

  @doc """
  Join a game room.
  """
  def join_room(room_id, user_id, username, max_players \\ @max_players) do
    GenServer.call(__MODULE__, {:join_room, room_id, user_id, username, max_players})
  end

  @doc """
  Leave a game room.
  """
  def leave_room(room_id, user_id) do
    GenServer.call(__MODULE__, {:leave_room, room_id, user_id})
  end

  @doc """
  Get room state.
  """
  def get_room_state(room_id) do
    GenServer.call(__MODULE__, {:get_room_state, room_id})
  end

  @doc """
  Set player as ready.
  """
  def set_player_ready(room_id, user_id) do
    GenServer.call(__MODULE__, {:set_ready, room_id, user_id})
  end

  @doc """
  Handle user disconnect - clean up room associations.
  """
  def handle_disconnect(user_id) do
    GenServer.cast(__MODULE__, {:handle_disconnect, user_id})
  end

  @doc """
  Join a tournament.
  """
  def join_tournament(tournament_id, user_id, username, max_participants \\ 1000) do
    GenServer.call(__MODULE__, {:join_tournament, tournament_id, user_id, username, max_participants})
  end

  @doc """
  Leave a tournament.
  """
  def leave_tournament(tournament_id, user_id) do
    GenServer.call(__MODULE__, {:leave_tournament, tournament_id, user_id})
  end

  @doc """
  Register player for tournament.
  """
  def register_player(tournament_id, user_id, username, buy_in) do
    GenServer.call(__MODULE__, {:register_player, tournament_id, user_id, username, buy_in})
  end

  @doc """
  Unregister player from tournament.
  """
  def unregister_player(tournament_id, user_id) do
    GenServer.call(__MODULE__, {:unregister_player, tournament_id, user_id})
  end

  @doc """
  Mute a user in a chat room.
  """
  def mute_user(room_id, user_id, duration) do
    GenServer.call(__MODULE__, {:mute_user, room_id, user_id, duration})
  end

  @doc """
  Unmute a user in a chat room.
  """
  def unmute_user(room_id, user_id) do
    GenServer.call(__MODULE__, {:unmute_user, room_id, user_id})
  end

  @doc """
  Ban a user from a chat room.
  """
  def ban_user(room_id, user_id, duration, reason) do
    GenServer.call(__MODULE__, {:ban_user, room_id, user_id, duration, reason})
  end

  @doc """
  Get all active rooms.
  """
  def get_all_rooms do
    GenServer.call(__MODULE__, :get_all_rooms)
  end

  # Server Callbacks

  @impl true
  def init(opts) do
    # Initialize room state from Redis if available
    rooms = %{}
    tournaments = %{}
    user_rooms = %{}
    muted = %{}
    banned = %{}

    # Schedule periodic cleanup
    schedule_cleanup()

    {:ok, %{
      rooms: rooms,
      tournaments: tournaments,
      user_rooms: user_rooms,
      muted: muted,
      banned: banned
    }}
  end

  @impl true
  def handle_call({:join_room, room_id, user_id, username, max_players}, _from, state) do
    # Check if user has reached max rooms
    user_room_count = Map.get(state.user_rooms, user_id, []) |> length()

    if user_room_count >= @max_rooms_per_user do
      {:reply, {:error, :max_rooms_reached}, state}
    else
      # Create room if not exists
      room = get_or_create_room(state.rooms, room_id, max_players)

      # Check if room is full
      if length(room.players) >= max_players do
        {:reply, {:error, :room_full}, state}
      else
        # Add player to room
        new_player = %{user_id: user_id, username: username, ready: false, joined_at: DateTime.utc_now()}
        updated_room = %{room | players: [new_player | room.players]}

        # Update user rooms mapping
        user_rooms = Map.update(state.user_rooms, user_id, [room_id], fn rooms ->
          [room_id | rooms]
        end)

        new_rooms = Map.put(state.rooms, room_id, updated_room)

        {:reply, {:ok, to_map(updated_room)}, %{state | rooms: new_rooms, user_rooms: user_rooms}}
      end
    end
  end

  @impl true
  def handle_call({:leave_room, room_id, user_id}, _from, state) do
    case Map.get(state.rooms, room_id) do
      nil ->
        {:reply, {:error, :room_not_found}, state}

      room ->
        updated_players = Enum.reject(room.players, fn p -> p.user_id == user_id end)

        # Remove room if empty
        {new_rooms, new_user_rooms} = if Enum.empty?(updated_players) do
          {Map.delete(state.rooms, room_id), Map.delete(state.user_rooms, user_id)}
        else
          updated_room = %{room | players: updated_rooms}
          {Map.put(state.rooms, room_id, updated_room), state.user_rooms}
        end

        # Update user rooms mapping
        new_user_rooms = Map.update!(new_user_rooms, user_id, fn rooms ->
          List.delete(rooms, room_id)
        end)

        {:reply, :ok, %{state | rooms: new_rooms, user_rooms: new_user_rooms}}
    end
  end

  @impl true
  def handle_call({:get_room_state, room_id}, _from, state) do
    case Map.get(state.rooms, room_id) do
      nil -> {:reply, nil, state}
      room -> {:reply, to_map(room), state}
    end
  end

  @impl true
  def handle_call({:set_ready, room_id, user_id}, _from, state) do
    case Map.get(state.rooms, room_id) do
      nil ->
        {:reply, {:error, :room_not_found}, state}

      room ->
        updated_players = Enum.map(room.players, fn p ->
          if p.user_id == user_id do
            %{p | ready: true}
          else
            p
          end
        end)

        updated_room = %{room | players: updated_players}
        new_rooms = Map.put(state.rooms, room_id, updated_room)

        {:reply, :ok, %{state | rooms: new_rooms}}
    end
  end

  @impl true
  def handle_call({:handle_disconnect, user_id}, _from, state) do
    # Get all rooms user is in
    user_room_ids = Map.get(state.user_rooms, user_id, [])

    # Remove user from all rooms
    {new_rooms, _} = Enum.reduce(user_room_ids, {state.rooms, state}, fn room_id, {rooms, _} ->
      case Map.get(rooms, room_id) do
        nil -> {rooms, state}
        room ->
          updated_players = Enum.reject(room.players, fn p -> p.user_id == user_id end)

          if Enum.empty?(updated_players) do
            {Map.delete(rooms, room_id), state}
          else
            {Map.put(rooms, room_id, %{room | players: updated_players}), state}
          end
      end
    end)

    new_state = state
      |> Map.put(:rooms, new_rooms)
      |> Map.put(:user_rooms, Map.delete(state.user_rooms, user_id))

    {:reply, :ok, new_state}
  end

  @impl true
  def handle_call({:join_tournament, tournament_id, user_id, username, max_participants}, _from, state) do
    # Get or create tournament
    tournament = get_or_create_tournament(state.tournaments, tournament_id, max_participants)

    if length(tournament.participants) >= max_participants do
      {:reply, {:error, :tournament_full}, state}
    else
      participant = %{user_id: user_id, username: username, status: :registered, joined_at: DateTime.utc_now()}
      updated_tournament = %{tournament | participants: [participant | tournament.participants]}

      new_tournaments = Map.put(state.tournaments, tournament_id, updated_tournament)

      {:reply, {:ok, to_map(updated_tournament)}, %{state | tournaments: new_tournaments}}
    end
  end

  @impl true
  def handle_call({:leave_tournament, tournament_id, user_id}, _from, state) do
    case Map.get(state.tournaments, tournament_id) do
      nil ->
        {:reply, {:error, :tournament_not_found}, state}

      tournament ->
        updated_participants = Enum.reject(tournament.participants, fn p -> p.user_id == user_id end)

        updated_tournament = %{tournament | participants: updated_participants}
        new_tournaments = Map.put(state.tournaments, tournament_id, updated_tournament)

        {:reply, :ok, %{state | tournaments: new_tournaments}}
    end
  end

  @impl true
  def handle_call({:register_player, tournament_id, user_id, username, buy_in}, _from, state) do
    case Map.get(state.tournaments, tournament_id) do
      nil ->
        {:reply, {:error, :tournament_not_found}, state}

      tournament ->
        # Check if already registered
        if Enum.any?(tournament.participants, fn p -> p.user_id == user_id end) do
          {:reply, {:error, :already_registered}, state}
        else
          participant = %{user_id: user_id, username: username, buy_in: buy_in, status: :registered, registered_at: DateTime.utc_now()}
          updated_tournament = %{tournament | participants: [participant | tournament.participants]}

          new_tournaments = Map.put(state.tournaments, tournament_id, updated_tournament)

          {:reply, {:ok, participant}, %{state | tournaments: new_tournaments}}
        end
    end
  end

  @impl true
  def handle_call({:unregister_player, tournament_id, user_id}, _from, state) do
    case Map.get(state.tournaments, tournament_id) do
      nil ->
        {:reply, {:error, :tournament_not_found}, state}

      tournament ->
        updated_participants = Enum.reject(tournament.participants, fn p -> p.user_id == user_id end)

        updated_tournament = %{tournament | participants: updated_participants}
        new_tournaments = Map.put(state.tournaments, tournament_id, updated_tournament)

        {:reply, :ok, %{state | tournaments: new_tournaments}}
    end
  end

  @impl true
  def handle_call({:mute_user, room_id, user_id, duration}, _from, state) do
    muted_key = "#{room_id}:#{user_id}"
    muted_until = DateTime.add(DateTime.utc_now(), duration, :second)

    new_muted = Map.put(state.muted, muted_key, muted_until)

    {:reply, :ok, %{state | muted: new_muted}}
  end

  @impl true
  def handle_call({:unmute_user, room_id, user_id}, _from, state) do
    muted_key = "#{room_id}:#{user_id}"
    new_muted = Map.delete(state.muted, muted_key)

    {:reply, :ok, %{state | muted: new_muted}}
  end

  @impl true
  def handle_call({:ban_user, room_id, user_id, duration, reason}, _from, state) do
    banned_key = "#{room_id}:#{user_id}"
    banned_until = DateTime.add(DateTime.utc_now(), duration, :second)

    new_banned = Map.put(state.banned, banned_key, %{until: banned_until, reason: reason})

    {:reply, :ok, %{state | banned: new_banned}}
  end

  @impl true
  def handle_call(:get_all_rooms, _from, state) do
    {:reply, Enum.map(state.rooms, fn {id, room} -> to_map(room) end), state}
  end

  @impl true
  def handle_info(:cleanup, state) do
    # Clean up idle rooms
    now = DateTime.utc_now()

    cleaned_rooms = Enum.reject(state.rooms, fn {_id, room} ->
      idle_time = DateTime.diff(now, room.updated_at, :millisecond)
      idle_time > @room_idle_timeout and Enum.empty?(room.players)
    end)
    |> Map.new()

    # Clean up expired mutes/bans
    cleaned_muted = Enum.reject(state.muted, fn {_key, until} ->
      DateTime.compare(now, until) == :gt
    end)
    |> Map.new()

    cleaned_banned = Enum.reject(state.banned, fn {_key, %{until: until}} ->
      DateTime.compare(now, until) == :gt
    end)
    |> Map.new()

    # Schedule next cleanup
    schedule_cleanup()

    {:noreply, %{state |
      rooms: cleaned_rooms,
      muted: cleaned_muted,
      banned: cleaned_banned
    }}
  end

  # Private functions

  defp get_or_create_room(rooms, room_id, max_players) do
    Map.get(rooms, room_id) || %{
      id: room_id,
      players: [],
      max_players: max_players,
      status: :waiting,
      created_at: DateTime.utc_now(),
      updated_at: DateTime.utc_now()
    }
  end

  defp get_or_create_tournament(tournaments, tournament_id, max_participants) do
    Map.get(tournaments, tournament_id) || %{
      id: tournament_id,
      participants: [],
      max_participants: max_participants,
      status: :open,
      created_at: DateTime.utc_now(),
      updated_at: DateTime.utc_now()
    }
  end

  defp to_map(struct) do
    Map.from_struct(struct)
  end

  defp schedule_cleanup do
    Process.send_after(self(), :cleanup, @room_cleanup_interval)
  end
end
