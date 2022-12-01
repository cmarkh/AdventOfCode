package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func main() {
	t := time.Now()

	fmt.Println("Part 1:")

	lines, err := readInput("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanners := parse(lines)
	for i, scan := range scanners {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scanners = fillScannerPermutations(scanners)
	scanners = findAllScannerOffsets(scanners, 12)

	for i := range scanners {
		fmt.Printf("scanner %v:\n", i)
		fmt.Printf("x: %v, y: %v, z: %v\n", scanners[i].offsets.x, scanners[i].offsets.y, scanners[i].offsets.z)
		fmt.Printf("matched with: %v\n", scanners[i].matchedWith)
		fmt.Println()
	}

	beacons := buildBeacons(scanners)
	for _, beacon := range beacons {
		fmt.Println(beacon)
	}
	fmt.Printf("%v beacons\n", len(beacons))

	fmt.Println()

	fmt.Println("Part 2:")

	d := greatestDistance(scanners)
	fmt.Printf("greatest distance between scanners: %v\n", d)

	fmt.Printf("sir slowness took %v\n", time.Since(t))
}

type scanner struct {
	beacons                []beacon
	beaconPermutations     [][]beacon
	beaconProperlyOriented []beacon
	offsets                //offset from scanner 0's position
	matched                bool
	matchedWith            int
}
type beacon [3]int
type offsets struct {
	x, y, z int
}

func readInput(path string) (lines []string, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(content), "\n")
	return
}

func parse(lines []string) (scanners []scanner) {
	beacons := []beacon{}
	var err error

	for _, line := range lines {
		if line == "" {
			scanners = append(scanners, scanner{beacons, [][]beacon{}, []beacon{}, offsets{}, false, 0})
			continue
		}
		if strings.Contains(line, "scanner") {
			beacons = []beacon{}
			continue
		}
		beacon := beacon{}

		coords := strings.Split(line, ",")
		beacon[0], err = strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}
		beacon[1], err = strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}
		if len(coords) > 2 {
			beacon[2], err = strconv.Atoi(coords[2])
			if err != nil {
				log.Fatal(err)
			}
		}

		beacons = append(beacons, beacon)

	}
	scanners = append(scanners, scanner{beacons, [][]beacon{}, []beacon{}, offsets{}, false, 0})
	return
}

func fillScannerPermutations(scanners []scanner) []scanner {
	for i, scanner := range scanners {
		scanners[i].beaconPermutations = scannerPermutations(scanner.beacons)
	}
	return scanners
}

func scannerPermutations(beacons []beacon) (permutations [][]beacon) {
	swap := func(attempt []beacon, a, b int) []beacon {
		for i, beacon := range attempt {
			attempt[i][a] = beacon[b]
			attempt[i][b] = beacon[a]
		}
		return attempt
	}

	var heap func(attempt []beacon, n int)
	heap = func(attempt []beacon, n int) {
		if n == 1 {
			new := make([]beacon, len(attempt))
			copy(new, attempt)
			permutations = append(permutations, new)
			return
		}
		for i := 0; i < n-1; i++ {
			heap(attempt, n-1)
			if n%2 == 0 {
				attempt = swap(attempt, i, n-1)
			} else {
				attempt = swap(attempt, 0, n-1)
			}
		}
		heap(attempt, n-1)
	}
	for _, negativePermutation := range negativePermutations(beacons) {
		heap(negativePermutation, len(negatives[0]))
	}

	return
}

func negativePermutations(beacons []beacon) (permutations [][]beacon) {
	for _, negative := range negatives {
		negativeBeacons := []beacon{}
		for _, beacon := range beacons {
			for n, num := range negative {
				beacon[n] *= num
			}
			negativeBeacons = append(negativeBeacons, beacon)
		}
		permutations = append(permutations, negativeBeacons)
	}
	return
}

var negatives = func(length int) (permutations [][]int) {
	var bases [][]int
	base := make([]int, length)
	for i := 0; i < length; i++ {
		base[i] = 1
	}
	bases = append(bases, base)

	for n := 1; n <= length; n++ {
		new := make([]int, length)
		copy(new, base)
		for i := 0; i < n; i++ {
			new[i] *= -1
		}
		bases = append(bases, new)
	}

	swap := func(arr []int, a, b int) []int {
		temp := arr[a]
		arr[a] = arr[b]
		arr[b] = temp
		return arr
	}

	shouldSwap := func(arr []int, start, curr int) bool {
		for i := start; i < curr; i++ {
			if arr[i] == arr[curr] {
				return false
			}
		}
		return true
	}

	for _, base := range bases {
		var permute func(arr []int, index, n int)
		permute = func(arr []int, index, n int) {
			if index >= n {
				new := make([]int, len(arr))
				copy(new, arr)
				permutations = append(permutations, new)
				return
			}
			for i := index; i < n; i++ {
				if !shouldSwap(arr, index, i) {
					continue
				}
				arr = swap(arr, index, i)
				permute(arr, index+1, n)
				arr = swap(arr, index, i)
			}
		}
		permute(base, 0, len(base))
	}

	return
}(3)

func findOffsets(beacons1, beacons2 []beacon, offsets1 offsets, matchesNeeded int) (success bool, offsets2 offsets) {
	// sensor can max see 1000 away
	// so 2 sensors can be max 2000 away from each other (each 1000 from the same beacon)
	maxDistance := 2000

	possibleX := []int{}
	for x := -maxDistance + offsets1.x; x < maxDistance+offsets1.x; x++ {
		matches := 0
		for _, beacon1 := range beacons1 {
			for _, beacon2 := range beacons2 {
				if beacon1[0]+offsets1.x == beacon2[0]+x {
					matches++
					if matches >= matchesNeeded {
						possibleX = append(possibleX, x)
					}
				}
			}
		}
	}
	possibleY := []int{}
	for y := -maxDistance + offsets1.y; y < maxDistance+offsets1.y; y++ {
		matches := 0
		for _, beacon1 := range beacons1 {
			for _, beacon2 := range beacons2 {
				if beacon1[1]+offsets1.y == beacon2[1]+y {
					matches++
					if matches >= matchesNeeded {
						possibleY = append(possibleY, y)
					}
				}
			}
		}
	}
	possibleZ := []int{}
	for z := -maxDistance + offsets1.z; z < maxDistance+offsets1.z; z++ {
		matches := 0
		for _, beacon1 := range beacons1 {
			for _, beacon2 := range beacons2 {
				if beacon1[2]+offsets1.z == beacon2[2]+z {
					matches++
					if matches >= matchesNeeded {
						possibleZ = append(possibleZ, z)
					}
				}
			}
		}
	}

	/*fmt.Printf("possible x: %v\n", possibleX)
	fmt.Printf("possible y: %v\n", possibleY)
	fmt.Printf("possible z: %v\n", possibleZ)
	fmt.Println()*/

	matches := 0
	func() {
		for _, x := range possibleX {
			for _, y := range possibleY {
				for _, z := range possibleZ {
					matches = 0 //need to reset when offset changes
					for _, beacon1 := range beacons1 {
						for _, beacon2 := range beacons2 {
							if beacon1[0]+offsets1.x != beacon2[0]+x {
								continue
							}
							if beacon1[1]+offsets1.y != beacon2[1]+y {
								continue
							}
							if beacon1[2]+offsets1.z != beacon2[2]+z {
								continue
							}
							matches++
							if matches >= matchesNeeded {
								offsets2.x = x
								offsets2.y = y
								offsets2.z = z
								success = true
								return
							}
						}
					}
				}
			}
		}
	}()

	return
}

func findAllScannerOffsets(scanners []scanner, matchesNeeded int) []scanner {
	scanners[0].beaconProperlyOriented = make([]beacon, len(scanners[0].beacons))
	copy(scanners[0].beaconProperlyOriented, scanners[0].beacons)
	scanners[0].matched = true

	matches := 1
	newMatches := []int{}
	newMatches2 := []int{0}
	for {
		if matches == len(scanners) {
			return scanners
		}
		newMatches = newMatches2
		newMatches2 = []int{}
		if !func() bool {
			//test each scanner for overlap with the other scanner
			for _, newMatch := range newMatches {
				for i, scanner := range scanners {
					if scanner.matched {
						continue
					}
					if i == newMatch {
						continue //same scanner
					}
					for _, beacons2 := range scanner.beaconPermutations {
						success, offsets := findOffsets(scanners[newMatch].beaconProperlyOriented, beacons2, scanners[newMatch].offsets, matchesNeeded)
						if success {
							scanners[i].beaconProperlyOriented = make([]beacon, len(beacons2))
							copy(scanners[i].beaconProperlyOriented, beacons2)
							scanners[i].offsets = offsets
							scanners[i].matchedWith = newMatch
							scanners[i].matched = true

							newMatches2 = append(newMatches2, i)
							matches++

							fmt.Printf("scanner %v:\n", i)
							fmt.Printf("x: %v, y: %v, z: %v\n", scanners[i].offsets.x, scanners[i].offsets.y, scanners[i].offsets.z)
							fmt.Printf("matched with: %v\n", scanners[i].matchedWith)
							fmt.Println()
						}
					}
				}
			}
			return len(newMatches2) > 0
		}() {
			log.Fatal("failure")
		}
	}
}

func buildBeacons(scanners []scanner) (unique []beacon) {
	unique = make([]beacon, len(scanners[0].beaconProperlyOriented))
	copy(unique, scanners[0].beaconProperlyOriented)

	for _, scanner := range scanners {
		for _, b := range scanner.beaconProperlyOriented {
			b[0] += scanner.offsets.x
			b[1] += scanner.offsets.y
			b[2] += scanner.offsets.z

			if !slices.Contains(unique, b) {
				unique = append(unique, b)
			}
		}
	}

	return
}

func greatestDistance(scanners []scanner) (distance int) {
	abs := func(n int) int {
		if n < 0 {
			return n * -1
		}
		return n
	}

	for m, scanner1 := range scanners {
		for n, scanner2 := range scanners {
			if m == n {
				continue
			}
			d := abs(scanner2.offsets.x-scanner1.offsets.x) +
				abs(scanner2.offsets.y-scanner1.offsets.y) +
				abs(scanner2.offsets.z-scanner1.offsets.z)
			if d > distance {
				distance = d
			}
		}
	}

	return
}
