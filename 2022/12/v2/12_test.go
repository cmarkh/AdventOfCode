package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	grid, start, end := parse(test1)
	grid.print()
	fmt.Printf("start: %v, end: %v\n", start, end)

	steps := grid.startAtEnd(start, end, false)
	fmt.Printf("steps: %v\n", steps)

	fmt.Println()
}

func TestPart2(t *testing.T) {
	grid, _, end := parse(test1)
	grid.print()

	steps := grid.startAtEnd(position{}, end, true)
	fmt.Printf("steps: %v\n", steps)

	fmt.Println()
}
