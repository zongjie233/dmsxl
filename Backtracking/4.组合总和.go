package backtracking

import "sort"

var (
	res4  [][]int
	path4 []int
)

func combinationSum(candidates []int, target int) [][]int {
	res4, path4 = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序,方便剪枝
	dfs4(candidates, 0, target)
	return res4
}

func dfs4(candidates []int, start int, target int) {
	if target == 0 {
		tmp := make([]int, len(path4))
		copy(tmp, path4)
		res4 = append(res4, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		path4 = append(path4, candidates[i])
		dfs4(candidates, i, target-candidates[i]) // 这里不是i+1, 因为可以重复选取
		path4 = path4[:len(path4)-1]
	}
}
