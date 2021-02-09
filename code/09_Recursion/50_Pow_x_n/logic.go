package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/powx-n/

const truncateNum = 100000

func myPow(x float64, n int) float64 {
	// 単純にxをn回かけるのが楽だが、それだとnが大きい場合に計算時間がかかりすぎるので
	// 2^11 = 4^5*2 = 16^2*4*2
	// のような形で同一の次数をまとめて再帰的に計算する必要がある。
	if n == 0 {
		return 1
	}

	// マイナスの場合、割り算に
	if n < 0 {
		n *= -1
		x = 1.0 / x
	}

	// 再帰的に計算
	sumX := pow(x, n)

	// 小さすぎる桁数をtruncate
	return math.Floor(sumX*truncateNum) / truncateNum
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	// x^nの場合、(x*x)^(n/2)にして計算
	sumX := 1.0
	x2 := pow(x*x, n/2)
	sumX *= x2

	// nが奇数の場合xを掛ける
	if n%2 == 1 {
		sumX *= x
	}
	return sumX
}
