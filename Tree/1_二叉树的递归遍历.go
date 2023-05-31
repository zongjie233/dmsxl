package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 中左右
func perorderTraversal(root *TreeNode) []int {
	res := []int{}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // 中
		traversal(node.Left)        // 左
		traversal(node.Right)       // 右
	}
	traversal(root)
	return res
}

// 中序
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)        // 左
		res = append(res, node.Val) // 中
		traversal(node.Right)       // 右
	}
	traversal(root)
	return res
}

// 后序
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)        // 左
		traversal(node.Right)       // 右
		res = append(res, node.Val) // 中
	}
	traversal(root)
	return res
}
