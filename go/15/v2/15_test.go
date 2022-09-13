package main

import (
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = advent.BasePath + "15/test.txt"
var testPath2 = advent.BasePath + "15/test2.txt"

func TestMain1(t *testing.T) {
	inputPath = testPath
	main()
}

func TestMain2(t *testing.T) {
	inputPath = testPath2
	main()
}
