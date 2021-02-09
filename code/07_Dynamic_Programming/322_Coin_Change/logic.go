package logic

import (
	"sort"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 降順ソート
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	searcher := NewSearcher(coins)
	count := searcher.Search(amount)
	return count
}

type Searcher struct {
	coins           []int
	minCombCount    int         // 最小組合せ数
	findAmountCache map[int]int // 同一の演算を行わないようにするためのキャッシュ (key=金額 value=組合せ数)
}

func NewSearcher(coins []int) *Searcher {
	return &Searcher{
		coins:           coins,
		minCombCount:    -1,
		findAmountCache: make(map[int]int),
	}
}

func (s *Searcher) Search(amount int) int {
	s.search(amount, 0, 0)
	return s.minCombCount
}

func (s *Searcher) search(amount int, combCount int, depth int) {
	if cacheCombCount, ok := s.findAmountCache[amount]; ok {
		if cacheCombCount <= combCount {
			return
		}
	}

	// 同一の演算を行わないようにするためにキャッシュに登録
	s.findAmountCache[amount] = combCount

	depth++
	// fmt.Printf("amount:%d, combCount:%d, depth:%d\n", amount, combCount, depth)

	// 合計金額から登録されたcoinの額を引いて、再帰的に演算させる。
	for _, coin := range s.coins {
		if amount == coin {
			// 一致した場合
			newCombCount := combCount + 1
			if s.minCombCount == -1 || s.minCombCount > newCombCount {
				s.minCombCount = newCombCount
			}
			continue
		}

		if amount > coin {
			// 合計金額をコインの額で引いて再帰呼び出し
			newAmount := amount - coin
			newCombCount := combCount + 1
			s.search(newAmount, newCombCount, depth)
		}
	}
}
