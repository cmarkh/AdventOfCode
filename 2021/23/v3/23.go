package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	bur := parse(input)
	bur.Print()

	fmt.Println("Part 1:")
	energy, _ := bur.shortestPath()
	fmt.Printf("used %v energy\n", energy)
	fmt.Println()

	fmt.Println("Part 2:")
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

		var found bool
		found, checked = contains(checked, move)
		if found {
			continue
		}

		fmt.Printf("len(moves): %v, len(checked): %v, least energy so far: %v, move energy: %v\n",
			len(moves), len(checked), energyUsed, move.energy)
		move.bur.Print()

		checked = append(checked, move)

		for r := range move.bur {
			for c := range move.bur[r] {
				if move.bur[r][c] == "" || move.bur[r][c] == "#" {
					continue
				}

				moves = append(moves, down(r, c, move)...)
				moves = append(moves, up(r, c, move)...)
				moves = append(moves, left(r, c, move)...)
				moves = append(moves, right(r, c, move)...)
			}
		}
	}

	return
}

func down(r, c int, move1 move) (newMoves []move) {
	if r+1 >= 2 {
		return
	}
	if move1.bur[r+1][c] != "" {
		return
	}
	if move1.bur[r][c] != sorted[c] { // Amphipods will never move from the hallway into a room unless that room is their destination room
		return
	}
	if move1.bur[1][c] != "" && move1.bur[1][c] != sorted[c] {
		return
	}
	if move1.bur[2][c] != "" && move1.bur[2][c] != sorted[c] {
		return
	}

	move2 := move1
	move2.previousStep = &move1

	if move1.bur[2][c] == "" {
		move2.bur[2][c] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]] * 2
	} else {
		move2.bur[1][c] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]]
	}

	return []move{move2}
}

func up(r, c int, move1 move) (newMoves []move) {
	if r-1 < 0 {
		return
	}
	if move1.bur[0][c] != "" { //if hall not empty, continue
		return
	}
	if move1.bur[r-1][c] != "" { //relevant if in row 2
		return
	}
	if r == 2 && move1.bur[r][c] == sorted[c] {
		return
	}
	if r == 1 && move1.bur[r][c] == sorted[c] && move1.bur[2][c] == sorted[c] {
		return
	}

	move2 := move1
	move2.previousStep = &move1
	//move2.bur[0][c] = move1.bur[r][c] //move into hall
	move2.bur[r][c] = ""
	move2.energy += energy[move1.bur[r][c]] * r //moves 2 places if coming from row 2

	for l := c - 1; l >= 0; l-- {
		if move2.bur[0][l] != "" {
			break
		}
		if slices.Contains(doors, l) { //no stopping in open doors
			continue
		}
		move3 := move2
		move3.bur[0][l] = move1.bur[r][c]
		move3.energy += energy[move1.bur[r][c]] * (c - l) //energy * number of moves to get there
		newMoves = append(newMoves, move3)
	}

	for l := c + 1; l < len(move2.bur[0]); l++ {
		if move2.bur[0][l] != "" {
			break
		}
		if slices.Contains(doors, l) { //no stopping in open doors
			continue
		}
		move3 := move2
		move3.bur[0][l] = move1.bur[r][c]
		move3.energy += energy[move1.bur[r][c]] * (l - c) //energy * number of moves to get there
		newMoves = append(newMoves, move3)
	}

	return
}

func left(r, c int, move1 move) (newMoves []move) {
	if r != 0 {
		return
	}

	//Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room
	for l := c - 1; l >= 0; l-- {
		if move1.bur[0][l] != "" {
			break
		}
		if sorted[l] != move1.bur[r][c] { //only bothering with moves I can move into
			continue
		}
		move2 := move1
		move2.previousStep = &move1
		move2.bur[0][l] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]] * (c - l) //energy * number of moves to get there

		newMoves = append(newMoves, down(0, l, move2)...) //let down handle whether we can move into the sorted column
	}

	return
}

func right(r, c int, move1 move) (newMoves []move) {
	if r != 0 {
		return
	}

	//Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room
	for l := c + 1; l < len(move1.bur[0]); l++ {
		if move1.bur[0][l] != "" {
			break
		}
		if sorted[l] != move1.bur[r][c] { //only bothering with moves I can move into
			continue
		}
		move2 := move1
		move2.previousStep = &move1
		move2.bur[0][l] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]] * (l - c) //energy * number of moves to get there

		newMoves = append(newMoves, down(0, l, move2)...) //let down handle whether we can move into the sorted column
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

func contains(checked []move, bur move) (bool, []move) {
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
			if check.energy == bur.energy || check.energy < bur.energy {
				return true, checked
			}
			checked = slices.Delete(checked, i, i+1) //will replace with the new lower energy version
		}
	}
	return false, checked
}

func trimChecked(checked []move, leastEnergy int) (new []move) {
	for _, checked := range checked {
		if checked.energy < leastEnergy {
			new = append(new, checked)
		}
	}
	return
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
