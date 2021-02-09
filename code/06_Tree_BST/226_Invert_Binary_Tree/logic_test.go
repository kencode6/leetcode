package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestInvertTree(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				pNums := input.([]*int)
				tn := tree.CreateTreeNode(pNums)
				log.Println("input tree:")
				tn.PrintTree()
				return tn
			},
		).
		ExpectedConverter(
			func(expected interface{}) interface{} {
				pNums := expected.([]*int)
				tn := tree.CreateTreeNode(pNums)
				log.Println("expected tree:")
				tn.PrintTree()
				return tn
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				tn := inputs[0].(*tree.TreeNode)
				invTn := invertTree(tn)
				log.Println("output tree:")
				invTn.PrintTree()
				return invTn
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
