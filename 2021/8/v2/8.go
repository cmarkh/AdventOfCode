package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/slices"
)

var inputPath = "../input.txt"

func main() {
	entries, _ := Input(inputPath)

	err := entries.Parse()
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, entry := range entries {
		v, err := strconv.Atoi(entry.Value)
		if err != nil {
			log.Fatal(err)
		}
		sum += v
		fmt.Println(entry)
	}
	fmt.Printf("sum of values: %d\n", sum)
}

var (
	zero  = Digit{1, 1, 1, 0, 1, 1, 1}
	one   = Digit{0, 0, 1, 0, 0, 1, 0}
	two   = Digit{1, 0, 1, 1, 1, 0, 1}
	three = Digit{1, 0, 1, 1, 0, 1, 1}
	four  = Digit{0, 1, 1, 1, 0, 1, 0}
	five  = Digit{1, 1, 0, 1, 0, 1, 1}
	six   = Digit{1, 1, 0, 1, 1, 1, 1}
	seven = Digit{1, 0, 1, 0, 0, 1, 0}
	eight = Digit{1, 1, 1, 1, 1, 1, 1}
	nine  = Digit{1, 1, 1, 1, 0, 1, 1}

	digits = map[int]Digit{
		0: {1, 1, 1, 0, 1, 1, 1},
		1: {0, 0, 1, 0, 0, 1, 0},
		2: {1, 0, 1, 1, 1, 0, 1},
		3: {1, 0, 1, 1, 0, 1, 1},
		4: {0, 1, 1, 1, 0, 1, 0},
		5: {1, 1, 0, 1, 0, 1, 1},
		6: {1, 1, 0, 1, 1, 1, 1},
		7: {1, 0, 1, 0, 0, 1, 0},
		8: {1, 1, 1, 1, 1, 1, 1},
		9: {1, 1, 1, 1, 0, 1, 1},
	}
)

type Entries []Entry
type Entry struct {
	Signals []string
	Outputs []string
	Value   string
}

type Display [4]Digit
type Digit [7]int

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

func (entry Entry) ParseSignals() (t Translation, found bool) {
	possibilities := make(Possibilities)
	for i := 0; i < 7; i++ {
		possibilities[i] = []string{"a", "b", "c", "d", "e", "f", "g"}
	}

	for _, code := range entry.Signals {
		possibilities.FilterKnownDigits(code)
	}

	translations := possibilities.PossibleTranslations()

	filtered := []Translation{}
	for _, t := range translations {
		if t.IsValid(entry.Signals) {
			filtered = append(filtered, t)
		}
	}

	/*for _, t := range filtered {
		fmt.Printf("possible: %v\n", t)
	}*/

	if len(filtered) == 1 {
		found = true
		t = filtered[0]
	}

	return
}

func (entries Entries) Parse() (err error) {
	for i, entry := range entries {
		t, found := entry.ParseSignals()
		if !found {
			err = fmt.Errorf("no translation found for: %v", entry)
			return
		}

		value, err := entry.ParseOutput(t)
		if err != nil {
			return err
		}
		entries[i].Value = value
	}
	return
}

func (entry Entry) ParseOutput(t Translation) (output string, err error) {
	for _, code := range entry.Outputs {
		digit, err := t.Translate(code)
		if err != nil {
			return "", err
		}
		//fmt.Println(digit.Digit())
		output += fmt.Sprint(digit.Digit())
	}
	return
}

type Possibilities map[int][]string

func (possibilities Possibilities) FilterKnownDigits(code string) {
	digit, err := WhichDigit(code)
	if err != nil {
		return
	}

	for i, on := range digits[digit] {
		filtered := []string{}
		for _, ch := range possibilities[i] {
			if (on == 1 && strings.Contains(code, ch)) || (on == 0 && !strings.Contains(code, ch)) {
				filtered = append(filtered, ch)
			}
		}
		possibilities[i] = filtered

		//fmt.Println(possibilities)
	}
}

type Translation map[string]int //given char, translated position

func (t Translation) Translate(code string) (d Digit, err error) {
	for _, ch := range code {
		d[t[string(ch)]] = 1
	}
	if d.Digit() == -1 {
		return d, fmt.Errorf("%v is not a valid digit. code: %s", d, code)
	}
	return
}

func (p Possibilities) FirstTranslation() (t Translation) {
	t = make(Translation)
	var used []string
	for i, pch := range p {
		for _, ch := range pch {
			if !slices.Contains(used, ch) {
				t[string(ch)] = i
				used = append(used, ch)
				break
			}
		}
	}
	return
}

func (p Possibilities) PossibleTranslations() (translations []Translation) {
	var combos [][]string
	var combinations func(partial []string, row, col int)
	combinations = func(partial []string, row, col int) {
		partial[row] = p[row][col]
		if row == len(p)-1 {
			var newcombo []string = make([]string, len(partial))
			copy(newcombo, partial)
			combos = append(combos, newcombo)
			//fmt.Println(partial)
		}
		if row < len(p)-1 {
			combinations(partial, row+1, 0)
		}
		if col < len(p[row])-1 {
			combinations(partial, row, col+1)
		}
	}

	combinations(make([]string, len(p)), 0, 0)

	norepeats := func(combo []string) bool {
		used := []string{}
		for _, ch := range combo {
			if slices.Contains(used, ch) {
				return false
			}
			used = append(used, ch)
		}
		return true
	}

	valid := [][]string{}
	for _, combo := range combos {
		if norepeats(combo) {
			valid = append(valid, combo)
		}
	}

	for _, combo := range valid {
		t := Translation{}
		for i, ch := range combo {
			t[ch] = i
		}
		translations = append(translations, t)
	}

	return
}

func (d Digit) Digit() int {
	for i, dig := range digits {
		if dig == d {
			return i
		}
	}
	return -1
}

func (t Translation) IsValid(codes []string) bool {
	for _, code := range codes {
		if _, err := t.Translate(code); err != nil {
			return false
		}
	}
	return true
}
