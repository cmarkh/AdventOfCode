package main

import (
	"fmt"
	"testing"
)

var testPath = "../test.txt"

func TestInput(t *testing.T) {
	ages, err := Input(testPath)
	if err != nil {
		return
	}
	//ages.Print()
	fmt.Println(ages)
}

func TestIncrementAges(t *testing.T) {
	ages, err := Input(testPath)
	if err != nil {
		return
	}

	days := 80
	ages = ages.Increment(days)
	//ages.Print()
	fmt.Printf("After %d turns, there are %d fish\n", days, ages.Count())

}
