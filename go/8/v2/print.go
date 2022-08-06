package main

import "fmt"

func (d Digit) String() (lines [7][6]string) {
	t := "."
	if d[0] == 1 {
		t = "a"
	}
	lines[0][1] = t
	lines[0][2] = t
	lines[0][3] = t
	lines[0][4] = t

	t = "."
	if d[1] == 1 {
		t = "b"
	}
	lines[1][0] = t
	lines[2][0] = t

	t = "."
	if d[2] == 1 {
		t = "c"
	}
	lines[1][5] = t
	lines[2][5] = t

	t = "."
	if d[3] == 1 {
		t = "d"
	}
	lines[3][1] = t
	lines[3][2] = t
	lines[3][3] = t
	lines[3][4] = t

	t = "."
	if d[4] == 1 {
		t = "e"
	}
	lines[4][0] = t
	lines[5][0] = t

	t = "."
	if d[5] == 1 {
		t = "f"
	}
	lines[4][5] = t
	lines[5][5] = t

	t = "."
	if d[6] == 1 {
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
