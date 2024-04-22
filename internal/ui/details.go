package ui

import "github.com/rivo/tview"

type Details struct{}

func NewDetails() *Details {
	return &Details{}
}

func (d *Details) Render() *tview.Grid {
	detailsContainer := tview.NewGrid().
		SetRows(-2, 0).
		SetColumns(0).
		SetBorders(true).
		AddItem(tview.NewTextView().SetText("Details"), 0, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("Commands"), 1, 0, 1, 1, 0, 0, false)

	return detailsContainer
}
