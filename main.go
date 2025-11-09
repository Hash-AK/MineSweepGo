package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type Cell struct {
	isMine        bool
	isRevealed    bool
	isFlagged     bool
	neighborMines int
	isSelected    bool
}
type Grid struct {
	grid [][]Cell
}

func (grid *Grid) revealCell(r, c int) {
	height := len(grid.grid)
	if height == 0 {
		return
	}
	width := len(grid.grid[0])
	if r < 0 || r >= height || c < 0 || c >= width {
		return
	}
	if grid.grid[r][c].isRevealed {
		return
	}
	grid.grid[r][c].isRevealed = true
	if grid.grid[r][c].neighborMines == 0 && grid.grid[r][c].isMine == false {
		grid.revealCell(r-1, c-1)
		grid.revealCell(r-1, c)
		grid.revealCell(r-1, c+1)
		grid.revealCell(r, c-1)
		grid.revealCell(r, c+1)
		grid.revealCell(r+1, c)
		grid.revealCell(r+1, c-1)
		grid.revealCell(r+1, c+1)
	}

}
func printString(screen tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		screen.SetContent(x, y, c, comb, style)
		x += w
	}
}
func main() {
	selectedX := 0
	selectedY := 0
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("Error initializing screen : %s\n", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("Failed to initialize screen : %s", err)
	}
	defer screen.Fini()
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
			randSeed := rand.IntN(5)
			if randSeed == 0 {
				game.grid[r][c].isMine = true

			}
		}

	}
	for r := 0; r < heightInt; r++ {
		for c := 0; c < widthInt; c++ {
			surroundingMines := 0
			if c > 0 {
				if game.grid[r][c-1].isMine {
					surroundingMines = surroundingMines + 1
				}
				if r > 0 {
					if game.grid[r-1][c-1].isMine {
						surroundingMines = surroundingMines + 1
					}
				}
				if r < heightInt-1 {
					if game.grid[r+1][c-1].isMine {
						surroundingMines = surroundingMines + 1
					}
				}

			}
			if c < widthInt-1 {
				if game.grid[r][c+1].isMine {
					surroundingMines = surroundingMines + 1
				}
				if r > 0 {
					if game.grid[r-1][c+1].isMine {
						surroundingMines = surroundingMines + 1
					}

				}
				if r < heightInt-1 {
					if game.grid[r+1][c+1].isMine {
						surroundingMines = surroundingMines + 1
					}
				}
			}
			if r > 0 {
				if game.grid[r-1][c].isMine {
					surroundingMines = surroundingMines + 1
				}
			}
			if r < heightInt-1 {
				if game.grid[r+1][c].isMine {
					surroundingMines = surroundingMines + 1
				}
			}
			game.grid[r][c].neighborMines = surroundingMines
		}
	}
	/*
		for r := 0; r < heightInt; r++ {
			for c := 0; c < widthInt; c++ {
				if game.grid[r][c].isMine {
					fmt.Print("* ")
				} else {
					fmt.Print(". ")
				}
			}
			fmt.Println()
		}
		for r := 0; r < heightInt; r++ {
			for c := 0; c < widthInt; c++ {
				fmt.Printf("%d ", game.grid[r][c].neighborMines)
			}
			fmt.Println()
		}
	*/
	defstyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	selectedStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlue)
	mineStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	flaggedStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrange)
	screen.SetStyle(defstyle)
	game.grid[selectedY][selectedX].isSelected = true
	for {
		screen.Clear()

		termWidth, termHeight := screen.Size()
		startX := termWidth/2 - widthInt
		startY := termHeight/2 - heightInt/2
		game.grid[selectedY][selectedX].isSelected = true

		for r := 0; r < heightInt; r++ {
			screenX := startX
			for c := 0; c < widthInt; c++ {

				styleToUse := defstyle
				var charToDraw rune
				if game.grid[r][c].isSelected {
					styleToUse = selectedStyle
				}
				if !game.grid[r][c].isRevealed {
					if game.grid[r][c].isFlagged {
						charToDraw = 'f'
						styleToUse = flaggedStyle
					} else {
						charToDraw = '■'
					}
				} else {
					if game.grid[r][c].isMine {
						charToDraw = '*'
						styleToUse = mineStyle
					} else if game.grid[r][c].neighborMines == 0 {
						charToDraw = ' '
					} else {
						charToDraw = rune(strconv.Itoa(game.grid[r][c].neighborMines)[0])
					}
				}
				screen.SetContent(screenX, r+startY, charToDraw, nil, styleToUse)

				screen.SetContent(screenX+1, r+startY, ' ', nil, defstyle)

				screenX += 2

			}
		}
		screen.Show()
		event := screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return

			}
			if ev.Key() == tcell.KeyRight {
				if selectedX < widthInt-1 {
					game.grid[selectedY][selectedX].isSelected = false

					selectedX = selectedX + 1
				}
			}
			if ev.Key() == tcell.KeyLeft {
				if selectedX > 0 {
					game.grid[selectedY][selectedX].isSelected = false
					selectedX = selectedX - 1
				}
			}
			if ev.Key() == tcell.KeyDown {
				if selectedY < heightInt-1 {
					game.grid[selectedY][selectedX].isSelected = false
					selectedY = selectedY + 1
				}

			}
			if ev.Key() == tcell.KeyUp {
				if selectedY > 0 {
					game.grid[selectedY][selectedX].isSelected = false
					selectedY = selectedY - 1
				}
			}
			if ev.Key() == tcell.KeyEnter {
				game.revealCell(selectedY, selectedX)
			}
			if ev.Key() == tcell.KeyRune {
				if ev.Rune() == 'f' {
					if game.grid[selectedY][selectedX].isFlagged == false {
						game.grid[selectedY][selectedX].isFlagged = true
					} else {
						game.grid[selectedY][selectedX].isFlagged = false
					}

				}
			}
		case *tcell.EventResize:
			screen.Clear()
			termWidth, termHeight = screen.Size()

		}
	}

}
