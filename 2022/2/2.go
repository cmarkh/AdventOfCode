package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go"
)

func main() {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strats := parse(lines)
	for _, strat := range strats {
		fmt.Printf("%+v\n", strat)
	}
	fmt.Println()

	fmt.Println("Part 1:")
	game := play(strats, yourHandPart1)
	fmt.Printf("opponent score: %v\n", game.scoreOpp)
	fmt.Printf("your score: %v\n", game.scoreYou)
	fmt.Println()

	fmt.Println("Part 2:")
	game = play(strats, yourHandPart2)
	fmt.Printf("opponent score: %v\n", game.scoreOpp)
	fmt.Printf("your score: %v\n", game.scoreYou)
	fmt.Println()
}

var hands = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

type game struct {
	//round              int
	scoreOpp, scoreYou int
}

type strategy struct {
	opponent  string
	suggested string
}

func parse(lines []string) (strats []strategy) {
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, " ")
		if len(s) != 2 {
			log.Fatalf("failed to parse line: %s", line)
		}
		strats = append(strats, strategy{s[0], s[1]})
	}

	return
}

func yourHandPart1(strat strategy) string {
	return strat.suggested
}

func yourHandPart2(strat strategy) string {
	//X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	switch strat.suggested {
	case "X":
		switch hands[strat.opponent] {
		case "rock":
			return "C"
		case "paper":
			return "A"
		case "scissors":
			return "B"
		}
	case "Y":
		return strat.opponent
	case "Z":
		switch hands[strat.opponent] {
		case "rock":
			return "B"
		case "paper":
			return "C"
		case "scissors":
			return "A"
		}
	default:
		log.Fatalf("suggested strat no understood: %+v", strat)
	}
	return ""
}

func play(strats []strategy, yourHand func(strat strategy) string) (g game) {
	for _, strat := range strats {
		opp, you := round(strat.opponent, yourHand(strat))
		g.scoreOpp += opp
		g.scoreYou += you
		//fmt.Printf("opp: %v, you: %v\n", strat.opponent, strat.suggested)
		//fmt.Printf("opp: %v, you: %v\n", opp, you)
		//fmt.Println()
	}
	return
}

func round(hand1, hand2 string) (wins1, wins2 int) {
	//The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	switch hands[hand1] {
	case "rock":
		switch hands[hand2] {
		case "rock":
			return 1 + 3, 1 + 3
		case "paper":
			return 1, 2 + 6
		case "scissors":
			return 1 + 6, 3
		default:
			log.Fatalf("unkown hand1: %v", hand1)
			return 0, 0
		}
	case "paper":
		switch hands[hand2] {
		case "rock":
			return 2 + 6, 1
		case "paper":
			return 2 + 3, 2 + 3
		case "scissors":
			return 2, 3 + 6
		default:
			log.Fatalf("unkown hand1: %v", hand1)
			return 0, 0
		}
	case "scissors":
		switch hands[hand2] {
		case "rock":
			return 3, 1 + 6
		case "paper":
			return 3 + 6, 2
		case "scissors":
			return 3 + 3, 3 + 3
		default:
			log.Fatalf("unkown hand1: %v", hand1)
			return 0, 0
		}
	default:
		log.Fatalf("unkown hand1: %v", hand1)
		return 0, 0
	}
}
