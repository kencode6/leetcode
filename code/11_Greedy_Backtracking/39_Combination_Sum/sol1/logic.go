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
	return &CombinationSearcher{
		candidates:   candidates,
		target:       target,
		combinations: [][]int{},
	}
}

func (c *CombinationSearcher) Search() [][]int {
	currentComb := newCombinationHolder([]int{}, 0)
	c.search(currentComb, 1)
	return c.combinations
}

func (c *CombinationSearcher) search(currentComb *combinationHolder, depth int) {
	depth++
	for _, num := range c.candidates {
		newComb := currentComb.copy(num)
		if newComb.sum == c.target {
			// targetと数値が一致
			isAdd := c.addCombination(newComb.combination)
			if isAdd {
				// fmt.Printf("match comb:%v sum:%d depth:%d\n", newComb.combination, newComb.sum, depth)
			}
			continue
		} else if newComb.sum > c.target {
			// targetを超えた
			continue
		}
		// targetを超えなければ再探索
		c.search(newComb, depth)
	}
}

func (c *CombinationSearcher) addCombination(newCombination []int) bool {
	// 組合せを一意にするためにソート
	sort.Sort(sort.IntSlice(newCombination))

	// 既存のcombinationの場合、追加しない
	for _, comb := range c.combinations {
		if isSameCombination(comb, newCombination) {
			return false
		}
	}

	// 新しいcombinationを追加
	c.combinations = append(c.combinations, newCombination)
	return true
}

func isSameCombination(sComb []int, dComb []int) bool {
	if len(sComb) != len(dComb) {
		return false
	}

	for i := 0; i < len(sComb); i++ {
		if sComb[i] != dComb[i] {
			return false
		}
	}
	return true
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
