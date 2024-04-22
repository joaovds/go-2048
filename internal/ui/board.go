package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type Board struct{}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Render() *tview.Grid {
	boardContainer := tview.NewGrid().
		SetRows(0, -10, 0).
		SetColumns(0, -4, 0).
		SetBorders(true)

	boardContainer.
		AddItem(b.makeCells(), 1, 1, 1, 1, 0, 0, false)

	return boardContainer
}

func (b *Board) makeCells() *tview.Grid {
	cellsContainer := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		SetBorders(true)

	cellsContainer.
		AddItem(b.makeCell(0), 0, 0, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 0, 1, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 0, 2, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 0, 3, 1, 1, 0, 0, false)

	cellsContainer.
		AddItem(b.makeCell(0), 1, 0, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 1, 1, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 1, 2, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 1, 3, 1, 1, 0, 0, false)

	cellsContainer.
		AddItem(b.makeCell(0), 2, 0, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 2, 1, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 2, 2, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 2, 3, 1, 1, 0, 0, false)

	cellsContainer.
		AddItem(b.makeCell(0), 3, 0, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 3, 1, 1, 1, 0, 0, false).
		AddItem(b.makeCell(0), 3, 2, 1, 1, 0, 0, false).
		AddItem(b.makeCell(8), 3, 3, 1, 1, 0, 0, false)

	return cellsContainer
}

func (b *Board) makeCell(value int) *tview.Grid {
	valueStr := ""
	if value != 0 {
		valueStr = fmt.Sprint(value)
	}

	cellContent := tview.NewTextView().SetText(valueStr).SetTextAlign(tview.AlignCenter)

	cell := tview.NewGrid().
		SetRows(-2, 0, 0).
		AddItem(cellContent, 1, 0, 1, 1, 0, 0, false)

	return cell
}
