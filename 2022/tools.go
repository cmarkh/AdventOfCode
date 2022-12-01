package advent

import (
	"log"
	"os"
	"strings"
)

func ReadInput(path string) (lines []string, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(content), "\n")
	return
}

func Purple(text string) string {
	return "\033[35m" + text + "\033[0m"
}
