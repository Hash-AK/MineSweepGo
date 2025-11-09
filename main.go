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

var gameState string
var totalMines, totalTiles int

func (grid *Grid) revealAllMine() {
	for r := range grid.grid {
		for c := range grid.grid[r] {
			if grid.grid[r][c].isMine {
				grid.grid[r][c].isRevealed = true
			}
		}
	}
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
	if grid.grid[r][c].isMine {
		gameState = "lost"

	}

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
func (grid *Grid) checkWin() bool {
	for r := range grid.grid {
		for c := range grid.grid[r] {
			if !grid.grid[r][c].isMine && !grid.grid[r][c].isRevealed {
				return false
			}
		}
	}
	return true
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
	totalMines = 0
	flagNumbers := 0
	totalTiles = 0
	gameState = "playing"
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
				totalMines = totalMines + 1

			}

		}

	}
	for r := 0; r < heightInt; r++ {
		for c := 0; c < widthInt; c++ {
			surroundingMines := 0
			totalTiles = totalTiles + 1
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

	defstyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	selectedStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlue)
	mineStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed)
	flaggedStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOrange)
	logoStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen)
	linkStyle := tcell.StyleDefault.Foreground(tcell.ColorAqua).Underline(true)
	screen.SetStyle(defstyle)
	game.grid[selectedY][selectedX].isSelected = true
	for {
		screen.Clear()
		termWidth, termHeight := screen.Size()
		startX := termWidth/2 - widthInt
		startY := termHeight/2 - heightInt/2

		logoX := termWidth/2 - len(" |_|  |_|_|_||_\\___|___/\\_/\\_/\\___\\___| .__/\\___\\___/")/2
		printString(screen, logoX, startY-6, logoStyle, "  __  __ _          ___                      ___     ")
		printString(screen, logoX, startY-5, logoStyle, " |  \\/  (_)_ _  ___/ __|_ __ _____ ___ _ __ / __|___ ")
		printString(screen, logoX, startY-4, logoStyle, " | |\\/| | | ' \\/ -_)__ \\ V  V / -_) -_) '_ \\ (_ / _ \\")
		printString(screen, logoX, startY-3, logoStyle, " |_|  |_|_|_||_\\___|___/\\_/\\_/\\___\\___| .__/\\___\\___/")
		printString(screen, logoX, startY-2, logoStyle, "                                      |_|            ")
		instructions := "Arrow key to move selection | 'F' to flag | Enter to reveal"
		instructionsX := termWidth/2 - len(instructions)/2
		printString(screen, instructionsX, startY+heightInt+2, logoStyle, instructions)
		credits := "By: github.com/Hash-AK"
		creditsX := termWidth/2 - len(credits)/2
		printString(screen, creditsX, startY+heightInt+4, linkStyle, credits)
		game.grid[selectedY][selectedX].isSelected = true
		status := fmt.Sprintf("Flags Place/Total Mines : %d/%d", flagNumbers, totalMines)
		statusX := startX + widthInt*2 + 5
		statusY := termHeight / 2
		printString(screen, statusX, statusY, logoStyle, status)
		screen.SetContent(startX-2, startY-1, '╭', nil, defstyle)
		screen.SetContent(startX+widthInt*2, startY-1, '╮', nil, defstyle)

		screen.SetContent(startX-2, startY+heightInt, '╰', nil, defstyle)

		screen.SetContent(startX+widthInt*2, startY+heightInt, '╯', nil, defstyle)

		for x := -1; x < widthInt*2; x++ {
			screen.SetContent(startX+x, startY-1, '─', nil, defstyle)

		}
		for x := -1; x < widthInt*2; x++ {
			screen.SetContent(startX+x, startY+heightInt, '─', nil, defstyle)
		}
		for y := 0; y < heightInt; y++ {
			screen.SetContent(startX-2, startY+y, '│', nil, defstyle)
		}
		for y := 0; y < heightInt; y++ {
			screen.SetContent(startX+widthInt*2, startY+y, '│', nil, defstyle)
		}
		for r := 0; r < heightInt; r++ {
			screenX := startX
			for c := 0; c < widthInt; c++ {

				styleToUse := defstyle
				var charToDraw rune
				if game.grid[r][c].isSelected {
					if game.grid[r][c].neighborMines == 0 && game.grid[r][c].isRevealed && !game.grid[r][c].isMine {
						styleToUse = selectedStyle.Background(tcell.ColorBlue)

					} else {
						styleToUse = selectedStyle

					}
				}
				if !game.grid[r][c].isRevealed {
					if game.grid[r][c].isFlagged {
						charToDraw = 'f'
						if game.grid[r][c].isSelected {
							styleToUse = flaggedStyle.Background(tcell.ColorDarkGreen)

						} else {
							styleToUse = flaggedStyle
						}
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
		switch gameState {
		case "lost":
			msg := "GAME OVER"
			msgX := termWidth/2 - len(msg)/2
			msgY := termHeight / 2
			printString(screen, msgX, msgY, mineStyle.Bold(true), msg)
		case "won":
			msg := "YOU WIN!"
			msgX := termWidth/2 - len(msg)/2
			msgY := termHeight / 2
			printString(screen, msgX, msgY, logoStyle.Bold(true), msg)

		}

		screen.Show()
		event := screen.PollEvent()
		if gameState == "playing" {
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
					if gameState == "lost" {
						game.revealAllMine()
					} else {
						if game.checkWin() {
							gameState = "won"
						}
					}
				}
				if ev.Key() == tcell.KeyRune {
					if ev.Rune() == 'f' {
						if game.grid[selectedY][selectedX].isFlagged == false {
							game.grid[selectedY][selectedX].isFlagged = true
							flagNumbers = flagNumbers + 1
						} else {
							game.grid[selectedY][selectedX].isFlagged = false
							flagNumbers = flagNumbers - 1
						}

					}
				}
			case *tcell.EventResize:
				screen.Clear()
				termWidth, termHeight = screen.Size()

			}
		} else {
			switch ev := event.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return

				}
			}

		}
	}

}
