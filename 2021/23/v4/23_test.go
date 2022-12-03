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
	bur := parse(test1, "")
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
	bur := parse(done, "")
	fmt.Println(bur.done())
}

func TestContains(t *testing.T) {
	bur := parse(done, "")
	bur2 := parse(done, "")
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
###A#B#C#D###
  #A#B#C#D#
  #########`

func TestDown(t *testing.T) {
	bur := parse(hall1, "")
	bur.Print()

	bur[1][5] = "B"
	bur[2][5] = ""
	//bur[3][5] = ""
	bur.Print()

	mov := move{bur, 0, nil}
	moves := down(1, 5, mov)
	if len(moves) > 0 {
		moves[0].bur.Print()
		fmt.Printf("used %v energy\n", moves[0].energy)
	}
	fmt.Println()
}

func TestUp(t *testing.T) {
	bur := parse(hall1, "")
	bur.Print()

	bur[2][5] = ""
	bur[3][5] = "C"
	bur[1][5] = "B"
	bur.Print()

	moves := up(3, 5, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}

func TestLeft(t *testing.T) {
	bur := parse(hall1, "")
	bur.Print()

	bur[2][5] = ""
	bur[3][5] = ""
	bur[1][8] = "B"
	bur.Print()

	moves := left(1, 8, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}

func TestRight(t *testing.T) {
	bur := parse(hall1, "")

	bur[2][5] = ""
	//bur[3][5] = ""
	bur[1][1] = "B"
	bur.Print()

	moves := right(1, 1, move{bur, 0, nil})
	for _, mov := range moves {
		mov.bur.Print()
		fmt.Printf("used %v energy\n\n", mov.energy)
	}

	fmt.Println()
}

func TestPart2(t *testing.T) {
	bur := parse(test1, part2input)
	bur.Print()

	energy, _ := bur.shortestPath()

	fmt.Printf("used %v energy\n", energy)

	fmt.Println()
}

func TestCopy(t *testing.T) {
	bur := parse(test1, part2input)
	bur.Print()

	bur2 := bur
	//bur2 = make(burrow, 2)
	//bur2[1][2] = "f"

	bur.Print()
	bur2.Print()

	fmt.Println()
}
