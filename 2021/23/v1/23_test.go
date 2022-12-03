package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var test1 = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

func TestPart1(t *testing.T) {
	bur := parse(test1)
	bur.Print()

	for i := range bur.rooms {
		fmt.Println(bur.energyToFillRoom(i))
	}
}
