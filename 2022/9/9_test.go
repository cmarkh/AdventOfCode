package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1(t *testing.T) {
	motions := parse(test1)
	for _, motion := range motions {
		fmt.Println(motion)
	}
	fmt.Println()

	tailPositions := move(motions)
	for _, pos := range tailPositions {
		fmt.Println(pos)
	}
	fmt.Printf("%v positions\n", len(tailPositions))
	fmt.Println()
}

var test2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestPart2(t *testing.T) {
	motions := parse(test2)
	for _, motion := range motions {
		fmt.Println(motion)
	}
	fmt.Println()

	tailPositions := move2(motions)
	for _, pos := range tailPositions {
		fmt.Println(pos)
	}
	fmt.Printf("%v positions\n", len(tailPositions))
	fmt.Println()
}
