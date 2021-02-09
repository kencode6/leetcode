package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	treeSearcher := NewTreeSearcher()
	return treeSearcher.walkTree(root, 1)
}

type TreeSearcher struct {
	minDepth int
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{
		minDepth: -1,
	}
}

func (t *TreeSearcher) walkTree(tn *TreeNode, depth int) int {

	// fmt.Printf("val:%d, depth:%d, minDepth:%d\n", tn.Val, depth, t.minDepth)

	// ノード終端
	if tn.Left == nil && tn.Right == nil {
		fmt.Printf("終端 val:%d, depth:%d, minDepth:%d\n", tn.Val, depth, t.minDepth)
		if t.minDepth == -1 || t.minDepth > depth {
			t.minDepth = depth
			fmt.Printf("終端min変更 val:%d, depth:%d, minDepth:%d\n", tn.Val, depth, t.minDepth)
		}
		return t.minDepth
	}

	// 左に移動
	if tn.Left != nil {
		t.walkTree(tn.Left, depth+1)
	}

	// 右に移動
	if tn.Right != nil {
		t.walkTree(tn.Right, depth+1)
	}

	// fmt.Printf("親ノードに戻る val:%d, depth:%d, minDepth:%d\n", tn.Val, depth, t.minDepth)

	return t.minDepth
}
