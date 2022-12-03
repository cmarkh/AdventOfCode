package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

var done = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########`

func TestPart1(t *testing.T) {
	bur := parse(test1)
	bur.Print()

	energy, final := bur.shortestPath()

	for {
		final.bur.Print()
		if final.previousStep == nil {
			break
		}
		fmt.Printf("energy: %v\n", final.energy-final.previousStep.energy)
		final = *final.previousStep
		fmt.Println("-----")
	}

	fmt.Printf("used %v energy\n", energy)

	fmt.Println()
}

func TestDone(t *testing.T) {
	bur := parse(done)
	fmt.Println(bur.done())
}

func TestContains(t *testing.T) {
	bur := parse(done)
	bur2 := parse(done)
	checked := []move{{bur, 10, nil}}

	mov2 := move{bur2, 5, nil}

	bur3 := bur2
	bur3[0][3] = "F"

	fmt.Println(checked)
	fmt.Println(contains(checked, mov2))
	fmt.Println(contains(checked, move{bur3, 0, nil}))
	fmt.Println(checked)
}
