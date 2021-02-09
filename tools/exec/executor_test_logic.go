package exec

import (
	"github.com/kencode6/leetcode/tools/linked"
)

// sumNums 配列をsumした結果を返却するテスト用関数
func sumNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// multiLindedList 掛け算を行ったLinkedlistを返却するテスト用関数
func multiLindedList(ln *linked.ListNode, multiNum int) *linked.ListNode {
	nums := ln.ToSlice()
	for i := 0; i < len(nums); i++ {
		nums[i] *= multiNum
	}
	return linked.SliceToListNode(nums)
}

// cycleListNodeInfos cycleしたListNodeの接続位置と終端長さを返却するテスト用関数
func cycleListNodeInfos(head *linked.ListNode) []int {
	if head == nil {
		return []int{-1, 0}
	}

	prevLns := make(map[*linked.ListNode]int)
	ln := head
	index := 0
	for ln.Next != nil {

		// 登録済みListNodeがあった場合、cycleになっている。
		if _, ok := prevLns[ln]; ok {
			jointIndex := prevLns[ln]
			// この時点でのindex番号は終端のindex+1になっているのでsizeと同じとなる。
			size := index
			return []int{jointIndex, size}
		}

		// 過去のListNodeを登録
		prevLns[ln] = index

		index++
		ln = ln.Next
	}
	jointIndex := -1
	size := index + 1
	return []int{jointIndex, size}
}

// groupingCharIndex 文字列を分解して文字ごとにindex配列に変換します
// aabcc → {"a":[1,2],"b":[3],"c":[4,5]}
func groupingCharIndex(line string) map[string][]int {
	chGroup := make(map[string][]int)
	index := 0
	for _, rn := range line {
		index++
		s := string(rn)
		indexes, ok := chGroup[s]
		if !ok {
			indexes = []int{}
		}
		indexes = append(indexes, index)
		chGroup[s] = indexes
	}
	return chGroup
}
