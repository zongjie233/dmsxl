package tree

import "math"

/*
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
*/

func getMinimumDifference(root *TreeNode) int {

	var prev *TreeNode
	min := math.MaxInt64

	var travel func(node *TreeNode)

	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}
