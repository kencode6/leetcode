package utils

import (
	"encoding/json"

	"github.com/theothertomelliott/acyclic"
)

// ToPointerNums int sliceをポインタのint sliceに変換します。
// nilNumに指定した値はnilとして登録されます。
func ToPointerNums(nums []int, nilNum int) []*int {
	pNums := []*int{}
	for _, num := range nums {
		if num == nilNum {
			pNums = append(pNums, nil)
			continue
		}
		var pNum *int
		pNum = new(int)
		*pNum = num
		pNums = append(pNums, pNum)
	}
	return pNums
}

// ToJSONString オブジェクトをjson文字列に変換します。変換できない場合は空文字を返します
func ToJSONString(data interface{}) string {
	// サイクル要素であるかチェックする。
	err := acyclic.Check(data)
	if err != nil {
		return ""
	}

	tnb, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(tnb)
}
