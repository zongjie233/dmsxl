package main

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

*/

// 不适用map，因为后期为了去重操作会超时。采用双指针，更快找到三元组
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 对a去重,这里是nums[i] == nums[i-1]做判断，以免因为三元组中的重复数组而漏掉三元组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})

				//找到一个三元组之后对b，c去重
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}

		}
	}
	return res
}
