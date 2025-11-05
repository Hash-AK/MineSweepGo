package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

type Cell struct {
	isMine        bool
	isRevealed    bool
	isFlagged     bool
	neighborMines int
}
type Grid struct {
	grid [][]Cell
}

func main() {
	var height string
	var width string
	if len(os.Args) > 2 {
		height = os.Args[1]
		width = os.Args[2]
	} else {
		fmt.Println("Usage : go run . <height> <width>. Using default 10x10.")
		height = "10"
		width = "10"
	}
	heightInt, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println("Invalid height. Using default 10.")
	}

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		fmt.Println("Invalid width. Using default 10.")
	}
	fmt.Printf("Height : %s", height)
	fmt.Println()
	fmt.Printf("Width : %s", width)
	fmt.Println()
	game := Grid{
		grid: make([][]Cell, heightInt),
	}
	for i := range game.grid {
		game.grid[i] = make([]Cell, widthInt)
	}
	for r := 0; r < heightInt; r++ {
		for c := 0; c < widthInt; c++ {
			game.grid[r][c].neighborMines = rand.IntN(5)
		}

	}
	for r := 0; r < heightInt; r++ {
		for c := 0; c < widthInt; c++ {
			fmt.Print(game.grid[r][c].neighborMines)
			fmt.Print("  ")
		}
		fmt.Println()
	}

}
