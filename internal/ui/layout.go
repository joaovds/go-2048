package ui

type Layout struct {
	board  *Board
	values [][]int
}

func NewLayout(values [][]int) *Layout {
	return &Layout{
		board:  NewBoard(),
		values: values,
	}
}

func (l *Layout) Render() {
	ClearScreen()
	l.board.Render(l.values)
}
