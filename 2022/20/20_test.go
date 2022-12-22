package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `1
2
-3
3
-2
0
4`

func TestPart1(t *testing.T) {
	file := parse(test1)
	for _, num := range file {
		fmt.Println(num)
	}
	fmt.Println()

	file = mix(file)
	file, sum, x, y, z := grove(file)
	print(file)
	fmt.Println()
	fmt.Printf("%v, %v, %v\nsum: %v\n", x, y, z, sum)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	file := parse(test1)

	file = mixPart2(file)
	file, sum, x, y, z := grove(file)
	print(file)
	fmt.Println()
	fmt.Printf("%v, %v, %v\nsum: %v\n", x, y, z, sum)
	fmt.Println()
}

func TestMix2(t *testing.T) {
	file := parse(test1)
	file = mix(file)
	print(file)
	fmt.Println()
}
