package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

const (
	width    = 50
	height   = 50
	cellSize = 15

	imageID = 1
)

var grid = [height][width]bool{
	{false, false, false, false, false, false},
	{false, true, false, false, true, false},
	{false, false, false, false, false, false},
	{false, false, false, false, false, false},
	{false, true, false, false, true, false},
	{false, false, true, true, false, false},
	{false, false, false, false, false, false},
}

func main() {
	generation := 0

	// Set stdin to raw mode for receiving key presses
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer func() {
		term.Restore(int(os.Stdin.Fd()), oldState)

		// Show cursor and reset terminal
		fmt.Print("\033[?25h")
		fmt.Print("\033[0m")
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}()

	// Clean screen and hide cursor
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Print("\033[?25l")

	// Reset terminal on SIGINT or SIGTERM
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		os.Exit(0)
	}()

	draw(imageID, generation)

	for {
		key, err := readKey()
		if err != nil {
			return
		}

		switch key {
		case "space":
			grid[0][0] = !grid[0][0]

		case "q", "Q":
			return

		default:
			continue
		}

		generation++
		draw(imageID, generation)

		time.Sleep(time.Millisecond)
	}
}

// draw generates the image and displays it in the terminal.
func draw(imageID int, generation int) {
	imageRGBA := createImage()

	data, err := encodePNG(imageRGBA)
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
	//
	// c=1, r=1 makes the image occupy one terminal cell regardless of its physical size
	fmt.Printf(
		"\033_Ga=T,f=100,i=%d,m=0,t=d;%s\033\\",
		imageID,
		data,
	)

	fmt.Printf(
		"\033[%d;%dH",
		height+2,
		1,
	)

	fmt.Printf("Space: play/pause   Q: quit   Generation: %d", generation)
}

// createImage creates an image of size:
// width * cellSize × height * cellSize pixels
func createImage() *image.RGBA {
	imgWidth := width * cellSize
	imgHeight := height * cellSize

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

	for y := range height {
		for x := range width {
			c := white

			if grid[y][x] {
				c = black
			}

			// The cell occupies exactly:
			// cellSize × cellSize pixels
			for py := range cellSize {
				for px := range cellSize {
					img.SetRGBA(
						x*cellSize+px,
						y*cellSize+py,
						c,
					)
				}
			}
		}
	}

	return img
}

func encodePNG(img image.Image) (string, error) {
	var buffer bytes.Buffer

	if err := png.Encode(&buffer, img); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// readKey converts ANSI keyboard sequences into simple events.
func readKey() (string, error) {
	buf := make([]byte, 1)

	if _, err := os.Stdin.Read(buf); err != nil {
		return "", err
	}

	switch buf[0] {
	case ' ':
		return "space", nil

	case 'q', 'Q':
		return string(buf[0]), nil
	}

	return "", nil
}
