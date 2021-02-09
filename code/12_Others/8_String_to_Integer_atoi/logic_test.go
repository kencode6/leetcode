package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestMyAtoi(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return myAtoi(inputs[0].(string))
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
