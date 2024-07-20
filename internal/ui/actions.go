package ui

import (
	"fmt"

	"github.com/joaovds/go-2048/internal/logic"
)

type Actions struct{}

func NewActions() *Actions {
	return new(Actions)
}

type Action struct {
	name    string
	command string
}

var actionsOptions []Action = []Action{
	{
		name:    "Restart",
		command: "r",
	},
	{
		name:    "Quit",
		command: "q",
	},
}

func (a *Actions) Render() {
	row, col := a.calcCursorPosition()
	MoveCursor(row, col)
	Colors.Red()

	fmt.Print(DoubleTopLeftCorner)
	for range 38 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Print(DoubleTopRightCorner)

	for i, action := range actionsOptions {
		MoveCursor(row+(i*2)+1, col)
		fmt.Print(DoubleVerticalLine)

		MoveCursor(row+(i*2)+1, col+4)
		fmt.Print(action.name)
		MoveCursor(row+(i*2)+1, col+34)
		fmt.Print(action.command)

		MoveCursor(row+(i*2)+1, col+39)
		fmt.Print(DoubleVerticalLine)

		var (
			line        string
			leftBorder  string
			rightBorder string
		)
		if i == len(actionsOptions)-1 {
			line = DoubleHorizontalLine
			leftBorder = DoubleBottomLeftCorner
			rightBorder = DoubleBottomRightCorner
		} else {
			line = HorizontalLine
			leftBorder = DoubleVerticalLine
			rightBorder = DoubleVerticalLine
		}

		MoveCursor(row+i+i+2, col)
		fmt.Print(leftBorder)

		MoveCursor(row+i+i+2, col+1)
		for range 38 {
			fmt.Print(line)
		}

		MoveCursor(row+i+i+2, col+39)
		fmt.Print(rightBorder)
	}
}

func (a *Actions) calcCursorPosition() (int, int) {
	col := logic.SIZE*TILE_WIDTH + 10
	row := TILE_OFFSET + 1

	return row, col
}
