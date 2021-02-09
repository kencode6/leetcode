package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	"fmt"

	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	ln := head

	builder := NewListNodeBuilder()

	var prevLn *ListNode
	isDuplicate := false

	for {
		fmt.Printf("ln.Val:%d, isDuplicate:%t\n", ln.Val, isDuplicate)

		if prevLn != nil {
			//　2回目以降
			if prevLn.Val != ln.Val {
				// 前の要素と重複無し
				if !isDuplicate {
					builder.AddVal(prevLn.Val)
					fmt.Printf("追加 ln.Val:%d, convLn.Val:%d isDuplicate:%t\n", ln.Val, builder.Current().Val, isDuplicate)
				}
				isDuplicate = false
			} else {
				// 前の要素と重複した場合フラグを立てる
				isDuplicate = true
			}
		}

		if ln.Next == nil {
			if prevLn == nil || ln.Val != prevLn.Val {
				builder.AddVal(ln.Val)
				fmt.Printf("追加 ln.Val:%d, convLn.Val:%d isDuplicate:%t\n", ln.Val, builder.Current().Val, isDuplicate)
			}
			break
		}
		prevLn = ln
		ln = ln.Next
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
