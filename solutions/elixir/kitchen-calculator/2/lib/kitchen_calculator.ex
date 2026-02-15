defmodule KitchenCalculator do
  @milliliter 1
  @cup 240
  @fluid_ounce 30
  @teaspoon 5
  @tablespoon 15

  def get_volume({_, volume}), do: volume

  def to_milliliter({:milliliter, volume}), do: {:milliliter, volume}
  def to_milliliter({:cup, volume}), do: {:milliliter, volume * @cup}
  def to_milliliter({:fluid_ounce, volume}), do: {:milliliter, volume * @fluid_ounce}
  def to_milliliter({:teaspoon, volume}), do: {:milliliter, volume * @teaspoon}
  def to_milliliter({:tablespoon, volume}), do: {:milliliter, volume * @tablespoon}

  def from_milliliter({_, volume}, :milliliter), do: {:milliliter, volume / @milliliter}
  def from_milliliter({_, volume}, :cup), do: {:cup, volume / @cup}
  def from_milliliter({_, volume}, :fluid_ounce), do: {:fluid_ounce, volume / @fluid_ounce}
  def from_milliliter({_, volume}, :teaspoon), do: {:teaspoon, volume / @teaspoon}
  def from_milliliter({_, volume}, :tablespoon), do: {:tablespoon, volume / @tablespoon}

  def convert(volume_pair, :cup) do
    volume_pair
    |> to_milliliter
    |> from_milliliter(:cup)
  end

  def convert(volume_pair, :fluid_ounce) do
    volume_pair
    |> to_milliliter
    |> from_milliliter(:fluid_ounce)
  end

  def convert(volume_pair, :teaspoon) do
    volume_pair
    |> to_milliliter
    |> from_milliliter(:teaspoon)
  end

  def convert(volume_pair, :tablespoon) do
    volume_pair
    |> to_milliliter
    |> from_milliliter(:tablespoon)
  end

  def convert(volume_pair, :milliliter) do
    volume_pair
    |> to_milliliter
    |> from_milliliter(:milliliter)
  end
end
