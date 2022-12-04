package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/slices"
	"golang.org/x/text/unicode/bidi"
)

var inputPath = "input.txt"

func main() {
	symbols, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	invalid, lines := symbols.Validate()
	/*for _, bracket := range invalid {
		fmt.Printf("%+v\n", bracket)
	}*/

	score, err := ScoreInvalid(invalid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("score: %d\n", score)

	symbols = symbols.RemoveInvalidLines(lines)

	completions := symbols.CompleteLines()

	scores, err := ScoreCompletions(completions)
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(scores)
	fmt.Printf("score: %d\n", scores[len(scores)/2])

}

type Symbols []Line
type Line []string

func Input(path string) (symbols Symbols, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		symbols = append(symbols, strings.Split(line, ""))
	}

	return
}

type Bracket struct {
	Symbol   string
	Position int
}

func (symbols Symbols) Validate() (invalid []string, lines []int) {
	for l, line := range symbols {
		/*for i, ch := range line {
			fmt.Printf("%s:%d ", ch, i)
		}
		fmt.Println()*/

		var closes []Bracket
		for i, ch := range line {
			if !IsOpeningBracket(ch) {
				closes = append(closes, Bracket{ch, i})
			}
		}

		for _, close := range closes {
			var testI int
			var skipOpens int
			for testI = close.Position - 1; ; testI -= 1 {
				if !IsOpeningBracket(line[testI]) {
					skipOpens++
					continue
				}
				if skipOpens == 0 {
					break
				}
				skipOpens--
			}
			//fmt.Printf("close: %s at %d, open: %s at %d\n", close.Symbol, close.Position, line[testI], testI)
			if ClosingBracket(line[testI]) != close.Symbol {
				invalid = append(invalid, close.Symbol)
				lines = append(lines, l)
				break
			}
		}
	}

	return
}

func (symbols Symbols) RemoveInvalidLines(invalid []int) (filtered Symbols) {
	for l, line := range symbols {
		if !slices.Contains(invalid, l) {
			filtered = append(filtered, line)
		}
	}
	return
}

type Center struct {
	Open, Close int
}

func (line Line) Centers() (centers []Center) {
	for i, ch := range line {
		if !IsOpeningBracket(ch) {
			if IsOpeningBracket(line[i-1]) {
				centers = append(centers, Center{i - 1, i})
			}
		}
	}
	return
}

func (symbols Symbols) CompleteLines() (completions [][]string) {
	for _, line := range symbols {
		newLine := []string{}

		var opens []Bracket
		for i, ch := range line {
			if IsOpeningBracket(ch) {
				opens = append(opens, Bracket{ch, i})
			}
		}

		for _, open := range opens {
			var testI int
			var skipCloses int
			for testI = open.Position + 1; ; testI += 1 {
				if testI >= len(line) {
					break
				}
				if IsOpeningBracket(line[testI]) {
					skipCloses++
					continue
				}
				if skipCloses == 0 {
					break
				}
				skipCloses--
			}

			if testI >= len(line) || OpeningBracket(line[testI]) != open.Symbol {
				newLine = append(newLine, ClosingBracket(open.Symbol))
			}
		}
		last := len(newLine) - 1
		for i := 0; i < len(newLine)/2; i++ {
			newLine[i], newLine[last-i] = newLine[last-i], newLine[i]
		}
		completions = append(completions, newLine)
	}
	return
}

func (symbols Symbols) Print() {
	for _, line := range symbols {
		fmt.Println(strings.Join(line, ""))
	}
}

func IsOpeningBracket(char string) bool {
	if char == "<" {
		return true
	}
	prop, _ := bidi.LookupString(char)
	return prop.IsOpeningBracket()
}

func ClosingBracket(char string) string {
	if char == "<" {
		return ">"
	}
	return bidi.ReverseString(char)
}

func OpeningBracket(char string) string {
	if char == ">" {
		return "<"
	}
	return bidi.ReverseString(char)
}

func ScoreInvalid(invalid []string) (score int, err error) {
	for _, ch := range invalid {
		switch ch {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		default:
			return 0, fmt.Errorf("unkown symbol: %s", ch)
		}
	}
	return
}

func ScoreCompletions(completions [][]string) (scores []int, err error) {
	scores = make([]int, len(completions))

	for i, line := range completions {
		for _, ch := range line {
			scores[i] *= 5
			switch ch {
			case ")":
				scores[i] += 1
			case "]":
				scores[i] += 2
			case "}":
				scores[i] += 3
			case ">":
				scores[i] += 4
			default:
				return scores, fmt.Errorf("unkown symbol: %s", ch)
			}
		}
	}
	return
}
