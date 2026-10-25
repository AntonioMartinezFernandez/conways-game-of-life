package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Receive height and width from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <height> <width>")
		return
	}

	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid height:", os.Args[1])
		return
	}

	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid width:", os.Args[2])
		return
	}

	// Create a text file full of 0s with the specified height and width
	file, err := os.Create("board")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for h := range height {
		for range width {
			file.WriteString("0")
		}
		if h < height-1 {
			file.WriteString("\n")
		}
	}

	fmt.Println("Empty board file created successfully.")
}
