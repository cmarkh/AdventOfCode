package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {
	stacks, instructions := parse(test1)
	for _, stack := range stacks {
		fmt.Printf("%+v\n", stack)
	}
	fmt.Println()
	for _, inst := range instructions {
		fmt.Printf("%+v\n", inst)
	}
	fmt.Println()

	stacks = rearrange(stacks, instructions)
	for _, stack := range stacks {
		fmt.Printf("%+v\n", stack)
	}
	fmt.Println()

	top := topCrates(stacks)
	fmt.Printf("crates at top of stacks: %s\n", top)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	stacks, instructions := parse(test1)
	for _, stack := range stacks {
		fmt.Printf("%+v\n", stack)
	}
	fmt.Println()
	for _, inst := range instructions {
		fmt.Printf("%+v\n", inst)
	}
	fmt.Println()

	stacks = rearrange9001(stacks, instructions)
	for _, stack := range stacks {
		fmt.Printf("%+v\n", stack)
	}
	fmt.Println()

	top := topCrates(stacks)
	fmt.Printf("crates at top of stacks: %s\n", top)
	fmt.Println()
}
