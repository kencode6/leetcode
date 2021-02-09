package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestMinMeetingRooms(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return minMeetingRooms(inputs[0].([][]int))
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
