package tree

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// if both roots are null, return null
	if root1 == nil && root2 == nil {
		return nil
	}

	// if root1 is null, return root2
	if root1 == nil {
		return root2
	}

	// if root2 is null, return root1
	if root2 == nil {
		return root1
	}

	// for overlapping nodes, add their values
	root1.Val += root2.Val

	// recursively merge the left subtrees
	root1.Left = mergeTrees(root1.Left, root2.Left)

	// recursively merge the right subtrees
	root1.Right = mergeTrees(root1.Right, root2.Right)

	// return the new root1 after merging
	return root1
}
