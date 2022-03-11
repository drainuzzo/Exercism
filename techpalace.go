package main

import (
	"fmt"
	"strings"
)

func main() {

	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	fmt.Println("WecomeMessage: ", WelcomeMessage("Judy"))
	fmt.Println("AddBorder:\n", AddBorder("Welcome!", 10))
	fmt.Println("before: ", message)
	fmt.Println("after Cleanup: ", CleanupMessage(strings.TrimSpace(message)))
}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	upCustomer := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + upCustomer
}

// MIA senza strings: AddBorder adds a border to a welcome message.
/*func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var line string
	for i := 0; i < numStarsPerLine; i++ {
		line += "*"
	}
	s := line + "\n" + welcomeMsg + "\n" + line
	return s
}*/

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// MIA skifo: CleanupMessage cleans up an old marketing message.
/*func CleanupMessage(oldMsg string) string {
	s := strings.TrimSpace(oldMsg)
	s = strings.Trim(s, "*")
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "*")
	s = strings.TrimSpace(s)
	return s
}*/
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
