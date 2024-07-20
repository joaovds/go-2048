package ui

import (
	"fmt"

	"github.com/joaovds/go-2048/internal/logic"
)

type GameOver struct{}

func NewGameOver() *GameOver {
	return new(GameOver)
}

func (g *GameOver) Render() {
	row, col := g.calcCursorPosition()
	MoveCursor(row, col)
	Colors.Red()

	fmt.Print(DoubleTopLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleTopRightCorner)

	MoveCursor(row+1, col)
	fmt.Print(DoubleVerticalLine)

	fmt.Print(" Game Over!!! ")

	MoveCursor(row+1, col+logic.SIZE*TILE_WIDTH-2)
	fmt.Print(DoubleVerticalLine)

	MoveCursor(row+2, col)
	fmt.Print(DoubleBottomLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleBottomRightCorner)
}

func (g *GameOver) calcCursorPosition() (int, int) {
	row := logic.SIZE*TILE_HEIGHT + TILE_OFFSET + 8

	return row, 5
}
