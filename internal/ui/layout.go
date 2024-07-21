package ui

import "github.com/joaovds/go-2048/internal/logic"

type Layout struct {
	game     *logic.Game
	board    *Board
	score    *Score
	Actions  *Actions
	record   *Record
	GameOver *GameOver
	Timer    *Timer
}

func NewLayout(game *logic.Game) *Layout {
	return &Layout{
		game:     game,
		board:    NewBoard(),
		score:    NewScore(),
		Actions:  NewActions(),
		record:   NewRecord(),
		GameOver: NewGameOver(),
		Timer:    NewTimer(),
	}
}

func (l *Layout) Render() {
	ClearScreen()
	l.board.Render(l.game.Values)
	l.score.Render(l.game.Score, l.game.NewPlayPoints)
	l.Timer.Render(l.game.Stopwatch)
	l.Actions.Render()
	l.record.Render(l.game.GameData.Record.Datetime, l.game.GameData.Record.Score, l.game.GameData.Record.Time)
}
