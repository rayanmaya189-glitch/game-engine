defmodule WebsocketGateway.Telemetry do
  @moduledoc """
  Telemetry instrumentation for metrics collection.
  """

  use Supervisor
  import Telemetry.Metrics

  def start_link(opts) do
    Supervisor.start_link(__MODULE__, opts, name: __MODULE__)
  end

  @impl true
  def init(_opts) do
    children = [
      {:telemetry_poller, measurements: periodic_measurements(), period: 10_000}
    ]

    Supervisor.init(children, strategy: :one_for_one)
  end

  def metrics do
    [
      # Phoenix metrics
      summary("phoenix.endpoint.start.system_time",
        unit: {:native, :millisecond}
      ),
      summary("phoenix.endpoint.stop.duration",
        unit: {:native, :millisecond}
      ),
      summary("phoenix.router_dispatch.start.system_time",
        tags: [:route],
        unit: {:native, :millisecond}
      ),
      summary("phoenix.router_dispatch.exception.duration",
        tags: [:route],
        unit: {:native, :millisecond}
      ),
      summary("phoenix.router_dispatch.start.system_time",
        tags: [:route],
        unit: {:native, :millisecond}
      ),
      summary("phoenix.socket_connected.duration",
        unit: {:native, :millisecond}
      ),
      summary("phoenix.channel_join.duration",
        unit: {:native, :millisecond}
      ),
      summary("phoenix.channel_handled_in.duration",
        tags: [:event],
        unit: {:native, :millisecond}
      ),

      # Database metrics
      summary("websocket_gateway.repo.query.total_time",
        unit: {:native, :millisecond},
        description: "The sum of the other database metrics"
      ),
      summary("websocket_gateway.repo.query.decode_time",
        unit: {:native, :millisecond},
        description: "The time spent decoding the data received from the database"
      ),
      summary("websocket_gateway.repo.query.query_time",
        unit: {:native, :millisecond},
        description: "The time spent executing the query"
      ),
      summary("websocket_gateway.repo.query.queue_time",
        unit: {:native, :millisecond},
        description: "The time spent waiting for a database connection"
      ),
      summary("websocket_gateway.repo.query.idle_time",
        unit: {:native, :millisecond},
        description: "The time the connection spent waiting before being checked out for the query"
      ),

      # VM metrics
      summary("vm.memory.total", unit: {:byte, :kilobyte}),
      summary("vm.memory.processes", unit: {:byte, :kilobyte}),
      summary("vm.memory.binary", unit: {:byte, :kilobyte}),
      summary("vm.memory.ets", unit: {:byte, :kilobyte}),
      summary("vm.run_queue.total"),
      summary("vm.process_count"),
      summary("vm.atom_count"),
      summary("vm.port_count"),

      # Connection metrics
      counter("phoenix.socket_connected.total"),
      counter("phoenix.socket_disconnected.total"),
      counter("phoenix.channel_join.total"),

      # WebSocket metrics
      summary("websocket_gateway.websocket.connect.duration",
        unit: {:native, :millisecond}
      ),
      counter("websocket_gateway.websocket.connect.total"),
      counter("websocket_gateway.websocket.disconnect.total"),
      counter("websocket_gateway.websocket.message.total")
    ]
  end

  defp periodic_measurements do
    []
  end
end

defmodule WebsocketGateway.Telemetry.Endpoint do
  use Phoenix.Endpoint, otp_app: :websocket_gateway

  @impl true
  def init(_opts, config) do
    {:ok, config}
  end
end
