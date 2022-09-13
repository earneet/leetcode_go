package Problem_101_200

import . "LeetCode_Go/DataStructure"

type NodeState struct {
	Node   *TreeNode
	Status int
}

type BSTIterator struct {
	tree  *TreeNode
	stack []*NodeState
}

func Constructor(root *TreeNode) BSTIterator {
	_s := make([]*NodeState, 0)
	_s = append(_s, &NodeState{root, 0})
	return BSTIterator{root, _s}
}

func (this *BSTIterator) Next() int {
	topState := this.stack[len(this.stack)-1]
	if topState.Status == 0 {
		for topState.Node.Left != nil {
			topState.Status = 1
			this.stack = append(this.stack, &NodeState{topState.Node.Left, 0})
			topState = this.stack[len(this.stack)-1]
		}
		result := topState.Node.Val
		this.stack = this.stack[:len(this.stack)-1]
		if topState.Node.Right != nil {
			this.stack = append(this.stack, &NodeState{topState.Node.Right, 0})
		}
		return result
	} else {
		result := topState.Node.Val
		this.stack = this.stack[:len(this.stack)-1]
		if topState.Node.Right != nil {
			this.stack = append(this.stack, &NodeState{topState.Node.Right, 0})
		}
		return result
	}
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}
