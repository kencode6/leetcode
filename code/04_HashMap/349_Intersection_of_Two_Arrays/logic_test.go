package logic

import (
	"reflect"
	"sort"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestIntersection(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return intersection(inputs[0].([]int), inputs[1].([]int))
			},
		).
		Validator(
			func(output interface{}, expected interface{}, inputs []interface{}) bool {
				// []intの結果を順不同でも一致判定できるようにソートする
				sort.Sort(sort.IntSlice(output.([]int)))
				sort.Sort(sort.IntSlice(expected.([]int)))
				return reflect.DeepEqual(output, expected)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
