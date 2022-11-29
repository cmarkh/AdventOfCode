package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func testLaunch(input string, velocity Velocity, t *testing.T) {
	target, err := ParseInput(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", target)

	hitTarget, _, location, path := Launch(target, velocity)
	path.Print(target)
	fmt.Println()
	fmt.Printf("hit target: %v ", hitTarget)
	if hitTarget {
		fmt.Printf("at %v\n", location)
	} else {
		fmt.Println()
	}
	fmt.Printf("max height: %v\n", path.maxHeight())
	fmt.Println()
}

func TestLaunch1(t *testing.T) {
	testLaunch("target area: x=20..30, y=-10..-5", Velocity{7, 2}, t)
}

func TestLaunch2(t *testing.T) {
	testLaunch("target area: x=20..30, y=-10..-5", Velocity{6, 3}, t)
}

func TestLaunch3(t *testing.T) {
	testLaunch("target area: x=20..30, y=-10..-5", Velocity{9, 0}, t)
}

func TestLaunch4(t *testing.T) {
	testLaunch("target area: x=20..30, y=-10..-5", Velocity{17, -4}, t)
}

func TestLaunch5(t *testing.T) {
	testLaunch("target area: x=150..171, y=-129..-70", Velocity{18, 63}, t)
}

func testMaxHeight(input string, t *testing.T) {
	target, err := ParseInput(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", target)

	maxH := MaxHeight(target)
	maxH.path.Print(target)
	fmt.Println()
	fmt.Printf("max height: %v\n", maxH.height)
	fmt.Printf("velocity: %v, %v\n", maxH.velocity.x, maxH.velocity.y)
	fmt.Printf("hit target at %v\n", maxH.hitLocation)
	fmt.Println()
}

func TestMaxHeight1(t *testing.T) {
	testMaxHeight("target area: x=20..30, y=-10..-5", t)
}

func TestMaxHeight2(t *testing.T) {
	for i := 7; i < 1000; i += rand.Intn(10) {
		target, err := ParseInput(fmt.Sprintf("target area: x=20..30, y=-%v..-5", i))
		if err != nil {
			t.Fatal(err)
		}
		maxH := MaxHeightOLD(target)
		fmt.Printf("x velocity: %v, y velocity: %v, tgt yMin: %v\tdiff: %v\n",
			maxH.velocity.x, maxH.velocity.y, target.yMin, target.yMin+maxH.velocity.y)
	}
}

func testValidShots(input string, t *testing.T) {
	target, err := ParseInput(input)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", target)

	velocities, count := AllValidShots(target)
	fmt.Println("Valid velocities:")
	for _, v := range velocities {
		fmt.Println(v)
	}
	fmt.Printf("\nNumber of valid velocities: %v\n", count)
}

func TestValidShots1(t *testing.T) {
	testValidShots("target area: x=20..30, y=-10..-5", t)
}
