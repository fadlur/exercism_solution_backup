package techpalace

import (
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

func CleanupMessage(oldMsg string) string {
	newMsg := oldMsg
	newMsg = strings.ReplaceAll(newMsg, "*", "")
	newMsg = strings.TrimSpace(newMsg)
	return newMsg
}
