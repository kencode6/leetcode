package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	"fmt"

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// 参考(Inorder, Preorder Postorderとは？)
// https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	builder := NewTreeBuilder()
	builder.build(preorder, inorder, 0)
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

func (b *TreeBuilder) Root() *TreeNode {
	return b.root
}

func NewTreeBuilder() *TreeBuilder {
	return &TreeBuilder{}
}

func (b *TreeBuilder) build(preorder []int, inorder []int, depth int) *TreeNode {

	depth++
	// fmt.Printf("preorder:%v, inorder:%v, depth:%d\n", preorder, inorder, depth)

	if len(preorder) != len(inorder) {
		panic(fmt.Sprintf("preorderとinorderの要素数が一致しません。preorder:%d inorder:%d\n", len(preorder), len(inorder)))
	}

	if len(preorder) == 0 {
		return nil
	}

	// preorderの先頭要素でnode作成
	nodeVal := preorder[0]
	tn := NewTreeNode(nodeVal)
	if b.root == nil {
		b.root = tn
	}

	// inorderindex取得
	inorderIndex := b.inorderIndex(inorder, nodeVal)

	// left, rightに分割
	lPreorder := preorder[1 : inorderIndex+1]
	lInorder := inorder[:inorderIndex]
	tn.Left = b.build(lPreorder, lInorder, depth)

	rPreorder := preorder[inorderIndex+1:]
	rInorder := inorder[inorderIndex+1:]
	tn.Right = b.build(rPreorder, rInorder, depth)

	return tn
}

func (b *TreeBuilder) inorderIndex(inorder []int, targetVal int) int {
	for index, val := range inorder {
		if val == targetVal {
			return index
		}
	}
	panic(fmt.Sprintf("targetVal:%dが見つかりませんでした。。", targetVal))
}
