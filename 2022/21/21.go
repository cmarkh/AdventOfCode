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
	monkies := parse(input)

	fmt.Println("Part 1:")
	answer := solve(monkies)
	fmt.Printf("root solves to: %v\n", answer)
	fmt.Println()

	fmt.Println("Part 2:")
	answer = solvePart2(monkies)
	fmt.Printf("I yell: %v\n", answer)
	fmt.Println()
}

type monkey struct {
	name        string
	left, right string //monky names
	operation   string
	specificNum bool //Each monkey is given a job: either to yell a specific number or to yell the result of a math operation
	num         int64
}
type monkies map[string]monkey //map[name]struct

var operations = map[string]func(int64, int64) int64{
	"=": func(a, b int64) int64 {
		fmt.Println(a, b)
		if a == b {
			return 1
		} else {
			return -1
		}
	},
	"+": func(a, b int64) int64 { return a + b },
	"-": func(a, b int64) int64 { return a - b },
	"*": func(a, b int64) int64 { return a * b },
	"/": func(a, b int64) int64 { return a / b },
}

func parse(input string) (monks monkies) {
	var err error
	monks = make(monkies)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		monk := monkey{}

		split := strings.Split(line, ": ")
		monk.name = split[0]

		monk.num, err = strconv.ParseInt(split[1], 10, 64)
		if err == nil {
			monk.specificNum = true
			monks[monk.name] = monk
			continue
		}

		split2 := strings.Split(split[1], " ")
		if len(split2) != 3 {
			log.Fatalf("error understanding line: %v", line)
		}
		monk.left = split2[0]
		monk.right = split2[2]

		monk.operation = split2[1]

		/*var ok bool
		monk.operation, ok = operations[split2[1]]
		if !ok {
			log.Fatalf("error understanding line: %v", line)
		}*/

		monks[monk.name] = monk
	}

	return
}

func solve(monkies monkies) (answer int64) {
	answers := make(map[string]int64) //map[monkey name]value
	queue := []monkey{}

	for _, monk := range monkies {
		if monk.specificNum {
			answers[monk.name] = monk.num
		} else {
			queue = append(queue, monk)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.specificNum {
			answers[current.name] = current.num
			continue
		}

		left, okL := answers[current.left]
		right, okR := answers[current.right]
		if okL && okR {
			answers[current.name] = operations[current.operation](left, right)
		} else {
			queue = append(queue, current)
			continue
		}

		if current.name == "root" {
			return answers[current.name]
		}
	}

	return answers["root"]
}

func solvePart2(monkies monkies) (answer int64) {
	root := monkies["root"]
	root.operation = "="
	monkies["root"] = root

	var solve func(monk monkey) (int64, bool)
	solve = func(monk monkey) (int64, bool) {
		if monk.name == "humn" {
			return -1, false
		}
		if monk.specificNum {
			return monk.num, true
		}
		left, ok := solve(monkies[monk.left])
		if !ok {
			return -1, false
		}
		right, ok := solve(monkies[monk.right])
		if !ok {
			return -1, false
		}
		return operations[monk.operation](left, right), true
	}

	var human int64 = -1
	var reverse func(monk monkey, answer int64)
	reverse = func(monk monkey, answer int64) {
		if monk.name == "humn" {
			human = answer
			return
		}
		if monk.specificNum {
			return
		}
		var newAnswer int64
		known, ok := solve(monkies[monk.left])
		solveSide := monk.right
		if ok {
			switch monk.operation {
			case "=":
				newAnswer = known
			case "+":
				newAnswer = answer - known
			case "*":
				newAnswer = answer / known
			case "-":
				newAnswer = known - answer
			case "/":
				newAnswer = known / answer
			}
		} else {
			known, ok = solve(monkies[monk.right])
			solveSide = monk.left
			if !ok {
				log.Fatal("couldn't figure it out")
			}
			switch monk.operation {
			case "=":
				newAnswer = known
			case "+":
				newAnswer = answer - known
			case "*":
				newAnswer = answer / known
			case "-":
				newAnswer = answer + known
			case "/":
				newAnswer = answer * known
			}
		}
		reverse(monkies[solveSide], newAnswer)
		//fmt.Printf("%s: %v using: %v, from: %v\n", monk.name, answer, monk.operation, known)
	}
	reverse(monkies["root"], 1)

	return human
}
