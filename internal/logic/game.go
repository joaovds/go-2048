package logic

import (
	"os"
	"sync"
	"time"

	"github.com/joaovds/go-2048/internal/logic/movements"
	"golang.org/x/term"
)

type Stopwatch struct {
	Hours, Minutes, Seconds int
}

func (s *Stopwatch) Update() {
	s.Seconds++
	if s.Seconds == 60 {
		s.Seconds = 0
		s.Minutes++

		if s.Minutes == 60 {
			s.Minutes = 0
			s.Hours++
		}
	}
}

type UpdateSignal struct{ GameOver bool }

type Game struct {
	Values        [][]int
	Score         int
	NewPlayPoints int
	Ticker        *time.Ticker
	Stopwatch     *Stopwatch
	Wg            *sync.WaitGroup
	UpdateSignal  chan *UpdateSignal
}

func NewGame() *Game {
	wg := &sync.WaitGroup{}
	updateSignal := make(chan *UpdateSignal)

	initialValues := make([][]int, SIZE)
	for i := range initialValues {
		initialValues[i] = make([]int, SIZE)
	}

	SetValueEmptyPosition(initialValues)
	SetValueEmptyPosition(initialValues)

	game := &Game{
		Values:        initialValues,
		Score:         0,
		NewPlayPoints: 0,
		Ticker:        time.NewTicker(time.Second),
		Stopwatch:     new(Stopwatch),
		Wg:            wg,
		UpdateSignal:  updateSignal,
	}

	wg.Add(1)
	go game.listenInput()

	return game
}

func (g *Game) listenInput() {
	defer g.Wg.Done()

	for {
		g.makeMove()
	}
}

func (g *Game) makeMove() {
	direction := g.getDirection(g.readInput())

	currentValues := make([][]int, len(g.Values))
	for row := range len(currentValues) {
		currentValues[row] = make([]int, len(g.Values[row]))

		for col := range g.Values[row] {
			currentValues[row][col] = g.Values[row][col]
		}
	}

	newPoints := movements.Move(
		direction,
		g.Values,
	)

	g.Score += newPoints
	g.NewPlayPoints = newPoints

	if (direction != movements.NONE) && HasEmptyCell(g.Values) && changed(currentValues, g.Values) {
		SetValueEmptyPosition(g.Values)
		g.UpdateSignal <- &UpdateSignal{}
	}

	if IsGameOver(g.Values) {
		g.Ticker.Stop()
		g.UpdateSignal <- &UpdateSignal{GameOver: true}
	}
}

func (g *Game) getDirection(char rune) movements.Direction {
	var direction movements.Direction

	switch char {
	case 'w', 'A':
		direction = movements.UP
	case 's', 'B':
		direction = movements.DOWN
	case 'a', 'D':
		direction = movements.LEFT
	case 'd', 'C':
		direction = movements.RIGHT
	case 'q':
		g.Ticker.Stop()
		os.Exit(0)
	default:
		direction = movements.NONE
	}

	return direction
}

func (g *Game) readInput() rune {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var buf [1]byte
	os.Stdin.Read(buf[:])
	return rune(buf[0])
}
