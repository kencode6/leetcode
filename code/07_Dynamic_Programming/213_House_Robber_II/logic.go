package logic

import (
	"math"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/house-robber-ii/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	// 循環しているので最後の要素を除いたものと、最初の要素を除いたもので計算し大きい方を取る
	// 最後の要素を除外
	calclator1 := NewCalclator(nums[0 : len(nums)-1])
	sum1 := calclator1.Sum()

	// 最初の要素を除外
	calclator2 := NewCalclator(nums[1:])
	sum2 := calclator2.Sum()

	sum := int(math.Max(float64(sum1), float64(sum2)))
	return sum
}

type Calclator struct {
	nums  []int
	cache map[int]int
}

func NewCalclator(nums []int) *Calclator {
	return &Calclator{
		nums:  nums,
		cache: make(map[int]int),
	}
}

func (c *Calclator) Sum() int {
	return c.sum(len(c.nums)-1, 1)
}

func (c *Calclator) sum(index int, depth int) int {
	if sum, ok := c.cache[index]; ok {
		// キャッシュを利用
		return sum
	}

	// S(n) = max(S(n-2) + a(n), S(n-1))
	// S(n) = 0 (n<0)
	// の漸化式から算出

	// S(n) = 0 (n<0)
	if index < 0 {
		return 0
	}

	depth++

	// S(n) = max(S(n-2) + a(n), S(n-1))
	sumPrev2 := c.sum(index-2, depth)
	sumPrev1 := c.sum(index-1, depth)
	sum := int(math.Max(float64(sumPrev2+c.nums[index]), float64(sumPrev1)))

	// キャッシュを登録
	c.cache[index] = sum
	return sum
}
