package advent1

import (
	"testing"
)

var testDepths = []int64{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestCountDepthIncreases(t *testing.T) {
	increases, err := CountDepthIncreases(testDepths)
	if err != nil {
		t.Fatal(err)
	}
	if increases != 7 {
		t.Fatalf("found incorrect number of increases: %d, should be 7", increases)
	}
}

/*
func ExportedCountDepthIncreases(depths []int64) (int64, error) {
	return countDepthIncreases(depths)
}
*/

func TestSlidingWindow(t *testing.T) {
	increases, err := SlidingWindow(testDepths)
	if err != nil {
		t.Fatal(err)
	}
	if increases != 5 {
		t.Fatalf("found incorrect number of increases: %d, should be 5", increases)
	}
}
