package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go"
)

var inputPath = "../input.txt"

func main() {
	grid, err := Input()
	if err != nil {
		log.Fatal(err)
	}
	grid.Print()

	grid = grid.FiveX()
	//grid.Print()

	leastRisk := grid.LowestRisk()
	fmt.Printf("least risk: %d\n\n", leastRisk)

}

type Grid [][]int

func Input() (grid Grid, err error) {
	input, err := advent.ReadInput(inputPath)
	if err != nil {
		return
	}

	for _, line := range input {
		row := []int{}
		for _, cell := range line {
			var num int
			num, err = strconv.Atoi(string(cell))
			if err != nil {
				return
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	return
}

func (grid Grid) Print() {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
	fmt.Println()
}

type Point struct {
	Row, Col int
	Risk     int
}

func (grid Grid) LowestRisk() (risk int) {
	cost := Grid{} //cost to get to a given point (ie sum of least risk path to that point)

	for r, row := range grid {
		cost = append(cost, make([]int, len(row)))
		for c := range row {
			if r == 0 && c == 0 {
				continue
			}
			if r == 0 {
				cost[r][c] = grid[r][c] + cost[r][c-1]
				continue
			}
			if c == 0 {
				cost[r][c] = grid[r][c] + cost[r-1][c]
				continue
			}
			if cost[r-1][c] < cost[r][c-1] {
				cost[r][c] = grid[r][c] + cost[r-1][c]
				continue
			}
			cost[r][c] = grid[r][c] + cost[r][c-1]
			//fmt.Printf("%d: %d\n", c, cost[r][c])
		}
		//for c := range grid[r-1] {
		//	fmt.Println(c)
		//}
	}

	return cost[len(grid)-1][len(grid[0])-1] //risk at end point
}

func (grid Grid) FiveX() Grid {
	for r, row := range grid {
		lastRow := make([]int, len(row))
		copy(lastRow, row)
		for i := 1; i < 5; i++ {
			newRow := make([]int, len(lastRow))
			copy(newRow, lastRow)
			for c := range newRow {
				newRow[c]++
				if newRow[c] > 9 {
					newRow[c] = 1
				}
			}
			row = append(row, newRow...)
			copy(lastRow, newRow)
		}
		grid[r] = row
	}

	original := [][]int{}
	for _, row := range grid {
		original = append(original, row)
	}

	for i := 1; i < 5; i++ {
		for _, row := range original {
			newRow := []int{}
			for _, c := range row {
				newC := c + i
				if newC > 9 {
					newC -= 9
				}
				newRow = append(newRow, newC)
			}
			grid = append(grid, newRow)
		}
	}

	return grid
}
