package main

/*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。
返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
*/
func reverseWords(s string) string {
	slow, fast := 0, 0
	res := []byte(s)
	// 删除头部多的空格
	for len(res) > 0 && fast < len(res) && res[fast] == ' ' {
		fast++
	}

	// 删除单词之间的多余空格
	for ; fast < len(res); fast++ {
		if fast-1 > 0 && res[fast-1] == res[fast] && res[fast] == ' ' {
			continue
		}
		res[slow] = res[fast]
		slow++
	}

	// 删除尾部多余空格,此时末尾最多有两个空格，故判断返回 :slow-1或者 :slow
	if slow-1 > 0 && res[slow-1] == ' ' {
		res = res[:slow-1]
	} else {
		res = res[:slow]
	}

	// 反转整个字符串
	reverse(&res, 0, len(res)-1)

	// 反转单个单词
	i := 0
	for i < len(res) {
		j := i
		for ; j < len(res) && res[j] != ' '; j++ {

		}
		reverse(&res, i, j-1)
		i = j
		i++
	}
	return string(res)
}

// reverse 双指针交换元素
func reverse(res *[]byte, left, right int) {
	for left < right {
		(*res)[left], (*res)[right] = (*res)[right], (*res)[left]
		left++
		right--
	}

}
