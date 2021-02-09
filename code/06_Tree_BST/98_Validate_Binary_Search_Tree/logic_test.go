package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestIsValidBST(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				pNums := input.([]*int)
				return tree.CreateTreeNode(pNums)
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				tn := inputs[0].(*tree.TreeNode)
				return isValidBST(tn)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
