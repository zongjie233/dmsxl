package backtracking

import "sort"

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

与4的区别在于数组中有重复元素,但不能有重复的组合

注意：解集不能包含重复的组合。
*/

var (
	res  [][]int
	path []int
	used []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	used = make([]bool, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		dfs(candidates, i+1, target-candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
}
