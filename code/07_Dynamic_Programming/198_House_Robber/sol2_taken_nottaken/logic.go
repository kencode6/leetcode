package logic

import (
	"math"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	calculator := NewCalculator(nums)
	sum := calculator.sumMoney(0, 1)
	return sum
}

type Calculator struct {
	nums  []int
	cache map[int]int
}

func NewCalculator(nums []int) *Calculator {
	return &Calculator{
		nums:  nums,
		cache: make(map[int]int),
	}
}

func (c *Calculator) sumMoney(index int, depth int) int {
	if sum, ok := c.cache[index]; ok {
		// キャッシュ利用
		return sum
	}

	if index == len(c.nums)-1 {
		//　終端
		return c.nums[index]
	}

	depth++

	// 取る
	takeSum := c.nums[index]
	if index+2 < len(c.nums) {
		takeSum += c.sumMoney(index+2, depth)
	}

	// 取らない
	notTakeSum := c.sumMoney(index+1, depth)

	// 取った場合と取らなかった場合を比較して大きい方を取得
	sum := int(math.Max(float64(takeSum), float64(notTakeSum)))

	// キャッシュする
	c.cache[index] = sum
	return sum
}
