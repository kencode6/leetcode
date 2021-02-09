package logic

import "fmt"

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	generator := NewParenthesisGenerator(n)
	return generator.Generate()
}

type ParenthesisGenerator struct {
	num           int
	parenthesises []string
}

func NewParenthesisGenerator(num int) *ParenthesisGenerator {
	return &ParenthesisGenerator{num: num}
}

func (p *ParenthesisGenerator) Generate() []string {
	p.generate("", 0, 0, 1)
	return p.parenthesises
}

func (p *ParenthesisGenerator) generate(currentStr string, leftNum int, rightNum int, depth int) {
	// 左カッコと右カッコのindexを別々に動かすのがコツ
	fmt.Printf("currentStr:%s, leftNum:%d, rightNum:%d, depth:%d\n", currentStr, leftNum, rightNum, depth)
	if len(currentStr) == p.num*2 {
		p.parenthesises = append(p.parenthesises, currentStr)
		return
	}

	if leftNum < p.num {
		// 左カッコを利用
		p.generate(currentStr+"(", leftNum+1, rightNum, depth+1)
	}
	if rightNum < leftNum {
		// 右カッコを利用。右カッコは左カッコを利用した個数以下であれば利用できる。
		p.generate(currentStr+")", leftNum, rightNum+1, depth+1)
	}
}
