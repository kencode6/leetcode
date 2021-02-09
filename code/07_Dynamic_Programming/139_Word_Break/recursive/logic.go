package logic

import (
	"fmt"
	"strings"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/word-break/

func wordBreak(s string, words []string) bool {
	searcher := NewWordSearcher(words)
	return searcher.IsWordBreak(s)
}

type WordSearcher struct {
	words []string
	cache map[string]bool // 同一の演算を行わないようにするためのキャッシュ
}

func NewWordSearcher(words []string) *WordSearcher {
	return &WordSearcher{
		words: words,
		cache: make(map[string]bool),
	}
}

func (w *WordSearcher) IsWordBreak(s string) bool {
	return w.isWordBreak(s, 1)
}

func (w *WordSearcher) isWordBreak(s string, depth int) bool {
	if isWordBleak, ok := w.cache[s]; ok {
		// すでにキャッシュに登録されている場合
		return isWordBleak
	}

	fmt.Printf("s:%s depth:%d\n", s, depth)
	depth++

	for _, word := range w.words {
		// 先頭文字列が一致したら一致した文字列をトリムして再帰的に検証する。
		if strings.HasPrefix(s, word) {
			if s == word {
				// 一致した場合
				return true
			}
			s2 := s[len(word):]
			isWordBleak := w.isWordBreak(s2, depth)
			if isWordBleak {
				// 一致した場合
				return true
			}
		}
	}
	// 上記のループ内でトリムしきれなかった場合はwordbreakでない。
	// 同一の検証を行わないようにするためにキャッシュに登録
	w.cache[s] = false
	return false
}
