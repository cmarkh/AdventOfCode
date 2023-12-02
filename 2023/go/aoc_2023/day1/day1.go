package day1

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) (sum int, err error) {
	for _, line := range strings.Split(input, "\n") {
		strNum := ""
		for _, char := range line {
			if unicode.IsNumber(char) {
				strNum += string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			if unicode.IsNumber(char) {
				strNum += string(char)
				break
			}
		}

		num, err := strconv.Atoi(strNum)
		if err != nil {
			return 0, err
		}
		sum += num
	}

	return
}

var NUMS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Part2(input string) (sum int, err error) {
	for _, line := range strings.Split(input, "\n") {
		nums, err := GetNums(line)
		if err != nil {
			return 0, err
		}

		strNum := fmt.Sprint(nums[0].num) + fmt.Sprint(nums[len(nums)-1].num)
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return
}

type numPosition struct {
	idx int
	num int
}

func GetNums(line string) (nums []numPosition, err error) {
	for i, char := range line {
		if unicode.IsNumber(char) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nums, err
			}
			nums = append(nums, numPosition{i, num})
		}
	}

	for str, num := range NUMS {
		idx := strings.Index(line, str)
		if idx != -1 {
			nums = append(nums, numPosition{idx, num})
		}
	}
	for str, num := range NUMS {
		idx := strings.LastIndex(line, str)
		if idx != -1 {
			nums = append(nums, numPosition{idx, num})
		}
	}

	slices.SortFunc(nums, func(a, b numPosition) int {
		return cmp.Compare(a.idx, b.idx)
	})

	return
}
