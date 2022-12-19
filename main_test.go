package main

import (
	"fmt"
	"testing"
)

func p(x, y int) Point {
	return Point{x: x, y: y}
}

// 0 0 0
// 0 0 0  + [[0]]
// 0 0 0
func isValidAlignment(pattern [][]int, board [][]int, x, y int) bool {
	maxy := min(len(board)-1, y+len(pattern))
	for i := y; i < maxy; i++ {
		maxx := min(len(board[i])-1, x+len(pattern[maxy-i+y]))
		for j := x; j < maxx; j++ {

		}
	}
	return true
}

func printMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf(" %d ", m[i][j])
		}
		fmt.Println()
	}
}

func TestApplyPattern(t *testing.T) {
	for _, tc := range []struct {
		name      string
		life      *Life
		pattern   [][]int
		pos       Point
		validator func(pattern [][]int, state [][]int, x, y int) bool
	}{
		{
			name: "simple",
			life: NewLife(parseFlags()),
			pos:  p(1, 1),
			pattern: [][]int{
				{0, 1},
				{1, 1},
			},
		},
	} {
		tc.life.ApplyPattern(tc.pattern, tc.pos.x, tc.pos.y)
		cells := tc.life.GetCells()
		printMatrix(cells)
	}
}
