package exec

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	usePointerKey           = "p*"
	cacheTypeKeyInputFormat = "input%d"
	cacheTypeKeyExpected    = "expected"
)

// TestCaseLoader TestCaseロード用
type TestCaseLoader struct {
	ymlFile   string
	converter *testCaseConverter
}

// NewTestCaseLoader TestCaseロード用
func NewTestCaseLoader(ymlFile string) *TestCaseLoader {

	converter := newTestCaseConverter()
	return &TestCaseLoader{
		ymlFile:   ymlFile,
		converter: converter,
	}
}

// Load ymlファイルを読み込みTestDataオブジェクトを返却します
func (t *TestCaseLoader) Load() (*TestData, error) {
	// ymlファイル読み込み
	testData, err := t.loadYmlFile(t.ymlFile)
	if err != nil {
		return nil, err
	}

	err = validateTestData(testData)
	if err != nil {
		return nil, err
	}

	for _, testCase := range testData.Cases {
		// testCaseのデータを加工
		err := t.convertTestCase(testCase)
		if err != nil {
			return nil, err
		}
	}
	return testData, nil
}

// loadYmlFile ymlファイル読み込み
func (t *TestCaseLoader) loadYmlFile(ymlFile string) (*TestData, error) {
	buf, err := ioutil.ReadFile(ymlFile)
	if err != nil {
		return nil, fmt.Errorf("ymlファイル読み込みに失敗しました。err:%v", err)
	}

	var testData TestData
	err = yaml.Unmarshal(buf, &testData)
	if err != nil {
		return nil, fmt.Errorf("ymlデータ変換に失敗しました。err:%v", err)
	}
	return &testData, nil
}

func validateTestData(testData *TestData) error {
	// testCaseの存在チェック
	if len(testData.Cases) == 0 {
		err := fmt.Errorf("testData:%sのTestCaseが存在しません", testData.Name)
		return err
	}

	// inputsの件数チェック
	var prevInputsCount int
	for _, testCase := range testData.Cases {
		if len(testCase.Inputs) == 0 {
			return fmt.Errorf("TestCase:%sのinputsを指定して下さい。", testData.Name)
		}
		if prevInputsCount != 0 && prevInputsCount != len(testCase.Inputs) {
			return fmt.Errorf("TestCase:%sのinputs数が異なります。 count%d diff count:%d", testData.Name, prevInputsCount, len(testCase.Inputs))
		}
		prevInputsCount = len(testCase.Inputs)

		if testCase.Expected == nil {
			return fmt.Errorf("TestCase:%sのexpectedを指定して下さい。", testData.Name)
		}

	}
	return nil
}

// convertTestCase テストケースのデータを加工
func (t *TestCaseLoader) convertTestCase(testCase *TestCase) error {
	// inputs変換
	for i := 0; i < len(testCase.Inputs); i++ {
		input := testCase.Inputs[i]
		input, isUserPtr := isUsePointer(input)

		cacheKey := fmt.Sprintf(cacheTypeKeyInputFormat, i)
		convInput, err := t.converter.convertValue(input, isUserPtr, cacheKey)
		if err != nil {
			return err
		}
		testCase.Inputs[i] = convInput
	}
	// expected変換
	expected, isUserPtr := isUsePointer(testCase.Expected)
	convExpected, err := t.converter.convertValue(expected, isUserPtr, cacheTypeKeyExpected)
	if err != nil {
		return err
	}
	testCase.Expected = convExpected
	return nil
}

func isUsePointer(input interface{}) (interface{}, bool) {
	castInput, ok := input.(map[interface{}]interface{})
	if !ok {
		return input, false
	}

	var key interface{} = usePointerKey
	ptrInput, ok2 := castInput[key]
	if !ok2 {
		// 通常のmap
		return input, false
	}
	return ptrInput, true
}
