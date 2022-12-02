package main

import (
	"fmt"
	"strings"
	"testing"
)

var test1 = `Player 1 starting position: 4
Player 2 starting position: 8`

func TestParse(t *testing.T) {
	players := parse(strings.Split(test1, "\n"))
	for i, player := range players {
		fmt.Printf("player %v: %+v\n", i, player)
	}
	fmt.Println()
}

func TestPlay(t *testing.T) {
	players := parse(strings.Split(test1, "\n"))
	for i, player := range players {
		fmt.Printf("player %v: %+v\n", i, player)
	}
	fmt.Println()

	g := game{players, 0, 0, 1000, 100}
	winner := g.play()
	fmt.Printf("%v won with %v\n", winner, g.players[winner].score)

	product := 1
	for i, player := range g.players {
		if i != winner {
			fmt.Printf("%v had %v points\n", i, player.score)
			product *= player.score
		}
	}
	product *= g.diceRolls
	fmt.Printf("the product of scores and rolls is %v\n", product)

	fmt.Println()
}

func TestPart2(t *testing.T) {
	players := parse(strings.Split(test1, "\n"))
	for i, player := range players {
		fmt.Printf("player %v: %+v\n", i, player)
	}
	fmt.Println()

	wins1, wins2 := playDirac(players)
	fmt.Printf("player 1 won %v games\n", wins1)
	fmt.Printf("player 2 won %v games\n", wins2)
	if wins1 > wins2 {
		fmt.Printf("player 1 is the master: %v\n", wins1)
	} else {
		fmt.Printf("player 2 is the master: %v\n", wins2)
	}

	fmt.Println()
}
