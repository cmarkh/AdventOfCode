package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	grid := parse(input)

	steps := grid.standstill()
	fmt.Printf("gridlock in %v steps\n", steps)
	fmt.Println()
}

type grid [][]string

func parse(input string) (grid grid) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		row := []string{}
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}
	return
}

func (grid grid) print() {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) step() (new grid, movement bool) {
	for _, row := range g {
		newRow := make([]string, len(row))
		copy(newRow, row)
		new = append(new, newRow)
	}

	for r, row := range g {
		for c, col := range row {
			if col == ">" {
				if c == len(row)-1 {
					if row[0] == "." {
						new[r][0] = col
						new[r][c] = "."
						movement = true
					}
				} else {
					if row[c+1] == "." {
						new[r][c+1] = col
						new[r][c] = "."
						movement = true
					}
				}
			}
		}
	}

	g = new
	new = grid{}
	for _, row := range g {
		newRow := make([]string, len(row))
		copy(newRow, row)
		new = append(new, newRow)
	}

	for c := 0; c < len(g[0]); c++ {
		for r := 0; r < len(g); r++ {
			if g[r][c] != "v" {
				continue
			}
			if r+1 == len(g) {
				if g[0][c] == "." {
					new[0][c] = g[r][c]
					new[r][c] = "."
					movement = true
				}
			} else {
				if g[r+1][c] == "." {
					new[r+1][c] = g[r][c]
					new[r][c] = "."
					movement = true
				}
			}
		}
	}

	return
}

func (g grid) standstill() (steps int) {
	movement := false
	for steps = 1; ; steps++ {
		g, movement = g.step()
		if !movement {
			return
		}
	}
}
