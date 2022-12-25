package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	sum := sum(toDecimal(parse(input)))
	snafu := snafu(sum)
	fmt.Printf("sum: %v, snafu: %v\n", sum, snafu)
}

func parse(input string) (numbers []string) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line == "" {
			slices.Delete(lines, i, i+1)
		}
	}
	return lines
}

func decimal(snafu string) (decimal int) {
	var err error
	for i, ch := range snafu {
		var digit int
		switch ch {
		case '=':
			digit = -2
		case '-':
			digit = -1
		default:
			digit, err = strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(err)
			}
		}
		digit *= pow(5, len(snafu)-i-1)
		decimal += digit
	}
	return
}

func snafu(decimal int) (snafu string) {
	digits := []int{}
	for {
		digits = append([]int{decimal % 5}, digits...)
		decimal /= 5
		if decimal == 0 {
			break
		}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] > 2 {
			if i-1 < 0 {
				digits = append([]int{0}, digits...)
				i++
			}
			digits[i-1]++
			switch digits[i] - 5 {
			case 0:
				snafu = "0" + snafu
			case -1:
				snafu = "-" + snafu
			case -2:
				snafu = "=" + snafu
			}
			continue
		}
		snafu = strconv.Itoa(digits[i]) + snafu
	}

	return
}

func pow(base, exponent int) (result int) {
	result = 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return
}

func sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func toDecimal(snafus []string) (decimals []int) {
	for _, snafu := range snafus {
		decimals = append(decimals, decimal(snafu))
	}
	return
}
