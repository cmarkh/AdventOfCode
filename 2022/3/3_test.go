package main

import (
	"fmt"
	"log"
	"testing"
)

var test1 = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestPart1(t *testing.T) {
	rucks := parse(test1)

	sum := 0
	for _, ruck := range rucks {
		common := commonItem(ruck)
		priority := priority(common)
		sum += priority
		fmt.Printf("ruck: %+v, common item: %+v, priority: %v\n", ruck, string(common), priority)
	}
	fmt.Printf("the sum of priorities is: %v\n", sum)

	fmt.Println()
}

func TestPart2(t *testing.T) {
	rucks := parse(test1)

	sum := 0
	badges := findBadges(rucks)
	for i, badge := range badges {
		fmt.Printf("group %v badge: %v, priority: %v\n", i+1, string(badge.letter), badge.priority)
		sum += badge.priority
	}
	fmt.Printf("the sum of priorities is: %v\n", sum)

	fmt.Println()
}
