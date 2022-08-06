package advent1

import (
	"log"
	"os"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "1/input.txt"

// input here are depth measurements
func Input() (depths []int64, err error) {
	content, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		depth, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}

	return
}
