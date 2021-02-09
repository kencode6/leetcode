package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		// 1つめの数値を固定
		num1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			// 2つめの数値を動かしてtargetと一致したらindexを返却
			num2 := nums[j]
			sumNum := num1 + num2
			if target == sumNum {
				return []int{i, j}
			}
		}
	}
	return nil
}
