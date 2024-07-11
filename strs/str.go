package strs

// EqualNoCase
//
//	@Description: 判断字符串是否相等,忽略大小写
//	@param str1
//	@param str2
//	@return bool
func EqualNoCase(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if (str1[i]-str2[i])%32 != 0 {
			return false
		}
	}
	return true
}
