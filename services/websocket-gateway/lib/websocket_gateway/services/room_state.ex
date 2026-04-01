defmodule WebsocketGateway.Services.RoomState do
  @moduledoc """
  Pure functions for querying and broadcasting room/tournament state.
  Used by RoomManager to keep its own module focused on GenServer callbacks.
  """

  @doc """
  Get a room map by id, or nil if not found.
  """
  def get_room(rooms, room_id) do
    case Map.get(rooms, room_id) do
      nil -> nil
      room -> to_map(room)
    end
  end

  @doc """
  Return a list of all rooms as plain maps.
  """
  def get_all_rooms(rooms) do
    Enum.map(rooms, fn {_id, room} -> to_map(room) end)
  end

  @doc """
  Get or create a room template (not yet stored in state).
  """
  def get_or_create_room(rooms, room_id, max_players) do
    Map.get(rooms, room_id) || %{
      id: room_id,
      players: [],
      max_players: max_players,
      status: :waiting,
      created_at: DateTime.utc_now(),
      updated_at: DateTime.utc_now()
    }
  end

  @doc """
  Get or create a tournament template.
  """
  def get_or_create_tournament(tournaments, tournament_id, max_participants) do
    Map.get(tournaments, tournament_id) || %{
      id: tournament_id,
      participants: [],
      max_participants: max_participants,
      status: :open,
      created_at: DateTime.utc_now(),
      updated_at: DateTime.utc_now()
    }
  end

  @doc """
  Convert a struct (or map) to a plain map, stripping internal fields.
  """
  def to_map(struct) do
    Map.from_struct(struct)
  end
end
