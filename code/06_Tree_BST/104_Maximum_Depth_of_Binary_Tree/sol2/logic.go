package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// tree内をwalkして探索する方法。sol1と比べると冗長だが、
	// tree系問題の色々なパターンに対応でき、汎用性は高い。
	searcher := NewTreeSearcher()
	return searcher.MaxDepth(root)
}

type TreeSearcher struct {
	depth    int
	maxDepth int
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{
		depth:    1,
		maxDepth: 0,
	}
}

func (t *TreeSearcher) MaxDepth(root *TreeNode) int {
	t.walkTree(root, 1)
	return t.maxDepth
}

func (t *TreeSearcher) walkTree(tn *TreeNode, depth int) {

	fmt.Printf("val:%d, depth:%d, maxDepth:%d\n", tn.Val, t.depth, t.maxDepth)

	// node終端
	if tn.Left == nil && tn.Right == nil {
		fmt.Printf("終端 val:%d, depth:%d, maxDepth:%d\n", tn.Val, t.depth, t.maxDepth)
		if t.maxDepth < depth {
			t.maxDepth = depth
			return
		}
	}

	depth++
	if tn.Left != nil {
		// 左に移動
		t.depth++
		t.walkTree(tn.Left, depth)
	}

	if tn.Right != nil {
		// 右に移動
		t.depth++
		t.walkTree(tn.Right, depth)
	}
}
