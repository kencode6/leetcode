package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"
	"math"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

func maxDepth(tn *TreeNode) int {
	if tn == nil {
		return 0
	}

	leftDepth := maxDepth(tn.Left)
	rightDepth := maxDepth(tn.Right)
	fmt.Printf("val:%d leftDepth:%d, rightDepth:%d", tn.Val, leftDepth, rightDepth)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
