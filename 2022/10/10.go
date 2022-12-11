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

const (
	importantCyclesStart     = 20
	importantCyclesFrequency = 40

	crtWidth  = 40 //pixels
	crtHeight = 6
)

func main() {
	instructions := parse(input)

	fmt.Println("Part 1:")
	sum := signalStrengths(instructions)
	fmt.Printf("sum of signal strengths: %v\n", sum)
	fmt.Println()

	fmt.Println("Part 2:")
	crt := drawCRT(instructions)
	crt.print()
	fmt.Println()
}

type instruction struct {
	instruction string
	value       int
}

func parse(input string) (instructions []instruction) {
	var err error
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		inst := instruction{}
		inst.instruction = split[0]

		if split[0] != "noop" {
			inst.value, err = strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		instructions = append(instructions, inst)
	}
	return
}

func signalStrengths(instructions []instruction) (sum int) {
	xRegister := 1
	cycle := 1
	for _, instruction := range instructions {
		if (cycle-importantCyclesStart)%importantCyclesFrequency == 0 {
			sum += cycle * xRegister
			//fmt.Printf("cycle: %v, register: %v, signal strenght: %v\n", cycle, xRegister, cycle*xRegister)
		}

		switch instruction.instruction {
		case "noop":
			cycle++
			continue

		case "addx":
			if (cycle+1-importantCyclesStart)%importantCyclesFrequency == 0 {
				sum += (cycle + 1) * xRegister
				//fmt.Printf("+1, cycle: %v, register: %v, signal strenght: %v\n", cycle+1, xRegister, (cycle+1)*xRegister)
			}

			cycle += 2
			xRegister += instruction.value
			//fmt.Printf("register: %v\n", xRegister)
			continue

		default:
			log.Fatalf("instruction not understood: %v", instructions[cycle-1])
		}
	}

	return
}

type crt [crtHeight][crtWidth]string

func drawCRT(instructions []instruction) (crt crt) {
	xRegister := 1
	cycle := 0
	row := 0
	add := 0

	for _, instruction := range instructions {
		for {
			if cycle >= crtWidth {
				cycle = cycle/crtWidth - 1
				row++
			}

			//fmt.Printf("cycle: %v, xRegister: %v\n", cycle, xRegister)

			if xRegister-1 == cycle || xRegister == cycle || xRegister+1 == cycle {
				crt[row][cycle] = "#"
			}
			cycle++

			if add != 0 {
				xRegister += instruction.value
				add = 0
				break
			}
			if instruction.instruction == "addx" {
				add = instruction.value
			}
			if instruction.instruction == "noop" {
				break
			}
		}
	}

	return
}

func (crt crt) print() {
	for _, row := range crt {
		for _, col := range row {
			if col == "" {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
