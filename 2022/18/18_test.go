package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPart1(t *testing.T) {
	cubes := parse(test1)
	for _, cube := range cubes {
		fmt.Printf("%+v\n", cube)
	}
	fmt.Println()

	area := surfaceArea(cubes)
	fmt.Printf("surface area: %v\n", area)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	cubes := parse(test1)
	//fmt.Println(grid)

	area := trappedDrops(cubes)
	fmt.Printf("surface area: %v\n", area)
	fmt.Println()
}
