package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/slices"
)

var inputPath = "../input.txt"

func main() {
	poly, pairs, err := input()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("template: %s\n", poly)
	for pair, insert := range pairs {
		fmt.Printf("%s -> %s\n", pair, insert)
	}

	for i := 1; i <= 6; i++ {
		poly, err = GrowPolymer(pairs, poly)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", i, poly)
	}
	fmt.Println(poly)

	most, least := CommonElements(poly)
	fmt.Printf("most: %s at %d\nleast: %s at %d\n", most.Char, most.Count, least.Char, least.Count)
	fmt.Printf("diff: %d\n", most.Count-least.Count)
}

type Pairs map[string]string //map[P1P2]insertion

func input() (template string, pairs Pairs, err error) {
	pairs = make(Pairs)

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
		pairs[split[0]] = split[1]
	}

	return
}

func GrowPolymer(pairs Pairs, polymer string) (poly string, err error) {
	for i := 0; i < len(polymer)-1; i++ {
		insert, ok := pairs[string(polymer[i])+string(polymer[i+1])]
		if !ok {
			err = fmt.Errorf("unkown pair: %s", string(polymer[i])+string(polymer[i+1]))
			return
		}
		poly += string(polymer[i]) + insert
	}
	poly += string(polymer[len(polymer)-1])

	return
}

type Element struct {
	Count int
	Char  string
}

func CommonElements(polymer string) (most, least Element) {
	least.Count = len(polymer)

	letters := []string{}
	for _, ch := range polymer {
		if !slices.Contains(letters, string(ch)) {
			letters = append(letters, string(ch))
		}
	}

	for _, ch := range letters {
		count := strings.Count(polymer, ch)
		if count > most.Count {
			most.Char = ch
			most.Count = count
		}
		if count < least.Count {
			least.Char = ch
			least.Count = count
		}
	}

	return
}
