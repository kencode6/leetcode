package exec

/*
******* CodeExecutorに関して *******
■概要
テスト実行用の共通処理を提供します。
testcase.ymlに入力値と期待値を設定し、実行することで各テストケースの結果を評価します。

■テストケースファイルの設定
以下のような形式のtestcase.ymlにinputsに入力値(複数)、expectedに出力値を設定します。
--- testcase.yml ---
name: "Converter Test (multi ListNode)"
cases:
  - name: "case1"
    inputs:
      - [1, 2, 3]
      - 2
    expected: [2, 4, 6]
---------------------

Executorハンドラにテストする関数を以下のように定義しExecを実行すると、
Executorハンドラの戻り値とexpectedの値の比較が行われ、一致していればテスト合格とします。

--- logic_test.go --
NewCodeExecutor(t).
Executor(
	func(inputs []interface{}) interface{} {
		return testFuncXXX(inputs[0].(type1), inputs[1].(type2))
	},
).
Exec()
---------------------

■testcase.yml読み込みの仕様
通常のyml読み込みではinputs, expectedのコレクション型の各要素は
interface型としてデシリアライズされてしまい使い勝手が悪いので、
内部ロジックで要素型を解決し以下のように変換をかけている。

例)
[1, 2, 3]: []interface → []int
{"a": 1, "b": 2, "c": 3}: map[interface{}]interface{} → map[string]int

以下のような要素がSliceまたはMapのコンポジットデータも指定可能
例)
[[1,2],[1,2,3],[1,2,3,4]]: [][]interface{} → [][]int
{"a": [] ,"b": [1,2,3], "c": [1,2,3]}: map[interface{}][]interface{} → map[string][]int

また、ポインタ型はp*という指定をすることによって表現している。
ポインタ型の場合は要素にnull値も含めることができる。
例)
p*: [1, 2, 3, null]: []interface → []*int
p*: {"a": 1, "b": 2, "c": 3, "d": null}: map[interface{}]interface{} → map[string]*int

--- testcase.yml ---
name: "Type Convert Sample"
cases:
  - name: "case1"
    inputs:
      - 2 → int
      - [1, 2, 3] → []int
	  - {"a": 1, "b": 2, "c": 3} → map[string]int
      - p*: [1, 2, 3, null] → []*int
	  - p*: {"a": 1, "b": 2, "c": 3} → map[string]*int
	  - [[1,2],[1,2,3],[1,2,3,4]] → [][]int
	  - {"a": [] ,"b": [1,2,3], "c": [1,2,3]} → map[string][]int
    expected: [2, 4, 6] → []int
---------------------

■テストの方針
・inputs、expectedが通常の値(スカラー値またはmap, slice)の場合
TestCaseのinputs、expectedでslice、mapはコンポジット構造も含め、全て表現できるので
その場合はconverterは不要でありExecutorの設定のみでテスト可能である。

・inputs、expectedがオブジェクト(ListNodeまたはTreeNode)の場合
ListNode、TreeNodeのオブジェクトは表現できないのでConverterでの変換が必要になる。
	・入力がオブジェクトの場合
		- inputsをInputConverterで変換する。
		- CycleListNodeのように複数のinputs要素でオブジェクトを生成する必要がある場合、
		　 inputsをMultiInputConverterで変換する。

	・出力がオブジェクトの場合
		ExpectedConverterでexpectedをオブジェクトに変換する

	・その他のテスト結果検証が難しいケース
		Validatorで独自の検証ロジックを実装する
*/

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/kencode6/leetcode/tools/utils"
)

// DefaultTestCaseYmlFile テストケースymlファイルのデフォルト名称
const DefaultTestCaseYmlFile = "testcase.yml"

// CodeExecutor コード実行用の共通処理を行います。
type CodeExecutor struct {
	testing *testing.T

	ymlFile  string
	testData *TestData // テストデータ

	inputConverter      func(index int, input interface{}) interface{}                            // 入力値の変換処理
	multiInputConverter func(inputs []interface{}) interface{}                                    // 入力値の変換処理(input引数を複数)
	expectedConverter   func(expected interface{}) interface{}                                    // 期待値の変換処理
	executor            func(inputs []interface{}) interface{}                                    // 実行処理
	validator           func(output interface{}, expected interface{}, inputs []interface{}) bool // 出力値と期待値の検証処理
}

// NewCodeExecutor Executorを生成します。
func NewCodeExecutor(testing *testing.T) *CodeExecutor {
	return &CodeExecutor{
		testing:   testing,
		ymlFile:   DefaultTestCaseYmlFile,
		validator: basicValidator,
	}
}

// InputConverter inputデータに変換が必要な場合に変換ロジックを定義します。
func (c *CodeExecutor) InputConverter(inputConverter func(index int, input interface{}) interface{}) *CodeExecutor {
	c.inputConverter = inputConverter
	return c
}

// MultiInputConverter 複数のinputデータを用いて1つのinputデータに変換が必要な場合に変換ロジックを定義します。
func (c *CodeExecutor) MultiInputConverter(multiInputConverter func(inputs []interface{}) interface{}) *CodeExecutor {
	c.multiInputConverter = multiInputConverter
	return c
}

// Executor 実行ロジックを定義します。
func (c *CodeExecutor) Executor(executor func(inputs []interface{}) interface{}) *CodeExecutor {
	c.executor = executor
	return c
}

// ExpectedConverter expectedデータに変換が必要な場合に変換ロジックを定義します。
func (c *CodeExecutor) ExpectedConverter(expectedConverter func(expected interface{}) interface{}) *CodeExecutor {
	c.expectedConverter = expectedConverter
	return c
}

// Validator 出力値と検証値の比較をカスタマイズしたい場合に検証ロジックを定義します。
func (c *CodeExecutor) Validator(validator func(output interface{}, expected interface{}, inputs []interface{}) bool) *CodeExecutor {
	c.validator = validator
	return c
}

func basicValidator(output interface{}, expected interface{}, inputs []interface{}) bool {
	return reflect.DeepEqual(output, expected)
}

// TestData テストデータを手動でセットする場合に利用します。
func (c *CodeExecutor) TestData(testData *TestData) *CodeExecutor {
	c.testData = testData
	return c
}

// YmlFile ロードするymlファイルを変更したい場合に利用します。
func (c *CodeExecutor) YmlFile(ymlFile string) *CodeExecutor {
	c.ymlFile = ymlFile
	return c
}

// Exec 検証コードを実行します。
func (c *CodeExecutor) Exec() error {
	// testDataがロードされていない場合、testcase.ymlからデータをロード
	if c.testData == nil {
		err := c.loadTestData(c.ymlFile)
		if err != nil {
			return err
		}
	}

	// テストデータを実行
	log.Printf("##### %s Start #####", c.testData.Name)
	for _, tc := range c.testData.Cases {
		err := c.execTestCase(tc)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadTestData testcase.ymlを読み込んでTestDataをロードします。
func (c *CodeExecutor) loadTestData(ymlFilePath string) error {
	loader := NewTestCaseLoader(ymlFilePath)
	testData, err := loader.Load()
	if err != nil {
		return err
	}
	c.testData = testData
	return nil
}

func (c *CodeExecutor) execTestCase(tc *TestCase) error {
	// 前提条件の確認
	if c.executor == nil {
		return fmt.Errorf("実行メソッド定義用のExecutorでメソッド定義を行って下さい。")
	}

	log.Printf("------TestCase:%s------", tc.Name)
	// inputパラメータの変換
	var convInputs []interface{}
	if c.inputConverter != nil {
		// 1つのinput要素に対して変換
		for i, input := range tc.Inputs {
			convInput := c.inputConverter(i, input)
			convInputs = append(convInputs, convInput)
		}
	} else if c.multiInputConverter != nil {
		// 複数のinput要素を統合して変換
		convResult := c.multiInputConverter(tc.Inputs)
		if reflect.TypeOf(convResult).Kind() == reflect.Slice {
			convInputs = convResult.([]interface{})
		} else {
			convInputs = append(convInputs, convResult)
		}
	} else {
		convInputs = tc.Inputs
	}

	// 期待値の変換
	var convExpected interface{}
	if c.expectedConverter != nil {
		convExpected = c.expectedConverter(tc.Expected)
	} else {
		convExpected = tc.Expected
	}

	// inputs, expectedログ出力
	log.Println("input:")
	if len(convInputs) == 1 {
		log.Printf("%v", convertString(convInputs[0]))
	} else if len(convInputs) > 1 {
		for i, convInput := range convInputs {
			log.Printf("param%d:%s", i+1, convertString(convInput))
		}
	}
	log.Println("expected:")
	log.Printf("%s", convertString(convExpected))

	// コードの実行
	start := time.Now()
	log.Printf("******************** %s　exec start ********************", tc.Name)
	output := c.executor(convInputs)
	end := time.Now()
	elapsedTime := (end.Sub(start)).Microseconds()
	log.Printf("******************** %s exec finished elapsedTime: %d microsec ********************", tc.Name, elapsedTime)

	log.Println("output:")
	log.Printf("%s", convertString(output))

	// 結果の検証
	isValid := c.validator(output, convExpected, convInputs)
	if !isValid {
		errMsg := fmt.Sprintf("TestCase:%s NG got: %s want: %s", tc.Name, convertString(output), convertString(convExpected))
		log.Printf(errMsg)
		c.testing.Errorf(errMsg)
	} else {
		log.Printf("TestCase:%s OK", tc.Name)
	}
	return nil
}

func convertString(data interface{}) string {
	str := utils.ToJSONString(data)
	if len(str) == 0 {
		return "[Not Convertible Data]"
	}
	return str
}
