defmodule NameBadge do
  def print(id, name, department) do
    suffix_id = if is_nil(id), do: "", else: "[#{id}] - "
    suffix_department = if is_nil(department), do: "Owner", else: department
    suffix_id<>name<>" - "<>String.upcase(suffix_department)
  end
end
