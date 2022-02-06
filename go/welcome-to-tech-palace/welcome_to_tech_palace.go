package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	// c := strings.ToUpper(customer)
	// return fmt.Sprintf("Welcome to the Tech Palace, %s", c)

}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)             // method repeat of funct strings
	return fmt.Sprintf("%s\n%s\n%s", star, welcomeMsg, star) // customize output

}

// CleanupMessage cleans up an old marketing message.
// Fisrt we replace all occurance of star by empty string but we must delete space too and it's possible with TrimSpace
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
