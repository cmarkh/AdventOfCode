package main

import (
	"fmt"
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = advent.BasePath + "10/test.txt"

func TestInput(t *testing.T) {
	symbols, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	symbols.Print()
}

func TestValidate(t *testing.T) {
	symbols, _ := Input(testPath)

	invalid, _ := symbols.Validate()
	for _, bracket := range invalid {
		fmt.Printf("%+v\n", bracket)
	}

	score, err := ScoreInvalid(invalid)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("score: %d\n", score)
}

func TestMain(m *testing.M) {
	inputPath = testPath
	main()
}
