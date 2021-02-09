package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	// m = 3, n = 7の場合、下に移動する回数は3-1=2回、右に移動する回数は7-1=6回なので
	// 下下右右右右右右の組合せ数を返却する。
	// この場合は8個の中から2個を選ぶ組合せを返却すれば良い。
	num1 := m - 1 + n - 1
	num2 := m - 1

	// num1の中からnum2選ぶ組合せを返却すれば良い。
	sum := 1.0
	for i := 0; i < num2; i++ {
		mulNum := num1 - i
		divNum := num2 - i
		sum *= float64(mulNum)
		sum /= float64(divNum)
	}
	return int(math.Round(sum))
}
