package logic

import (
	"reflect"
	"sort"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestSubsets(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return subsets(inputs[0].([]int))
			},
		).
		Validator(
			func(output interface{}, expected interface{}, inputs []interface{}) bool {
				// [][]intの結果を順不同でも一致判定できるようにソートする
				sortIntSlices(output.([][]int))
				sortIntSlices(expected.([][]int))
				return reflect.DeepEqual(output, expected)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

func sortIntSlices(numSlices [][]int) {
	// sliceをソート
	sort.Slice(numSlices, func(i, j int) bool {
		sNums := numSlices[i]
		dNums := numSlices[j]

		// 要素数の数が違う場合は少ない順
		if len(sNums) != len(dNums) {
			return len(sNums) < len(dNums)
		}
		// 要素数の数が同じ場合は初回の異なる要素の大小で比較
		for k := 0; k < len(sNums); k++ {
			if sNums[k] != dNums[k] {
				return sNums[k] < dNums[k]
			}
		}
		return false
	})
}
