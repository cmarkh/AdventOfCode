package advent2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "2/input.txt"

type Instruction struct {
	Direction string
	Amount    int64
}

// input here are depth measurements
func Input() (instructions []Instruction, err error) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		direction, amount, found := strings.Cut(line, " ")
		if !found {
			err = fmt.Errorf("error reading line of input: %s", line)
			return
		}

		instruction := Instruction{}
		instruction.Direction = direction

		instruction.Amount, err = strconv.ParseInt(amount, 10, 64)
		if err != nil {
			err = fmt.Errorf("error parsing direction amount: %s", line)
			return
		}

		instructions = append(instructions, instruction)
	}

	return
}
