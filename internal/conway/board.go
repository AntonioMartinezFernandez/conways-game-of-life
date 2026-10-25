package conway

import "sync"

type Board struct {
	mu         sync.Mutex
	height     int
	width      int
	generation int
	grid       Grid
}

func NewBoard(height, width int) *Board {
	return &Board{
		mu:         sync.Mutex{},
		height:     height,
		width:      width,
		generation: 0,
		grid:       newGrid(height, width),
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

func (b *Board) Grid() map[int]map[int]bool {
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
