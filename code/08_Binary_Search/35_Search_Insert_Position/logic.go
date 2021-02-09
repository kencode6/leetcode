package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	// 開始と終了から中央の要素を比較する。
	// targetが中央要素と等しい場合、終了
	// targetが中央要素より大きい場合 中央+1〜終了 で検索
	// targetが中央要素より小さい場合 開始〜中央-1 で検索
	startIndex := 0
	endIndex := len(nums) - 1

	for {
		if startIndex > endIndex {
			break
		}
		// 中央のindexを取得
		midIndex := (startIndex + endIndex)
		midVal := nums[midIndex]

		if midVal == target {
			return midIndex
		} else if target < midVal {
			endIndex = midIndex - 1
		} else {
			startIndex = midIndex + 1
		}
	}
	// 一致しなかった場合、終了+1が挿入位置
	return endIndex + 1
}
