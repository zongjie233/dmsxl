package tree

//使用栈模拟递归过程

// 前序
func perorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil { // 先将右子树压入栈中，便于左子树的数据弹出
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

// 后序 对比前序，先把左子树压入栈中，这样出栈之后的顺序为中右左，此时反转res顺序为左右中，即为后序遍历
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			res = append(res, node.Val)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(res []int) {
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

}

// 中序遍历

func inorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		// 循环访问左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 访问根节点并弹出栈顶元素
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)

		// 访问右子树
		root = root.Right
	}

	return result
}
