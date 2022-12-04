package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pairs := parse(input)

	fmt.Println("Part 1:")
	overlaps := fullyOverlaps(pairs)
	fmt.Printf("%v overlaps\n", len(overlaps))
	fmt.Println()

	fmt.Println("Part 2:")
	overlaps = partiallyOverlaps(pairs)
	fmt.Printf("%v overlaps\n", len(overlaps))
	fmt.Println()
}

type pair struct {
	sections1, sections2 []int
}

func parse(input string) (pairs []pair) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		pair := pair{}

		split := strings.Split(line, ",")
		s1 := strings.Split(split[0], "-")
		s2 := strings.Split(split[1], "-")

		start1, err := strconv.Atoi(s1[0])
		if err != nil {
			log.Fatal(err)
		}
		end1, err := strconv.Atoi(s1[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := start1; i <= end1; i++ {
			pair.sections1 = append(pair.sections1, i)
		}

		start2, err := strconv.Atoi(s2[0])
		if err != nil {
			log.Fatal(err)
		}
		end2, err := strconv.Atoi(s2[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := start2; i <= end2; i++ {
			pair.sections2 = append(pair.sections2, i)
		}

		pairs = append(pairs, pair)
	}

	return
}

func fullyOverlaps(pairs []pair) (overlapping []pair) {
	for _, pair := range pairs {
		if len(pair.sections1) >= len(pair.sections2) {
			if pair.sections1[0] <= pair.sections2[0] &&
				pair.sections1[len(pair.sections1)-1] >= pair.sections2[len(pair.sections2)-1] {
				overlapping = append(overlapping, pair)
				continue
			}
		}
		if pair.sections2[0] <= pair.sections1[0] &&
			pair.sections2[len(pair.sections2)-1] >= pair.sections1[len(pair.sections1)-1] {
			overlapping = append(overlapping, pair)
		}
	}
	return
}

func partiallyOverlaps(pairs []pair) (overlapping []pair) {
	for _, pair := range pairs {
		if func() bool {
			for _, s1 := range pair.sections1 {
				for _, s2 := range pair.sections2 {
					if s1 == s2 {
						return true
					}
				}
			}
			return false
		}() {
			overlapping = append(overlapping, pair)
		}
	}

	return
}
