package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// this is a naive / brute force attempt

func main() {
	strInstructions := sanitize(input)
	for _, instruction := range strInstructions {
		fmt.Println(instruction)
	}
	fmt.Println()

	modelNum := highestValidModelNum(strInstructions)
	fmt.Printf("largest valid model number: %v\n", modelNum)
	fmt.Println()
}

type instruction [3]string

func sanitize(input string) (instructions []instruction) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		switch len(split) {
		case 2:
			instructions = append(instructions, [3]string{split[0], split[1]})
		case 3:
			instructions = append(instructions, [3]string{split[0], split[1], split[2]})
		default:
			log.Fatalf("instructions not understood: %v", line)
		}
	}
	return
}

type alu map[string]int //map[variable]value

func executeInstructions(instructions []instruction, modelNum string) (vars alu) {
	fmt.Printf("testing %v\n", modelNum)

	vars = make(alu)
	modelUsed := 0

	for _, instruction := range instructions {
		var b int
		var err error
		b, err = strconv.Atoi(instruction[2])
		if err != nil {
			b = vars[instruction[2]]
		}

		switch instruction[0] {
		case "inp":
			vars[instruction[1]], err = strconv.Atoi(string(modelNum[modelUsed]))
			if err != nil {
				log.Fatal(err)
			}
			modelUsed++
		case "add":
			vars[instruction[1]] += b
		case "mul":
			vars[instruction[1]] *= b
		case "div":
			if b == 0 {
				vars["z"] = 1 //meaning modelNum is invalid
				return
				//log.Fatalf("can't divide by zero: %+v", instruction)
			}
			var quotient float64 = float64(vars[instruction[1]]) / float64(b)
			vars[instruction[1]] = int(math.Floor(quotient))
		case "mod":
			if vars[instruction[1]] < 0 || b <= 0 {
				vars["z"] = 1 //meaning modelNum is invalid
				return
				//log.Fatalf("can't mod with a<0 or b<=0: %+v", instruction)
			}
			vars[instruction[1]] %= b
		case "eql":
			if vars[instruction[1]] == b {
				vars[instruction[1]] = 1
			} else {
				vars[instruction[1]] = 0
			}
		default:
			log.Fatalf("instruction not understood: %+v", instruction)
		}
	}

	return
}

func (vars alu) valid() bool {
	return vars["z"] == 0
}

func highestValidModelNum(instructions []instruction) (modelNum int) {
	for i := 99999999999999; i > 9999999999999; i-- {
		if executeInstructions(instructions, fmt.Sprint(i)).valid() {
			return i
		}
	}
	return
}
