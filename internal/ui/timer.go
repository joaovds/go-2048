package ui

import (
	"fmt"
	"time"

	"github.com/joaovds/go-2048/internal/logic"
)

type Timer struct{}

func NewTimer() *Timer {
	return new(Timer)
}

func (t *Timer) Render(duration time.Duration) {
	row, col := t.calcCursorPosition()
	MoveCursor(row, col)
	Colors.BrightBlack()

	fmt.Print(DoubleTopLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleTopRightCorner)

	MoveCursor(row+1, col)
	fmt.Print(DoubleVerticalLine)

	fmt.Printf(" Time: %02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)

	MoveCursor(row+1, col+logic.SIZE*TILE_WIDTH-2)
	fmt.Print(DoubleVerticalLine)

	MoveCursor(row+2, col)
	fmt.Print(DoubleBottomLeftCorner)
	for range logic.SIZE*TILE_WIDTH - 3 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Println(DoubleBottomRightCorner)
}

func (t *Timer) calcCursorPosition() (int, int) {
	row := logic.SIZE*TILE_HEIGHT + TILE_OFFSET + 5

	return row, 5
}
