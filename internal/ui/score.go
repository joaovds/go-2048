package ui

import (
	"fmt"

	"github.com/joaovds/go-2048/internal/logic"
)

type Score struct{}

func NewScore() *Score {
	return new(Score)
}

func (s *Score) Render(score int, newPoints int) {
	row, col := s.calcCursorPosition()
	MoveCursor(row, col)
	Colors.Orange()

	fmt.Print(DoubleTopLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleTopRightCorner)

	MoveCursor(row+1, col)
	fmt.Print(DoubleVerticalLine)

	fmt.Printf(" Score: %d ", score)

	newPointsStr := fmt.Sprintf(" +%d", newPoints)
	MoveCursor(row+1, col+logic.SIZE*TILE_WIDTH-len(newPointsStr)-3)
	fmt.Println(newPointsStr)

	MoveCursor(row+1, col+logic.SIZE*TILE_WIDTH-2)
	fmt.Print(DoubleVerticalLine)

	MoveCursor(row+2, col)
	fmt.Print(DoubleBottomLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleBottomRightCorner)
}

func (s *Score) calcCursorPosition() (int, int) {
	row := logic.SIZE*TILE_HEIGHT + TILE_OFFSET + 2

	return row, 5
}
