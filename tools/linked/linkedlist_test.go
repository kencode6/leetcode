package linked

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSliceToListNode(t *testing.T) {
	nums := []int{5, 3, -1, 0, 6, -10}
	ln := SliceToListNode(nums)

	json, err := json.Marshal(ln)
	if err != nil {
		t.Errorf("json変換エラー :%v", err)
		return
	}
	lnStr := string(json)
	expected := `{"Val":5,"Next":{"Val":3,"Next":{"Val":-1,"Next":{"Val":0,"Next":{"Val":6,"Next":{"Val":-10,"Next":null}}}}}}`
	if lnStr != expected {
		t.Errorf("想定された結果と一致しません。 :%v", lnStr)
	}
}

func TestSliceToCycleListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	cycleLn, lastLn, jointLn := SliceToCycleListNode(nums, 3)

	lastVal := lastLn.Val
	jointVal := jointLn.Val
	expectedLastVal := nums[len(nums)-1]
	expectedJointVal := nums[3]
	if lastVal != expectedLastVal || jointVal != expectedJointVal {
		t.Errorf("想定された結果と一致しません。 lastVal:%d jointVal:%d", lastVal, jointVal)
		return
	}

	// jsonシリアライズはcycle構造に対応していないので、終端を切断してjson化して比較
	lastLn.Next = nil
	json, err := json.Marshal(cycleLn)
	if err != nil {
		t.Errorf("json変換エラー :%v", err)
		return
	}
	lnStr := string(json)
	expected := `{"Val":1,"Next":{"Val":2,"Next":{"Val":3,"Next":{"Val":4,"Next":{"Val":5,"Next":null}}}}}`
	if lnStr != expected {
		t.Errorf("想定された結果と一致しません。 :%v", lnStr)
	}
}

func TestToSlice(t *testing.T) {
	{
		nums := []int{-1, -2, -3, -4, -5, 0, 1, 2, 3, 4, 5}
		ln := SliceToListNode(nums)
		retNums := ln.ToSlice()
		if !reflect.DeepEqual(nums, retNums) {
			t.Errorf("想定された結果と一致しません。 :%v", retNums)
		}
	}
	{
		nums := []int{}
		ln := SliceToListNode(nums)
		retNums := ln.ToSlice()
		if !reflect.DeepEqual(nums, retNums) {
			t.Errorf("想定された結果と一致しません。 :%v", retNums)
		}
	}
}

func TestToString(t *testing.T) {
	{
		nums := []int{-1, -2, -3, -4, -5, 0, 1, 2, 3, 4, 5}
		ln := SliceToListNode(nums)
		retStr := ln.ToString()
		expectedStr := "[-1, -2, -3, -4, -5, 0, 1, 2, 3, 4, 5]"
		if retStr != expectedStr {
			t.Errorf("想定された結果と一致しません。 :%v", retStr)
		}
	}
	{
		nums := []int{}
		ln := SliceToListNode(nums)
		retStr := ln.ToString()
		expectedStr := "[]"
		if retStr != expectedStr {
			t.Errorf("想定された結果と一致しません。 :%v", retStr)
		}
	}
}

func TestAtIndex(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	ln := SliceToListNode(nums)
	retLn := ln.AtIndex(3)

	json, err := json.Marshal(retLn)
	if err != nil {
		t.Errorf("json変換エラー :%v", err)
		return
	}
	lnStr := string(json)
	expected := `{"Val":4,"Next":{"Val":5,"Next":null}}`
	if lnStr != expected {
		t.Errorf("想定された結果と一致しません。 :%v", lnStr)
	}
}
