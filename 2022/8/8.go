package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	grid := parse(input)

	fmt.Println("Part 1:")
	grid.visibleTrees()
	count := grid.countVisible()
	fmt.Printf("%v visible trees\n", count)
	fmt.Println()

	fmt.Println("Part 2:")
	best := grid.scenicScores()
	fmt.Printf("best scenic score: %v\n", best)
	fmt.Println()
}

type grid [][]tree

type tree struct {
	height  int
	visible bool
}

func parse(input string) (grid grid) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		row := []tree{}
		for _, ch := range line {
			height, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, tree{height, false})
		}
		grid = append(grid, row)
	}

	return
}

func (g grid) print() {
	for _, row := range g {
		for _, col := range row {
			fmt.Print(col.height)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) printVisible() {
	for _, row := range g {
		for _, col := range row {
			fmt.Print(col.visible)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) visibleTrees() {
	//top down
	highest := 0
	for c := 0; c < len(g[0]); c++ {
		for r := 0; r < len(g); r++ {
			if r == 0 {
				g[r][c].visible = true
				highest = g[r][c].height
				continue
			}
			if g[r][c].height > highest {
				g[r][c].visible = true
				highest = g[r][c].height
			}
		}
	}

	//bottom up
	highest = 0
	for c := 0; c < len(g[0]); c++ {
		for r := len(g) - 1; r >= 0; r-- {
			if r == len(g)-1 {
				g[r][c].visible = true
				highest = g[r][c].height
				continue
			}
			if g[r][c].height > highest {
				g[r][c].visible = true
				highest = g[r][c].height
			}
		}
	}

	//left to right
	highest = 0
	for r := 0; r < len(g); r++ {
		for c := 0; c < len(g[r]); c++ {
			if c == 0 {
				g[r][c].visible = true
				highest = g[r][c].height
				continue
			}
			if g[r][c].height > highest {
				g[r][c].visible = true
				highest = g[r][c].height
			}
		}
	}

	//right to left
	highest = 0
	for r := 0; r < len(g); r++ {
		for c := len(g[r]) - 1; c >= 0; c-- {
			if c == len(g[r])-1 {
				g[r][c].visible = true
				highest = g[r][c].height
				continue
			}
			if g[r][c].height > highest {
				g[r][c].visible = true
				highest = g[r][c].height
			}
		}
	}
}

func (g grid) countVisible() (count int) {
	for _, row := range g {
		for _, col := range row {
			if col.visible {
				count++
			}
		}
	}
	return
}

func (g grid) scenicScores() (best int) {
	for r := range g {
		for c := range g[r] {
			treeScore := g.treeScenicScore(r, c)
			if treeScore > best {
				best = treeScore
			}
		}
	}

	return
}

func (g grid) treeScenicScore(row, col int) (score int) {
	if row == 0 || row == len(g)-1 {
		return 0
	}
	if col == 0 || col == len(g[0])-1 {
		return 0
	}
	tree := g[row][col]

	up := 0
	for r := row - 1; r >= 0; r-- {
		up++
		if g[r][col].height >= tree.height {
			break
		}
	}

	down := 0
	for r := row + 1; r < len(g); r++ {
		down++
		if g[r][col].height >= tree.height {
			break
		}
	}

	left := 0
	for c := col - 1; c >= 0; c-- {
		left++
		if g[row][c].height >= tree.height {
			break
		}
	}

	right := 0
	for c := col + 1; c < len(g[0]); c++ {
		right++
		if g[row][c].height >= tree.height {
			break
		}
	}

	return up * down * left * right
}
