package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	advent "temp/adventofcode/go/2021"
)

func main() {
	line, err := input()
	if err != nil {
		log.Fatal(err)
	}
	target, err := ParseInput(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", target)

	fmt.Println("Part 1:")
	maxH := MaxHeight(target)
	maxH.path.Print(target)
	fmt.Println()
	fmt.Printf("max height: %v\n", maxH.height)
	fmt.Printf("velocity: %v, %v\n", maxH.velocity.x, maxH.velocity.y)
	fmt.Printf("hit target at %v\n", maxH.hitLocation)
	fmt.Println()

	fmt.Println("Part 2:")
	velocities, count := AllValidShots(target)
	fmt.Println("Valid velocities:")
	for _, v := range velocities {
		fmt.Println(v)
	}
	fmt.Printf("\nNumber of valid velocities: %v\n", count)
}

type TargetArea struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func input() (line string, err error) {
	lines, err := advent.ReadInput("input.txt")
	if err != nil {
		return
	}
	if len(lines) != 1 {
		err = fmt.Errorf("expected just one line of input. got: %v", lines)
		return
	}
	return lines[0], nil
}

func ParseInput(input string) (area TargetArea, err error) {
	splitXY := strings.Split(strings.TrimPrefix(input, "target area: "), ", ")
	if len(splitXY) != 2 {
		err = fmt.Errorf("failed to parse target area: %v", input)
		return
	}

	splitX := strings.Split(strings.TrimPrefix(splitXY[0], "x="), "..")
	if len(splitX) != 2 {
		err = fmt.Errorf("failed to parse target area: %v", input)
		return
	}
	area.xMin, err = strconv.Atoi(splitX[0])
	if err != nil {
		return
	}
	area.xMax, err = strconv.Atoi(splitX[1])
	if err != nil {
		return
	}

	splitY := strings.Split(strings.TrimPrefix(splitXY[1], "y="), "..")
	if len(splitY) != 2 {
		err = fmt.Errorf("failed to parse target area: %v", input)
		return
	}
	area.yMin, err = strconv.Atoi(splitY[0])
	if err != nil {
		return
	}
	area.yMax, err = strconv.Atoi(splitY[1])
	if err != nil {
		return
	}

	return
}

type Position struct {
	x int
	y int
}
type Velocity struct {
	x int
	y int
}
type Path []Position

func Launch(target TargetArea, velocity Velocity) (hitTarget bool, long bool, hitPosition Position, path Path) {
	probe := Position{}
	path = append(path, probe) //starting position

	for i := 0; ; i++ {
		probe.x += velocity.x
		probe.y += velocity.y

		path = append(path, probe)

		if velocity.x > 0 {
			velocity.x--
		} else if velocity.x < 0 {
			velocity.x++
		}

		velocity.y--

		if probe.withinTarget(target) {
			return true, false, probe, path
		}
		if missed, long := probe.missed(target); missed {
			return false, long, hitPosition, path
		}
	}
}

func (probe Position) withinTarget(target TargetArea) bool {
	if probe.x < target.xMin || probe.x > target.xMax {
		return false
	}
	if probe.y < target.yMin || probe.y > target.yMax {
		return false
	}
	return true
}

func (probe Position) missed(target TargetArea) (missed bool, long bool) {
	if probe.x > target.xMax {
		return true, true
	}
	if probe.y < target.yMin {
		return true, false
	}
	return false, false
}

func (path Path) inPath(point Position) bool {
	for _, p := range path {
		if point.x == p.x && point.y == p.y {
			return true
		}
	}
	return false
}

func (path Path) Print(tgt TargetArea) {
	var maxX, maxY, minY int
	for _, point := range path {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.y < minY {
			minY = point.y
		}
	}
	if maxX < tgt.xMax {
		maxX = tgt.xMax
	}
	if maxY < tgt.yMax {
		maxY = tgt.yMax
	}
	if minY > tgt.yMin {
		minY = tgt.yMin
	}

	for r := maxY; r >= minY; r-- {
		for c := 0; c <= maxX; c++ {
			if r == 0 && c == 0 {
				fmt.Print("S")
				continue
			}
			if path.inPath(Position{c, r}) {
				fmt.Print("#")
				continue
			}
			if (Position{c, r}).withinTarget(tgt) {
				fmt.Print("T")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

type MaxH struct {
	velocity    Velocity
	height      int
	hitLocation Position
	path        Path
}

// Turns out that the y velocity that achieves the greatest height is always target area minY-1
func MaxHeight(target TargetArea) (maxH MaxH) {
	yVelocity := absoluteValue(target.yMin) - 1

	for vX := 1; vX <= target.xMax; vX++ {
		//fmt.Printf("testing %v, %v\n", vX, vY)
		hit, _, hitPosition, path := Launch(target, Velocity{vX, yVelocity})
		if hit {
			maxH.height = path.maxHeight()
			fmt.Println(yVelocity)
			maxH.velocity = Velocity{vX, yVelocity}
			maxH.hitLocation = hitPosition
			maxH.path = path
			return
		}
	}
	return //should never reach here
}

func absoluteValue(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func MaxHeightOLD(target TargetArea) (maxH MaxH) {
	//target x min can't be 0 or y velocity would go toward infinity
	for vY := 0; vY < 1000; vY++ {
		var hit, long bool
		var hitPosition Position
		var path Path
		var vX int

		//looking for minimum x velocity that manages to hit the target without going long
		for vX = 1; vX <= target.xMax; vX++ {
			//fmt.Printf("testing %v, %v\n", vX, vY)
			hit, long, hitPosition, path = Launch(target, Velocity{vX, vY})
			if hit {
				if path.maxHeight() > maxH.height {
					maxH.height = path.maxHeight()
					maxH.velocity = Velocity{vX, vY}
					maxH.hitLocation = hitPosition
					maxH.path = path
				}
				break
			}
			if long {
				break
			}
		}

		//if !hit {
		//	return //return once can no longer hit target at any x velocity
		//}
	}
	return
}

func (path Path) maxHeight() (height int) {
	//starting position is 0 so height intial value 0 doesn't need to be checked
	for _, p := range path {
		if p.y > height {
			height = p.y
		}
	}
	return
}

func AllValidShots(target TargetArea) (velocities []Velocity, count int) {
	for vY := absoluteValue(target.yMin) * -1; vY < absoluteValue(target.yMin); vY++ {
		for vX := 1; vX <= target.xMax; vX++ {
			hit, _, _, _ := Launch(target, Velocity{vX, vY})
			if hit {
				velocities = append(velocities, Velocity{vX, vY})
			}
		}
	}
	return velocities, len(velocities)
}
