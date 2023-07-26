package backtracking

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/
var (
	m     []string
	path1 []byte
	res1  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path1, res1 = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res1
	}

	dfs1(digits, 0)
	return res1
}

func dfs1(digits string, idx int) {
	if len(path) == len(digits) {
		tmp := string(path1)
		res1 = append(res1, tmp)
		return
	}

	digit := int(digits[idx] - '0')
	str := m[digit-2]
	for j := 0; j < len(str); j++ {
		path1 = append(path1, str[j])
		dfs1(digits, idx+1)
		path1 = path1[:len(path1)-1]
	}
}
