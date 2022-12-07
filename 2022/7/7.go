package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var (
	totalDiskSpace = 70000000
	spaceRequired  = 30000000
)

func main() {
	topDir := parse(input)

	fmt.Println("Part 1:")

	dirs := topDir.directoriesWithMaxSize(100000)
	for _, dir := range dirs {
		fmt.Printf("%s: %v\n", dir.name, dir.size())
	}
	fmt.Println()

	size := sumSizes(dirs)
	fmt.Printf("sum of total sizes: %v\n", size)
	fmt.Println()

	fmt.Println("Part 2:")
	dirs = topDir.listDirectories()
	for _, dir := range dirs {
		fmt.Printf("%s: %v\n", dir.name, dir.size())
	}
	fmt.Println()

	spaceRequired := spaceRequired - topDir.unusedSpace()
	fmt.Printf("space required: %v\n", spaceRequired)

	smallest, size := smallestDeletion(dirs, spaceRequired)
	fmt.Printf("%s: %v\n", smallest.name, size)

	fmt.Println()
}

type directories map[string]directory //map[dir name]contents
type directory struct {
	name           string
	subDirectories directories
	files          []file
}
type file struct {
	name string
	size int
}

func parse(input string) (top directory) {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var walk func(line int) (dir directory, endLine int)
	walk = func(line int) (dir directory, endLine int) {
		dir.subDirectories = make(directories)
		for line < len(lines) {
			split := strings.Split(lines[line], " ")

			if split[1] == "cd" {
				if split[2] == ".." {
					line++
					return dir, line
				}
				var sub directory
				sub, line = walk(line + 1)
				sub.name = split[2]
				dir.subDirectories[split[2]] = sub
				continue
			}
			if split[1] == "ls" {
				line++
				for line < len(lines) {
					split := strings.Split(lines[line], " ")

					if split[1] == "cd" {
						if split[2] == ".." {
							line++
							return dir, line
						}
						var sub directory
						sub, line = walk(line + 1)
						sub.name = split[2]
						dir.subDirectories[split[2]] = sub
						continue
					}
					if split[0] == "dir" {
						line++
						continue
					} else {
						size, err := strconv.Atoi(split[0])
						if err != nil {
							log.Fatal(err)
						}
						dir.files = append(dir.files, file{split[1], size})
						line++
					}
				}
			}

		}
		return dir, line
	}
	top, _ = walk(0)

	return
}

func (dir directory) print(indent int) {
	for name, dir := range dir.subDirectories {
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(name)
		dir.print(indent + 1)
	}
	for _, file := range dir.files {
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(file)
	}
}

func (dir directory) size() (sum int) {
	for _, dir := range dir.subDirectories {
		sum += dir.size()
	}
	for _, file := range dir.files {
		sum += file.size
	}
	return
}

func (dir directory) directoriesWithMaxSize(max int) (dirs []directory) {
	if dir.size() <= max {
		dirs = append(dirs, dir)
	}
	for _, sub := range dir.subDirectories {
		dirs = append(dirs, sub.directoriesWithMaxSize(max)...)
	}
	return
}

func sumSizes(dirs []directory) (sum int) {
	for _, dir := range dirs {
		sum += dir.size()
	}
	return
}

func (dir directory) listDirectories() (dirs []directory) {
	dirs = append(dirs, dir)
	for _, sub := range dir.subDirectories {
		dirs = append(dirs, sub.listDirectories()...)
	}
	return
}

func smallestDeletion(dirs []directory, spaceRequired int) (smallest directory, size int) {
	size = math.MaxInt / 2

	for _, dir := range dirs {
		s := dir.size()
		if s >= spaceRequired && s < size {
			smallest = dir
			size = s
		}
	}

	return
}

func (top directory) unusedSpace() (size int) {
	return totalDiskSpace - top.size()
}
