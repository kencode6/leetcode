package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestHasPathSum(t *testing.T) {
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
		Executor(
			func(inputs []interface{}) interface{} {
				root := inputs[0].(*tree.TreeNode)
				targetSum := inputs[1].(int)
				return hasPathSum(root, targetSum)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
