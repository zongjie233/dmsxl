package tree

/*
给定一个二叉树，检查它是否是镜像对称的。 后序遍历，左右中
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && (right.Left == left.Right) && (left.Right == right.Left)
}
