package main

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
*/
func reverseStr2(s string, k int) string {
	res := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k { // for循环中，i一次前进2k个
		if k+i <= length {
			reverse1(res[i : i+k])
		} else {
			reverse1(res[i:length])
		}
	}
	return string(res)
}

func reverse1(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right-- //右指针--操作
	}
}
