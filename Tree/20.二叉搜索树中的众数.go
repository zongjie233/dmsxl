package tree

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
*/
func findMode(root *TreeNode) []int {
	res := []int{}

	var prev *TreeNode
	max, count := 0, 1

	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)

		if prev != nil {
			if prev.Val == node.Val {
				count++
			} else {
				count = 1
			}
		}

		if count > max {
			max = count
			res = []int{node.Val}
		} else if count == max {
			res = append(res, node.Val)
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
