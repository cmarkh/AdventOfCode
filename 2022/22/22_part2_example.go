package main

import "log"

func (grid grid) makeCubeExample() (cube cube, side int) {
	side = len(grid) / 3
	for i := range cube {
		cube[i] = make([][]string, side)
		for r := range cube[i] {
			cube[i][r] = make([]string, side)
		}
	}

	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[0][r][c] = grid[r][c+side*2]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[1][r][c] = grid[r+side][c]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[2][r][c] = grid[r+side][c+side]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[3][r][c] = grid[r+side][c+side*2]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[4][r][c] = grid[r+side*2][c+side*2]
		}
	}
	for r := 0; r < side; r++ {
		for c := 0; c < side; c++ {
			cube[5][r][c] = grid[r+side*2][c+side*3]
		}
	}

	return
}

func (cube cube) assembleGridExample() (grid grid) {
	side := len(cube[0])
	for i := 0; i < side*3; i++ {
		grid = append(grid, make([]string, side*4))
	}

	for r, row := range cube[0] { //face 1
		for c, col := range row {
			grid[r][c+side*2] = col
		}
	}
	for r, row := range cube[1] { //face 2
		//lint:ignore S1001 consistency
		for c, col := range row {
			grid[r+side][c] = col
		}
	}
	for r, row := range cube[2] { //face 3
		for c, col := range row {
			grid[r+side][c+side] = col
		}
	}
	for r, row := range cube[3] { //face 4
		for c, col := range row {
			grid[r+side][c+side*2] = col
		}
	}
	for r, row := range cube[4] { //face 5
		for c, col := range row {
			grid[r+side*2][c+side*2] = col
		}
	}
	for r, row := range cube[5] { //face 6
		for c, col := range row {
			grid[r+side*2][c+side*3] = col
		}
	}

	return
}

func (cube cube) printExample() {
	grid := cube.assembleGridExample()
	grid.print()
}

func (pos cubePos) gridPosExample(side int) (gPos position) {
	switch pos.cube {
	case 0:
		gPos.r = pos.r
		gPos.c = pos.c + side*2
	case 1:
		gPos.r = pos.r + side
		gPos.c = pos.c
	case 2:
		gPos.r = pos.r + side
		gPos.c = pos.c + side
	case 3:
		gPos.r = pos.r + side
		gPos.c = pos.c + side*2
	case 4:
		gPos.r = pos.r + side*2
		gPos.c = pos.c + side*2
	case 5:
		gPos.r = pos.r + side*2
		gPos.c = pos.c + side*3
	}
	return
}

func (grid grid) move2Example(instructions []instruction) (gridPos position, gridFacing int, g2 grid) {
	pos := cubePos{0, 0, 0, 90}
	cube, side := grid.makeCubeExample()

	wrap := func() {
		switch pos.cube {
		case 0: //face 1
			switch pos.facing {
			case 0:
				pos.cube = 1 //face 2
				pos.r = 0
				pos.facing = 180
			case 90:
				pos.cube = 5 //face 6
				pos.r = side - 1 - pos.r
				pos.c = side - 1
				pos.facing = 270
			case 180:
				pos.cube = 3 //face 4
				pos.r = 0
			case 270:
				pos.cube = 2 //face 3
				pos.c = pos.r
				pos.r = 0
				pos.facing = 180
			}
		case 1: //face 2
			switch pos.facing {
			case 0:
				pos.cube = 0 //face 1
				pos.r = 0
				pos.facing = 180
			case 90:
				pos.cube = 2 //face 3
				pos.c = 0
			case 180:
				pos.cube = 4 //face 5
				pos.r = side - 1
				pos.facing = 0
			case 270:
				pos.cube = 5 //face 6
				pos.c = pos.r
				pos.r = side - 1
				pos.facing = 0
			}
		case 2: //face 3
			switch pos.facing {
			case 0:
				pos.cube = 0 //face 1
				pos.r = pos.c
				pos.c = 0
				pos.facing = 90
			case 90:
				pos.cube = 3 //face 4
				pos.c = 0
			case 180:
				pos.cube = 4 //face 5
				pos.r = pos.c
				pos.c = 0
				pos.facing = 90
			case 270:
				pos.cube = 1 //face 2
				pos.c = side - 1
			}
		case 3: //face 4
			switch pos.facing {
			case 0:
				pos.cube = 0 //face 1
				pos.r = side - 1
			case 90:
				pos.cube = 5 //face 6
				pos.c = side - 1 - pos.r
				pos.r = 0
				pos.facing = 180
			case 180:
				pos.cube = 4 //face 5
				pos.r = 0
			case 270:
				pos.cube = 2 //face 3
				pos.c = side - 1
			}
		case 4: //face 5
			switch pos.facing {
			case 0:
				pos.cube = 3 //face 4
				pos.r = side - 1
			case 90:
				pos.cube = 5 //face 6
				pos.c = pos.r
				pos.r = 0
			case 180:
				pos.cube = 1 //face 2
				pos.c = side - 1 - pos.c
				pos.r = side - 1
				pos.facing = 0
			case 270:
				pos.cube = 2 //face 3
				pos.c = pos.r
				pos.r = side - 1
				pos.facing = 0
			}
		case 5: //face 6
			switch pos.facing {
			case 0:
				pos.cube = 3 //face 4
				pos.r = pos.c
				pos.c = side - 1
				pos.facing = 270
			case 90:
				pos.cube = 4 //face 5
				pos.c = side - 1
			case 180:
				pos.cube = 1 //face 2
				pos.r = pos.c
				pos.c = side - 1
				pos.facing = 270
			case 270:
				pos.cube = 4 //face 5
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
		cube.printExample()
	}

	return pos.gridPosExample(side), pos.facing, cube.assembleGridExample()
}
