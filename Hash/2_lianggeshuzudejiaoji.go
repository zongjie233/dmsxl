package main

/*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
*/

// 使用map模拟set数据结构。1. 无序:Set 中的元素没有任何顺序。 2. 不重复:Set 中的元素都是唯一的。
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) //map模拟set，
	res := make([]int, 0)

	//遍历第一个数组，如果set中没有该数据则存入set，
	//2. 尽管每个键有一个默认值 struct{},但这只是技术实现,该键实际上并不存在。
	//3. 所以访问任何键,ok 都会是 false,表示该键不存在。
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{} // struct{}{}构建一个空的结构体
		}
	}

	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v) // 从set map 中删掉v对应的值
		}
	}
	return res
}
