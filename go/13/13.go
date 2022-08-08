package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "13/input.txt"

func main() {
	coordinates, folds, err := input()
	if err != nil {
		log.Fatal(err)
	}

	for _, coord := range coordinates {
		fmt.Printf("%+v\n", coord)
	}
	for _, fold := range folds {
		fmt.Printf("fold: %+v\n", fold)
	}
	fmt.Println()

	grid := NewGrid(coordinates)
	grid.Print()

	fmt.Println()

	for _, fold := range folds {
		grid = grid.Fold(fold)
		//break //part one
	}
	grid.Print()
	fmt.Printf("grid has %d dots\n", grid.CountDots())
	fmt.Println()
}

type Coordinates []Coordinate
type Coordinate struct {
	x, y int
}

type Fold struct {
	IsX      bool
	Position int
}

func input() (coords Coordinates, folds []Fold, err error) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(content), "\n")

	foldStart := 0
	for i, line := range input {
		if line == "" {
			foldStart = i + 1
			break
		}

		split := strings.Split(line, ",")
		if len(split) != 2 {
			err = fmt.Errorf("unkown coordinate: %s", line)
			return
		}

		var x, y int
		x, err = strconv.Atoi(split[0])
		if err != nil {
			return
		}
		y, err = strconv.Atoi(split[1])
		if err != nil {
			return
		}
		coords = append(coords, Coordinate{x, y})
	}

	for line := foldStart; line < len(input); line++ {
		var fold Fold
		strFold := strings.Split(input[line], "=")
		if len(strFold) != 2 {
			err = fmt.Errorf("unkown fold: %s at %d", input[line], line)
			return
		}
		fold.Position, err = strconv.Atoi(strFold[1])
		if err != nil {
			return
		}
		if string(strFold[0][len(strFold[0])-1]) == "x" {
			fold.IsX = true
		}
		folds = append(folds, fold)
	}

	return
}

type Grid [][]string

func NewGrid(coords Coordinates) (grid Grid) {
	maxX, maxY := coords.Maxes()

	basicLine := []string{}
	for x := 0; x <= maxX; x++ {
		basicLine = append(basicLine, ".")
	}
	for y := 0; y <= maxY; y++ {
		grid = append(grid, append([]string{}, basicLine...))
	}

	for _, coord := range coords {
		grid[coord.y][coord.x] = "#"
	}

	return
}

func (coords Coordinates) Maxes() (maxX, maxY int) {
	for _, coord := range coords {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}
	return
}

func (grid Grid) Print() {
	for _, line := range grid {
		for _, cell := range line {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func (grid Grid) CountDots() (count int) {
	for _, line := range grid {
		for _, cell := range line {
			if cell == "#" {
				count++
			}
		}
	}
	return
}

func (grid Grid) Fold(fold Fold) Grid {
	for r, row := range grid {
		for c, col := range row {
			if col != "#" {
				continue
			}
			if fold.IsX {
				if c > fold.Position {
					grid[r][fold.Position-(c-fold.Position)] = "#"
				}
			} else {
				if r > fold.Position {
					grid[fold.Position-(r-fold.Position)][c] = "#"
				}
			}
		}
		if fold.IsX {
			grid[r] = grid[r][:fold.Position]
		}
	}

	if !fold.IsX {
		grid = grid[:fold.Position]
	}

	return grid
}
