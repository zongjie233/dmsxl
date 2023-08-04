package main

/*
你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。

实现 MyLinkedList 类：

MyLinkedList() 初始化 MyLinkedList 对象。
int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
*/

import "fmt"

// 单链表写法 //

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type SingleLinkedList struct {
	dummyHead *SingleNode
	Size      int
}

// 初始化链表
func (list *SingleLinkedList) Init() *SingleLinkedList {
	list.Size = 0
	list.dummyHead = new(SingleNode)
	return list
}

// 获取第index个节点数值
func (list *SingleLinkedList) get(index int) int {
	if list != nil || index < 0 || index > list.Size {
		return -1
	}
	// 让cur等于真正头节点
	cur := list.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// 在链表最前面插入一个节点
func (list *SingleLinkedList) addAtHead(val int) {
	// 以下两行代码可用一行代替
	// newNode := new(SingleNode)
	// newNode.Val = val
	newNode := &SingleNode{Val: val}

	newNode.Next = list.dummyHead.Next
	list.dummyHead.Next = newNode
	list.Size++
}

// 在链表最后面插入一个节点
func (list *SingleLinkedList) addAtTail(val int) {
	newNode := &SingleNode{Val: val}
	cur := list.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	list.Size++
}

// 打印链表
func (list *SingleLinkedList) printLinkedList() {
	cur := list.dummyHead
	for cur.Next != nil {
		fmt.Println(cur.Next.Val)
		cur = cur.Next
	}
}

// 在第index个节点之前插入新节点
func (list *SingleLinkedList) addAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > list.Size {
		return
	}

	newNode := &SingleNode{Val: val}
	cur := list.dummyHead //用虚拟头节点不用考虑在头部插入的情况
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	list.Size++
}

func (list *SingleLinkedList) deleteAtIndex(index int) {
	if list.dummyHead == nil {
		return
	}
	if index < 0 || index > list.Size {
		return
	}
	cur := list.dummyHead
	if index == 0 {
		list.dummyHead = list.dummyHead.Next
		list.Size--
		return
	}
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	list.Size--
}

// 采用模拟的方式进行code

type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 同时完成新建节点和赋值给头结点的工作
	this.head = &Node{val, this.head}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	// 如果此时链表长度为0,则只需要使用addathead方法代替
	if this.size == 0 {
		this.AddAtHead(val)
		return
	}
	curr := this.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node{val: val}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
	} else {
		curr := this.head
		// index-1,到目标位置的前一个节点
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		curr.next = &Node{val, curr.next}
		this.size++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	this.size--
	// 如果index =0 则直接通过操作头结点完成
	if index == 0 {
		this.head = this.head.next
		return
	}
	curr := this.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = curr.next.next

}
