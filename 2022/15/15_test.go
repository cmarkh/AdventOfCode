package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestPart1(t *testing.T) {
	sensors := parse(test1)
	for sensor, beacon := range sensors {
		fmt.Printf("sensor %v beacon is at %v\n", sensor, beacon)
	}
	count := noBeaconsCount(sensors, 10)
	fmt.Printf("%v positions where no beacon can be\n", count)
	fmt.Println()
}

func TestLineByLine1(t *testing.T) {
	sensors := parse(test1)
	pos := lineByLine(sensors, 20)
	fmt.Printf("beacon must be at %+v\n", pos)
}

func TestLineByLine2(t *testing.T) {
	sensors := parse(input)
	pos := lineByLine(sensors, 4000000)
	fmt.Printf("beacon must be at %+v\n", pos)
}

func TestPart2(t *testing.T) {
	sensors := parse(test1)
	pos := lineByLine(sensors, 20)
	fmt.Printf("beacon must be at %+v\n", pos)
	tuningFreq := tuningFrequency(pos)
	fmt.Printf("tuning frequency: %v\n", tuningFreq)
}
