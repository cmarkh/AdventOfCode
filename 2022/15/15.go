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
	sensors := parse(input)

	fmt.Println("Part 1:")
	count := noBeaconsCount(sensors, 2000000)
	fmt.Printf("%v positions where no beacon can be\n", count)
	fmt.Println()

	fmt.Println("Part 2:")
	pos := lineByLine(sensors, 4000000)
	fmt.Printf("beacon must be at %+v\n", pos)
	tuningFreq := tuningFrequency(pos)
	fmt.Printf("tuning frequency: %v\n", tuningFreq)
	fmt.Println()
}

type sensors map[position]beacon //map[sensor pos]closest beacon pos

type beacon struct {
	pos      position
	distance int
}

type position struct {
	x, y int
}

func parse(input string) (sens sensors) {
	var err error
	sens = make(sensors)
	lines := strings.Split(input, "\n")

	xy := func(str string) (pos position) {
		split := strings.Split(str, ", ")
		x := strings.TrimPrefix(split[0], "x=")
		y := strings.TrimPrefix(split[1], "y=")
		pos.x, err = strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		pos.y, err = strconv.Atoi(y)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ": ")
		sensor1 := strings.TrimPrefix(split[0], "Sensor at ")
		beacon1 := strings.TrimPrefix(split[1], "closest beacon is at ")
		sensorPos := xy(sensor1)
		beaconPos := xy(beacon1)
		sens[sensorPos] = beacon{beaconPos, distance(sensorPos, beaconPos)}
	}

	return
}

func noBeaconsCount(sensors sensors, row int) (count int) {
	type grid map[position]bool
	none := make(grid)

	for sensor, beacon := range sensors {
		beaconDistance := distance(sensor, beacon.pos)
		for r := sensor.y - beaconDistance; r <= sensor.y+beaconDistance; r++ {
			if r != row {
				continue
			}
			for c := sensor.x - beaconDistance; c <= sensor.x+beaconDistance; c++ {
				if distance(position{c, r}, sensor) > beaconDistance {
					continue
				}
				none[position{c, r}] = true
			}
		}
		none[beacon.pos] = false
	}

	for pos, empty := range none {
		if !empty {
			continue
		}
		if pos.y == row {
			count++
		}
	}
	return
}

func distance(pos1, pos2 position) int {
	return abs(pos1.x-pos2.x) + abs(pos1.y-pos2.y)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func lineByLine(sensors sensors, searchSpace int) (notSensed position) {
	type coverage struct {
		startX, endX int
	}

	for y := 0; y <= searchSpace; y++ {
		coverages := []coverage{}
		for sensor, beacon := range sensors {
			offset := abs(sensor.y - y)
			if offset > beacon.distance {
				continue
			}
			coverage := coverage{}
			coverage.startX = sensor.x - beacon.distance + offset
			coverage.endX = sensor.x + beacon.distance - offset
			coverages = append(coverages, coverage)
		}

		slices.SortStableFunc(coverages, func(a, b coverage) bool {
			return a.startX < b.startX
		})

		x := 0
		for _, coverage := range coverages {
			if coverage.startX > x {
				return position{x, y}
			}
			if coverage.endX+1 > x {
				x = coverage.endX + 1
			}
		}
		if x < searchSpace {
			return position{x, y}
		}
	}

	return
}

func tuningFrequency(pos position) int {
	return pos.x*4000000 + pos.y
}
