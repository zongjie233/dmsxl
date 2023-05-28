package main

type MyDqueue struct {
	queue []int
}

func NewMyDqueue() *MyDqueue {
	return &MyDqueue{
		queue: make([]int, 0),
	}
}

func (m *MyDqueue) Front() int {
	return m.queue[0]
}

func (m *MyDqueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyDqueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyDqueue) Pop(val int) { // 如果要pop的val等于第一个，则弹出
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}

}

func (m *MyDqueue) Push(val int) { //将元素 val 插入到合适的位置,维持优先队列的性质。
	for !m.Empty() && val > m.Back() { // 这里是for，不是if
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)

}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyDqueue()
	length := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front()) // 先记录目前的最大值

	for i := k; i < length; i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}
