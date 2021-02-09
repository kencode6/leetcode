package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	prevMaxSum := 0
	currMaxSum := 0
	for i := 0; i < len(nums); i++ {
		// 前回のmaxが要素を取った場合と今回のmaxを比較
		prevMaxSum += nums[i]
		maxSum := int(math.Max(float64(prevMaxSum), float64(currMaxSum)))

		// 次の前回のmaxを前回のcurrent max、次の今回のmaxを最大maxにする
		prevMaxSum = currMaxSum
		currMaxSum = maxSum
	}
	return currMaxSum
}
