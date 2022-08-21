package Problem_701_800

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	for i := 2; i < size; i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[size-1], cost[size-2])
}
