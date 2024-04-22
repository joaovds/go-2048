package ui

import "github.com/rivo/tview"

type Layout struct {
	BoardComponent   *Board
	DetailsComponent *Details
}

func NewLayout() *Layout {
	return &Layout{
		BoardComponent:   NewBoard(),
		DetailsComponent: NewDetails(),
	}
}

func (l *Layout) Render() {
	layout := tview.NewGrid().
		SetRows(0).
		SetColumns(-2, 0).
		SetBorders(true).
		AddItem(l.BoardComponent.Render(), 0, 0, 1, 1, 0, 0, false).
		AddItem(l.DetailsComponent.Render(), 0, 1, 1, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(layout, true).
		SetFocus(layout).Run(); err != nil {
		panic(err)
	}
}
