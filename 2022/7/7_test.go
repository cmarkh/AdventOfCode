package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestParse(t *testing.T) {
	topDir := parse(test1)
	topDir.print(0)
	fmt.Println()
}

func TestSize(t *testing.T) {
	topDir := parse(test1)
	size := topDir.size()
	fmt.Printf("size: %v\n", size)
	fmt.Println()
}

func TestMinSize(t *testing.T) {
	topDir := parse(test1)
	dirs := topDir.directoriesWithMaxSize(100000)
	for _, dir := range dirs {
		fmt.Printf("%s: %v\n", dir.name, dir.size())
	}
	fmt.Println()
}

func TestPart1(t *testing.T) {
	topDir := parse(test1)

	dirs := topDir.directoriesWithMaxSize(100000)
	for _, dir := range dirs {
		fmt.Printf("%s: %v\n", dir.name, dir.size())
	}

	size := sumSizes(dirs)
	fmt.Printf("sum of total sizes: %v\n", size)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	topDir := parse(test1)
	dirs := topDir.listDirectories()
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

func TestParse2(t *testing.T) {
	topDir := parse(input)
	topDir.print(0)
	fmt.Println()
}
