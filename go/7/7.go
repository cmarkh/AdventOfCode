package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "7/input.txt"

func main() {
	positions, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	newPos, fuelUsed := positions.OptimizeFuel()
	fmt.Printf("New Position: %d, Fuel Used: %d\n", newPos, fuelUsed)
}

type Positions []int

func Input(path string) (positions Positions, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		strPositions := strings.Split(line, ",")
		for _, pos := range strPositions {
			pos, err := strconv.Atoi(pos)
			if err != nil {
				return positions, err
			}
			positions = append(positions, pos)
		}
	}

	return
}

func (positions Positions) AveragePosition() int {
	sum := 0
	for _, p := range positions {
		sum += p
	}
	return sum / len(positions)
}

func (positions Positions) OptimizeFuel() (newPos, fuelUsed int) {
	min, max := positions.MinMax()

	for i := min; i <= max; i++ {
		fuel := positions.FuelUsed(i)
		if fuel < fuelUsed || fuelUsed == 0 {
			fuelUsed = fuel
			newPos = i
		}
	}

	return
}

func (positions Positions) MinMax() (min, max int) {
	for _, p := range positions {
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
	}
	return
}

func (positions Positions) FuelUsed(tgtPosition int) (fuel int) {
	for _, p := range positions {
		for i := 1; i <= int(math.Abs(float64(p-tgtPosition))); i++ {
			fuel += i
		}
		//fuel += int(math.Abs(float64(p - tgtPosition)))
	}
	//fmt.Printf("pos: %d, fuel: %d\n", tgtPosition, fuel)
	return
}
