package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	sandIngress = position{500, 0}
)

func main() {
	fmt.Println("Part 1:")
	grid := parse(input)
	sand := grid.produceSand()
	fmt.Printf("%v units of sand stacked up\n", sand)
	fmt.Println()

	fmt.Println("Part 2:")
	grid = parse(input)
	sand = grid.produceSand2()
	fmt.Printf("%v units of sand stacked up\n", sand)
	fmt.Println()
}

type grid map[position]string
type position struct {
	x, y int
}

func parse(input string) (g grid) {
	g = make(grid)
	g[sandIngress] = "+"

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " -> ")
		points := []position{}
		for _, point := range split {
			xy := strings.Split(point, ",")
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				log.Fatal(err)
			}
			points = append(points, position{x, y})
		}

		for i, point := range points {
			if i == 0 {
				continue
			}

			switch {
			case point.x == points[i-1].x:
				if point.y > points[i-1].y {
					for y := points[i-1].y; y <= point.y; y++ {
						g[position{point.x, y}] = "#"
					}
				} else {
					for y := point.y; y <= points[i-1].y; y++ {
						g[position{point.x, y}] = "#"
					}
				}

			case point.y == points[i-1].y:
				if point.x > points[i-1].x {
					for x := points[i-1].x; x <= point.x; x++ {
						g[position{x, point.y}] = "#"
					}
				} else {
					for x := point.x; x <= points[i-1].x; x++ {
						g[position{x, point.y}] = "#"
					}
				}

			default:
				log.Fatalf("not a straight line: %+v", points)
			}
		}
	}

	return
}

func (g grid) print() {
	maxX, maxY := 0, 0
	minX, minY := math.MaxInt/2, math.MaxInt/2
	for pos := range g {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.x < minX {
			minX = pos.x
		}

		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.y < minY {
			minY = pos.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, ok := g[position{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) produceSand() (stackedSand int) {
	maxY := 0
	for pos := range g {
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	for {
		sand := sandIngress
		for {
			if sand.y >= maxY {
				return
			}
			if _, ok := g[position{sand.x, sand.y + 1}]; !ok { //if down one position is empty
				sand.y++
				continue
			}
			if _, ok := g[position{sand.x - 1, sand.y + 1}]; !ok { //if down and left one position is empty
				sand.x--
				sand.y++
				continue
			}
			if _, ok := g[position{sand.x + 1, sand.y + 1}]; !ok { //if down and right one position is empty
				sand.x++
				sand.y++
				continue
			}
			//if reach here, sand couldn't move further
			g[sand] = "o"
			stackedSand++
			break
		}
	}
}

func (g grid) produceSand2() (stackedSand int) {
	maxY := 0
	for pos := range g {
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	maxY += 2 //assume the floor is an infinite horizontal line with a y coordinate equal to two plus the highest y coordinate of any point in your scan

	for {
		sand := sandIngress
		for {
			if sand.y == maxY-1 {
				g[sand] = "o"
				stackedSand++
				break
			}
			if _, ok := g[position{sand.x, sand.y + 1}]; !ok { //if down one position is empty
				sand.y++
				continue
			}
			if _, ok := g[position{sand.x - 1, sand.y + 1}]; !ok { //if down and left one position is empty
				sand.x--
				sand.y++
				continue
			}
			if _, ok := g[position{sand.x + 1, sand.y + 1}]; !ok { //if down and right one position is empty
				sand.x++
				sand.y++
				continue
			}
			//if reach here, sand couldn't move further
			g[sand] = "o"
			stackedSand++
			break
		}
		//simulate falling sand until a unit of sand comes to rest at 500,0, blocking the source entirely and stopping the flow of sand into the cave
		if sand == sandIngress {
			return
		}
	}
}
