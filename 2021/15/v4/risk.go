package main

import (
	"fmt"
	"math"
)

type Point struct {
	Row int
	Col int
}
type RiskPoint struct {
	Risk             int
	Predecessor      Point
	OnLowestRiskPath bool
}
type RiskGrid [][]RiskPoint

func (grid Grid) LowestRisk() (riskGrid RiskGrid, totalRisk int) {
	rowLength := len(grid[0])
	riskGrid = make(RiskGrid, len(grid))
	for r := range riskGrid {
		riskGrid[r] = make([]RiskPoint, rowLength)
		for c := 0; c < rowLength; c++ {
			riskGrid[r][c].Risk = math.MaxInt //start all risk at ~infinite and will override with lowest risk available
		}
	}
	riskGrid[0][0].Risk = 0 //starting point

	for r := range grid {
		//traverse down
		if r > 0 {
			for c := 0; c < rowLength; c++ {
				if riskGrid[r-1][c].Risk+grid[r][c] < riskGrid[r][c].Risk {
					riskGrid[r][c].Risk = riskGrid[r-1][c].Risk + grid[r][c]
					riskGrid[r][c].Predecessor = Point{r - 1, c}
				}
			}
		}

		/* //traverse up
		if r < rowLength-1 {
			for c := 0; c < rowLength; c++ {
				if riskGrid[r+1][c].Risk+grid[r][c] < riskGrid[r][c].Risk {
					riskGrid[r][c].Risk = riskGrid[r+1][c].Risk + grid[r][c]
					riskGrid[r][c].Predecessor = Point{r + 1, c}
				}
				fmt.Printf("%v, %v\n", r, c)
			}
		}*/

		//traverse right
		for c := 1; c < rowLength; c++ {
			if riskGrid[r][c-1].Risk+grid[r][c] < riskGrid[r][c].Risk {
				riskGrid[r][c].Risk = riskGrid[r][c-1].Risk + grid[r][c]
				riskGrid[r][c].Predecessor = Point{r, c - 1}
			}
		}

		//traverse left
		for c := rowLength - 2; c >= 0; c-- {
			if riskGrid[r][c+1].Risk+grid[r][c] < riskGrid[r][c].Risk {
				riskGrid[r][c].Risk = riskGrid[r][c+1].Risk + grid[r][c]
				riskGrid[r][c].Predecessor = Point{r, c + 1}
			}
		}
	}

	riskGrid = riskGrid.MarkPath()

	return riskGrid, riskGrid[len(grid)-1][rowLength-1].Risk
}

func (riskGrid RiskGrid) MarkPath() RiskGrid {
	point := Point{len(riskGrid) - 1, len(riskGrid[0]) - 1}
	for {
		riskGrid[point.Row][point.Col].OnLowestRiskPath = true
		if point.Row == 0 && point.Col == 0 {
			break
		}
		point = riskGrid[point.Row][point.Col].Predecessor
	}

	return riskGrid
}

func (riskGrid RiskGrid) Print() {
	for r := range riskGrid {
		for c := range riskGrid[r] {
			if riskGrid[r][c].OnLowestRiskPath {
				fmt.Printf("%v ", riskGrid[r][c].Risk)
			} else {
				fmt.Printf("\033[35m%v \033[0m", riskGrid[r][c].Risk)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintGrids(riskGrid RiskGrid, grid Grid) {
	riskGrid.Print()

	for r := range grid {
		for c := range grid[r] {
			if riskGrid[r][c].OnLowestRiskPath {
				fmt.Printf("%v", grid[r][c])
			} else {
				fmt.Printf("\033[35m%v\033[0m", grid[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
