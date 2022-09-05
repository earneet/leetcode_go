package Problem_401_500

import . "LeetCode_Go/DataStructure"

func levelOrder(root *Node) [][]int {
	queue := make([]*Node, 0)
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := make([]*Node, 0)
		levelList := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			levelList = append(levelList, node.Val)
			for _, n := range node.Children {
				newQueue = append(newQueue, n)
			}
		}
		queue = newQueue
		ans = append(ans, levelList)
	}
	return ans
}
