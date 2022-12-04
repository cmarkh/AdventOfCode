package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	rucks := parse(input)

	fmt.Println("Part 1:")
	sum := 0
	for _, ruck := range rucks {
		common := commonItem(ruck)
		priority := priority(common)
		sum += priority
		//fmt.Printf("ruck: %+v, common item: %+v, priority: %v\n", ruck, string(common), priority)
	}
	fmt.Printf("the sum of priorities is: %v\n", sum)
	fmt.Println()

	fmt.Println("Part 2:")
	sum = 0
	badges := findBadges(rucks)
	for _, badge := range badges {
		//fmt.Printf("group %v badge: %v, priority: %v\n", i+1, string(badge.letter), badge.priority)
		sum += badge.priority
	}
	fmt.Printf("the sum of priorities is: %v\n", sum)
	fmt.Println()
}

type rucksack struct {
	contents           string
	compart1, compart2 string
}

func parse(input string) (rucks []rucksack) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		ruck := rucksack{}
		ruck.contents = line
		ruck.compart1 = line[:len(line)/2]
		ruck.compart2 = line[len(line)/2:]
		rucks = append(rucks, ruck)
	}
	return
}

func commonItem(ruck rucksack) rune {
	for _, c1 := range ruck.compart1 {
		for _, c2 := range ruck.compart2 {
			if c1 == c2 {
				return c1
			}
		}
	}
	log.Fatalf("nothing in common found in: %+v", ruck)
	return 0
}

func priority(ch rune) int {
	if unicode.IsUpper(ch) {
		return int(ch - 'A' + 27)
	}
	return int(ch - 'a' + 1)
}

type badge struct {
	letter   rune
	priority int
}

func findBadges(rucks []rucksack) (badges []badge) {
	for i := range rucks {
		if (i+1)%3 == 0 { //every third ruck
			badges = append(badges, findBadge(rucks[i-2:i+1]))
		}
	}
	return
}

func findBadge(rucks []rucksack) badge {
	for _, r1 := range rucks[0].contents {
		for _, r2 := range rucks[1].contents {
			if r1 != r2 {
				continue
			}
			for _, r3 := range rucks[2].contents {
				if r3 == r1 {
					return badge{r1, priority(r1)}
				}
			}
		}
	}
	return badge{}
}
