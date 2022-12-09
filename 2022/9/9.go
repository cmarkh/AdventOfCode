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
	motions := parse(input)

	fmt.Println("Part 1:")
	tailPositions := move(motions)
	fmt.Printf("%v positions\n", len(tailPositions))
	fmt.Println()

	fmt.Println("Part 2:")
	tailPositions = move2(motions)
	fmt.Printf("%v positions\n", len(tailPositions))
	fmt.Println()
}

type motion struct {
	direction string
	steps     int
}

func parse(input string) (motions []motion) {
	var err error
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		motion := motion{}
		motion.direction = split[0]
		motion.steps, err = strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		motions = append(motions, motion)
	}
	return
}

type position struct {
	x, y int
}

func move(motions []motion) (tailPositions []position) {
	head, tail := position{}, position{}

	for _, motion := range motions {
		for s := 0; s < motion.steps; s++ {
			tailPositions = append(tailPositions, tail)

			switch motion.direction {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			default:
				log.Fatalf("direction not understood: %v", motion)
			}

			distance := position{}
			distance.x = abs(head.x - tail.x)
			distance.y = abs(head.y - tail.y)

			if distance.x <= 1 && distance.y <= 1 {
				continue
			}

			if distance.x == 0 {
				if head.y < tail.y {
					tail.y--
				} else {
					tail.y++
				}
				continue
			}
			if distance.y == 0 {
				if head.x < tail.x {
					tail.x--
				} else {
					tail.x++
				}
				continue
			}

			if head.x < tail.x {
				tail.x--
			} else {
				tail.x++
			}
			if head.y < tail.y {
				tail.y--
			} else {
				tail.y++
			}

		}
	}
	tailPositions = append(tailPositions, tail)

	trimmed := []position{}
	for _, pos := range tailPositions {
		if !slices.Contains(trimmed, pos) {
			trimmed = append(trimmed, pos)
		}
	}
	return trimmed
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func move2(motions []motion) (tailPositions []position) {
	ropes := [10]position{}
	tailPositions = append(tailPositions, ropes[0])

	for _, motion := range motions {
		for s := 0; s < motion.steps; s++ {
			switch motion.direction {
			case "U":
				ropes[9].y++
			case "D":
				ropes[9].y--
			case "L":
				ropes[9].x--
			case "R":
				ropes[9].x++
			default:
				log.Fatalf("direction not understood: %v", motion)
			}

			for i := 8; i >= 0; i-- {

				distance := position{}
				distance.x = abs(ropes[i+1].x - ropes[i].x)
				distance.y = abs(ropes[i+1].y - ropes[i].y)

				if distance.x <= 1 && distance.y <= 1 {
					continue
				}

				if distance.x == 0 {
					if ropes[i+1].y < ropes[i].y {
						ropes[i].y--
					} else {
						ropes[i].y++
					}
				} else if distance.y == 0 {
					if ropes[i+1].x < ropes[i].x {
						ropes[i].x--
					} else {
						ropes[i].x++
					}
				} else {

					if ropes[i+1].x < ropes[i].x {
						ropes[i].x--
					} else {
						ropes[i].x++
					}
					if ropes[i+1].y < ropes[i].y {
						ropes[i].y--
					} else {
						ropes[i].y++
					}
				}

				if i == 0 {
					tailPositions = append(tailPositions, ropes[i])
				}
			}
		}
	}

	trimmed := []position{}
	for _, pos := range tailPositions {
		if !slices.Contains(trimmed, pos) {
			trimmed = append(trimmed, pos)
		}
	}
	return trimmed
}
