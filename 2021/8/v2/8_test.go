package main

import (
	"fmt"
	"testing"
)

var testPath = "../test.txt"

func TestInput(t *testing.T) {
	entries, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("%+v\n", entry)
	}
}

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

func TestParseEntry(t *testing.T) {
	entries, _ := Input(testPath)

	translation, found := entries[0].ParseSignals()
	if !found {
		t.Fatal(fmt.Errorf("no translation found :("))
	}
	fmt.Println(translation)
}

func TestParseOutput(t *testing.T) {
	entries, _ := Input(testPath)

	translation, found := entries[0].ParseSignals()
	if !found {
		t.Fatal(fmt.Errorf("no translation found :("))
	}

	output, err := entries[0].ParseOutput(translation)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v - %s\n", entries[0].Outputs, output)
}

func TestTranslation(t *testing.T) {
	entries, _ := Input(testPath)

	translation, found := entries[0].ParseSignals()
	if !found {
		t.Fatal(fmt.Errorf("no translation found :("))
	}

	digit, err := translation.Translate("eafb")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(digit.Digit())
}

func TestRecursion(t *testing.T) {
	d2 := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	var combinations func(partial []string, row, col int)
	combinations = func(partial []string, row, col int) {
		partial[row] = d2[row][col]

		if row == len(d2)-1 {
			fmt.Println(partial)
		}
		if row < len(d2)-1 {
			combinations(partial, row+1, 0)
		}
		if col < len(d2[row])-1 {
			combinations(partial, row, col+1)
		}
	}

	combinations(make([]string, len(d2)), 0, 0)
}

func TestParseEntries(t *testing.T) {
	input, _ := Input(testPath)

	err := input.Parse()
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range input {
		fmt.Println(entry)
	}
}
