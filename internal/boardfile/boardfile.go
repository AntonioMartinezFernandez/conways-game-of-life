package boardfile

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/AntonioMartinezFernandez/conways-game-of-life/internal/conway"
)

// Load returns grid, height, width, err
func Load(path string) (conway.Grid, int, int, error) {
	file, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, 0, 0, nil
	}
	if err != nil {
		return nil, 0, 0, fmt.Errorf("open board file: %w", err)
	}
	defer file.Close()

	var grid conway.Grid
	var height int
	var width int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\r")
		if line == "" {
			return nil, 0, 0, fmt.Errorf("line %d: empty row", height+1)
		}

		if width == 0 {
			width = len(line)
		} else if len(line) != width {
			return nil, 0, 0, fmt.Errorf(
				"line %d: expected %d columns, got %d",
				height+1,
				width,
				len(line),
			)
		}

		row := make([]bool, width)
		for column := 0; column < len(line); column++ {
			switch line[column] {
			case '0':
				row[column] = false
			case '1':
				row[column] = true
			default:
				return nil, 0, 0, fmt.Errorf(
					"line %d, column %d: invalid character %q",
					height+1,
					column+1,
					line[column],
				)
			}
		}

		grid = append(grid, row)
		height++
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("read board file: %w", err)
	}
	if height == 0 {
		return nil, 0, 0, fmt.Errorf("board file %q is empty", path)
	}

	return grid, height, width, nil
}
