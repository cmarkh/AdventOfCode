package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go/2021"
)

var (
	//lint:ignore U1000 ignore unused
	inputBothParts = "input_2.txt"
	//lint:ignore U1000 ignore unused
	testPart1 = "input_1.txt"
	//lint:ignore U1000 ignore unused
	testPart2 = "test_2.txt"
	//lint:ignore U1000 ignore unused
	test3 = "test_3.txt"
)

func main() {
	err := Part1()
	if err != nil {
		log.Fatal(err)
	}

	err = Part2()
	if err != nil {
		log.Fatal(err)
	}
}

func Part1() (err error) {
	fmt.Printf("Part 1:\n\n")

	grid, err := input(inputBothParts)
	if err != nil {
		return
	}
	//grid.Print()

	riskGrid, totalRisk := grid.LowestRisk()
	grid.PrintPath(riskGrid)
	fmt.Printf("Lowest possible risk: %v\n", totalRisk)
	fmt.Println()

	return
}

func Part2() (err error) {
	fmt.Printf("Part 2: \n\n")

	grid, err := input(inputBothParts)
	if err != nil {
		return
	}
	//grid.Print()

	grid = grid.FiveX()

	riskGrid, totalRisk := grid.LowestRisk()
	grid.PrintPath(riskGrid)
	fmt.Printf("Lowest possible risk: %v\n", totalRisk)

	return
}

type Grid [][]int

func input(path string) (grid Grid, err error) {
	input, err := advent.ReadInput(path)
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
