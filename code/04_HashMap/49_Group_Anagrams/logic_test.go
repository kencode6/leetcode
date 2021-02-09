package logic

import (
	"reflect"
	"sort"
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestGroupAnagrams(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				return groupAnagrams(inputs[0].([]string))
			},
		).
		Validator(
			func(output interface{}, expected interface{}, inputs []interface{}) bool {
				// [][]stringの結果を順不同でも一致判定できるようにソートする
				sortStringSlices(output.([][]string))
				sortStringSlices(expected.([][]string))
				return reflect.DeepEqual(output, expected)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

func sortStringSlices(strSlices [][]string) {
	// まず各グループをソート
	for _, strSlice := range strSlices {
		sort.Sort(sort.StringSlice(strSlice))
	}

	// 次にsliceをソート
	sort.Slice(strSlices, func(i, j int) bool {
		sStrs := strSlices[i]
		dStrs := strSlices[j]

		// 要素数の数が違う場合は少ない順
		if len(sStrs) != len(dStrs) {
			return len(sStrs) < len(dStrs)
		}
		// 要素数の数が同じ場合は初回の異なる要素の大小で比較
		for k := 0; k < len(sStrs); k++ {
			if sStrs[k] != dStrs[k] {
				return sStrs[k] < dStrs[k]
			}
		}
		return false
	})
}
