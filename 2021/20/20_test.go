package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func TestParse(t *testing.T) {
	enhancer, input := parse(strings.Split(test1, "\n"))
	fmt.Printf("enchancer: %v\n", enhancer)
	fmt.Printf("\ninput:\n")
	input.Print()
}

func TestEnhance1(t *testing.T) {
	enhancer, input := parse(strings.Split(test1, "\n"))

	output := enhance(input, enhancer, 0)
	output.Print()

	fmt.Println()
}

func TestEnhance2(t *testing.T) {
	enhancer, input := parse(strings.Split(test1, "\n"))

	output := enhance(input, enhancer, 0)
	output = enhance(output, enhancer, 0)
	fmt.Println("output:")
	output.Print()

	fmt.Println()
}

func TestCount(t *testing.T) {
	enhancer, input := parse(strings.Split(test1, "\n"))

	output := enhance(input, enhancer, 0)
	fmt.Println("output:")
	output.Print()
	fmt.Printf("number of lit pixels: %v\n\n", countLit(output))

	output = enhance(output, enhancer, 1)
	fmt.Println("output:")
	output.Print()
	fmt.Printf("number of lit pixels: %v\n\n", countLit(output))

	fmt.Println()
}

func TestCount2(t *testing.T) {
	enhancer, input := parse(strings.Split(test1, "\n"))

	iter := func(n int) {
		output := input
		for i := 0; i < n; i++ {
			output = enhance(output, enhancer, i)
		}
		fmt.Println("checked border:")
		output.Print()
		fmt.Printf("number of lit pixels: %v\n\n", countLit(output))
		fmt.Println()
	}

	iter(50)
}
