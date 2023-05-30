package main

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums2 := []int{1, 1, 1, 2, 2, 3}
	k := 2
	maxSlidingWindow(nums, 3)
	topKFrequent(nums2, k)
}
