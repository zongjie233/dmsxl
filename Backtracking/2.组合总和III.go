package backtracking

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/

/*
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
*/
var (
	path []int
	res  [][]int
)

func combinationSum3(k int, n int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(k, n, 1, 0)
	return res
}

func dfs(k, n int, start int, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}

	for i := start; i <= n-(k-len(path))+1; i++ {
		// 剪枝
		if sum+i > n || k-len(path) > 9-i+1 {
			break
		}

		path = append(path, i)
		dfs(k, n, i+1, sum+i)
		path = path[:len(path)-1]

	}
}
