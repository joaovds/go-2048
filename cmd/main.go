package main

import (
	"github.com/joaovds/go-2048/internal/logic"
	"github.com/joaovds/go-2048/internal/ui"
)

func main() {
	game := logic.NewGame()

	layout := ui.NewLayout(game)
	layout.Render()
	layout.Timer.Render(game.Stopwatch)

	go listenUpdates(game, layout)

	game.Wg.Wait()
}

func listenUpdates(game *logic.Game, layout *ui.Layout) {
	for {
		select {
		case us := <-game.UpdateSignal:
			layout.Render()
			if us.GameOver {
				// Layout Game Over Render
			}
		case <-game.Ticker.C:
			game.Stopwatch.Update()
			layout.Timer.Render(game.Stopwatch)
		}
	}
}
