package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {

}

type burrow [3][11]string

var energy = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

var sorted = map[int]string{ //map[col]character if sorted
	2: "A",
	4: "B",
	6: "C",
	8: "D",
}

func parse(input string) (bur burrow) {
	lines := strings.Split(input, "\n")

	bur[1][0] = "#"
	bur[2][0] = "#"

	bur[1][1] = "#"
	bur[2][1] = "#"

	bur[1][2] = string(lines[2][3])
	bur[2][2] = string(lines[3][3])

	bur[1][3] = "#"
	bur[2][3] = "#"

	bur[1][4] = string(lines[2][5])
	bur[2][4] = string(lines[3][5])

	bur[1][5] = "#"
	bur[2][5] = "#"

	bur[1][6] = string(lines[2][7])
	bur[2][6] = string(lines[3][7])

	bur[1][7] = "#"
	bur[2][7] = "#"

	bur[1][8] = string(lines[2][9])
	bur[2][8] = string(lines[3][9])

	bur[1][9] = "#"
	bur[2][9] = "#"

	bur[1][10] = "#"
	bur[2][10] = "#"

	return
}

func (bur burrow) Print() {
	dot := func(ch string) string {
		if ch == "" {
			return "."
		}
		return ch
	}

	for _, row := range bur {
		for _, col := range row {
			fmt.Print(dot(col))
		}
		fmt.Println()
	}

	fmt.Println()
}

type move struct {
	bur          burrow
	energy       int
	previousStep *move
}

var doors = []int{2, 4, 6, 8}

func (bur burrow) shortestPath() (energyUsed int, final move) {
	var moves = []move{{bur, 0, nil}}
	var checked = []move{}
	energyUsed = math.MaxInt / 2

	for len(moves) > 0 {
		move := moves[len(moves)-1]
		moves = moves[:len(moves)-1]

		if move.energy > energyUsed {
			continue
		}
		if move.bur.done() {
			if move.energy < energyUsed {
				energyUsed = move.energy
				final = move
				checked = trimChecked(checked, energyUsed)
			}
			continue
		}
		/*if move.bur.sumHall() > 4 {
			continue
		}*/
		if contains(checked, move) {
			continue
		}

		fmt.Printf("len(moves): %v, checked: %v, least energy so far: %v, move energy: %v\n",
			len(moves), len(checked), energyUsed, move.energy)
		move.bur.Print()

		checked = append(checked, move)

		for r := range move.bur {
			for c := range move.bur[r] {
				if move.bur[r][c] == "" || move.bur[r][c] == "#" {
					continue
				}
				if r == 2 && move.bur[r][c] == sorted[c] {
					continue
				}
				if r == 1 && move.bur[r][c] == sorted[c] && move.bur[r+1][c] == sorted[c] {
					continue
				}

				if move2, new := down(r, c, move); new {
					moves = append(moves, move2)
				}

				if move2, new := up(r, c, move); new {
					moves = append(moves, move2)

					/*
						if move2, new := left(r-1, c, move2); new {
							moves = append(moves, move2)
							fmt.Println("here")
						}
						if move2, new := right(r-1, c, move2); new {
							moves = append(moves, move2)
						}*/
				}

				if move2, new := left(r, c, move); new {
					moves = append(moves, move2)
				}

				if move2, new := right(r, c, move); new {
					moves = append(moves, move2)
				}

			}
		}

		//sort so prioritize positions with highest number of rooms filled correctly
		sort.SliceStable(moves, func(i, j int) bool {
			both1, bottom1, top1 := countSorted(moves[i].bur)
			both2, bottom2, top2 := countSorted(moves[j].bur)
			if both1 != both2 {
				return both1 < both2
			}
			if bottom1 != bottom2 {
				return bottom1 < bottom2
			}
			if top1 != top2 {
				return top1 < top2
			}
			return moves[i].energy > moves[j].energy
		})
	}

	return
}

func down(r, c int, move move) (move, bool) {
	//move down
	if r+1 < len(move.bur) && move.bur[r+1][c] == "" {
		if sorted[c] != move.bur[r][c] { // Amphipods will never move from the hallway into a room unless that room is their destination room
			return move, false
		}
		move2 := move
		move2.previousStep = &move
		move2.bur[r+1][c] = move.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move2.bur[r+1][c]]
		return move2, true
	}
	return move, false
}

func up(r, c int, move move) (move, bool) {
	//move up
	if r-1 >= 0 && move.bur[r-1][c] == "" {
		move2 := move
		move2.previousStep = &move
		move2.bur[r-1][c] = move.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move2.bur[r-1][c]]
		return move2, true
	}
	return move, false
}

func left(r, c int, move move) (move, bool) {
	//move left
	if c-1 >= 0 && move.bur[r][c-1] == "" {
		left := c - 1
		doubleEnergy := 1
		if slices.Contains(doors, left) && sorted[left] != move.bur[r][c] {
			if left-1 < 0 {
				return move, false
			}
			left--
			doubleEnergy = 2 //since took 2 steps
		}
		if move.bur[r][left] != "" {
			return move, false
		}
		move2 := move
		move2.previousStep = &move
		move2.bur[r][left] = move.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move2.bur[r][left]] * doubleEnergy
		return move2, true
	}
	return move, false
}

func right(r, c int, move move) (move, bool) {
	//move right
	if c+1 < len(move.bur[0]) && move.bur[r][c+1] == "" {
		right := c + 1
		doubleEnergy := 1
		if slices.Contains(doors, right) && sorted[right] != move.bur[r][c] {
			if right+1 >= len(move.bur[0]) {
				return move, false
			}
			right++
			doubleEnergy = 2 //since took 2 steps
		}
		if move.bur[r][right] != "" {
			return move, false
		}
		move2 := move
		move2.previousStep = &move
		move2.bur[r][right] = move.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move2.bur[r][right]] * doubleEnergy
		return move2, true
	}
	return move, false
}

func countSorted(bur burrow) (both, bottom, top int) {
	for c, ch := range sorted {
		if bur[1][c] == ch {
			top++
		}
		if bur[2][c] == ch {
			bottom++
		}
		if bur[1][c] == ch && bur[2][c] == ch {
			both++
		}
	}
	return
}

func (bur burrow) done() bool {
	if bur[1][2] != "A" || bur[2][2] != "A" {
		return false
	}
	if bur[1][4] != "B" || bur[2][4] != "B" {
		return false
	}
	if bur[1][6] != "C" || bur[2][6] != "C" {
		return false
	}
	if bur[1][8] != "D" || bur[2][8] != "D" {
		return false
	}
	return true
}

func contains(checked []move, bur move) bool {
	for i, check := range checked {
		if func() bool {
			for r := range bur.bur {
				for c := range bur.bur[r] {
					if bur.bur[r][c] != check.bur[r][c] {
						return false
					}
				}
			}
			return true
		}() {
			if check.energy > bur.energy {
				checked[i].energy = bur.energy //if position exists but now has a shorter energy to get there
			}
			return true
		}
	}
	return false
}

func trimChecked(checked []move, leastEnergy int) (new []move) {
	for _, checked := range checked {
		if checked.energy < leastEnergy {
			new = append(new, checked)
		}
	}
	return
}

func (bur burrow) sumHall() (sum int) {
	for _, ch := range bur[0] {
		if ch != "" {
			sum++
		}
	}
	return
}
