defmodule WebsocketGateway.MixProject do
  use Mix.Project

  def project do
    [
      app: :websocket_gateway,
      version: "0.1.0",
      elixir: "~> 1.15",
      elixirc_paths: elixirc_paths(Mix.env()),
      start_permanent: Mix.env() == :prod,
      aliases: aliases(),
      deps: deps()
    ]
  end

  # Configuration for the OTP application.
  #
  # Type `mix help compile.app` for more information.
  def application do
    [
      mod: {WebsocketGateway.Application, []},
      extra_applications: [:logger, :runtime_tools]
    ]
  end

  # Specifies which paths to compile per environment.
  defp elixirc_paths(:test), do: ["lib", "test/support"]
  defp elixirc_paths(:dev), do: ["lib"]
  defp elixirc_paths(:prod), do: ["lib"]

  # Dependencies can be Hex packages:
  #
  #   {:phoenix, "~> 1.7.0"}
  #
  # Or git/path repositories:
  #
  #   {:phoenix, git: "https://github.com/phoenixframework/phoenix.git", tag: "v1.7.0"}
  #
  # Type `mix help deps` for examples and options.
  defp deps do
    [
      # Phoenix framework
      {:phoenix, "~> 1.7.0"},
      {:phoenix_pubsub, "~> 2.1.3"},
      {:phoenix_presence_tracking, "~> 0.1.0"},
      {:phoenix_html, "~> 3.3.0"},

      # WebSocket
      {:phoenix_live_reload, "~> 1.5.0", only: :dev},
      {:phoenix_live_view, "~> 0.20.0", optional: true},

      # Database
      {:ecto_sql, "~> 3.10"},
      {:postgrex, ">= 0.0.0"},

      # Redis
      {:redix, ">= 0.0.0"},

      # JWT
      {:jose, "~> 1.jwt_ver11"},
      {:ifier, "~> 0.2.0"},

      # NATS
      {:gnat, "~> 1.5"},

      # gRPC (inter-service communication)
      {:grpc, "~> 0.9"},
      {:protobuf, "~> 0.13"},
      {:cowboy, "~> 2.10"},

      # Monitoring
      {:telemetry_metrics, "~> 0.6"},
      {:telemetry_poller, "~> 1.0"},

      # UUID
      {:uuid, "~> 2.0"},

      # Config
      {:distillery, "~> 2.1", runtime: false},
      {:config_parser_ex, "~> 4.0"}
    ]
  end

  # Aliases are shortcuts or tasks specific to the current project.
  # For example, to create default migrations run:
  #
  #     mix ecto.create
  #
  # See the documentation for `Mix` for more info on aliases.
  defp aliases do
    [
      setup: ["deps.get", "ecto.setup"],
      "ecto.setup": ["ecto.create", "ecto.migrate", "run priv/repo/seeds.exs"],
      "ecto.reset": ["ecto.drop", "ecto.setup"],
      test: ["ecto.create --quiet", "ecto.migrate --quiet", "test"]
    ]
  end
end
