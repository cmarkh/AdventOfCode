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

}

type burrow struct {
	hall  [11]string
	rooms [4]room
}

type room [2]string

var energy = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

var hallPos = map[int]int{ //map[room]position in hall
	0: 3,
	1: 5,
	2: 7,
	3: 9,
}

var sorted = map[int]string{ //map[room]character if sorted
	0: "A",
	1: "B",
	2: "C",
	3: "D",
}

type position struct {
	inRoom bool //in a room or if false then in hall
	pos    int  //room or hall position
}

func parse(input string) (b burrow) {
	lines := strings.Split(input, "\n")

	b.rooms[0] = [2]string{string(lines[2][3]), string(lines[3][3])}
	b.rooms[1] = [2]string{string(lines[2][5]), string(lines[3][5])}
	b.rooms[2] = [2]string{string(lines[2][7]), string(lines[3][7])}
	b.rooms[3] = [2]string{string(lines[2][9]), string(lines[3][9])}

	return
}

// energyToFillRoom is same as fillRoom except takes a copy of burrow so doesn't actually do it, just checks the energy required
func (bur burrow) energyToFillRoom(room int) (energy int, blocked, done bool) {
	return bur.fillRoom(room)
}

func (bur *burrow) fillRoom(room int) (energy int, blocked, done bool) {
	if bur.rooms[room][0] == sorted[room] && bur.rooms[room][1] == sorted[room] {
		done = true
		return
	}

	matches := bur.findMatches(room, sorted[room])
	fmt.Printf("match %s is in room %v %v, distance: %v\n", sorted[room], matches[0].inRoom, matches[0].pos, distance(room, matches[0]))
	fmt.Printf("match %s is in room %v %v, distance: %v\n", sorted[room], matches[1].inRoom, matches[1].pos, distance(room, matches[1]))

	return
}

func (bur burrow) findMatches(fromRoom int, ch string) (matches []position) {
	for r, room := range bur.rooms {
		if room[0] == ch || room[1] == ch {
			matches = append(matches, position{true, r})
		}
	}

	for i, content := range bur.hall {
		if content == ch {
			matches = append(matches, position{false, i})
		}
	}

	if len(matches) != 2 {
		log.Fatal("couldn't find " + ch)
	}

	return
}

func distance(fromRoom int, pos position) int {
	if pos.inRoom {
		if fromRoom == pos.pos {
			return 0
		}
		return abs(fromRoom-pos.pos) + 1
	}
	return abs(hallPos[fromRoom] - pos.pos)
}

func (b burrow) Print() {
	dot := func(ch string) string {
		if ch == "" {
			return "."
		}
		return ch
	}

	for i := 0; i < 13; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	fmt.Print("#")
	for _, ch := range b.hall {
		fmt.Print(dot(ch))
	}
	fmt.Println("#")

	fmt.Print("###")
	fmt.Print(dot(b.rooms[0][0]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[1][0]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[2][0]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[3][0]))
	fmt.Println("###")

	fmt.Print("  #")
	fmt.Print(dot(b.rooms[0][1]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[1][1]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[2][1]))
	fmt.Print("#")
	fmt.Print(dot(b.rooms[3][1]))
	fmt.Println("#  ")

	fmt.Println("  #########")

	fmt.Println()
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
