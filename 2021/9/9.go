package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var inputPath = "input.txt"

func main() {
	grid, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	lowPoints := grid.LowPoints()

	grid.Print(lowPoints)

	fmt.Printf("risk at low points: %d\n", RiskLevel(lowPoints))

	basins := grid.Basins(lowPoints)
	if err != nil {
		log.Fatal(err)
	}

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

type Grid [][]int

func Input(path string) (grid Grid, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for row, line := range lines {
		grid = append(grid, []int{})
		numbers := strings.Split(line, "")
		for _, n := range numbers {
			intN, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			grid[row] = append(grid[row], intN)
		}
	}

	return
}

func (grid Grid) Print(lowPoints LowPoint) {
	for r, row := range grid {
		for c, cell := range row {
			fmt.Print(cell)
			if _, ok := lowPoints[Point{r, c}]; ok {
				fmt.Print("'")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type LowPoint map[Point]int //map[x,y]value
type Point struct {
	r, c int
}

func (grid Grid) LowPoints() (lows LowPoint) {
	lows = make(LowPoint)

	for r, row := range grid {
		for c, cell := range row {
			if r > 0 {
				if cell >= grid[r-1][c] {
					continue
				}
			}
			if r < len(grid)-1 {
				if cell >= grid[r+1][c] {
					continue
				}
			}
			if c > 0 {
				if cell >= grid[r][c-1] {
					continue
				}
			}
			if c < len(row)-1 {
				if cell >= grid[r][c+1] {
					continue
				}
			}
			lows[Point{r, c}] = cell
		}
	}

	return
}

type Basins map[Point][]Point //map[low point]points in basin

func RiskLevel(points LowPoint) (risk int) {
	for _, p := range points {
		risk += 1 + p
	}
	return
}

func (grid Grid) Basins(lows LowPoint) (basins Basins) {
	basins = make(Basins)

	var flow func(current Point)

	partial := []Point{}
	testNext := func(next Point) {
		if !slices.Contains(partial, next) &&
			next.c >= 0 &&
			next.r >= 0 &&
			next.r < len(grid) &&
			next.c < len(grid[next.r]) &&
			grid[next.r][next.c] != 9 {
			//fmt.Printf("last: %v, current: %v, next: %v\n", last, current, next)
			flow(next)
		}
	}

	flow = func(current Point) {
		partial = append(partial, current)
		//fmt.Printf("p: %v\n", partial)

		next := Point{current.r, current.c + 1}
		testNext(next)

		next = Point{current.r, current.c - 1}
		testNext(next)

		next = Point{current.r + 1, current.c}
		testNext(next)

		next = Point{current.r - 1, current.c}
		testNext(next)

	}

	for low := range lows {
		flow(low)
		//fmt.Printf("final: %v\n", partial)
		basins[low] = append(basins[low], partial...)
		partial = []Point{}
	}

	return
}

/*
func (grid Grid) Basins3(lows LowPoint) (basins Basins) {
	var flow func(point Point)
	flow = func(point Point) {
		fmt.Println(point)
		for c := point.c + 1; c < len(grid[point.r]); c++ {
			if grid[point.r][c] == 9 {
				break
			}
			basins[point] = append(basins[point], Point{point.r, c})
			flow(Point{point.r, c})
		}
		for c := point.c - 1; c >= 0; c-- {
			if grid[point.r][c] == 9 {
				break
			}
			basins[point] = append(basins[point], Point{point.r, c})
			flow(Point{point.r, c})
		}
		for r := point.r + 1; r < len(grid); r++ {
			if grid[r][point.c] == 9 {
				break
			}
			basins[point] = append(basins[point], Point{r, point.c})
			flow(Point{r, point.c})
		}
		for r := point.r - 1; r >= 0; r-- {
			if grid[r][point.c] == 9 {
				break
			}
			basins[point] = append(basins[point], Point{r, point.c})
			flow(Point{r, point.c})
		}
	}

	basins = make(Basins)
	for low := range lows {
		basins[low] = []Point{low}
		flow(low)
	}

	return
}


func (grid Grid) Basins2(lows LowPoint) (basins Basins) {
	basins = make(Basins)
	for low := range lows {
		basins[low] = []Point{low}
	}

	for r, row := range grid {
		for c, cell := range row {
			if cell == 9 {
				continue
			}
			if _, ok := basins[Point{r, c}]; ok {
				continue
			}
			shortestD := len(grid) + len(grid[0])
			closestLow := Point{}
			for low := range basins {
				if shortestD > Distance(low, Point{r, c}) {
					if grid.FlowBlocked(low, Point{r, c}) {
						fmt.Printf("flow blocked: %v: %d, %v%d\n", low, grid[low.r][low.c], Point{r, c}, grid[r][c])
						continue
					}
					shortestD = Distance(low, Point{r, c})
					closestLow = low
				}
			}
			basins[closestLow] = append(basins[closestLow], Point{r, c})
		}
	}

	return
}

func Distance(one, two Point) int {
	return int(math.Abs(float64(one.r)-float64(two.r)) + math.Abs(float64(one.c)-float64(two.c)))
}

func (grid Grid) FlowBlocked(one, two Point) bool {
	horizontalBlocked := func(one, two Point) bool {
		iter := 1
		if one.c > two.c {
			iter = -1
		}
		for c := one.c; c != two.c; c += iter {
			if grid[one.r][c] == 9 {
				return true
			}
		}
		return false
	}

	verticalBlocked := func(one, two Point) bool {
		iter := 1
		if one.r > two.r {
			iter = -1
		}
		for r := one.r; r != two.r; r += iter {
			if grid[r][one.c] == 9 {
				return true
			}
		}
		return false
	}

	horizontalFirst := func(one, two Point) bool {
		if horizontalBlocked(one, two) {
			if verticalBlocked(one, two) {
				return true
			}
			one.r = two.r
			return horizontalBlocked(one, two)
		}
		one.c = two.c
		return horizontalBlocked(one, two)
	}

	verticalFirst := func(one, two Point) bool {
		if verticalBlocked(one, two) {
			if horizontalBlocked(one, two) {
				return true
			}
			one.c = two.c
			return verticalBlocked(one, two)
		}
		one.r = two.r
		return horizontalBlocked(one, two)
	}

	return horizontalFirst(one, two) || verticalFirst(one, two)
}
*/
