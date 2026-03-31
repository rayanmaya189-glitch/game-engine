defmodule WebsocketGateway.GRPC.ChannelManager do
  @moduledoc """
  GenServer that manages gRPC channel connections to backend services.

  Maintains persistent gRPC channels for each service and provides
  channel lookup. Channels are created on startup from application config
  and can be reconnected on failure.
  """

  use GenServer

  require Logger

  @service_keys [
    :wallet_service,
    :game_registry_service,
    :tournament_service,
    :jackpot_service,
    :auth_service
  ]

  def start_link(opts \\ []) do
    GenServer.start_link(__MODULE__, opts, name: __MODULE__)
  end

  @doc """
  Get the gRPC channel for a given service key.
  Returns {:ok, channel} or {:error, :not_connected}.
  """
  def get_channel(service_key) when is_atom(service_key) do
    GenServer.call(__MODULE__, {:get_channel, service_key})
  end

  @doc """
  List all connected service channels.
  """
  def list_channels do
    GenServer.call(__MODULE__, :list_channels)
  end

  # GenServer callbacks

  @impl true
  def init(_opts) do
    Process.send_after(self(), :connect_channels, 0)
    {:ok, %{channels: %{}}}
  end

  @impl true
  def handle_info(:connect_channels, state) do
    channels = connect_all_services()
    {:noreply, %{state | channels: channels}}
  end

  @impl true
  def handle_info({:reconnect, service_key}, state) do
    case connect_service(service_key) do
      {:ok, channel} ->
        Logger.info("Reconnected to gRPC service: #{service_key}")
        {:noreply, put_in(state, [:channels, service_key], channel)}

      {:error, reason} ->
        Logger.warning("Failed to reconnect to #{service_key}: #{inspect(reason)}")
        Process.send_after(self(), {:reconnect, service_key}, 5_000)
        {:noreply, state}
    end
  end

  @impl true
  def handle_call({:get_channel, service_key}, _from, state) do
    case Map.get(state.channels, service_key) do
      nil -> {:reply, {:error, :not_connected}, state}
      channel -> {:reply, {:ok, channel}, state}
    end
  end

  @impl true
  def handle_call(:list_channels, _from, state) do
    {:reply, state.channels, state}
  end

  # Private functions

  defp connect_all_services do
    services_config = Application.get_env(:websocket_gateway, :services, [])

    Enum.reduce(@service_keys, %{}, fn key, acc ->
      case connect_service(key, services_config) do
        {:ok, channel} ->
          Map.put(acc, key, channel)

        {:error, reason} ->
          Logger.warning("Failed to connect to gRPC service #{key}: #{inspect(reason)}")
          Process.send_after(self(), {:reconnect, key}, 5_000)
          acc
      end
    end)
  end

  defp connect_service(service_key, services_config \\ nil) do
    services_config = services_config || Application.get_env(:websocket_gateway, :services, [])

    host = Keyword.get(services_config, :"#{service_key}_host", default_host(service_key))
    port = Keyword.get(services_config, :"#{service_key}_port", default_port(service_key))

    opts = [
      cred: GRPC.Credential.new(%{ssl: false})
    ]

    case GRPC.Stub.connect("#{host}:#{port}", opts) do
      {:ok, channel} ->
        Logger.info("Connected to gRPC service #{service_key} at #{host}:#{port}")
        {:ok, channel}

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp default_host(:wallet_service), do: "localhost"
  defp default_host(:game_registry_service), do: "localhost"
  defp default_host(:tournament_service), do: "localhost"
  defp default_host(:jackpot_service), do: "localhost"
  defp default_host(:auth_service), do: "localhost"
  defp default_host(_), do: "localhost"

  defp default_port(:wallet_service), do: 9081
  defp default_port(:game_registry_service), do: 9082
  defp default_port(:tournament_service), do: 9083
  defp default_port(:jackpot_service), do: 9084
  defp default_port(:auth_service), do: 9080
  defp default_port(_), do: 9080
end
