package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	cubes := parse(input)

	fmt.Println("Part 1:")
	area := surfaceArea(cubes)
	fmt.Printf("surface area: %v\n", area)
	fmt.Println()

	fmt.Println("Part 2:")
	area = trappedDrops(cubes)
	fmt.Printf("surface area: %v\n", area)
	fmt.Println()
}

type cube struct {
	x, y, z int
}

func parse(input string) (cubes []cube) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ",")
		if len(split) != 3 {
			log.Fatalf("line not understood: %v", line)
		}
		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}
		cubes = append(cubes, cube{x, y, z})
	}

	return
}

func surfaceArea(cubes []cube) (area int) {
	for _, cube1 := range cubes {
		cubeArea := 6
		for _, cube2 := range cubes {
			if abs(cube1.x-cube2.x) == 1 && cube1.y == cube2.y && cube1.z == cube2.z {
				cubeArea--
			}
		}
		for _, cube2 := range cubes {
			if abs(cube1.y-cube2.y) == 1 && cube1.x == cube2.x && cube1.z == cube2.z {
				cubeArea--
			}
		}
		for _, cube2 := range cubes {
			if abs(cube1.z-cube2.z) == 1 && cube1.x == cube2.x && cube1.y == cube2.y {
				cubeArea--
			}
		}
		area += cubeArea
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func size(cubes []cube) (minX, minY, minZ int, maxX, maxY, maxZ int) {
	for _, cube := range cubes {
		if cube.x < minX {
			minX = cube.x
		}
		if cube.y < minY {
			minY = cube.y
		}
		if cube.z < minZ {
			minZ = cube.z
		}

		if cube.x > maxX {
			maxX = cube.x
		}
		if cube.y > maxY {
			maxY = cube.y
		}
		if cube.z > maxZ {
			maxZ = cube.z
		}
	}

	return
}

func cubeAt(c1 cube, cubes []cube) bool {
	for _, c2 := range cubes {
		if c1.x == c2.x && c1.y == c2.y && c1.z == c2.z {
			return true
		}
	}
	return false
}

func trappedDrops(cubes []cube) (area int) {
	minX, minY, minZ, maxX, maxY, maxZ := size(cubes)

	queue := []cube{{0, 0, 0}}
	visited := []cube{}

	for len(queue) > 0 {
		c1 := queue[0]
		queue = queue[1:]

		if slices.Contains(visited, c1) {
			continue
		}
		visited = append(visited, c1)

		if cubeAt(c1, cubes) {
			continue
		}
		if c1.x < minX-1 || c1.x > maxX+1 {
			continue
		}
		if c1.y < minY-1 || c1.y > maxY+1 {
			continue
		}
		if c1.z < minZ-1 || c1.z > maxZ+1 {
			continue
		}

		if cubeAt(cube{c1.x - 1, c1.y, c1.z}, cubes) {
			area++
		}
		if cubeAt(cube{c1.x + 1, c1.y, c1.z}, cubes) {
			area++
		}
		if cubeAt(cube{c1.x, c1.y - 1, c1.z}, cubes) {
			area++
		}
		if cubeAt(cube{c1.x, c1.y + 1, c1.z}, cubes) {
			area++
		}
		if cubeAt(cube{c1.x, c1.y, c1.z - 1}, cubes) {
			area++
		}
		if cubeAt(cube{c1.x, c1.y, c1.z + 1}, cubes) {
			area++
		}

		queue = append(queue, cube{c1.x - 1, c1.y, c1.z})
		queue = append(queue, cube{c1.x + 1, c1.y, c1.z})
		queue = append(queue, cube{c1.x, c1.y - 1, c1.z})
		queue = append(queue, cube{c1.x, c1.y + 1, c1.z})
		queue = append(queue, cube{c1.x, c1.y, c1.z - 1})
		queue = append(queue, cube{c1.x, c1.y, c1.z + 1})
	}

	return
}
