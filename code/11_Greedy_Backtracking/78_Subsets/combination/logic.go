package logic

import "fmt"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/subsets/

func subsets(nums []int) [][]int {
	newSubsets := [][]int{}
	newSubsets = append(newSubsets, []int{})
	for i := 0; i <= len(nums); i++ {
		// numsに対して組合せ数がiのサブセットを返却
		tmpSubsets := subsetsByCombNum(nums, i)
		for _, tmpNums := range tmpSubsets {
			newSubsets = append(newSubsets, tmpNums)
		}
	}
	return newSubsets
}

// subsetsByCombNum numsに対して組合せ数がcombNumの組合せを全て返却します
// 例)
// nums=[1,2,3], combNum=2の場合
// [1,2], [1,3], [2,3]
func subsetsByCombNum(nums []int, combNum int) [][]int {
	if combNum == 0 {
		return [][]int{}
	}
	return subsetsByIndex(nums, combNum, 0)
}

// subsetsByIndex 開始indexがstartIndexでnumsの中から組合せ数combNumのサブセットを再帰的に返却します。
// 例)
// nums=[1,2,3,4,5], cmobNum=3の場合 10通り存在するので再帰手続きを用いて以下のように要素を返却できる処理を検討する。
// 1,2,3
// 1,2,  4
// 1,2,    5
// 1,  3,4
// 1,  3,  5
// 1,    4,5
//   2,3,4
//   2,3,  5
//   2,  4,5
//     3,4,5
func subsetsByIndex(nums []int, combNum int, startIndex int) [][]int {
	if combNum == 1 {
		// 組合せ数が1つの場合
		newSubsets := [][]int{}
		for i := startIndex; i < len(nums); i++ {
			num := nums[i]
			newSubsets = append(newSubsets, []int{num})
		}
		fmt.Printf("first combNum:%d, startIndex:%d newSubsets:%v\n", combNum, startIndex, newSubsets)
		return newSubsets
	}

	newSubsets := [][]int{}
	maxIndex := len(nums) - combNum
	fmt.Printf("combNum:%d, startIndex:%d maxIndex:%d\n", combNum, startIndex, maxIndex)
	for i := startIndex; i <= maxIndex; i++ {
		num := nums[i]
		// 組合せ数を1つ引いて、indexを一つ進めたサブセットの取得を再帰呼び出し
		tmpSubsets := subsetsByIndex(nums, combNum-1, i+1)
		for _, tmpNums := range tmpSubsets {
			newNums := []int{num}
			newNums = append(newNums, tmpNums...)
			newSubsets = append(newSubsets, newNums)
		}
	}
	fmt.Printf("combNum:%d, startIndex:%d maxIndex:%d newSubsets:%v\n", combNum, startIndex, maxIndex, newSubsets)
	return newSubsets
}
