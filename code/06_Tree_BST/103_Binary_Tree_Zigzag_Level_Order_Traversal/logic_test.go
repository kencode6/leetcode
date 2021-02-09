package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestZigzagLevelOrder(t *testing.T) {
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
				ln := inputs[0].(*tree.TreeNode)
				return zigzagLevelOrder(ln)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
