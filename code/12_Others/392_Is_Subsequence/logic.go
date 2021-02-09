package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}

	sRns := []rune(s)
	sIndex := 0
	for _, tRn := range t {
		// サブシーケンスチェック対象の文字を取得
		sRn := sRns[sIndex]
		if sRn == tRn {
			// 一致したら次のindexに進める
			sIndex++
			if sIndex == len(sRns) {
				// indexが終端までいったらサブシーケンスに含まれる
				return true
			}
		}
	}
	return false
}
