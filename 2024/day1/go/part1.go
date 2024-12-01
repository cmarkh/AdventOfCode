package main

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) (left, right []int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}

		split := strings.Split(line, " ")
		ln, err := strconv.ParseInt(strings.TrimSpace(split[0]), 10, 0)
		if err != nil {
			panic(err)
		}
		rs := ""
		for _, ch := range split[1:] {
			if ch == "" {
				continue
			}
			rs = ch
			break
		}
		rn, err := strconv.ParseInt(strings.TrimSpace(rs), 10, 0)
		if err != nil {
			panic(err)
		}

		left = append(left, int(ln))
		right = append(right, int(rn))
	}

	return left, right
}

func part1(left, right []int) int {
	count := 0

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		count += abs(left[i] - right[i])
	}

	return count
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
