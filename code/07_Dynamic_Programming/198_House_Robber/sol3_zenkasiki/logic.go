package logic

import (
	"fmt"
	"math"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	cache := make(map[int]int)
	sum := sum(nums, len(nums)-1, 1, cache)
	fmt.Printf("%d", sum)
	return sum
}

func sum(nums []int, index int, depth int, cache map[int]int) int {
	if sum, ok := cache[index]; ok {
		// キャッシュを利用
		return sum
	}

	depth++

	// S(n) = max(S(n-2) + a(n), S(n-1))
	// S(n) = 0 (n<0)
	// の漸化式から算出

	// S(n) = 0 (n<0)
	if index < 0 {
		return 0
	}

	// S(n) = max(S(n-2) + a(n), S(n-1))
	sumPrev2 := sum(nums, index-2, depth, cache)
	sumPrev1 := sum(nums, index-1, depth, cache)
	sum := int(math.Max(float64(nums[index]+sumPrev2), float64(sumPrev1)))

	// 答えをキャッシュ
	cache[index] = sum
	return sum
}
