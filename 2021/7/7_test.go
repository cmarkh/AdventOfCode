package main

import (
	"fmt"
	"testing"
)

var testPath = "test.txt"

func TestInput(t *testing.T) {
	positions, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(positions)
}

func TestAveragePosition(t *testing.T) {
	positions, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}
	avgPos := positions.AveragePosition()
	fmt.Printf("Avg Position: %d\n", avgPos)
}

func TestMinMax(t *testing.T) {
	positions, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}
	min, max := positions.MinMax()
	fmt.Printf("min: %d, max: %d\n", min, max)
}

func TestFuelUsed(t *testing.T) {
	positions, _ := Input(testPath)

	tgt := 2
	fuel := positions.FuelUsed(tgt)
	fmt.Printf("fuel used to get to %d: %d\n", tgt, fuel)

}

func TestOptimizeFuel(t *testing.T) {
	positions, _ := Input(testPath)

	newPos, fuelUsed := positions.OptimizeFuel()
	fmt.Printf("New Position: %d, Fuel Used: %d\n", newPos, fuelUsed)
}
