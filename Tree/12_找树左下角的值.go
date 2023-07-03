package tree

/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。
*/

// DFS

type Result struct {
	Value     int
	MaxHeight int
}

func findBottomLeftValue(root *TreeNode) int {
	result := findBottomLeftValueDFS(root, 0)
	return result.Value
}

func findBottomLeftValueDFS(root *TreeNode, curHeight int) Result {
	if root.Left == nil && root.Right == nil {
		return Result{Value: root.Val, MaxHeight: curHeight}
	}

	var leftResult, rightResult Result

	if root.Left != nil {
		leftResult = findBottomLeftValueDFS(root.Left, curHeight+1)
	}
	if root.Right != nil {
		rightResult = findBottomLeftValueDFS(root.Right, curHeight+1)
	}

	if leftResult.MaxHeight >= rightResult.MaxHeight {
		return leftResult
	}

	return rightResult
}

// BFS 层序遍历
func findBottomLeftValueBFS(root *TreeNode) int {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		next := []*TreeNode{}

		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if len(next) == 0 {
			return queue[0].Val
		}
		queue = next
	}
	return 0

}
