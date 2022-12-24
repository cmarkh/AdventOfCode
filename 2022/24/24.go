package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	t := time.Now()

	fmt.Println("Part 1:")
	grid := parse(input)
	start := position{0, 1}
	end := position{len(grid) - 1, len(grid[0]) - 2}
	grid, steps1 := grid.expedition(start, end)
	grid.print(position{len(grid) - 1, len(grid[0]) - 2})
	fmt.Printf("reached the end in %v minutes\n", steps1)
	fmt.Println()

	fmt.Println("Part 2:")
	grid, steps2 := grid.expedition(end, start)
	grid, steps3 := grid.expedition(start, end)
	grid.print(end)
	fmt.Printf("there and back and there again in %v minutes\n", steps1+steps2+steps3)
	fmt.Println()

	fmt.Printf("I meandered through in %v\n", time.Since(t))
}

type grid [][][]string

func parse(input string) (grid grid) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		row := [][]string{}
		for _, ch := range line {
			if ch == '.' {
				row = append(row, []string{})
			} else {
				row = append(row, []string{string(ch)})
			}
		}
		grid = append(grid, row)
	}
	return
}

func (grid grid) print(pos position) {
	for r, row := range grid {
		for c, col := range row {
			if r == pos.r && c == pos.c {
				fmt.Print("E")
				continue
			}
			if len(col) == 0 {
				fmt.Print(".")
			} else if len(col) > 1 {
				fmt.Print(len(col))
			} else {
				fmt.Print(col[0])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type position struct {
	r, c int
}
type q struct {
	g     grid
	pos   position
	steps int
	path  []position
}

func (g grid) expedition(start, end position) (newGrid grid, steps int) {
	steps = math.MaxInt
	queue := []q{{g, start, 0, []position{}}}
	visited := []q{}

	up := func(current q) {
		if current.pos.r > 0 && len(current.g[current.pos.r-1][current.pos.c]) == 0 {
			next := q{pos: current.pos, steps: current.steps}
			next.pos.r--
			if !slices.Contains(next.path, next.pos) {
				next.path = append(next.path, next.pos)
				next.g = current.g.copy()
				queue = append(queue, next)
			}
		}
	}
	down := func(current q) {
		if current.pos.r < len(current.g)-1 && len(current.g[current.pos.r+1][current.pos.c]) == 0 {
			next := q{pos: current.pos, steps: current.steps}
			next.pos.r++
			if !slices.Contains(next.path, next.pos) {
				next.path = append(next.path, next.pos)
				next.g = current.g.copy()
				queue = append(queue, next)
			}
		}
	}
	left := func(current q) {
		if current.pos.c > 0 && len(current.g[current.pos.r][current.pos.c-1]) == 0 {
			next := q{pos: current.pos, steps: current.steps}
			next.pos.c--
			if !slices.Contains(next.path, next.pos) {
				next.path = append(next.path, next.pos)
				next.g = current.g.copy()
				queue = append(queue, next)
			}
		}
	}
	right := func(current q) {
		if current.pos.c < len(current.g[current.pos.r])-1 && len(current.g[current.pos.r][current.pos.c+1]) == 0 {
			next := q{pos: current.pos, steps: current.steps}
			next.pos.c++
			if !slices.Contains(next.path, next.pos) {
				next.path = append(next.path, next.pos)
				next.g = current.g.copy()
				queue = append(queue, next)
			}
		}
	}
	wait := func(current q) {
		if len(current.g[current.pos.r][current.pos.c]) == 0 {
			next := q{pos: current.pos, steps: current.steps}
			next.g = current.g.copy()
			queue = append(queue, next)
		}
	}

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if current.pos == end {
			if current.steps < steps {
				steps = current.steps
				newGrid = current.g
				continue
			}
		}

		lowerBound := current.steps + abs(end.r-current.pos.r) + abs(end.c-current.pos.c)
		if lowerBound >= steps {
			continue
		}

		if current.duplicate(visited) {
			continue
		}
		visited = append(visited, current)

		fmt.Printf("current steps: %v, min steps so far: %v\n", current.steps, steps)
		//current.g.print(current.pos)

		current.steps++
		current.g = current.g.blizzardsMove()

		if end.r > start.r {
			up(current)
			left(current)
			wait(current)
			down(current)
			right(current)
		} else {
			down(current)
			right(current)
			wait(current)
			up(current)
			left(current)
		}
	}

	return
}

func (grid grid) blizzardsMove() (newGrid grid) {
	newGrid = make([][][]string, len(grid))
	for r := range newGrid {
		newGrid[r] = make([][]string, len(grid[0]))
	}

	for r, row := range grid {
		for c, col := range row {
			for _, bliz := range col {
				switch bliz {
				case ">":
					if len(grid[r][c+1]) > 0 && grid[r][c+1][0] == "#" {
						newGrid[r][1] = append(newGrid[r][1], bliz)
					} else {
						newGrid[r][c+1] = append(newGrid[r][c+1], bliz)
					}
				case "<":
					if len(grid[r][c-1]) > 0 && grid[r][c-1][0] == "#" {
						newGrid[r][len(newGrid[r])-2] = append(newGrid[r][len(newGrid[r])-2], bliz)
					} else {
						newGrid[r][c-1] = append(newGrid[r][c-1], bliz)
					}
				case "v":
					if len(grid[r+1][c]) > 0 && grid[r+1][c][0] == "#" || r+1 >= len(grid) {
						newGrid[1][c] = append(newGrid[1][c], bliz)
					} else {
						newGrid[r+1][c] = append(newGrid[r+1][c], bliz)
					}
				case "^":
					if len(grid[r-1][c]) > 0 && grid[r-1][c][0] == "#" || r-1 < 0 {
						newGrid[len(newGrid)-2][c] = append(newGrid[len(newGrid)-2][c], bliz)
					} else {
						newGrid[r-1][c] = append(newGrid[r-1][c], bliz)
					}
				case "#":
					newGrid[r][c] = append(newGrid[r][c], bliz)
				}
				// grid[r][c] = slices.Delete(grid[r][c], b, b+1)
			}
		}
	}

	return
}

func (grid grid) copy() (newGrid grid) {
	newGrid = make([][][]string, len(grid))
	for r := range grid {
		newGrid[r] = make([][]string, len(grid[r]))
		for c := range grid[r] {
			newGrid[r][c] = make([]string, len(grid[r][c]))
			copy(newGrid[r][c], grid[r][c])
		}
	}
	return
}

func (q q) duplicate(visited []q) bool {
	for _, vis := range visited {
		if vis.pos != q.pos {
			continue
		}
		if func() bool {
			for r := range q.g {
				for c := range q.g[r] {
					if slices.Compare(q.g[r][c], vis.g[r][c]) != 0 {
						return false
					}
				}
			}
			return true
		}() {
			if q.steps >= vis.steps {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
