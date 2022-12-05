package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//lint:ignore U1000 ignore unused
var test1 = `inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`

func TestZVal(t *testing.T) {
	strInstructions := sanitize(input)
	sets := splitInstructions(strInstructions)

	for _, instruction := range sets[13] {
		fmt.Println(instruction)
	}
	fmt.Println()

	for i := 1; i <= 9; i++ {
		for b := 0; b < 18; b++ {
			z, valid := zValue(sets[0], i, b)
			if z == 0 && valid {
				fmt.Printf("z: %v, model: %v\n", b, i)
				//fmt.Printf("z: %v, valid: %v\n", z, valid)
				//return
			}
		}
	}

	fmt.Println()
}

func TestZVal2(t *testing.T) {
	strInstructions := sanitize(input)
	sets := splitInstructions(strInstructions)

	for i := 1; i <= 9; i++ {
		z, valid := zValue(sets[0], i, 0)
		if valid {
			fmt.Printf("z: %v, model: %v\n", z, i)
			//fmt.Printf("z: %v, valid: %v\n", z, valid)
			//return
		}
	}

	fmt.Println()
}

func TestZVal3(t *testing.T) {
	strInstructions := sanitize(input)
	sets := splitInstructions(strInstructions)

	z, valid := zValue(sets[5], 3, 2820)
	//fmt.Printf("z: %v, model: %v\n", z, i)
	fmt.Printf("z: %v, valid: %v\n", z, valid)
	//return

	fmt.Println()
}

func TestZFormulas(t *testing.T) {
	strInstructions := sanitize(input)
	sets := splitInstructions(strInstructions)

	inputZ := 0
	for s := 0; s < 14; s++ {
		zFormula := zValueFormulas(s, 3, inputZ)
		zInstructions, _ := zValue(sets[s], 3, inputZ)
		if zFormula != zInstructions {
			fmt.Println("uh oh")
		}
		inputZ = zFormula
	}

}

func TestZFormulas2(t *testing.T) {
	for i := 1; i <= 9; i++ {
		zFormula := zValueFormulas(0, i, 0)
		fmt.Println(zFormula)
	}

}

func TestZFormulas3(t *testing.T) {
	inputZ := 0
	for s := 0; s < 14; s++ {
		zFormula := zValueFormulas(s, 4, inputZ)
		fmt.Println(zFormula)
		inputZ = zFormula
	}
}

func TestZFormulas4(t *testing.T) {
	for d := 1; d <= 9; d++ {
		for z := 0; z < 100000; z++ {
			zFormula := zValueFormulas(13, d, z)
			if zFormula == 0 {
				fmt.Printf("d: %v, inZ: %v, outZ: %v\n", d, z, zFormula)
			}
		}
	}
}

func TestZFormulas5(t *testing.T) {
	zFormula := zValueFormulas(13, 3, 11)
	fmt.Println(zFormula)
}

func TestZValue4(t *testing.T) {
	strInstructions := sanitize(input)
	sets := splitInstructions(strInstructions)

	for d := 1; d <= 9; d++ {
		for z := 0; z < 100000; z++ {
			zOut, _ := zValue(sets[13], d, z)
			if zOut == 0 {
				fmt.Printf("d: %v, inZ: %v, outZ: %v\n", d, z, zOut)
			}
		}
	}
}

func TestZFormulas6(t *testing.T) {
	for z := 0; z < 100000; z++ {
		zFormula := zValueFormulas(12, 3, z)
		if zFormula == 11 {
			fmt.Printf("d: %v, inZ: %v, outZ: %v\n", 12, z, zFormula)
		}
	}
}

func TestHighest(t *testing.T) {
	zs := zsNeeded()
	highest := highestMONAD(zs)
	fmt.Printf("highest MONAD: %s\n", highest)
	fmt.Println()
}

func TestLowest(t *testing.T) {
	zs := zsNeeded()
	lowest := lowestMONAD(zs)
	fmt.Printf("lowest MONAD: %s\n", lowest)
	fmt.Println()
}
