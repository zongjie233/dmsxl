package main

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。


*/

// 使用数组占用的空间小于map
func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)

	for _, word := range magazine { // 遍历magazine，将其中出现的字符进行记录
		record[word-'a']++
	}

	for _, word := range ransomNote {
		record[word-'a']--        // 这里进行了--操作，故下一行的判断为<0,而不是<=0
		if record[word-'a'] < 0 { // 如果ransomNote中的字符不是magazine里的，则对应值小于零
			return false
		}
	}
	return true
}
