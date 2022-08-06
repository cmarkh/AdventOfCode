package advent1

import (
	"fmt"
)

func CountDepthIncreases(depths []int64) (increases int64, err error) {
	var prior int64
	for i, d := range depths {
		if i == 0 {
			continue
		}
		if d > prior {
			increases++
			fmt.Printf("%d (increased)\n", d)
		} else if d == prior {
			fmt.Printf("%d (no change)\n", d)
		} else {
			fmt.Printf("%d (decreased)\n", d)
		}
		prior = d
	}

	fmt.Printf("%d increases\n", increases)

	return
}

func SlidingWindow(depths []int64) (increases int64, err error) {
	var sums []int64
	for i := 0; i+2 < len(depths); i++ {
		sums = append(sums, depths[i]+depths[i+1]+depths[i+2])
	}
	return CountDepthIncreases(sums)
}
