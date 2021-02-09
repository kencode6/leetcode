package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// 初めの要素より最後の要素の方が大きい場合はロテートしていないので初めが最小
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	sIndex := 0
	eIndex := len(nums) - 1
	for {
		if sIndex > eIndex {
			break
		}

		mIndex := (sIndex + eIndex) / 2

		// 左の要素が右の要素より大きかったら答え
		if nums[mIndex] > nums[mIndex+1] {
			return nums[mIndex+1]
		}

		if nums[mIndex-1] > nums[mIndex] {
			return nums[mIndex]
		}

		// 次のINDEXへ
		if nums[mIndex] > nums[eIndex] {
			// midより右にあり
			sIndex = mIndex + 1
		} else {
			// midより左にあり
			eIndex = mIndex - 1
		}
	}
	return -1
}
