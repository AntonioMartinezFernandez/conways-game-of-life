package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/conway"
	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/tui"
)

func main() {
	board := conway.NewBoard(150, 150)
	tuiDrawer := tui.NewDrawer(1, 10)

	tui.SetupTerm()
	defer tui.ResetTerm()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalChan)

	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	tuiDrawer.Draw(board.Generation(), board.Height(), board.Width(), board.Grid())

	for {
		select {
		case <-tick.C:
			board.NextGeneration()
			tuiDrawer.Draw(
				board.Generation(),
				board.Height(),
				board.Width(),
				board.Grid(),
			)

		case <-signalChan:
			return
		}
	}
}
