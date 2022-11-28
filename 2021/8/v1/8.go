package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = "../input.txt"

func main() {
	Part1()
	Part2()
}

func Part1() {
	entries, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	count := entries.CountDigits()
	fmt.Printf("Number of 1,4,7,8s: %d\n", count)
}

type Display [4]Digit

type Digit struct {
	A, B, C, D, E, F, G bool
}

var (
	zero  = Digit{A: true, B: true, C: true, E: true, F: true, G: true}
	one   = Digit{C: true, F: true}
	two   = Digit{A: true, C: true, D: true, E: true, G: true}
	three = Digit{A: true, C: true, D: true, F: true, G: true}
	four  = Digit{B: true, C: true, D: true, F: true}
	five  = Digit{A: true, B: true, D: true, F: true, G: true}
	six   = Digit{A: true, B: true, D: true, E: true, F: true, G: true}
	seven = Digit{A: true, C: true, F: true}
	eight = Digit{A: true, B: true, C: true, D: true, E: true, F: true, G: true}
	nine  = Digit{A: true, B: true, C: true, D: true, F: true, G: true}
)

func (d Digit) String() (lines [7][6]string) {
	t := "."
	if d.A {

		t = "a"
	}
	lines[0][1] = t
	lines[0][2] = t
	lines[0][3] = t
	lines[0][4] = t

	t = "."
	if d.B {
		t = "b"
	}
	lines[1][0] = t
	lines[2][0] = t

	t = "."
	if d.C {
		t = "c"
	}
	lines[1][5] = t
	lines[2][5] = t

	t = "."
	if d.D {
		t = "d"
	}
	lines[3][1] = t
	lines[3][2] = t
	lines[3][3] = t
	lines[3][4] = t

	t = "."
	if d.E {
		t = "e"
	}
	lines[4][0] = t
	lines[5][0] = t

	t = "."
	if d.F {
		t = "f"
	}
	lines[4][5] = t
	lines[5][5] = t

	t = "."
	if d.G {
		t = "g"
	}
	lines[6][1] = t
	lines[6][2] = t
	lines[6][3] = t
	lines[6][4] = t

	return
}

func (d Digit) Print() {
	lines := d.String()

	for _, line := range lines {
		for _, c := range line {
			if c == "" {
				fmt.Print(" ")
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (display Display) Print() {
	//digitWidth := 8
	//var lines [8 * 4][7 * 4]string
	var lines [7][]string

	for i, digit := range display {
		digiLines := digit.String()
		for l, line := range digiLines {
			if i != 0 {
				lines[l] = append(lines[l], "  ")
			}
			lines[l] = append(lines[l], line[:]...)
		}
	}

	for _, line := range lines {
		for _, c := range line {
			if c == "" {
				fmt.Print(" ")
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func WhichDigit(code string) (digit int, err error) {
	switch len(code) {
	case 2:
		return 1, nil
	case 4:
		return 4, nil
	case 3:
		return 7, nil
	case 7:
		return 8, nil
	default:
		return 0, fmt.Errorf("unkown digit: %s", code)
	}
}

type Entries []Entry
type Entry struct {
	Signals []string
	Outputs []string
	Value   string
}

func Input(path string) (entries Entries, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		signals, outputs, found := strings.Cut(line, " | ")
		if !found {
			err = fmt.Errorf("input not recognized")
			return
		}
		entries = append(entries, Entry{
			Signals: strings.Split(signals, " "),
			Outputs: strings.Split(outputs, " "),
		})
	}
	return
}

func (entries Entries) CountDigits() (count int) {
	for _, e := range entries {
		for _, code := range e.Outputs {
			if d, err := WhichDigit(code); err == nil {
				switch d {
				case 1, 4, 7, 8:
					//fmt.Printf("%s: %d\n", code, d)
					count++
				}
			}
		}
	}
	return
}
