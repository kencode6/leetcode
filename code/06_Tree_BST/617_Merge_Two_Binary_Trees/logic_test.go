package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestMergeTrees(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				pNums := input.([]*int)
				tn := tree.CreateTreeNode(pNums)
				log.Printf("input[%d] tree:", index)
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
				tn1 := inputs[0].(*tree.TreeNode)
				tn2 := inputs[1].(*tree.TreeNode)
				mergeTn := mergeTrees(tn1, tn2)
				log.Println("output tree:")
				mergeTn.PrintTree()
				return mergeTn
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
