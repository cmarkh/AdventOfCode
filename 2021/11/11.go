package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
)

var inputPath = "input.txt"

func main() {
	octopi, err := input()
	if err != nil {
		log.Fatal(err)
	}
	octopi.Print()
	fmt.Println()

	var flashes int
	steps := 100
	for i := 0; i < steps; i++ {
		//fmt.Println(i + 1)
		flashes += octopi.Increment()
		//octopi.Print()
		//fmt.Println()
	}
	fmt.Printf("flashes after %d steps: %d\n\n", steps, flashes)

	octopi, err = input()
	if err != nil {
		log.Fatal(err)
	}
	iteration := octopi.EveryoneFlashes()
	fmt.Printf("everyone flashes after %d steps\n", iteration)
}

type Octopi [][]int

func input() (octopi Octopi, err error) {
	input, err := advent.ReadInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range input {
		octs := []int{}
		for _, oct := range strings.Split(line, "") {
			intOct, err := strconv.Atoi(oct)
			if err != nil {
				return nil, err
			}
			octs = append(octs, intOct)
		}
		octopi = append(octopi, octs)
	}
	return
}

func (octopi Octopi) Print() {
	for _, line := range octopi {
		for _, oct := range line {
			fmt.Print(oct)
		}
		fmt.Println()
	}
}

func (octopi Octopi) Increment() (flashes int) {
	for l, line := range octopi {
		for o := range line {
			octopi[l][o]++
		}
	}

	neighbors := func(row, col int) {
		if row >= 0 && row < len(octopi) &&
			col >= 0 && col < len(octopi[row]) {
			if octopi[row][col] != 0 {
				octopi[row][col]++
			}
		}
	}

	anyNines := true
	for anyNines {
		anyNines = false
		for l, line := range octopi {
			for o, oct := range line {
				if oct > 9 {
					flashes++
					anyNines = true
					octopi[l][o] = 0
					neighbors(l+1, o)
					neighbors(l-1, o)
					neighbors(l, o+1)
					neighbors(l, o-1)
					neighbors(l+1, o+1)
					neighbors(l+1, o-1)
					neighbors(l-1, o+1)
					neighbors(l-1, o-1)
				}
			}
		}
	}
	return
}

func (octopi Octopi) EveryoneFlashes() (iteration int) {
	for iteration = 1; ; iteration++ {
		flashes := octopi.Increment()
		fmt.Printf("%d: %d flashes\n", iteration, flashes)
		octopi.Print()
		fmt.Println()
		if flashes == len(octopi)*len(octopi[0]) {
			return
		}
	}
}
