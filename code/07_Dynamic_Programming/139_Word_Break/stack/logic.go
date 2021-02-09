package logic

import (
	"fmt"
	"strings"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/word-break/

func wordBreak(s string, words []string) bool {
	stack := []*SelectedData{}
	stack = append(stack, NewSelectedData(s, []string{}))

	cache := make(map[string]bool)
	for len(stack) > 0 {

		currentData := stack[0]
		stack = stack[1:]

		if _, ok := cache[currentData.s]; ok {
			// すでに評価してNGワードの場合
			continue
		}

		fmt.Printf("stackNum:%d, s:%s\n", len(stack), currentData.s)

		for _, word := range words {
			if currentData.s == word {
				// 完全一致した場合、完了
				newWords := append(currentData.selectedWords, word)
				fmt.Printf("splitWords:%v", newWords)
				return true
			}

			if strings.HasPrefix(currentData.s, word) {
				//先頭が一致した場合新しいデータをスタックに登録
				newS := currentData.s[len(word):]
				newWords := append(currentData.selectedWords, word)
				newData := NewSelectedData(newS, newWords)
				stack = append(stack, newData)
			}
		}

		// 一致しなかった場合NGワードとしてキャッシュに登録
		cache[currentData.s] = false
	}
	return false
}

type SelectedData struct {
	s             string
	selectedWords []string // 選択されたワードを保持
}

func NewSelectedData(s string, selectedWords []string) *SelectedData {
	cWords := copyWords(selectedWords)
	return &SelectedData{
		s:             s,
		selectedWords: cWords,
	}
}

func copyWords(words []string) []string {
	cWords := make([]string, len(words))
	copy(cWords, words)
	return cWords
}
