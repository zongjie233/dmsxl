package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return abs(rightHeight-leftHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

/*
首先，在isBalanced函数中，我们首先判断根节点是否为空。如果根节点为空，则它本身是一个平衡的二叉树，所以我们返回true。

接下来，我们用depth函数计算二叉树的深度。这里的depth函数是一个递归函数，它的作用是计算以当前节点为根的子树的深度。如果当前节点为空，则子树深度为0。否则，我们分别计算左子树和右子树的深度，并取其中较大的一个深度加1，即当前节点所在子树的深度。

接下来，我们通过比较左子树的深度和右子树的深度的差的绝对值，看是否小于等于1。如果小于等于1，则说明当前节点所在子树是平衡的。

最后，我们递归调用isBalanced函数，对当前节点的左子节点和右子节点进行判断，看它们是否都是平衡的。

综上所述，这段代码使用递归的方式，通过判断每个节点所在子树的深度，并比较左右子树的深度差，来判断给定的二叉树是否是平衡的。

*/
