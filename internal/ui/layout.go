package ui

import "github.com/joaovds/go-2048/internal/logic"

type Layout struct {
	board *Board
	game  *logic.Game
}

func NewLayout(game *logic.Game) *Layout {
	return &Layout{
		board: NewBoard(),
		game:  game,
	}
}

func (l *Layout) Render() {
	ClearScreen()
	l.board.Render(l.game.Values)
}
