package logic

import "sort"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	// 文字列グルーピング用マッピング生成
	sortedStrMaps := make(map[string][]string)

	// 順序保持用のソートキー生成
	sortedStrs := []string{}
	for _, str := range strs {

		// 文字列を一意判定するためのソートキーを取得
		sortedStr := sortedStr(str)

		strMapVals := sortedStrMaps[sortedStr]
		if strMapVals == nil {
			// 初回登録の場合
			strMapVals = []string{}
			sortedStrMaps[sortedStr] = strMapVals

			sortedStrs = append(sortedStrs, sortedStr)
		}
		strMapVals = append(strMapVals, str)
		sortedStrMaps[sortedStr] = strMapVals
	}

	// sortedStrMapsを結果用の[][]stringに変換
	strGroups := [][]string{}
	for _, sortedStr := range sortedStrs {
		strMapVals := sortedStrMaps[sortedStr]
		strGroups = append(strGroups, strMapVals)
	}
	return strGroups
}

// sortedStr 文字をアルファベット順にソートします。
func sortedStr(s string) string {
	runes := []rune{}
	for _, rn := range s {
		runes = append(runes, rn)
	}
	sort.Slice(runes,
		func(i, j int) bool {
			return runes[i] < runes[j]
		})
	sortStr := string(runes)
	return sortStr
}
