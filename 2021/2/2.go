package advent2

import "fmt"

func Travel(directions []Instruction) (distance int64, err error) {
	var horizontal, depth int64
	for _, d := range directions {
		switch d.Direction {
		case "forward":
			horizontal += d.Amount
		case "down":
			depth += d.Amount
		case "up":
			depth -= d.Amount
		default:
			err = fmt.Errorf("unkown direction: %v", d)
			return
		}
	}

	fmt.Printf("Horizontal Positon: %d, Depth: %d\n", horizontal, depth)

	distance = horizontal * depth
	fmt.Printf("Distance travelled: %d\n", distance)

	return
}

func Aim(directions []Instruction) (distance int64, err error) {
	var horizontal, depth, aim int64
	for _, d := range directions {
		switch d.Direction {
		case "forward":
			horizontal += d.Amount
			depth += aim * d.Amount
		case "down":
			aim += d.Amount
		case "up":
			aim -= d.Amount
		default:
			err = fmt.Errorf("unkown direction: %v", d)
			return
		}
	}

	fmt.Printf("Horizontal Positon: %d, Depth: %d\n", horizontal, depth)

	distance = horizontal * depth
	fmt.Printf("Distance travelled: %d\n", distance)

	return
}
