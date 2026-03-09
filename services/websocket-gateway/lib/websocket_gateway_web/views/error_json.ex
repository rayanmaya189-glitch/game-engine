defmodule WebsocketGateway.ErrorJSON do
  @moduledoc """
  This module is invoked by your endpoint in case of errors on JSON requests.
  """

  def render(template, _assigns) do
    %{errors: %{detail: Phoenix.Controller.status_message_from_template(template)}}
  end

  def error(%{message: message, status: status}) do
    %{
      error: %{
        code: status,
        message: message
      }
    }
  end
end
