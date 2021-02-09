package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	//　縦または横の文字列で、折れ曲がれ無しの場合
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	// zigzagの文字列を横に格納用
	rnlines := [][]rune{}
	for i := 1; i <= numRows; i++ {
		rnlines = append(rnlines, []rune{})
	}

	// zigzag文字をrnlinesに格納
	rowIndex := 0
	isUp := false // 上方向に移動
	rns := []rune(s)
	for _, rn := range rns {
		rnlines[rowIndex] = append(rnlines[rowIndex], rn)
		if isUp {
			rowIndex--
		} else {
			rowIndex++
		}

		if rowIndex == numRows-1 {
			isUp = true
		} else if rowIndex == 0 {
			isUp = false
		}
	}

	// 返却用文字列を統合
	retRns := []rune{}
	for _, rns := range rnlines {
		retRns = append(retRns, rns...)
	}
	return string(retRns)
}
