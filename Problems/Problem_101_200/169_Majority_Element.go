package Problem_101_200

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		dict[nums[i]] += 1
		if dict[nums[i]] > size/2 {
			return nums[i]
		}
	}
	return 0
}
