package leetcode

func pivotIndex(nums []int) int {
	size := len(nums)
	for i := 1; i < size; i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	total := nums[size-1]
	if total == nums[0] {
		return 0
	}
	for i := 1; i < size; i++ {
		if remain := total - nums[i]; remain == nums[i-1] {
			return i
		}
	}
	return -1
}
