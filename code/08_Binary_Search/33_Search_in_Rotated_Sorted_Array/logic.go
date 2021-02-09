package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	// 要素1の場合は特別対応
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	// まずロテートインデックスを見つける
	rotateIndex := rotateIndex(nums)

	// ロテートインデックスに応じた挟み込みのindexを取得
	sIndex, eIndex := startEndIndex(nums, rotateIndex, target)

	// 通常の２分探索
	return binarySearch(nums, sIndex, eIndex, target)
}

// rotateIndexを取得
func rotateIndex(nums []int) int {
	// 初めの要素より最後の要素の方が大きい場合はロテートしていないので初めが最小
	if nums[0] < nums[len(nums)-1] {
		return 0
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
			return mIndex + 1
		}

		if nums[mIndex-1] > nums[mIndex] {
			return mIndex
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

// 開始、終了indexを取得
func startEndIndex(nums []int, rotateIndex int, target int) (startIndex int, endIndex int) {
	if rotateIndex == 0 {
		// 開始〜終了
		startIndex = 0
		endIndex = len(nums) - 1
		return
	}

	if target <= nums[len(nums)-1] {
		// rotateIndex〜終了
		startIndex = rotateIndex
		endIndex = len(nums) - 1
	} else {
		// 開始〜rotateIndex
		startIndex = 0
		endIndex = rotateIndex - 1
	}
	return
}

// binarySearch 二分探索を行う
func binarySearch(nums []int, sIndex int, eIndex int, target int) int {
	for {
		if sIndex > eIndex {
			break
		}
		mIndex := (sIndex + eIndex) / 2
		if nums[mIndex] == target {
			return mIndex
		}
		if target < nums[mIndex] {
			eIndex = mIndex - 1
		} else {
			sIndex = mIndex + 1
		}
	}
	return -1
}
