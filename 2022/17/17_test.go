package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPart1(t *testing.T) {
	jets := parse(test1)
	g, height := fall(jets, 2022)
	g.print()
	fmt.Printf("height: %v\n", height)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	jets := parse(test1)
	height := fall2(jets, 1000000000000)
	//height := fall5(jets, 2022)
	fmt.Printf("height: %v\n", height)
	fmt.Println()
}

func TestHitSomething(t *testing.T) {
	g := grid{
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
		{"", "x", "", "", "", "", ""},
		{"", "", "", "", "", "", ""},
	}
	fmt.Println(g.rockHitSomething(rockShapes[3], 5, 2))
}

func TestFall21(t *testing.T) {
	jets := parse("<<<<<<<<>>>><<<<><><><><>")
	height := fall2(jets, 2022)
	//height := fall5(jets, 2022)
	fmt.Printf("height: %v\n", height)
	fmt.Println()

	if height != 3437 {
		t.Fail()
	}
}

func TestFall22(t *testing.T) {
	jets := parse("<<<<>>>>>>>>>>>>>>>>><<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>")
	height := fall2(jets, 2022)
	//height := fall5(jets, 2022)
	fmt.Printf("height: %v\n", height)
	fmt.Println()

	if height != 2641 {
		t.Fail()
	}
}
