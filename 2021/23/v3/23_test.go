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

var hall1 = `#############
#...........#
###A#Bs#C#D###
  #A#B#C#D#
  #########`

func TestDown(t *testing.T) {
	bur := parse(hall1)

	bur[0][4] = "B"
	bur[1][4] = ""
	//bur[2][4] = ""
	bur.Print()

	moves := down(0, 4, move{bur, 0, nil})
	if len(moves) > 0 {
		moves[0].bur.Print()
		fmt.Printf("used %v energy\n", moves[0].energy)
	}

	fmt.Println()
}

func TestUp(t *testing.T) {
	bur := parse(hall1)

	bur[1][4] = ""
	bur[2][4] = "C"
	bur[0][8] = "B"
	bur.Print()

	moves := up(2, 4, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}

func TestLeft(t *testing.T) {
	bur := parse(hall1)

	bur[1][4] = ""
	//bur[2][4] = ""
	bur[0][8] = "B"
	bur.Print()

	moves := left(0, 8, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}

func TestRight(t *testing.T) {
	bur := parse(hall1)

	bur[1][4] = ""
	//bur[2][4] = ""
	bur[0][1] = "B"
	bur.Print()

	moves := right(0, 1, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}
