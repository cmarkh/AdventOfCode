package main

import (
	"fmt"
	"os"
	"testing"
)

func ex1(t *testing.T) (left, right []int) {
	input, err := os.ReadFile("ex1.txt")
	if err != nil {
		t.Fatal(err)
	}
	left, right = parseInput(string(input))
	return left, right
}

func input(t *testing.T) (left, right []int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	left, right = parseInput(string(input))
	return left, right
}

func TestParseInput(t *testing.T) {
	left, right := ex1(t)
	fmt.Println(left)
	fmt.Println(right)
}

func TestPart1Ex1(t *testing.T) {
	left, right := ex1(t)
	output := part1(left, right)
	fmt.Println(output)
	if output != 11 {
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	left, right := input(t)
	output := part1(left, right)
	fmt.Println(output)
	if output != 1651298 {
		t.Fail()
	}
}
