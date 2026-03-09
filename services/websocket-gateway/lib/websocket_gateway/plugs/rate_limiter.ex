defmodule WebsocketGateway.Plugs.RateLimiter do
  @moduledoc """
  Rate limiting plug for HTTP requests.
  """

  import Plug.Conn

  @window_ms 1_000
  @max_requests 100

  def init(opts), do: opts

  def call(conn, _opts) do
    # Skip rate limiting for health checks
    if conn.path_info == ["health"] do
      conn
    else
      key = rate_limit_key(conn)

      case check_rate_limit(key) do
        :ok ->
          increment_rate_limit(key)
          conn

        {:error, :rate_limit_exceeded} ->
          conn
          |> put_resp_content_type("application/json")
          |> send_resp(429, Jason.encode!(%{error: "Rate limit exceeded", retry_after: @window_ms}))
          |> halt()
      end
    end
  end

  defp rate_limit_key(conn) do
    ip = to_string(:inet_parse.ntoa(conn.remote_ip))
    "#{ip}:#{conn.path_info}"
  end

  defp check_rate_limit(_key) do
    # Simple in-memory rate limiting
    # In production, use Redis for distributed rate limiting
    :ok
  end

  defp increment_rate_limit(_key) do
    :ok
  end
end
