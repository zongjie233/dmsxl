package tree

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
*/

func levelOrder(root *TreeNode) [][]int {

	// 避免根节点为空
	if root == nil {
		return [][]int{}
	}

	// 存入根节点
	queue := []*TreeNode{root}
	res := make([][]int, 0)

	// 队列不为0则一直进行
	for len(queue) > 0 {

		// 记录当前队列长度，其值在一直变化
		l := len(queue)
		tmp := make([]int, 0, l)
		//遍历该层，左右节点不为空则添加到队列中
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 存放队列中的值，便于下面做清空队列的操作
			tmp = append(tmp, queue[i].Val)
		}
		// 清空队列
		queue = queue[l:]
		//存放数据
		res = append(res, tmp)
	}
	return res
}
