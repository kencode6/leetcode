package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	calculator := NewSumCalculator(nums)
	sum := calculator.sum(len(nums)-1, 1)
	return sum
}

type SumCalculator struct {
	nums  []int
	cache map[int]int
}

func NewSumCalculator(nums []int) *SumCalculator {
	return &SumCalculator{
		nums:  nums,
		cache: make(map[int]int),
	}
}

func (c *SumCalculator) sum(index int, depth int) int {
	if sum, ok := c.cache[index]; ok {
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
	sumPrev2 := c.sum(index-2, depth)
	sumPrev1 := c.sum(index-1, depth)
	sum := int(math.Max(float64(c.nums[index]+sumPrev2), float64(sumPrev1)))

	// 答えをキャッシュ
	c.cache[index] = sum
	return sum
}
