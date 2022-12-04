package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1(t *testing.T) {
	pairs := parse(test1)
	for _, pair := range pairs {
		fmt.Printf("%+v\n", pair)
	}
	fmt.Println()

	overlaps := fullyOverlaps(pairs)
	for _, overlap := range overlaps {
		fmt.Println(overlap)
	}
	fmt.Printf("%v overlaps\n", len(overlaps))

	fmt.Println()
}

func TestPart2(t *testing.T) {
	pairs := parse(test1)
	for _, pair := range pairs {
		fmt.Printf("%+v\n", pair)
	}
	fmt.Println()

	overlaps := partiallyOverlaps(pairs)
	for _, overlap := range overlaps {
		fmt.Println(overlap)
	}
	fmt.Printf("%v overlaps\n", len(overlaps))

	fmt.Println()
}
