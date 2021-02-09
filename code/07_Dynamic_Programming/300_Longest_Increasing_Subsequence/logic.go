package logic

import "math"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	calculator := NewSubsequenceCalculator(nums)
	maxSeq := calculator.calcLength(-1, 0)
	return maxSeq
}

type SubsequenceCalculator struct {
	nums           []int
	maxSubSequence int

	// 計算速度向上のために計算したシーケンスをキャッシュしておく
	seqCache [][]int
}

func NewSubsequenceCalculator(nums []int) *SubsequenceCalculator {
	seqCache := [][]int{}
	for i := 0; i < len(nums)+1; i++ {

		tmpNums := []int{}
		for j := 0; j < len(nums); j++ {
			tmpNums = append(tmpNums, -1)
		}
		seqCache = append(seqCache, tmpNums)
	}

	return &SubsequenceCalculator{
		nums:     nums,
		seqCache: seqCache,
	}
}

func (c *SubsequenceCalculator) calcLength(prevIndex int, currentIndex int) int {
	if currentIndex == len(c.nums) {
		// 終端
		return 0
	}

	if c.seqCache[prevIndex+1][currentIndex] > 0 {
		// すでに計算済みのシーケンスの場合
		return c.seqCache[prevIndex+1][currentIndex]
	}

	takenSeq := 0
	if prevIndex == -1 || c.nums[prevIndex] < c.nums[currentIndex] {
		// 次の数値比較時
		takenSeq = c.calcLength(currentIndex, currentIndex+1)
		takenSeq++
	}

	// 次の数値スキップ時
	notTakenSeq := c.calcLength(prevIndex, currentIndex+1)

	// 取得した場合と取得しなかった場合で大きい方を選択
	maxSeq := int(math.Max(float64(takenSeq), float64(notTakenSeq)))

	// キャッシュに保持
	c.seqCache[prevIndex+1][currentIndex] = maxSeq
	return maxSeq
}
