package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func Part2() {

}

func sort(str string) string {
	run := []rune(str)
	slices.Sort(run)
	return string(run)
}

func WhichDigitRemmapped(code string) (digit int, err error) {
	switch sort(code) {
	case sort("acedgfb"):
		return 8, nil
	case sort("cdfbe"):
		return 5, nil
	case sort("gcdfa"):
		return 2, nil
	case sort("fbcad"):
		return 3, nil
	case sort("dab"):
		return 7, nil
	case sort("cefabd"):
		return 9, nil
	case sort("cdfgeb"):
		return 6, nil
	case sort("eafb"):
		return 4, nil
	case sort("cagedb"):
		return 0, nil
	case sort("ab"):
		return 1, nil
	default:
		return 0, fmt.Errorf("unkown code: %s", code)
	}
}

func (entry *Entry) OutputValue() (err error) {
	for _, code := range entry.Outputs {
		digit, err := WhichDigit(code)
		if err != nil {
			digit, err = WhichDigitRemmapped(code)
			if err != nil {
				return err
			}
		}
		entry.Value += fmt.Sprint(digit)
		fmt.Println(entry)
	}
	return
}

func (entries *Entries) OutputValues() (err error) {
	for i := range *entries {
		err = (*entries)[i].OutputValue()
		if err != nil {
			return
		}
	}
	return
}
