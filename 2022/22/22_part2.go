package main

import (
	"log"
)

type cube [6][][]string

func (grid grid) makeCubeInput() (cube cube, side int) {
	side = len(grid)
	if len(grid[0]) > side {
		side = len(grid[0])
	}
	side /= 4

	for i := range cube {
		cube[i] = make([][]string, side)
		for r := range cube[i] {
			cube[i][r] = make([]string, side)
		}
	}

	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[0][r][c] = grid[r][c+side]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[1][r][c] = grid[r+side][c+side]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[2][r][c] = grid[r+side*2][c+side]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[3][r][c] = grid[r+side*2][c]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[4][r][c] = grid[r+side*3][c]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[5][r][c] = grid[r][c+side*2]
		}
	}

	return
}

func (cube cube) assembleGridInput() (grid grid) {
	side := len(cube[0])
	for i := 0; i < side*4; i++ {
		grid = append(grid, make([]string, side*3))
	}

	for r, row := range cube[0] {
		for c, col := range row {
			grid[r][c+side] = col
		}
	}
	for r, row := range cube[1] {
		for c, col := range row {
			grid[r+side][c+side] = col
		}
	}
	for r, row := range cube[2] {
		for c, col := range row {
			grid[r+side*2][c+side] = col
		}
	}
	for r, row := range cube[3] {
		//lint:ignore S1001 consistency
		for c, col := range row {
			grid[r+side*2][c] = col
		}
	}
	for r, row := range cube[4] {
		//lint:ignore S1001 consistency
		for c, col := range row {
			grid[r+side*3][c] = col
		}
	}
	for r, row := range cube[5] {
		for c, col := range row {
			grid[r][c+side*2] = col
		}
	}

	return
}

func (cube cube) printInput() {
	grid := cube.assembleGridInput()
	grid.print()
}

type cubePos struct {
	cube   int
	r, c   int
	facing int
}

func (pos cubePos) gridPosInput(side int) (gPos position) {
	switch pos.cube {
	case 0:
		gPos.r = pos.r
		gPos.c = pos.c + side
	case 1:
		gPos.r = pos.r + side
		gPos.c = pos.c + side
	case 2:
		gPos.r = pos.r + side*2
		gPos.c = pos.c + side
	case 3:
		gPos.r = pos.r + side*2
		gPos.c = pos.c
	case 4:
		gPos.r = pos.r + side*3
		gPos.c = pos.c
	case 5:
		gPos.r = pos.r
		gPos.c = pos.c + side*2
	}
	return
}

func (grid grid) move2Input(instructions []instruction) (gridPos position, gridFacing int, g2 grid) {
	pos := cubePos{0, 0, 0, 90}
	cube, side := grid.makeCubeInput()

	wrap := func() {
		switch pos.cube {
		case 0:
			switch pos.facing {
			case 0:
				pos.cube = 4
				pos.r = pos.c
				pos.c = 0
				pos.facing = 90
			case 90:
				pos.cube = 5
				pos.c = 0
			case 180:
				pos.cube = 1
				pos.r = 0
			case 270:
				pos.cube = 3
				pos.c = 0
				pos.r = side - 1 - pos.r
				pos.facing = 90
			}
		case 1:
			switch pos.facing {
			case 0:
				pos.cube = 0
				pos.r = side - 1
			case 90:
				pos.cube = 5
				pos.c = pos.r
				pos.r = side - 1
				pos.facing = 0
			case 180:
				pos.cube = 2
				pos.r = 0
			case 270:
				pos.cube = 3
				pos.c = pos.r
				pos.r = 0
				pos.facing = 180
			}
		case 2:
			switch pos.facing {
			case 0:
				pos.cube = 1
				pos.r = side - 1
			case 90:
				pos.cube = 5
				pos.r = side - 1 - pos.r
				pos.c = side - 1
				pos.facing = 270
			case 180:
				pos.cube = 4
				pos.r = pos.c
				pos.c = side - 1
				pos.facing = 270
			case 270:
				pos.cube = 3
				pos.c = side - 1
			}
		case 3:
			switch pos.facing {
			case 0:
				pos.cube = 1
				pos.r = pos.c
				pos.c = 0
				pos.facing = 90
			case 90:
				pos.cube = 2
				pos.c = 0
			case 180:
				pos.cube = 4
				pos.r = 0
			case 270:
				pos.cube = 0
				pos.r = side - 1 - pos.r
				pos.c = 0
				pos.facing = 90
			}
		case 4:
			switch pos.facing {
			case 0:
				pos.cube = 3
				pos.r = side - 1
			case 90:
				pos.cube = 2
				pos.c = pos.r
				pos.r = side - 1
				pos.facing = 0
			case 180:
				pos.cube = 5
				pos.r = 0
			case 270:
				pos.cube = 0
				pos.c = pos.r
				pos.r = 0
				pos.facing = 180
			}
		case 5:
			switch pos.facing {
			case 0:
				pos.cube = 4
				pos.r = side - 1
			case 90:
				pos.cube = 2
				pos.c = side - 1
				pos.r = side - 1 - pos.r
				pos.facing = 270
			case 180:
				pos.cube = 1
				pos.r = pos.c
				pos.c = side - 1
				pos.facing = 270
			case 270:
				pos.cube = 0
				pos.c = side - 1
			}

		}
	}

	drawPos := func() {
		switch pos.facing {
		case 0:
			cube[pos.cube][pos.r][pos.c] = "^"
		case 90:
			cube[pos.cube][pos.r][pos.c] = ">"
		case 180:
			cube[pos.cube][pos.r][pos.c] = "v"
		case 270:
			cube[pos.cube][pos.r][pos.c] = "<"
		}
	}
	drawPos()

	for _, instruction := range instructions {
		for i := 0; i < instruction.move; i++ {
			formerPos := pos

			switch pos.facing {
			case 0:
				pos.r--
			case 90:
				pos.c++
			case 180:
				pos.r++
			case 270:
				pos.c--
			default:
				log.Fatalf("unkown facing: %v", pos.facing)
			}

			if pos.r < 0 || pos.r == side || pos.c < 0 || pos.c == side {
				wrap()
			}

			if cube[pos.cube][pos.r][pos.c] == "#" {
				pos = formerPos
				drawPos()
				break
			}

			drawPos()
		}

		if instruction.turn == "" {
			continue
		}
		switch instruction.turn {
		case "R":
			pos.facing += 90
			if pos.facing == 360 {
				pos.facing = 0
			}
		case "L":
			pos.facing -= 90
			if pos.facing < 0 {
				pos.facing += 360
			}
		default:
			log.Fatalf("turn not understood: %v", instruction.turn)
		}

		drawPos()
		//cube.printInput()
	}

	return pos.gridPosInput(side), pos.facing, cube.assembleGridInput()
}
