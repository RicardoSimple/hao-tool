package now

import (
	"time"
)

// TodayOfDate
//
//	@Description: 返回当天0点
//	@return time.Time
func TodayOfDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// TodayOfDateFormat
//
//	@Description: 返回当天0点format后的字符串
//	@param layout
//	@return string
func TodayOfDateFormat(layout string) string {
	return TodayOfDate().Format(layout)
}
