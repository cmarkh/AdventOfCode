package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`
var test2 = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

func TestPart1(t *testing.T) {
	marker := marker(test1, 4)
	fmt.Printf("marker at %v character\n", marker)
	fmt.Println()
}

func TestPart12(t *testing.T) {
	marker := startOfPacket(test2)
	fmt.Printf("marker at %v character\n", marker)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	marker := marker(test1, 14)
	fmt.Printf("marker at %v character\n", marker)
	fmt.Println()
}

func TestPart22(t *testing.T) {
	marker := startOfMessage(test2)
	fmt.Printf("marker at %v character\n", marker)
	fmt.Println()
}
