package main

import (
	"github.com/joaovds/go-2048/internal/logic"
	"github.com/joaovds/go-2048/internal/ui"
)

func main() {
	game := logic.NewGame()

	layout := ui.NewLayout(game)
	layout.Render()

	go listenUpdates(game, layout)

	game.Wg.Wait()
}

func listenUpdates(game *logic.Game, layout *ui.Layout) {
	for {
		select {
		case <-game.UpdateSignal:
			layout.Render()
		}
	}
}
