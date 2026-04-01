defmodule WebsocketGateway.Services.RoomLifecycle do
  @moduledoc """
  Handles room lifecycle operations: cleanup, idle timeout detection,
  and expired mute/ban pruning for the RoomManager.
  """

  @room_idle_timeout Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:room_idle_timeout, 300_000)

  @room_cleanup_interval Application.get_env(:websocket_gateway, :room_manager, [])
    |> Keyword.get(:room_cleanup_interval, 60_000)

  @doc """
  Clean up idle rooms, expired mutes, and expired bans from state.
  Returns the updated state map.
  """
  def cleanup(state) do
    now = DateTime.utc_now()

    cleaned_rooms =
      Enum.reject(state.rooms, fn {_id, room} ->
        idle_time = DateTime.diff(now, room.updated_at, :millisecond)
        idle_time > @room_idle_timeout and Enum.empty?(room.players)
      end)
      |> Map.new()

    cleaned_muted =
      Enum.reject(state.muted, fn {_key, until} ->
        DateTime.compare(now, until) == :gt
      end)
      |> Map.new()

    cleaned_banned =
      Enum.reject(state.banned, fn {_key, %{until: until}} ->
        DateTime.compare(now, until) == :gt
      end)
      |> Map.new()

    %{state | rooms: cleaned_rooms, muted: cleaned_muted, banned: cleaned_banned}
  end

  @doc """
  Schedule the next cleanup tick to be sent to the calling process.
  """
  def schedule_cleanup do
    Process.send_after(self(), :cleanup, @room_cleanup_interval)
  end
end
