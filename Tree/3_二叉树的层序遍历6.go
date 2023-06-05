package tree

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
*/
type Node struct {
	Val      int
	Children []*Node
}

func levelOrde2(root *Node) [][]int {
	var res [][]int  // 定义存储结果的二维切片
	var temp []int   // 定义存储每一层节点值的一维切片
	if root == nil { // 如果根节点为空，则直接返回空的二维切片
		return res
	}
	queue := []*Node{root, nil} // 使用队列来存储待遍历的节点，初始时只有根节点。在根节点后添加空节点 nil，用于标识一层的结束
	for len(queue) > 1 {        // 只要队列中有超过一个元素就执行循环
		node := queue[0]  // 获取队首节点
		queue = queue[1:] // 弹出队首节点
		if node == nil {  // 如果队首节点是空节点，则说明该层遍历已完成，需要存储当前层的结果，并添加一个新的空节点来标识下一层的结束
			queue = append(queue, nil)
			res = append(res, temp)
			temp = []int{} // 清空 temp，为下一层的累加值做准备
		} else { // 如果队首节点不是空节点，则将当前节点的值存入 temp 中，并将当前节点的孩子节点（如果有的话）加入到队列中
			temp = append(temp, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
	}
	res = append(res, temp) // 将最后一层的节点值存入 res 中
	return res
}
