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

func main() {
	monkies := parse(input)

	fmt.Println("Part 1:")
	inspections, _ := takeTurns(monkies, 20)
	monkeyBiz := mostActive(inspections)
	fmt.Printf("level of monkey biz: %v\n", monkeyBiz)
	fmt.Println()

	fmt.Println("Part 2:")
	monkies = parse(input)
	multiple := leastCommonMultiple(monkies)
	inspections, _ = takeTurns2(monkies, 10000, multiple)
	monkeyBiz = mostActive(inspections)
	fmt.Printf("level of monkey biz: %v\n", monkeyBiz)
	fmt.Println()
}

type monkey struct {
	items     []int         //int is each item's worry level
	operation func(int) int //operation done to input item's worry returns item's new worry
	test      func(int) int //if test true return throw to monkey x, if false return throw to monkey y
	divisor   int           //divisor used int test
}

func parse(input string) (monkies []monkey) {
	lines := strings.Split(input, "\n")

	current := monkey{}
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])

		switch {
		case lines[i] == "":
			monkies = append(monkies, current)

		case strings.HasPrefix(lines[i], "Monkey "):
			current = monkey{}

		case strings.HasPrefix(lines[i], "Starting items: "):
			split := strings.Split(strings.TrimPrefix(lines[i], "Starting items: "), ", ")
			for _, item := range split {
				worry, err := strconv.Atoi(item)
				if err != nil {
					log.Fatalf("error interpreting line: %s with %v", lines[i], item)
				}
				current.items = append(current.items, worry)
			}

		case strings.HasPrefix(lines[i], "Operation: "):
			op := strings.TrimPrefix(lines[i], "Operation: new = ")
			split := strings.Split(op, " ")
			if len(split) != 3 {
				log.Fatalf("error interpreting line: %s", lines[i])
			}
			switch split[1] {
			case "+":
				switch split[2] {
				case "old":
					current.operation = func(item int) int {
						return item + item
					}
				default:
					val, err := strconv.Atoi(split[2])
					if err != nil {
						log.Fatalf("error interpreting line: %s", lines[i])
					}
					current.operation = func(item int) int {
						return item + val
					}
				}
			case "*":
				switch split[2] {
				case "old":
					current.operation = func(item int) int {
						return item * item
					}
				default:
					val, err := strconv.Atoi(split[2])
					if err != nil {
						log.Fatalf("error interpreting line: %s", lines[i])
					}
					current.operation = func(item int) int {
						return item * val
					}
				}
			default:
				log.Fatalf("error interpreting line: %s", lines[i])
			}

		case strings.HasPrefix(lines[i], "Test: "):
			if !strings.HasPrefix(strings.TrimSpace(lines[i+1]), "If true: throw to monkey") {
				log.Fatalf("error interpreting line: %s", lines[i+1])
			}
			if !strings.HasPrefix(strings.TrimSpace(lines[i+2]), "If false: throw to monkey") {
				log.Fatalf("error interpreting line: %s", lines[i+2])
			}

			divisor, err := strconv.Atoi(strings.TrimPrefix(lines[i], "Test: divisible by "))
			if err != nil {
				log.Fatalf("error interpreting line: %s", lines[i])
			}
			current.divisor = divisor

			trueMonkey, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(lines[i+1]), "If true: throw to monkey "))
			if err != nil {
				log.Fatalf("error interpreting line: %s", lines[i+1])
			}
			falseMonkey, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(lines[i+2]), "If false: throw to monkey "))
			if err != nil {
				log.Fatalf("error interpreting line: %s", lines[i+2])
			}

			current.test = func(item int) int {
				if item%divisor == 0 {
					return trueMonkey
				} else {
					return falseMonkey
				}
			}

			i += 2

		default:
			log.Fatalf("error interpreting line: %s", lines[i])
		}
	}
	if lines[len(lines)-1] != "" {
		monkies = append(monkies, current)
	}

	return
}

func takeTurns(monkies []monkey, rounds int) (inspections []int, newMonkies []monkey) {
	inspections = make([]int, len(monkies))

	for round := 0; round < rounds; round++ {
		for m, monkey := range monkies {
			for _, item := range monkey.items {
				inspections[m]++
				worry := monkey.operation(item) / 3
				monkies[monkey.test(worry)].items = append(monkies[monkey.test(worry)].items, worry)
			}
			monkies[m].items = []int{}
		}
	}

	return inspections, monkies
}

func mostActive(inspections []int) (monkeyBiz int) {
	slices.Sort(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func takeTurns2(monkies []monkey, rounds int, multiple int) (inspections []int, newMonkies []monkey) {
	inspections = make([]int, len(monkies))

	for round := 0; round < rounds; round++ {
		for m, monkey := range monkies {
			for _, item := range monkey.items {
				inspections[m]++
				worry := monkey.operation(item)
				if worry > multiple {
					worry %= multiple
				}
				monkies[monkey.test(worry)].items = append(monkies[monkey.test(worry)].items, worry)
			}
			monkies[m].items = []int{}
		}
	}

	return inspections, monkies
}

func leastCommonMultiple(monkies []monkey) (multiple int) {
	for i := 1; ; i++ {
		found := true
		for _, monkey := range monkies {
			if i < monkey.divisor {
				found = false
				break
			}
			if i%monkey.divisor != 0 {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
}
