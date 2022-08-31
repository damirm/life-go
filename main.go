package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	HEIGHT = 10
	WIDTH  = 10
)

type Point struct {
	x, y int
}

type Matrix [][]int

type Life struct {
	cells  [][]int
	width  int
	height int
	alive  int
}

func (l *Life) GetWidth() int {
	return l.width
}

func (l *Life) GetHeight() int {
	return l.height
}

func (l *Life) GetCells() [][]int {
	return l.cells
}

func (l *Life) Tick() int {
	l.alive = 0

	var updated int
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			nbors := l.CountAliveNeighbors(x, y)
			cstate := l.cells[y][x]
			nstate := 0

			if cstate == 0 && nbors == 3 {
				nstate = 1
				l.alive++
			} else if cstate == 1 && (nbors == 2 || nbors == 3) {
				nstate = 1
				l.alive++
			}

			if cstate != nstate {
				updated++
			}

			l.SetCell(x, y, nstate)
		}
	}
	return updated
}

func (l *Life) SetCellAlive(x, y int) {
	l.SetCell(x, y, 1)
}

func (l *Life) SetCell(x, y, value int) {
	l.cells[y][x] = value
}

func (l *Life) IsAlive(x, y int) bool {
	return l.cells[y][x] > 0
}

func (l *Life) IsAnybodyAlive() bool {
	return l.alive > 0
}

func (l *Life) CountAliveNeighbors(x, y int) int {
	var res int

	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			if iy != 0 || ix != 0 {
				nx := x + ix
				ny := y + iy

				if nx >= 0 && nx < l.width && ny >= 0 && ny < l.height {
					if l.cells[ny][nx] == 1 {
						res++
					}
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func random(from, to int) int {
	return rand.Intn(to-from) + from
}

// ApplyPattern applies given pattern matrix to cells matrix.
// 0 0 0 0 0   0 1 0   0 1 0 0 0
// 0 0 0 0 0 + 1 0 1 = 1 0 1 0 0
// 0 0 0 0 0   0 1 0   0 1 0 0 0
// 0 0 0 0 0           0 0 0 0 0
func (l *Life) ApplyPattern(pattern Matrix, x, y int) {
	my := min(len(l.cells)-1, y+len(pattern)-1)
	pi, pj := 0, 0
	for i := y; i <= my; i++ {
		mx := min(len(l.cells[i])-1, x+len(pattern[pi])-1)
		for j := x; j <= mx; j++ {
			l.cells[i][j] = pattern[pi][pj]
			pj++
		}
		pj, pi = 0, pi+1
	}
}

func (l *Life) canPutPatternThere(pattern Matrix, x, y int) bool {
	for i := y; i < y+len(pattern); i++ {
		for j := x; j < x+len(pattern); j++ {
			if l.cells[i][j] == 1 {
				return false
			}
		}
	}
	return true
}

func (l *Life) ApplyPatternToRandomPoint(pattern Matrix, maxTries int) bool {
	for try := 0; try < maxTries; try++ {
		x, y := random(0, len(l.cells)-len(pattern)-2), random(0, len(l.cells)-len(pattern)-2)
		if l.canPutPatternThere(pattern, x, y) {
			l.ApplyPattern(pattern, x, y)
			return true
		}
	}

	return false
}

func NewLife(width, height int) *Life {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
	}
	return &Life{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func printLife(life *Life) {
	for y := 0; y < life.GetHeight(); y++ {
		fmt.Println()
		for x := 0; x < life.GetWidth(); x++ {
			chr := " "
			if life.IsAlive(x, y) {
				chr = "."
			}
			fmt.Print(chr)
		}
	}
}

var patterns = []Matrix{
	{
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	},
	{
		{1},
		{1},
		{1},
	},
	{
		{1, 1},
		{1, 0},
	},
	{
		{0, 1, 0},
		{1, 1, 1},
	},
}

func loop(w, h int) {
	life := NewLife(w, h)
	reader := bufio.NewReader(os.Stdin)

	for _, pattern := range patterns {
		if ok := life.ApplyPatternToRandomPoint(pattern, 10); !ok {
			fmt.Println("Failed to apply pattern :-(")
		}
	}

	for {
		printLife(life)

		life.Tick()

		if !life.IsAnybodyAlive() {
			break
		}

		reader.ReadString('\n')
	}
}

func main() {
	loop(WIDTH, HEIGHT)
}
