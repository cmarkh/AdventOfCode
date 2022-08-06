package main

import (
	"log"
	advent2 "temp/adventofcode/go/2"
)

func main() {
	directions, err := advent2.Input()
	if err != nil {
		log.Fatal(err)
	}

	_, err = advent2.Aim(directions)
	if err != nil {
		log.Fatal(err)
	}
}
