package main

import (
	"fmt"
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = "../test.txt"

func TestCommon(t *testing.T) {
	input, err := advent.ReadInput(testPath)
	if err != nil {
		t.Fatal(err)
	}

	most, least, err := Common(input, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("most: %d\nleast: %d\n", most, least)
}

func TestRates(t *testing.T) {
	input, err := advent.ReadInput(testPath)
	if err != nil {
		t.Fatal(err)
	}

	gamma, epsilon, err := Rates(input)
	if err != nil {
		t.Fatal(err)
	}

	product, err := MultiplyRates(gamma, epsilon)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("gamma: %s\nepsilon: %s\nmultiplied: %d\n", gamma, epsilon, product)
}

func TestLifeSupport(t *testing.T) {
	input, err := advent.ReadInput(testPath)
	if err != nil {
		t.Fatal(err)
	}

	oxygen, co2, product, err := LifeSupport(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("oxygen: %d\nco2: %d\nproduct: %d\n", oxygen, co2, product)
}
