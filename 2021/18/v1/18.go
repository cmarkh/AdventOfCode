package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
)

func main() {
	lines, err := advent.ReadInput("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}

type Pair struct {
	childX, childY *Pair
	value          int //note if has value, should not have child pairs
	text           string
}

func ParseInput(input []string) (outermost []Pair) {
	for _, line := range input {
		outermost = append(outermost, buildPair(line))
		break //just focus on line 0 for now
	}

	return
}

func buildPair(str string) (pair Pair) {
	pair.text = str
	//fmt.Printf("str: %s\tcaller:%s\n", str)

	if !strings.Contains(str, ",") {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		pair.value = val
		return
	}

	openBraceCount := 0
	for i, ch := range str {
		if ch == '[' {
			openBraceCount++
			continue
		}
		if ch == ']' {
			openBraceCount--
		}
		if ch == ',' && openBraceCount == 1 {
			x := buildPair(str[1:i])              //first half of string less the opening [
			y := buildPair(str[i+1 : len(str)-1]) //second half of string less the splitting comma and closing ]
			pair.childX = &x
			pair.childY = &y
		}
	}
	return
}

// START by sending nestlevels = 1
func (pair Pair) explode(nestLevels int) {
	if pair.value != 0 {
		fmt.Printf("done: %+v\n", pair.text)
		return
	}

	if nestLevels == 4 {
		fmt.Printf("kaboom: %s\n", pair.text)
		return
	}

	pair.childX.explode(nestLevels + 1)
	pair.childY.explode(nestLevels + 1)
}

func (pair Pair) Print() {
	if pair.value != 0 {
		fmt.Print(pair.value)
		return
	}
	fmt.Print("[")
	pair.childX.Print()
	fmt.Print(",")
	pair.childY.Print()
	fmt.Print("]")
}

func Add(input []string) {
	var previousSum string
	for i, line := range input {
		if i == 0 {
			previousSum = line
			continue
		}
		newPair := "[" + previousSum + "," + line + "]"
		fmt.Println(newPair)
	}
}

func reduce(sum string) {
	for {

		return
	}
}
