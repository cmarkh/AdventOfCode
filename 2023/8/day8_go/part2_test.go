package day8

import "testing"

func TestPart2Ex3(t *testing.T) {
	instructions, mapIn := input("ex3.txt")
	steps := Part2(instructions, mapIn)
	expected := 6
	if steps != uint64(expected) {
		t.Fatalf("expected %d, found %d", expected, steps)
	}
}

func TestPart2(t *testing.T) {
	instructions, mapIn := input("input.txt")
	steps := Part2(instructions, mapIn)
	expected := 13289612809129
	if steps != uint64(expected) {
		t.Fatalf("expected %d, found %d", expected, steps)
	}
}
