package logic

import (
	"testing"

	"github.com/kencode6/leetcode/tools/exec"
)

func TestKthLargest(t *testing.T) {
	err := exec.NewCodeExecutor(t).
		Executor(
			func(inputs []interface{}) interface{} {
				// structのテストの為、他のテストケースと検証方法が異なる
				operationNames := inputs[0].([]string)
				inputNums := inputs[1].([][][]int)
				if len(operationNames) != len(inputNums) {
					t.Errorf("operationNamesとinputsの要素数数が一致しません operationNames:%d, inputs:%d", len(operationNames), len(inputs))
				}

				outputNums := []*int{}

				var kthLargest *KthLargest
				for i := 0; i < len(operationNames); i++ {
					operationName := operationNames[i]
					if operationName == "KthLargest" {
						// インスタンス生成オペレーション
						initInput := inputNums[i]
						k := initInput[0][0]
						nums := initInput[1]
						instance := Constructor(k, nums)
						kthLargest = &instance
						outputNums = append(outputNums, nil)
					} else if operationName == "add" {
						// 関数呼び出しオペレーション
						val := inputNums[i][0][0]
						retNum := kthLargest.Add(val)
						outputNums = append(outputNums, &retNum)
					} else {
						t.Errorf("未定義のoperation名です:%s", operationName)
					}
				}
				return outputNums
			},
		).Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
