package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	advent "temp/adventofcode/go"
)

var inputPart1 = "input_1.txt"

func main() {
	err := Part1()
	if err != nil {
		log.Fatal(err)
	}
}

func Part1() (err error) {
	grid, err := input(inputPart1)
	if err != nil {
		return
	}
	grid.Print()

	//grid.dijkstra()
	_, lowestRisks := grid.riskLevels()
	lowestRisks.Print()

	return
}

type RisksAtPoint struct {
	up, down, left, right int //one risk level for each possible direction, //you cannot move diagonally
}
type Risks [][]*RisksAtPoint

func (risks RisksAtPoint) lowest() int {
	lowest := risks.up
	if risks.down < lowest {
		lowest = risks.down
	}
	if risks.left < lowest {
		lowest = risks.left
	}
	if risks.right < lowest {
		lowest = risks.right
	}
	return lowest
}

func (risks Risks) toGrid() (grid Grid) {
	for r := range risks {
		row := []int{}
		for c := range risks {
			row = append(row, risks[r][c].lowest())
		}
		grid = append(grid, row)
	}
	return
}

func (grid Grid) NewRiskLevels() (risks Risks) {
	rowLength := len(grid[0])
	for range grid {
		//lowestRisks = append(lowestRisks, make([]int, rowLength))
		row := []*RisksAtPoint{}
		for c := 0; c < rowLength; c++ {
			row = append(row, &RisksAtPoint{})
		}
		risks = append(risks, row)
	}

	for r := range grid {
		risks[r][0].left = math.MaxInt //no left point so using max risk will serve as an ignore function
		risks[r][rowLength-1].right = math.MaxInt
	}
	for c := 0; c < rowLength; c++ {
		risks[0][c].up = math.MaxInt
		risks[len(grid)-1][c].down = math.MaxInt
	}

	return
}

// to determine the total risk of an entire path, add up the risk levels of each position you enter
func (grid Grid) riskLevels() (risks Risks, lowestRisks Grid) {
	risks = grid.NewRiskLevels()

	for {
		done := true
		for r := range risks {
			for c := range risks {
				if r == 0 && c == 0 {
					continue //skip starting point
				}
				if risks[r][c].up != 0 && //if found all risk levels for this point
					risks[r][c].down != 0 &&
					risks[r][c].left != 0 &&
					risks[r][c].right != 0 {
					continue
				}
			}
		}

		if done {
			break
		}
	}

	return risks, risks.toGrid()
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
