package ui

import "github.com/joaovds/go-2048/internal/logic"

type Layout struct {
	game     *logic.Game
	board    *Board
	score    *Score
	GameOver *GameOver
	Timer    *Timer
}

func NewLayout(game *logic.Game) *Layout {
	return &Layout{
		game:     game,
		board:    NewBoard(),
		score:    NewScore(),
		GameOver: NewGameOver(),
		Timer:    NewTimer(),
	}
}

func (l *Layout) Render() {
	ClearScreen()
	l.board.Render(l.game.Values)
	l.score.Render(l.game.Score, l.game.NewPlayPoints)
	l.Timer.Render(l.game.Stopwatch)
}
