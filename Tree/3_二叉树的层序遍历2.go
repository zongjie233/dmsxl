package tree

/*
leet:107
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
*/
func levelOrderBottom(root *TreeNode) [][]int {
	tmp := levelOrder(root)
	res := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}
