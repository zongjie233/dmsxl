package arr

/*
https://leetcode.cn/problems/spiral-matrix-ii/
给你一个正整数 n ，生成一个包含 1 到 n^2所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

// 重点：边界条件，指定同一条循环条件

// (startx,starty) (i,j)
func luoxuan(n int) [][]int {
	startx, starty := 0, 0
	var center, loop int = n / 2, n / 2
	count := 1
	offset := 1
	// 创建二维数组
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for loop > 0 {
		i, j := startx, starty

		// 行数不边列数变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}

		// 列不变，行变
		for i = starty; i < n-offset; i++ {
			res[i][j] = count
			count++
		}

		//行不变，列变
		for ; j > starty; j-- { // 这里的边界条件是>starty
			res[i][j] = count
			count++
		}

		//列不变 行变
		for ; i > startx; i-- { // 这里的边界条件是>startx
			res[i][j] = count
			count++
		}
		// 一圈结束
		startx++
		starty++
		offset++
		loop-- // loop--
	}
	if n%2 == 1 {
		res[center][center] = n * n

	}
	return res
}
