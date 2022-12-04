package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	advent "temp/adventofcode/go/2022"
)

func main() {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	elves := parse(lines)

	elves = totalFoodPerElf(elves)

	fmt.Println("Part 1:")

	elf := mostFood(elves)
	fmt.Printf("elf %v has most food: %v\n", elf.index, elf.totalFood)
	fmt.Println()

	fmt.Println("Part 2: ")

	topElves, sum := top3(elves)
	for _, e := range topElves {
		fmt.Printf("elf %v has %v food\n", e.index, e.totalFood)
	}
	fmt.Println()
	fmt.Printf("together they have %v food\n", sum)
	fmt.Println()
}

type elf struct {
	food      []int
	totalFood int
	index     int
}

func parse(lines []string) (elves []elf) {
	e := elf{}
	count := 1
	for _, line := range lines {
		if line == "" {
			elves = append(elves, e)
			e = elf{}
			count++
			continue
		}
		food, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		e.food = append(e.food, food)
		e.index = count
	}
	elves = append(elves, e)

	return
}

func totalFoodPerElf(elves []elf) []elf {
	for i, elf := range elves {
		for _, f := range elf.food {
			elves[i].totalFood += f
		}
	}
	return elves
}

func mostFood(elves []elf) (elfy elf) {
	for _, elf := range elves {
		if elf.totalFood > elfy.totalFood {
			elfy = elf
		}
	}
	return
}

func top3(elves []elf) (topElves []elf, sum int) {
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].totalFood > elves[j].totalFood
	})

	topElves = elves[:3]

	for _, e := range topElves {
		sum += e.totalFood
	}

	return
}
