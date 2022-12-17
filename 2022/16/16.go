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

	valves := parse(input)

	fmt.Println("Part 1:")
	flow := openValves(valves)
	fmt.Printf("most flow: %v\n", flow)
	fmt.Println()

	fmt.Println("Part 2:")
	flow = openValvesWithElephant2(valves)
	fmt.Printf("most flow: %v\n", flow)
	fmt.Println()

	fmt.Printf("I took %v\n", time.Since(t))
}

type valve struct {
	name    string
	flow    int
	tunnels []string
}

func parse(input string) (valves map[string]valve) {
	var err error
	valves = make(map[string]valve)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		valve := valve{}
		split := strings.Split(line, "; ")

		splitValve := strings.Split(split[0], " ")
		strRate := splitValve[len(splitValve)-1]
		strRate = strings.TrimPrefix(strRate, "rate=")
		valve.flow, err = strconv.Atoi(strRate)
		if err != nil {
			log.Fatal(err)
		}

		strTunnels := strings.TrimPrefix(split[1], "tunnels lead to valves ")
		strTunnels = strings.TrimPrefix(strTunnels, "tunnel leads to valve ")
		splitTunnels := strings.Split(strTunnels, ", ")
		valve.tunnels = append(valve.tunnels, splitTunnels...)

		valve.name = splitValve[1]
		valves[splitValve[1]] = valve
	}

	return
}

func stepsToReachValve(start, destination string, valves map[string]valve) (steps int) {
	var step func(start, destination string, visited []string, stepsSoFar int) (steps int)
	step = func(start, destination string, visited []string, stepsSoFar int) (steps int) {
		//fmt.Printf("start: %v, destination: %v\n", start, destination)
		if start == destination {
			return stepsSoFar
		}
		visited = append(visited, start)
		steps = math.MaxInt / 2
		for _, tunnel := range valves[start].tunnels {
			if slices.Contains(visited, tunnel) {
				continue
			}
			s := step(tunnel, destination, visited, stepsSoFar+1)
			if s < steps {
				steps = s
			}
		}
		return steps
	}
	return step(start, destination, []string{}, 0)
}

func stepsToReachValves(valves map[string]valve) (steps map[string]map[string]int) {
	steps = make(map[string]map[string]int)
	for name1 := range valves {
		steps[name1] = make(map[string]int)
		for name2 := range valves {
			steps[name1][name2] = stepsToReachValve(name1, name2, valves)
			//fmt.Printf("%s to %s - %v\n", name1, name2, steps[name1][name2])
		}
	}
	return
}

func openValves(valves map[string]valve) (totalFlow int) {
	stepsToEachValve := stepsToReachValves(valves)

	type partial struct {
		name    string
		opened  []string
		minutes int
		flow    int
	}
	queue := []partial{{"AA", []string{}, 30, 0}}

	valvesWithFlow := func() (count int) {
		for _, valve := range valves {
			if valve.flow > 0 {
				count++
			}
		}
		return
	}()

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if len(current.opened) == valvesWithFlow || current.minutes <= 0 {
			if current.flow > totalFlow {
				totalFlow = current.flow
				/*fmt.Println(current.opened)
				fmt.Println(current.minutes)
				fmt.Println(current.flow)
				fmt.Println()*/
			}
			continue
		}

		for name, steps := range stepsToEachValve[current.name] {
			if slices.Contains(current.opened, name) {
				continue
			}
			if steps+1 > current.minutes {
				if current.flow > totalFlow {
					totalFlow = current.flow
					/*fmt.Println(current.opened)
					fmt.Println(current.minutes)
					fmt.Println(current.flow)
					fmt.Println()*/
				}
				continue
			}
			if valves[name].flow == 0 {
				continue
			}

			next := partial{}
			next.name = name
			next.minutes = current.minutes - steps - 1
			next.flow = current.flow + valves[name].flow*next.minutes

			next.opened = make([]string, len(current.opened))
			copy(next.opened, current.opened)
			next.opened = append(next.opened, name)

			queue = append(queue, next)
		}

	}

	return
}

func openValvesWithElephant2(valves map[string]valve) (totalFlow int) {
	stepsToEachValve := stepsToReachValves(valves)

	type partial struct {
		me, ele struct {
			name    string
			minutes int
		}
		opened []string
		flow   int
	}
	queue := []partial{}

	p := partial{}
	p.me.name = "AA"
	p.ele.name = "AA"
	p.me.minutes = 26
	p.ele.minutes = 26

	queue = append(queue, p)

	valvesWithFlow := func() (count int) {
		for _, valve := range valves {
			if valve.flow > 0 {
				count++
			}
		}
		return
	}()

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if len(current.opened) == valvesWithFlow || (current.me.minutes <= 0 && current.ele.minutes <= 0) {
			if current.flow > totalFlow {
				totalFlow = current.flow
				fmt.Println(current.opened)
				fmt.Println(current.flow)
				fmt.Println()
			}
			continue
		}

		for name, steps := range stepsToEachValve[current.me.name] {
			if slices.Contains(current.opened, name) {
				continue
			}
			if steps+1 > current.me.minutes {
				if current.flow > totalFlow {
					totalFlow = current.flow
					fmt.Println(current.opened)
					fmt.Println(current.flow)
					fmt.Println()
				}
				continue
			}
			if valves[name].flow == 0 {
				continue
			}

			next := current
			next.me.name = name
			next.me.minutes = current.me.minutes - steps - 1
			next.flow = current.flow + valves[name].flow*next.me.minutes

			next.opened = make([]string, len(current.opened))
			copy(next.opened, current.opened)
			next.opened = append(next.opened, name)

			for name, steps := range stepsToEachValve[next.ele.name] {
				if slices.Contains(next.opened, name) {
					continue
				}
				if steps+1 > next.ele.minutes {
					if next.flow > totalFlow {
						totalFlow = next.flow
						fmt.Println(next.opened)
						fmt.Println(next.flow)
						fmt.Println()
					}
					continue
				}
				if valves[name].flow == 0 {
					continue
				}

				next2 := next
				next2.ele.name = name
				next2.ele.minutes = next.ele.minutes - steps - 1
				next2.flow = next.flow + valves[name].flow*next2.ele.minutes

				next2.opened = make([]string, len(next.opened))
				copy(next2.opened, next.opened)
				next2.opened = append(next2.opened, name)

				queue = append(queue, next2)
			}

		}
	}

	return
}
