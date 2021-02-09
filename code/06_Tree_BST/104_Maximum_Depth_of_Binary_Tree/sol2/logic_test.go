package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestMaxDepth(t *testing.T) {
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
		Executor(
			func(inputs []interface{}) interface{} {
				tn := inputs[0].(*tree.TreeNode)
				return maxDepth(tn)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
