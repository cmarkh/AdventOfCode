package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
)

func main() {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	players := parse(lines)
	for i, player := range players {
		fmt.Printf("player %v: %+v\n", i, player)
	}
	fmt.Println()

	fmt.Println("Part 1:")

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

	fmt.Println("Part 2:")

	players = parse(lines)
	wins1, wins2 := playDirac(players)
	fmt.Printf("player 1 won %v games\n", wins1)
	fmt.Printf("player 2 won %v games\n", wins2)
	if wins1 > wins2 {
		fmt.Printf("player 1 is the master: %v\n", wins1)
	} else {
		fmt.Printf("player 2 is the master: %v\n", wins2)
	}

}

type player struct {
	position int
	score    int
}
type game struct {
	players    []player
	diceRolls  int
	lastRoll   int
	scoreToWin int
	sidedDie   int
}

func parse(lines []string) (players []player) {
	for _, line := range lines {
		pos, err := strconv.Atoi(strings.Split(line, ": ")[1])
		if err != nil {
			log.Fatal(err)
		}
		players = append(players, player{pos, 0})
	}
	return
}

func (g *game) roll() {
	for i := range g.players {
		totalRoll := 0
		for d := 0; d < 3; d++ {
			g.lastRoll += 1
			g.diceRolls++
			totalRoll += g.lastRoll
			if g.lastRoll >= g.sidedDie {
				g.lastRoll = 0
			}
		}

		g.players[i].position += totalRoll
		g.players[i].score += wrapAround(g.players[i].position)

		if g.players[i].score >= g.scoreToWin {
			return
		}
	}
}

func (g *game) play() (winner int) {
	for {
		for i, player := range g.players {
			if player.score >= g.scoreToWin {
				return i
			}
		}
		g.roll()
		//fmt.Printf("%+v\n", g)
	}
}

func wrapAround(position int) (wrapped int) {
	wrapped = position
	if position > 10 {
		wrapped -= int(position/10) * 10
		if position%10 == 0 { //if divisible by 10, subtracted an extra 10 in line above
			wrapped += 10
		}
	}
	return
}

func playDirac(players []player) (wins1, wins2 int) {
	rolls := make(map[int]int) //map[roll]occurances
	rolls[3] = 1
	rolls[4] = 3
	rolls[5] = 6
	rolls[6] = 7
	rolls[7] = 6
	rolls[8] = 3
	rolls[9] = 1

	type game struct {
		position   [2]int
		score      [2]int
		occurances int
		turn       int
	}
	games := []game{{[2]int{players[0].position, players[1].position}, [2]int{0, 0}, 1, 1}}

	for len(games) > 0 {
		g := games[len(games)-1]
		games = games[:len(games)-1]

		if g.score[0] >= 21 {
			wins1 += g.occurances
			continue
		}
		if g.score[1] >= 21 {
			wins2 += g.occurances
			continue
		}

		for roll, occurances := range rolls {
			g2 := g
			g2.position[1-g2.turn] = wrapAround(g2.position[1-g2.turn] + roll)
			g2.score[1-g2.turn] += g2.position[1-g2.turn]
			g2.occurances *= occurances
			g2.turn ^= 1
			games = append(games, g2)
		}
	}

	return
}
