package main

import (
	"fmt"
	"testing"
)

var testPath = "../test.txt"

func TestDigitPrint(t *testing.T) {
	fmt.Println("\n0:")
	zero.Print()

	fmt.Println("\n1:")
	one.Print()

	fmt.Println("\n2:")
	two.Print()

	fmt.Println("\n3:")
	three.Print()

	fmt.Println("\n4:")
	four.Print()

	fmt.Println("\n5:")
	five.Print()

	fmt.Println("\n6:")
	six.Print()

	fmt.Println("\n7:")
	seven.Print()

	fmt.Println("\n8:")
	eight.Print()

	fmt.Println("\n9:")
	nine.Print()
}

func TestDisplayPrint(t *testing.T) {
	display := Display{zero, one, two, three}
	display.Print()
}

func TestInput(t *testing.T) {
	entries, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

func TestCountDigits(t *testing.T) {
	entries, _ := Input(testPath)

	count := entries.CountDigits()
	fmt.Printf("Number of 1,4,7,8s: %d\n", count)
}

func TestOutputValues(t *testing.T) {
	entries, _ := Input(testPath)

	err := entries.OutputValues()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", entries[0])
}
