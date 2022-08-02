package leetcode

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	list := make([]int, 0, len(matrix)*len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		list = append(list, matrix[i]...)
	}
	sort.Ints(list)
	return list[k-1]
}
