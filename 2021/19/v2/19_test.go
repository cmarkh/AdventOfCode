package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `--- scanner 0 ---
0,2
4,1
3,3

--- scanner 1 ---
-1,-1
-5,0
-2,1`

var test2 = `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`

var test3 = `--- scanner 0 ---
404,-588,-901

--- scanner 1 ---
686,422,578`

var test4 = `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7`

func Test2D(t *testing.T) {
	scans := parse(strings.Split(test1, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}

	fmt.Println()
}

func TestNegativePermutations(t *testing.T) {
	scans := parse(strings.Split(test2, "\n"))

	for _, negative := range negatives {
		fmt.Println(negative)
	}
	fmt.Println()

	negativePermutes := negativePermutations(scans[0].beacons)
	for _, beacon := range negativePermutes {
		fmt.Println(beacon)
	}
	fmt.Printf("length of negative permutations: %v\n", len(negativePermutes))
	fmt.Printf("original length of beacons: %v\n", len(scans[0].beacons))

	fmt.Println()
	fmt.Println("original beacons:")
	for _, beacon := range scans[0].beacons {
		fmt.Println(beacon)
	}

	fmt.Println()
}

func TestPermutations1(t *testing.T) {
	scans := parse(strings.Split(test2, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	permutes := scannerPermutations(scans[0].beacons)
	for _, beacon := range permutes {
		fmt.Println(beacon)
	}
	fmt.Printf("length of permutations: %v\n", len(permutes))
	fmt.Printf("original length of beacons: %v\n", len(scans[0].beacons))

	fmt.Println()
}

func TestPermutations2(t *testing.T) {
	scans := parse(strings.Split(test3, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scans = fillScannerPermutations(scans)
	for i, scanner := range scans {
		fmt.Printf("scanner %v\n", i)
		for _, beacon := range scanner.beaconPermutations {
			fmt.Println(beacon)
		}
		fmt.Println()
	}

	fmt.Println()
}

func TestPermutations3(t *testing.T) {
	scans := parse(strings.Split(test4, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scans = fillScannerPermutations(scans)
	for i, scanner := range scans {
		fmt.Printf("scanner %v\n", i)
		for _, beacon := range scanner.beaconPermutations {
			fmt.Println(beacon)
		}
		fmt.Println()
	}

	fmt.Println()
}

func TestOffsets2(t *testing.T) {
	scans := parse(strings.Split(test1, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	success, offsets := findOffsets(scans[0].beacons, scans[1].beacons, scans[0].offsets, 3)
	fmt.Printf("success: %v\n", success)
	fmt.Printf("offsets: x: %v, y: %v, z: %v\n", offsets.x, offsets.y, offsets.z)

	fmt.Println()
}

func TestOffsets3(t *testing.T) {
	scans := parse(strings.Split(test2, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scans = fillScannerPermutations(scans)

	findAllScannerOffsets(scans[:2], 12)

	fmt.Println()
}

func TestOffsets4(t *testing.T) {
	scanners := parse(strings.Split(test2, "\n"))
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

	fmt.Println()
}

func TestOffsets5(t *testing.T) {
	scans := parse(strings.Split(test2, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scans[1].offsets.x = 68
	scans[1].offsets.y = -1246
	scans[1].offsets.z = -43

	scans = fillScannerPermutations([]scanner{scans[1], scans[3]})

	findAllScannerOffsets(scans, 12)

	fmt.Println()
}

func TestOffsets6(t *testing.T) {
	scans := parse(strings.Split(test2, "\n"))
	for i, scan := range scans {
		fmt.Printf("scanner %d: %v\n", i, scan)
	}
	fmt.Println()

	scans = fillScannerPermutations([]scanner{scans[1], scans[3]})

	scans = findAllScannerOffsets(scans, 12)

	for _, beacon := range scans[1].beaconProperlyOriented {
		fmt.Println(beacon)
	}

	fmt.Println()
}

func TestBuildBeacons(t *testing.T) {
	scanners := parse(strings.Split(test2, "\n"))
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
}

func TestDistance(t *testing.T) {
	scanners := parse(strings.Split(test2, "\n"))
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

	d := greatestDistance(scanners)
	fmt.Printf("greatest distance between scanners: %v\n", d)

	fmt.Println()
}
