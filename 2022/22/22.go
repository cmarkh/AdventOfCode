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
	grid, instructions := parse(input)

	fmt.Println("Part 1:")
	pos, facing := grid.move(instructions)
	//grid.print()
	pos, facing, pswd := password(pos, facing)
	fmt.Printf("final pos: %v, facing: %v\n", pos, facing)
	fmt.Printf("password: %v\n", pswd)
	fmt.Println()

	fmt.Println("Part 2:")
	grid, instructions = parse(input)
	pos, facing, _ = grid.move2Input(instructions)
	//grid.print()
	pos, facing, pswd = password(pos, facing)
	fmt.Printf("final pos: %v, facing: %v\n", pos, facing)
	fmt.Printf("password: %v\n", pswd)
	fmt.Println()
}

type grid [][]string

type instruction struct {
	move int
	turn string
}

func parse(input string) (grid grid, instructions []instruction) {
	var err error
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	maxLen := func() (max int) {
		for _, line := range lines {
			if line == "" {
				break
			}
			if len(line) > max {
				max = len(line)
			}
		}
		return
	}()

	parseInstructions := func(line string) {
		strMove := ""
		for i := 0; i < len(line); i++ {
			if unicode.IsNumber(rune(line[i])) {
				strMove += string(line[i])
				if i == len(line)-1 {
					instruction := instruction{}
					instruction.move, err = strconv.Atoi(strMove)
					if err != nil {
						log.Fatal(err)
					}
					instructions = append(instructions, instruction)
				}
			} else {
				instruction := instruction{}
				instruction.move, err = strconv.Atoi(strMove)
				if err != nil {
					log.Fatal(err)
				}
				strMove = ""

				if unicode.IsNumber(rune(line[i])) {
					log.Fatalf("unexpected number, expected direction: %v", string(line[i]))
				}
				instruction.turn = string(line[i])
				instructions = append(instructions, instruction)
			}
		}
	}

	for i, line := range lines {
		if line == "" {
			continue
		}
		if i == len(lines)-1 {
			parseInstructions(line)
			continue
		}
		gridLine := make([]string, maxLen)
		for i, ch := range line {
			//gridLine = append(gridLine, string(ch))
			gridLine[i] = string(ch)
		}
		grid = append(grid, gridLine)
	}

	return
}

func (grid grid) print() {
	for _, row := range grid {
		for _, col := range row {
			if col == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type position struct{ r, c int }

func (grid grid) move(instructions []instruction) (pos position, facing int) {
	pos = position{0, 0}
	facing = 90

	firstDot := func() position {
		switch facing {
		case 0:
			for i := len(grid) - 1; i >= 0; i-- {
				if grid[i][pos.c] != " " && grid[i][pos.c] != "" {
					return position{i, pos.c}
				}
			}
		case 90:
			for i := 0; i < len(grid[pos.r]); i++ {
				if grid[pos.r][i] != " " && grid[pos.r][i] != "" {
					return position{pos.r, i}
				}
			}
		case 180:
			for i := 0; i < len(grid); i++ {
				if grid[i][pos.c] != " " && grid[i][pos.c] != "" {
					return position{i, pos.c}
				}
			}
		case 270:
			for i := len(grid[pos.r]) - 1; i >= 0; i-- {
				if grid[pos.r][i] != " " && grid[pos.r][i] != "" {
					return position{pos.r, i}
				}
			}
		default:
			log.Fatalf("unkown facing: %v", facing)
		}
		return position{-1, -1}
	}

	drawPos := func() {
		switch facing {
		case 0:
			grid[pos.r][pos.c] = "^"
		case 90:
			grid[pos.r][pos.c] = ">"
		case 180:
			grid[pos.r][pos.c] = "v"
		case 270:
			grid[pos.r][pos.c] = "<"
		}
	}

	pos = firstDot()
	drawPos()

	for _, instruction := range instructions {
		for i := 0; i < instruction.move; i++ {
			formerPos := pos

			switch facing {
			case 0:
				pos.r--
			case 90:
				pos.c++
			case 180:
				pos.r++
			case 270:
				pos.c--
			default:
				log.Fatalf("unkown facing: %v", facing)
			}

			if pos.r < 0 || pos.r == len(grid) || pos.c < 0 || pos.c == len(grid[pos.r]) ||
				grid[pos.r][pos.c] == " " || grid[pos.r][pos.c] == "" {
				pos = firstDot()
			}

			if grid[pos.r][pos.c] == "#" {
				pos = formerPos
				drawPos()
				break
			}

			drawPos()
		}

		if instruction.turn == "" {
			continue
		}
		switch instruction.turn {
		case "R":
			facing += 90
			if facing == 360 {
				facing = 0
			}
		case "L":
			facing -= 90
			if facing < 0 {
				facing += 360
			}
		default:
			log.Fatalf("turn not understood: %v", instruction.turn)
		}

		drawPos()
		//grid.print()
	}

	return
}

func password(pos position, facing int) (pos2 position, facing2 int, sum int) {
	pos.r++
	pos.c++ //1 based indexing
	switch facing {
	case 0:
		facing = 3
	case 90:
		facing = 0
	case 180:
		facing = 1
	case 270:
		facing = 2
	}
	return pos, facing, 1000*(pos.r) + 4*(pos.c) + facing
}
