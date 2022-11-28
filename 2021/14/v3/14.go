package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go"
	"time"
)

var inputPath = "../input.txt"

func main() {
	start := time.Now()

	poly, rules, err := input()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("template: %s\n\n", poly)

	for pair, insert := range rules {
		fmt.Printf("%s -> %s\n", pair, insert)
	}
	fmt.Println()

	tripletRules := TripletRules(rules)
	tripletRules.Print()

	run := func(iter int) {
		count := Iterate(poly, rules, tripletRules, iter)
		for letter, c := range count {
			fmt.Printf("%s: %d\n", letter, c)
		}
		fmt.Println()

		most, least := CommonElements(count)
		fmt.Printf("most: %s at %d\nleast: %s at %d\n", most.Char, most.Count, least.Char, least.Count)
		fmt.Printf("diff: %d\n", most.Count-least.Count)
		fmt.Println()
	}

	fmt.Printf("Part 1:\n\n")
	run(10)

	fmt.Printf("Part 2:\n\n")
	run(40)

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

type TripletProducts map[string][]string //map[triplet]products on insertions into triplet

func TripletRules(rules Rules) (products TripletProducts) {
	products = make(TripletProducts)

	for pair, insert := range rules {
		triplet := string(pair[0]) + insert + string(pair[1])

		prods := []string{} //each triplet produces exactly 2 new triplets
		insert := rules[string(triplet[0])+string(triplet[1])]
		prods = append(prods, string(triplet[0])+insert+string(triplet[1]))

		insert = rules[string(triplet[1])+string(triplet[2])]
		prods = append(prods, string(triplet[1])+insert+string(triplet[2]))

		products[triplet] = prods
	}

	return
}

type Triplets map[string]int //map[triplet]count

func Iterate(poly string, rules Rules, tripletRules TripletProducts, iterations int) (count Count) {
	count = make(Count)
	triplets := make(Triplets)

	//craft first set of triplets from poly
	for i := range poly {
		if i == len(poly)-1 {
			break
		}
		insert := rules[string(poly[i])+string(poly[i+1])]
		triplets[string(poly[i])+insert+string(poly[i+1])]++
	}

	//build all remaining triplets from original triplets (iteration 1 is already done above)
	for i := 0; i < iterations-1; i++ {
		newTriplets := make(Triplets)
		for tripl, c := range triplets {
			for _, product := range tripletRules[tripl] {
				newTriplets[product] += c
			}
		}
		triplets = newTriplets
	}
	triplets.Print()

	for tripl, c := range triplets {
		// don't count tripl[2] as it is contained in the neighboring triplet
		// ie one triplet's tripl[2] is also another's tripl[0]
		count[string(tripl[0])] += c
		count[string(tripl[1])] += c
	}
	//need to count the last letter in the whole poly (not counted when counting triples)
	count[string(poly[len(poly)-1])]++

	return
}

func (rules TripletProducts) Print() {
	for tripl, rule := range rules {
		fmt.Printf("%s -> %s\n", tripl, rule)
	}
	fmt.Println()
}

func (triplets Triplets) Print() {
	for tripl, count := range triplets {
		fmt.Printf("%s: %d\n", tripl, count)
	}
	fmt.Println()
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
