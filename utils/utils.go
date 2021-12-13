package utils

import (
	"fmt"
	"math/big"
	"os"
)

// ############################ CONSTANTS ##############################
// ### this section holds all constants needed.
const (
	black = "\u001b[30m"
	red = "\u001b[31m"
	yellow = "\u001b[33m"
	green = "\u001b[32m"
	blue = "\u001b[34m"
	reset = "\u001b[0m"
)

// ############################ ERRORS #################################
// ### this section handles error checking, printing, etc.

// checks and error and panics. Used primarily for debugging
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// handles an error in a pretty way for the user.
func HandleError(e error) {
	if e != nil {
		PrintError(e.Error())
		os.Exit(1)
	}
}

// ############################ PRINTING FUNCTIONS #########################
// ### this section contains all printing functions

func PrintDebug(msg string) {
	fmt.Println(blue + msg + reset)
}

func PrintInfo(msg string) {
	fmt.Println(green + msg + reset)
}

func PrintWarning(msg string) {
	fmt.Println(yellow + msg + reset)
}

func PrintError(msg string) {
	fmt.Println(red + msg + reset)
}

// Formats numbers with commas separating thousands
func PrintNumber(num *big.Int) {
	text := num.String()
	for i, d := range text {
		if i > 0 && (len(text)-i) % 3 == 0 {
			fmt.Print(",")
		}
		fmt.Print(string(d))
	}
	fmt.Println()
	fmt.Println("digit count:", len(text))
}