package main

import (
	"fmt"
	"testing"
)

func TestPart2Ex1(t *testing.T) {
	left, right := ex1(t)
	output := part2(left, right)
	fmt.Println(output)
	if output != 31 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	left, right := input(t)
	output := part2(left, right)
	fmt.Println(output)
	if output != 21306195 {
		t.Fail()
	}
}
