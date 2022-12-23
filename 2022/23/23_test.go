package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

//lint:ignore U1000 unused
var test2 = `.....
..##.
..#..
.....
..##.
.....`

func TestPart1(t *testing.T) {
	elves := parse(test1)
	elves.print()

	elves, _ = elves.move(10)
	elves.print()
	empty := elves.countEmpty()
	fmt.Printf("empty ground tiles: %v\n", empty)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	elves := parse(test1)
	elves.print()

	elves, rounds := elves.move(0)
	elves.print()
	fmt.Printf("done after %v rounds\n", rounds)
	fmt.Println()
}
