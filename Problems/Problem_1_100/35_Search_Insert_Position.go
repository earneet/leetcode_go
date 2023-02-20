package Problem_1_100

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, len(nums)/2
	for left <= right {
		mid = (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
