package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestPart1(t *testing.T) {
	grid := parse(test1)
	grid.print()
	sand := grid.produceSand()
	grid.print()
	fmt.Printf("%v units of sand stacked up\n", sand)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	grid := parse(test1)
	grid.print()
	sand := grid.produceSand2()
	grid.print()
	fmt.Printf("%v units of sand stacked up\n", sand)
	fmt.Println()
}

func TestPart22(t *testing.T) {
	grid := parse(input)
	grid.print()
	sand := grid.produceSand2()
	grid.print()
	fmt.Printf("%v units of sand stacked up\n", sand)
	fmt.Println()
}
