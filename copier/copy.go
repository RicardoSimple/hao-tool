package copier

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"
	"unsafe"
)

// DeepCopyWithJson
//
//	@Description: 深拷贝，按照json的相同tag解析，不支持私有类型拷贝
//	@param src 原始数据
//	@param dst 目标数据，指针
//	@return error
func DeepCopyWithJson(src, dst interface{}) error {
	marshal, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, dst)
	if err != nil {
		return err
	}
	return nil
}

// DeepCopy
//
//	@Description: 深拷贝 不拷贝私有类型
//	@param src 原始数据
//	@param dst 目标数据 指针
//	@return error
func DeepCopy(src, dst interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(src); err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}

// CopySame
//
//	@Description: 深拷贝,复制相同类型,支持复制私有类型,调用了unsafe包
//	@param src
//	@return interface{}
func CopySame(src interface{}) interface{} {
	srcVal := reflect.ValueOf(src)

	// 检查源是否为指针
	if srcVal.Kind() != reflect.Ptr {
		return nil
	}
	dstVal := reflect.New(srcVal.Elem().Type()).Elem()
	// 调用内部的深拷贝函数
	deepCopyValue(srcVal.Elem(), dstVal)
	return dstVal.Addr().Interface()
}

func deepCopyValue(srcVal, dstVal reflect.Value) {
	switch srcVal.Kind() {
	case reflect.Struct:
		for i := 0; i < srcVal.NumField(); i++ {
			fieldValue := srcVal.Field(i)
			dstField := dstVal.Field(i)

			// 使用 reflect 包的 UnexportedStructField 方法处理不可导出字段
			if !fieldValue.CanSet() {
				fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
			}

			if !dstField.CanSet() {
				dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
			}

			// 拷贝字段值
			deepCopyValue(fieldValue, dstField)
		}
	case reflect.Slice:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.MakeSlice(srcVal.Type(), srcVal.Len(), srcVal.Cap()))
			for i := 0; i < srcVal.Len(); i++ {
				deepCopyValue(srcVal.Index(i), dstVal.Index(i))
			}
		}
	case reflect.Map:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.MakeMapWithSize(srcVal.Type(), srcVal.Len()))
			for _, key := range srcVal.MapKeys() {
				newVal := reflect.New(srcVal.MapIndex(key).Type()).Elem()
				deepCopyValue(srcVal.MapIndex(key), newVal)
				dstVal.SetMapIndex(key, newVal)
			}
		}
	case reflect.Ptr:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.New(srcVal.Type().Elem()))
			deepCopyValue(srcVal.Elem(), dstVal.Elem())
		}
	default:
		dstVal.Set(srcVal)
	}
}
