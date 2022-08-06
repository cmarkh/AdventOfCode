package main

import (
	advent "temp/adventofcode/go"
	"testing"
)

func TestGamma(t *testing.T) {
	inputPath = advent.BasePath + "3/test.txt"
	transposed, err := Input()
	if err != nil {
		t.Fatal(err)
	}

	gamma, epsilon, power, err := GammaRate(transposed)
	if err != nil {
		t.Fatal(err)
	}

	if gamma != 22 {
		t.Fatalf("Incorrect gamma rate: %d, Should have been: 22\n", gamma)
	}
	if epsilon != 9 {
		t.Fatalf("Incorrect epsilon rate: %d, Should have been: 9\n", gamma)
	}
	if power != 198 {
		t.Fatalf("Incorrect power rate: %d, Should have been: 198\n", gamma)
	}
}

func TestLifeSupport(t *testing.T) {
	inputPath = advent.BasePath + "3/test.txt"

	err := LifeSuport()
	if err != nil {
		t.Fatal(err)
	}
}
