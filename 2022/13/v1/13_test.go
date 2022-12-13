package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

var check1 = map[int]bool{
	1: true,
	2: true,
	3: false,
	4: true,
	5: false,
	6: true,
	7: false,
	8: false,
}

func TestOrdered1(t *testing.T) {
	pairs := parse(test1)
	for i, pair := range pairs {
		ordered := ordered(pair[0], pair[1])
		fmt.Println(i + 1)
		fmt.Println(pair[0])
		fmt.Println(pair[1])
		fmt.Printf("ordered: %v\n", ordered)
		fmt.Println()
		if ordered != check1[i+1] {
			t.Errorf("%v should be %v\n", i+1, check1[i+1])
		}
	}
	fmt.Println()

}

func TestOrdered2(t *testing.T) {
	pairs := parse(input)
	i := 8
	for _, pair := range pairs[i : i+1] {
		ordered := ordered(pair[0], pair[1])
		fmt.Println(pair[0])
		fmt.Println(pair[1])
		fmt.Printf("ordered: %v\n", ordered)
		fmt.Println()
	}
	fmt.Println()

}

func TestSplit(t *testing.T) {
	split := split("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	//split := split("[[]]")
	for _, s := range split {
		fmt.Println(s)
	}
	fmt.Println()
}

func TestPart1(t *testing.T) {
	pairs := parse(test1)
	sum := sumOrderedIndices(pairs)
	fmt.Printf("sum of properly ordered pairs' indices: %v\n", sum)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	packets := parse2(test1)
	sort(packets)
	for _, packet := range packets {
		fmt.Println(packet)
	}
	fmt.Println()
	d1, d2, product := divisorIndices(packets)
	fmt.Printf("d1: %v, d2: %v, product: %v\n", d1, d2, product)
}
