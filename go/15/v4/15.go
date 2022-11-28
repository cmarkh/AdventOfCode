package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go"
)

var (
	//lint:ignore U1000 ignore unused
	inputPart1 = "input_1.txt"
	//lint:ignore U1000 ignore unused
	inputPart2 = "input_2.txt"
	//lint:ignore U1000 ignore unused
	testPart2 = "test_2.txt"
)

func main() {
	grid, err := input(inputPart1)
	if err != nil {
		log.Fatal(err)
	}
	grid.Print()

	riskGrid, totalRisk := grid.LowestRisk()
	PrintGrids(riskGrid, grid)
	fmt.Printf("Lowest possible risk: %v\n", totalRisk)
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
