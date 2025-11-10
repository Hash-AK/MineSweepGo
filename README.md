## MineSweepGo
A recreation of the iconic MineSweeper in Golang, in the terminal! This was made for [Siege by HackClub](https://siege.hackclub.com). It uses the Tcell Go module to facilitate operation with the terminal and the TUI interface.

![demo img](/minesweepgo.png)

## How To Play
Install a binary from the Release page, or do the following :
```bash
git clone https://github.com/Hash-AK/MineSweepGo

cd MineSweepGo

go run . {ARGUMENTS HERE}
```
The arguments workboth for the bianries and the manual way, and they are the height and width of the grid in blocks, example :
```bash
go run . 10 10
```
or :
```bash
./MineSweepGo-linux-amd64 10 10
```
if no entry is provided it defaults to 10x10

In the game, you can reveal tiles using Enter, flag a tile with 'f', and quit using either Ctrl+C or Escape.

Have fun!
