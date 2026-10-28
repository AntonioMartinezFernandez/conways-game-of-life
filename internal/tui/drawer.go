package tui

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
)

type Drawer struct {
	imageID  int
	cellSize int
}

func NewDrawer(imageID int, cellSize int) *Drawer {
	return &Drawer{
		imageID:  imageID,
		cellSize: cellSize,
	}
}

// Draw generates the image and displays it in the terminal.
func (t *Drawer) Draw(generation int, boardHeight int, boardWidth int, boardGrid [][]bool) {
	imageRGBA := t.createImage(boardHeight, boardWidth, boardGrid)

	data, err := t.encodePNG(imageRGBA)
	if err != nil {
		return
	}

	// Set cursor to top-left corner of the terminal before drawing the image
	fmt.Print("\033[H")

	// Kitty Graphics Protocol.
	//
	// a=T   -> transmit
	// f=100 -> PNG
	// t=d   -> direct
	// i=1   -> image ID
	// m=0   -> single transmission
	// q=2   -> suppress all terminal responses (they get echoed on screen)
	//
	// c=1, r=1 makes the image occupy one terminal cell regardless of its physical size
	fmt.Printf(
		"\033_Ga=T,f=100,i=%d,m=0,q=2,t=d;%s\033\\",
		t.imageID,
		data,
	)

	fmt.Printf(
		"\033[%d;%dH",
		boardHeight+2,
		1,
	)

	fmt.Printf("Ctrl+C: quit   Generation: %d", generation)
}

// createImage creates an image of size:
// width * cellSize × height * cellSize pixels
func (t *Drawer) createImage(boardHeight int, boardWidth int, grid [][]bool) *image.RGBA {
	imgHeight := boardHeight * t.cellSize
	imgWidth := boardWidth * t.cellSize

	img := image.NewRGBA(
		image.Rect(0, 0, imgWidth, imgHeight),
	)

	white := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	black := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}

	for y := range boardHeight {
		for x := range boardWidth {
			c := white

			if grid[y][x] {
				c = black
			}

			// The cell occupies exactly:
			// cellSize × cellSize pixels
			for py := range t.cellSize {
				for px := range t.cellSize {
					img.SetRGBA(
						x*t.cellSize+px,
						y*t.cellSize+py,
						c,
					)
				}
			}
		}
	}

	return img
}

func (t *Drawer) encodePNG(img image.Image) (string, error) {
	var buffer bytes.Buffer

	if err := png.Encode(&buffer, img); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}
