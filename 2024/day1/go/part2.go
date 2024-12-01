package main

func part2(left, right []int) int {
	similarity := 0

	right_map := make(map[int]int)
	for _, r := range right {
		right_map[r] += 1
	}

	for _, l := range left {
		similarity += right_map[l] * l
	}

	return similarity
}
