package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(s int, nums []int) int {
	isFirst := true
	minCount := 0
	for i := 0; i < len(nums); i++ {
		// startIndexから加算してlimitSumを超えるまでの要素数を取得
		startIndex := i
		count, isFind := subArrayCount(nums, startIndex, s)
		if !isFind {
			// 見つからなかった場合
			continue
		}

		if isFirst {
			// 初回登録
			isFirst = false
			minCount = count
			continue
		}

		if minCount > count {
			// 最小値更新
			minCount = count
		}
	}
	return minCount
}

// subArrayCount startIndexから加算してlimitSumを超えるまでの要素数を返却します
func subArrayCount(nums []int, startIndex int, limitSum int) (count int, isFind bool) {
	sum := 0
	for i := startIndex; i < len(nums); i++ {
		count++
		sum += nums[i]
		if sum >= limitSum {
			isFind = true
			return
		}
	}
	isFind = false
	return
}
