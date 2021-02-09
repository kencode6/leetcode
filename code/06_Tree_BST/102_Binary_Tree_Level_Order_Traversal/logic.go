package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	treeSearcher := NewTreeSearcher()
	treeSearcher.walkTree(root, 0)
	return treeSearcher.convDepthVals()
}

type TreeSearcher struct {
	depthMappingVals map[int][]int // key=depth, val=対象depthのノード値のslice
	maxDepth         int
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{
		depthMappingVals: make(map[int][]int),
		maxDepth:         0,
	}
}

func (s *TreeSearcher) walkTree(tn *TreeNode, depth int) {

	depth++

	fmt.Printf("val:%d, depth:%d\n", tn.Val, depth)

	// ノード登録
	s.setNodeVal(tn.Val, depth)

	// 終端
	if tn.Left == nil && tn.Right == nil {
		return
	}

	// 左へ移動
	if tn.Left != nil {
		s.walkTree(tn.Left, depth)
	}

	// 右へ移動
	if tn.Right != nil {
		s.walkTree(tn.Right, depth)
	}
}

func (s *TreeSearcher) setNodeVal(val int, depth int) {

	// ノード登録
	vals, ok := s.depthMappingVals[depth]
	if !ok {
		vals = []int{}
	}
	vals = append(vals, val)
	s.depthMappingVals[depth] = vals

	//　最大深さ更新
	if s.maxDepth < depth {
		s.maxDepth = depth
	}
}

func (s *TreeSearcher) convDepthVals() [][]int {
	valsSlice := [][]int{}
	for depth := 1; depth <= s.maxDepth; depth++ {
		vals := s.depthMappingVals[depth]
		valsSlice = append(valsSlice, vals)
	}
	return valsSlice
}
