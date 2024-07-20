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

	game.Wg.Add(1)
	go listenUpdates(game, layout)

	game.Wg.Wait()
}

func listenUpdates(game *logic.Game, layout *ui.Layout) {
	for {
		select {
		case us := <-game.UpdateSignal:
			if us.GameOver {
				layout.GameOver.Render()
				game.Wg.Done()
				break
			} else if us.Restart {
				game.Reset()
			}
			layout.Render()
		case <-game.Ticker.C:
			game.Stopwatch.Update()
			layout.Timer.Render(game.Stopwatch)
		}
	}
}
