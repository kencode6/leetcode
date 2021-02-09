package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/subarray-sum-equals-k/

func subarraySum(nums []int, k int) int {
	totalCount := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		// i番目の要素から合計してkになる配列数をカウントする
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				// 一致したらtotalCountを加算する。要素がマイナスの場合もあるのでbreakしない
				totalCount++
			}
		}
	}
	return totalCount
}
