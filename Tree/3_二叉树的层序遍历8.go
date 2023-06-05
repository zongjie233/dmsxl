package tree

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
*/
type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return root
	}

	queue := []*Node1{root}

	for len(queue) > 0 {
		var pot []*Node1

		for i, node := range queue {
			if i+1 < len(queue) { // 如果当前节点不是该层的最后一个节点，则将next指向下一个节点
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
