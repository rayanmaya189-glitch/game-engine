defmodule WebsocketGateway.Services.ServiceClient do
  @moduledoc """
  HTTP client for communicating with backend services.
  Used by WebSocket channels to call wallet, game, and other services.
  """

  @timeout 10_000

  @doc """
  Get user balance from wallet service.
  """
  def get_balance(user_id) do
    service_url(:wallet_service, "/api/v1/wallet/balance")
    |> make_request(%{userId: user_id})
  end

  @doc """
  Check if user has sufficient balance for a bet.
  """
  def check_balance(user_id, amount) do
    service_url(:wallet_service, "/api/v1/wallet/check-balance")
    |> make_request(%{userId: user_id, amount: amount})
  end

  @doc """
  Deduct bet amount from user balance.
  """
  def place_bet(user_id, game_id, amount, bet_type) do
    service_url(:wallet_service, "/api/v1/wallet/bet")
    |> make_request(%{
      userId: user_id,
      gameId: game_id,
      amount: amount,
      betType: bet_type
    })
  end

  @doc """
  Process game result and settle winnings.
  """
  def settle_game(user_id, game_id, bet_id, result, multiplier) do
    service_url(:wallet_service, "/api/v1/wallet/settle")
    |> make_request(%{
      userId: user_id,
      gameId: game_id,
      betId: bet_id,
      result: result,
      multiplier: multiplier
    })
  end

  @doc """
  Get game details from game service.
  """
  def get_game(game_id) do
    service_url(:game_service, "/api/v1/games/#{game_id}")
    |> make_request(%{})
  end

  @doc """
  Validate if game is available for playing.
  """
  def validate_game(game_id, user_id) do
    service_url(:game_service, "/api/v1/games/#{game_id}/validate")
    |> make_request(%{userId: user_id})
  end

  @doc """
  Get tournament details.
  """
  def get_tournament(tournament_id) do
    service_url(:tournament_service, "/api/v1/tournaments/#{tournament_id}")
    |> make_request(%{})
  end

  @doc """
  Get current jackpot amounts.
  """
  def get_jackpot(jackpot_id) do
    service_url(:jackpot_service, "/api/v1/jackpots/#{jackpot_id}")
    |> make_request(%{})
  end

  @doc """
  List all active jackpots.
  """
  def list_jackpots do
    service_url(:jackpot_service, "/api/v1/jackpots")
    |> make_request(%{})
  end

  @doc """
  Get featured games from game registry.
  """
  def get_featured_games(limit \\ 10) do
    service_url(:game_registry_service, "/api/v1/games/featured")
    |> make_request(%{limit: limit})
  end

  @doc """
  Get new games from game registry.
  """
  def get_new_games(limit \\ 10) do
    service_url(:game_registry_service, "/api/v1/games/new")
    |> make_request(%{limit: limit})
  end

  @doc """
  Get popular games from game registry.
  """
  def get_popular_games(limit \\ 10) do
    service_url(:game_registry_service, "/api/v1/games/popular")
    |> make_request(%{limit: limit})
  end

  @doc """
  Get game categories from game registry.
  """
  def get_game_categories do
    service_url(:game_registry_service, "/api/v1/games/categories")
    |> make_request(%{})
  end

  @doc """
  Search games in game registry.
  """
  def search_games(query, limit \\ 20) do
    service_url(:game_registry_service, "/api/v1/games/search")
    |> make_request(%{query: query, limit: limit})
  end

  @doc """
  Get active tournaments.
  """
  def list_active_tournaments do
    service_url(:tournament_service, "/api/v1/tournaments/active")
    |> make_request(%{})
  end

  @doc """
  Get tournament schedule.
  """
  def get_tournament_schedule do
    service_url(:tournament_service, "/api/v1/tournaments/schedule")
    |> make_request(%{})
  end

  @doc """
  Get tournament leaderboard.
  """
  def get_leaderboard(tournament_id) do
    service_url(:tournament_service, "/api/v1/tournaments/#{tournament_id}/leaderboard")
    |> make_request(%{})
  end

  @doc """
  Get tournament standings.
  """
  def get_standings(tournament_id) do
    service_url(:tournament_service, "/api/v1/tournaments/#{tournament_id}/standings")
    |> make_request(%{})
  end

  @doc """
  Authenticate user token.
  """
  def authenticate_user(token) do
    service_url(:auth_service, "/api/v1/auth/validate")
    |> make_request(%{token: token})
  end

  # Private functions

  defp service_url(service_key, path) do
    services = Application.get_env(:websocket_gateway, :services, [])
    base_url = Keyword.get_lazy(services, service_key, fn ->
      System.get_env("DEFAULT_SERVICE_URL") || "http://localhost:8080"
    end)

    "#{base_url}#{path}"
  end

  defp make_request(url, body) do
    case HTTPoison.post(url, Jason.encode!(body), [
      {"Content-Type", "application/json"},
      {"Accept", "application/json"}
    ], timeout: @timeout, recv_timeout: @timeout) do
      {:ok, %{status_code: 200, body: response_body}} ->
        case Jason.decode(response_body) do
          {:ok, %{"success" => true} = response} -> {:ok, response}
          {:ok, %{"success" => false, "error" => error}} -> {:error, error}
          {:ok, response} -> {:ok, response}
          {:error, _} -> {:error, :decode_error}
        end
      {:ok, %{status_code: status, body: body}} when status >= 500 ->
        {:error, :service_error}
      {:ok, %{status_code: status, body: body}} when status >= 400 ->
        case Jason.decode(body) do
          {:ok, %{"error" => error}} -> {:error, error}
          _ -> {:error, :client_error}
        end
      {:ok, %{status_code: status}} ->
        {:error, "Unexpected status: #{status}"}
      {:error, %HTTPoison.Error{reason: :timeout}} ->
        {:error, :timeout}
      {:error, %HTTPoison.Error{reason: :connect_timeout}} ->
        {:error, :connect_timeout}
      {:error, %HTTPoison.Error{reason: reason}} ->
        {:error, reason}
      {:error, reason} ->
        {:error, reason}
    end
  rescue
    e in HTTPoison.Error -> {:error, e.reason}
    e in Jason.EncodeError -> {:error, :encode_error}
    _ -> {:error, :unknown}
  end
end
