package exec

import (
	"fmt"
	"reflect"
)

// TestData testcase.ymlから読み込まれるデータ
type TestData struct {
	Name  string
	Cases []*TestCase
}

// TestCase testcase.ymlから読み込まれる各テストケース
type TestCase struct {
	Name     string
	Inputs   []interface{}
	Expected interface{}
}

// NewTestData TestDataをyamlファイル読み込みでなく、手動作成の場合に利用します。
func NewTestData(name string) *TestData {
	return &TestData{
		Name: name,
	}
}

func newTestCase(name string, inputs []interface{}, expected interface{}) *TestCase {
	return &TestCase{
		Name:     name,
		Inputs:   inputs,
		Expected: expected,
	}
}

// AddTestCase テストケースを追加します。
func (t *TestData) AddTestCase(name string, inputs []interface{}, expected interface{}) *TestData {
	testData := newTestCase(name, inputs, expected)
	t.Cases = append(t.Cases, testData)
	return t
}

// PrintTestCaseType テストケースの型を出力します。(デバッグ用)
func (t *TestData) PrintTestCaseType() {
	// type確認
	for _, testCase := range t.Cases {
		inputs := testCase.Inputs
		fmt.Printf("--- testCase %s ---\n", testCase.Name)
		for i, input := range inputs {
			fmt.Printf("input%d: %s\n", i, reflect.TypeOf(input).String())
		}
		fmt.Printf("expected: %s\n", reflect.TypeOf(testCase.Expected).String())
	}
}
