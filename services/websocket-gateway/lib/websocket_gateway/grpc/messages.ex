defmodule WebsocketGateway.GRPC.Messages do
  @moduledoc """
  Protobuf message definitions for gRPC service communication.

  These define the wire-format messages matching the proto definitions
  in services/common-service/proto/game_engine/.
  """

  # =========================================================================
  # Common messages
  # =========================================================================

  defmodule Money do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :amount, 1, type: :int64
    field :currency, 2, type: :int32
    field :display_amount, 3, type: :string
  end

  # =========================================================================
  # Wallet Service messages
  # =========================================================================

  defmodule Wallet.GetBalanceRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :user_id, 1, type: :string, json_name: "userId"
    field :balance_type, 2, type: :int32, json_name: "balanceType"
  end

  defmodule Wallet.GetBalanceResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :balance, 1, type: WebsocketGateway.GRPC.Messages.Money
    field :locked_amount, 2, type: WebsocketGateway.GRPC.Messages.Money, json_name: "lockedAmount"
    field :available_amount, 3, type: WebsocketGateway.GRPC.Messages.Money, json_name: "availableAmount"
    field :bonus_amount, 4, type: WebsocketGateway.GRPC.Messages.Money, json_name: "bonusAmount"
  end

  defmodule Wallet.PlaceBetRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :user_id, 1, type: :string, json_name: "userId"
    field :game_id, 2, type: :string, json_name: "gameId"
    field :amount, 3, type: WebsocketGateway.GRPC.Messages.Money
    field :bet_type, 4, type: :string, json_name: "betType"
    field :selection, 5, type: :string
    field :odds, 6, type: :string
  end

  defmodule Wallet.PlaceBetResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :bet, 1, type: :map
    field :new_balance, 2, type: WebsocketGateway.GRPC.Messages.Money, json_name: "newBalance"
    field :locked_amount, 3, type: WebsocketGateway.GRPC.Messages.Money, json_name: "lockedAmount"
    field :bet_id, 4, type: :string, json_name: "betId"
    field :message, 5, type: :string
  end

  defmodule Wallet.SettleBetRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :bet_id, 1, type: :string, json_name: "betId"
    field :settlement_type, 2, type: :int32, json_name: "settlementType"
    field :win_amount, 3, type: WebsocketGateway.GRPC.Messages.Money, json_name: "winAmount"
    field :result, 4, type: :string
  end

  defmodule Wallet.SettleBetResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :success, 1, type: :bool
    field :new_balance, 4, type: WebsocketGateway.GRPC.Messages.Money, json_name: "newBalance"
    field :message, 5, type: :string
  end

  # =========================================================================
  # Game Registry Service messages
  # =========================================================================

  defmodule Game.GetGameRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :game_id, 1, type: :string, json_name: "gameId"
  end

  defmodule Game.GetGameResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :game, 1, type: :map
  end

  defmodule Game.SearchGamesRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :query, 1, type: :string
    field :limit, 2, type: :int32
    field :category_id, 3, type: :string, json_name: "categoryId"
  end

  defmodule Game.SearchGamesResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :games, 1, repeated: true, type: :map
    field :total_count, 2, type: :int32, json_name: "totalCount"
  end

  defmodule Game.GetFeaturedGamesRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :limit, 1, type: :int32
    field :category_id, 2, type: :string, json_name: "categoryId"
  end

  defmodule Game.GetFeaturedGamesResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :games, 1, repeated: true, type: :map
  end

  defmodule Game.GetPopularGamesRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :limit, 1, type: :int32
    field :category_id, 2, type: :string, json_name: "categoryId"
  end

  defmodule Game.GetPopularGamesResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :games, 1, repeated: true, type: :map
  end

  defmodule Game.GetNewGamesRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :limit, 1, type: :int32
    field :category_id, 2, type: :string, json_name: "categoryId"
  end

  defmodule Game.GetNewGamesResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :games, 1, repeated: true, type: :map
  end

  defmodule Game.GetCategoriesRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :include_games_count, 1, type: :bool, json_name: "includeGamesCount"
  end

  defmodule Game.GetCategoriesResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :categories, 1, repeated: true, type: :map
  end

  # =========================================================================
  # Tournament Service messages
  # =========================================================================

  defmodule Tournament.GetTournamentRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :tournament_id, 1, type: :string, json_name: "tournamentId"
  end

  defmodule Tournament.GetTournamentResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :tournament, 1, type: :map
  end

  defmodule Tournament.ListTournamentsRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :status, 1, type: :string
    field :page, 2, type: :int32
    field :limit, 3, type: :int32
  end

  defmodule Tournament.ListTournamentsResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :tournaments, 1, repeated: true, type: :map
    field :total, 2, type: :int32
  end

  defmodule Tournament.GetLeaderboardRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :tournament_id, 1, type: :string, json_name: "tournamentId"
    field :limit, 2, type: :int32
  end

  defmodule Tournament.GetLeaderboardResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :entries, 1, repeated: true, type: :map
  end

  # =========================================================================
  # Jackpot Service messages
  # =========================================================================

  defmodule Jackpot.GetJackpotRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :jackpot_id, 1, type: :string, json_name: "jackpotId"
  end

  defmodule Jackpot.GetJackpotResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :jackpot, 1, type: :map
  end

  defmodule Jackpot.ListJackpotsRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :status, 1, type: :string
  end

  defmodule Jackpot.ListJackpotsResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :jackpots, 1, repeated: true, type: :map
  end

  # =========================================================================
  # Auth Service messages
  # =========================================================================

  defmodule Auth.ValidateTokenRequest do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :token, 1, type: :string
    field :expected_type, 2, type: :string, json_name: "expectedType"
  end

  defmodule Auth.ValidateTokenResponse do
    @moduledoc false
    use Protobuf, syntax: :proto3, protoc_gen_elixir_version: "0.13.0"

    field :valid, 1, type: :bool
    field :user_id, 2, type: :string, json_name: "userId"
    field :session_id, 3, type: :string, json_name: "sessionId"
    field :token_type, 4, type: :string, json_name: "tokenType"
    field :expires_at, 5, type: :int64, json_name: "expiresAt"
    field :roles, 6, repeated: true, type: :int32
  end
end
