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

func TestExplode1(t *testing.T) {
	exploded, newPairs := explode("[[[[[9,8],1],2],3],4]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
	if newPairs != "[[[[0,9],2],3],4]" {
		t.Fail()
	}
}
func TestExplode2(t *testing.T) {
	exploded, newPairs := explode("[7,[6,[5,[4,[3,2]]]]]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
	if newPairs != "[7,[6,[5,[7,0]]]]" {
		t.Fail()
	}
}
func TestExplode3(t *testing.T) {
	exploded, newPairs := explode("[[6,[5,[4,[3,2]]]],1]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
	if newPairs != "[[6,[5,[7,0]]],3]" {
		t.Fail()
	}
}

func TestExplode4(t *testing.T) {
	exploded, newPairs := explode("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
	if newPairs != "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]" {
		t.Fail()
	}
}

func TestExplode5(t *testing.T) {
	exploded, newPairs := explode("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
	if newPairs != "[[3,[2,[8,0]]],[9,[5,[7,0]]]]" {
		t.Fail()
	}
}

func TestSplit1(t *testing.T) {
	fmt.Printf("old pairs: %s\n", "[[3,[2,[8,10]]],[9,[5,[7,0]]]]")
	split, newPairs := split("[[3,[2,[8,10]]],[9,[5,[7,0]]]]")
	if split {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
}

func TestSplit2(t *testing.T) {
	fmt.Printf("old pairs: %s\n", "[[3,[2,[11,8]]],[9,[5,[7,0]]]]")
	split, newPairs := split("[[3,[2,[11,8]]],[9,[5,[7,0]]]]")
	if split {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
}

func TestAdd(t *testing.T) {
	newPairs := add("[1,2]", "[[3,4],5]")
	fmt.Printf("new pairs: %s\n", newPairs)
	if newPairs != "[[1,2],[[3,4],5]]" {
		t.Fail()
	}
}

func TestAddReduce(t *testing.T) {
	newPairs := add("[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]")
	fmt.Printf("added: %s\n", newPairs)

	newPairs = reduce(newPairs)
	fmt.Printf("reduced: %s\n", newPairs)

	if newPairs != "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]" {
		t.Fail()
	}
}

func TestSumList(t *testing.T) {
	test := `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`

	fmt.Println(strings.Split(test, "\n"))

	answer := sumList(strings.Split(test, "\n"))
	fmt.Printf("answer: %s\n", answer)
	if answer != "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]" {
		t.Fail()
	}
}

func TestSplit3(t *testing.T) {
	fmt.Printf("old pairs: %s\n", "[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[10,[[11,9],[11,0]]]]")
	split, newPairs := split("[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[10,[[11,9],[11,0]]]]")
	if split {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
}

func TestExplode6(t *testing.T) {
	fmt.Printf("old pairs: %s\n", "[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[10,[[11,9],[11,0]]]]")
	exploded, newPairs := explode("[[[[4,0],[5,4]],[[7,7],[0,[6,7]]]],[10,[[11,9],[11,0]]]]")
	if exploded {
		fmt.Printf("new pairs: %s\n", newPairs)
	}
	fmt.Println()
}

func TestMagnitude1(t *testing.T) {
	mag := magnitude("[[9,1],[1,9]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 129 {
		t.Fail()
	}
}

func TestMagnitude3(t *testing.T) {
	mag := magnitude("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 1384 {
		t.Fail()
	}
}

func TestMagnitude4(t *testing.T) {
	mag := magnitude("[[[[1,1],[2,2]],[3,3]],[4,4]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 445 {
		t.Fail()
	}
}

func TestMagnitude5(t *testing.T) {
	mag := magnitude("[[1,2],[[3,4],5]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 143 {
		t.Fail()
	}
}

func TestMagnitude6(t *testing.T) {
	mag := magnitude("[[[[3,0],[5,3]],[4,4]],[5,5]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 791 {
		t.Fail()
	}
}

func TestMagnitude7(t *testing.T) {
	mag := magnitude("[[[[5,0],[7,4]],[5,5]],[6,6]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 1137 {
		t.Fail()
	}
}

func TestMagnitude8(t *testing.T) {
	mag := magnitude("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 3488 {
		t.Fail()
	}
}

func TestSumList2(t *testing.T) {
	test := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	answer := sumList(strings.Split(test, "\n"))
	fmt.Printf("answer: %s\n", answer)
	if answer != "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]" {
		t.Fail()
	}

	mag := magnitude(answer)
	fmt.Printf("magnitude: %v\n", mag)
	if mag != 4140 {
		t.Fail()
	}
}

func TestGreatestMagnitude(t *testing.T) {
	test := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	pair1, pair2, mag := greatestMagnitude(strings.Split(test, "\n"))

	fmt.Println("Greatest Magnitude is between:")
	fmt.Printf("Pair 1: %s\n", pair1)
	fmt.Printf("Pair 2: %s\n", pair2)
	fmt.Printf("Magnitude: %v\n", mag)

	if mag != 3993 {
		t.Fail()
	}
}
