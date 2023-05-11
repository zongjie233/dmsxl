package main

/*
leetcode:977
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]
*/

// 双指针
func demo3(nums []int) []int {
	n := len(nums)
	k := n - 1
	res := make([]int, n)
	left, right := 0, n-1
	for left <= right { //这里是<=
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[k] = nums[left] * nums[left]
			left++
		} else {
			res[k] = nums[right] * nums[right]
			right--
		}
		k--
	}
	return res
}

/*
非递减顺序的数组因为包含负数，则平方之后两边大，中间小，最大值只能出现在两边，故使用双指针
*/
