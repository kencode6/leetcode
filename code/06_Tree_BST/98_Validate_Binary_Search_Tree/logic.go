package logic

import (
	// leetcode環境にコードを貼り付けて実行可能にする為にパッケージ名省略記述を利用しています。
	. "github.com/kencode6/leetcode/tools/tree"
)

// 下記のleetcodeサイトにアクセスし、以下のコードを貼り付けて実行可能です。
// https://leetcode.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	treeSearcher := NewTreeSearcher()
	treeSearcher.walkTree(root, 0)
	return treeSearcher.IsValid
}

type TreeSearcher struct {
	IsValid bool
	nums    []int // 2分探索木の構造からソートsliceを作成
}

func NewTreeSearcher() *TreeSearcher {
	return &TreeSearcher{
		IsValid: true,
		nums:    []int{},
	}
}

func (s *TreeSearcher) walkTree(tn *TreeNode, depth int) {

	// 不整合の場合は終了
	if !s.IsValid {
		return
	}

	depth++

	// fmt.Printf("val:%d, depth:%d\n", tn.Val, depth)

	// 左
	if tn.Left != nil {
		s.walkTree(tn.Left, depth)

		// 左から戻ってきたら順序確定
		s.appendVal(tn.Val)
		// fmt.Printf("順序確定(左から戻った) val:%d, depth:%d nums:%v isValid:%t\n", tn.Val, depth, s.nums, s.IsValid)
	} else {
		// 左が無かったら順序確定
		s.appendVal(tn.Val)
		// fmt.Printf("順序確定(左無し) val:%d, depth:%d nums:%v isValid:%t\n", tn.Val, depth, s.nums, s.IsValid)
	}

	// 不整合の場合は終了
	if !s.IsValid {
		return
	}

	// 右
	if tn.Right != nil {
		s.walkTree(tn.Right, depth)
	}

	// 終端
	// if tn.Left == nil && tn.Right == nil {
	// 	fmt.Printf("終端 val:%d, depth:%d\n", tn.Val, depth)
	// }
}

func (s *TreeSearcher) appendVal(val int) {
	// 終端要素の方がval以上の場合は不整合
	if len(s.nums) > 0 && s.nums[len(s.nums)-1] >= val {
		s.IsValid = false
	}
	s.nums = append(s.nums, val)
}
