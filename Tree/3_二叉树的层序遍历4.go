package tree

/*
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
*/
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}

	queue := []*TreeNode{root}
	res := []float64{}
	sum := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			sum += queue[i].Val
		}
		res = append(res, float64(sum)/float64(l))
		queue = queue[l:]
		sum = 0
	}
	return res
}
