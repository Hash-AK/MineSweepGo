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
	if len(os.Args[1]) > 0 && len(os.Args[2]) > 0 {
		height = os.Args[1]
		width = os.Args[2]
	}
	widthInt, _ := strconv.Atoi(width)
	heightInt, _ := strconv.Atoi(height)
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
