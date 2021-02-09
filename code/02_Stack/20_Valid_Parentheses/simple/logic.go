package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	blacketStarts := make(map[string]string)
	blacketStarts["("] = ")"
	blacketStarts["{"] = "}"
	blacketStarts["["] = "]"

	blacketEnds := make(map[string]string)
	for key, val := range blacketStarts {
		blacketEnds[val] = key
	}

	tokenStack := []string{}

	for _, ch := range s {
		token := string(ch)

		// 開始トークンの場合はstackに積む
		if _, ok := blacketStarts[token]; ok {
			tokenStack = append(tokenStack, token)
			continue
		}
		// 終了トークンの場合
		startToken, ok := blacketEnds[token]
		if !ok {
			// 不明なトークン
			return false
		}
		// 開始トークンなしで終了トークンがあった場合
		if len(tokenStack) == 0 {
			return false
		}

		//  tokenStackをpopして終端要素と比較
		lastToken := tokenStack[len(tokenStack)-1]
		tokenStack = tokenStack[:len(tokenStack)-1]
		if startToken != lastToken {
			// 挟み込みが不整合
			return false
		}
	}

	// スタックが残っていた場合
	if len(tokenStack) > 0 {
		return false
	}
	return true
}
