package logic

import "sort"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	countNumMap := make(map[int]int)
	for _, num := range nums {
		// マッピングカウント登録
		count, ok := countNumMap[num]
		if !ok {
			count = 0
		}
		count++
		countNumMap[num] = count
	}

	// オブジェクトに変換
	numVals := []*NumVal{}
	for num, count := range countNumMap {
		numVal := NewNumVal(num, count)
		numVals = append(numVals, numVal)
	}

	//　ソート
	sort.Slice(numVals, func(i, j int) bool {
		return numVals[i].count > numVals[j].count
	})

	retNums := []int{}
	for _, numVal := range numVals {
		if len(retNums) < k {
			retNums = append(retNums, numVal.num)
		} else {
			break
		}
	}
	return retNums
}

type NumVal struct {
	num   int
	count int
}

func NewNumVal(num int, count int) *NumVal {
	return &NumVal{
		num:   num,
		count: count,
	}
}
