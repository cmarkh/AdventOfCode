package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"

	advent "temp/adventofcode/go/2021"
)

var inputPath = "../input.txt"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	ages, err := Input(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	days := 200
	ages = ages.Increment(days)
	fmt.Printf("After %d turns, there are %d fish\n", days, ages.Count())
}

type Ages []int

func Input(path string) (ages Ages, err error) {
	lines, err := advent.ReadInput(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		strAges := strings.Split(line, ",")
		for _, a := range strAges {
			age, err := strconv.Atoi(a)
			if err != nil {
				return nil, err
			}
			ages = append(ages, age)
		}
	}

	if len(ages) < 2 {
		err = fmt.Errorf("possible issue reading input")
	}

	return
}

func (ages Ages) Increment(days int) Ages {
	for i := 0; i < days; i++ {
		for n, age := range ages {
			switch age {
			case 0:
				(ages)[n] = 6
				ages = append(ages, 8)
			default:
				(ages)[n]--
			}
		}
		fmt.Printf("day: %d, fish: %d\n", i, ages.Count())
	}
	return ages
}

func (ages *Ages) Count() (count int) {
	return len(*ages)
}

func (ages *Ages) Print() {
	for i, age := range *ages {
		if i < len(*ages)-1 {
			fmt.Printf("%d,", age)
		} else {
			fmt.Print(age)
		}

	}
	fmt.Println()
}
