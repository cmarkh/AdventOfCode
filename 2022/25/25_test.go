package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`

func TestPart1(t *testing.T) {
	sum := sum(toDecimal(parse(test1)))
	snafu := snafu(sum)
	fmt.Printf("sum: %v, snafu: %v\n", sum, snafu)
}

func TestPow(t *testing.T) {
	if pow(5, 2) != 25 {
		t.Fail()
	}
}

func TestDecimal(t *testing.T) {
	if decimal("1") != 1 {
		t.Fail()
	}
	if decimal("2") != 2 {
		t.Fail()
	}
	if decimal("1=") != 3 {
		t.Fail()
	}
	if decimal("1-") != 4 {
		t.Fail()
	}
	if decimal("20") != 10 {
		t.Fail()
	}
	if decimal("1=11-2") != 2022 {
		t.Fail()
	}
	if decimal("1-0---0") != 12345 {
		t.Fail()
	}
	if decimal("1121-1110-1=0") != 314159265 {
		t.Fail()
	}
}

func TestSnafu(t *testing.T) {
	fmt.Println(snafu(1))
	if snafu(1) != "1" {
		t.Fail()
	}
	fmt.Println(snafu(2))
	if snafu(2) != "2" {
		t.Fail()
	}
	fmt.Println(snafu(3))
	if snafu(3) != "1=" {
		t.Fail()
	}
	fmt.Println(snafu(4))
	if snafu(4) != "1-" {
		t.Fail()
	}
	fmt.Println(snafu(5))
	if snafu(5) != "10" {
		t.Fail()
	}
	fmt.Println(snafu(6))
	if snafu(6) != "11" {
		t.Fail()
	}
	fmt.Println(snafu(7))
	if snafu(7) != "12" {
		t.Fail()
	}
	fmt.Println(snafu(8))
	if snafu(8) != "2=" {
		t.Fail()
	}
	fmt.Println(snafu(9))
	if snafu(9) != "2-" {
		t.Fail()
	}
	fmt.Println(snafu(10))
	if snafu(10) != "20" {
		t.Fail()
	}
	fmt.Println(snafu(15))
	if snafu(15) != "1=0" {
		t.Fail()
	}
	fmt.Println(snafu(20))
	if snafu(20) != "1-0" {
		t.Fail()
	}
	fmt.Println(snafu(2022))
	if snafu(2022) != "1=11-2" {
		t.Fail()
	}
	fmt.Println(snafu(12345))
	if snafu(12345) != "1-0---0" {
		t.Fail()
	}
	fmt.Println(snafu(314159265))
	if snafu(314159265) != "1121-1110-1=0" {
		t.Fail()
	}
}
