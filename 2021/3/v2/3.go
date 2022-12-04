package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var inputPath = "../input.txt"

func main() {
	input, err := advent.ReadInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	gamma, epsilon, err := Rates(input)
	if err != nil {
		log.Fatal(err)
	}
	product, err := MultiplyRates(gamma, epsilon)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gamma: %s\nepsilon: %s\nmultiplied: %d\n", gamma, epsilon, product)

	oxygen, co2, product, err := LifeSupport(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("oxygen: %d\nco2: %d\nproduct: %d\n", oxygen, co2, product)
}

// Common returns the most and least common digit in the given index
func Common(input []string, index int) (most, least byte, err error) {
	count := make(map[byte]int)
	for _, line := range input {
		count[line[index]]++
	}

	counts := maps.Values(count)
	slices.Sort(counts)

	highCount := counts[len(counts)-1]
	lowCount := counts[0]

	if highCount == lowCount {
		most = 49
		least = 48
		return
	}

	for byt, count := range count {
		if count == highCount {
			most = byt
		}
		if count == lowCount {
			least = byt
		}
	}

	return
}

func Rates(input []string) (gamma, epsilon string, err error) {
	for i := 0; i < len(input[0]); i++ {
		most, least, err := Common(input, i)
		if err != nil {
			return "", "", err
		}
		gamma += string(most)
		epsilon += string(least)
	}
	return
}

func MultiplyRates(gamma, epsilon string) (product int64, err error) {
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return
	}

	return g * e, nil
}

func LifeSupport(input []string) (oxygen, co2, product int64, err error) {
	inputOxy := input
	inputCO2 := input

	for i := 0; i < len(inputOxy[0]); i++ {
		most, _, err := Common(inputOxy, i)
		if err != nil {
			return 0, 0, 0, err
		}
		filtered := []string{}
		for _, line := range inputOxy {
			if line[i] == most {
				filtered = append(filtered, line)
			}
		}
		inputOxy = filtered
		if len(inputOxy) == 1 {
			break
		}
	}
	oxygen, err = strconv.ParseInt(inputOxy[0], 2, 64)
	if err != nil {
		return
	}

	for i := 0; i < len(inputCO2[0]); i++ {
		_, least, err := Common(inputCO2, i)
		if err != nil {
			return 0, 0, 0, err
		}
		filtered := []string{}
		for _, line := range inputCO2 {
			if line[i] == least {
				//fmt.Printf("i: %d, least: %v, line: %s\n", i, least, line)
				filtered = append(filtered, line)
			}
		}
		inputCO2 = filtered
		if len(inputCO2) == 1 {
			break
		}
	}
	co2, err = strconv.ParseInt(inputCO2[0], 2, 64)
	if err != nil {
		return
	}

	product = oxygen * co2

	//fmt.Printf("oxygen: %s\nco2: %s\n", inputOxy, inputCO2)

	return
}
