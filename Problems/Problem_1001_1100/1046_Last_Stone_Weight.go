package Problem_1001_1100

import "sort"

func insert(nums []int, stone int) []int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == stone {
			break
		} else if nums[mid] > stone {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	mid := start + (end-start)/2
	nums = append(nums[:mid], append([]int{stone}, nums[mid:]...)...)
	return nums
}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		x := stones[len(stones)-2]
		y := stones[len(stones)-1]
		n := y - x
		stones = stones[:len(stones)-2]
		if n > 0 {
			stones = insert(stones, n)
		}
	}

	if len(stones) == 0 {
		return 0
	} else {
		return stones[0]
	}
}
