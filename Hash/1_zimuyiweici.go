package main

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-rune('a')]++ // 单引号，rune把字母转化成ascall码对应的整数

	}

	for _, r := range t {
		record[r-rune('a')]--
	}
	return record == [26]int{}
}
