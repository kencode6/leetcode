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

	searcher := NewTreeSearcher(root)
	searcher.walkTree()
	return searcher.minDepth
}

type TreeNodeData struct {
	tn    *TreeNode
	depth int
}

func NewTreeNodeData(tn *TreeNode, depth int) *TreeNodeData {
	return &TreeNodeData{
		tn:    tn,
		depth: depth,
	}
}

type TreeSearcher struct {
	queue    []*TreeNodeData
	minDepth int
}

func NewTreeSearcher(root *TreeNode) *TreeSearcher {
	rootData := NewTreeNodeData(root, 1)
	return &TreeSearcher{
		queue:    []*TreeNodeData{rootData},
		minDepth: -1,
	}
}

func (s *TreeSearcher) walkTree() {

	for len(s.queue) > 0 {

		tnd := s.queue[0]
		s.queue = s.queue[1:]

		tn := tnd.tn
		fmt.Printf("val:%d, depth:%d", tn.Val, tnd.depth)

		// 左
		if tn.Left != nil {
			newTnd := NewTreeNodeData(tn.Left, tnd.depth+1)
			s.queue = append(s.queue, newTnd)
		}

		// 右
		if tn.Right != nil {
			newTnd := NewTreeNodeData(tn.Right, tnd.depth+1)
			s.queue = append(s.queue, newTnd)
		}

		// 終端
		if tn.Left == nil && tn.Right == nil {
			fmt.Printf("end val:%d, depth:%d", tn.Val, tnd.depth)
			if s.minDepth == -1 || s.minDepth > tnd.depth {
				s.minDepth = tnd.depth
			}
		}
	}
}
