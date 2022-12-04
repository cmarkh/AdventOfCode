package main

import (
	"fmt"
	"log"
	"strconv"
	advent "temp/adventofcode/go/2021"
)

func main() {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	enhancer, input := parse(lines)
	fmt.Printf("enchancer: %v\n", enhancer)
	fmt.Printf("\ninput:\n")
	//input.Print()
	fmt.Println()

	iter := func(n int) {
		output := input
		for i := 0; i < n; i++ {
			output = enhance(output, enhancer, i)
		}
		output.Print()
		fmt.Printf("number of lit pixels: %v\n\n", countLit(output))
		fmt.Println()
	}

	fmt.Println("Part 1:")
	iter(2)

	fmt.Println("Part 2:")
	iter(50)

}

type image [][]int

func parse(lines []string) (enhancer []int, input image) {
	for i, line := range lines {
		if i == 0 {
			enhancer = decode(line)
			continue
		}
		if line == "" {
			continue
		}
		input = append(input, decode(line))
	}
	return
}

func decode(line string) (decoded []int) {
	for _, ch := range line {
		switch ch {
		case '.':
			decoded = append(decoded, 0)
		case '#':
			decoded = append(decoded, 1)
		default:
			log.Panicf("unkown symbol: %v\n", ch)
		}
	}
	return
}

func (image image) Print() {
	for _, line := range image {
		for _, ch := range line {
			switch ch {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func enhance(input image, enhancer []int, step int) (output image) {
	infinite := func(r, c int) string {
		var infinite string
		if enhancer[0] == 0 {
			infinite = "0"
		} else if step%2 == 0 {
			infinite = "0"
		} else {
			infinite = "1"
		}
		if r < 0 || r >= len(input) {
			return infinite
		}
		if c < 0 || c >= len(input[0]) {
			return infinite
		}
		return fmt.Sprint(input[r][c])
	}

	border := 1
	for r := -border; r < len(input)+border; r++ {
		outLine := []int{}
		for c := -border; c < len(input[0])+border; c++ {
			binary := ""

			binary += infinite(r-1, c-1)
			binary += infinite(r-1, c)
			binary += infinite(r-1, c+1)

			binary += infinite(r, c-1)
			binary += infinite(r, c)
			binary += infinite(r, c+1)

			binary += infinite(r+1, c-1)
			binary += infinite(r+1, c)
			binary += infinite(r+1, c+1)

			decimal, err := strconv.ParseInt(binary, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			outLine = append(outLine, enhancer[decimal])

		}
		output = append(output, outLine)
	}

	return
}

func countLit(im image) (count int) {
	for _, row := range im {
		for _, col := range row {
			if col == 1 {
				count++
			}
		}
	}
	return
}
