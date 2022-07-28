package leetcode

func runningSum(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)
	copy(ans, nums)
	for i := 1; i < size; i++ {
		ans[i] = ans[i-1] + ans[i]
	}
	return ans
}
