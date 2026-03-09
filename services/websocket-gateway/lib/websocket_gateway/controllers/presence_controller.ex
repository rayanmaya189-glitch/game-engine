defmodule WebsocketGateway.PresenceController do
  use WebsocketGateway, :controller

  alias WebsocketGateway.Services.Presence

  def index(conn, _params) do
    # Get presence stats
    lobby_count = Presence.count("lobby")
    global_count = Presence.global_online_count()

    json(conn, %{
      lobby: lobby_count,
      global: global_count,
      rooms: get_room_presence()
    })
  end

  defp get_room_presence do
    # Get presence by room type
    %{
      games: Presence.count("game:*"),
      chat: Presence.count("chat:*"),
      tournament: Presence.count("tournament:*")
    }
  end
end
