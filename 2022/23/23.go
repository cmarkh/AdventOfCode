package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:")
	elves := parse(input)
	elves, _ = elves.move(10)
	empty := elves.countEmpty()
	fmt.Printf("empty ground tiles: %v\n", empty)
	fmt.Println()

	fmt.Println("Part 2:")
	elves = parse(input)
	_, rounds := elves.move(0)
	fmt.Printf("done after %v rounds\n", rounds)
	fmt.Println()
}

type position struct {
	r, c int
}
type elf struct {
	proposed position
}
type elves map[position]elf

func parse(input string) (elves elves) {
	elves = make(map[position]elf)
	lines := strings.Split(input, "\n")

	for r, line := range lines {
		if line == "" {
			continue
		}
		for c, ch := range line {
			if ch == '#' {
				elves[position{r, c}] = elf{}
			}
		}
	}

	return
}

func (elves elves) gridSize() (minR, minC, maxR, maxC int) {
	for pos := range elves {
		if pos.r < minR {
			minR = pos.r
		}
		if pos.r > maxR {
			maxR = pos.r
		}
		if pos.c < minC {
			minC = pos.c
		}
		if pos.c > maxC {
			maxC = pos.c
		}
	}
	return
}

func (elves elves) print() {
	minR, minC, maxR, maxC := elves.gridSize()
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if _, ok := elves[position{r, c}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (elves elves) move(maxRounds int) (newElves elves, round int) {
	north := func(n, s, e, w, ne, nw, se, sw bool, pos position) (position, bool) {
		if !n && !ne && !nw {
			return position{pos.r - 1, pos.c}, true
		}
		return position{}, false
	}
	south := func(n, s, e, w, ne, nw, se, sw bool, pos position) (position, bool) {
		if !s && !se && !sw {
			return position{pos.r + 1, pos.c}, true
		}
		return position{}, false
	}
	west := func(n, s, e, w, ne, nw, se, sw bool, pos position) (position, bool) {
		if !w && !nw && !sw {
			return position{pos.r, pos.c - 1}, true
		}
		return position{}, false
	}
	east := func(n, s, e, w, ne, nw, se, sw bool, pos position) (position, bool) {
		if !e && !ne && !se {
			return position{pos.r, pos.c + 1}, true
		}
		return position{}, false
	}
	directions := []func(n, s, e, w, ne, nw, se, sw bool, pos position) (position, bool){
		north, south, west, east,
	}

	noMoves := 0
	for round = 0; noMoves < len(elves); round++ {
		if maxRounds > 0 && round == maxRounds {
			return
		}

		proposals := make(map[position]int) //count of position proposed x times
		noMoves = 0

		for pos, elf := range elves {
			_, n := elves[position{pos.r - 1, pos.c}]
			_, s := elves[position{pos.r + 1, pos.c}]
			_, e := elves[position{pos.r, pos.c + 1}]
			_, w := elves[position{pos.r, pos.c - 1}]

			_, ne := elves[position{pos.r - 1, pos.c + 1}]
			_, nw := elves[position{pos.r - 1, pos.c - 1}]
			_, se := elves[position{pos.r + 1, pos.c + 1}]
			_, sw := elves[position{pos.r + 1, pos.c - 1}]

			func() {
				//no movement
				if !n && !s && !e && !w && !ne && !nw && !se && !sw {
					noMoves++
					elf.proposed = pos
					elves[pos] = elf
					proposals[elf.proposed] += 1
					return
				}

				for _, direction := range directions {
					if proposal, valid := direction(n, s, e, w, ne, nw, se, sw, pos); valid {
						elf.proposed = proposal
						elves[pos] = elf
						proposals[elf.proposed] += 1
						return
					}
				}

				//can't move
				elf.proposed = pos
				elves[pos] = elf
				proposals[elf.proposed] += 1
			}()
		}

		//at the end of the round, the first direction the Elves considered is moved to the end of the list of directions
		directions = append(directions[1:], directions[0])

		newElves = make(map[position]elf)
		for pos, elf := range elves {
			if proposals[elf.proposed] > 1 {
				newElves[pos] = elf //no movement
				continue
			}
			newElves[elf.proposed] = elf
		}

		elves = newElves
		//fmt.Println(round + 1)
		//elves.print()
	}

	return
}

func (elves elves) countEmpty() (empty int) {
	minR, minC, maxR, maxC := elves.gridSize()
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if _, ok := elves[position{r, c}]; !ok {
				empty++
			}
		}
	}
	return
}
