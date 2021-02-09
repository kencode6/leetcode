package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	treeSearcher := NewTreeSearcher()
	treeSearcher.walkTree(root, 0)
	return treeSearcher.convDepthVals()
}

type TreeSearcher struct {
	depthMapVals map[int][]int
	maxDepth     int
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{
		depthMapVals: make(map[int][]int),
	}
}

func (s *TreeSearcher) walkTree(tn *TreeNode, depth int) {

	depth++

	fmt.Printf("val:%d, depth:%d\n", tn.Val, depth)
	s.setNode(tn.Val, depth)

	// 終端
	if tn.Left == nil && tn.Right == nil {
		return
	}

	// 左
	if tn.Left != nil {
		s.walkTree(tn.Left, depth)
	}

	// 右
	if tn.Right != nil {
		s.walkTree(tn.Right, depth)
	}
}

func (s *TreeSearcher) setNode(val int, depth int) {
	// ノード登録
	vals, ok := s.depthMapVals[depth]
	if !ok {
		vals = []int{}
	}

	isReverse := depth%2 == 0
	if isReverse {
		vals = append([]int{val}, vals...)
	} else {
		vals = append(vals, val)
	}

	s.depthMapVals[depth] = vals

	// 深さ更新
	if s.maxDepth < depth {
		s.maxDepth = depth
	}
}

func (s *TreeSearcher) convDepthVals() [][]int {
	valsSlice := [][]int{}
	for depth := 1; depth <= s.maxDepth; depth++ {
		vals := s.depthMapVals[depth]
		valsSlice = append(valsSlice, vals)
	}
	return valsSlice
}
