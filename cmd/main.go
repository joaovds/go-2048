package main

import (
	"time"

	"github.com/joaovds/go-2048/internal/logic"
	"github.com/joaovds/go-2048/internal/ui"
)

func main() {
	game := logic.NewGame()

	layout := ui.NewLayout(game)
	layout.Render()
	layout.Timer.Render(game.GetGameDuration())

	game.Wg.Add(1)
	go listenUpdates(game, layout)

	game.Wg.Wait()
}

func listenUpdates(game *logic.Game, layout *ui.Layout) {
	for {
		select {
		case us := <-game.UpdateSignal:
			if us.GameOver {
				if game.GameData.IsNewRecord(game.Score, game.GetGameDuration()) {
					datetime := time.Now()
					game.GameData.SaveGameRecord(&logic.Data{Record: &logic.Record{Score: game.Score, Time: game.GetGameDuration(), Datetime: &datetime}})
				}
				layout.GameOver.Render()
				game.Wg.Done()
				break
			} else if us.Restart {
				game.Reset()
			}
			layout.Render()
		case <-game.Ticker.C:
			layout.Timer.Render(game.GetGameDuration())
		}
	}
}
