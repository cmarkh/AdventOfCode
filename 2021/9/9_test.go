package main

import (
	"fmt"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var testPath = "test.txt"

func TestInput(t *testing.T) {
	grid, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	grid.Print(nil)
}

func TestLowest(t *testing.T) {
	grid, _ := Input(testPath)

	lowest := grid.LowPoints()
	grid.Print(lowest)

	fmt.Printf("sum: %d\n", RiskLevel(lowest))
}

func TestBasins(t *testing.T) {
	grid, _ := Input(testPath)

	lows := grid.LowPoints()
	grid.Print(lows)

	basins := grid.Basins(lows)

	for low, b := range basins {
		fmt.Printf("%v: %v\n", low, b)
	}
}

func TestLargestBasins(t *testing.T) {
	grid, _ := Input(testPath)

	lows := grid.LowPoints()
	grid.Print(lows)

	basins := grid.Basins(lows)

	basinValues := maps.Values(basins)

	slices.SortFunc(basinValues, func(a, b []Point) bool {
		return len(a) > len(b)
	})

	for _, v := range basinValues {
		fmt.Println(v)
	}

	product := 1
	for i, basin := range basinValues {
		if i == 3 {
			break
		}
		fmt.Println(len(basin))
		product *= len(basin)
	}

	fmt.Printf("largest 3 product: %d\n", product)
}

/*
func TestFlowBlocked(t *testing.T) {
	grid, _ := Input(testPath)

	fmt.Println(grid[4][6])
	fmt.Println(grid.FlowBlocked(Point{3, 4}, Point{4, 6}))
}
*/
