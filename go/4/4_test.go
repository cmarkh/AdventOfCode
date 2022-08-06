package main

import (
	"fmt"
	advent "temp/adventofcode/go"
	"testing"
)

var testPath = advent.BasePath + "4/test.txt"

func TestInput(t *testing.T) {
	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(draws)

	boards.Print()
}

func TestMark(t *testing.T) {
	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	boards.Mark(draws[0])
	fmt.Println(draws[0])
	boards.Print()
}

func TestCheckWinnerHorizontal(t *testing.T) {
	_, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	draws := []string{"22", "13", "17", "11", "0"}
	for _, draw := range draws {
		boards.Mark(draw)
	}
	winner, boards := boards.CheckWinners()
	if !winner {
		t.Fatal("should have been a winner (horizontal)")
	}
	fmt.Printf("winning boards (horizontal): %d\n", len(boards))
	boards.Print()
}

func TestCheckWinnerVertical(t *testing.T) {
	_, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	draws := []string{"22", "8", "21", "6", "1",
		"24", "9", "26", "6", "3"}
	for _, draw := range draws {
		boards.Mark(draw)
	}
	winner, boards := boards.CheckWinners()
	if !winner {
		t.Fatal("should have been a winner")
	}
	fmt.Printf("winning boards (vertical): %d\n", len(boards))
	boards.Print()
}

func TestCheckWinnerMany(t *testing.T) {
	_, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	draws := []string{
		"22", "8", "21", "6", "1",
		"19", "8", "7", "25", "23",
		"22", "8", "21", "6", "1",
		"24", "9", "26", "6", "3",
	}
	for _, draw := range draws {
		boards.Mark(draw)
	}
	winner, boards := boards.CheckWinners()
	if !winner {
		t.Fatal("should have been a winner 1")
	}
	fmt.Printf("winning boards (many): %d\n", len(boards))
	boards.Print()
}

func TestFiveDraws(t *testing.T) {
	fmt.Println("After five draws:")
	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	winner, _ := boards.PlayXRounds(5, draws)

	fmt.Printf("winner: %v\n", winner)
	boards.Print()
}

func TestElevenDraws(t *testing.T) {
	fmt.Println("After 11 draws:")
	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	winner, _ := boards.PlayXRounds(5+6, draws)

	fmt.Printf("winner: %v\n", winner)
	boards.Print()
}

func Test12Draws(t *testing.T) {
	fmt.Println("After 12 draws:")
	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	winner, winners := boards.PlayXRounds(5+6+1, draws)

	fmt.Printf("winner: %v\n", winner)
	boards.Print()

	fmt.Printf("Last draw: %s\n\n", draws[5+6+1-1])

	for _, w := range winners {
		score, err := w.Score(draws[5+6+1-1])
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("Score: %d\n", score)
		Boards{w}.Print()
	}
}

func TestFirstWinner(t *testing.T) {
	fmt.Println("First winner: ")

	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	winners, lastDraw := boards.FirstWinner(draws)
	fmt.Printf("last draw: %s\n", lastDraw)

	for _, w := range winners {
		score, err := w.Score(lastDraw)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("score: %d\n", score)
		Boards{w}.Print()
	}
}

func TestLastWinner(t *testing.T) {
	fmt.Println("Last winner: ")

	draws, boards, err := Input(testPath)
	if err != nil {
		t.Fatal(err)
	}

	loser, lastDraw := boards.LastWinner(draws)
	score, err := loser.Score(lastDraw)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("last draw: %s\n", lastDraw)
	fmt.Printf("score: %d\n", score)
	fmt.Printf("last winner: \n")
	Boards{loser}.Print()
}
