package main

import (
	"fmt"
	"testing"
)

var testPath = "test.txt"

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
