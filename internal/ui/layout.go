package ui

type Layout struct {
	board  *Board
	values [][]int
}

func NewLayout() *Layout {
	values := make([][]int, SIZE)
	for i := range values {
		values[i] = make([]int, SIZE)
	}

	return &Layout{
		board:  NewBoard(),
		values: values,
	}
}

func (l *Layout) Render() {
	ClearScreen()
	l.board.Render(l.values)
}
