package exec

import (
	"testing"

	"github.com/kencode6/leetcode/tools/linked"
)

func TestCodeExecutorSimple(t *testing.T) {
	err := NewCodeExecutor(t).
		YmlFile("testcase/exec_simple_test.yml"). // ymlファイルを指定しない場合はtestcase.ymlが読み込まれる
		Executor(
			// 評価する関数の実行方法を定義(必須)
			func(inputs []interface{}) interface{} {
				// InputConverterで変換された値が引数として渡される。
				nums := inputs[0].([]int)
				return sumNums(nums)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

// コンバータを用いたCodeExecutor実行テスト
func TestCodeExecutorConverter(t *testing.T) {
	err := NewCodeExecutor(t).
		YmlFile("testcase/exec_converter_test.yml"). // ymlファイルを指定しない場合はtestcase.ymlが読み込まれる
		InputConverter(
			// inputs値を変換したい場合に設定(任意)
			func(index int, input interface{}) interface{} {
				// TestCaseのInputsの各要素がindex番号付きで引数として渡される。
				if index == 0 {
					// 第一引数をsliceからLinkedListに変換
					nums := input.([]int)
					return linked.SliceToListNode(nums)
				}
				return input
			},
		).
		Executor(
			// 評価する関数の実行方法を定義(必須)
			func(inputs []interface{}) interface{} {
				// InputConverterで変換された値が引数として渡される。
				ln := inputs[0].(*linked.ListNode)
				multiNum := inputs[1].(int)
				return multiLindedList(ln, multiNum)
			},
		).
		ExpectedConverter(
			// expected値を変換したい場合に設定(任意)
			func(expected interface{}) interface{} {
				// Executorで返却した戻り値が引数として渡される。
				nums := expected.([]int)
				return linked.SliceToListNode(nums)
			},
		).
		Validator(
			// 値の検証方法をカスタマイズしたい場合に設定(任意)
			// 設定しない場合はreflect.DeepEqual(convOutput, expected)で検証される
			// このテストケースではValidatorを設定しなくても検証できるが、テスト用に設定
			func(output interface{}, expected interface{}, input []interface{}) bool {
				// convOutput: OutputConverterの戻り値
				// expected: TestCaseのExpectedの値
				// convInputs: InputConverterの戻り値
				ln := output.(*linked.ListNode)
				expectedLn := expected.(*linked.ListNode)

				nums := ln.ToSlice()
				expectedNums := expectedLn.ToSlice()

				if len(nums) != len(expectedNums) {
					return false
				}
				for i := 0; i < len(nums); i++ {
					if nums[i] != expectedNums[i] {
						return false
					}
				}
				return true
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

// MultiInputConverterを用いたCodeExecutor実行テスト
func TestCodeExecutorMultiInput(t *testing.T) {
	err := NewCodeExecutor(t).
		YmlFile("testcase/exec_multiconverter_test.yml").
		MultiInputConverter(
			// inputs値を個別に変換ではなく、第一引数と第二引数を統合して変換したい場合に設定(任意)
			func(inputs []interface{}) interface{} {
				head := inputs[0].([]int)
				pos := inputs[1].(int)
				ln, _, _ := linked.SliceToCycleListNode(head, pos)
				return ln
			},
		).
		Executor(
			// 評価する関数の実行方法を定義(必須)
			func(inputs []interface{}) interface{} {
				// InputConverterで変換された値が引数として渡される。
				ln := inputs[0].(*linked.ListNode)
				return cycleListNodeInfos(ln)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

// expectedがcompositデータのCodeExecutor実行テスト
func TestExecConverterCompositData(t *testing.T) {
	err := NewCodeExecutor(t).
		YmlFile("testcase/exec_composit_data_test.yml").
		Executor(
			func(inputs []interface{}) interface{} {
				return groupingCharIndex(inputs[0].(string))
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}

func TestManualExecConverter(t *testing.T) {
	// テストケースをyamlファイルでなく手動で作成
	testData := NewTestData("Manual Converter Test (multi ListNode)").
		AddTestCase("case1",
			// inputs
			[]interface{}{
				[]int{1, 2, 3},
				2,
			},
			// expected
			[]int{2, 4, 6},
		).
		AddTestCase("case2",
			// inputs
			[]interface{}{
				[]int{},
				2,
			},
			// expected
			[]int{},
		)

	err := NewCodeExecutor(t).
		TestData(testData).
		InputConverter(
			// inputs値を変換したい場合に設定(任意)
			func(index int, input interface{}) interface{} {
				// TestCaseのInputsの各要素がindex番号付きで引数として渡される。
				if index == 0 {
					// 第一引数をsliceからLinkedListに変換
					nums := input.([]int)
					return linked.SliceToListNode(nums)
				}
				return input
			},
		).
		Executor(
			// 評価する関数の実行方法を定義(必須)
			func(inputs []interface{}) interface{} {
				// InputConverterで変換された値が引数として渡される。
				ln := inputs[0].(*linked.ListNode)
				multiNum := inputs[1].(int)
				return multiLindedList(ln, multiNum)
			},
		).
		ExpectedConverter(
			// expected値を変換したい場合に設定(任意)
			func(expected interface{}) interface{} {
				// Executorで返却した戻り値が引数として渡される。
				nums := expected.([]int)
				return linked.SliceToListNode(nums)
			},
		).
		Exec()

	if err != nil {
		t.Errorf("コードテスト実行中にエラーが発生しました。 error:%v", err)
	}
}
