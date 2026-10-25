package tui

import "fmt"

func SetupTerm() {
	// Clean screen and hide cursor
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Print("\033[?25l")
}

func ResetTerm() {
	// Show cursor and reset terminal
	fmt.Print("\033[?25h")
	fmt.Print("\033[0m")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
