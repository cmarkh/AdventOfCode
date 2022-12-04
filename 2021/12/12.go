package main

import (
	"fmt"
	"log"
	"strings"
	advent "temp/adventofcode/go/2021"

	"golang.org/x/exp/slices"
)

var inputPath = "input.txt"

func main() {
	segments, err := input()
	if err != nil {
		log.Fatal(err)
	}

	/*for key, seg := range segments {
		fmt.Printf("%s: %v\n", key, seg)
	}
	fmt.Println()*/

	paths := PathsPartOne(segments)
	/*for _, p := range paths {
		fmt.Println(p)
	}*/
	fmt.Printf("\n%d possible paths\n\n", len(paths))

	paths = PathsPartTwo(segments)
	fmt.Printf("\n%d possible paths\n\n", len(paths))

}

type Segments map[string][]string //map[letter]connecting letters

func input() (segments Segments, err error) {
	segments = make(Segments)

	input, err := advent.ReadInput(inputPath)
	if err != nil {
		return
	}

	for _, line := range input {
		given := strings.Split(line, "-")
		if len(given) != 2 {
			err = fmt.Errorf("unkown inpit: %v", given)
			return
		}
		segments[given[0]] = append(segments[given[0]], given[1])
		segments[given[1]] = append(segments[given[1]], given[0])
	}

	return
}

type Path []string

func PathsPartTwo(segments Segments) (paths []Path) {
	var walk func(partial []string, connection string, visitedTwice bool)
	walk = func(partial []string, connection string, visitedTwice bool) {
		partial = append(partial, connection)
		if len(partial) != 0 && connection == "end" {
			paths = append(paths, append([]string{}, partial...))
			//fmt.Println(strings.Join(partial, ","))
			return
		}
		for _, con := range segments[connection] {
			if con == "start" {
				continue
			}
			//fmt.Printf("%v - %v - %v - %v\n", strings.Join(partial, ","), visitedTwice, twice, con)
			if strings.ToLower(con) == con && slices.Contains(partial, con) {
				if visitedTwice {
					continue
				}
				walk(partial, con, true)
				continue
			}
			walk(partial, con, visitedTwice)
		}
	}
	walk([]string{}, "start", false)
	return
}

func PathsPartOne(segments Segments) (paths []Path) {
	var walk func(partial []string, connection string)
	walk = func(partial []string, connection string) {
		partial = append(partial, connection)
		if len(partial) != 0 && connection == "end" {
			paths = append(paths, append([]string{}, partial...))
			return
		}
		for _, con := range segments[connection] {
			if strings.ToLower(con) == con && slices.Contains(partial, con) {
				continue
			}
			if con == "start" {
				continue
			}
			walk(partial, con)
		}
	}
	walk([]string{}, "start")
	return
}
