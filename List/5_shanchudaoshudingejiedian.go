package main

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
*/
// 双指针，快指针先移动n+1步，然后快慢指针同时移动，直到快指针指向nil，此时慢指针在要删除元素的前边一位
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyhead := &ListNode{}
	dummyhead.Next = head
	fast := dummyhead
	slow := dummyhead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyhead.Next //返回真正的头节点
}
