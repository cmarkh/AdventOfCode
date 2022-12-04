package main

import (
	"log"

	advent1 "temp/adventofcode/go/2021/1"
)

func main() {
	depths, err := advent1.Input()
	if err != nil {
		log.Fatal(err)
	}

	//_, err = countDepthIncreases(depths)
	if err != nil {
		log.Fatal(err)
	}

	_, err = advent1.SlidingWindow(depths)
	if err != nil {
		log.Fatal(err)
	}
}
