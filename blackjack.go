package main

import "fmt"

func main() {
	value := ParseCard("ace")
	fmt.Println(value)
	// Output: 11

	isBlackjack := IsBlackjack("queen", "ace")
	fmt.Println(isBlackjack)
	// Output: true

	isBlackJack := true
	dealerScore := 7
	choice := LargeHand(isBlackJack, dealerScore)
	fmt.Println(choice)
	// Output: "W"

	handScore := 15
	dealerScore = 12
	choice = SmallHand(handScore, dealerScore)
	fmt.Println(choice)
	// Output: "H"
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	deck := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	return deck[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
// mia soluzione senza considerare PaseCard
/*func IsBlackjack(card1, card2 string) bool {

	switch {
	case card1 == "ace" &&
		(card2 == "ten" || card2 == "jack" ||
			card2 == "queen" || card2 == "king"):
		return true
	case card2 == "ace" &&
		(card1 == "ten" || card1 == "jack" ||
			card1 == "queen" || card1 == "king"):
		return true
	default:
		return false

	}
}*/

//
func IsBlackjack(card1, card2 string) bool {
	var isBJ bool
	if ParseCard(card1)+ParseCard(card2) == 21 {
		isBJ = true
	}
	return isBJ
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack == true {
		if dealerScore < 10 {
			return "W"
		}
	} else {
		return "P"
	}
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 { //&& handScore < 21 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else if dealerScore < 7 {
		return "S"
	}
	return "H"
}
