package main

func cubesWithIntersections(steps []step) (cubes []cube) {
	for _, step := range steps {
		for _, existing := range cubes {
			if yes, intersection := existing.intersection(step); yes {
				cubes = append(cubes, intersection)
			}
		}
		if step.on {
			cubes = append(cubes, step)
		}
	}
	return
}

func (c1 cube) intersection(c2 cube) (intersects bool, intersection cube) {
	intersection.start[0] = max(c1.start[0], c2.start[0])
	intersection.end[0] = min(c1.end[0], c2.end[0])

	intersection.start[1] = max(c1.start[1], c2.start[1])
	intersection.end[1] = min(c1.end[1], c2.end[1])

	intersection.start[2] = max(c1.start[2], c2.start[2])
	intersection.end[2] = min(c1.end[2], c2.end[2])

	if intersection.start[0] > intersection.end[0] {
		return
	}
	if intersection.start[1] > intersection.end[1] {
		return
	}
	if intersection.start[2] > intersection.end[2] {
		return
	}

	if c1.on == c2.on {
		intersection.on = !c2.on //intersection state cancels out existing state
	} else {
		intersection.on = c2.on
	}

	return true, intersection
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countOn2(cubes []cube) (on int) {
	for _, cube := range cubes {
		volume := (cube.end[0] - cube.start[0] + 1) * (cube.end[1] - cube.start[1] + 1) * (cube.end[2] - cube.start[2] + 1)
		if cube.on {
			on += volume
		} else {
			on -= volume
		}
	}
	return
}
