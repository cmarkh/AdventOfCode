package main

import (
	"log"
)

func main() {
	numbers, err := Input()
	if err != nil {
		log.Fatal(err)
	}

	_, _, _, err = GammaRate(numbers)
	if err != nil {
		log.Fatal(err)
	}

	err = LifeSuport()
	if err != nil {
		log.Fatal(err)
	}
}
