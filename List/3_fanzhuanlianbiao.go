package main

/*
206
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

// 双指针法
func fanzhuanlianbiao(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil { // 确定循环终止条件
		tmp := cur.Next // 需要一个临时变量来存储cur.Next，不然该值会丢失
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre

}

// 递归
func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	return reverse(cur, temp)
}

func reverseList(head *ListNode) *ListNode {
	return reverse(head, nil)
}
