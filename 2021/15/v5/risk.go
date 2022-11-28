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
			riskGrid[r][c].Risk = math.MaxInt / 2 //start all risk at ~infinite and will override with lowest risk available
		}
	}
	riskGrid[0][0].Risk = 0 //starting point

	riskGrid = RecalculateRisk(riskGrid, grid)

	riskGrid = riskGrid.MarkPath()

	return riskGrid, riskGrid[len(grid)-1][rowLength-1].Risk
}

func RecalculateRisk(riskGrid RiskGrid, grid Grid) RiskGrid {
	repeat := true
	for repeat {
		repeat = false
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				if c > 0 {
					if riskGrid[r][c-1].Risk+grid[r][c] < riskGrid[r][c].Risk {
						riskGrid[r][c].Risk = riskGrid[r][c-1].Risk + grid[r][c]
						riskGrid[r][c].Predecessor = Point{r, c - 1}
						repeat = true
					}
				}
				if c < len(grid[r])-1 {
					if riskGrid[r][c+1].Risk+grid[r][c] < riskGrid[r][c].Risk {
						riskGrid[r][c].Risk = riskGrid[r][c+1].Risk + grid[r][c]
						riskGrid[r][c].Predecessor = Point{r, c + 1}
						repeat = true
					}
				}
				if r > 0 {
					if riskGrid[r-1][c].Risk+grid[r][c] < riskGrid[r][c].Risk {
						riskGrid[r][c].Risk = riskGrid[r-1][c].Risk + grid[r][c]
						riskGrid[r][c].Predecessor = Point{r - 1, c}
						repeat = true
					}
				}
				if r < len(grid)-1 {
					if riskGrid[r+1][c].Risk+grid[r][c] < riskGrid[r][c].Risk {
						riskGrid[r][c].Risk = riskGrid[r+1][c].Risk + grid[r][c]
						riskGrid[r][c].Predecessor = Point{r + 1, c}
						repeat = true
					}
				}
			}
		}
	}

	return riskGrid
}

func (riskGrid RiskGrid) MarkPath() RiskGrid {
	point := Point{len(riskGrid) - 1, len(riskGrid[0]) - 1}
	for {
		//fmt.Printf("%+v\t", point)
		riskGrid[point.Row][point.Col].OnLowestRiskPath = true
		if point.Row == 0 && point.Col == 0 {
			break
		}
		point = riskGrid[point.Row][point.Col].Predecessor
		//fmt.Printf("%+v\n", point)
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

func (grid Grid) PrintPath(riskGrid RiskGrid) {
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
