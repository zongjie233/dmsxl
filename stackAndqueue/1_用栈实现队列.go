package main

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

// func Constructor() MyQueue {
// 	return MyQueue{
// 		stackIn:  make([]int, 0),
// 		stackOut: make([]int, 0),
// 	}
// }

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

// 如果输出栈的数据为空，则需要将输入栈的数据全部导出，如果非空，则直接使用
func (this *MyQueue) Pop() int {
	inLen, outLen := len(this.stackIn), len(this.stackOut)
	if outLen == 0 {
		if inLen == 0 {
			return -1
		}
		for i := inLen - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}      //导出后清空
		outLen = len(this.stackOut) // 更新长度

	}
	val := this.stackOut[outLen-1]
	this.stackOut = this.stackOut[:outLen-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, val)
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
