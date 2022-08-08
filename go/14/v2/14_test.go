package main

import (
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = advent.BasePath + "14/test.txt"

func TestMain(m *testing.M) {
	inputPath = testPath
	main()
}
