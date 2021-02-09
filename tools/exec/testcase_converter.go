package exec

import (
	"fmt"
	"reflect"
)

type testCaseConverter struct {
	valueConverters    map[reflect.Kind]func(value interface{}, isUserPtr bool, cacheKey string) (interface{}, error)
	fixedTypers        map[reflect.Kind]func(rValue reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Type, error)
	fixedValueCreaters map[reflect.Kind]func(rValue reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Value, error)
	cacheType          map[string]reflect.Type // 初回のTestCaseのTypeをキャッシュ
}

func newTestCaseConverter() *testCaseConverter {
	t := &testCaseConverter{
		valueConverters:    make(map[reflect.Kind]func(value interface{}, isUserPtr bool, cacheKey string) (interface{}, error)),
		fixedTypers:        make(map[reflect.Kind]func(rValue reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Type, error)),
		fixedValueCreaters: make(map[reflect.Kind]func(rValue reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Value, error)),
		cacheType:          make(map[string]reflect.Type),
	}
	t.registFunc()
	return t
}

func (t *testCaseConverter) registFunc() {
	t.valueConverters[reflect.Slice] = t.convertFixedTypeSlice
	t.valueConverters[reflect.Map] = t.convertFixedTypeMap

	t.fixedTypers[reflect.Slice] = t.getFixedSliceType
	t.fixedTypers[reflect.Map] = t.getFixedMapType

	t.fixedValueCreaters[reflect.Slice] = t.createFixedTypeSlice
	t.fixedValueCreaters[reflect.Map] = t.createFixedTypeMap
}

// convertValue valueConvertersに対応する型の値を変換
func (t *testCaseConverter) convertValue(value interface{}, isUserPtr bool, cacheKey string) (interface{}, error) {
	kind := reflect.ValueOf(value).Kind()
	converter := t.valueConverters[kind]
	if converter != nil {
		// slice、mapは変換
		convVal, err := converter(value, isUserPtr, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("type:%vの値変換に失敗しました。err:%v", kind, err)
		}
		return convVal, nil
	}

	if isUserPtr {
		err := fmt.Errorf("ポインタ型の指定はsliceまたはmapのvalueのみ対応しています。 value:%v", value)
		return nil, err
	}

	// その他の値は変換せず、前回TestCaseとTypeが同じであるかを確認
	err := t.cacheAndValidateType(reflect.ValueOf(value).Type(), cacheKey, 1)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// cacheAndValidateType valueの型情報をキャッシュが無ければキャッシュし、キャッシュがあればTypeを比較します。
func (t *testCaseConverter) cacheAndValidateType(valueType reflect.Type, cacheKey string, depth int) error {
	cacheDepthKey := cacheDepthKey(cacheKey, depth)
	cacheType, ok := t.cacheType[cacheDepthKey]
	if !ok {
		t.cacheType[cacheDepthKey] = valueType
	} else {
		if cacheType.String() != valueType.String() {
			return fmt.Errorf("型情報が他のTestCaseまたは他のCollection要素と異なります。 type:%s different type:%s cacheKey:%s", cacheType.String(), valueType.String(), cacheKey)
		}
	}
	return nil
}

func cacheDepthKey(cacheKey string, depth int) string {
	return fmt.Sprintf("%s-%d", cacheKey, depth)
}

// convertFixedTypeSlice
// []interface{}{inteface{}{Type}, inteface{}{Type} ...} の各要素がinterface型のsliceを
// []Type{val1, val2 ...} の型を固定したsliceに変換
func (t *testCaseConverter) convertFixedTypeSlice(iSlice interface{}, isUserPtr bool, cacheKey string) (interface{}, error) {
	// Slice Typeの取得
	rSlice := reflect.ValueOf(iSlice)
	_, err := t.getFixedSliceType(rSlice, isUserPtr, cacheKey, 1)
	if err != nil {
		return nil, err
	}

	// Sliceの生成
	newRSlice, err := t.createFixedTypeSlice(rSlice, isUserPtr, cacheKey, 1)
	if err != nil {
		return nil, err
	}
	return newRSlice.Interface(), nil
}

// getAndCacheSliceType typeを固定したスライスタイプを取得します。
func (t *testCaseConverter) getFixedSliceType(rSlice reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Type, error) {
	sliceLen := rSlice.Len()
	if sliceLen == 0 {
		// 空配列の場合
		cacheDepthKey := cacheDepthKey(cacheKey, depth)
		cacheType, ok := t.cacheType[cacheDepthKey]
		if ok {
			return cacheType, nil
		}
		return nil, nil
	}

	// typeの取得とnilの利用を確認
	isExistsNil := false
	var elemType reflect.Type
	for i := 0; i < sliceLen; i++ {
		rIndexVal := rSlice.Index(i)
		if rIndexVal.IsNil() {
			isExistsNil = true
			continue
		}

		var tmpElemType reflect.Type
		if t.isCollectionType(rIndexVal.Elem().Type()) {
			// collectionの場合、再帰的に取得
			typer := t.fixedTypers[rIndexVal.Elem().Type().Kind()]
			tmpCollElemType, err := typer(rIndexVal.Elem(), isUserPtr, cacheKey, depth+1)
			if err != nil {
				return nil, err
			}
			if tmpCollElemType == nil {
				continue
			}
			tmpElemType = tmpCollElemType
		} else {
			// プリミティブ要素の場合
			tmpElemType = rIndexVal.Elem().Type()
		}

		if elemType == nil {
			// 初回のelemTypeを登録
			elemType = tmpElemType
			continue
		}
		if elemType != tmpElemType {
			// 2回目以降のelemTypeが初回と不一致の場合エラー
			err := fmt.Errorf("index:%dの要素の型が不一致です。 elemType:%v differentType:%v slice:%v", i, elemType, tmpElemType, rSlice.Interface())
			return nil, err
		}
	}

	// 要素が全てnilだった場合
	if elemType == nil {
		return nil, nil
	}

	// slice or mapの場合
	isCollType := t.isCollectionType(elemType)

	// nilが存在するが、ポインタ利用で無い場合
	if !isCollType && isExistsNil && !isUserPtr {
		err := fmt.Errorf("nil値の存在するsliceにはポインタ型を適用して下さい。 slice:%v", rSlice.Interface())
		return nil, err
	}

	// スライス型の生成
	sliceElemType := elemType
	if !isCollType && isUserPtr {
		// ポインタ利用の場合、ポインタ型にする。
		sliceElemType = reflect.PtrTo(elemType)
	}
	sliceType := reflect.SliceOf(sliceElemType)
	// fmt.Printf("sliceType:%s cacheKey:%s depth:%d \n", sliceType.String(), cacheKey, depth)

	// 型のキャッシュと検証
	err := t.cacheAndValidateType(sliceType, cacheKey, depth)
	if err != nil {
		err := fmt.Errorf("sliceの型が他のTestCaseまたは他のCollection要素と異なります。 slice:%v error:%v", rSlice.Interface(), err)
		return nil, err
	}
	return sliceType, nil
}

func (t *testCaseConverter) isCollectionType(valueType reflect.Type) bool {
	_, ok := t.valueConverters[valueType.Kind()]
	return ok
}

// createFixedTypeSlice typeを固定したSliceを生成します。
func (t *testCaseConverter) createFixedTypeSlice(rSlice reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Value, error) {
	// この時点で型は確定しキャッシュされている
	cacheDepthKey := cacheDepthKey(cacheKey, depth)
	sliceType, ok := t.cacheType[cacheDepthKey]
	if !ok {
		err := fmt.Errorf("Slice型が判別できませんでした。 空Sliceを設定する場合は空で無いテストケースの後に設定して下さい。")
		return reflect.Zero(reflect.TypeOf(0)), err
	}

	sliceLen := rSlice.Len()
	if sliceLen == 0 {
		// 空配列の場合
		newRSlice := reflect.MakeSlice(sliceType, 0, 0)
		return newRSlice, nil
	}

	// Sliceの生成
	newRSlice := reflect.MakeSlice(sliceType, sliceLen, sliceLen)

	// 各要素をコピー
	sliceElemType := sliceType.Elem()
	for i := 0; i < sliceLen; i++ {
		targetRVal := newRSlice.Index(i)
		rVal := rSlice.Index(i)

		if rVal.IsNil() {
			// nil値の場合
			targetRVal.Set(reflect.Zero(sliceElemType))
		} else if t.isCollectionType(rVal.Elem().Type()) {
			// コレクション要素の場合、再帰的に取得
			creater := t.fixedValueCreaters[rVal.Elem().Type().Kind()]
			rSliceVal, err := creater(rVal.Elem(), isUserPtr, cacheKey, depth+1)
			if err != nil {
				return reflect.Zero(reflect.TypeOf(0)), err
			}
			targetRVal.Set(rSliceVal)
		} else if sliceElemType.Kind() == reflect.Ptr {
			// ポインタ型はint型とstring型のみ対応
			if rVal.Elem().Kind() == reflect.Int {
				intVal := rVal.Elem().Interface().(int)
				targetRVal.Set(reflect.ValueOf(&intVal))
			} else if rVal.Elem().Kind() == reflect.String {
				strVal := rVal.Elem().Interface().(string)
				targetRVal.Set(reflect.ValueOf(&strVal))
			} else {
				err := fmt.Errorf("未対応のnil許容型です。rVal:%v type:%v slice:%v", rVal, rVal.Kind(), rSlice.Interface())
				return reflect.Zero(reflect.TypeOf(0)), err
			}
		} else {
			// プリミティブ型の場合
			targetRVal.Set(rVal.Elem())
		}
	}
	return newRSlice, nil
}

// convertFixedTypeMap
// map[interface{}]interface{}{inteface{}{KeyType}:inteface{}{ValueType}, inteface{}{KeyType}:inteface{}{ValueType} ...} の各要素がinterface型のsliceを
// map[KeyType]ValueType{key1:value1, key2:value2 ...} の型を固定したmapに変換
func (t *testCaseConverter) convertFixedTypeMap(iMap interface{}, isUserPtr bool, cacheKey string) (interface{}, error) {
	// map Typeの取得
	rMap := reflect.ValueOf(iMap)
	_, err := t.getFixedMapType(rMap, isUserPtr, cacheKey, 1)
	if err != nil {
		return nil, err
	}

	// mapの生成
	newRMap, err := t.createFixedTypeMap(rMap, isUserPtr, cacheKey, 1)
	if err != nil {
		return nil, err
	}
	return newRMap.Interface(), nil
}

// getAndCacheSliceType Typeを固定したMap Typeを取得します。
func (t *testCaseConverter) getFixedMapType(rMap reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Type, error) {
	mapLen := rMap.Len()
	if mapLen == 0 {
		// 空配列の場合
		cacheDepthKey := cacheDepthKey(cacheKey, depth)
		cacheType, ok := t.cacheType[cacheDepthKey]
		if ok {
			return cacheType, nil
		}
		return nil, nil
	}

	// typeの取得とnilの利用を確認
	isExistsNil := false
	rKeys := rMap.MapKeys()
	var keyType reflect.Type
	var valueType reflect.Type
	for i := 0; i < mapLen; i++ {
		rKey := rKeys[i]
		rVal := rMap.MapIndex(rKey)

		// keyの評価
		if rKey.IsNil() {
			return nil, fmt.Errorf("値がnilのkeyには対応していません。map:%v", rMap.Interface())
		}

		if i == 0 {
			keyType = rKey.Elem().Type()
		}

		if keyType != rKey.Elem().Type() {
			err := fmt.Errorf("key:%vのkeyの型が不一致です。 key type:%v different key type:%v map:%v", rKey.Elem(), keyType, rKey.Elem().Type(), rMap.Interface())
			return nil, err
		}

		// valueの評価
		if rVal.IsNil() {
			isExistsNil = true
			continue
		}

		var tmpValueType reflect.Type
		if t.isCollectionType(rVal.Elem().Type()) {
			// collectionの場合、再帰的に取得
			typer := t.fixedTypers[rVal.Elem().Type().Kind()]
			tmpCollElemType, err := typer(rVal.Elem(), isUserPtr, cacheKey, depth+1)
			if err != nil {
				return nil, err
			}
			if tmpCollElemType == nil {
				continue
			}
			tmpValueType = tmpCollElemType
		} else {
			// プリミティブ要素の場合
			tmpValueType = rVal.Elem().Type()
		}

		if valueType == nil {
			valueType = tmpValueType
			continue
		}

		if valueType != tmpValueType {
			err := fmt.Errorf("key:%vのvalueの型が不一致です。 value type:%v different value type:%v map:%v", rKey.Elem(), valueType, tmpValueType, rMap.Interface())
			return nil, err
		}
	}

	// 要素が全てnilだった場合
	if valueType == nil {
		return nil, nil
	}

	// slice or mapの場合
	isCollType := t.isCollectionType(valueType)

	// nilが存在するが、ポインタ利用で無い場合
	if !isCollType && isExistsNil && !isUserPtr {
		err := fmt.Errorf("nil値の存在するmapにはポインタ型を適用して下さい。 mao:%v", rMap.Interface())
		return nil, err
	}

	// Map型の生成
	mapValueType := valueType
	if !isCollType && isUserPtr {
		// ポインタ利用の場合、ポインタ型にする。
		mapValueType = reflect.PtrTo(valueType)
	}
	mapType := reflect.MapOf(keyType, mapValueType)
	// fmt.Printf("mapType:%s cacheKey:%s depth:%d \n", mapType.String(), cacheKey, depth)

	// 型のキャッシュと検証
	err := t.cacheAndValidateType(mapType, cacheKey, depth)
	if err != nil {
		err := fmt.Errorf("mapの型が他のTestCaseまたは他のCollection要素と異なります。 map:%v error:%v", rMap.Interface(), err)
		return nil, err
	}
	return mapType, nil
}

// createFixedTypeSlice typeを固定したMapを生成します。
func (t *testCaseConverter) createFixedTypeMap(rMap reflect.Value, isUserPtr bool, cacheKey string, depth int) (reflect.Value, error) {
	// この時点で型は確定しキャッシュされている
	cacheDepthKey := cacheDepthKey(cacheKey, depth)
	mapType, ok := t.cacheType[cacheDepthKey]
	if !ok {
		err := fmt.Errorf("Map型が判別できませんでした。 空Mapを設定する場合は空で無いテストケースの後に設定して下さい。")
		return reflect.Zero(reflect.TypeOf(0)), err
	}

	mapLen := rMap.Len()
	if mapLen == 0 {
		// 空mapの場合
		newRMap := reflect.MakeMap(mapType)
		return newRMap, nil
	}

	// mapの生成
	newRMap := reflect.MakeMap(mapType)

	// 各要素をコピー
	rKeys := rMap.MapKeys()
	mapValueType := mapType.Elem()
	for i := 0; i < mapLen; i++ {
		rKey := rKeys[i]
		rVal := rMap.MapIndex(rKey)

		if rVal.IsNil() {
			// nilの場合
			newRMap.SetMapIndex(rKey.Elem(), reflect.Zero(mapValueType))
		} else if t.isCollectionType(rVal.Elem().Type()) {
			// コレクション要素の場合、再帰的に取得
			creater := t.fixedValueCreaters[rVal.Elem().Type().Kind()]
			rCollVal, err := creater(rVal.Elem(), isUserPtr, cacheKey, depth+1)
			if err != nil {
				return reflect.Zero(reflect.TypeOf(0)), err
			}
			newRMap.SetMapIndex(rKey.Elem(), rCollVal)
		} else if mapValueType.Kind() == reflect.Ptr {
			// ポインタのプリミティブ型はint型とstring型のみ対応
			if rVal.Elem().Kind() == reflect.Int {
				intVal := rVal.Elem().Interface().(int)
				newRMap.SetMapIndex(rKey.Elem(), reflect.ValueOf(&intVal))
			} else if rVal.Elem().Kind() == reflect.String {
				strVal := rVal.Elem().Interface().(string)
				newRMap.SetMapIndex(rKey.Elem(), reflect.ValueOf(&strVal))
			} else {
				err := fmt.Errorf("未対応のnil許容型です。rVal:%v type:%v map:%v", rVal, rVal.Kind(), rMap.Interface())
				return reflect.Zero(reflect.TypeOf(0)), err
			}
		} else {
			// プリミティブ型の場合
			newRMap.SetMapIndex(rKey.Elem(), rVal.Elem())
		}
	}
	return newRMap, nil
}
