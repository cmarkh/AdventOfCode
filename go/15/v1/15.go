package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go"

	"golang.org/x/exp/slices"
)

var inputPath = advent.BasePath + "15/input.txt"

func main() {
	grid, err := Input()
	if err != nil {
		log.Fatal(err)
	}
	grid.Print()

	leastRisk := grid.ChartPaths()
	fmt.Println("least risk:")
	leastRisk.Print()

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

type Path struct {
	Points    []Point
	TotalRisk int
}

type Point struct {
	Row, Col int
	Risk     int
}

func (grid Grid) ChartPaths() (leastRisk Path) {
	var move func(partial Path, next Point)

	lowestRisk := len(grid) * len(grid[0])

	nextPoint := func(partial Path, next Point) {
		if partial.TotalRisk > lowestRisk {
			return
		}
		if next.Row < 0 || next.Col < 0 {
			return
		}
		if next.Row >= len(grid) || next.Col >= len(grid[next.Row]) {
			return
		}
		next.Risk = grid[next.Row][next.Col]
		if slices.Contains(partial.Points, next) {
			return
		}
		//partial.Print()
		move(partial, next)
	}

	move = func(partial Path, next Point) {
		partial.Points = append(partial.Points, next)
		if next.Row != 0 || next.Col != 0 { //don't count risk at start point 0,0
			partial.TotalRisk += next.Risk
		}

		if next.Row == len(grid)-1 && next.Col == len(grid[next.Row])-1 {
			if partial.TotalRisk < lowestRisk {
				lowestRisk = partial.TotalRisk
				leastRisk = partial
			}
			return
		}

		last := next

		next = Point{last.Row + 1, last.Col, 0}
		nextPoint(partial, next)

		next = Point{last.Row, last.Col + 1, 0}
		nextPoint(partial, next)

		next = Point{last.Row - 1, last.Col, 0}
		nextPoint(partial, next)

		next = Point{last.Row, last.Col - 1, 0}
		nextPoint(partial, next)
	}

	move(Path{}, Point{0, 0, grid[0][0]})

	return
}

func (path Path) Print() {
	for i, point := range path.Points {
		if i != 0 {
			fmt.Print(" -> ")
		}
		fmt.Printf("%d,%d", point.Row, point.Col)
	}
	fmt.Printf("\ntotal risk: %d\n", path.TotalRisk)
}
