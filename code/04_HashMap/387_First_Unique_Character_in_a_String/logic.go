package logic

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	// 各runeの出現数を保持するmapを用意
	rnCountMap := make(map[rune]int)
	for _, rn := range s {
		if count, ok := rnCountMap[rn]; ok {
			rnCountMap[rn] = count + 1
		} else {
			rnCountMap[rn] = 1
		}
	}

	// 出現数が1つのruneを見つける
	minIndex := -1
	for i, rn := range s {
		count := rnCountMap[rn]
		if count == 1 {
			minIndex = i
			break
		}
	}
	return minIndex
}
