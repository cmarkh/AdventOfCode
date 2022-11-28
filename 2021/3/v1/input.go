package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var inputPath = "input.txt"

// input here are depth measurements
func Input() (transposed [][]byte, err error) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")
	lineLen := len(lines[0])
	transposed = make([][]byte, lineLen)

	for _, line := range lines {
		if fmt.Sprint(line) == "" {
			continue
		}

		/*n, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			err = fmt.Errorf("error parsing line: %s", line)
			return numbers, err
		}
		numbers = append(numbers, n)
		*/

		for i := 0; i < lineLen; i++ {
			transposed[i] = append(transposed[i], line[i])
		}

	}

	return
}
