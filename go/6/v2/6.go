package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "6/input.txt"

func main() {
	ages, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	days := 256
	ages = ages.Increment(days)
	fmt.Printf("After %d turns, there are %d fish\n", days, ages.Count())
}

type Ages [9]int

func Input(path string) (ages Ages, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		strAges := strings.Split(line, ",")
		for _, a := range strAges {
			age, err := strconv.Atoi(a)
			if err != nil {
				return ages, err
			}
			if age >= len(ages) {
				err = fmt.Errorf("possible issue reading input")
				return ages, err
			}
			ages[age]++
		}
	}

	return
}

func (ages Ages) Increment(days int) Ages {
	for d := 0; d < days; d++ {
		incremented := Ages{}
		for i, age := range ages {
			switch i {
			case 0:
				incremented[6] += age
				incremented[8] += age //new fish
			default:
				incremented[i-1] += age
			}
		}
		ages = incremented
		fmt.Printf("day: %d, fish: %d, list: %v\n", d+1, ages.Count(), ages)
	}
	return ages
}

func (ages *Ages) Count() (count int) {
	for _, fish := range ages {
		count += fish
	}
	return
}
