package logic

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/joaovds/go-2048/internal/logic/movements"
	"golang.org/x/term"
)

type UpdateSignal struct {
	GameOver bool
	Restart  bool
}

type Game struct {
	Values        [][]int
	Score         int
	NewPlayPoints int
	StartTime     time.Time
	Ticker        *time.Ticker
	Wg            *sync.WaitGroup
	UpdateSignal  chan *UpdateSignal
	GameData      *Data
}

func NewGame() *Game {
	wg := &sync.WaitGroup{}
	updateSignal := make(chan *UpdateSignal)

	initialValues := createInitValues()
	gameData := NewData()
	if err := gameData.GetGameData(); err != nil {
		log.Fatalf("Error when read file data: %s", err)
	}

	game := &Game{
		Values:        initialValues,
		Score:         0,
		NewPlayPoints: 0,
		StartTime:     time.Now(),
		Ticker:        time.NewTicker(time.Second),
		Wg:            wg,
		UpdateSignal:  updateSignal,
		GameData:      gameData,
	}

	wg.Add(1)
	go game.listenInput()

	return game
}

func createInitValues() [][]int {
	initialValues := make([][]int, SIZE)
	for i := range initialValues {
		initialValues[i] = make([]int, SIZE)
	}
	SetValueEmptyPosition(initialValues)
	SetValueEmptyPosition(initialValues)

	return initialValues
}

func (g *Game) Reset() {
	g.Values = createInitValues()
	g.Score = 0
	g.NewPlayPoints = 0
	g.StartTime = time.Now()
}

func (g *Game) GetGameDuration() time.Duration {
	return time.Now().Sub(g.StartTime)
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
	case 'r', 'R':
		g.UpdateSignal <- &UpdateSignal{Restart: true}
	case 'q':
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
