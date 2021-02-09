package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestNextPermutation(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				nums := inputs[0].([]int)
				nextPermutation(nums)
				return nums
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
