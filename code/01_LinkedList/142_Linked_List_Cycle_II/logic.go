package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/linked-list-cycle-ii/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prevLns := make(map[*ListNode]interface{})

	ln := head
	var targetLn *ListNode
	for ln.Next != nil {
		if _, ok := prevLns[ln]; ok {
			targetLn = ln
			break
		}

		prevLns[ln] = new(interface{})
		ln = ln.Next
	}
	return targetLn
}
