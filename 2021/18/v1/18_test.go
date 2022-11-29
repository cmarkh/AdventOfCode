package main

import (
	"fmt"
	"strings"
	"testing"
)

func testAdd(input string, t *testing.T) {
	//fmt.Println(len(strings.Split(input, "\n")))
	Add(strings.Split(input, "\n"))
}

func TestAdd1(t *testing.T) {
	testAdd(`[1,2]
[[1,2],3]
[9,[8,7]]
[[1,9],[8,5]]
[[[[1,2],[3,4]],[[5,6],[7,8]]],9]
[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]`, t)
}

func TestAdd2(t *testing.T) {
	testAdd(`[1,2]
[[3,4],5]`, t)
}

func TestBuildPair1(t *testing.T) {
	pair := buildPair("[[1,2],3]")
	fmt.Printf("top: %+v\n", pair)
	fmt.Printf("x: %+v\n", pair.childX)
	fmt.Printf("x child: %+v\n", pair.childX.childX)
	fmt.Printf("y: %+v\n", pair.childY)
	//pair.Print()
}

func TestBuildPair2(t *testing.T) {
	pair := buildPair("[[1,2],3]")
	fmt.Printf("top: %+v\n", pair)
	pair.Print()
	fmt.Println()
}

func TestBuildPair3(t *testing.T) {
	pair := buildPair("[[[[1,2],[3,4]],[[5,6],[7,8]]],9]")
	fmt.Printf("top: %+v\n", pair)
	pair.Print()
	fmt.Println()
}

func TestExplode1(t *testing.T) {
	pair := buildPair("[[[[[9,8],1],2],3],4]")
	pair.Print()
	fmt.Println()
	pair.explode(1)
}
