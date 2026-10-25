package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/boardfile"
	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/conway"
	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/tui"
)

const (
	boardPath     = "board"
	defaultHeight = 10
	defaultWidth  = 10
)

func main() {
	grid, height, width, err := boardfile.Load(boardPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load %s: %v\n", boardPath, err)
		os.Exit(1)
	}
	if grid == nil {
		height, width = defaultHeight, defaultWidth
	}

	board := conway.NewBoard(height, width, grid)
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
