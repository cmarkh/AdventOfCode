package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

// Notes on input data
// there are 14 inp w (one for each digit of the model number) - I will split the instructions into 14 "sets"
// ergo w is overwritten on each digit
// x and y are also overwritten - mul x 0 and mul y 0 are always called before first use
// z does appear to be dependent on the value from prior instruction sets

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	t := time.Now()

	zs := zsNeeded()

	fmt.Println("Part 1:")
	highest := highestMONAD(zs)
	fmt.Printf("highest MONAD: %s\n", highest)
	fmt.Println()

	fmt.Println("Part 2:")
	lowest := lowestMONAD(zs)
	fmt.Printf("lowest MONAD: %s\n", lowest)
	fmt.Println()

	fmt.Printf("I'm not getting any younger: %v\n", time.Since(t))
}

type ZInNeeded [15][]int //map[step][digit]zNeeded

// calculates what the input z needs to be for each instruction set, used later to determine valid digits for the set prior
func zsNeeded() (zInNeeded ZInNeeded) {
	zInNeeded[14] = []int{0}

	//sets := splitInstructions(sanitize(input))

	for s := 13; s > 0; s-- {
		for d := 1; d <= 9; d++ {
			for z := -1000; z < 200000; z++ {
				//zOut, _ := zValue(sets[s], d, z)
				zOut := zValueFormulas(s, d, z)
				if slices.Contains(zInNeeded[s+1], zOut) {
					zInNeeded[s] = append(zInNeeded[s], z)
				}
			}
		}
	}

	return
}

func highestMONAD(zInNeeded ZInNeeded) (highest string) {
	zIn := 0
	for s := 0; s < 14; s++ {
		for d := 9; d >= 1; d-- {
			zOut := zValueFormulas(s, d, zIn)
			if slices.Contains(zInNeeded[s+1], zOut) {
				highest += fmt.Sprint(d)
				zIn = zOut
				break
			}
		}
	}
	return
}

func lowestMONAD(zInNeeded ZInNeeded) (lowest string) {
	zIn := 0
	for s := 0; s < 14; s++ {
		for d := 1; d <= 9; d++ {
			zOut := zValueFormulas(s, d, zIn)
			if slices.Contains(zInNeeded[s+1], zOut) {
				lowest += fmt.Sprint(d)
				zIn = zOut
				break
			}
		}
	}
	return
}

type instruction [3]string
type instructionSets [14][]instruction

func sanitize(input string) (instructions []instruction) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		switch len(split) {
		case 2:
			instructions = append(instructions, [3]string{split[0], split[1]})
		case 3:
			instructions = append(instructions, [3]string{split[0], split[1], split[2]})
		default:
			log.Fatalf("instructions not understood: %v", line)
		}
	}
	return
}

func splitInstructions(instructions []instruction) (sets instructionSets) {
	s := -1
	for i := 0; i < len(instructions); i++ {
		if instructions[i][0] == "inp" {
			s++
		}
		sets[s] = append(sets[s], instructions[i])
	}

	return
}

type alu map[string]int //map[variable]value

func zValue(set []instruction, modelDigit int, inputZ int) (z int, valid bool) {
	vars := make(alu)
	vars["z"] = inputZ

	for _, instruction := range set {
		var b int
		var err error
		b, err = strconv.Atoi(instruction[2])
		if err != nil {
			b = vars[instruction[2]]
		}

		switch instruction[0] {
		case "inp":
			vars[instruction[1]] = modelDigit
		case "add":
			vars[instruction[1]] += b
		case "mul":
			vars[instruction[1]] *= b
		case "div":
			if b == 0 {
				return 0, false
			}
			var quotient float64 = float64(vars[instruction[1]]) / float64(b)
			vars[instruction[1]] = int(math.Floor(quotient))
		case "mod":
			if vars[instruction[1]] < 0 || b <= 0 {
				return 0, false
			}
			vars[instruction[1]] %= b
		case "eql":
			if vars[instruction[1]] == b {
				vars[instruction[1]] = 1
			} else {
				vars[instruction[1]] = 0
			}
		default:
			log.Fatalf("instruction not understood: %+v", instruction)
		}
	}

	return vars["z"], true
}

// in each step only 3 variables change
// x = z % 26 + a != digit
// z = (z/b)(25(x) + 1) + (digit + c)(x)
type instructionVars struct {
	a, b, c int
}

var stepVars = map[int]instructionVars{
	0:  {12, 1, 1},
	1:  {12, 1, 1},
	2:  {15, 1, 16},
	3:  {-8, 26, 5},
	4:  {-4, 26, 9},
	5:  {15, 1, 3},
	6:  {14, 1, 2},
	7:  {14, 1, 15},
	8:  {-13, 26, 5},
	9:  {-3, 26, 11},
	10: {-7, 26, 7},
	11: {10, 1, 1},
	12: {-6, 26, 10},
	13: {-8, 26, 3},
}

func zValueFormulas(set int, digit int, inputZ int) (z int) {
	x := func() int {
		if inputZ%26+stepVars[set].a != digit {
			return 1
		}
		return 0
	}()
	return (inputZ/stepVars[set].b)*(25*x+1) + (digit+stepVars[set].c)*x
}
