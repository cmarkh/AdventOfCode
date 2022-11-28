package main

import (
	"testing"
)

var testPath = "../test.txt"
var testPath2 = "../test2.txt"

func TestMain1(t *testing.T) {
	inputPath = testPath
	main()
}

func TestMain2(t *testing.T) {
	inputPath = testPath2
	main()
}
