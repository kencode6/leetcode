package logic

import "strings"

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
	atIndex := strings.Index(email, "@")
	leftStr := email[0:atIndex]
	rightStr := email[atIndex:]

	plusIndex := strings.Index(email, "+")
	if plusIndex != -1 {
		leftStr = leftStr[:plusIndex]
	}
	leftStr = strings.ReplaceAll(leftStr, ".", "")

	convEmail := leftStr + rightStr
	return convEmail
}
