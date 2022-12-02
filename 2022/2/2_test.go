package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	strats := parse(strings.Split(test1, "\n"))
	for _, strat := range strats {
		fmt.Printf("%+v\n", strat)
	}
	fmt.Println()

	game := play(strats, yourHandPart1)
	fmt.Printf("opponent score: %v\n", game.scoreOpp)
	fmt.Printf("your score: %v\n", game.scoreYou)

	fmt.Println()
}

func TestPart2(t *testing.T) {
	strats := parse(strings.Split(test1, "\n"))
	for _, strat := range strats {
		fmt.Printf("%+v\n", strat)
	}
	fmt.Println()

	game := play(strats, yourHandPart2)
	fmt.Printf("opponent score: %v\n", game.scoreOpp)
	fmt.Printf("your score: %v\n", game.scoreYou)

	fmt.Println()
}
