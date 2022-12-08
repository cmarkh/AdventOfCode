package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	grid := parse(test1)
	grid.print()
	grid.visibleTrees()
	grid.printVisible()
	count := grid.countVisible()
	fmt.Printf("%v visible trees\n", count)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	grid := parse(test1)
	best := grid.scenicScores()
	fmt.Printf("best scenic score: %v\n", best)
	fmt.Println()
}
