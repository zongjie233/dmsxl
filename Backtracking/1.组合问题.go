package backtracking

/*
模板：
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

func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	backtrack(n, k, 1)
	return res
}

func backtrack(n int, k int, start int) {
	if len(path) == k { // 找到一组结果,进行保存
		tmp := make([]int, k)
		copy(tmp, path) // copy方法:复制切片到另一个切片
		res = append(res, tmp)
		return
	}
	for i := start; i <= n-(k-len(path))+1; i++ { // 剪枝
		path = append(path, i)
		backtrack(n, k, i+1)
		path = path[:len(path)-1] // 踢出第一个元素
	}
}
