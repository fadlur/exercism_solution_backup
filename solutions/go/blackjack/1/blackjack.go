package blackjack

func ParseCard(card string) int {
	switch card {
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	// - If you have a pair of aces you must always split them.
	// - If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	// - If your cards sum up to a value within the range [17, 20] you should always stand.
	// - If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	// - If your cards sum up to 11 or lower you should always hit.
	switch {
	case ParseCard(card1)+ParseCard(card2) > 20:
		return LargeHand(IsBlackJack(card1, card2), ParseCard(dealerCard))
	default:
		return SmallHand(ParseCard(card1)+ParseCard(card2), ParseCard(dealerCard))
	}
}

func IsBlackJack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

func LargeHand(IsBlackJack bool, dealerScore int) string {
	switch {
	case !IsBlackJack:
		return "P"
	case dealerScore < 10:
		return "W"
	default:
		return "S"
	}
}

func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore <= 11 || (dealerScore >= 7 && handScore < 17):
		return "H"
	default:
		return "S"
	}
}
