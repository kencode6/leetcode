package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
	"github.com/kencode6/leetcode/tools/linked"
)

func TestDetectCycle(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		MultiInputConverter(
			func(inputs []interface{}) interface{} {
				head := inputs[0].([]int)
				pos := inputs[1].(int)
				ln, _, _ := linked.SliceToCycleListNode(head, pos)
				return ln
			},
		).
		Executor(
			func(inputs []interface{}) interface{} {
				head := inputs[0].(*linked.ListNode)
				return detectCycle(head)
			},
		).
		Validator(
			func(output interface{}, expected interface{}, convInputs []interface{}) bool {
				outputLn := output.(*linked.ListNode)
				convInputLn := convInputs[0].(*linked.ListNode)
				expectedLnIndex := expected.(int)
				expectedLn := convInputLn.AtIndex(expectedLnIndex)
				return outputLn == expectedLn
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
