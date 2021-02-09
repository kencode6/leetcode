package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestWordBreak(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return wordBreak(inputs[0].(string), inputs[1].([]string))
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
