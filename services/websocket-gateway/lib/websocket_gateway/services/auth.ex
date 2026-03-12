defmodule WebsocketGateway.Services.Auth do
  @moduledoc """
  Authentication service for JWT token validation.

  Provides functionality for:
  - Token validation
  - Token refresh
  - Claims extraction
  """

  @jwt_config Application.get_env(:websocket_gateway, :jwt, [])

  @secret_key Keyword.get(@jwt_config, :secret_key)
  @algorithm Keyword.get(@jwt_config, :algorithm, "HS256")
  @issuer Keyword.get(@jwt_config, :issuer, "game_engine")
  @audience Keyword.get(@jwt_config, :audience, "game_engine")
  @expiration Keyword.get(@jwt_config, :expiration, 86_400)

  @doc """
  Validate a JWT token.

  ## Examples

      iex> Auth.validate_token("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
      {:ok, %{sub: "user123", username: "john", ...}}
  """
  @spec validate_token(String.t()) :: {:ok, map()} | {:error, atom()}
  def validate_token(token) when is_binary(token) do
    case JOSE.JWT.decode(token) do
      {:ok, jwt, _jws} ->
        # Verify claims
        verify_claims(jwt)

      {:error, _reason} ->
        {:error, :invalid_token}
    end
  end

  def validate_token(_), do: {:error, :invalid_token}

  @doc """
  Generate a JWT token for testing or internal use.
  """
  @spec generate_token(map()) :: String.t()
  def generate_token(claims) do
    now = DateTime.utc_now() |> DateTime.to_unix()

    default_claims = %{
      "iss" => @issuer,
      "aud" => @audience,
      "iat" => now,
      "exp" => now + @expiration
    }

    merged_claims = Map.merge(default_claims, claims)

    jws = %{
      "alg" => @algorithm
    }

    jwt = JOSE.JWT.sign(@secret_key, jws, merged_claims)
    {_, token} = JOSE.JWS.compact(jws)
    token
  end

  @doc """
  Refresh a JWT token.
  """
  @spec refresh_token(String.t()) :: {:ok, String.t()} | {:error, atom()}
  def refresh_token(token) do
    case validate_token(token) do
      {:ok, claims} ->
        # Generate new token with updated expiration
        new_claims = claims
          |> Map.delete("exp")
          |> Map.delete("iat")

        {:ok, generate_token(new_claims)}

      error ->
        error
    end
  end

  @doc """
  Extract user ID from token.
  """
  @spec get_user_id(String.t()) :: {:ok, String.t()} | {:error, atom()}
  def get_user_id(token) do
    case validate_token(token) do
      {:ok, claims} ->
        user_id = Map.get(claims, "sub") || Map.get(claims, "user_id")
        if user_id, do: {:ok, user_id}, else: {:error, :no_user_id}

      error ->
        error
    end
  end

  @doc """
  Verify token belongs to specific user.
  """
  @spec verify_user(String.t(), String.t()) :: boolean()
  def verify_user(token, user_id) do
    case validate_token(token) do
      {:ok, claims} ->
        claims_sub = Map.get(claims, "sub") || Map.get(claims, "user_id")
        claims_sub == user_id

      _ ->
        false
    end
  end

  @doc """
  Check if token is expired.
  """
  @spec expired?(String.t()) :: boolean()
  def expired?(token) do
    case JOSE.JWT.decode(token) do
      {:ok, jwt, _} ->
        case JOSE.JWT.expired?(jwt) do
          true -> true
          false -> false
        end

      _ ->
        true
    end
  end

  @doc """
  Get token expiration timestamp.
  """
  @spec get_expiration(String.t()) :: {:ok, integer()} | {:error, atom()}
  def get_expiration(token) do
    case JOSE.JWT.decode(token) do
      {:ok, jwt, _} ->
        case Map.get(jwt.fields, "exp") do
          nil -> {:error, :no_expiration}
          exp -> {:ok, exp}
        end

      _ ->
        {:error, :invalid_token}
    end
  end

  @doc """
  Verify token type (access vs refresh).
  """
  @spec verify_token_type(String.t(), String.t()) :: boolean()
  def verify_token_type(token, expected_type) do
    case validate_token(token) do
      {:ok, claims} ->
        Map.get(claims, "type", "access") == expected_type

      _ ->
        false
    end
  end

  # Private functions

  defp verify_claims(jwt) do
    now = DateTime.utc_now() |> DateTime.to_unix()

    # Check expiration
    if JOSE.JWT.expired?(jwt) do
      {:error, :token_expired}
    else
      claims = jwt.fields

      # Verify issuer if configured
      if @issuer != nil do
        if Map.get(claims, "iss") != @issuer do
          return {:error, :invalid_issuer}
        end
      end

      # Verify audience if configured
      if @audience != nil do
        aud = Map.get(claims, "aud")
        if is_list(aud) do
          if @audience not in aud do
            return {:error, :invalid_audience}
          end
        else
          if aud != @audience do
            return {:error, :invalid_audience}
          end
        end
      end

      {:ok, claims}
    end
  end
end
