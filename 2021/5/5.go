package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
)

var inputPath = "input.txt"

func main() {
	vents, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	grid, err := DrawVents(vents)
	if err != nil {
		log.Fatal(err)
	}
	//grid.Print()

	_, count := grid.DangerousPoints()
	//points.Print()
	fmt.Printf("Dangerous Points: %d\n", count)
}

type Grid [][]int

type Vents []Vent
type Vent struct {
	Start Coordinates
	End   Coordinates
}

type Points []Coordinates
type Coordinates struct {
	X, Y int
}

func Input(path string) (vents Vents, err error) {
	input, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range input {
		vent := Vent{}

		start, end, found := strings.Cut(line, " -> ")
		if !found {
			err = fmt.Errorf("error parsing coordinates: %s", line)
			return
		}

		startX, startY, found := strings.Cut(start, ",")
		if !found {
			err = fmt.Errorf("error parsing coordinates: %s", line)
			return
		}
		endX, endY, found := strings.Cut(end, ",")
		if !found {
			err = fmt.Errorf("error parsing coordinates: %s", line)
			return
		}

		vent.Start.X, err = strconv.Atoi(startX)
		if err != nil {
			return
		}
		vent.Start.Y, err = strconv.Atoi(startY)
		if err != nil {
			return
		}
		vent.End.X, err = strconv.Atoi(endX)
		if err != nil {
			return
		}
		vent.End.Y, err = strconv.Atoi(endY)
		if err != nil {
			return
		}

		vents = append(vents, vent)
	}

	return
}

func (vent Vent) Equal() bool {
	return vent.Start.X == vent.End.X && vent.Start.Y == vent.End.Y
}
func (vent Vent) Vertical() bool {
	return vent.Start.X == vent.End.X
}
func (vent Vent) Horizontal() bool {
	return vent.Start.Y == vent.End.Y
}
func (vent Vent) Diagonal() bool {
	return math.Abs(float64(vent.End.X-vent.Start.X)) == math.Abs(float64(vent.End.Y-vent.Start.Y))
}

func (vent Vent) Print() {
	fmt.Printf("%d,%d -> %d,%d\n", vent.Start.X, vent.Start.Y, vent.End.X, vent.End.Y)
}
func (vents Vents) Print() {
	for _, vent := range vents {
		vent.Print()
	}
}

func DrawVents(vents Vents) (grid Grid, err error) {
	maxX, maxY := 0, 0
	for _, vent := range vents {
		if vent.Start.X > maxX {
			maxX = vent.Start.X
		}
		if vent.End.X > maxX {
			maxX = vent.End.X
		}
		if vent.Start.Y > maxY {
			maxY = vent.Start.Y
		}
		if vent.End.Y > maxY {
			maxY = vent.End.Y
		}
	}
	grid = make([][]int, maxY+1)
	for row := range grid {
		grid[row] = make([]int, maxX+1)
	}

	for _, vent := range vents {
		moveX, moveY := 1, 1
		if vent.Start.X > vent.End.X {
			moveX = -1
		}
		if vent.Start.Y > vent.End.Y {
			moveY = -1
		}

		grid[vent.End.Y][vent.End.X]++

		switch {
		case vent.Equal():
			//grid[vent.Start.Y][vent.Start.X]++

		case vent.Vertical():
			for i := vent.Start.Y; i != vent.End.Y; i += moveY {
				grid[i][vent.Start.X]++
			}

		case vent.Horizontal():
			for i := vent.Start.X; i != vent.End.X; i += moveX {
				grid[vent.Start.Y][i]++
			}

		case vent.Diagonal():
			for x, y := vent.Start.X, vent.Start.Y; x != vent.End.X; {
				grid[y][x]++

				x += moveX
				y += moveY
			}

		default:
			err = fmt.Errorf("unkown vent type: %v", vent)
		}

	}
	return
}

func (grid Grid) DangerousPoints() (points Points, count int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell >= 2 {
				points = append(points, Coordinates{x, y})
				count++
			}
		}
	}

	return
}

func (grid Grid) Print() {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}

func (points Points) Print() {
	for _, p := range points {
		fmt.Printf("%d,%d\n", p.X, p.Y)
	}
}
