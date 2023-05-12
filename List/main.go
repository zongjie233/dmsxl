package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list = new(SingleLinkedList)
	list.Init()
	list.addAtHead(1)
	list.addAtTail(3)
	list.addAtIndex(1, 2)
	list.get(1)
	list.deleteAtIndex(1)
	list.get(1)
	fmt.Println("---------")
	list.printLinkedList()
}
