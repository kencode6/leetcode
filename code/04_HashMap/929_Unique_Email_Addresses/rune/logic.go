package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/unique-email-addresses/

func numUniqueEmails(emails []string) int {
	emailmap := make(map[string]interface{})
	for _, email := range emails {
		convEmail := convEmail(email)
		if _, ok := emailmap[convEmail]; !ok {
			emailmap[convEmail] = new(interface{})
		}
	}
	return len(emailmap)
}

func convEmail(email string) string {
	convRns := []rune{}
	isPlusSkip := false
	isAfterAtmark := false
	for _, rn := range email {
		if isAfterAtmark {
			// @以降はappend
			convRns = append(convRns, rn)
			continue
		}

		// @以前
		if rn == '@' {
			// @を発見したらフラグを立てる
			isAfterAtmark = true
			convRns = append(convRns, rn)
			continue
		}

		if isPlusSkip {
			continue
		}

		if rn == '.' {
			continue
		}

		if rn == '+' {
			isPlusSkip = true
			continue
		}

		// @前の通常の文字列
		convRns = append(convRns, rn)
	}
	return string(convRns)
}
