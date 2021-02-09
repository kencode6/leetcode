package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	builder := NewTreeBuilder()
	builder.walkTree(root, 0)
	return builder.Root()
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

type TreeBuilder struct {
	root *TreeNode
}

func NewTreeBuilder() *TreeBuilder {
	return &TreeBuilder{}
}

func (b *TreeBuilder) Root() *TreeNode {
	return b.root
}

func (b *TreeBuilder) walkTree(tn *TreeNode, depth int) *TreeNode {

	depth++

	fmt.Printf("val:%d, depth:%d \n", tn.Val, depth)

	newTn := NewTreeNode(tn.Val)
	if b.root == nil {
		b.root = newTn
	}

	if tn.Left != nil {
		fmt.Printf("左行きます\n")
		lNewTn := b.walkTree(tn.Left, depth)
		newTn.Right = lNewTn
	}

	if tn.Right != nil {
		fmt.Printf("右行きます\n")
		rNewTn := b.walkTree(tn.Right, depth)
		newTn.Left = rNewTn
	}

	if tn.Left == nil && tn.Right == nil {
		fmt.Printf("終端です\n")
	}
	return newTn
}
