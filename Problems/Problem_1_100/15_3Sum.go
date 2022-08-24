package Problem_1_100

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j += 1
				}
				for j < k && nums[k] == nums[k-1] {
					k -= 1
				}
				j += 1
				k -= 1
			} else if sum < 0 {
				j += 1
			} else {
				k -= 1
			}
		}
	}
	return ans
}
