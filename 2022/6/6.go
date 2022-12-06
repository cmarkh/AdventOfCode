package main

import (
	_ "embed"
	"fmt"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:")
	marker := startOfPacket(input)
	fmt.Printf("start of packet marker at %v character\n", marker)
	fmt.Println()

	fmt.Println("Part 2:")
	marker = startOfMessage(input)
	fmt.Printf("start of message marker at %v character\n", marker)
	fmt.Println()
}

func startOfPacket(input string) (mark int) {
	return marker(input, 4)
}

func startOfMessage(input string) (mark int) {
	return marker(input, 14)
}

func marker(input string, length int) (mark int) {
	mark = -1
	start := 0
	soFar := []rune{}
	for start < len(input) {
		for ch := start; ch < len(input); ch++ {
			if slices.Contains(soFar, rune(input[ch])) {
				soFar = []rune{}
				start++
				break
			}
			soFar = append(soFar, rune(input[ch]))
			if len(soFar) >= length {
				return start + length
			}
		}
	}
	return
}
