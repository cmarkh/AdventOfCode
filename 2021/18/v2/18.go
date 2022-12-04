package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
	"unicode"
)

func main() {
	lines, err := advent.ReadInput("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(lines)

	fmt.Println("Part 1:")
	answer := sumList(lines)
	fmt.Printf("answer: %s\n", answer)

	mag := magnitude(answer)
	fmt.Printf("magnitude: %v\n", mag)

	fmt.Println()

	fmt.Println("Part 2:")
	pair1, pair2, mag := greatestMagnitude(lines)
	fmt.Println("Greatest Magnitude is between:")
	fmt.Printf("Pair 1: %s\n", pair1)
	fmt.Printf("Pair 2: %s\n", pair2)
	fmt.Printf("Magnitude: %v\n", mag)

}

func explode(pairs string) (exploded bool, newPairs string) {

	left := func(leftStr string, explodingNum int) (newLeftString string) {
		for l := len(leftStr) - 1; l >= 0; l-- {
			if unicode.IsNumber(rune(leftStr[l])) {
				strNum := string(leftStr[l])
				if unicode.IsNumber(rune(leftStr[l-1])) {
					strNum = string(leftStr[l-1]) + strNum
				}

				leftNum, err := strconv.Atoi(string(strNum))
				if err != nil {
					log.Fatal(err)
				}
				return leftStr[:l-len(strNum)+1] + fmt.Sprint(explodingNum+leftNum) + leftStr[l+1:]
			}
		}
		return leftStr //no change if no number on left was found
	}
	right := func(rightStr string, explodingNum int) (newRightString string) {
		for r := 0; r < len(rightStr); r++ {
			if unicode.IsNumber(rune(rightStr[r])) {
				strNum := string(rightStr[r])
				if unicode.IsNumber(rune(rightStr[r+1])) {
					strNum = strNum + string(rightStr[r+1])
				}

				rightNum, err := strconv.Atoi(string(strNum))
				if err != nil {
					log.Fatal(err)
				}
				return rightStr[:r] + fmt.Sprint(explodingNum+rightNum) + rightStr[r+len(strNum):]
			}
		}
		return rightStr //no change if no number on left was found
	}

	openBraceCount := 0
	for i, ch := range pairs {
		if ch == '[' {
			openBraceCount++
		}
		if ch == ']' {
			openBraceCount--
		}
		if yes, leftNum, rightNum, lenPair := isPairStart(pairs[i:]); yes && openBraceCount > 4 {
			return true, left(pairs[:i], leftNum) + "0" + right(pairs[i+lenPair+1:], rightNum)
		}
	}
	return false, pairs
}

func isPairStart(str string) (yes bool, leftNum, rightNum int, lenPair int) {
	if str[0] != '[' {
		return
	}
	var comma int
	for i := 1; i < len(str); i++ {
		if str[i] == ',' {
			comma = i
		}
		if str[i] == ']' {
			yes = true
			lenPair = i
			break
		}
		if str[i] == '[' { //this would indicate opening a new pair, not closing the same one
			return
		}
	}

	leftNum, err := strconv.Atoi(str[1:comma]) //trim the leading [
	if err != nil {
		log.Fatal(err)
	}

	rightNum, err = strconv.Atoi(str[comma+1 : lenPair]) //split ending ]
	if err != nil {
		log.Fatal(err)
	}

	return
}

func split(pairs string) (split bool, newPairs string) {
	for i := 0; i < len(pairs)-2; i++ {
		if !unicode.IsNumber(rune(pairs[i])) || !unicode.IsNumber(rune(pairs[i+1])) {
			continue
		}
		num, err := strconv.Atoi(pairs[i : i+2])
		if err != nil {
			continue
		}
		if num < 10 { //shouldn't be possible since num is 2 chars long
			continue
		}

		newLeftNum := int(math.Floor(float64(num) / 2))
		newRightNum := int(math.Ceil(float64(num) / 2))
		return true, pairs[:i] + "[" + fmt.Sprint(newLeftNum) + "," + fmt.Sprint(newRightNum) + "]" + pairs[i+2:]
	}

	return false, pairs
}

func add(leftPair, rightPair string) (newPairs string) {
	return "[" + leftPair + "," + rightPair + "]"
}

func reduce(pairs string) (newPairs string) {
	newPairs = pairs
	for {
		var exploded, splited bool

		exploded, newPairs = explode(newPairs)
		if exploded {
			continue
		}

		splited, newPairs = split(newPairs)
		if splited {
			continue
		}

		if !exploded && !splited {
			return
		}
	}
}

func sumList(list []string) (answer string) {
	for i, line := range list {
		if i == 0 {
			answer = line
			continue
		}
		answer = reduce(add(answer, line))
	}
	return
}

func magnitude(answer string) int {
	for {
		for i := 0; i < len(answer); i++ {
			if yes, leftNum, rightNum, lenPair := isPairStart(answer[i:]); yes {
				answer = answer[:i] + fmt.Sprint(leftNum*3+rightNum*2) + answer[i+lenPair+1:]
				break
			}
		}
		if !strings.Contains(answer, ",") {
			if mag, err := strconv.Atoi(answer); err == nil {
				return mag
			}
		}
	}
}

func greatestMagnitude(lines []string) (pair1, pair2 string, greatestMag int) {
	for t, testLine := range lines {
		for i, line := range lines {
			if t == i {
				continue //don't add the same line together
			}
			mag := magnitude(reduce(add(testLine, line)))
			if mag > greatestMag {
				pair1 = testLine
				pair2 = line
				greatestMag = mag
			}

			mag = magnitude(reduce(add(line, testLine))) //order matters
			if mag > greatestMag {
				pair1 = testLine
				pair2 = line
				greatestMag = mag
			}
		}
	}

	return
}
