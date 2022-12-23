package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var input string

func main() {
	t := time.Now()

	part1()
	part2()

	fmt.Printf("gramps took %v\n", time.Since(t))
}

func part1() {
	fmt.Println("Part 1:")
	blueprints := parse(input)
	totalQuality, _, maxes := maxOpenGeodes(blueprints, 24)
	for i, max := range maxes {
		fmt.Printf("max: %v for %v\n", max, i)
	}
	fmt.Println()
	fmt.Printf("total quality: %v\n", totalQuality)
	fmt.Println()
}

func part2() {
	fmt.Println("Part 2:")
	blueprints := parse(input)
	blueprints = blueprints[:3]
	_, product, maxes := maxOpenGeodes(blueprints, 32)
	for i, max := range maxes {
		fmt.Printf("max: %v for %v\n", max, i)
	}
	fmt.Println()
	fmt.Printf("product of geodes: %v\n", product)
	fmt.Println()
}

type contents struct {
	ore, clay, obsidian, geode int
}

type blueprint struct {
	id       int
	ore      cost //eg. Each ore robot costs 4 ore.
	clay     cost
	obsidian cost
	geode    cost
}

type cost struct {
	contents
}

func parse(input string) (blueprints []blueprint) {
	var err error
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		print := blueprint{}

		split := strings.Split(line, ": ")
		if len(split) != 2 {
			log.Fatalf("error reading line: %s", line)
		}

		splitName := strings.Split(split[0], " ")
		if len(splitName) != 2 {
			log.Fatalf("error reading line: %s", line)
		}
		print.id, err = strconv.Atoi(splitName[1])
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}

		splitBots := strings.Split(split[1], ". ")
		if len(splitBots) != 4 {
			log.Fatalf("error reading line: %s", line)
		}

		print.ore.ore, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(
			splitBots[0], "Each ore robot costs "), " ore"))
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}

		print.clay.ore, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(
			splitBots[1], "Each clay robot costs "), " ore"))
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}

		splitCost := strings.Split(splitBots[2], " and ")
		print.obsidian.ore, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(
			splitCost[0], "Each obsidian robot costs "), " ore"))
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}
		print.obsidian.clay, err = strconv.Atoi(strings.TrimSuffix(splitCost[1], " clay"))
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}

		splitCost = strings.Split(splitBots[3], " and ")
		print.geode.ore, err = strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(
			splitCost[0], "Each geode robot costs "), " ore"))
		if err != nil {
			log.Fatalf("error reading line: %s", line)
		}
		print.geode.obsidian, err = strconv.Atoi(strings.TrimSuffix(splitCost[1], " obsidian."))
		if err != nil {
			log.Fatal(err)
		}

		blueprints = append(blueprints, print)
	}

	return
}

func maxOpenGeodes(blueprints []blueprint, timeCap int) (totalQuality int, product int, maxes []int) {
	for _, print := range blueprints {
		var geodes int
		for min := 1; min >= 0; min-- {
			geodes = openGeodes(print, timeCap, min)
			if geodes > 0 {
				break
			}
		}
		totalQuality += geodes * print.id
		maxes = append(maxes, geodes)
	}
	for _, max := range maxes {
		product *= max
	}
	return
}

func min(ints ...int) (min int) {
	min = math.MaxInt
	for _, num := range ints {
		if num < min {
			min = num
		}
	}
	return
}

func max(ints ...int) (max int) {
	for _, num := range ints {
		if num > max {
			max = num
		}
	}
	return
}

type q struct {
	resources contents
	robots    contents
	minute    int
}

func openGeodes(blueprint blueprint, timeCap int, minGeodes int) (maxGeodes int) {
	maxOre := maxOre(blueprint)
	queue := []q{{contents{}, contents{ore: 1}, 0}}
	visited := []q{}

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if slices.Contains(visited, current) {
			continue
		}
		visited = append(visited, current)

		if current.minute == timeCap {
			if current.resources.geode > maxGeodes {
				maxGeodes = current.resources.geode
			}
			continue
		}

		test, obs, geo := test(current, timeCap, blueprint)
		if test <= maxGeodes {
			continue
		}
		if test <= minGeodes { //this is sort of a hack to speed things up. min gets decremented if first pass didn't work
			continue
		}

		fmt.Printf("%+v, max: %v, bp: %v, test: %v, min: %v\n", current, maxGeodes, blueprint.id, test, minGeodes)

		next := current
		next.resources.ore += current.robots.ore // Each robot can collect 1 of its resource type per minute
		next.resources.clay += current.robots.clay
		next.resources.obsidian += current.robots.obsidian
		next.resources.geode += current.robots.geode
		next.minute++

		if current.resources.ore >= blueprint.geode.ore && current.resources.obsidian >= blueprint.geode.obsidian {
			next2 := next
			next2.robots.geode++
			next2.resources.ore -= blueprint.geode.ore
			next2.resources.obsidian -= blueprint.geode.obsidian
			queue = append(queue, next2)
			continue
		}

		queue = append(queue, next) //here we just increment the resources but don't build anything this round

		if current.robots.obsidian < blueprint.geode.obsidian &&
			geo > 0 { //it's only worth building an obs bot if we can later use it
			if current.resources.ore >= blueprint.obsidian.ore && current.resources.clay >= blueprint.obsidian.clay {
				next2 := next
				next2.robots.obsidian++
				next2.resources.ore -= blueprint.obsidian.ore
				next2.resources.clay -= blueprint.obsidian.clay
				queue = append(queue, next2)
			}
		}

		if current.robots.clay < blueprint.obsidian.clay && current.robots.obsidian < blueprint.geode.obsidian &&
			geo > 0 && obs > 0 { //it's only worth building a clay bot if we can later use it
			if current.resources.ore >= blueprint.clay.ore {
				next2 := next
				next2.robots.clay++
				next2.resources.ore -= blueprint.clay.ore
				queue = append(queue, next2)
			}
		}

		if current.robots.ore < maxOre { //the most bots we should ever build == the max ore we can ever use in one turn
			if current.resources.ore >= blueprint.ore.ore {
				next2 := next
				next2.robots.ore++
				next2.resources.ore -= blueprint.ore.ore
				queue = append(queue, next2)
			}
		}
	}

	return
}

// test upper bound of geodes made after full time has run
func test(test q, timeCap int, blueprint blueprint) (geodes, obs, geo int) {
	remaining := timeCap - test.minute

	ore := remaining - 1
	for i := 0; i < remaining; i++ {
		test.resources.ore += test.robots.ore + max(ore-i, 0)
	}

	clay := min(remaining-1, test.resources.ore/blueprint.clay.ore)
	for i := 0; i < remaining; i++ {
		test.resources.clay += test.robots.clay + max(clay-i, 0)
	}

	obs = min(remaining-1, test.resources.clay/blueprint.obsidian.clay, test.resources.ore/blueprint.obsidian.ore)
	for i := 0; i < remaining; i++ {
		test.resources.obsidian += test.robots.obsidian + max(obs-i, 0)
	}

	geo = min(remaining, test.resources.obsidian/blueprint.geode.obsidian, test.resources.ore/blueprint.geode.ore)
	for i := 0; i < remaining; i++ {
		test.resources.geode += test.robots.geode + max(geo-i, 0)
	}
	geodes = test.resources.geode

	return
}

// the most bots we should ever build == the max ore we can ever use in one turn
func maxOre(blueprint blueprint) (max int) {
	if blueprint.ore.ore > max {
		max = blueprint.ore.ore
	}
	if blueprint.clay.ore > max {
		max = blueprint.clay.ore
	}
	if blueprint.obsidian.ore > max {
		max = blueprint.obsidian.ore
	}
	if blueprint.geode.ore > max {
		max = blueprint.geode.ore
	}
	return
}
