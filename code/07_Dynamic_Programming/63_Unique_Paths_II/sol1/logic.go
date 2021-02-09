package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 初回がブロックの場合、終了
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// パス数格納配列を複製する。要素は-1で初期化する。
	pathNumGrid := [][]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		row := []int{}
		for j := 0; j < len(obstacleGrid[i]); j++ {
			row = append(row, -1)
		}
		pathNumGrid = append(pathNumGrid, row)
	}

	maxRow := len(obstacleGrid) - 1
	maxCol := len(obstacleGrid[0]) - 1

	// 左と上の数値を加算するというロジックにする
	pathNumGrid[0][0] = 1
	if maxRow > 0 {
		// 初回の行
		for i := 1; i <= maxRow; i++ {
			if obstacleGrid[i][0] == 1 {
				// 岩を見つけたら対象ブロックを0に
				pathNumGrid[i][0] = 0
			} else {
				// 岩以外は前の数値と同じ
				pathNumGrid[i][0] = pathNumGrid[i-1][0]
			}
		}
	}

	if maxCol > 0 {
		// 初回の列
		for i := 1; i <= maxCol; i++ {
			if obstacleGrid[0][i] == 1 {
				// 岩を見つけたら対象ブロックを0に
				pathNumGrid[0][i] = 0
			} else {
				// 岩以外は前の数値と同じ
				pathNumGrid[0][i] = pathNumGrid[0][i-1]
			}
		}
	}

	// gridが1行または1列の場合
	if maxRow == 0 || maxCol == 0 {
		return pathNumGrid[maxRow][maxCol]
	}

	// 2回目以降の行
	for i := 1; i <= maxRow; i++ {
		for j := 1; j <= maxCol; j++ {
			if obstacleGrid[i][j] == 1 {
				// 岩を見つけたら対象ブロックを0に
				pathNumGrid[i][j] = 0
			} else {
				// 岩以外は左と上の数値を加算
				pathNumGrid[i][j] = pathNumGrid[i-1][j] + pathNumGrid[i][j-1]
			}
		}
	}
	return pathNumGrid[maxRow][maxCol]
}
