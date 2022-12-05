package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

func TestPart1(t *testing.T) {
	grid := parse(test1)
	grid.print()

	steps := grid.standstill()
	fmt.Printf("gridlock in %v steps\n", steps)
	fmt.Println()
}
