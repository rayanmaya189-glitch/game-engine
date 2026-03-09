defmodule WebsocketGateway.Plugs.HealthCheck do
  @moduledoc """
  Health check plug for Kubernetes liveness/readiness probes.
  """

  import Plug.Conn

  def init(opts), do: opts

  def call(conn, _opts) do
    if conn.path_info == ["health"] || conn.path_info == ["health", "live"] do
      conn
      |> put_resp_content_type("application/json")
      |> send_resp(200, Jason.encode!(%{status: "ok", timestamp: DateTime.utc_now()}))
      |> halt()
    else
      conn
    end
  end
end
