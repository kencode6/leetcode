package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	. "github.com/kencode6/leetcode/tools/linked"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/linked-list-cycle/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	prevLns := make(map[*ListNode]interface{})
	ln := head
	for ln.Next != nil {

		// 登録済みListNodeがあった場合、cycleになっている。
		if _, ok := prevLns[ln]; ok {
			return true
		}

		// 過去のListNodeを登録
		prevLns[ln] = new(interface{})

		ln = ln.Next
	}
	return false
}
