package boardfile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeBoard(t *testing.T, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "board")
	err := os.WriteFile(path, []byte(content), 0o600)
	assert.NoError(t, err)

	return path
}

func TestLoad(t *testing.T) {
	grid, height, width, err := Load(writeBoard(t, "01\n10\n11\n"))

	assert.NoError(t, err)
	assert.Equal(t, 3, height)
	assert.Equal(t, 2, width)

	cells := []struct {
		row  int
		col  int
		want bool
	}{
		{0, 0, false},
		{0, 1, true},
		{1, 0, true},
		{1, 1, false},
		{2, 0, true},
		{2, 1, true},
	}
	for _, cell := range cells {
		assert.Equal(
			t,
			cell.want,
			grid[cell.row][cell.col],
			"cell (%d,%d)",
			cell.row,
			cell.col,
		)
	}
}

func TestLoadMissingFile(t *testing.T) {
	grid, height, width, err := Load(filepath.Join(t.TempDir(), "missing"))

	assert.NoError(t, err)
	assert.Nil(t, grid)
	assert.Zero(t, height)
	assert.Zero(t, width)
}

func TestLoadCRLF(t *testing.T) {
	grid, height, width, err := Load(writeBoard(t, "01\r\n10\r\n"))

	assert.NoError(t, err)
	assert.Equal(t, 2, height)
	assert.Equal(t, 2, width)
	assert.True(t, grid[0][1])
	assert.True(t, grid[1][0])
}

func TestLoadInvalidBoards(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{name: "ragged rows", content: "010\n01\n"},
		{name: "invalid character", content: "01\n0x\n"},
		{name: "empty row", content: "01\n\n10\n"},
		{name: "empty file", content: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := Load(writeBoard(t, tt.content))
			assert.Error(t, err)
		})
	}
}
