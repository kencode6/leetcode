package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/merge-two-binary-trees/

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	builder := NewTreeSearcher()
	builder.mergeTrees(t1, t2)
	return builder.root
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

type TreeSearcher struct {
	root *TreeNode
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{}
}

func (s *TreeSearcher) mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	val1 := s.treeNodeVal(t1)
	val2 := s.treeNodeVal(t2)

	tn := NewTreeNode(val1 + val2)
	fmt.Printf("val1:%d, val2:%d\n", val1, val2)
	if s.root == nil {
		s.root = tn
	}

	tn.Left = mergeTrees(s.leftTreeNode(t1), s.leftTreeNode(t2))
	tn.Right = mergeTrees(s.rightTreeNode(t1), s.rightTreeNode(t2))
	return tn
}

func (s *TreeSearcher) treeNodeVal(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return t.Val
}

func (s *TreeSearcher) leftTreeNode(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	return t.Left
}

func (s *TreeSearcher) rightTreeNode(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	return t.Right
}
