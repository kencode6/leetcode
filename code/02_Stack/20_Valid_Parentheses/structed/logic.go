package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	blacketValidator := NewBlacketValidator().
		Regist("(", ")").
		Regist("{", "}").
		Regist("[", "]")

	isValid := blacketValidator.Validate(s)
	return isValid
}

// BlacketValidator Blacket検証バリデーター
type BlacketValidator struct {
	blacketStarts map[string]string
	blacketEnds   map[string]string
}

// NewBlacketValidator Blacket検証用インスタンスを生成します。
func NewBlacketValidator() *BlacketValidator {
	return &BlacketValidator{
		blacketStarts: make(map[string]string),
		blacketEnds:   make(map[string]string),
	}
}

// Regist Blacketの対を登録します。
func (b *BlacketValidator) Regist(start string, end string) *BlacketValidator {
	b.blacketStarts[start] = end
	b.blacketEnds[end] = start
	return b
}

// Validate 与えられた文字列に対してBlacketの整合性を検証します。
func (b *BlacketValidator) Validate(s string) bool {

	tokenStack := []string{}

	for _, ch := range s {
		token := string(ch)

		// 開始トークンの場合はstackに積む
		if _, ok := b.blacketStarts[token]; ok {
			tokenStack = append(tokenStack, token)
			continue
		}
		// 終了トークンの場合
		startToken, ok := b.blacketEnds[token]
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
