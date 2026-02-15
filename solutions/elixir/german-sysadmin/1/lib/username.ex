defmodule Username do
  def sanitize('') do
    ''
  end

  def sanitize([first_letter | tail]) do
    sanitized =
      case first_letter do
        ?ä -> ~c"ae"
        ?ö -> ~c"oe"
        ?ü -> ~c"ue"
        ?ß -> ~c"ss"
        x when x >= ?a and x <= ?z -> [x]
        ?_ -> ~c"_"
        _ -> ~c""
      end

    sanitized ++ sanitize(tail)
  end
end
