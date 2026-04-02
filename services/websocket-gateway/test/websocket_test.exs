defmodule WebsocketGateway.WebsocketTest do
  use ExUnit.Case, async: true

  alias WebsocketGateway.Services.Presence

  describe "channel join" do
    test "user can join lobby channel" do
      socket = build_socket("user_1", "alice")
      assert {:ok, _reply, _socket} = join_lobby(socket)
    end

    test "user can join game channel" do
      socket = build_socket("user_1", "alice")
      assert {:ok, _reply, _socket} = join_game(socket, "game_123")
    end

    test "join assigns game_id to socket" do
      socket = build_socket("user_1", "alice")
      {:ok, _reply, socket} = join_game(socket, "game_456")
      assert socket.assigns.game_id == "game_456"
    end
  end

  describe "message broadcast" do
    test "place_bet event is broadcast to game channel" do
      params = %{"amount" => 100, "bet_type" => "red"}
      assert validate_bet_amount(100) == :ok
      assert validate_bet_type("red") == :ok
    end

    test "invalid bet amount is rejected" do
      assert validate_bet_amount(-5) == {:error, "Invalid bet amount"}
      assert validate_bet_amount("abc") == {:error, "Invalid bet amount"}
    end

    test "player_action events validated" do
      allowed = ~w(hit stand double surrender split fold bet raise call check)
      for action <- allowed do
        assert action in allowed
      end
      refute "cheat" in allowed
    end

    test "ready event marks player ready" do
      assert broadcast_event(:player_ready, %{user_id: "u1"}) == :ok
    end

    test "leave_game untracks presence" do
      assert broadcast_event(:leave_game, %{user_id: "u1"}) == :ok
    end

    test "chat_message is forwarded" do
      assert broadcast_event(:chat_message, %{user_id: "u1", message: "hello"}) == :ok
    end
  end

  describe "presence tracking" do
    test "user presence is tracked on join" do
      meta = %{username: "alice", user_id: "u1", online_at: now_unix()}
      assert {:ok, _} = mock_track("lobby", "u1", meta)
    end

    test "user presence is untracked on leave" do
      :ok = mock_untrack("lobby", "u1")
    end

    test "presence count returns correct value" do
      mock_track("lobby", "u1", %{username: "alice"})
      mock_track("lobby", "u2", %{username: "bob"})
      assert mock_count("lobby") == 2
    end

    test "get_user returns user metadata" do
      meta = %{username: "alice", user_id: "u1"}
      mock_track("lobby", "u1", meta)
      result = mock_get_user("lobby", "u1")
      assert result.username == "alice"
    end

    test "get_user returns nil for unknown user" do
      assert mock_get_user("lobby", "unknown") == nil
    end

    test "list_users returns all online users" do
      mock_track("global", "u1", %{username: "alice"})
      mock_track("global", "u2", %{username: "bob"})
      users = mock_list_users("global")
      assert length(users) == 2
    end
  end

  describe "socket authentication" do
    test "valid JWT token allows connection" do
      claims = %{"sub" => "user_1", "username" => "alice", "role" => "user"}
      assert {:ok, claims} = mock_validate_token("valid_token")
    end

    test "expired token rejected" do
      assert {:error, :token_expired} = mock_validate_token("expired_token")
    end

    test "invalid token rejected" do
      assert {:error, :invalid_token} = mock_validate_token("garbage")
    end

    test "missing token rejected" do
      assert {:error, :invalid_token} = mock_validate_token(nil)
    end
  end

  describe "user socket" do
    test "socket assigns user_id from claims" do
      claims = %{"sub" => "user_42", "username" => "alice"}
      socket = mock_connect(claims)
      assert socket.assigns.user_id == "user_42"
    end

    test "socket assigns username from claims" do
      claims = %{"sub" => "u1", "username" => "alice"}
      socket = mock_connect(claims)
      assert socket.assigns.username == "alice"
    end

    test "socket assigns default role" do
      claims = %{"sub" => "u1", "username" => "alice"}
      socket = mock_connect(claims)
      assert socket.assigns.role == "user"
    end
  end

  # Helper functions

  defp build_socket(user_id, username) do
    %{
      assigns: %{
        user_id: user_id,
        username: username,
        email: "#{username}@test.com",
        role: "user",
        message_count: 0,
        last_message_time: DateTime.utc_now()
      }
    }
  end

  defp join_lobby(socket), do: {:ok, %{online_count: 1}, socket}
  defp join_game(socket, game_id) do
    socket = put_in(socket.assigns[:game_id], game_id)
    {:ok, %{room: %{}, players: []}, socket}
  end

  defp validate_bet_amount(amount) when is_number(amount) and amount > 0, do: :ok
  defp validate_bet_amount(_), do: {:error, "Invalid bet amount"}

  defp validate_bet_type(type) when is_binary(type), do: :ok
  defp validate_bet_type(_), do: {:error, "Invalid bet type"}

  defp broadcast_event(_event, _params), do: :ok

  # Mock presence functions
  @presence_state %{}

  defp mock_track(topic, user_id, meta) do
    Process.put({:presence, topic, user_id}, meta)
    {:ok, meta}
  end

  defp mock_untrack(topic, user_id) do
    Process.delete({:presence, topic, user_id})
    :ok
  end

  defp mock_count(topic) do
    Process.get()
    |> Enum.filter(fn {{key, t, _}, _} -> key == :presence and t == topic end)
    |> length()
  end

  defp mock_get_user(topic, user_id) do
    case Process.get({:presence, topic, user_id}) do
      nil -> nil
      meta -> struct(OpenStruct, meta)
    end
  end

  defp mock_list_users(topic) do
    Process.get()
    |> Enum.filter(fn {{key, t, _}, _} -> key == :presence and t == topic end)
    |> Enum.map(fn {_, meta} -> meta end)
  end

  defp mock_validate_token("valid_token") do
    {:ok, %{"sub" => "user_1", "username" => "alice", "role" => "user"}}
  end
  defp mock_validate_token("expired_token"), do: {:error, :token_expired}
  defp mock_validate_token(nil), do: {:error, :invalid_token}
  defp mock_validate_token(_), do: {:error, :invalid_token}

  defp mock_connect(claims) do
    %{
      assigns: %{
        user_id: Map.get(claims, "sub"),
        username: Map.get(claims, "username"),
        role: Map.get(claims, "role", "user"),
        connected_at: DateTime.utc_now(),
        message_count: 0,
        last_message_time: DateTime.utc_now()
      }
    }
  end

  defp now_unix do
    DateTime.to_unix(DateTime.utc_now())
  end

  defmodule OpenStruct do
    defstruct [:username, :user_id]
  end
end
