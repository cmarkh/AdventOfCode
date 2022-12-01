package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

}

type scanner struct {
	beacons     []beacon
	offsets     //each offset is relative to scanner 0
	matched     bool
	matchedWith int
}
type offsets struct {
	x, y, z int
}
type beacon [3]int

func parse(lines []string) (scanners []scanner) {
	beacons := []beacon{}
	var err error

	for _, line := range lines {
		if line == "" {
			scanners = append(scanners, scanner{beacons, offsets{}, false, 0})
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
	scanners = append(scanners, scanner{beacons, offsets{}, false, 0}) //append last scanner
	return
}

func rotateScanner(scanner1, scanner2 scanner, matchesNeeded int) (success bool, newScanner2 scanner) {
	negativePermutations := negativePermutations(len(scanner2.beacons[0]))

	swap := func(attempt scanner, a, b int) scanner {
		for i, beacon := range attempt.beacons {
			attempt.beacons[i][a] = beacon[b]
			attempt.beacons[i][b] = beacon[a]
		}
		return attempt
	}

	var heap func(attempt scanner, n int)
	heap = func(attempt scanner, n int) {
		if success {
			return
		}
		if n == 1 {
			for _, negatives := range negativePermutations {
				negatedAttemp := scanner{}
				negatedAttemp.beacons = make([]beacon, len(attempt.beacons))
				copy(negatedAttemp.beacons, attempt.beacons)

				for i := range attempt.beacons {
					for n := range negatives {
						negatedAttemp.beacons[i][n] *= negatives[n]
					}
				}
				if found, offsets := findOffsets(scanner1, negatedAttemp, matchesNeeded); found {
					negatedAttemp.offsets = offsets
					newScanner2 = negatedAttemp
					success = found
					return
				}
			}
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
	heap(scanner2, len(scanner2.beacons[0]))
	return
}

func negativePermutations(length int) (permutations [][]int) {
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
}

func findOffsets(scanner1, scanner2 scanner, matchesNeeded int) (success bool, offsets offsets) {
	// sensor can max see 1000 away
	// so 2 sensors can be max 2000 away from each other (each 1000 from the same beacon)
	maxDistance := 2000

	possibleX := []int{}
	for x := -maxDistance; x < maxDistance; x++ {
		matches := 0
		for _, beacon1 := range scanner1.beacons {
			for _, beacon2 := range scanner2.beacons {
				if beacon1[0]+scanner1.offsets.x == beacon2[0]+x {
					if x == -92 {
						fmt.Println("here")
						fmt.Println(beacon2[0])
					}
					matches++
					if matches >= matchesNeeded {
						possibleX = append(possibleX, x)
					}
				}
			}
		}
	}
	possibleY := []int{}
	for y := -maxDistance; y < maxDistance; y++ {
		matches := 0
		for _, beacon1 := range scanner1.beacons {
			for _, beacon2 := range scanner2.beacons {
				if beacon1[1]+scanner1.offsets.y == beacon2[1]+y {
					matches++
					if matches >= matchesNeeded {
						possibleY = append(possibleY, y)
					}
				}
			}
		}
	}
	possibleZ := []int{}
	for z := -maxDistance; z < maxDistance; z++ {
		matches := 0
		for _, beacon1 := range scanner1.beacons {
			for _, beacon2 := range scanner2.beacons {
				if beacon1[2]+scanner1.offsets.z == beacon2[2]+z {
					matches++
					if matches >= matchesNeeded {
						possibleZ = append(possibleZ, z)
					}
				}
			}
		}
	}

	/* fmt.Printf("possible x: %v\n", possibleX)
	fmt.Printf("possible y: %v\n", possibleY)
	fmt.Printf("possible z: %v\n", possibleZ)
	fmt.Println() */

	matches := 0
	func() {
		for _, x := range possibleX {
			for _, y := range possibleY {
				for _, z := range possibleZ {
					matches = 0 //need to reset when offset changes
					for _, beacon1 := range scanner1.beacons {
						for _, beacon2 := range scanner2.beacons {
							if beacon1[0]+scanner1.offsets.x != beacon2[0]+x {
								continue
							}
							if beacon1[1]+scanner1.offsets.y != beacon2[1]+y {
								continue
							}
							if beacon1[2]+scanner1.offsets.z != beacon2[2]+z {
								continue
							}
							matches++
							if matches >= matchesNeeded {
								offsets.x = x
								offsets.y = y
								offsets.z = z
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
	scanners[0].matched = true
	matches := 0
	for {
		if matches == len(scanners) {
			return scanners
		}
		if !func() bool {
			//test each scanner for overlap with the other scanners
			for i := 1; i < len(scanners); i++ {
				if scanners[i].matched {
					continue
				}
				for j := 0; j < len(scanners); j++ {
					if i == j {
						continue
					}
					if !scanners[j].matched {
						continue
					}
					if i == 3 {
						fmt.Println(j)
						fmt.Println(scanners[1].offsets)
					}
					var scanner scanner
					success, scanner := rotateScanner(scanners[j], scanners[i], matchesNeeded)
					if success {
						scanners[i] = scanner
						scanners[i].matched = true
						scanners[i].matchedWith = j

						fmt.Printf("scanner %v:\n", i)
						fmt.Printf("x: %v, y: %v, z: %v\n", scanners[i].offsets.x, scanners[i].offsets.y, scanners[i].offsets.z)
						fmt.Printf("matched with: %v\n", scanners[i].matchedWith)
						fmt.Println()
						return true
					}
				}
			}
			return false
		}() {
			log.Fatal("failure")
		}
	}
}
