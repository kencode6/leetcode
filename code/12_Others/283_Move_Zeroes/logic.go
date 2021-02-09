package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	// slice返却でなく引数のnumsを書き換える。
	// ゼロで無い値をsliceに保持する。
	nonZeroNums := []int{}
	for _, num := range nums {
		if num != 0 {
			nonZeroNums = append(nonZeroNums, num)
		}
	}

	// 前方にゼロでない値をつめて後方にゼロを入れる
	for i := 0; i < len(nums); i++ {
		if i < len(nonZeroNums) {
			nums[i] = nonZeroNums[i]
		} else {
			nums[i] = 0
		}
	}
}
