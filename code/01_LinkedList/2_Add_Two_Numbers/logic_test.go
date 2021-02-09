package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/linked"
)

func TestAddTwoNumbers(t *testing.T) {
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
				ln1 := inputs[0].(*linked.ListNode)
				ln2 := inputs[1].(*linked.ListNode)
				return addTwoNumbers(ln1, ln2)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
