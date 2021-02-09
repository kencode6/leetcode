package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/path-sum/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	searcher := NewTreeSearcher(sum)
	searcher.walkTree(root, 0, 0)
	return searcher.isFind
}

type TreeSearcher struct {
	targetSum int
	isFind    bool
}

func NewTreeSearcher(targetSum int) *TreeSearcher {
	return &TreeSearcher{targetSum: targetSum, isFind: false}
}

func (t *TreeSearcher) walkTree(tn *TreeNode, depth int, sum int) {

	// 見つけた場合はwalk終了
	if t.isFind {
		return
	}

	sum += tn.Val
	depth++

	fmt.Printf("val:%d depth:%d sum:%d\n", tn.Val, depth, sum)

	// 終端
	if tn.Left == nil && tn.Right == nil {
		fmt.Printf("終端 val:%d depth:%d sum:%d\n", tn.Val, depth, sum)
		if sum == t.targetSum {
			fmt.Printf("ノード発見\n")
			t.isFind = true
		}
		return
	}

	// 左へ移動
	if tn.Left != nil {
		t.walkTree(tn.Left, depth, sum)
	}

	//  右へ移動
	if tn.Right != nil {
		t.walkTree(tn.Right, depth, sum)
	}
}
