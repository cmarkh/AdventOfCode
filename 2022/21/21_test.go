package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestPart1(t *testing.T) {
	monkies := parse(test1)
	for name, monkey := range monkies {
		fmt.Printf("%s: %+v\n", name, monkey)
	}
	fmt.Println()

	answer := solve(monkies)
	fmt.Printf("root solves to: %v\n", answer)
	fmt.Println()
}

func TestPart2(t *testing.T) {
	monkies := parse(test1)

	answer := solvePart2(monkies)
	fmt.Printf("I yell: %v\n", answer)
	fmt.Println()
}

func check(monk monkey, monkies monkies) int64 {
	if monk.specificNum {
		return monk.num
	}
	left := check(monkies[monk.left], monkies)
	right := check(monkies[monk.right], monkies)
	return operations[monk.operation](left, right)
}

func TestCheck(t *testing.T) {
	monkies := parse(test2)

	root := monkies["root"]
	root.operation = "="
	monkies["root"] = root

	human := monkies["humn"]
	human.num = 19
	monkies["humn"] = human

	answer := check(monkies["root"], monkies)
	fmt.Printf("I yell: %v\n", human.num)
	fmt.Printf("%v\n", answer)
	if answer == 1 {
		fmt.Println("and it worked")
	} else {
		t.Fail()
	}
	fmt.Println()
}

var test2 = `root: juli + josi
juli: amee + alex
amee: buki * abby
buki: 5
abby: 4
alex: 4
josi: benj / mark
benj: 360
mark: emly - humn
emly: 34
humn: 0`
