package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestUniquePaths(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return uniquePaths(inputs[0].(int), inputs[1].(int))
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
