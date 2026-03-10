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
  Authenticate user token.
  """
  def authenticate_user(token) do
    service_url(:auth_service, "/api/v1/auth/validate")
    |> make_request(%{token: token})
  end

  # Private functions

  defp service_url(service_key, path) do
    base_url = Application.get_env(:websocket_gateway, :services, [])
               |> Keyword.get(service_key, "http://localhost:8080")

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
