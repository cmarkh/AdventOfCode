package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`

func TestPart1(t *testing.T) {
	strInstructions := sanitize(test1)
	for _, instruction := range strInstructions {
		fmt.Println(instruction)
	}
	fmt.Println()

	alu := executeInstructions(strInstructions, "13579246899999")
	fmt.Printf("%+v\n", alu)
	fmt.Printf("valid: %v\n", alu.valid())

	fmt.Println()
}
