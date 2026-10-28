package conway

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGen(t *testing.T) {
	t.Run("empty grid", func(t *testing.T) {
		grid := NewGrid(5, 5)
		nextGen := calculateNextgen(grid)
		fmt.Println(grid)
		assert.Equal(t, grid, nextGen)
	})
}
