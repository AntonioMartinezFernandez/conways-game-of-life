package conway

import "sync"

type Board struct {
	mu         sync.Mutex
	height     int
	width      int
	generation int
	grid       Grid
}

// NewBoard creates a board of the given size. When grid is nil, a blank
// grid is created; otherwise the provided grid is used as the initial state.
func NewBoard(height, width int, grid Grid) *Board {
	if grid == nil {
		grid = NewGrid(height, width)
	}

	return &Board{
		mu:         sync.Mutex{},
		height:     height,
		width:      width,
		generation: 0,
		grid:       grid,
	}
}

func (b *Board) Height() int {
	return b.height
}

func (b *Board) Width() int {
	return b.width
}

func (b *Board) Generation() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.generation
}

func (b *Board) Grid() [][]bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.grid
}

func (b *Board) NextGeneration() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.generation++
	b.grid = calculateNextgen(b.grid)
}
