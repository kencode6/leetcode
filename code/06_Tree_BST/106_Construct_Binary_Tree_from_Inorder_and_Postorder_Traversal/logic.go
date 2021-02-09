package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// 参考(Inorder, Preorder Postorderとは？)
// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	builder := NewTreeBuilder()
	builder.build(inorder, postorder, 0)
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

func (b *TreeBuilder) build(inorder []int, postorder []int, depth int) *TreeNode {

	if len(inorder) != len(postorder) {
		panic(fmt.Errorf("len(inorder) != len(postorder) %d, %d", len(inorder), len(postorder)))
	}

	if len(inorder) == 0 {
		return nil
	}

	depth++

	// fmt.Printf("inorder:%v, postorder:%v, depth:%d\n", inorder, postorder, depth)

	// ノード追加
	nodeVal := postorder[len(postorder)-1]
	tn := NewTreeNode(nodeVal)
	if b.root == nil {
		b.root = tn
	}

	// index検索
	inorderIndex := b.searchIndex(inorder, nodeVal)

	// 左
	lInorder := inorder[:inorderIndex]
	lPostorder := postorder[:inorderIndex]
	tn.Left = b.build(lInorder, lPostorder, depth)

	// 右
	rInorder := inorder[inorderIndex+1:]
	rPostorder := postorder[inorderIndex : len(postorder)-1]
	tn.Right = b.build(rInorder, rPostorder, depth)

	return tn
}

func (b *TreeBuilder) searchIndex(inorder []int, targetVal int) int {

	for index, val := range inorder {
		if val == targetVal {
			return index
		}
	}
	panic(fmt.Errorf("targetVal not found inorder:%v, targetVal:%d", inorder, targetVal))
}
