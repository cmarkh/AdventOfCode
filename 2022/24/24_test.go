package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`

var test2 = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func TestPart1(t *testing.T) {
	grid := parse(test2)
	grid.print(position{0, 1})

	start := position{0, 1}
	end := position{len(grid) - 1, len(grid[0]) - 2}
	grid, steps := grid.expedition(start, end)
	grid.print(position{len(grid) - 1, len(grid[0]) - 2})
	fmt.Printf("reached the end in %v steps\n", steps)
	fmt.Println()
}

func TestBlizzards(t *testing.T) {
	grid := parse(test1)
	grid.print(position{0, 1})

	for i := 0; i < 5; i++ {
		grid = grid.blizzardsMove()
		fmt.Println(i)
		grid.print(position{0, 1})
	}
}

func TestCopy(t *testing.T) {
	grid := parse(test2)
	grid.print(position{0, 1})

	newGrid := grid.copy()
	newGrid.print(position{0, 1})
}

func TestPart2(t *testing.T) {
	grid := parse(test2)
	grid.print(position{0, 1})

	start := position{0, 1}
	end := position{len(grid) - 1, len(grid[0]) - 2}
	grid, steps1 := grid.expedition(start, end)
	grid, steps2 := grid.expedition(end, start)
	grid, steps3 := grid.expedition(start, end)
	grid.print(end)
	fmt.Printf("there and back and there again in %v steps\n", steps1+steps2+steps3)
	fmt.Println()
}
