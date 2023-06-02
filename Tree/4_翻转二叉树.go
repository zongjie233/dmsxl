package tree

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
*/

// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 下面几行的顺序无所谓
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
