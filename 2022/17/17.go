package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	jets := parse(input)

	fmt.Println("Part 1:")
	_, height := fall(jets, 2022)
	fmt.Printf("height: %v\n", height)
	fmt.Println()

	fmt.Println("Part 2:")
	height = fall2(jets, 1000000000000)
	fmt.Printf("height: %v\n", height)
	fmt.Println()
}

func parse(input string) (jets jets) {
	jets = newJets()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		jets.jets = append(jets.jets, strings.Split(line, "")...)
	}
	return
}

type grid [][7]string //The tall, vertical chamber is exactly seven units wide

func (g grid) print() {
	for _, row := range g {
		for _, col := range row {
			if col == "" {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) rockHitSomething(rock rock, rockBottom, rockLeft int) bool {
	for r := 0; r < len(rock); r++ {
		for c := 0; c < len(rock[r]); c++ {
			if rock[len(rock)-r-1][c] == "" {
				continue
			}
			if g[rockBottom-r][c+rockLeft] != "" {
				//fmt.Printf("r: %v, c: %v\n", rockBottom-r, c+rockLeft)
				return true
			}
		}
	}
	return false
}

type jets struct {
	i    int
	jets []string
}

func newJets() jets {
	return jets{-1, []string{}}
}

func (jets *jets) next() string {
	jets.i++
	if jets.i == len(jets.jets) {
		jets.i = 0
	}
	return jets.jets[jets.i]
}

type rock [][]string

var rockShapes = []rock{
	{{"#", "#", "#", "#"}},

	{{"", "#", ""},
		{"#", "#", "#"},
		{"", "#", ""}},

	{{"", "", "#"},
		{"", "", "#"},
		{"#", "#", "#"}},

	{{"#"},
		{"#"},
		{"#"},
		{"#"}},

	{{"#", "#"},
		{"#", "#"}},
}

type rocks struct {
	i     int
	rocks []rock
}

func newRocks() rocks {
	return rocks{-1, rockShapes}
}

func (rock rock) width() (width int) {
	for _, row := range rock {
		if len(row) > width {
			width = len(row)
		}
	}
	return
}

func (rocks *rocks) next() rock {
	rocks.i++
	if rocks.i == len(rocks.rocks) {
		rocks.i = 0
	}
	return rocks.rocks[rocks.i]
}

func fall(jets jets, rockCount int) (g grid, height int) {
	rockI := -1
	jetI := -1
	for i := 0; i < rockCount; i++ {
		//Each rock appears so that its left edge is two units away from the left wall and
		//its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one)

		rockI++
		if rockI >= 5 {
			rockI = 0
		}
		rock := rockShapes[rockI]

		g = append(make([][7]string, len(rock)+3), g...)
		rockLeft := 2
		rockBottom := len(rock) - 1

		func() {
			for {
				jetI++
				if jetI == len(jets.jets) {
					jetI = 0
				}

				//fmt.Println(jets[jetI])
				switch jets.jets[jetI] {
				case "<":
					if rockLeft > 0 && !g.rockHitSomething(rock, rockBottom, rockLeft-1) {
						rockLeft--
					}
				case ">":
					if rockLeft+rock.width()-1 < 6 && !g.rockHitSomething(rock, rockBottom, rockLeft+1) {
						rockLeft++
					}
				default:
					log.Fatalf("jet not understood: %v", jets.jets[jetI])
				}

				rockBottom++
				if rockBottom == len(g) || g.rockHitSomething(rock, rockBottom, rockLeft) {
					rockBottom--
					return
				}
			}
		}()

		for r := len(rock) - 1; r >= 0; r-- {
			for c, col := range rock[len(rock)-r-1] {
				if col == "" {
					continue
				}
				g[rockBottom-r][c+rockLeft] = col
			}
		}

		//remove extra lines at top of grid
		func() {
			for r := 0; ; {
				for _, col := range g[r] {
					if col != "" {
						return
					}
				}
				g = g[r+1:]
			}
		}()
	}

	return g, len(g)
}

func fall2(jets jets, rockCount int) (height int) {
	type pattern struct {
		rock, jet int
		grid20    [20][7]string
	}
	patterns := make(map[pattern]struct{})

	g := grid{}
	rocks := newRocks()

	dropRock := func() {
		rock := rocks.next()
		g = append(make([][7]string, len(rock)+3), g...)
		rockLeft := 2
		rockBottom := len(rock) - 1

		func() {
			for {
				//fmt.Println(jets[jetI])
				switch jets.next() {
				case "<":
					if rockLeft > 0 && !g.rockHitSomething(rock, rockBottom, rockLeft-1) {
						rockLeft--
					}
				case ">":
					if rockLeft+rock.width()-1 < 6 && !g.rockHitSomething(rock, rockBottom, rockLeft+1) {
						rockLeft++
					}
				default:
					log.Fatalf("jet not understood: %v", jets.jets[jets.i])
				}

				rockBottom++
				if rockBottom == len(g) || g.rockHitSomething(rock, rockBottom, rockLeft) {
					rockBottom--
					return
				}
			}
		}()

		//add rock to grid
		for r := len(rock) - 1; r >= 0; r-- {
			for c, col := range rock[len(rock)-r-1] {
				if col == "" {
					continue
				}
				g[rockBottom-r][c+rockLeft] = col
			}
		}

		//remove extra lines at top of grid
		func() {
			for r := 0; ; {
				for _, col := range g[r] {
					if col != "" {
						return
					}
				}
				g = g[r+1:]
			}
		}()
	}

	//find first instance of repition
	var i int
	var key pattern
	for i = 0; i < rockCount; i++ {
		dropRock()
		if len(g) > 20 {
			var cutset [20][7]string
			//lint:ignore S1001 using array not slice
			for r, row := range g[:20] {
				cutset[r] = row
			}
			key = pattern{rocks.i, jets.i, cutset}
			if _, ok := patterns[key]; ok {
				break
			} else {
				patterns[key] = struct{}{}
			}
		}
	}
	i++

	//calculate height of repeating pattern
	h1 := len(g)
	length := 0
	for length = i; length < rockCount; length++ {
		dropRock()
		var cutset [20][7]string
		//lint:ignore S1001 using array not slice
		for r, row := range g[:20] {
			cutset[r] = row
		}
		if (key == pattern{rocks.i, jets.i, cutset}) {
			break
		}
	}
	length++
	h2 := len(g) - h1

	repeats := (rockCount - i) / (length - i)
	remaining := (rockCount - i) % (length - i)

	height = h2*repeats + h1

	/*fmt.Printf("i: %v, length: %v\n", i, length)
	fmt.Printf("h1: %v, h2: %v\n", h1, h2)
	fmt.Printf("repeats: %v\n", repeats)
	fmt.Printf("remaining i: %v\n", remaining)*/

	h1 = len(g)
	for i := 0; i < remaining; i++ {
		dropRock()
	}
	height += len(g) - h1

	return
}
