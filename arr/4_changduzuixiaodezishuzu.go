package arr

/*
leet:209
给定一个含有 n 个正整数的数组和一个正整数 target 。找出该数组中满足其和 ≥ target 的长度最小的连续子数组
 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

*/

//滑动窗口
func huadongchuangkou(nums []int, target int) int {
	i := 0
	length := len(nums)
	res := length + 1
	sum := 0 // 窗口内数据的和
	for j := 0; j < length; j++ {
		sum += nums[j]
		for sum >= target {
			sublength := j - i + 1
			if sublength < res {
				res = sublength
			}
			sum -= nums[i]
			i++
		}
	}
	if res == length+1 {
		return 0
	} else {
		return res

	}

}
