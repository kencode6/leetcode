package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/reverse-linked-list/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	lns := []*ListNode{}
	ln := head
	for {
		lns = append(lns, ln)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}

	var startLn *ListNode
	var prevLn *ListNode
	for i := 0; i < len(lns); i++ {
		l := lns[len(lns)-1-i]

		l.Next = nil
		if i == 0 {
			startLn = l
		}

		if prevLn != nil {
			prevLn.Next = l
		}
		prevLn = l
	}
	return startLn
}
