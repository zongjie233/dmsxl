package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val { // if key value is less than the root value, move to left subtree
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { // if key value is greater than the root value, move to right subtree
		root.Right = deleteNode(root.Right, key)
	} else { // if key is equal to root value
		if root.Left == nil { // node with only one child or no child
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// node with two children, get the inorder successor (smallest in the right subtree)
		minNode := findMin(root.Right)
		root.Val = minNode.Val                        // copy the inorder successor's value to this node
		root.Right = deleteNode(root.Right, root.Val) // delete the inorder successor
	}
	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
