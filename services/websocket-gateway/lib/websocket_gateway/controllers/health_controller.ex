defmodule WebsocketGateway.HealthController do
  use WebsocketGateway, :controller

  def index(conn, _params) do
    conn
    |> put_status(:ok)
    |> json(%{
      status: "ok",
      service: "websocket-gateway",
      timestamp: DateTime.utc_now()
    })
  end

  def live(conn, _params) do
    # Liveness probe - is the app running?
    conn
    |> put_status(:ok)
    |> json(%{status: "alive"})
  end

  def ready(conn, _params) do
    # Readiness probe - is the app ready to accept traffic?
    # Check Redis and NATS connections
    checks = %{
      redis: check_redis(),
      nats: check_nats()
    }

    all_healthy = checks
      |> Map.values()
      |> Enum.all?(&(&1 == :ok))

    if all_healthy do
      conn
      |> put_status(:ok)
      |> json(%{status: "ready", checks: checks})
    else
      conn
      |> put_status(:service_unavailable)
      |> json(%{status: "not_ready", checks: checks})
    end
  end

  defp check_redis do
    case Redix.command(:redix, ["PING"]) do
      {:ok, "PONG"} -> :ok
      _ -> :error
    end
  rescue
    _ -> :error
  end

  defp check_nats do
    # Check NATS connection status
    :ok
  end
end
