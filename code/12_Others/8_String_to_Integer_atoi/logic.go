package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/string-to-integer-atoi/

var (
	maxInt         int
	minInt         int
	maxIntOneDigit int // maxIntの1桁目
)

func init() {
	maxInt = int(math.Pow(2.0, 31.0)) - 1
	minInt = -maxInt - 1
	maxIntOneDigit = maxInt % 10
}

func myAtoi(s string) int {
	// 文字列操作系のAPIを利用しないようにする
	rns := []rune(s)

	// 先頭空白をトリム
	sIndex := 0
	for i, rn := range rns {
		if rn != ' ' {
			sIndex = i
			break
		}
	}
	rns = rns[sIndex:]

	// 空文字の場合
	if len(rns) == 0 {
		return 0
	}

	// 符号を検索
	sign := 1
	if rns[0] == '-' {
		sign = -1
		rns = rns[1:]
	} else if rns[0] == '+' {
		rns = rns[1:]
	}

	// 数値を計算
	sum := 0
	for _, rn := range rns {
		if rn < '0' || '9' < rn {
			// 数値以外の場合、終了
			break
		}
		num := int(rn - '0')

		// オーバーフロー対策
		if sum > maxInt/10 || (sum == maxInt/10 && num > maxIntOneDigit) {
			if sign == 1 {
				return maxInt
			} else {
				return minInt
			}
		}
		// 先頭から計算していくので×10していく
		sum = sum*10 + num
	}
	return sum * sign
}
