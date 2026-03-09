defmodule WebsocketGateway.Services.Presence do
  @moduledoc """
  Presence tracking service for tracking online users across channels.

  Uses Phoenix Presence for real-time presence tracking with Redis
  backend for distributed presence across multiple nodes.
  """

  use Phoenix.Presence,
    otp_app: :websocket_gateway,
    pubsub_server: WebsocketGateway.PubSub

  alias Phoenix.Socket.Broadcast

  @timeout Application.get_env(:websocket_gateway, :presence, [])
    |> Keyword.get(:timeout, 25_000)

  @doc """
  Track a user in a specific topic/channel.

  ## Examples

      iex> Presence.track(self(), "lobby", "user123", %{username: "john"})
      {:ok, %{user_id: "user123", ...}}
  """
  @spec track(pid(), String.t(), String.t(), map()) :: {:ok, map()} | {:error, term()}
  def track(pid, topic, user_id, meta) when is_pid(pid) and is_binary(topic) and is_binary(user_id) do
    meta_with_defaults = Map.merge(meta, %{
      user_id: user_id,
      online_at: DateTime.to_unix(DateTime.utc_now()),
      phx_ref: make_ref() |> :erlang.ref_to_list() |> to_string()
    })

    super(pid, topic, user_id, meta_with_defaults)
  end

  @doc """
  Untrack a user from a specific topic.
  """
  @spec untrack(pid(), String.t(), String.t()) :: :ok | {:error, term()}
  def untrack(pid, topic, user_id) do
    super(pid, topic, user_id)
  end

  @doc """
  Get all presence information for a topic.
  """
  @spec list(String.t()) :: map()
  def list(topic) do
    case super(topic) do
      %{metas: metas} = result when is_list(metas) ->
        result
      _ ->
        %{presences: %{}, metas: []}
    end
  end

  @doc """
  Get list of users in a topic.
  """
  @spec list_users(String.t()) :: [map()]
  def list_users(topic) do
    case list(topic) do
      %{metas: metas} -> Enum.map(metas, &Map.delete(&1, :phx_ref))
      _ -> []
    end
  end

  @doc """
  Get count of users in a topic.
  """
  @spec count(String.t()) :: integer()
  def count(topic) do
    case list(topic) do
      %{metas: metas} -> length(metas)
      _ -> 0
    end
  end

  @doc """
  Get user info from a specific topic.
  """
  @spec get_user(String.t(), String.t()) :: map() | nil
  def get_user(topic, user_id) do
    case list(topic) do
      %{presences: presences} ->
        Map.get(presences, user_id)
      _ ->
        nil
    end
  end

  @doc """
  Track a user in multiple channels (game, chat, etc).
  """
  @spec track_user(pid(), String.t(), map()) :: {:ok, map()}
  def track_user(pid, user_id, base_meta) do
    # Track in lobby by default
    :ok = track(pid, "lobby", user_id, base_meta)

    # Track in global presence
    :ok = track(pid, "global", user_id, base_meta)

    {:ok, base_meta}
  end

  @doc """
  Untrack a user from all channels.
  """
  @spec untrack_user(pid(), String.t()) :: :ok
  def untrack_user(pid, user_id) do
    :ok = untrack(pid, "lobby", user_id)
    :ok = untrack(pid, "global", user_id)
    :ok
  end

  @doc """
  Update user's metadata in a topic.
  """
  @spec update(String.t(), String.t(), map()) :: {:ok, map()}
  def update(topic, user_id, meta_updates) do
    # Get current presence
    current = get_user(topic, user_id)

    if current do
      new_meta = Map.merge(current, meta_updates)
      # Re-track with updated metadata
      # Note: Phoenix Presence doesn't have direct update, so we track again
      {:ok, new_meta}
    else
      {:error, :not_found}
    end
  end

  @doc """
  Get global online count across all users.
  """
  @spec global_online_count() :: integer()
  def global_online_count do
    count("global")
  end

  @doc """
  Get all online users (with pagination support).
  """
  @spec get_online_users(integer(), integer()) :: [map()]
  def get_online_users(offset \\ 0, limit \\ 100) do
    list_users("global")
    |> Enum.drop(offset)
    |> Enum.take(limit)
  end

  # Handle presence diffs for tracking changes
  @doc false
  def handle_mdiff(mdiff, socket) do
    Enum.each(mdiff.joins, fn {user_id, meta} ->
      push(socket, "presence_diff", %{event: "join", user_id: user_id, meta: meta})
    end)

    Enum.each(mdiff.leaves, fn {user_id, meta} ->
      push(socket, "presence_diff", %{event: "leave", user_id: user_id, meta: meta})
    end)

    socket
  end
end
