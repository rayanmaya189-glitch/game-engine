defmodule WebsocketGateway.AdminController do
  use WebsocketGateway, :controller

  alias WebsocketGateway.Services.{Presence, RoomManager}

  def stats(conn, _params) do
    json(conn, %{
      timestamp: DateTime.utc_now(),
      connections: %{
        total: Presence.global_online_count(),
        lobby: Presence.count("lobby"),
        games: Presence.count("game:*"),
        chat: Presence.count("chat:*"),
        tournament: Presence.count("tournament:*")
      },
      rooms: %{
        active: length(RoomManager.get_all_rooms())
      },
      system: %{
        node: Node.self(),
        uptime: get_uptime(),
        memory: :erlang.memory()
      }
    })
  end

  defp get_uptime do
    # Get system uptime in seconds
    {uptime, _} = :erlang.statistics(:wall_clock)
    div(uptime, 1000)
  end
end
