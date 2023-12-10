package day5

import (
	"fmt"
	"os"

	"testing"
)

func TestParseAlmanac(t *testing.T) {
	input, err := os.ReadFile("ex1.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	almanac := ParseAlmanac(string(input))
	fmt.Println(almanac)
}

func TestPart2Ex1(t *testing.T) {
	input, err := os.ReadFile("ex1.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	almanac := ParseAlmanac(string(input))
	seeds, err := ParseSeeds(string(input))
	if err != nil {
		t.Fatalf("Failed to parse seeds: %v", err)
	}

	res := Part2(*seeds, almanac)
	if res != 46 {
		t.Fatalf("Expected 46, got %d", res)
	}
}

func TestPart2(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	almanac := ParseAlmanac(string(input))
	seeds, err := ParseSeeds(string(input))
	if err != nil {
		t.Fatalf("Failed to parse seeds: %v", err)
	}

	res := Part2(*seeds, almanac)
	if res != 5200543 {
		t.Fatalf("Expected 5200543, got %d", res)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		b.Fatalf("Failed to read file: %v", err)
	}

	almanac := ParseAlmanac(string(input))
	seeds, err := ParseSeeds(string(input))
	if err != nil {
		b.Fatalf("Failed to parse seeds: %v", err)
	}

	for n := 0; n < b.N; n++ {
		Part2(*seeds, almanac)
	}
}
