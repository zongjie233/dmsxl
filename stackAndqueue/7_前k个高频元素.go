package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	// 记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}

	h := &IHeap{}
	heap.Init(h)

	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)

	for i := 0; i < k; i++ { // 从小顶堆中弹出k个元素，并从大到小排列放入res中
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 构建小顶堆

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h) // 返回子数组的个数
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
