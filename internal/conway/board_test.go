package conway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoardWithNilGrid(t *testing.T) {
	board := NewBoard(2, 3, nil)

	if board.Height() != 2 || board.Width() != 3 {
		t.Fatalf("expected 2x3 board, got %dx%d", board.Height(), board.Width())
	}

	grid := board.Grid()
	assert.Equal(t, 2, len(grid))

	for row := range 2 {
		assert.Equal(t, 3, len(grid[row]))

		for column := range 3 {
			assert.False(t, grid[row][column])
		}
	}
}

func TestNewBoardWithGrid(t *testing.T) {
	grid := Grid{
		0: {0: true, 1: false},
		1: {0: false, 1: true},
	}

	board := NewBoard(2, 2, grid)
	assert.True(t, board.Grid()[0][0])
	assert.True(t, board.Grid()[1][1])
	assert.False(t, board.Grid()[0][1])
	assert.False(t, board.Grid()[1][0])
}
