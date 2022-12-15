package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

var (
	divider1, divider2 = `[[2]]`, `[[6]]`
)

func main() {
	pairs := parse(input)

	fmt.Println("Part 1:")
	sum := sumOrderedIndices(pairs)
	fmt.Printf("sum of properly ordered pairs' indices: %v\n", sum)
	fmt.Println()

	fmt.Println("Part 2:")
	packets := parse2(input)
	sort(packets)
	d1, d2, product := divisorIndices(packets)
	fmt.Printf("d1: %v, d2: %v, product: %v\n", d1, d2, product)
	fmt.Println()
}

type pair [2]string

func parse(input string) (pairs []pair) {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		pairs = append(pairs, pair{lines[i], lines[i+1]})
	}

	return
}

func parse2(input string) (packets []string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets = append(packets, line)
	}
	packets = append(packets, divider1)
	packets = append(packets, divider2)
	return
}

func ordered(p1, p2 string) bool {
	var ordered func(p1, p2 string) int
	ordered = func(p1, p2 string) int { //-1 = false, 0 = equal, 1 = true
		//fmt.Printf("0: %v\n", p1)
		//fmt.Printf("1: %v\n", p2)
		//fmt.Println()

		if p2 == "" {
			if p1 == "" {
				return 0
			} else {
				return -1
			}
		}
		if p1 == "" {
			return 1
		}

		if !strings.HasPrefix(p1, "[") && !strings.HasPrefix(p2, "[") {
			num0, err := strconv.Atoi(p1)
			if err != nil {
				log.Fatal(err)
			}
			num1, err := strconv.Atoi(p2)
			if err != nil {
				log.Fatal(err)
			}
			if num0 < num1 {
				return 1
			}
			if num0 == num1 {
				return 0
			}
			return -1
		}

		if strings.HasPrefix(p1, "[") && strings.HasPrefix(p2, "[") {
			split1 := split(p1)
			split2 := split(p2)
			lastRound := 0
			for i := 0; ; i++ {
				if i == len(split2) && i == len(split1) {
					return 0
				}
				if i >= len(split2) {
					return -1
				}
				if i >= len(split1) {
					return 1
				}
				lastRound = ordered(split1[i], split2[i])
				if lastRound == 0 {
					continue
				}
				return lastRound
			}
		}

		if !strings.HasPrefix(p1, "[") {
			p1 = "[" + p1 + "]"
			return ordered(p1, p2)
		}
		if !strings.HasPrefix(p2, "[") {
			p2 = "[" + p2 + "]"
			return ordered(p1, p2)
		}

		log.Fatalf("comparison of %v, %v failed", p1, p2)
		return -1 //should never get here
	}

	return ordered(p1, p2) != -1
}

func split(str string) (split []string) {
	str = strings.TrimSuffix(strings.TrimPrefix(str, "["), "]")
	open := 0
	elementStart := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			open++
			continue
		}
		if str[i] == ']' {
			open--
			continue
		}
		if str[i] == ',' {
			if open != 0 {
				continue
			}
			split = append(split, str[elementStart:i])
			elementStart = i + 1
		}
	}
	split = append(split, str[elementStart:])

	return
}

func sumOrderedIndices(pairs []pair) (sum int) {
	for i, p := range pairs {
		if ordered(p[0], p[1]) {
			sum += i + 1
		}
	}
	return
}

func sort(packets []string) {
	slices.SortStableFunc(packets, ordered)
}

func divisorIndices(packets []string) (d1, d2 int, product int) {
	for i, packet := range packets {
		if packet == divider1 {
			d1 = i + 1 //1 based indexing
		}
		if packet == divider2 {
			d2 = i + 1
		}
	}
	product = d1 * d2
	return
}
