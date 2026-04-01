defmodule WebsocketGateway.Services.RoomManager do
  @moduledoc """
  Room Manager service for managing game rooms.

  Provides functionality for:
  - Creating/joining/leaving game rooms
  - Managing room state
  - Player management
  """

  use GenServer
  require Logger

  alias WebsocketGateway.Services.RoomLifecycle
  alias WebsocketGateway.Services.RoomState

  @max_rooms_per_user Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:max_rooms_per_user, 10)

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
  Get all active rooms.
  """
  def get_all_rooms do
    GenServer.call(__MODULE__, :get_all_rooms)
  end

  # Server Callbacks

  @impl true
  def init(opts) do
    rooms = %{}
    tournaments = %{}
    user_rooms = %{}
    muted = %{}
    banned = %{}

    RoomLifecycle.schedule_cleanup()

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
    user_room_count = Map.get(state.user_rooms, user_id, []) |> length()

    if user_room_count >= @max_rooms_per_user do
      {:reply, {:error, :max_rooms_reached}, state}
    else
      room = RoomState.get_or_create_room(state.rooms, room_id, max_players)

      if length(room.players) >= max_players do
        {:reply, {:error, :room_full}, state}
      else
        new_player = %{user_id: user_id, username: username, ready: false, joined_at: DateTime.utc_now()}
        updated_room = %{room | players: [new_player | room.players]}

        user_rooms = Map.update(state.user_rooms, user_id, [room_id], fn rooms ->
          [room_id | rooms]
        end)

        new_rooms = Map.put(state.rooms, room_id, updated_room)

        {:reply, {:ok, RoomState.to_map(updated_room)}, %{state | rooms: new_rooms, user_rooms: user_rooms}}
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

        {new_rooms, new_user_rooms} = if Enum.empty?(updated_players) do
          {Map.delete(state.rooms, room_id), Map.delete(state.user_rooms, user_id)}
        else
          updated_room = %{room | players: updated_players}
          {Map.put(state.rooms, room_id, updated_room), state.user_rooms}
        end

        new_user_rooms = Map.update!(new_user_rooms, user_id, fn rooms ->
          List.delete(rooms, room_id)
        end)

        {:reply, :ok, %{state | rooms: new_rooms, user_rooms: new_user_rooms}}
    end
  end

  @impl true
  def handle_call({:get_room_state, room_id}, _from, state) do
    {:reply, RoomState.get_room(state.rooms, room_id), state}
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
    user_room_ids = Map.get(state.user_rooms, user_id, [])

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
  def handle_call(:get_all_rooms, _from, state) do
    {:reply, RoomState.get_all_rooms(state.rooms), state}
  end

  @impl true
  def handle_info(:cleanup, state) do
    RoomLifecycle.schedule_cleanup()
    {:noreply, RoomLifecycle.cleanup(state)}
  end
end
