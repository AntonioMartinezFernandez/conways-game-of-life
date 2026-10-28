package conway

func calculateNextgen(grid Grid) Grid {
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	nextGenGrid := NewGrid(height, width)

	//! TODO: Implement real logic with Conway's Game of Life
	for h := range height {
		for w := range width {
			nextGenGrid[h][w] = !grid[h][w]
		}
	}

	return nextGenGrid
}
