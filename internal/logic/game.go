package logic

import (
	"os"
	"sync"

	"github.com/joaovds/go-2048/internal/logic/movements"
	"golang.org/x/term"
)

type Game struct {
	Values [][]int
	Wg     *sync.WaitGroup
}

func NewGame() *Game {
	wg := &sync.WaitGroup{}

	initialValues := make([][]int, SIZE)
	for i := range initialValues {
		initialValues[i] = make([]int, SIZE)
	}

	SetValueEmptyPosition(initialValues)
	SetValueEmptyPosition(initialValues)

	game := &Game{
		Values: initialValues,
		Wg:     wg,
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
	movements.Move(g.getDirection(g.readInput()))
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
