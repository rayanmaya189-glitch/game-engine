defmodule WebsocketGateway do
  @moduledoc """
  WebsocketGateway keeps the contexts that define your domain
  and business logic.

  Contexts are also the boundary between your business logic and the
  database, and can depend on other contexts.
  """

  def controller do
    quote do
      use Phoenix.Controller, namespace: WebsocketGateway

      import Plug.Conn

      alias WebsocketGateway.Repo
      alias WebsocketGateway.Router.Helpers, as: Routes
    end
  end

  def channel do
    quote do
      use Phoenix.Channel
      alias WebsocketGateway.Repo
      alias WebsocketGateway.Services.Presence
      alias WebsocketGateway.Services.RoomManager

      import Ecto
      import Ecto.Query
    end
  end

  def model do
    quote do
      use Ecto.Schema

      import Ecto
      import Ecto.Changeset
      import Ecto.Query
    end
  end

  @doc """
  When used, dispatch to the appropriate controller/view/etc.
  """
  defmacro __using__(which) when is_atom(which) do
    apply(__MODULE__, which, [])
  end
end
