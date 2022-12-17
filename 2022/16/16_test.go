package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func TestPart1(t *testing.T) {
	valves := parse(test1)
	for name, valve := range valves {
		fmt.Printf("%v: %+v\n", name, valve)
	}
	fmt.Println()

	flow := openValves(valves)
	fmt.Printf("most flow: %v\n", flow)
	fmt.Println()
}

func TestStepsToReachValve(t *testing.T) {
	valves := parse(input)
	//steps := stepsToReachValve("AA", "HH", valves)
	//fmt.Println(steps)

	stepsToEachValve := stepsToReachValves(valves)
	fmt.Println(len(stepsToEachValve))
	for name := range stepsToEachValve {
		fmt.Println(len(stepsToEachValve[name]))
		for name2, steps := range stepsToEachValve[name] {
			fmt.Printf("%s to %s in %v steps\n", name, name2, steps)
		}
		fmt.Println()
	}
	fmt.Println()
}

var test2 = `Valve AA has flow rate=0; tunnels lead to valves BA
Valve BA has flow rate=2; tunnels lead to valves AA, CA
Valve CA has flow rate=4; tunnels lead to valves BA, DA
Valve DA has flow rate=6; tunnels lead to valves CA, EA
Valve EA has flow rate=8; tunnels lead to valves DA, FA
Valve FA has flow rate=10; tunnels lead to valves EA, GA
Valve GA has flow rate=12; tunnels lead to valves FA, HA
Valve HA has flow rate=14; tunnels lead to valves GA, IA
Valve IA has flow rate=16; tunnels lead to valves HA, JA
Valve JA has flow rate=18; tunnels lead to valves IA, KA
Valve KA has flow rate=20; tunnels lead to valves JA, LA
Valve LA has flow rate=22; tunnels lead to valves KA, MA
Valve MA has flow rate=24; tunnels lead to valves LA, NA
Valve NA has flow rate=26; tunnels lead to valves MA, OA
Valve OA has flow rate=28; tunnels lead to valves NA, PA
Valve PA has flow rate=30; tunnels lead to valves OA`

func Test2(t *testing.T) {
	valves := parse(test2)
	for name, valve := range valves {
		fmt.Printf("%v: %+v\n", name, valve)
	}
	fmt.Println()

	flow := openValves(valves)
	fmt.Printf("most flow: %v\n", flow)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	valves := parse(test2)

	flow := openValvesWithElephant2(valves)
	fmt.Printf("most flow: %v\n", flow)
	fmt.Println()
}
