package main

import (
	"github.com/joaovds/go-2048/internal/logic"
	"github.com/joaovds/go-2048/internal/ui"
)

func main() {
	game := logic.NewGame()

	layout := ui.NewLayout(game)
	layout.Render()

	game.Wg.Wait()
}
