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

func Load(path string) (grid conway.Grid, height int, width int, err error) {
	file, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, 0, 0, nil
	}
	if err != nil {
		return nil, 0, 0, fmt.Errorf("open board file: %w", err)
	}
	defer file.Close()

	grid = make(conway.Grid)

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

		row := make(map[int]bool, width)
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

		grid[height] = row
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
