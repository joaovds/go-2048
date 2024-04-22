package ui

import "fmt"

func MoveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func ClearScreen() {
	fmt.Print("\033[H\033[J")
}
