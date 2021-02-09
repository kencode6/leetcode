package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/next-permutation/

/*
解法
[1, 5, 8, 4, 7, 6, 5, 3, 1]
を例に考える

1. 後ろから検索して前の要素が後の要素より値が小さくなるindexを見つける
[1, 5, 8, (4), 7, 6, 5, 3, 1]

2. 後ろから検索して先ほど検索した数(4)より大きいindexを見つける
[1, 5, 8, (4), 7, 6, (5), 3, 1]

3. 1,2で見つけた数値をswapする
[1, 5, 8, (5), 7, 6, (4), 3, 1]

4. 1で見つけたindex+1以降を開始と終了のindexを1つずつ狭めながらswapする
[1, 5, 8, [5], (7), 6, 4, 3, (1)]
[1, 5, 8, [5], (1), 6, 4, 3, (7)]
[1, 5, 8, [5], 1, (6), 4, (3), 7]
[1, 5, 8, [5], 1, (3), 4, (6), 7]	←答え
*/

func nextPermutation(nums []int) {
	// 1. 後ろから検索して前の要素が後の要素より値が小さくなるindexを見つける
	sIndex := 0
	isFind := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			sIndex = i - 1
			isFind = true
			break
		}
	}

	if isFind {
		// 2. 後ろから検索して先ほど検索した数(4)より大きいindexを見つける
		eIndex := len(nums) - 1
		sNum := nums[sIndex]
		for i := eIndex; i > 0; i-- {
			if sNum < nums[i] {
				eIndex = i
				break
			}
		}
		// 3. 1,2で見つけた数値をswapする
		if sIndex != eIndex {
			swap(nums, sIndex, eIndex)
		}
		// 4. 1で見つけたindex+1以降を開始と終了のindexを1つずつ狭めながら反転させる
		reverse(nums, sIndex+1)
	} else {
		// sIndexが見つからなかった場合は先頭から反転させる
		reverse(nums, 0)
	}
}

// reverse 開始と終了のindexを1つずつ狭めながらswapする
func reverse(nums []int, sIndex int) {
	i := sIndex
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

// swap numsの要素i,jを反転させる
func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
