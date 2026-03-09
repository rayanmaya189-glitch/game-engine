defmodule WebsocketGateway.Repo do
  use Ecto.Repo,
    otp_app: :websocket_gateway,
    adapter: Ecto.Adapters.Postgres
end
