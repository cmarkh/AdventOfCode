package main

import (
	_ "embed"
	"fmt"
	"log"
	"testing"
)

//go:embed test1.txt
var test1 string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestPart1(t *testing.T) {
	grid, instructions := parse(test1)
	grid.print()
	for _, ins := range instructions {
		fmt.Println(ins)
	}

	pos, facing := grid.move(instructions)
	grid.print()
	pos, facing, password := password(pos, facing)
	fmt.Printf("final pos: %v, facing: %v\n", pos, facing)
	fmt.Printf("password: %v\n", password)
}

func TestPart2(t *testing.T) {
	grid, instructions := parse(test1)
	grid.print()
	for _, ins := range instructions {
		fmt.Println(ins)
	}

	pos, facing, grid := grid.move2Example(instructions)
	grid.print()
	pos, facing, password := password(pos, facing)
	fmt.Printf("final pos: %v, facing: %v\n", pos, facing)
	fmt.Printf("password: %v\n", password)
}

func TestCubeTest(t *testing.T) {
	grid, _ := parse(test1)
	grid.print()

	cube, _ := grid.makeCubeExample()
	cube.printExample()
}

func TestCubeInput(t *testing.T) {
	grid, _ := parse(input)
	//grid.print()

	cube, _ := grid.makeCubeInput()
	cube.printInput()
}
