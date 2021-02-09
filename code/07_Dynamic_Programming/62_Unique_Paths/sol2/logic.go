package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	// m行n列の配列を用意して左と上の合計値を検討する
	// m = 3, n = 7の場合
	// [
	// 	[1, 1, 1, 1, 1, 1, 1]
	// 	[1, 2, 3, 4, 5, 6, 7]
	// 	[1, 3, 6,10,15,21,28]
	// ]
	// 右下要素の数が答えとなる

	// m行n列の要素が全て1の2次元配列を用意
	grid := [][]int{}
	for i := 0; i < m; i++ {
		nums := []int{}
		for j := 0; j < n; j++ {
			nums = append(nums, 1)
		}
		grid = append(grid, nums)
	}

	// 左と上の数を加算する
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	sum := grid[m-1][n-1]
	return sum
}
