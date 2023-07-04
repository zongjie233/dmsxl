package tree

/*
给你二叉树的根节点 `root` 和一个整数目标和 `targetSum` ，找出所有 **从根节点到叶子节点** 路径总和等于给定目标和的路径。

**叶子节点** 是指没有子节点的节点。
*/

// pathSum 函数用于找出所有从根节点到叶子节点的路径，使得路径上节点值的和等于目标和
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{} // 结果列表，用于存储所有满足条件的路径
	path := []int{}     // 当前路径
	dfs(root, targetSum, path, &result)
	return result
}

// dfs 函数用于深度优先搜索二叉树的所有路径，并记录路径和
func dfs(node *TreeNode, targetSum int, path []int, result *[][]int) {
	if node == nil {
		return
	}

	// 将当前节点的值加入路径和
	path = append(path, node.Val)

	// 当前节点是叶子节点且路径和等于目标和
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		// 将当前路径加入结果列表
		*result = append(*result, append([]int{}, path...))
	}

	// 递归遍历左子树和右子树
	dfs(node.Left, targetSum-node.Val, path, result)
	dfs(node.Right, targetSum-node.Val, path, result)

	// 回溯，将当前节点从路径和中移除
	path = path[:len(path)-1]
}
