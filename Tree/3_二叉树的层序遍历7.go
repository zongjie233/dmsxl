package tree

import "math"

/*
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		l := len(queue)
		max := math.MinInt32 // 定义一个最小整数常量 max，用于比较 TreeNode 中的节点值, 需要再for循环里边定义
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
