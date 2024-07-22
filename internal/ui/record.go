package ui

import (
	"fmt"
	"time"

	"github.com/joaovds/go-2048/internal/logic"
)

type Record struct{}

func NewRecord() *Record {
	return new(Record)
}

func (r *Record) Render(timeSinceLastRecord *time.Time, recordScore int, recordTime time.Duration) {
	var timeSinceLastRecordFormated string
	if timeSinceLastRecord != nil {
		timeSinceLastRecordDays := int((time.Now().Sub(*timeSinceLastRecord)).Hours() / 24)
		if timeSinceLastRecordDays == 0 {
			timeSinceLastRecordFormated = "today!"
		} else {
			timeSinceLastRecordFormated = fmt.Sprintf("%d days", timeSinceLastRecordDays)
		}
	}

	recordTimeFormated := fmt.Sprintf("at %s", fmt.Sprintf(" Time: %02d:%02d:%02d", int(recordTime.Hours()), int(recordTime.Minutes())%60, int(recordTime.Seconds())%60))

	row, col := r.calcCursorPosition()
	MoveCursor(row, col)
	Colors.BrightBlue()

	fmt.Print(DoubleTopLeftCorner)
	for range 38 {
		fmt.Print(DoubleHorizontalLine)
	}
	fmt.Print(DoubleTopRightCorner)

	for i := range 2 {
		MoveCursor(row+(i*2)+1, col)
		fmt.Print(DoubleVerticalLine)

		MoveCursor(row+(i*2)+1, col+4)

		if i == 0 {
			fmt.Printf("Last record in %s", timeSinceLastRecordFormated)
		} else {
			fmt.Printf("Score: %d", recordScore)
			MoveCursor(row+(i*2)+1, col+34-(len(recordTimeFormated)-1))
			fmt.Printf(recordTimeFormated)
		}

		MoveCursor(row+(i*2)+1, col+39)
		fmt.Print(DoubleVerticalLine)

		var (
			leftBorder  string
			rightBorder string
		)
		if i == len(ActionsOptions)-1 {
			leftBorder = DoubleBottomLeftCorner
			rightBorder = DoubleBottomRightCorner

			MoveCursor(row+i+i+2, col+1)
			for range 38 {
				fmt.Print(DoubleHorizontalLine)
			}
		} else {
			leftBorder = DoubleVerticalLine
			rightBorder = DoubleVerticalLine
		}

		MoveCursor(row+i+i+2, col)
		fmt.Print(leftBorder)

		MoveCursor(row+i+i+2, col+39)
		fmt.Print(rightBorder)
	}
}

func (r *Record) calcCursorPosition() (int, int) {
	col := logic.SIZE*TILE_WIDTH + 10
	row := TILE_OFFSET + 2 + (len(ActionsOptions) * 2)

	return row, col
}
