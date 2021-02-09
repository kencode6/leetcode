package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestIsValid(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				s := inputs[0].(string)
				return isValid(s)
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
