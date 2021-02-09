package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	// len(nums)は1以上なので空チェック不要
	maxSumNum := nums[0]
	for i := 0; i < len(nums); i++ {
		sumNum := 0
		for j := i; j < len(nums); j++ {
			sumNum += nums[j]
			// fmt.Printf("i:%d, j:%d, sumNum:%d", i, j, sumNum)
			if maxSumNum < sumNum {
				maxSumNum = sumNum
			}
		}
	}
	return maxSumNum
}
