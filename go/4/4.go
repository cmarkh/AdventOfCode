package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	advent "temp/adventofcode/go"
)

var inputPath = advent.BasePath + "4/input.txt"

func main() {
	GetFirstWinner()
	GetLastWinner()
}

func GetFirstWinner() {
	fmt.Println("First winner: ")

	draws, boards, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	winners, lastDraw := boards.FirstWinner(draws)
	fmt.Printf("last draw: %s\n", lastDraw)

	for _, w := range winners {
		score, err := w.Score(lastDraw)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("score: %d\n", score)
		Boards{w}.Print()
	}
}

func GetLastWinner() {
	fmt.Println("Last winner: ")

	draws, boards, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	loser, lastDraw := boards.LastWinner(draws)
	score, err := loser.Score(lastDraw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("last draw: %s\n", lastDraw)
	fmt.Printf("score: %d\n", score)
	fmt.Printf("last winner: \n")
	Boards{loser}.Print()
}

type Boards []Board
type Board [][]Square
type Square struct {
	Number string
	Marked bool
}

func Input(path string) (draws []string, boards Boards, err error) {
	//input, err := advent.ReadInput(path)
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")

	var board Board
	for i := 0; i+5 < len(lines); {
		if i == 0 {
			draws = strings.Split(lines[i], ",")
			i++
			continue
		}
		if lines[i] == "" {
			board = make(Board, 5)
			for j := 0; j < 5; j++ {
				line := strings.Fields(lines[i+j+1])
				for _, n := range line {
					board[j] = append(board[j], Square{n, false})
				}
				if len(board[j]) != 5 {
					err = fmt.Errorf("failed to parse board: %v", board[j])
					return
				}
			}
			boards = append(boards, board)
			i += 6
		}
	}

	return
}

func (boards Boards) Mark(drawn string) {
	for b, board := range boards {
		for l, line := range board {
			for n, num := range line {
				if num.Number == drawn {
					boards[b][l][n].Marked = true
				}
			}
		}
	}
}

func (board Board) Winner() bool {
	horizontal := func(line []Square) bool {
		for _, num := range line {
			if !num.Marked {
				return false
			}
		}
		return true
	}

	vertical := func(column int, board Board) bool {
		//fmt.Println(column)
		//Boards{board}.Print()
		for row := 0; row < len(board); row++ {
			if !board[row][column].Marked {
				return false
			}
		}
		return true
	}

	for _, line := range board {
		if horizontal(line) {
			return true
		}
	}
	for c := 0; c < len(board[0]); c++ {
		if vertical(c, board) {
			return true
		}
	}
	return false
}

func (boards Boards) CheckWinners() (winner bool, won Boards) {
	for _, board := range boards {
		if board.Winner() {
			winner = true
			won = append(won, board)
		}
	}
	return
}

func (boards Boards) Print() {
	for _, board := range boards {
		for _, line := range board {
			for _, num := range line {
				if num.Marked {
					fmt.Printf("%s'\t", num.Number)
				} else {
					fmt.Printf("%s\t", num.Number)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (boards Boards) PlayXRounds(rounds int, draws []string) (winner bool, won Boards) {
	for i := 0; i < rounds; i++ {
		boards.Mark(draws[i])
	}
	winner, won = boards.CheckWinners()
	return
}

func (board Board) Score(lastDraw string) (score int64, err error) {
	for _, line := range board {
		for _, num := range line {
			if !num.Marked {
				n, err := strconv.ParseInt(num.Number, 10, 64)
				if err != nil {
					return 0, err
				}
				score += n
			}
		}
	}

	last, err := strconv.ParseInt(lastDraw, 10, 64)
	if err != nil {
		return 0, err
	}
	score *= last

	return
}

func (boards Boards) FirstWinner(draws []string) (winners Boards, lastDraw string) {
	for _, draw := range draws {
		boards.Mark(draw)
		winner, winners := boards.CheckWinners()
		if winner {
			return winners, draw
		}
	}
	return
}

func (boards Boards) LastWinner(draws []string) (last Board, lastDraw string) {
	filtered := boards
	for _, draw := range draws {
		boards.Mark(draw)
		losers := Boards{}
		for _, board := range filtered {
			if !board.Winner() {
				losers = append(losers, board)
			}
		}
		filtered = losers
		if len(filtered) == 1 {
			last = filtered[0]
			_, lastDraw = filtered.FirstWinner(draws)
			return
		}
	}

	return
}
