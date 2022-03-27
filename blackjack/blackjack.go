package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	if card1 == "ace" && card2 == card1 {
		return "P"
	} else if SumUserCards(card1, card2) >= 21 && !DealerHasFigureOrTen(dealerCard) {
		return "W"
	} else if SumUserCards(card1, card2) >= 21 && DealerHasFigureOrTen(dealerCard) {
		return "S"
	} else if ParseCard(dealerCard) >= 7 && SumUserCards(card1, card2) <= 16 && SumUserCards(card1, card2) >= 12 {
		return "H"
	} else if SumUserCards(card1, card2) <= 20 && SumUserCards(card1, card2) >= 12 {
		return "S"
	}
	return "H"
}

func DealerHasFigureOrTen(card string) bool {
	return card == "ten" || card == "jack" || card == "queen" || card == "king" || card == "ace"
}

func SumUserCards(card1, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

/*

If you have a pair of aces you must always split them.

If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you
automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.

If your cards sum up to a value within the range [17, 20] you should always stand.

If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.

If your cards sum up to 11 or lower you should always hit.


    Stand (S)
    Hit (H)
    Split (P)
    Automatically win (W)

*/
