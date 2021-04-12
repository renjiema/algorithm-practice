package union_find

// https://leetcode-cn.com/problems/surrounded-regions/
// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

// 方向数组 d 是上下左右搜索的常用手法
var d [][]int = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}

// 方法一：DFT
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	if m == 0 || n == 0 {
		return
	}
	queue := [][]int{}
	// 将第一列和最后一列中的'O'找出
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		if board[i][n-1] == 'O' {
			queue = append(queue, []int{i, n - 1})
		}
	}
	// 将第一行和最后一行中的'O'找出
	for i := 1; i < n-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
		}
		if board[m-1][i] == 'O' {
			queue = append(queue, []int{m - 1, i})
		}
	}
	// DFT修改所有未被围绕的'O'
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			nx, ny := x+d[i][0], y+d[i][1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || board[nx][ny] != 'O' {
				continue
			}
			queue = append(queue, []int{nx, ny})
		}
	}
	// 修改所有的'O'为'X','A'为'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// 二维数组转一维：n*x+y
// 方法二：UF
func solveUF(board [][]byte) {
	m, n := len(board), len(board[0])
	if m == 0 || n == 0 {
		return
	}
	// m*n设置为虚拟节点，不修改的'O'都和虚拟节点联通
	dummy := m * n
	uf := NewUF(dummy + 1)
	// 第一列和最后一列
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	// 行
	for i := 1; i < n-1; i++ {
		if board[0][i] == 'O' {
			uf.Union(i, dummy)
		}
		if board[m-1][i] == 'O' {
			uf.Union((m-1)*n+i, dummy)
		}
	}
	// 联通和边界上的 'O'相连的'O'
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x, y := i+d[k][0], j+d[k][1]
					if board[x][y] == 'O' {
						uf.Union(i*n+j, x*n+y)
					}
				}
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !uf.Connected(i*n+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
}
