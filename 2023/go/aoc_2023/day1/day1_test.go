package day1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	res, err := Part1(string(input))
	if err != nil {
		t.Fatal(err)
	}
	if res != 56049 {
		t.Logf("res: %d", res)
		t.Fatalf("Expected 56049, got %d", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	res, err := Part2(string(input))
	if err != nil {
		t.Fatal(err)
	}
	if res != 54530 {
		t.Logf("res: %d", res)
		t.Fatalf("Expected 54530, got %d", res)
	}
}

func TestEx1(t *testing.T) {
	input, err := os.ReadFile("ex1.txt")
	if err != nil {
		t.Fatal(err)
	}
	res, err := Part1(string(input))
	if err != nil {
		t.Fatal(err)
	}
	if res != 142 {
		t.Fatalf("Expected 142, got %d", res)
	}
}

func TestEx2(t *testing.T) {
	input, err := os.ReadFile("ex2.txt")
	if err != nil {
		t.Fatal(err)
	}
	res, err := Part2(string(input))
	if err != nil {
		t.Fatal(err)
	}
	if res != 281 {
		t.Fatalf("Expected 3, got %d", res)
	}
}

func TestGetNums(t *testing.T) {
	input, err := os.ReadFile("ex2.txt")
	if err != nil {
		t.Fatal(err)
	}
	nums, err := GetNums(string(input))
	if err != nil {
		t.Fatal(err)
	}
	for num := range nums {
		t.Logf("num: %v", num)
	}
}
