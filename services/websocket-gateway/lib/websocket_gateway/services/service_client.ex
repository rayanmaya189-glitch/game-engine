defmodule WebsocketGateway.Services.ServiceClient do
  @moduledoc """
  gRPC client for communicating with backend services.
  Used by WebSocket channels to call wallet, game, tournament,
  jackpot, and auth services via gRPC.

  All inter-service communication uses gRPC exclusively (no HTTP/REST).
  """

  alias WebsocketGateway.GRPC.ChannelManager
  alias WebsocketGateway.GRPC.Messages

  require Logger

  @timeout 10_000

  # =========================================================================
  # Wallet Service
  # =========================================================================

  @doc """
  Get user balance from wallet service.
  """
  def get_balance(user_id) do
    request = %Messages.Wallet.GetBalanceRequest{user_id: user_id}

    call_rpc(:wallet_service, "game_engine.wallet.v1.WalletService/GetBalance", request, Messages.Wallet.GetBalanceResponse)
    |> format_balance_response()
  end

  @doc """
  Check if user has sufficient balance for a bet.
  Calls GetBalance and performs the comparison locally.
  """
  def check_balance(user_id, amount) do
    case get_balance(user_id) do
      {:ok, %{"balance" => balance}} ->
        available = Map.get(balance, "available_amount", 0)
        if available >= amount do
          {:ok, %{"sufficient" => true, "available" => available}}
        else
          {:ok, %{"sufficient" => false, "available" => available, "required" => amount}}
        end

      {:ok, response} ->
        available = get_in(response, ["available_amount"]) || 0
        if available >= amount do
          {:ok, %{"sufficient" => true, "available" => available}}
        else
          {:ok, %{"sufficient" => false, "available" => available, "required" => amount}}
        end

      error ->
        error
    end
  end

  @doc """
  Deduct bet amount from user balance.
  """
  def place_bet(user_id, game_id, amount, bet_type) do
    request = %Messages.Wallet.PlaceBetRequest{
      user_id: user_id,
      game_id: game_id,
      amount: %Messages.Money{amount: amount},
      bet_type: bet_type
    }

    call_rpc(:wallet_service, "game_engine.wallet.v1.WalletService/PlaceBet", request, Messages.Wallet.PlaceBetResponse)
  end

  @doc """
  Process game result and settle winnings.
  """
  def settle_game(user_id, game_id, bet_id, result, multiplier) do
    win_amount = trunc(multiplier * 100)

    request = %Messages.Wallet.SettleBetRequest{
      bet_id: bet_id,
      settlement_type: if(result == "win", do: 2, else: 3),
      win_amount: %Messages.Money{amount: win_amount},
      result: result
    }

    call_rpc(:wallet_service, "game_engine.wallet.v1.WalletService/SettleBet", request, Messages.Wallet.SettleBetResponse)
  end

  # =========================================================================
  # Game Registry Service
  # =========================================================================

  @doc """
  Get game details from game registry service.
  """
  def get_game(game_id) do
    request = %Messages.Game.GetGameRequest{game_id: game_id}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/GetGame", request, Messages.Game.GetGameResponse)
  end

  @doc """
  Validate if game is available for playing.
  Calls GetGame and performs validation locally.
  """
  def validate_game(game_id, user_id) do
    case get_game(game_id) do
      {:ok, %{"game" => game}} ->
        status = Map.get(game, "status", "")
        if status == "ACTIVE" do
          {:ok, %{"valid" => true, "game" => game}}
        else
          {:error, "Game is not available (status: #{status})"}
        end

      {:ok, response} ->
        {:ok, %{"valid" => true, "response" => response}}

      error ->
        error
    end
  end

  @doc """
  Get featured games from game registry.
  """
  def get_featured_games(limit \\ 10) do
    request = %Messages.Game.GetFeaturedGamesRequest{limit: limit}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/GetFeaturedGames", request, Messages.Game.GetFeaturedGamesResponse)
  end

  @doc """
  Get new games from game registry.
  """
  def get_new_games(limit \\ 10) do
    request = %Messages.Game.GetNewGamesRequest{limit: limit}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/GetNewGames", request, Messages.Game.GetNewGamesResponse)
  end

  @doc """
  Get popular games from game registry.
  """
  def get_popular_games(limit \\ 10) do
    request = %Messages.Game.GetPopularGamesRequest{limit: limit}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/GetPopularGames", request, Messages.Game.GetPopularGamesResponse)
  end

  @doc """
  Get game categories from game registry.
  """
  def get_game_categories do
    request = %Messages.Game.GetCategoriesRequest{include_games_count: true}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/GetCategories", request, Messages.Game.GetCategoriesResponse)
  end

  @doc """
  Search games in game registry.
  """
  def search_games(query, limit \\ 20) do
    request = %Messages.Game.SearchGamesRequest{query: query, limit: limit}

    call_rpc(:game_registry_service, "game_engine.game.v1.GameRegistryService/SearchGames", request, Messages.Game.SearchGamesResponse)
  end

  # =========================================================================
  # Tournament Service
  # =========================================================================

  @doc """
  Get tournament details.
  """
  def get_tournament(tournament_id) do
    request = %Messages.Tournament.GetTournamentRequest{tournament_id: tournament_id}

    call_rpc(:tournament_service, "game_engine.tournament.v1.TournamentService/GetTournament", request, Messages.Tournament.GetTournamentResponse)
  end

  @doc """
  Get active tournaments.
  """
  def list_active_tournaments do
    request = %Messages.Tournament.ListTournamentsRequest{status: "active", limit: 50}

    call_rpc(:tournament_service, "game_engine.tournament.v1.TournamentService/ListTournaments", request, Messages.Tournament.ListTournamentsResponse)
  end

  @doc """
  Get tournament schedule.
  Lists all tournaments regardless of status for scheduling view.
  """
  def get_tournament_schedule do
    request = %Messages.Tournament.ListTournamentsRequest{status: "", limit: 100}

    call_rpc(:tournament_service, "game_engine.tournament.v1.TournamentService/ListTournaments", request, Messages.Tournament.ListTournamentsResponse)
  end

  @doc """
  Get tournament leaderboard.
  """
  def get_leaderboard(tournament_id) do
    request = %Messages.Tournament.GetLeaderboardRequest{tournament_id: tournament_id, limit: 50}

    call_rpc(:tournament_service, "game_engine.tournament.v1.TournamentService/GetLeaderboard", request, Messages.Tournament.GetLeaderboardResponse)
  end

  @doc """
  Get tournament standings.
  Same as leaderboard, returned as standings view.
  """
  def get_standings(tournament_id) do
    request = %Messages.Tournament.GetLeaderboardRequest{tournament_id: tournament_id, limit: 100}

    call_rpc(:tournament_service, "game_engine.tournament.v1.TournamentService/GetLeaderboard", request, Messages.Tournament.GetLeaderboardResponse)
  end

  # =========================================================================
  # Jackpot Service
  # =========================================================================

  @doc """
  Get current jackpot amounts.
  """
  def get_jackpot(jackpot_id) do
    request = %Messages.Jackpot.GetJackpotRequest{jackpot_id: jackpot_id}

    call_rpc(:jackpot_service, "game_engine.jackpot.v1.JackpotService/GetJackpot", request, Messages.Jackpot.GetJackpotResponse)
  end

  @doc """
  List all active jackpots.
  """
  def list_jackpots do
    request = %Messages.Jackpot.ListJackpotsRequest{status: "active"}

    call_rpc(:jackpot_service, "game_engine.jackpot.v1.JackpotService/ListJackpots", request, Messages.Jackpot.ListJackpotsResponse)
  end

  # =========================================================================
  # Auth Service
  # =========================================================================

  @doc """
  Authenticate user token via auth service gRPC.
  """
  def authenticate_user(token) do
    request = %Messages.Auth.ValidateTokenRequest{token: token, expected_type: "access"}

    call_rpc(:auth_service, "game_engine.auth.v1.AuthService/ValidateToken", request, Messages.Auth.ValidateTokenResponse)
  end

  # =========================================================================
  # Private - gRPC call helpers
  # =========================================================================

  defp call_rpc(service_key, method, request, response_module) do
    with {:ok, channel} <- ChannelManager.get_channel(service_key),
         {:ok, response} <- do_grpc_call(channel, method, request, response_module, @timeout) do
      {:ok, response}
    else
      {:error, :not_connected} ->
        Logger.error("gRPC channel not connected for service: #{service_key}")
        {:error, :service_unavailable}

      {:error, %GRPC.RPCError{status: status, message: message}} ->
        Logger.error("gRPC error from #{service_key}##{method}: #{status} - #{message}")
        {:error, grpc_error_to_atom(status, message)}

      {:error, %GRPC.Client.Stream{reason: reason}} ->
        Logger.error("gRPC stream error from #{service_key}##{method}: #{inspect(reason)}")
        {:error, reason}

      {:error, reason} ->
        Logger.error("gRPC call failed #{service_key}##{method}: #{inspect(reason)}")
        {:error, normalize_error(reason)}
    end
  rescue
    e ->
      Logger.error("gRPC call exception #{service_key}##{method}: #{inspect(e)}")
      {:error, :internal_error}
  end

  defp do_grpc_call(channel, method, request, response_module, timeout) do
    codec = GRPC.Codec.Proto

    GRPC.Stub.call(channel, method, request, codec: codec, timeout: timeout, return_headers: false)
    |> case do
      {:ok, response} ->
        {:ok, decode_response(response, response_module)}

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp decode_response(%{__struct__: _} = response, _response_module), do: struct_to_map(response)
  defp decode_response(response, _response_module), do: response

  defp struct_to_map(struct) when is_struct(struct) do
    struct
    |> Map.from_struct()
    |> Enum.reduce(%{}, fn
      {:__unknown_fields__, _}, acc -> acc
      {key, value}, acc -> Map.put(acc, to_string(key), sanitize_value(value))
    end)
  end

  defp struct_to_map(value), do: value

  defp sanitize_value(nil), do: nil
  defp sanitize_value(value) when is_struct(value), do: struct_to_map(value)
  defp sanitize_value(values) when is_list(values), do: Enum.map(values, &sanitize_value/1)
  defp sanitize_value(value), do: value

  defp format_balance_response({:ok, response}) do
    {:ok, response}
  end

  defp format_balance_response(error), do: error

  defp grpc_error_to_atom(status, _message) when status in [2, 3, 5, 11, 13], do: :service_error
  defp grpc_error_to_atom(status, _message) when status == 4, do: :deadline_exceeded
  defp grpc_error_to_atom(status, _message) when status == 7, do: :permission_denied
  defp grpc_error_to_atom(status, _message) when status == 12, do: :unimplemented
  defp grpc_error_to_atom(status, _message) when status == 14, do: :service_unavailable
  defp grpc_error_to_atom(status, _message) when status == 16, do: :unauthenticated
  defp grpc_error_to_atom(_status, message) when is_binary(message) and message != "", do: message
  defp grpc_error_to_atom(_status, _), do: :unknown_error

  defp normalize_error(:timeout), do: :timeout
  defp normalize_error(:connect_timeout), do: :connect_timeout
  defp normalize_error(reason) when is_atom(reason), do: reason
  defp normalize_error(reason) when is_binary(reason), do: reason
  defp normalize_error(_), do: :unknown_error
end
