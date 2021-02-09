package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestMaxProfit(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return maxProfit(inputs[0].([]int))
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
