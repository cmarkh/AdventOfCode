package main

import (
	advent "temp/adventofcode/go"
	"testing"
)

func TestSmall1(t *testing.T) {
	inputPath = advent.BasePath + "12/test.txt"
	main()
}

func TestSmall2(t *testing.T) {
	inputPath = advent.BasePath + "12/test2.txt"
	main()
}

func TestSmall3(t *testing.T) {
	inputPath = advent.BasePath + "12/test3.txt"
	main()
}
