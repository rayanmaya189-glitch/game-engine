defmodule WebsocketGateway.NATS.Client do
  @moduledoc """
  NATS client for subscribing to game, tournament, and jackpot events.
  """

  use GenServer
  require Logger

  alias Phoenix.PubSub

  @reconnect_interval 5_000

  def start_link(opts \\ []) do
    GenServer.start_link(__MODULE__, opts, name: __MODULE__)
  end

  @doc """
  Publish a message to a NATS subject.
  """
  def publish(subject, payload) do
    case :gnat.pub(:websocket_gateway_nats, subject, Jason.encode!(payload)) do
      :ok -> {:ok, "published"}
      error -> {:error, error}
    end
  end

  @doc """
  Request a response from a NATS subject (request-reply pattern).
  Returns {:ok, response_map} or {:error, reason}
  """
  def request(subject, payload, timeout \\ 5000) do
    encoded = Jason.encode!(payload)

    case :gnat.req(:websocket_gateway_nats, subject, encoded, request_timeout: timeout) do
      {:ok, %{body: body}} ->
        case Jason.decode(body) do
          {:ok, response} -> {:ok, response}
          {:error, _} -> {:error, :decode_error}
        end
      {:error, :timeout} ->
        {:error, :timeout}
      {:error, reason} ->
        {:error, reason}
    end
  rescue
    e in Jason.EncodeError -> {:error, :encode_error}
    e in ArgumentError -> {:error, :invalid_request}
    _ -> {:error, :unknown}
  end

  @impl true
  def init(opts) do
    config = Application.get_env(:websocket_gateway, :nats, [])

    state = %{
      config: config,
      connection: nil,
      subscriptions: []
    }

    # Connect to NATS
    {:ok, state, {:continue, :connect}}
  end

  @impl true
  def handle_continue(:connect, state) do
    case connect(state.config) do
      {:ok, conn} ->
        subscriptions = subscribe_all(state.config)
        Logger.info("Connected to NATS")
        {:noreply, %{state | connection: conn, subscriptions: subscriptions}}

      {:error, reason} ->
        Logger.warning("Failed to connect to NATS: #{inspect(reason)}, retrying...")
        Process.send_after(self(), :retry_connect, @reconnect_interval)
        {:noreply, state}
    end
  end

  @impl true
  def handle_info(:retry_connect, state) do
    {:noreply, state, {:continue, :connect}}
  end

  @impl true
  def handle_info({:nats, _msg}, state) do
    # Handle incoming NATS messages
    {:noreply, state}
  end

  @impl true
  def terminate(_reason, state) do
    if state.connection do
      Gnat.close(state.connection)
    end
    :ok
  end

  # Private functions

  defp connect(config) do
    host = Keyword.get(config, :host, "localhost")
    port = Keyword.get(config, :port, 4222)
    username = Keyword.get(config, :username, "")
    password = Keyword.get(config, :password, "")
    timeout = Keyword.get(config, :connect_timeout, 5000)

    connection_opts = [
      host: host,
      port: port,
      name: :websocket_gateway_nats
    ]
    |> maybe_add_auth(username, password)

    Gnat.start_link(connection_opts)
  end

  defp maybe_add_auth(opts, "", ""), do: opts
  defp maybe_add_auth(opts, username, password) do
    Keyword.put(opts, :auth, {String.to_charlist(username), String.to_charlist(password)})
  end

  defp subscribe_all(config) do
    subscriptions = Keyword.get(config, :subscriptions, [])

    Enum.map(subscriptions, fn pattern ->
      topic = pattern
      |> String.replace("*", "**")

      {:ok, sid} = Gnat.sub(:websocket_gateway_nats, self(), topic)
      %{sid: sid, topic: topic}
    end)
  end
end
