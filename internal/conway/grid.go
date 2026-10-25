package conway

type Grid map[int]map[int]bool

func newGrid(height int, width int) Grid {
	grid := make(map[int]map[int]bool)
	for h := range height {
		grid[h] = make(map[int]bool)
		for w := range width {
			grid[h][w] = false
		}
	}
	return grid
}
