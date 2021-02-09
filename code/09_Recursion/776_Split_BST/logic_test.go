package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestSplitBST(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				if index == 0 {
					pNums := input.([]*int)
					tn := tree.CreateTreeNode(pNums)
					log.Println("input tree:")
					tn.PrintTree()
					return tn
				}
				return input
			},
		).
		ExpectedConverter(
			func(expected interface{}) interface{} {
				pNums := expected.([][]*int)
				lowerTn := tree.CreateTreeNode(pNums[0])
				log.Println("expected lower tree:")
				lowerTn.PrintTree()

				upperTn := tree.CreateTreeNode(pNums[1])
				log.Println("expected upper tree:")
				upperTn.PrintTree()
				return []*tree.TreeNode{lowerTn, upperTn}
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				tns := splitBST(inputs[0].(*tree.TreeNode), inputs[1].(int))
				lowerTn := tns[0]
				log.Println("output lower tree:")
				lowerTn.PrintTree()

				upperTn := tns[1]
				log.Println("output upper tree:")
				upperTn.PrintTree()
				return tns
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
