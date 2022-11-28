package main

import (
	"fmt"
	"math"
	"strconv"
	advent "temp/adventofcode/go"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func GammaRate(transposed [][]byte) (gamma, epsilon, power int64, err error) {
	strGamma := ""
	strEpsilon := ""

	for i, row := range transposed {
		most := make(map[byte]int)
		for _, col := range row {
			most[col]++
		}

		counts := maps.Values(most)
		slices.Sort(counts)

		var mostCommon, leastCommon byte
		for num, count := range most {
			if count == counts[len(counts)-1] {
				mostCommon = num
				break
			}
		}
		for num, count := range most {
			if count == counts[0] {
				leastCommon = num
				break
			}
		}

		fmt.Printf("%d most common: %s\n", i, string(mostCommon))
		strGamma += string(mostCommon)

		fmt.Printf("%d least common: %s\n", i, string(leastCommon))
		strEpsilon += string(leastCommon)
	}

	gamma, err = strconv.ParseInt(strGamma, 2, 64)
	if err != nil {
		return
	}
	fmt.Printf("Gamma rate: %d\n", gamma)

	epsilon, err = strconv.ParseInt(strEpsilon, 2, 64)
	if err != nil {
		return
	}
	fmt.Printf("Epsilon rate: %d\n", epsilon)

	power = gamma * epsilon
	fmt.Printf("Power: %d\n", power)

	return
}

func LifeSuport() (err error) {
	lines, err := advent.ReadInput(inputPath)
	if err != nil {
		return
	}

	transposed, err := Input()
	if err != nil {
		return
	}

	oxygen := lines
	carbon := lines

	for i, t := range transposed {
		var sum int
		for _, byt := range t {
			sum += int(byt)
		}
		avg := math.Round(float64(sum) / float64(len(t)))
		least := func() int {
			if avg == 48 {
				return 49
			} else {
				return 48
			}
		}()
		fmt.Printf("avg: %v, sum: %v, len: %v\n", avg, sum, len(t))

		if len(oxygen) > 1 {
			filtered := []string{}
			for _, line := range oxygen {
				if line[i] == byte(avg) {
					filtered = append(filtered, line)
				}
			}
			oxygen = filtered
		}

		if len(carbon) > 1 {
			filtered := []string{}
			for _, line := range carbon {
				if line[i] == byte(least) {
					filtered = append(filtered, line)
				}
			}
			carbon = filtered
		}
	}

	o, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		return
	}
	c, err := strconv.ParseInt(carbon[0], 2, 64)
	if err != nil {
		return
	}

	fmt.Printf("Oxygen: %s - %d\n", oxygen, o)
	fmt.Printf("CO2: %s - %d\n", carbon, c)

	return
}
