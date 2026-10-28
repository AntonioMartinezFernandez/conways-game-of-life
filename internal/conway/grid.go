package conway

type Grid [][]bool

func NewGrid(height int, width int) Grid {

	grid := make([][]bool, height)
	for h := range height {
		grid[h] = make([]bool, width)
		for w := range width {
			grid[h][w] = false
		}
	}
	return grid
}
