package day8

import (
	"strings"
)

const (
	Left = iota
	Right
)

type Instructions []int

type Map map[string]Destinations

type Destinations struct {
	Left  string
	Right string
}

func ParseInput(input string) (instructions Instructions, mapIn Map) {
	mapIn = make(Map)

	for i, line := range strings.Split(input, "\n") {
		switch i {
		case 0:
			for _, char := range line {
				switch char {
				case 'L':
					instructions = append(instructions, Left)
				case 'R':
					instructions = append(instructions, Right)
				}
			}
		case 1:
			continue
		default:
			split := strings.Split(line, " = ")

			start := split[0]

			destinations := strings.Split(split[1], ", ")

			left := strings.TrimLeft(destinations[0], "(")
			left = strings.TrimRight(left, ")")

			right := strings.TrimLeft(destinations[1], "(")
			right = strings.TrimRight(right, ")")

			mapIn[start] = Destinations{left, right}
		}
	}

	return
}

func Part1(instructions Instructions, mapIn Map) (steps uint64) {
	position := "AAA"

	for i := 0; ; i = (i + 1) % len(instructions) {
		if position == "ZZZ" {
			return steps
		}
		switch instructions[i] {
		case Left:
			position = mapIn[position].Left
		case Right:
			position = mapIn[position].Right
		}
		steps++
	}

}
