package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestCountComponents(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return ladderLength(inputs[0].(string), inputs[1].(string), inputs[2].([]string))
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
