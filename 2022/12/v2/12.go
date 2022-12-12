package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	t := time.Now()

	grid, start, end := parse(input)

	fmt.Println("Part 1:")
	steps1 := grid.startAtEnd(start, end, false)
	fmt.Printf("steps: %v\n", steps1)
	fmt.Println()

	fmt.Println("Part 2:")
	steps2 := grid.startAtEnd(start, end, true)
	fmt.Printf("steps: %v\n", steps2)
	fmt.Println()

	//fmt.Printf("Part 1 again: steps: %v\n", steps1)

	fmt.Printf("time ellapsed: %v\n", time.Since(t))
}

type grid [][]int
type position struct {
	row, col int
}

func parse(input string) (grid grid, start, end position) {
	lines := strings.Split(input, "\n")

	for r, line := range lines {
		if line == "" {
			continue
		}
		row := []int{}
		for c, ch := range line {
			if ch == 'S' {
				start.row = r
				start.col = c
				row = append(row, 0)
				continue
			}
			if ch == 'E' {
				end.row = r
				end.col = c
				row = append(row, int('z'-'a'))
				continue
			}
			row = append(row, int(ch-'a'))
		}
		grid = append(grid, row)
	}

	return
}

func (g grid) print() {
	for _, row := range g {
		for _, col := range row {
			fmt.Print(col)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) startAtEnd(start, end position, anyStart bool) (leastSteps int) {
	leastSteps = math.MaxInt
	stepsToEnd := make(map[position]int) //map[position]steps to end
	type next struct {
		position
		steps int
	}
	nextPositions := []next{{end, 0}}

	for len(nextPositions) > 0 {
		current := nextPositions[len(nextPositions)-1]
		nextPositions = nextPositions[:len(nextPositions)-1]

		if current.steps > leastSteps {
			continue
		}

		if steps, ok := stepsToEnd[current.position]; ok {
			if current.steps >= steps {
				continue
			}
		}
		stepsToEnd[current.position] = current.steps

		//fmt.Printf("len(next): %v, len(visited): %v, least steps: %v\n", len(nextPositions), len(stepsToEnd), leastSteps)

		if (anyStart && g[current.row][current.col] == 0) || (!anyStart && current.position == start) {
			if current.steps < leastSteps {
				leastSteps = current.steps
			}
			continue
		}

		current.steps++
		if current.row-1 >= 0 && g[current.row-1][current.col] >= g[current.row][current.col]-1 {
			nextPositions = append(nextPositions, next{position{current.row - 1, current.col}, current.steps})
		}
		if current.row+1 < len(g) && g[current.row+1][current.col] >= g[current.row][current.col]-1 {
			nextPositions = append(nextPositions, next{position{current.row + 1, current.col}, current.steps})
		}
		if current.col-1 >= 0 && g[current.row][current.col-1] >= g[current.row][current.col]-1 {
			nextPositions = append(nextPositions, next{position{current.row, current.col - 1}, current.steps})
		}
		if current.col+1 < len(g[current.row]) && g[current.row][current.col+1] >= g[current.row][current.col]-1 {
			nextPositions = append(nextPositions, next{position{current.row, current.col + 1}, current.steps})
		}
	}

	return
}
