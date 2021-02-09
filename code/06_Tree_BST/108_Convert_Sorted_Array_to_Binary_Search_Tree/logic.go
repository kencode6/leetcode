package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。

	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	treeBuilder := NewTreeBuilder(nums)
	root := treeBuilder.sortedArrayToBST()
	return root
}

type TreeBuilder struct {
	nums []int
	root *TreeNode
}

func NewTreeBuilder(nums []int) *TreeBuilder {
	return &TreeBuilder{
		nums: nums,
	}
}

func (t *TreeBuilder) sortedArrayToBST() *TreeNode {
	sIndex := 0
	eIndex := len(t.nums) - 1
	t.createTreeNode(sIndex, eIndex, 1)
	return t.root
}

func (t *TreeBuilder) createTreeNode(sIndex, eIndex, depth int) *TreeNode {

	// fmt.Printf("sIndex:%d eIndex:%d depth:%d\n", sIndex, eIndex, depth)

	if sIndex > eIndex {
		return nil
	}

	mIndex := (sIndex + eIndex) / 2
	mVal := t.nums[mIndex]

	// fmt.Printf("index:%d val:%d depth:%d\n", mIndex, mVal, depth)

	tn := NewTreeNode(mVal)
	if t.root == nil {
		t.root = tn
	}

	depth++

	tn.Left = t.createTreeNode(sIndex, mIndex-1, depth)
	tn.Right = t.createTreeNode(mIndex+1, eIndex, depth)

	return tn
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
