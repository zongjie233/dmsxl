package tree

import "strconv"

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。
*/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}

	// 检查当前节点是否为叶子节点，即没有左子节点和右子节点。如果是叶子节点，表示我们已经到达一条完整的路径
	if root.Left == nil && root.Right == nil {
		// 将当前节点的值转换为字符串，并将它作为独立的路径返回
		res = append(res, strconv.Itoa(root.Val))
	}

	// 如果不是，则向下递归遍历
	templeft := binaryTreePaths(root.Left)
	for i := 0; i < len(templeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+templeft[i])
	}

	tempright := binaryTreePaths(root.Right)
	for i := 0; i < len(tempright); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tempright[i])
	}
	return res
}
