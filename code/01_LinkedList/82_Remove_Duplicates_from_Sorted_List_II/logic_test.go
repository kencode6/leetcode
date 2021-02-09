package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/linked"
)

func TestDeleteDuplicates(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		InputConverter(
			func(index int, input interface{}) interface{} {
				return linked.SliceToListNode(input.([]int))
			},
		).
		ExpectedConverter(
			func(expected interface{}) interface{} {
				return linked.SliceToListNode(expected.([]int))
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				head := inputs[0].(*linked.ListNode)
				return deleteDuplicates(head)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
