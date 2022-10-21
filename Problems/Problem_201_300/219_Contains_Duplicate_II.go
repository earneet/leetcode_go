package Problem_201_300

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		dict[key] += 1
		if dict[key] > 1 {
			return true
		}
		if i >= k {
			key = nums[i-k]
			dict[key] -= 1
		}
	}
	return false
}
