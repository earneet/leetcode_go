package Problem_101_200

func findPeakElement(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}
	for i := 0; i < size-1; i++ {
		if nums[i+1] < nums[i] {
			return i
		}
	}
	return size - 1
}
