package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValue := 11
	switch card {
    case "Ace":
        cardValue = 11
    case  "ten", "jack", "queen", "king":
        cardValue = 10
    case "nine":
        cardValue = 9
    case "eight":
        cardValue = 8
    case "seven":
        cardValue = 7
    case "six":
        cardValue = 6
    case "five":
        cardValue = 5
    case "four":
        cardValue = 4
    case "three":
    	cardValue = 3
    case "two":
        cardValue = 2
    case "joker":
        cardValue = 0
    }
    return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Inline helper to convert card string to its blackjack value
	value := func(card string) int {
		switch card {
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
		case "ten", "jack", "queen", "king":
			return 10
		default:
			return 0 // invalid, but won't happen in tests
		}
	}

	playerSum := value(card1) + value(card2)
	dealerValue := value(dealerCard)

	switch {
	// Pair of aces (sum == 22) → always split
	case playerSum == 22:
		return "P"

	// Blackjack (sum == 21)
	case playerSum == 21:
		if dealerValue >= 10 { // dealer has 10, J, Q, K, or ace (11)
			return "S"
		}
		return "W"

	// Hard 17-20 → stand
	case playerSum >= 17 && playerSum <= 20:
		return "S"

	// 12-16 → hit if dealer has 7+, else stand
	case playerSum >= 12 && playerSum <= 16:
		if dealerValue >= 7 {
			return "H"
		}
		return "S"

	// 11 or less → always hit
	default:
		return "H"
	}
}