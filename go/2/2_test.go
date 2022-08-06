package advent2

import "testing"

var testDirections []Instruction = []Instruction{
	{Direction: "forward", Amount: 5},
	{Direction: "down", Amount: 5},
	{Direction: "forward", Amount: 8},
	{Direction: "up", Amount: 3},
	{Direction: "down", Amount: 8},
	{Direction: "forward", Amount: 2},
}

func TestTravel(t *testing.T) {
	distance, err := Travel(testDirections)
	if err != nil {
		t.Fatal(err)
	}

	if distance != 150 {
		t.Fatalf("Incorrect distance: %d, Should have been: 150\n", distance)
	}
}

func TestAim(t *testing.T) {
	distance, err := Aim(testDirections)
	if err != nil {
		t.Fatal(err)
	}

	if distance != 900 {
		t.Fatalf("Incorrect distance: %d, Should have been: 900\n", distance)
	}
}
