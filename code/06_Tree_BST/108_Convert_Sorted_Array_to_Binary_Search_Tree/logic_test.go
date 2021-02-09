package logic

import (
	"log"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/tree"
)

func TestSortedArrayToBST(t *testing.T) {
	err := exec.NewCodeExecutor(t).
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
				tn := sortedArrayToBST(inputs[0].([]int))
				log.Println("output tree:")
				tn.PrintTree()
				return tn
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
