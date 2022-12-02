package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

func main() {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	steps := parse(lines)
	fmt.Println("Steps:")
	for _, step := range steps {
		fmt.Printf("%+v\n", step)
	}
	fmt.Println()

	fmt.Println("Part 1:")

	grid := initializeGrid(steps, 50)
	//grid.listOn()
	on := grid.countOn()
	fmt.Printf("%v cubes on\n", on)

	fmt.Println()

	fmt.Println("Part 2:")

	cubes := cubesWithIntersections(steps)
	on = countOn2(cubes)
	fmt.Printf("%v cubes on\n", on)

	fmt.Println()
}

type reactor map[int]map[int]map[int]bool

type step struct {
	start [3]int
	end   [3]int
	on    bool
}
type cube = step

func parse(lines []string) (steps []step) {
	var err error

	for _, line := range lines {
		step := step{}

		split := strings.Split(line, " ")
		if len(split) != 2 {
			log.Fatal("error parsing line: " + line)
		}
		switch split[0] {
		case "on":
			step.on = true
		case "off":
			step.on = false
		default:
			log.Fatal("error parsing line: " + line)
		}

		split = strings.Split(split[1], ",")
		if len(split) != 3 {
			log.Fatal("error parsing line: " + line)
		}
		for i, axis := range split {
			coords := strings.Split(axis, "..")
			step.start[i], err = strconv.Atoi(strings.Split(coords[0], "=")[1])
			if err != nil {
				log.Fatal("error parsing line: " + line)
			}
			step.end[i], err = strconv.Atoi(coords[1])
			if err != nil {
				log.Fatal("error parsing line: " + line)
			}
		}

		steps = append(steps, step)
	}
	return
}

func initializeGrid(procedure []step, region int) (grid reactor) {
	size := region*2 + 1 //at least -50 and at most 50
	grid = make(reactor, size)
	for x := -region; x <= region; x++ {
		grid[x] = make(map[int]map[int]bool, size)
		for y := -region; y <= region; y++ {
			grid[x][y] = make(map[int]bool, size)
		}
	}

	for _, step := range procedure {
		for x := range grid {
			if x < step.start[0] || x > step.end[0] {
				continue
			}
			for y := range grid[x] {
				if y < step.start[1] || y > step.end[1] {
					continue
				}
				for z := range grid[y] {
					if z < step.start[2] || z > step.end[2] {
						continue
					}
					grid[x][y][z] = step.on
				}
			}
		}
	}

	return
}

func (grid reactor) listOn() {
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] {
					fmt.Printf("%v,%v,%v\n", x, y, z)
				}
			}
		}
	}
}

func (grid reactor) countOn() (on int) {
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] {
					on++
				}
			}
		}
	}
	return
}

func (grid reactor) Print() {
	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] {
					fmt.Print(1)
				}
			}
		}
	}
}
