package ui

import (
	"github.com/rivo/tview"
)

type Board struct{}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) Render() *tview.Grid {
	boardContainer := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		SetBorders(true)

	boardContainer.
		AddItem(tview.NewTextView().SetText("2"), 0, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 0, 1, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 0, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 0, 3, 1, 1, 0, 0, false)

	boardContainer.
		AddItem(tview.NewTextView().SetText("2"), 1, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 1, 1, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 1, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 1, 3, 1, 1, 0, 0, false)

	boardContainer.
		AddItem(tview.NewTextView().SetText("2"), 2, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 2, 1, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 2, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 2, 3, 1, 1, 0, 0, false)

	boardContainer.
		AddItem(tview.NewTextView().SetText("2"), 3, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 3, 1, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 3, 2, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("2"), 3, 3, 1, 1, 0, 0, false)

	return boardContainer
}

func (b *Board) MakeCell() *tview.Box {
	cell := tview.NewBox()

	return cell
}
