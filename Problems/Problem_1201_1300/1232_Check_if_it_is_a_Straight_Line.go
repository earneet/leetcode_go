package Problem_1201_1300

import "sort"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	sort.Slice(coordinates, func(ai, bi int) bool {
		a := coordinates[ai]
		b := coordinates[bi]
		if a[0] == b[0] {
			return a[1] < b[1]
		} else {
			return a[0] < b[0]
		}
	})
	slope := float32(coordinates[1][1]-coordinates[0][1]) / float32(coordinates[1][0]-coordinates[0][0])
	for i := 2; i < len(coordinates); i += 1 {
		if float32(coordinates[i][1]-coordinates[i-1][1])/float32(coordinates[i][0]-coordinates[i-1][0]) != slope {
			return false
		}
	}
	return true
}
