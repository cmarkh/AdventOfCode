package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func TestPart1(t *testing.T) {
	monkies := parse(test1)
	for _, monkey := range monkies {
		fmt.Printf("%+v\n", monkey)
	}
	fmt.Println()
}

func TestPart12(t *testing.T) {
	monkies := parse(test1)
	inspections, monkies := takeTurns(monkies, 20)
	for m, inspection := range inspections {
		fmt.Printf("monkey %v inspections: %v\n", m, inspection)
	}
	fmt.Println()

	for m, monkey := range monkies {
		fmt.Printf("monkey %v holds %v\n", m, monkey.items)
	}
	fmt.Println()

	monkeyBiz := mostActive(inspections)
	fmt.Printf("level of monkey biz: %v\n", monkeyBiz)
	fmt.Println()
}

func TestPart13(t *testing.T) {
	monkies := parse(input)
	inspections, monkies := takeTurns(monkies, 20)
	for m, inspection := range inspections {
		fmt.Printf("monkey %v inspections: %v\n", m, inspection)
	}
	fmt.Println()

	for m, monkey := range monkies {
		fmt.Printf("monkey %v holds %v\n", m, monkey.items)
	}
	fmt.Println()

	monkeyBiz := mostActive(inspections)
	fmt.Printf("level of monkey biz: %v\n", monkeyBiz)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	monkies := parse(test1)

	multiple := leastCommonMultiple(monkies)
	fmt.Printf("least common multiple: %v\n\n", multiple)

	inspections, monkies := takeTurns2(monkies, 10000, multiple)
	for m, inspection := range inspections {
		fmt.Printf("monkey %v inspections: %v\n", m, inspection)
	}
	fmt.Println()

	for m, monkey := range monkies {
		fmt.Printf("monkey %v holds %v\n", m, monkey.items)
	}
	fmt.Println()

	monkeyBiz := mostActive(inspections)
	fmt.Printf("level of monkey biz: %v\n", monkeyBiz)
	fmt.Println()
}
