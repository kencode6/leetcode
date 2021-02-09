package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	// runes に変換
	rns := []rune(s)
	longestLen := 0
	for i := 0; i < len(rns); i++ {
		startIndex := i

		// startIndexから検索して重複していない文字列を取得
		noRepStr := noRepeatString(rns, startIndex)

		if longestLen < len(noRepStr) {
			longestLen = len(noRepStr)
		}
	}
	return longestLen
}

// noRepeatString　startIndexから検索して重複していない文字列を返却する
func noRepeatString(rns []rune, startIndex int) string {
	// 重複文字判定用のmap
	runeMap := make(map[rune]interface{})

	for j := startIndex; j < len(rns); j++ {
		rn := rns[j]
		if _, exists := runeMap[rn]; exists {
			// 重複文字を発見
			return string(rns[startIndex:j])
		}
		runeMap[rn] = new(interface{})
	}
	// 重複が無い場合は終端まで
	return string(rns[startIndex:])
}
