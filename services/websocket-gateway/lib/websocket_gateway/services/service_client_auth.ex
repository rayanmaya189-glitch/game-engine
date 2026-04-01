defmodule WebsocketGateway.Services.ServiceClientAuth do
  @moduledoc """
  gRPC client for auth service calls.
  """

  alias WebsocketGateway.GRPC.ChannelManager
  alias WebsocketGateway.GRPC.Messages

  require Logger

  @timeout 10_000

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
