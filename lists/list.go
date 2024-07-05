package lists

// InList [T comparable]
// @Author ricardo
// @Date 17:03 2024/7/5
// @Description:判断item是否在arr中存在
// @param arr 数组/切片
// @param item 元素
// @return bool
func InList[T comparable](arr []T, item T) bool {
	return InListWithFn(arr, item, func(a, b T) bool {
		return a == b
	})
}

// InListWithFn [T any]
//
//	@Author ricardo
//	@Date 17:03 2024/7/5
//	@Description: 通过传入自定义方法，判断item是否在arr中存在
//	@param arr 数组/切片
//	@param item 目标元素
//	@param equal 判断方法
//	@return bool
func InListWithFn[T any](arr []T, item T, equal func(a, b T) bool) bool {
	for _, a := range arr {
		if equal(a, item) {
			return true
		}
	}
	return false
}

// FindIndex [T comparable]
//
//	@Description: 查询元素在数组/切片中的索引
//	@param arr 数组/切片
//	@param item 元素
//	@return []int
func FindIndex[T comparable](arr []T, item T) []int {
	return FindIndexWithFn(arr, item, func(a, b T) bool {
		return a == b
	})
}

// FindIndexWithFn [T any]
//
//	@Description: 根据传入的自定义方法,查询元素在数组/切片中的索引
//	@param arr 数组/切片
//	@param item 元素
//	@param equal 判断方法
//	@return []int
func FindIndexWithFn[T any](arr []T, item T, equal func(a, b T) bool) []int {
	res := make([]int, 0)
	for i, a := range arr {
		if equal(a, item) {
			res = append(res, i)
		}
	}
	return res
}
