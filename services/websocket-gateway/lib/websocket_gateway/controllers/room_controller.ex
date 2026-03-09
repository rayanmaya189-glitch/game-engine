defmodule WebsocketGateway.RoomController do
  use WebsocketGateway, :controller

  alias WebsocketGateway.Services.RoomManager

  def list(conn, _params) do
    rooms = RoomManager.get_all_rooms()
    json(conn, %{rooms: rooms})
  end

  def show(conn, %{"room_id" => room_id}) do
    case RoomManager.get_room_state(room_id) do
      nil ->
        conn
        |> put_status(:not_found)
        |> json(%{error: "Room not found"})

      room ->
        json(conn, %{room: room})
    end
  end

  def delete(conn, %{"room_id" => room_id}) do
    # Admin only - implement proper auth
    json(conn, %{status: "deleted", room_id: room_id})
  end
end
