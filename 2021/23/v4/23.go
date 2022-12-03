package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

var part2input = `  #D#C#B#A#
  #D#B#A#C#`

func main() {
	t := time.Now()

	bur := parse(input, "")
	bur.Print()

	fmt.Println("Part 1:")
	energy1, _ := bur.shortestPath()
	fmt.Printf("used %v energy\n", energy1)
	fmt.Println()

	fmt.Println("Part 2:")
	bur = parse(input, part2input)
	bur.Print()

	energy2, _ := bur.shortestPath()
	fmt.Printf("used %v energy\n", energy2)
	fmt.Println()

	fmt.Println("(and part 1 again):")
	//fmt.Printf("used %v energy\n", energy1)

	fmt.Printf("You took %v to finish. I fell asleep waiting...\n", time.Since(t))
}

type burrow [7][13]string

var energy = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

var sorted = map[int]string{ //map[col]character if sorted
	3: "A",
	5: "B",
	7: "C",
	9: "D",
}

var doors = []int{3, 5, 7, 9}

func (bur burrow) shortestPath() (energyUsed int, final move) {
	var moves = []move{{bur, 0, nil}}
	var checked = []move{}
	energyUsed = math.MaxInt / 2

	for len(moves) > 0 {
		move := moves[len(moves)-1].copy()
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

		sortMoves(moves)
	}

	return
}

func sortMoves(moves []move) {
	s := func(i, j int) bool {
		//sum each row's position with a fully sorted char
		var sumi, sumj int
		for _, door := range doors {
			done1, done2 := false, false
			for r := 2; r < len(moves[i].bur); r++ {
				if done1 && done2 {
					break
				}
				ch1, ch2 := moves[i].bur[r][door], moves[j].bur[r][door]
				if ch1 == "#" || ch2 == "#" {
					break
				}
				if ch1 == sorted[door] {
					sumi++
				} else {
					done1 = true
				}
				if ch2 == sorted[door] {
					sumj++
				} else {
					done2 = true
				}
			}
		}
		if sumi == sumj {
			return moves[i].energy < moves[j].energy
		}
		return sumi < sumj
	}
	sort.SliceStable(moves, s)
}

func down(r, c int, move1 move) (newMoves []move) {
	if r+1 >= len(move1.bur) {
		return
	}
	if move1.bur[r+1][c] != "" {
		return
	}
	if move1.bur[r][c] != sorted[c] { // Amphipods will never move from the hallway into a room unless that room is their destination room
		return
	}

	//make sure all pods in the room are sorted or can't move into room
	for i := r + 1; i < len(move1.bur); i++ {
		if move1.bur[i][c] == "#" {
			break
		}
		if move1.bur[i][c] != "" && move1.bur[i][c] != sorted[c] {
			return
		}
	}

	move2 := move1.copy()
	move2.previousStep = &move1
	move2.bur[r][c] = ""

	lowest := r + 1
	for ; lowest < len(move1.bur); lowest++ {
		if move2.bur[lowest][c] == "#" {
			break
		}
		if move2.bur[lowest][c] == "" {
			move2.energy += energy[move1.bur[r][c]]
		} else {
			break
		}
	}
	move2.bur[lowest-1][c] = move1.bur[r][c]

	return []move{move2}
}

func up(r, c int, move1 move) (newMoves []move) {
	if r-1 < 1 {
		return
	}

	//check all slots up to the hall are empty (so can move up)
	for i := r - 1; i >= 1; i-- {
		if move1.bur[i][c] != "" {
			return
		}
	}

	if move1.bur[r][c] == sorted[c] { //if already sorted, check that those below aren't sorted or already done here
		allSorted := true
		for i := r + 1; i < len(move1.bur); i++ {
			if move1.bur[i][c] == "#" {
				break
			}
			if move1.bur[i][c] != sorted[c] {
				allSorted = false
				break
			}
		}
		if allSorted {
			return
		}
	}

	move2 := move1.copy()
	move2.previousStep = &move1
	//move2.bur[0][c] = move1.bur[r][c] //move into hall
	move2.bur[r][c] = ""
	move2.energy += energy[move1.bur[r][c]] * (r - 1) //moves 2 places if coming from row 3

	for l := c - 1; l > 0; l-- {
		if move2.bur[1][l] != "" {
			break
		}
		if slices.Contains(doors, l) { //no stopping in open doors
			continue
		}
		move3 := move2.copy()
		move3.bur[1][l] = move1.bur[r][c]
		move3.energy += energy[move1.bur[r][c]] * (c - l) //energy * number of moves to get there
		newMoves = append(newMoves, move3)
	}

	for l := c + 1; l < len(move2.bur[0]); l++ {
		if move2.bur[1][l] != "" {
			break
		}
		if slices.Contains(doors, l) { //no stopping in open doors
			continue
		}
		move3 := move2.copy()
		move3.bur[1][l] = move1.bur[r][c]
		move3.energy += energy[move1.bur[r][c]] * (l - c) //energy * number of moves to get there
		newMoves = append(newMoves, move3)
	}

	return
}

func left(r, c int, move1 move) (newMoves []move) {
	if r != 1 { //can only move laterally in the hall
		return
	}

	//Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room
	for l := c - 1; l > 0; l-- {
		if move1.bur[1][l] != "" {
			break
		}
		if sorted[l] != move1.bur[r][c] { //only bothering with moves I can move into
			continue
		}
		move2 := move1.copy()
		move2.previousStep = &move1
		move2.bur[1][l] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]] * (c - l) //energy * number of moves to get there

		newMoves = append(newMoves, down(1, l, move2)...) //let down handle whether we can move into the sorted column
	}

	return
}

func right(r, c int, move1 move) (newMoves []move) {
	if r != 1 {
		return
	}

	//Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room
	for l := c + 1; l < len(move1.bur[0]); l++ {
		if move1.bur[1][l] != "" {
			break
		}
		if sorted[l] != move1.bur[r][c] { //only bothering with moves I can move into
			continue
		}
		move2 := move1.copy()
		move2.previousStep = &move1
		move2.bur[1][l] = move1.bur[r][c]
		move2.bur[r][c] = ""
		move2.energy += energy[move1.bur[r][c]] * (l - c) //energy * number of moves to get there

		newMoves = append(newMoves, down(1, l, move2)...) //let down handle whether we can move into the sorted column
	}

	return
}

func (bur burrow) done() bool {
	for _, row := range bur {
		for c, ch := range row {
			if ch == "" || ch == "#" {
				continue
			}
			if ch != sorted[c] {
				return false
			}
		}
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

func parse(input string, input2 string) (bur burrow) {
	lines := strings.Split(input, "\n")

	if input2 != "" {
		lines2 := strings.Split(input2, "\n")
		lines1 := make([]string, len(lines[3:]))
		copy(lines1, lines[3:])
		lines = lines[:3]
		lines = append(lines, lines2...)
		lines = append(lines, lines1...)
	}

	//bur = make(burrow, len(lines))

	for r, line := range lines {
		//bur[r] = make([]string, len(line))
		for c, ch := range line {
			if string(ch) == "." {
				bur[r][c] = ""
			} else {
				bur[r][c] = strings.TrimSpace(string(ch))
			}
		}
	}

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

// deep copy move
// I was going to use this to deep copy the slices but that was too slow so I'm using arrays instead
func (old move) copy() (new move) {
	new = old
	//new.bur = make(burrow, len(old.bur))
	//copy(new.bur, old.bur)
	return
}
