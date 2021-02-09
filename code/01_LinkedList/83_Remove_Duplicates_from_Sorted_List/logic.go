package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	"fmt"

	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	builder := NewListNodeBuilder()

	ln := head
	builder.AddVal(ln.Val)
	prevVal := ln.Val

	for ln.Next != nil {

		ln = ln.Next

		fmt.Printf("val:%d prevVal:%d", ln.Val, prevVal)
		if prevVal != ln.Val {
			prevVal = ln.Val
			builder.AddVal(ln.Val)
		}
	}
	return builder.Head()
}

// NewListNode ListNodeを生成します。
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

// ListNodeBuilder ListNodeの生成処理を行います。
type ListNodeBuilder struct {
	headLn    *ListNode
	currentLn *ListNode
}

// NewListNodeBuilder ListNodeBuilderを生成します。
func NewListNodeBuilder() *ListNodeBuilder {
	return &ListNodeBuilder{}
}

// AddVal ListNodeの要素を追加します。
func (b *ListNodeBuilder) AddVal(val int) {
	newLn := NewListNode(val)
	if b.headLn == nil {
		// 初回
		b.headLn = newLn
	} else {
		// 2回目以降
		b.currentLn.Next = newLn
	}
	b.currentLn = newLn
}

// Head HeadのListNodeを返却します。
func (b *ListNodeBuilder) Head() *ListNode {
	return b.headLn
}

// Current CurrentのListNodeを返却します。
func (b *ListNodeBuilder) Current() *ListNode {
	return b.currentLn
}
