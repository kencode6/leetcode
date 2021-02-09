package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/subsets/

// 前回のサブセットの各sliceに新しい数値を加算していく
// nums [1,2,3]の場合
// 初回
// []
// 1を加算
// [],[1]
// 2を加算
// [],[1], [2], [1,2]
// 3を加算
// [],[1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]

func subsets(nums []int) [][]int {

	prevSubsets := [][]int{}
	prevSubsets = append(prevSubsets, []int{})

	var newSubsets [][]int
	for i := 0; i < len(nums); i++ {
		// 前回のsubsetの各sliceにnumを追加したものを新しいsubsetとする
		num := nums[i]
		newSubsets = copySubsets(prevSubsets)
		for _, ns := range prevSubsets {
			newNums := copyNums(ns)
			newNums = append(newNums, num)
			newSubsets = append(newSubsets, newNums)
		}
		prevSubsets = newSubsets
	}
	return newSubsets
}

func copySubsets(subsets [][]int) [][]int {
	cSubsets := [][]int{}
	for _, ns := range subsets {
		cSubsets = append(cSubsets, copyNums(ns))
	}
	return cSubsets
}

func copyNums(nums []int) []int {
	cNums := make([]int, len(nums))
	copy(cNums, nums)
	return cNums
}
