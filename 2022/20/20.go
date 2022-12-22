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

var decryptionKey = 811589153

func main() {
	fmt.Println("Part 1:")
	file := parse(input)
	file = mix(file)
	_, sum, x, y, z := grove(file)
	fmt.Printf("%v, %v, %v\nsum: %v\n", x, y, z, sum)
	fmt.Println()

	fmt.Println("Part 2:")
	file = parse(input)
	file = mixPart2(file)
	_, sum, x, y, z = grove(file)
	fmt.Printf("%v, %v, %v\nsum: %v\n", x, y, z, sum)
	fmt.Println()
}

func parse(input string) (file []int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		file = append(file, num)
	}
	return
}

type num struct {
	num   int
	mixed bool
}

func mix(file []int) []int {
	nums := []num{}
	for _, n := range file {
		nums = append(nums, num{num: n})
	}

	done := false
	for !done {
		done = true
		for i, num := range nums {
			if !num.mixed {
				done = false
				num.mixed = true

				newI := (i + num.num) % (len(nums) - 1)
				if newI < 0 {
					newI = len(nums) - 1 + newI
					if newI >= len(nums) {
						log.Fatal("too big")
					}
				}
				//fmt.Printf("num: %v, oldI: %v, newI: %v\n", num.num, i, newI)

				if newI > i {
					for i := i + 1; i <= newI; i++ {
						nums[i-1] = nums[i]
					}
					nums[newI] = num
				}
				if newI == i {
					nums[newI] = num
				}
				if newI < i {
					for i := i - 1; i >= newI; i-- {
						nums[i+1] = nums[i]
					}
					nums[newI] = num
				}

				//printNums(nums)
				//fmt.Println()
				break
			}
		}
	}

	for i, num := range nums {
		file[i] = num.num
	}
	return file
}

func grove(file []int) (f []int, sum int, x, y, z int) {
	index0 := slices.Index(file, 0)
	x, y, z = file[(1000+index0)%(len(file))], file[(2000+index0)%(len(file))], file[(3000+index0)%(len(file))]
	sum = x + y + z
	return file, sum, x, y, z
}

func print(file []int) {
	for i, num := range file {
		fmt.Print(num)
		if i < len(file)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

//lint:ignore U1000 unused
func printNums(nums []num) {
	for i, num := range nums {
		fmt.Print(num.num)
		if i < len(nums)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func applyDecryptionKey(file []int) []int {
	for i := range file {
		file[i] *= decryptionKey
	}
	return file
}

type num2 struct {
	num      int
	mixOrder int
}

func mixPart2(file []int) []int {
	file = applyDecryptionKey(file)

	nums := []num2{}
	for i, n := range file {
		nums = append(nums, num2{n, i})
	}

	for l := 0; l < 10; l++ {
		mixNext := 0
		done := false
		for !done {
			done = true
			for i, num := range nums {
				if num.mixOrder == mixNext {
					done = false
					mixNext++

					newI := (i + num.num) % (len(nums) - 1)
					if newI < 0 {
						newI = len(nums) - 1 + newI
						if newI >= len(nums) {
							log.Fatal("too big")
						}
					}
					//fmt.Printf("num: %v, oldI: %v, newI: %v\n", num.num, i, newI)

					if newI > i {
						for i := i + 1; i <= newI; i++ {
							nums[i-1] = nums[i]
						}
						nums[newI] = num
					}
					if newI == i {
						nums[newI] = num
					}
					if newI < i {
						for i := i - 1; i >= newI; i-- {
							nums[i+1] = nums[i]
						}
						nums[newI] = num
					}

					//printNums(nums)
					//fmt.Println()
					break
				}
			}
		}
	}

	for i, num := range nums {
		file[i] = num.num
	}
	return file
}
