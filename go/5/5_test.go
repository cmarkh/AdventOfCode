package main

import (
	"fmt"
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = advent.BasePath + "5/test.txt"

func TestInput(t *testing.T) {
	vents, err := Input(testPath)
	if err != nil {
		t.Error(err)
	}

	vents.Print()
}

func TestGrid(t *testing.T) {
	vents, err := Input(testPath)
	if err != nil {
		t.Error(err)
	}

	grid, err := DrawVents(vents)
	if err != nil {
		t.Error(err)
	}
	grid.Print()
}

func TestDangerousPoints(t *testing.T) {
	vents, err := Input(testPath)
	if err != nil {
		t.Error(err)
	}

	grid, err := DrawVents(vents)
	if err != nil {
		t.Error(err)
	}

	points, count := grid.DangerousPoints()
	points.Print()
	fmt.Printf("Dangerous Points: %d\n", count)

}
