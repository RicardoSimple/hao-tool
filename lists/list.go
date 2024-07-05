package lists

import (
	"fmt"
)

// InList
//
//	@Description:判断item是否在arr中存在
//	@param arr 数组/切片
//	@param item 元素
//	@return bool
func InList(arr []interface{}, item interface{}) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func Hello() {
	fmt.Println("Hello World")
}
