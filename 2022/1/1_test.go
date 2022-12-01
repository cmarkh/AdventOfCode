package main

import (
	"fmt"
	"strings"
	"testing"
)

var test1 = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestParse(t *testing.T) {
	elves := parse(strings.Split(test1, "\n"))
	for _, elf := range elves {
		fmt.Println(elf)
	}
	fmt.Println()
}

func TestTotalFoodPerElf(t *testing.T) {
	elves := parse(strings.Split(test1, "\n"))
	elves = totalFoodPerElf(elves)
	for _, elf := range elves {
		fmt.Println(elf)
	}
	fmt.Println()
}

func TestMostFood(t *testing.T) {
	elves := parse(strings.Split(test1, "\n"))
	elves = totalFoodPerElf(elves)
	elf := mostFood(elves)
	fmt.Printf("elf %v has most food: %v\n", elf.index, elf.totalFood)
	fmt.Println()
}

func TestTop3(t *testing.T) {
	elves := parse(strings.Split(test1, "\n"))
	elves = totalFoodPerElf(elves)
	topElves, sum := top3(elves)
	for _, e := range topElves {
		fmt.Printf("elf %v has %v food\n", e.index, e.totalFood)
	}
	fmt.Println()
	fmt.Printf("together they have %v food\n", sum)
	fmt.Println()
}
