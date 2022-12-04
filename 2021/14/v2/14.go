package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go/2021"
	"time"
)

var inputPath = "../input.txt"

func main() {
	start := time.Now()

	poly, rules, err := input()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("template: %s\n", poly)
	for pair, insert := range rules {
		fmt.Printf("%s -> %s\n", pair, insert)
	}
	fmt.Println()

	count, err := GrowPolymer(poly, rules, 20)
	if err != nil {
		log.Fatal(err)
	}

	for letter, c := range count {
		fmt.Printf("%s: %d\n", letter, c)
	}
	fmt.Println()

	most, least := CommonElements(count)
	fmt.Printf("most: %s at %d\nleast: %s at %d\n", most.Char, most.Count, least.Char, least.Count)
	fmt.Printf("diff: %d\n", most.Count-least.Count)

	fmt.Println("\n", time.Since(start).Seconds(), "seconds")

}

type Rules map[string]string //map[P1P2]insertion
type Count map[string]int    //map[letter]count

func input() (template string, rules Rules, err error) {
	rules = make(Rules)

	input, err := advent.ReadInput(inputPath)
	if err != nil {
		return
	}

	template = input[0]

	for _, line := range input[1:] {
		split := strings.Split(line, " -> ")
		if len(split) != 2 || len(split[0]) != 2 || len(split[1]) != 1 {
			err = fmt.Errorf("unkown line format: %s", line)
			return
		}
		rules[split[0]] = split[1]
	}

	return
}

func GrowPolymer(poly string, rules Rules, maxIterations int) (count Count, err error) {
	count = make(Count)
	for _, ch := range poly { //initial count
		count[string(ch)]++
	}

	getPairs := func(poly string) (pairs []string) {
		for i, ch := range poly {
			if i == len(poly)-1 {
				break
			}
			pairs = append(pairs, string(ch)+string(poly[i+1]))
		}
		return
	}

	var grow func(pair string, iterations int)
	grow = func(pair string, iterations int) {
		if iterations == maxIterations {
			return
		}
		iterations++

		for _, newPair := range getPairs(pair) {
			insert, ok := rules[newPair]
			if !ok {
				log.Fatal("bad pair: ", pair)
			}
			count[insert]++
			grow(string(newPair[0])+insert+string(newPair[1]), iterations)
		}
	}

	grow(poly, 0)

	return
}

type Element struct {
	Char  string
	Count int
}

func CommonElements(count Count) (most, least Element) {
	for char, count := range count {
		if least.Count == 0 || least.Count > count {
			least.Count = count
			least.Char = char
		}
		if most.Count < count {
			most.Count = count
			most.Char = char
		}
	}
	return
}
