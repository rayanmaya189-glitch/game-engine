defmodule WebsocketGateway.Plugs.RequestLogger do
  @moduledoc """
  Request logging plug.
  """

  import Plug.Conn
  require Logger

  def init(opts), do: opts

  def call(conn, _opts) do
    start_time = System.monotonic_time(:millisecond)

    conn
    |> register_before_send(fn conn ->
      stop_time = System.monotonic_time(:millisecond)
      duration = stop_time - start_time

      Logger.info(
        "[REQUEST] #{conn.method} #{conn.request_path} - #{conn.status} - #{duration}ms",
        request_id: get_session(conn, :request_id) || "unknown",
        ip: to_string(:inet_parse.ntoa(conn.remote_ip)),
        method: conn.method,
        path: conn.request_path,
        status: conn.status,
        duration: duration
      )

      conn
    end)
  end
end
