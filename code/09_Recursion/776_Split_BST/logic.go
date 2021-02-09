package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/split-bst/

func splitBST(root *TreeNode, v int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	}

	builder := NewTreeBuilder(v)
	return builder.BuildSplitTree(root)
}

type TreeBuilder struct {
	splitValue        int
	rootLowerTreeNode *TreeNode
	rootUpperTreeNode *TreeNode

	currentLowerTreeNode *TreeNode
	currentUpperTreeNode *TreeNode
}

func NewTreeBuilder(splitValue int) *TreeBuilder {
	return &TreeBuilder{splitValue: splitValue}
}

func (t *TreeBuilder) BuildSplitTree(root *TreeNode) []*TreeNode {
	t.walkTree(root, 1)
	return []*TreeNode{t.rootLowerTreeNode, t.rootUpperTreeNode}
}

func (t *TreeBuilder) walkTree(tn *TreeNode, depth int) {
	// 計算量は木の深さである。分割値に応じて切断しながら左右を探索していく。
	depth++
	if tn.Val <= t.splitValue {
		t.currentLowerTreeNode = tn

		// ルートが未設定の場合
		if t.rootLowerTreeNode == nil {
			t.rootLowerTreeNode = tn
		}

		// 右検証
		// ※左側はsplitValueより小さいので検証不要
		if tn.Right != nil {
			tnRight := tn.Right

			// 右切り離し判定
			if tn.Right.Val > t.splitValue {
				// 切断
				fmt.Printf("lower右切断 val:%d depth:%d \n", tn.Val, depth)
				tn.Right = nil

				if t.currentUpperTreeNode != nil {
					// 接続
					fmt.Printf("upper左接続 val:%d depth:%d \n", tn.Val, depth)
					t.currentUpperTreeNode.Left = tnRight
					t.currentUpperTreeNode = tnRight
				}
			}
			// 右に入る
			t.walkTree(tnRight, depth)
		}

	} else {
		t.currentUpperTreeNode = tn

		// ルートが未設定の場合
		if t.rootUpperTreeNode == nil {
			t.rootUpperTreeNode = tn
		}

		// 左検証
		// ※右側はsplitValueより大きいので検証不要
		if tn.Left != nil {
			tnLeft := tn.Left

			// 左切り離し判定
			if tn.Left.Val <= t.splitValue {
				// 切断
				fmt.Printf("upper左切断 val:%d depth:%d \n", tn.Val, depth)
				tn.Left = nil

				if t.currentLowerTreeNode != nil {
					//　接続
					fmt.Printf("lower右接続 val:%d depth:%d \n", tn.Val, depth)
					t.currentLowerTreeNode.Right = tnLeft
					t.currentLowerTreeNode = tnLeft
				}
			}
			// 左に入る
			t.walkTree(tnLeft, depth)
		}
	}
}
