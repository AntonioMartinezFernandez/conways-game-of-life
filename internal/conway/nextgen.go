package conway

func calculateNextgen(grid Grid) Grid {
	nextGenGrid := grid

	//! TODO: Implement real logic with Conway's Game of Life
	if nextGenGrid[0][0] {
		nextGenGrid[0][0] = false
	} else {
		nextGenGrid[0][0] = true
	}

	return nextGenGrid
}
