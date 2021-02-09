package exec

import (
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestLoadTestCaseSlice(t *testing.T) {
	ymlFile := "testcase/load_slice_test.yml"
	loader := NewTestCaseLoader(ymlFile)
	testData, err := loader.Load()
	if err != nil {
		t.Error(err)
		return
	}

	// データが正しく読み込まれているか確認
	err = validateLoadTestData(testData, ymlFile)
	if err != nil {
		t.Error(err)
		return
	}

	// TestCaseのType出力
	testData.PrintTestCaseType()

	// Slice系型変換の確認
	for _, testCase := range testData.Cases {
		inputs := testCase.Inputs
		oksCount := len(inputs) + 1
		oks := make([]bool, oksCount, oksCount)
		_, oks[0] = inputs[0].(int)
		_, oks[1] = inputs[1].(string)
		_, oks[2] = inputs[2].(bool)
		_, oks[3] = inputs[3].([]int)
		_, oks[4] = inputs[4].([]*int)
		_, oks[5] = inputs[5].([]*int)
		_, oks[6] = inputs[6].([]bool)
		_, oks[7] = inputs[7].([]string)
		_, oks[8] = inputs[8].([]*string)
		_, oks[9] = inputs[9].([]*string)
		_, oks[10] = testData.Cases[0].Expected.([]*int)

		for i, ok := range oks {
			if !ok {
				t.Errorf("inputs[%d]が想定された形で形変換されていません。", i)
			}
		}
	}
}

func TestLoadTestCaseMap(t *testing.T) {
	ymlFile := "testcase/load_map_test.yml"
	loader := NewTestCaseLoader(ymlFile)
	testData, err := loader.Load()
	if err != nil {
		t.Error(err)
		return
	}

	// データが正しく読み込まれているか確認
	err = validateLoadTestData(testData, ymlFile)
	if err != nil {
		t.Error(err)
		return
	}

	// TestCaseのType出力
	testData.PrintTestCaseType()

	// Map系型変換の確認
	for _, testCase := range testData.Cases {
		inputs := testCase.Inputs
		oksCount := len(inputs) + 1
		oks := make([]bool, oksCount, oksCount)
		_, oks[0] = inputs[0].(map[string]int)
		_, oks[1] = inputs[1].(map[string]*int)
		_, oks[2] = inputs[2].(map[string]*int)
		_, oks[3] = inputs[3].(map[string]bool)
		_, oks[4] = inputs[4].(map[string]string)
		_, oks[5] = inputs[5].(map[string]*string)
		_, oks[6] = inputs[6].(map[int]int)
		_, oks[7] = inputs[7].(map[int]*int)
		_, oks[8] = inputs[8].(map[int]bool)
		_, oks[9] = inputs[9].(map[int]string)
		_, oks[10] = inputs[10].(map[int]*string)
		_, oks[11] = testCase.Expected.(map[string]int)

		for i, ok := range oks {
			if !ok {
				t.Errorf("inputs[%d]が想定された形で形変換されていません。", i)
			}
		}
	}
}

func TestLoadTestCaseComposit(t *testing.T) {

	// compositのテストケース作成
	// executorでcompositが必要なケースを追加
	// abebbcdefabdf で文字と出現位置にする a:[1,10], b:[2,3,4].. みたいな

	// TODO:mapとsliceの複合型の実装
	// reflect.kindをkeyとしたmap[kind]funcをgettype, createの2組作成し呼び出す。
	ymlFile := "testcase/load_composit_test.yml"
	loader := NewTestCaseLoader(ymlFile)
	testData, err := loader.Load()
	if err != nil {
		t.Error(err)
		return
	}

	// データが正しく読み込まれているか確認
	err = validateLoadTestData(testData, ymlFile)
	if err != nil {
		t.Error(err)
		return
	}

	// TestCaseのType出力
	testData.PrintTestCaseType()
}

// validateLoadTestData 読み込んだデータが正しいかを
// パース前のTestCaseとパース後のTestCaseを用意し
// 各項目をyml文字列にシリアライズして比較することで実施する。
func validateLoadTestData(testData *TestData, ymlFile string) error {
	// ymlFileからprimitiveなデータを読み込み
	pTestData, err := loadPrimitiveYmlData(ymlFile)
	if err != nil {
		return err
	}

	if len(testData.Cases) != len(pTestData.Cases) {
		return fmt.Errorf("テストケース数が一致しません。 data:%d primitiveData:%d", len(testData.Cases), len(pTestData.Cases))
	}

	for i := 0; i < len(testData.Cases); i++ {
		testCase := testData.Cases[i]
		pTestCase := pTestData.Cases[i]
		inputs := testCase.Inputs
		pInputs := pTestCase.Inputs

		if len(inputs) != len(pInputs) {
			return fmt.Errorf("テストケース:%dのInput数が一致しません。 data:%d primitiveData:%d", i, len(inputs), len(pInputs))
		}

		fmt.Printf("--- testCase %s ---\n", testCase.Name)

		for j := 0; j < len(inputs); j++ {
			input := inputs[j]
			pInput := pInputs[j]

			isEqual, err := compareTestCaseData(input, pInput, fmt.Sprintf("inputs[%d]:", j))
			if err != nil {
				return err
			}
			if !isEqual {
				return fmt.Errorf("テストケース:%d inputs[%d]のデータが一致しません。 input:%v primitiveInput:%v", i, j, input, pInput)
			}
		}

		isEqual, err := compareTestCaseData(testCase.Expected, pTestCase.Expected, "expected:")
		if err != nil {
			return err
		}
		if !isEqual {
			return fmt.Errorf("テストケース:%d expectedのデータが一致しません。 input:%v primitiveInput:%v", i, testCase.Expected, pTestCase.Expected)
		}
	}
	return nil
}

func loadPrimitiveYmlData(ymlFile string) (*TestData, error) {
	buf, err := ioutil.ReadFile(ymlFile)
	if err != nil {
		return nil, err
	}

	var testData TestData
	err = yaml.Unmarshal(buf, &testData)
	if err != nil {
		return nil, err
	}
	return &testData, nil
}

// compareTestCaseData dataとpDataをyml文字列に変換して比較
func compareTestCaseData(data interface{}, pData interface{}, title string) (bool, error) {
	// ポインタ利用キーワードがあった場合は比較対象データを抽出して比較
	compPdata, _ := isUsePointer(pData)

	dataStr, err := toMarshalYamlString(data)
	if err != nil {
		return false, err
	}
	fmt.Println(title)
	fmt.Println(dataStr)

	pDataStr, err := toMarshalYamlString(compPdata)
	if err != nil {
		return false, err
	}
	isEqualData := dataStr == pDataStr
	return isEqualData, nil
}

func toMarshalYamlString(data interface{}) (string, error) {
	bData, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bData), nil
}
