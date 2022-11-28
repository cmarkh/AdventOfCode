package main

import (
	"fmt"
	"testing"
)

var testPath = "../test.txt"

func TestMain(t *testing.T) {
	inputPath = testPath
	main()
}

func TestPairs(t *testing.T) {
	inputPath = testPath

	_, rules, err := input()
	if err != nil {
		t.Fatal(err)
	}

	for pair, insert := range rules {
		fmt.Printf("%s -> %s\n", pair, insert)
	}
	fmt.Println()

	poly := "NCB"
	for i := 1; i <= 5; i++ {
		poly, err = GrowPolymer(rules, poly)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%d: ", i)
		for n, ch := range poly {
			if n%3 == 0 {
				fmt.Print(" ")
			}
			fmt.Print(string(ch))
		}
		fmt.Println()
	}

	fmt.Println()
}
