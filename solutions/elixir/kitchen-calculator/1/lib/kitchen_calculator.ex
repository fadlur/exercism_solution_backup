defmodule KitchenCalculator do
  def get_volume(volume_pair) do
    elem(volume_pair, 1)
  end

  def to_milliliter({:milliliter, num}), do: {:milliliter, num}
  def to_milliliter({:cup, num}), do: {:milliliter, num * 240}
  def to_milliliter({:fluid_ounce, num}), do: {:milliliter, num * 30}
  def to_milliliter({:teaspoon, num}), do: {:milliliter, num * 5}
  def to_milliliter({:tablespoon, num}), do: {:milliliter, num * 15}

  def from_milliliter(volume_pair, :cup), do: {:cup, elem(volume_pair, 1) / 240}
  def from_milliliter(volume_pair, :fluid_ounce), do: {:fluid_ounce, elem(volume_pair, 1) / 30}
  def from_milliliter(volume_pair, :teaspoon), do: {:teaspoon, elem(volume_pair, 1) / 5}
  def from_milliliter(volume_pair, :tablespoon), do: {:tablespoon, elem(volume_pair, 1) / 15}
  def from_milliliter(volume_pair, :milliliter), do: {:milliliter, elem(volume_pair, 1)}

  def convert(volume_pair, :cup) do
    to_milliliter_pair = volume_pair
    |> to_milliliter
    {:cup, elem(to_milliliter_pair, 1) / 240}
  end
  def convert(volume_pair, :fluid_ounce) do
    to_milliliter_pair = volume_pair
    |> to_milliliter
    {:fluid_ounce, elem(to_milliliter_pair, 1) / 30}
  end
  def convert(volume_pair, :teaspoon) do
    to_milliliter_pair = volume_pair
    |> to_milliliter
    {:teaspoon, elem(to_milliliter_pair, 1) / 5}
  end
  def convert(volume_pair, :tablespoon) do
    to_milliliter_pair = volume_pair
    |> to_milliliter
    {:tablespoon, elem(to_milliliter_pair, 1) / 15}
  end
  def convert(volume_pair, :milliliter) do
    to_milliliter_pair = volume_pair
    |> to_milliliter
    {:milliliter, elem(to_milliliter_pair, 1)}
  end
end
