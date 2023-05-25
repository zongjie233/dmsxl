package main

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

*/

// 双指针法，不需要额外的辅助空间，从后向前填充
func replaceSpace(s string) string {
	res := make([]byte, 0)
	arrs := []byte(s) // 字符串转切片
	for i := 0; i < len(arrs); i++ {
		if arrs[i] == ' ' {
			res = append(res, []byte("%20")...) // ... 将%20拆开，如果不添加...，则传给append是一个切片，加了之后则是%20拆开的元素
		} else {
			res = append(res, arrs[i])
		}
	}
	return string(res)
}
