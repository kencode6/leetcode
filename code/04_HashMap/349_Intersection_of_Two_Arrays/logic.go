package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	intersectNums := []int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			if isContains(intersectNums, num1) {
				break
			}
			if num1 == num2 {
				intersectNums = append(intersectNums, num1)
				break
			}
		}
	}
	return intersectNums
}

func isContains(nums []int, targetNum int) bool {
	for _, num := range nums {
		if targetNum == num {
			return true
		}
	}
	return false
}
