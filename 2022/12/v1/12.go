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
	steps1 := grid.walk(start, end, 0)
	fmt.Printf("steps: %v\n", steps1)
	fmt.Println()

	fmt.Println("Part 2:")
	steps2 := grid.shortestStart(end)
	fmt.Printf("steps: %v\n", steps2)
	fmt.Println()

	fmt.Printf("Part 1 again: steps: %v\n", steps1)

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

func (g grid) walk(start, end position, stepsToBeat int) (steps int) {
	if stepsToBeat == 0 {
		stepsToBeat = math.MaxInt
	}
	steps = stepsToBeat

	type next struct {
		position
		steps int
	}
	nextPositions := []next{{start, 0}}
	visited := []next{}

	alreadyVisited := func(current next) bool {
		for i, pos := range visited {
			if current.position == pos.position {
				if pos.steps <= current.steps {
					return true
				} else {
					visited = append(visited[:i], visited[i+1:]...)
					return false
				}
			}
		}
		return false
	}

	clearVisited := func() {
		for i := 0; i < len(visited); i++ {
			if visited[i].steps > steps {
				visited = append(visited[:i], visited[i+1:]...)
				i--
			}
		}
	}

	for len(nextPositions) > 0 {
		current := nextPositions[len(nextPositions)-1]
		nextPositions = nextPositions[:len(nextPositions)-1]

		if current.steps > steps {
			continue
		}

		if alreadyVisited(current) {
			continue
		}
		visited = append(visited, current)

		fmt.Printf("len(next): %v, len(visited): %v, least steps: %v\n", len(nextPositions), len(visited), steps)

		if current.position == end {
			if current.steps < steps {
				steps = current.steps
				clearVisited()
				continue
			}
		}

		current.steps++
		if current.row-1 >= 0 && g[current.row-1][current.col] <= g[current.row][current.col]+1 {
			nextPositions = append(nextPositions, next{position{current.row - 1, current.col}, current.steps})
		}
		if current.row+1 < len(g) && g[current.row+1][current.col] <= g[current.row][current.col]+1 {
			nextPositions = append(nextPositions, next{position{current.row + 1, current.col}, current.steps})
		}
		if current.col-1 >= 0 && g[current.row][current.col-1] <= g[current.row][current.col]+1 {
			nextPositions = append(nextPositions, next{position{current.row, current.col - 1}, current.steps})
		}
		if current.col+1 < len(g[current.row]) && g[current.row][current.col+1] <= g[current.row][current.col]+1 {
			nextPositions = append(nextPositions, next{position{current.row, current.col + 1}, current.steps})
		}
	}

	return
}

func (g grid) shortestStart(end position) (steps int) {
	steps = math.MaxInt

	for r, row := range g {
		for c, col := range row {
			if col == 0 {
				start := position{r, c}
				s := g.walk(start, end, steps)
				if s < steps {
					steps = s
				}
			}
		}
	}

	return
}
