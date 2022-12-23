package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.

Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestPart1(t *testing.T) {
	blueprints := parse(test1)
	for _, print := range blueprints {
		fmt.Println(print)
	}
	fmt.Println()

	totalQuality, _, maxes := maxOpenGeodes(blueprints, 24)
	for i, max := range maxes {
		fmt.Printf("max: %v for %v\n", max, i)
	}
	fmt.Println()
	fmt.Printf("total quality: %v\n", totalQuality)
	fmt.Println()

	if totalQuality != 33 {
		t.Fail()
	}
}

func TestOpenGeodes(t *testing.T) {
	blueprints := parse(test1)

	totalQuality := openGeodes(blueprints[1], 24, 1)
	fmt.Printf("total quality: %v\n", totalQuality)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	blueprints := parse(test1)

	_, product, maxes := maxOpenGeodes(blueprints, 32)
	for i, max := range maxes {
		fmt.Printf("max: %v for %v\n", max, i)
	}
	fmt.Println()
	fmt.Printf("product of geodes: %v\n", product)
	fmt.Println()

	if product != 56 {
		t.Fail()
	}
}
