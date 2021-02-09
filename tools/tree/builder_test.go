package tree

import (
	"reflect"
	"testing"

	"github.com/kencode6/leetcode/tools/utils"
)

const nilNum = -10000

func TestTreeNode(t *testing.T) {

	nums := []int{3, 9, 20, nilNum, nilNum, 15, 7, 11, 12}
	pNums := utils.ToPointerNums(nums, nilNum)
	/*
	   		想定されるTreeNode
	                  3
	      ┌───────────┴───────────┐
	      9                      20
	                        ┌─────┴─────┐
	                       15           7
	                     ┌──┴──┐
	                    11    12
	*/

	// 手動で答え合わせ用のTreeNode作成
	ansTn := newTreeNode(3)
	tn9 := newTreeNode(9)
	tn20 := newTreeNode(20)
	ansTn.Left = tn9
	ansTn.Right = tn20
	tn15 := newTreeNode(15)
	tn7 := newTreeNode(7)
	tn20.Left = tn15
	tn20.Right = tn7
	tn11 := newTreeNode(11)
	tn12 := newTreeNode(12)
	tn15.Left = tn11
	tn15.Right = tn12

	// pNumsからTreeNodeに変換
	builder := newTreeBuilder()
	tn, err := builder.Build(pNums)
	if err != nil {
		t.Error(err)
	}

	// 一致しているか確認
	if !reflect.DeepEqual(tn, ansTn) {
		t.Errorf("TreeNodeのデータが生成元データと一致しません :%s", utils.ToJSONString(tn))
	}

	// 最大深さ確認
	depth := tn.MaxDepth()
	if depth != 4 {
		t.Errorf("TreeNodeの深さが正しくありません")
	}

	// treenodeを標準出力
	tn.PrintTree()

	// TreeNodeからpNumsに変換
	convPNums := tn.ToPointerNums()

	// 一致しているか確認
	if !reflect.DeepEqual(pNums, convPNums) {
		t.Errorf("TreeNodeのデータが生成元データと一致しません :%v", convPNums)
	}
}

func TestPrintTree(t *testing.T) {
	nums := []int{0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1}
	pNums := utils.ToPointerNums(nums, nilNum)
	tn := CreateTreeNode(pNums)
	tn.PrintTree()
}
