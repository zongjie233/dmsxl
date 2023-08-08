package main

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil { // 一定先判断cur.Next != nil，不然可能会报空指针异常
		// 以前三个节点为例，做以下注释
		temp := cur.Next            // 保存1号节点
		temp1 := cur.Next.Next.Next // 保存3号节点

		cur.Next = cur.Next.Next
		cur.Next.Next = temp
		temp.Next = temp1

		cur = cur.Next.Next
	}
	return dummyHead.Next
}
