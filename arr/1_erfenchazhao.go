package arr

/*
leetcode:704
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

*/

// 左闭右闭
func erfenchazhao(slice []int, target int) int {
	left, right := 0, len(slice)-1
	for left <= right { // 左闭右闭时为小于等于
		mid := (left + right) / 2
		if slice[mid] == target {
			return mid
		} else if slice[mid] > target {
			right = mid - 1
		} else if slice[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

/*
左闭右闭区间： left<=right

左闭右开区间：left < right
*/
