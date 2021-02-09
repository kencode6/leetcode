package logic

import "sort"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	searcher := NewCombinationSearcher(candidates, target)
	return searcher.Search()
}

type CombinationSearcher struct {
	candidates   []int
	target       int
	combinations [][]int
}

func NewCombinationSearcher(candidates []int, target int) *CombinationSearcher {
	// 初回にcandidatesにソートをかけると重複した要素を評価せずに効率的に探索できる。
	sort.Sort(sort.IntSlice(candidates))

	return &CombinationSearcher{
		candidates:   candidates,
		target:       target,
		combinations: [][]int{},
	}
}

func (c *CombinationSearcher) Combinations() [][]int {
	return c.combinations
}

func (c *CombinationSearcher) Search() [][]int {
	currentComb := newCombinationHolder([]int{}, 0)
	c.search(currentComb, 0, 1)
	return c.combinations
}

func (c *CombinationSearcher) search(currentComb *combinationHolder, currentIndex int, depth int) {
	depth++
	for i := currentIndex; i < len(c.candidates); i++ {
		num := c.candidates[i]
		newComb := currentComb.copy(num)
		if newComb.sum == c.target {
			// targetと数値が一致
			c.combinations = append(c.combinations, newComb.combination)
			continue
		} else if newComb.sum > c.target {
			// targetを超えた
			continue
		}
		// targetを超えなければ再探索
		c.search(newComb, i, depth)
	}
}

type combinationHolder struct {
	combination []int
	sum         int
}

func newCombinationHolder(combination []int, sum int) *combinationHolder {
	return &combinationHolder{
		combination: combination,
		sum:         sum,
	}
}

func (c *combinationHolder) copy(num int) *combinationHolder {
	// sliceをコピーして要素を追加
	newComb := make([]int, len(c.combination))
	copy(newComb, c.combination)
	newComb = append(newComb, num)

	// sumを加算
	newSum := c.sum + num
	return newCombinationHolder(newComb, newSum)
}

func newCombination(currentComb []int, num int) []int {
	newConb := make([]int, len(currentComb))
	copy(newConb, currentComb)
	newConb = append(newConb, num)
	return newConb
}
