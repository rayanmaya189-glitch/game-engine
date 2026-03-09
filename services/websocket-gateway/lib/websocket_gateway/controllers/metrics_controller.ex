defmodule WebsocketGateway.MetricsController do
  use WebsocketGateway, :controller

  def index(conn, _params) do
    # Return prometheus-formatted metrics
    metrics = [
      "# TYPE phoenix_socket_connected counter",
      "phoenix_socket_connected_total 0",
      "# TYPE websocket_gateway_rooms_active gauge",
      "websocket_gateway_rooms_active 0",
      "# TYPE websocket_gateway_players_online gauge",
      "websocket_gateway_players_online 0"
    ]

    conn
    |> put_resp_content_type("text/plain")
    |> send_resp(200, Enum.join(metrics, "\n"))
  end
end
