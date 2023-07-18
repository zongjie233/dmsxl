package tree

/*
给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。
你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	root.Right = trimBST(root.Right, low, high)
	root.Left = trimBST(root.Left, low, high)
	return root
}
