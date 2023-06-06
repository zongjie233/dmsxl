package tree

/*
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func connect2(root *Node1) *Node1 {
	if root == nil {
		return root
	}

	queue := []*Node1{root}

	for len(queue) > 0 {
		var pot []*Node1

		for i, node := range queue {
			if i+1 < len(queue) {
				node.Next = queue[i+1]
			}
			if node.Left != nil {
				pot = append(pot, node.Left)
			}
			if node.Right != nil {
				pot = append(pot, node.Right)
			}
		}
		queue = pot
	}
	return root
}
