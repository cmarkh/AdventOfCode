package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:")
	stacks, instructions := parse(input)
	stacks = rearrange(stacks, instructions)
	top := topCrates(stacks)
	fmt.Printf("crates at top of stacks: %s\n", top)
	fmt.Println()

	fmt.Println("Part 2:")
	stacks, instructions = parse(input)
	stacks = rearrange9001(stacks, instructions)
	top = topCrates(stacks)
	fmt.Printf("crates at top of stacks: %s\n", top)
	fmt.Println()
}

type stack []string //each stack contains a slice of crates designated by letter

type instruction struct {
	moveCount int
	from, to  int
}

func parse(input string) (stacks []stack, instructions []instruction) {
	var err error
	lines := strings.Split(input, "\n")

	stackCount := 0
	countLine := 0
	for i, line := range lines {
		if unicode.IsNumber(rune(line[1])) {
			stackCount, err = strconv.Atoi(string(line[len(line)-2]))
			if err != nil {
				log.Fatal(err)
			}
			countLine = i
			break
		}
	}
	stacks = make([]stack, stackCount)

	s := 0
	for c := 0; c < len(lines[countLine]); c++ {
		if !unicode.IsNumber(rune(lines[countLine][c])) {
			continue
		}
		for r := countLine - 1; r >= 0; r-- { //starting with the line containing stack numbers, work up and add the crates in each stack
			if strings.TrimSpace(string(lines[r][c])) == "" {
				continue
			}
			stacks[s] = append(stacks[s], string(lines[r][c]))
		}
		s++
	}

	//instructions:
	for i := countLine + 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		split := strings.Split(lines[i], " ")
		move := instruction{}

		for i := range split {
			switch split[i] {
			case "move":
				move.moveCount, err = strconv.Atoi(split[i+1])
				if err != nil {
					log.Fatal(err)
				}
			case "from":
				move.from, err = strconv.Atoi(split[i+1])
				if err != nil {
					log.Fatal(err)
				}
				move.from-- //instructions are 1 based but my slice is 0 based
			case "to":
				move.to, err = strconv.Atoi(split[i+1])
				if err != nil {
					log.Fatal(err)
				}
				move.to--
			}
		}

		instructions = append(instructions, move)
	}

	return
}

func rearrange(stacks []stack, instructions []instruction) []stack {
	for _, instruction := range instructions {
		for c := 0; c < instruction.moveCount; c++ {
			stacks[instruction.to] = append(stacks[instruction.to],
				stacks[instruction.from][len(stacks[instruction.from])-1])
			stacks[instruction.from] = stacks[instruction.from][:len(stacks[instruction.from])-1]
		}
	}
	return stacks
}

func rearrange9001(stacks []stack, instructions []instruction) []stack {
	for _, instruction := range instructions {
		move := make([]string, instruction.moveCount)
		copy(move, stacks[instruction.from][len(stacks[instruction.from])-instruction.moveCount:])

		stacks[instruction.to] = append(stacks[instruction.to], move...)
		stacks[instruction.from] = stacks[instruction.from][:len(stacks[instruction.from])-instruction.moveCount]
	}
	return stacks
}

func topCrates(stacks []stack) (crates string) {
	for _, stack := range stacks {
		crates += stack[len(stack)-1]
	}

	return
}
